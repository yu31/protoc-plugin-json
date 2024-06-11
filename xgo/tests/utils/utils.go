package utils

import (
	"encoding/base64"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func MustNewAny(src proto.Message) *anypb.Any {
	v, err := anypb.New(src)
	if err != nil {
		panic(err)
	}
	return v
}

func PointerString(s string) *string {
	return &s
}
func PointerInt32(i int32) *int32 {
	return &i
}
func PointerInt64(i int64) *int64 {
	return &i
}
func PointerUint32(i uint32) *uint32 {
	return &i
}
func PointerUint64(i uint64) *uint64 {
	return &i
}
func PointerFloat32(i float32) *float32 {
	return &i
}
func PointerFloat64(i float64) *float64 {
	return &i
}
func PointerBool(i bool) *bool {
	return &i
}

func ConvertJSONBytesToString(b string) string {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(dst, []byte(b))
	if err != nil {
		panic(err)
	}
	return string(dst[:n])
}
