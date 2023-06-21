package jsondecoder

import (
	"fmt"
	"strconv"
	"unsafe"
)

func (dec *Decoder) unquoteString(b []byte) (t []byte, err error) {
	var ok bool
	if t, ok = unquoteStdBytes(b); !ok {
		err = &SyntaxError{
			reason: "invalid string of JSON input",
			Offset: dec.offset,
		}
		return nil, err
	}
	return
}

func (dec *Decoder) convertToString(jsonKey string, b []byte) (vv string, err error) {
	if vv, err = bytesToString(b); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(b), jsonKey)
		return
	}
	return
}

func (dec *Decoder) convertToBool(jsonKey string, b []byte, unquote bool) (vv bool, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseBool(*(*string)(unsafe.Pointer(&b))); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(b), jsonKey)
		return
	}
	return
}
func (dec *Decoder) convertToInt32(jsonKey string, b []byte, unquote bool) (vv int32, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var i64 int64
	if i64, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(b), jsonKey)
		return
	}
	vv = int32(i64)
	return
}
func (dec *Decoder) convertToInt64(jsonKey string, b []byte, unquote bool) (vv int64, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 64); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(b), jsonKey)
		return
	}
	return
}
func (dec *Decoder) convertToUint32(jsonKey string, b []byte, unquote bool) (vv uint32, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var u64 uint64
	if u64, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 32); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(b), jsonKey)
		return
	}
	vv = uint32(u64)
	return
}
func (dec *Decoder) convertToUint64(jsonKey string, b []byte, unquote bool) (vv uint64, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(b), jsonKey)
		return
	}
	return
}

func (dec *Decoder) convertToFloat32(jsonKey string, b []byte, unquote bool) (vv float32, err error) {
	if unquote && floatIsNumber(b) {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var f64 float64
	if f64, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 32); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(b), jsonKey)
		return
	}
	vv = float32(f64)
	return
}
func (dec *Decoder) convertToFloat64(jsonKey string, b []byte, unquote bool) (vv float64, err error) {
	if unquote && floatIsNumber(b) {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	if vv, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 64); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(b), jsonKey)
		return
	}
	return
}
func (dec *Decoder) convertToEnum(jsonKey string, b []byte, unquote bool) (vv int32, err error) {
	if unquote {
		if b, err = dec.unquoteString(b); err != nil {
			return
		}
	}
	var i64 int64
	if i64, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type Enum", string(b), jsonKey)
		return
	}
	vv = int32(i64)
	return
}
