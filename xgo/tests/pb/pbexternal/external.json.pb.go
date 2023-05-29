// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/module/external.proto

package pbexternal

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message Embed in file tests/proto/module/external.proto
func (x *Embed) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed in file tests/proto/module/external.proto
func (x *Embed) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal.(*Embed) is nil")
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Message in file tests/proto/module/external.proto
func (x *Embed_Message) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Message in file tests/proto/module/external.proto
func (x *Embed_Message) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal.(*Embed_Message) is nil")
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Message1 in file tests/proto/module/external.proto
func (x *Message1) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Message1 in file tests/proto/module/external.proto
func (x *Message1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal.(*Message1) is nil")
	}
	decoder, err := jsondecoder.New(b)
	if err != nil {
		return err
	}
	if ok, _err := decoder.CheckJSONBegin(); _err != nil || ok {
		return _err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		if decoder.ReadObjectKeyBefore() { // before read object key
			break LOOP_OBJECT
		}
		// Read JSON key.
		jsonKey := decoder.ReadObjectKey()
		decoder.ReadObjectValueBefore() // Before read object value
		// match field with JSON key.
		switch {
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
			_ = decoder.ReadItem() // discard unknown field
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

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/module/external.proto
func (x *Message1_Embed1) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/module/external.proto
func (x *Message1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal.(*Message1_Embed1) is nil")
	}
	decoder, err := jsondecoder.New(b)
	if err != nil {
		return err
	}
	if ok, _err := decoder.CheckJSONBegin(); _err != nil || ok {
		return _err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		if decoder.ReadObjectKeyBefore() { // before read object key
			break LOOP_OBJECT
		}
		// Read JSON key.
		jsonKey := decoder.ReadObjectKey()
		decoder.ReadObjectValueBefore() // Before read object value
		// match field with JSON key.
		switch {
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
			_ = decoder.ReadItem() // discard unknown field
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

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/module/external.proto
func (x *Message1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/module/external.proto
func (x *Message1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal.(*Message1_Embed1_Embed2) is nil")
	}
	decoder, err := jsondecoder.New(b)
	if err != nil {
		return err
	}
	if ok, _err := decoder.CheckJSONBegin(); _err != nil || ok {
		return _err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		if decoder.ReadObjectKeyBefore() { // before read object key
			break LOOP_OBJECT
		}
		// Read JSON key.
		jsonKey := decoder.ReadObjectKey()
		decoder.ReadObjectValueBefore() // Before read object value
		// match field with JSON key.
		switch {
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
			_ = decoder.ReadItem() // discard unknown field
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
