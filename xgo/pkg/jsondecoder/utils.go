package jsondecoder

import (
	"bytes"
	"errors"
	"unsafe"
)

func bytesToStringUnsafe(b []byte) (string, error) {
	v, ok := unquoteJSONBytes(b)
	if !ok {
		return "", errors.New("invalid string format")
	}
	s := *(*string)(unsafe.Pointer(&v))
	return s, nil
}

func assertInterface(vv interface{}) {
	if vv == nil {
		panic("jsondecoder: the interface is nil")
	}
}

func floatIsNumber(b []byte) bool {
	if bytes.Equal(b, []byte("NaN")) ||
		bytes.Equal(b, []byte("+Inf")) ||
		bytes.Equal(b, []byte("-Inf")) {
		return false
	}
	return true
}
