// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/types/type_common.proto

package pbtypes

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message MessageCommon1 in file tests/proto/cases/types/type_common.proto
func (x *MessageCommon1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)
	encoder.AppendJSONKey("f_string3")
	encoder.AppendValueString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageCommon1 in file tests/proto/cases/types/type_common.proto
func (x *MessageCommon1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessageCommon1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		case jsonKey == "f_string3":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString3 = vv
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/types/type_common.proto
func (x *MessageCommon1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)
	encoder.AppendJSONKey("f_string3")
	encoder.AppendValueString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/types/type_common.proto
func (x *MessageCommon1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessageCommon1_Embed1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		case jsonKey == "f_string3":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString3 = vv
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/types/type_common.proto
func (x *MessageCommon1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)
	encoder.AppendJSONKey("f_string3")
	encoder.AppendValueString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/types/type_common.proto
func (x *MessageCommon1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessageCommon1_Embed1_Embed2) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		case jsonKey == "f_string3":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString3 = vv
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}
