package jsondecoder

// PhasePanicMsg is used as a panic message when we end up with something that
// shouldn't happen. It can indicate a bug in the JSON decoder, or that
// something is editing the data slice while the decoder executes.
const PhasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

type Decoder struct {
	data []byte
	off  int // next read offset in data
	scan scanner

	OpCode OpCode // last read result
}

func New(data []byte) (*Decoder, error) {
	d := &Decoder{
		data: data,
		off:  0,
	}
	err := checkValid(d.data, &d.scan)
	if err != nil {
		return nil, err
	}
	d.scan.reset()
	return d, nil
}

func (dec *Decoder) ScanError() error {
	return dec.scan.err
}

// ScanWhile scanWhile processes bytes in d.data[d.off:] until it
// receives a scan code not equal to op.
func (dec *Decoder) ScanWhile(op OpCode) {
	s, data, i := &dec.scan, dec.data, dec.off
	for i < len(data) {
		newOp := s.step(s, data[i])
		i++
		if newOp != op {
			dec.OpCode = newOp
			dec.off = i
			return
		}
	}

	dec.off = len(data) + 1 // mark processed EOF with len+1
	dec.OpCode = dec.scan.eof()
}

// Skip scans to the end of what was started.
func (dec *Decoder) Skip() {
	s, data, i := &dec.scan, dec.data, dec.off
	depth := len(s.parseState)
	for {
		op := s.step(s, data[i])
		i++
		if len(s.parseState) < depth {
			dec.off = i
			dec.OpCode = op
			return
		}
	}
}

// ScanNext processes the byte at d.data[d.off].
func (dec *Decoder) ScanNext() {
	if dec.off < len(dec.data) {
		dec.OpCode = dec.scan.step(&dec.scan, dec.data[dec.off])
		dec.off++
	} else {
		dec.OpCode = dec.scan.eof()
		dec.off = len(dec.data) + 1 // mark processed EOF with len+1
	}
}

// RescanLiteral is similar to scanWhile(ScanContinue), but it specialises the
// gojson case where we're decoding a literal. The decoder scans the input
// twice, once for syntax errors and to check the length of the value, and the
// second to perform the decoding.
//
// Only in the second step do we use decodeState to tokenize literals, so we
// know there aren't any syntax errors. We can take advantage of that knowledge,
// and scan a literal's bytes much more quickly.
func (dec *Decoder) RescanLiteral() {
	data, i := dec.data, dec.off
Switch:
	switch data[i-1] {
	case '"': // string
		for ; i < len(data); i++ {
			switch data[i] {
			case '\\':
				i++ // escaped char
			case '"':
				i++ // tokenize the closing quote too
				break Switch
			}
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-': // number
		for ; i < len(data); i++ {
			switch data[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
				'.', 'e', 'E', '+', '-':
			default:
				break Switch
			}
		}
	case 't': // true
		i += len("rue")
	case 'f': // false
		i += len("alse")
	case 'n': // null
		i += len("ull")
	}
	if i < len(data) {
		dec.OpCode = stateEndValue(&dec.scan, data[i])
	} else {
		dec.OpCode = ScanEnd
	}
	dec.off = i + 1
}
