package jsondecoder

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

func bytesToString(b []byte) (string, error) {
	v, ok := unquoteJSONBytes(b)
	if !ok {
		return "", fmt.Errorf("bytesToString: parsing %s: invalid syntax", strconv.Quote(string(b)))
	}
	s := string(v)
	return s, nil
}

func bytesToStringUnsafe(b []byte) (string, error) {
	v, ok := unquoteJSONBytes(b)
	if !ok {
		return "", fmt.Errorf("bytesToString: parsing %s: invalid syntax", strconv.Quote(string(b)))
	}
	s := *(*string)(unsafe.Pointer(&v))
	return s, nil
}

func parseEnumString(b []byte, em map[string]int32) (int32, error) {
	v1, err := bytesToStringUnsafe(b)
	if err != nil {
		return 0, err
	}
	if v2, ok := em[v1]; ok {
		return v2, nil
	}
	// Not found enum number in map. return the native number value.
	v3, err := strconv.ParseInt(v1, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v3), nil
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
