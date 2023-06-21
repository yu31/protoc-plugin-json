package jsonencoder

import (
	"encoding/base64"
)

// AppendLiteralNULL can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralNULL() {
	enc.appendCommaSeparator()
	enc.writeString("null")
}

// AppendLiteralString can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralString(v string) {
	enc.appendCommaSeparator()
	enc.appendString(v)
}

// AppendLiteralBytes can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralBytes(v []byte) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
		return
	}

	enc.writeByte('"')
	if len(v) != 0 {
		encodedLen := base64.StdEncoding.EncodedLen(len(v))
		// TODO: Improved the alloc logic.
		if encodedLen <= 1024 {
			// The encoded bytes are short enough to allocate for, and
			// Encoding.Encode is still cheaper.
			dst := make([]byte, encodedLen)
			base64.StdEncoding.Encode(dst, v)
			enc.writeBytes(dst)
		} else {
			// The encoded bytes are too long to cheaply allocate, and
			// Encoding.Encode is no longer noticeably cheaper.
			be := base64.NewEncoder(base64.StdEncoding, enc)
			if _, err := be.Write(v); err != nil {
				panic(err)
			}
			if err := be.Close(); err != nil {
				panic(err)
			}
		}
	}
	enc.writeByte('"')
}

// AppendLiteralInterface can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralInterface(v interface{}) error {
	enc.appendCommaSeparator()
	return enc.appendInterface(v)
}

// AppendLiteralBool can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralBool(v bool, quote bool) {
	enc.appendCommaSeparator()
	enc.appendBool(v, quote)
}

// AppendLiteralInt32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralInt32(v int32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(int64(v), quote)
}

// AppendLiteralInt64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralInt64(v int64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(v, quote)
}

// AppendLiteralUint32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralUint32(v uint32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(uint64(v), quote)
}

// AppendLiteralUint64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralUint64(v uint64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(v, quote)
}

// AppendLiteralFloat32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralFloat32(v float32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendFloat(float64(v), 32, quote)
}

// AppendLiteralFloat64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralFloat64(v float64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendFloat(v, 64, quote)
}
