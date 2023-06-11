package jsondecoder

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"unsafe"
)

func bytesToString(b []byte) (string, error) {
	v, ok := unquoteBytes(b)
	if !ok {
		return "", fmt.Errorf("bytesToString: parsing %s: invalid syntax", strconv.Quote(string(b)))
	}
	s := string(v)
	return s, nil
}
func bytesToStringUnsafe(b []byte) (string, error) {
	v, ok := unquoteBytes(b)
	if !ok {
		return "", fmt.Errorf("bytesToString: parsing %s: invalid syntax", strconv.Quote(string(b)))
	}
	s := *(*string)(unsafe.Pointer(&v))
	return s, nil
}

func bytesToBytes(b []byte) ([]byte, error) {
	t, ok := unquoteBytes(b)
	if !ok {
		return nil, fmt.Errorf("bytesToBytes: parsing %s: invalid syntax", strconv.Quote(string(b)))
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(t)))
	n, err := base64.StdEncoding.Decode(dst, t)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
func bytesToInt32(b []byte) (int32, error) {
	v, err := strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}
func bytesToInt64(b []byte) (int64, error) {
	return strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 64)
}
func bytesToUint32(b []byte) (uint32, error) {
	v, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}
func bytesToUint64(b []byte) (uint64, error) {
	return strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64)
}
func bytesToFloat32(b []byte) (float32, error) {
	v, err := strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}
func bytesToFloat64(b []byte) (float64, error) {
	return strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 64)
}
func bytesToBool(b []byte) (bool, error) {
	return strconv.ParseBool(*(*string)(unsafe.Pointer(&b)))
}
func parseEnumNumber(b []byte, em map[int32]string) (int32, error) {
	_ = em
	v1, err := bytesToInt32(b)
	if err != nil {
		return 0, err
	}
	return v1, nil
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
	v3, err := strconv.ParseInt(v1, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(v3), nil
}
