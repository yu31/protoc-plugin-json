package jsondecoder

import (
	"encoding/base64"
	"errors"
	"strconv"
	"unsafe"
)

func (dec *Decoder) unquoteBytes(b []byte) (t []byte, err error) {
	var ok bool
	if t, ok = unquoteStdBytes(b); !ok {
		err = errors.New("invalid bytes format")
		return nil, err
	}
	return
}
func (dec *Decoder) unquoteString(b []byte) (t []byte, err error) {
	var ok bool
	if t, ok = unquoteStdBytes(b); !ok {
		err = errors.New("invalid string format")
		return nil, err
	}
	return
}

func (dec *Decoder) convertToBytes(b []byte) (vv []byte, err error) {
	if b, err = dec.unquoteBytes(b); err != nil {
		return
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(dst, b)
	if err != nil {
		return nil, err
	}
	vv = dst[:n]
	return
}
func (dec *Decoder) convertToString(b []byte) (vv string, err error) {
	v, ok := unquoteJSONBytes(b)
	if !ok {
		err = errors.New("unable to convert to the type string")
		return
	}
	vv = string(v)
	return
}
func (dec *Decoder) convertToBool(unquote bool, b []byte) (vv bool, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseBool(*(*string)(unsafe.Pointer(&b))); err != nil {
		err = errors.New("unable to convert to the type bool")
		return
	}
	return
}
func (dec *Decoder) convertToInt32(unquote bool, b []byte) (vv int32, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var i64 int64
	if i64, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32); err != nil {
		err = errors.New("unable to convert to the type int32")
		return
	}
	vv = int32(i64)
	return
}
func (dec *Decoder) convertToInt64(unquote bool, b []byte) (vv int64, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 64); err != nil {
		err = errors.New("unable to convert to the type int64")
		return
	}
	return
}
func (dec *Decoder) convertToUint32(unquote bool, b []byte) (vv uint32, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var u64 uint64
	if u64, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 32); err != nil {
		err = errors.New("unable to convert to the type uint32")
		return
	}
	vv = uint32(u64)
	return
}
func (dec *Decoder) convertToUint64(unquote bool, b []byte) (vv uint64, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64); err != nil {
		err = errors.New("unable to convert to the type uint64")
		return
	}
	return
}

func (dec *Decoder) convertToFloat32(unquote bool, b []byte) (vv float32, err error) {
	if unquote && floatIsNumber(b) {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var f64 float64
	if f64, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 32); err != nil {
		err = errors.New("unable to convert to the type float32")
		return
	}
	vv = float32(f64)
	return
}
func (dec *Decoder) convertToFloat64(unquote bool, b []byte) (vv float64, err error) {
	if unquote && floatIsNumber(b) {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 64); err != nil {
		err = errors.New("unable to convert to the type float64")
		return
	}
	return
}
func (dec *Decoder) convertToEnum(unquote bool, b []byte) (vv int32, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var i64 int64
	if i64, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32); err != nil {
		err = errors.New("unable to convert to the type Enum")
		return
	}
	vv = int32(i64)
	return
}
