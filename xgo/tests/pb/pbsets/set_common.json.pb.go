// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/sets/set_common.proto

package pbsets

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message Message1 in file tests/proto/cases/sets/set_common.proto
func (x *Message1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)
	encoder.AppendObjectKey("f_string3")
	encoder.AppendLiteralString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Message1 in file tests/proto/cases/sets/set_common.proto
func (x *Message1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbsets.(*Message1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/sets/set_common.proto
func (x *Message1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)
	encoder.AppendObjectKey("f_string3")
	encoder.AppendLiteralString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/sets/set_common.proto
func (x *Message1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbsets.(*Message1_Embed1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/sets/set_common.proto
func (x *Message1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)
	encoder.AppendObjectKey("f_string3")
	encoder.AppendLiteralString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/sets/set_common.proto
func (x *Message1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbsets.(*Message1_Embed1_Embed2) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
