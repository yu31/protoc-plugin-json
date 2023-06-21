package jsonencoder

import (
	"encoding/json"
	"math"
	"strconv"
)

type Encoder struct {
	escapeHTML bool
	buf        []byte
}

// New return an Encoder.
// TODO: Add escapeHTML in func arguments.
func New(bufLen int) *Encoder {
	return &Encoder{
		escapeHTML: true,
		buf:        make([]byte, 0, bufLen),
	}
}

func (enc *Encoder) writeByte(v byte) {
	enc.buf = append(enc.buf, v)
}

func (enc *Encoder) writeBytes(v []byte) {
	enc.buf = append(enc.buf, v...)
}

func (enc *Encoder) writeString(v string) {
	enc.buf = append(enc.buf, v...)
}

// Impls interface io.Writer
func (enc *Encoder) Write(p []byte) (n int, err error) {
	enc.writeBytes(p)
	return len(p), nil
}

// Bytes Implements encoder.
func (enc *Encoder) Bytes() []byte {
	b := enc.buf
	enc.buf = nil
	return b
}

// AppendObjectBegin can be used for begin of JSON, map and struct.
func (enc *Encoder) AppendObjectBegin() {
	enc.writeByte('{')
}

// AppendObjectEnd can be used for end of JSON, map and struct.
func (enc *Encoder) AppendObjectEnd() {
	enc.writeByte('}')
}

// AppendArrayBegin can be used for begin of array and slice.
func (enc *Encoder) AppendArrayBegin() {
	enc.writeByte('[')
}

// AppendArrayEnd can be used for end of array and slice.
func (enc *Encoder) AppendArrayEnd() {
	enc.writeByte(']')
}

func (enc *Encoder) AppendObjectKey(k string) {
	enc.appendCommaSeparator()
	enc.appendString(k)
	enc.writeByte(':')
}

// Add elements separator.
func (enc *Encoder) appendCommaSeparator() {
	last := len(enc.buf) - 1
	if last < 0 {
		return
	}

	switch enc.buf[last] {
	case '{', '[', ':', ',':
		return
	default:
		enc.writeByte(',')
	}
}

func (enc *Encoder) appendBool(v bool, quote bool) {
	if quote {
		// As json.Marshal in standard library. The bool will be converted to string in map key.
		enc.buf = append(enc.buf, '"')
		enc.buf = strconv.AppendBool(enc.buf, v)
		enc.buf = append(enc.buf, '"')
	} else {
		enc.buf = strconv.AppendBool(enc.buf, v)
	}
}
func (enc *Encoder) appendInt(v int64, quote bool) {
	enc.appendCommaSeparator()
	if quote {
		// As json.Marshal in standard library. The integer will be converted to string in map key.
		enc.buf = append(enc.buf, '"')
		enc.buf = strconv.AppendInt(enc.buf, v, 10)
		enc.buf = append(enc.buf, '"')
	} else {
		enc.buf = strconv.AppendInt(enc.buf, v, 10)
	}
}
func (enc *Encoder) appendUint(v uint64, quote bool) {
	enc.appendCommaSeparator()
	if quote {
		// As json.Marshal in standard library. The integer will be converted to string in map key.
		enc.buf = append(enc.buf, '"')
		enc.buf = strconv.AppendUint(enc.buf, v, 10)
		enc.buf = append(enc.buf, '"')
	} else {
		enc.buf = strconv.AppendUint(enc.buf, v, 10)
	}
}
func (enc *Encoder) appendFloat(v float64, bitSize int, quote bool) {
	switch {
	case math.IsNaN(v):
		enc.appendString(`"NaN"`)
	case math.IsInf(v, 1):
		enc.appendString(`"+Inf"`)
	case math.IsInf(v, -1):
		enc.appendString(`"-Inf"`)
	default:
		if quote {
			// As json.Marshal in standard library. The float will be converted to string in map key.
			enc.buf = append(enc.buf, '"')
			enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, bitSize)
			enc.buf = append(enc.buf, '"')
		} else {
			enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, bitSize)
		}
	}
}

func (enc *Encoder) appendInterface(i interface{}) error {
	var err error
	var b []byte

	switch v := i.(type) {
	case nil:
		enc.writeString("null")
		return nil
	case json.Marshaler:
		b, err = v.MarshalJSON()
	default:
		b, err = json.Marshal(i)
	}
	if err != nil {
		return err
	}
	enc.writeBytes(b)
	return nil
}
