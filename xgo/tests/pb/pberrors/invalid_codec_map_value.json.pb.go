// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/errors/invalid_codec_map_value.proto

package pberrors

import (
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
)

// MarshalJSON implements interface json.Marshaler for proto message InvalidCodecMapValue in file tests/proto/cases/errors/invalid_codec_map_value.proto
func (x *InvalidCodecMapValue) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message InvalidCodecMapValue in file tests/proto/cases/errors/invalid_codec_map_value.proto
func (x *InvalidCodecMapValue) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pberrors", "InvalidCodecMapValue")
	}
	return nil
}
