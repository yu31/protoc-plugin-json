package jsonencoder

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"strconv"

	"github.com/golang/protobuf/ptypes/any"
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
	enc.appendItemSeparator()
	enc.appendString(k)
	enc.writeByte(':')
}

// AppendLiteralNULL can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralNULL() {
	enc.appendItemSeparator()
	enc.writeString("null")
}

// AppendLiteralString can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralString(v string) {
	enc.appendItemSeparator()
	enc.appendString(v)
}

// AppendLiteralBytes can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralBytes(v []byte) {
	enc.appendItemSeparator()
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

// AppendLiteralBool can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralBool(v bool) {
	enc.appendItemSeparator()
	enc.buf = strconv.AppendBool(enc.buf, v)
}

// AppendLiteralInt32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralInt32(v int32) {
	enc.appendItemSeparator()
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

// AppendLiteralInt64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralInt64(v int64) {
	enc.appendItemSeparator()
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}

// AppendLiteralUint32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralUint32(v uint32) {
	enc.appendItemSeparator()
	enc.buf = strconv.AppendUint(enc.buf, uint64(v), 10)
}

// AppendLiteralUint64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralUint64(v uint64) {
	enc.appendItemSeparator()
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
}

// AppendLiteralFloat32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralFloat32(v float32) {
	enc.appendItemSeparator()
	enc.appendFloat32(v)
}

// AppendLiteralFloat64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralFloat64(v float64) {
	enc.appendItemSeparator()
	enc.appendFloat64(v)
}

// AppendLiteralInterface can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendLiteralInterface(v interface{}) error {
	enc.appendItemSeparator()
	return enc.appendInterface(v)
}

func (enc *Encoder) AppendWKTAnyByProto(v *any.Any) error {
	b, err := pMarshal.Marshal(v)
	if err != nil {
		return err
	}
	enc.appendItemSeparator()
	enc.writeBytes(b)
	return nil
}

// AppendPointerString can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerString(v *string) {
	enc.appendItemSeparator()
	if v != nil {
		enc.appendString(*v)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerBool can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerBool(v *bool) {
	enc.appendItemSeparator()
	if v != nil {
		enc.buf = strconv.AppendBool(enc.buf, *v)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerInt32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerInt32(v *int32) {
	enc.appendItemSeparator()
	if v != nil {
		enc.buf = strconv.AppendInt(enc.buf, int64(*v), 10)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerInt64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerInt64(v *int64) {
	enc.appendItemSeparator()
	if v != nil {
		enc.buf = strconv.AppendInt(enc.buf, *v, 10)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerUint32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerUint32(v *uint32) {
	enc.appendItemSeparator()
	if v != nil {
		enc.buf = strconv.AppendUint(enc.buf, uint64(*v), 10)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerUint64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerUint64(v *uint64) {
	enc.appendItemSeparator()
	if v != nil {
		enc.buf = strconv.AppendUint(enc.buf, *v, 10)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerFloat32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerFloat32(v *float32) {
	enc.appendItemSeparator()
	if v != nil {
		enc.appendFloat32(*v)
	} else {
		enc.writeString("null")
	}
}

// AppendPointerFloat64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerFloat64(v *float64) {
	enc.appendItemSeparator()
	if v != nil {
		enc.appendFloat64(*v)
	} else {
		enc.writeString("null")
	}
}

func (enc *Encoder) AppendMapKeyString(v string) {
	enc.appendItemSeparator()
	enc.appendString(v)
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyInt32(v int32) {
	enc.appendItemSeparator()
	// As json.Marshal in standard library. The integer will be converted to string in map key.
	enc.buf = append(enc.buf, '"')
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
	enc.buf = append(enc.buf, '"')
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyInt64(v int64) {
	enc.appendItemSeparator()
	// As json.Marshal in standard library. The integer will be converted to string in map key.
	enc.buf = append(enc.buf, '"')
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
	enc.buf = append(enc.buf, '"')
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyUInt32(v uint32) {
	enc.appendItemSeparator()
	// As json.Marshal in standard library. The integer will be converted to string in map key.
	enc.buf = append(enc.buf, '"')
	enc.buf = strconv.AppendUint(enc.buf, uint64(v), 10)
	enc.buf = append(enc.buf, '"')
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyUInt64(v uint64) {
	enc.appendItemSeparator()
	// As json.Marshal in standard library. The integer will be converted to string in map key.
	enc.buf = append(enc.buf, '"')
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
	enc.buf = append(enc.buf, '"')
	enc.writeByte(':')
}

func (enc *Encoder) AppendMapKeyBool(v bool) {
	enc.appendItemSeparator()
	enc.buf = append(enc.buf, '"')
	enc.buf = strconv.AppendBool(enc.buf, v)
	enc.buf = append(enc.buf, '"')
	enc.writeByte(':')
}

// Add elements separator.
func (enc *Encoder) appendItemSeparator() {
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

func (enc *Encoder) appendFloat64(v float64) {
	switch {
	case math.IsNaN(v):
		enc.appendString(`"NaN"`)
	case math.IsInf(v, 1):
		enc.appendString(`"+Inf"`)
	case math.IsInf(v, -1):
		enc.appendString(`"-Inf"`)
	default:
		enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 64)
	}
}

func (enc *Encoder) appendFloat32(x float32) {
	v := float64(x)
	switch {
	case math.IsNaN(v):
		enc.appendString(`"NaN"`)
	case math.IsInf(v, 1):
		enc.appendString(`"+Inf"`)
	case math.IsInf(v, -1):
		enc.appendString(`"-Inf"`)
	default:
		enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 32)
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

func (enc *Encoder) writeByte(v byte) {
	enc.buf = append(enc.buf, v)
}

func (enc *Encoder) writeBytes(v []byte) {
	enc.buf = append(enc.buf, v...)
}

func (enc *Encoder) writeString(v string) {
	enc.buf = append(enc.buf, v...)
}
