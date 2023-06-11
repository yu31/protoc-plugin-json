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
func (dec *Decoder) scanWhile(op OpCode) error {
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
		return dec.scan.err
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

// Check if the object is NULL before read.
func (dec *Decoder) beforeReadObject() (isNULL bool, err error) {
	// If the object is a value of other key. The opCode maybe scanSkipSpace.
	// - e.g: {"m1"  :  {"k1":"v1"}}
	if dec.opCode == scanSkipSpace {
		// Skip the space character and advance to the next token.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}

	// Advance scanner to the next token that literal beginning(`null`) or object beginning character(`{`).
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
				reason: "unexpected object of JSON input, expecting object null",
				Offset: dec.offset,
			}
			return false, err
		}
		return true, nil
	case scanObjectBegin:
		// The token is object beginning character(`{`) - that represents the object is not NULL.
	default:
		// The object is not NULL. And the token must be object beginning character(`{`).
		// FIXME: use error instead of panics ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}
	return false, nil
}

// Check if it has more keys in the object before read.
func (dec *Decoder) beforeReadKey() (isEnd bool, err error) {
	/*
		The following cases will be matched:
		1. The token is object closing character(`}`).
			- that represents no more key/value in the object.
			- That happens when there is no any spaces between the previous valid character.
			- e.g: {"k1":"v1"}
		2. The token is space character and opCode is scanSkipSpace
			- that happens when has some spaces between previous valid character and next valid character.
			- And its next token is comma separators(`,`) or object closing character(`}`).
			- e.g: {"k1":"v1"   } or - e.g: {"k1":"v1"   ,"k2":"v2"}
		3. The token is object beginning character(`{`)
			- only happens when reading the first key in the object.
		    - And its next token is literal beginning or object closing character(`}`).
		4. The token is comma separators(`,`).
			- that represents has more key/value in the object.
	*/

	if dec.opCode == scanSkipSpace { // In cases 2.
		// Skip the space character and advance to the next token.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}

	if dec.opCode == scanObjectEnd { // In cases 1 or 2.
		// The token is object closing character(`}`) - That represents no more key/value in the object.
		return true, nil
	}

	// For process the cases 2, 3, 4
	// Advance scanner to the next token literal beginning or closing character(`}`).
	if err = dec.scanWhile(scanSkipSpace); err != nil {
		return
	}

	if dec.opCode == scanObjectEnd { // In cases 2 or 3
		// The token is object closing character(`}`) - That represents no more key/value in the object.
		return true, nil
	}

	// Assert the token - In cases  2 or 4.
	switch dec.opCode {
	case scanLiteralBegin, scanObjectValue:
	default:
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	/* The current token is the literal beginning of the double quotation marks(`"`). */

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

	/* The current token is colon separators(`:`) or space characters. */

	return key, err
}

// Read the value in object.
func (dec *Decoder) readObjectValue() (value []byte, err error) {
	if dec.opCode == scanSkipSpace {
		// In the cases - that represents has some spaces between key and colon separators(`:`). e.g {"k1"  :"v1"}
		// Skip the space characters and advance to the token colon separators(`:`).
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}
	// Assert the current op.
	if dec.opCode != scanObjectKey {
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	// Skip the colon separators(`:`) and advance the scanner to the next token.
	// Any spaces also will be skipped between colon separators(`:`) and next token.
	//
	// And the next token is literal beginning or object beginning or array beginning.
	if err = dec.scanWhile(scanSkipSpace); err != nil {
		return
	}

	// Assert the token.
	switch dec.opCode {
	case scanLiteralBegin, scanArrayBegin, scanObjectBegin:
	default:
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	// Read the valid characters as object value.
	if value, err = dec.readItem(); err != nil {
		return
	}

	/* The current token is space character or comma separators(`,`) or object closing character(`}`). */

	return value, nil
}

// Check if the array is NULL or is empty before read.
func (dec *Decoder) beforeReadArray() (isNULL bool, err error) {
	// - e.g: {"m1"  :  [1,2,3]}
	if dec.opCode == scanSkipSpace {
		// Skip the space character and advance to the next token.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}

	// Advance scanner to the next token literal beginning(`null`) or array beginning character(`[`).
	if err = dec.scanWhile(scanSkipSpace); err != nil {
		return
	}

	switch dec.opCode {
	case scanLiteralBegin:
		// The token is literal(`null`) - that represents the array is NULL.
		var item []byte
		if item, err = dec.readItem(); err != nil {
			return
		}
		if item[0] != 'n' {
			// TODO: Optimize the error output.
			err = &SyntaxError{
				reason: "unexpected object of JSON input, expecting array null",
				Offset: dec.offset,
			}
			return false, err
		}
		return true, nil
	case scanArrayBegin:
		// The token is array beginning character(`[`) - that represents the array is not NULL.
	default:
		// The array is not NULL. And the token must be array beginning character(`[`).
		// FIXME: use error instead of panics ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	/* The current token is the beginning of array element */

	return false, nil
}

// Check if it has more elements in the array before read.
func (dec *Decoder) beforeReadElem() (isEnd bool, err error) {
	/*
		The following cases will be matched:
		1. The token is array closing character(`]`).
			- that represents no more elements in the array.
			- That happens when there is no any spaces between the previous valid character.
			- e.g: [1,2]
		2. The token is space character and opCode is scanSkipSpace
			- that happens when has some spaces between previous valid character and next valid character.
			- And its next token is comma separators(`,`) or array closing character(`]`).
			- e.g: [1,2   ] or - e.g: [1,   2]
		3. The token is array beginning character(`[`)
			- only happens when reading the first key in the object.
		    - And its next token is literal beginning or array closing character(`]`).
		4. The token is comma separators(`,`).
			- that represents has more elements in the array.
	*/

	if dec.opCode == scanSkipSpace { // In cases 2.
		// Skip the space character and advance to the next token.
		if err = dec.scanWhile(scanSkipSpace); err != nil {
			return
		}
	}

	if dec.opCode == scanArrayEnd { // In cases 1 or 2.
		// The token is array closing character(`]`) - That represents no more elements in the array.
		return true, nil
	}

	// For process the cases 2, 3, 4
	// Advance scanner to the next token literal beginning or closing character(`]`).
	if err = dec.scanWhile(scanSkipSpace); err != nil {
		return
	}

	if dec.opCode == scanArrayEnd { // In cases 2 or 3
		// The token is array closing character(`]`) - That represents no more elements in the array.
		return true, nil
	}

	// Assert the token - In cases  2 or 4.
	switch dec.opCode {
	case scanLiteralBegin, scanArrayElem, scanObjectBegin, scanArrayBegin:
	default:
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	/* The current token is the literal beginning of array element */

	return false, nil
}

// Read the array element.
func (dec *Decoder) readArrayElem() (elem []byte, err error) {
	// Assert the opCode.
	switch dec.opCode {
	case scanLiteralBegin, scanArrayBegin, scanObjectBegin:
	default:
		// FIXME: return error instead of panic ?
		panic(fmt.Sprintf("%s, opCode: %d", PhasePanicMsg, dec.opCode))
	}

	// Read the valid literal characters as array element.
	if elem, err = dec.readItem(); err != nil {
		return
	}

	/*
		The current token is space character or comma separators(`,`) or array closing character(`]`)
	*/
	return
}
