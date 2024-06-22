// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/inline/private.proto

package pbinline

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message MessageEmpty1 in file tests/proto/cases/inline/private.proto
func (x *MessageEmpty1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageEmpty1 in file tests/proto/cases/inline/private.proto
func (x *MessageEmpty1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbinline.(*MessageEmpty1) is nil")
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message MessageIgnoreFields in file tests/proto/cases/inline/private.proto
func (x *MessageIgnoreFields) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageIgnoreFields in file tests/proto/cases/inline/private.proto
func (x *MessageIgnoreFields) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbinline.(*MessageIgnoreFields) is nil")
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message MessageIgnoreByRef in file tests/proto/cases/inline/private.proto
func (x *MessageIgnoreByRef) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(64)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "ignore_by_ref_string1", x.IgnoreByRefString1, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageIgnoreByRef in file tests/proto/cases/inline/private.proto
func (x *MessageIgnoreByRef) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbinline.(*MessageIgnoreByRef) is nil")
	}
	var (
		err    error
		isNULL bool
		dec    *jsondecoder.Decoder
	)
	if dec, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = dec.BeforeScanJSON(); err != nil || isNULL {
		return err
	}
LOOP_SCAN:
	for { // Loop to read the JSON objects
		var (
			jsonKey string
			isEnd   bool
		)

		if isEnd, err = dec.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}

		if jsonKey, err = dec.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "ignore_by_ref_string1":
			if x.IgnoreByRefString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		default:
			if err = dec.DiscardValue(); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
