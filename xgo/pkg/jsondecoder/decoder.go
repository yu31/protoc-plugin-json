package jsondecoder

import "fmt"

// PhasePanicMsg is used as a panic message when we end up with something that
// shouldn't happen. It can indicate a bug in the JSON decoder, or that
// something is editing the data slice while the decoder executes.
const PhasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

type Decoder struct {
	data   []byte
	offset int // The offset of read.

	scan   scanner
	opCode OpCode // last read result
}

func New(data []byte) (*Decoder, error) {
	d := &Decoder{
		data:   data,
		offset: 0,
		opCode: scanContinue,
		scan: scanner{
			step:       stateBeginValue,
			endTop:     false,
			parseState: nil,
			err:        nil,
			bytes:      0,
		},
	}
	return d, nil
}

// ReadItem read current valid value.
func (dec *Decoder) readItem() ([]byte, error) {
	var err error
	start := dec.offset - 1 // start index.
	switch dec.opCode {
	case scanLiteralBegin:
		if err = dec.scanWhile(scanContinue); err != nil {
			return nil, err
		}
	case scanArrayBegin, scanObjectBegin:
		if err = dec.skip(); err != nil {
			return nil, err
		}
		if err = dec.scanNext(); err != nil {
			return nil, err
		}
	default:
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}
	end := dec.offset - 1 // end index.
	bb := dec.data[start:end]
	return bb, nil
}

// scanWhile processes bytes in d.data[d.offset:] until it
// receives a scan code not equal to op.
func (dec *Decoder) scanWhile(op OpCode) (err error) {
	var ok bool
	s, data, i := &dec.scan, dec.data, dec.offset
	for i < len(data) {
		s.bytes++
		newOp := s.step(s, data[i])
		i++
		if newOp != op {
			dec.opCode = newOp
			dec.offset = i
			ok = true
			break
		}
	}
	if !ok {
		dec.offset = len(data) + 1 // mark processed EOF with len+1
		dec.opCode = dec.scan.eof()
	}
	if dec.opCode == scanError {
		err = dec.scan.err
		return
	}
	return nil
}

// skip scans to the end of what was started.
func (dec *Decoder) skip() (err error) {
	s, data, i := &dec.scan, dec.data, dec.offset
	depth := len(s.parseState)
	for {
		s.bytes++
		op := s.step(s, data[i])
		i++
		if len(s.parseState) < depth {
			dec.offset = i
			dec.opCode = op
			break
		}
	}
	if dec.opCode == scanError {
		err = dec.scan.err
		return
	}
	return err
}

// scanNext processes the byte at d.data[d.offset].
// notice: used to after read object or array if them not NULL.
func (dec *Decoder) scanNext() (err error) {
	if dec.offset < len(dec.data) {
		dec.scan.bytes++
		dec.opCode = dec.scan.step(&dec.scan, dec.data[dec.offset])
		dec.offset++
	} else {
		dec.opCode = dec.scan.eof()
		dec.offset = len(dec.data) + 1 // mark processed EOF with len+1
	}
	if dec.opCode == scanError {
		err = dec.scan.err
		return
	}
	return nil
}

// Check if the object or the array or the map is NULL before reading.
func (dec *Decoder) beforeReadObject() (isNULL bool, err error) {
	// Advance scanner to the next token.
	// The next toke is literal beginning(`null`) or object beginning character(`{`) or array beginning character(`[`).
	if err = dec.scanWhile(scanSkipSpace); err != nil {
		return
	}

	switch dec.opCode {
	case scanLiteralBegin:
		// The token is literal(`null`) - that represents the object is NULL.
		var item []byte
		if item, err = dec.readItem(); err != nil {
			return
		}
		if item[0] != 'n' {
			// TODO: Optimize the error output.
			err = &SyntaxError{
				reason: fmt.Sprintf("unexpected of JSON input <%s>, expecting literanl null", string(item)),
				Offset: dec.offset,
			}
			return false, err
		}
		return true, nil
	case scanObjectBegin, scanArrayBegin:
		// The token is object beginning character(`{`) or array beginning character(`[`).
		// And it represents the object or array is not NULL.
	default:
		// FIXME: use error instead of panics ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	/* The current token is the beginning of literal */

	return false, nil
}

// Check if it has more items in the object or array before reading.
func (dec *Decoder) beforeReadNext() (isEnd bool, err error) {
	/*
		The following cases will be matched:
		1. The token is object closing character(`}`) or array closing character(`]`).
			- That represents there no more items in the object or array.
			- That happens when there is no any spaces between the previous valid character.
			- e.g: {"k1":"v1"} or [1,2]
		2. The token is space character and opCode is scanSkipSpace
			- That happens when has some spaces between previous valid character and next valid character.
			- And its next token is comma separators(`,`) or object closing character(`}`) or array closing character(`]`).
			- e.g: {"k1":"v1"   } or {"k1":"v1"   ,"k2":"v2"} or [1,2   ] or [1,   2].
		3. The token is object beginning character(`{`) or array beginning character(`]`).
			- Only happens when reading the first key in the object or first element in the arrray.
		    - And its next token is literal beginning or object closing character(`}`) or array closing character(`]`).
		4. The token is comma separators(`,`).
			- That represents there has more items in the object or array.
	*/

	if dec.opCode == scanSkipSpace { // In cases 2.
		// Skip the space character and advance to the next token.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}

	switch dec.opCode {
	case scanObjectEnd, scanArrayEnd: // In cases 1 or 2.
		// The token is object closing character(`}`) or array closing character(`]`).
		// And that represents there no more items in the object or array.
		return true, nil
	}

	// For process the cases 2, 3, 4
	// Advance scanner to the next token literal beginning or object closing character(`}`) or array closing character(`]`).
	// The operation will skip the comma separators(`,`) is next token is literal.
	if err = dec.scanWhile(scanSkipSpace); err != nil {
		return
	}

	switch dec.opCode {
	case scanObjectEnd, scanArrayEnd: // In cases 2 or 3.
		// The token is object closing character(`}`) or array closing character(`]`).
		// And that represents there no more items in the object or array.
		return true, nil
	}

	// Assert the token - In cases  2 or 4.
	switch dec.opCode {
	//case scanLiteralBegin, scanObjectBegin, scanObjectValue, scanArrayElem, scanArrayBegin:
	case scanLiteralBegin, scanObjectBegin:
	default:
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	/* The current token is the literal beginning or object beginning. */

	return false, nil
}

// Read the key in an object.
func (dec *Decoder) readObjectKey() (key []byte, err error) {
	// Assert the current token.
	if dec.opCode != scanLiteralBegin {
		// FIXME: Use error instead of panic.
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	// The token is literal beginning, Read the valid literal characters as object key.
	if key, err = dec.readItem(); err != nil {
		return
	}

	/* The current token is space characters or colon separators(`:`). */

	if dec.opCode == scanSkipSpace {
		// In the cases - that represents has some spaces between key and colon separators(`:`). e.g {"k1"  :"v1"}
		// Skip the space characters and advance to the token colon separators(`:`).
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return nil, err
		}
	}
	// Assert the current op.
	if dec.opCode != scanObjectKey {
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	/* The next token is colon separators(`:`)  */

	return key, err
}

// Read the next item as object value or array element.
func (dec *Decoder) readLiteralValue() (value []byte, err error) {
	if dec.opCode == scanObjectKey {
		// The opCode is scanObjectKey represents the last operation is readObjectKey.
		//
		// Skip the colon separators(`:`) and advance the scanner to the next token.
		// Any spaces also will be skipped between colon separators(`:`) and next token.
		//
		// And the next token is literal beginning or object beginning or array beginning.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}

	// Read the valid characters as object value or array element.
	if value, err = dec.readItem(); err != nil {
		return
	}

	/*
		The current token is space character or comma separators(`,`) or object closing character(`}`) or array closing character(`]`).
	*/
	return value, nil
}

func (dec *Decoder) nextLiteralIsNULL() (isNULL bool, err error) {
	if dec.opCode == scanObjectKey {
		// The opCode is scanObjectKey represents the last operation is readObjectKey.
		//
		// Skip the colon separators(`:`) and advance the scanner to the next token.
		// Any spaces also will be skipped between colon separators(`:`) and next token.
		//
		// And the next token is literal beginning or object beginning or array beginning.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}
	if dec.offset >= len(dec.data) {
		return
	}
	if b := dec.data[dec.offset-1]; b == 'n' {
		return true, nil
	}
	return
}
