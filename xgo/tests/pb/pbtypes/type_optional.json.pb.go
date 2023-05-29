// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/types/type_optional.proto

package pbtypes

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	pbexternal "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message MessageOptional1 in file tests/proto/cases/types/type_optional.proto
func (x *MessageOptional1) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message MessageOptional1 in file tests/proto/cases/types/type_optional.proto
func (x *MessageOptional1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessageOptional1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/types/type_optional.proto
func (x *MessageOptional1_Embed1) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/types/type_optional.proto
func (x *MessageOptional1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessageOptional1_Embed1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/types/type_optional.proto
func (x *MessageOptional1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/types/type_optional.proto
func (x *MessageOptional1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessageOptional1_Embed1_Embed2) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message TypeOptional1 in file tests/proto/cases/types/type_optional.proto
func (x *TypeOptional1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(868)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendPointerString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendPointerString(x.FString2)
	encoder.AppendJSONKey("f_int32")
	encoder.AppendPointerInt32(x.FInt32)
	encoder.AppendJSONKey("f_int64")
	encoder.AppendPointerInt64(x.FInt64)
	encoder.AppendJSONKey("f_uint32")
	encoder.AppendPointerUint32(x.FUint32)
	encoder.AppendJSONKey("f_uint64")
	encoder.AppendPointerUint64(x.FUint64)
	encoder.AppendJSONKey("f_sint32")
	encoder.AppendPointerInt32(x.FSint32)
	encoder.AppendJSONKey("f_sint64")
	encoder.AppendPointerInt64(x.FSint64)
	encoder.AppendJSONKey("f_sfixed32")
	encoder.AppendPointerInt32(x.FSfixed32)
	encoder.AppendJSONKey("f_sfixed64")
	encoder.AppendPointerInt64(x.FSfixed64)
	encoder.AppendJSONKey("f_fixed32")
	encoder.AppendPointerUint32(x.FFixed32)
	encoder.AppendJSONKey("f_fixed64")
	encoder.AppendPointerUint64(x.FFixed64)
	encoder.AppendJSONKey("f_float")
	encoder.AppendPointerFloat32(x.FFloat)
	encoder.AppendJSONKey("f_double")
	encoder.AppendPointerFloat64(x.FDouble)
	encoder.AppendJSONKey("f_bool1")
	encoder.AppendPointerBool(x.FBool1)
	encoder.AppendJSONKey("f_bytes1")
	encoder.AppendValueBytes(x.FBytes1)
	encoder.AppendJSONKey("f_enum1")
	if x.FEnum1 != nil {
		encoder.AppendValueInt32(int32(x.FEnum1.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_enum2")
	if x.FEnum2 != nil {
		encoder.AppendValueInt32(int32(x.FEnum2.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_enum3")
	if x.FEnum3 != nil {
		encoder.AppendValueInt32(int32(x.FEnum3.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_enum4")
	if x.FEnum4 != nil {
		encoder.AppendValueInt32(int32(x.FEnum4.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_enum5")
	if x.FEnum5 != nil {
		encoder.AppendValueInt32(int32(x.FEnum5.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_enum6")
	if x.FEnum6 != nil {
		encoder.AppendValueInt32(int32(x.FEnum6.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_duration1")
	err = encoder.AppendValueInterface(x.FDuration1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_duration2")
	err = encoder.AppendValueInterface(x.FDuration2)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_timestamp1")
	err = encoder.AppendValueInterface(x.FTimestamp1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_timestamp2")
	err = encoder.AppendValueInterface(x.FTimestamp2)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_any1")
	err = encoder.AppendValueInterface(x.FAny1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_any2")
	err = encoder.AppendValueInterface(x.FAny2)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message1")
	err = encoder.AppendValueInterface(x.FMessage1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message2")
	err = encoder.AppendValueInterface(x.FMessage2)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message3")
	err = encoder.AppendValueInterface(x.FMessage3)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message4")
	err = encoder.AppendValueInterface(x.FMessage4)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message5")
	err = encoder.AppendValueInterface(x.FMessage5)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message6")
	err = encoder.AppendValueInterface(x.FMessage6)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message7")
	err = encoder.AppendValueInterface(x.FMessage7)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message8")
	err = encoder.AppendValueInterface(x.FMessage8)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message9")
	err = encoder.AppendValueInterface(x.FMessage9)
	if err != nil {
		return nil, err
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOptional1 in file tests/proto/cases/types/type_optional.proto
func (x *TypeOptional1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*TypeOptional1) is nil")
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
			vv, _err := decoder.ReadPointerString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadPointerString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		case jsonKey == "f_int32":
			vv, _err := decoder.ReadPointerInt32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FInt32 = vv
		case jsonKey == "f_int64":
			vv, _err := decoder.ReadPointerInt64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FInt64 = vv
		case jsonKey == "f_uint32":
			vv, _err := decoder.ReadPointerUint32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FUint32 = vv
		case jsonKey == "f_uint64":
			vv, _err := decoder.ReadPointerUint64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FUint64 = vv
		case jsonKey == "f_sint32":
			vv, _err := decoder.ReadPointerInt32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSint32 = vv
		case jsonKey == "f_sint64":
			vv, _err := decoder.ReadPointerInt64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSint64 = vv
		case jsonKey == "f_sfixed32":
			vv, _err := decoder.ReadPointerInt32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSfixed32 = vv
		case jsonKey == "f_sfixed64":
			vv, _err := decoder.ReadPointerInt64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSfixed64 = vv
		case jsonKey == "f_fixed32":
			vv, _err := decoder.ReadPointerUint32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FFixed32 = vv
		case jsonKey == "f_fixed64":
			vv, _err := decoder.ReadPointerUint64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FFixed64 = vv
		case jsonKey == "f_float":
			vv, _err := decoder.ReadPointerFloat32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FFloat = vv
		case jsonKey == "f_double":
			vv, _err := decoder.ReadPointerFloat64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FDouble = vv
		case jsonKey == "f_bool1":
			vv, _err := decoder.ReadPointerBool(jsonKey)
			if _err != nil {
				return _err
			}
			x.FBool1 = vv
		case jsonKey == "f_bytes1":
			vv, _err := decoder.ReadValueBytes(jsonKey)
			if _err != nil {
				return _err
			}
			x.FBytes1 = vv
		case jsonKey == "f_enum1":
			var vv *EnumOptional1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, EnumOptional1_name)
			if v1 != nil {
				_vv := EnumOptional1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum1 = vv
		case jsonKey == "f_enum2":
			var vv *pbexternal.Enum1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, pbexternal.Enum1_name)
			if v1 != nil {
				_vv := pbexternal.Enum1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum2 = vv
		case jsonKey == "f_enum3":
			var vv *pbexternal.Embed_Enum1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, pbexternal.Embed_Enum1_name)
			if v1 != nil {
				_vv := pbexternal.Embed_Enum1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum3 = vv
		case jsonKey == "f_enum4":
			var vv *pbexternal.Embed_Message_Enum1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, pbexternal.Embed_Message_Enum1_name)
			if v1 != nil {
				_vv := pbexternal.Embed_Message_Enum1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum4 = vv
		case jsonKey == "f_enum5":
			var vv *EnumCommon1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, EnumCommon1_name)
			if v1 != nil {
				_vv := EnumCommon1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum5 = vv
		case jsonKey == "f_enum6":
			var vv *MessageCommon1_Enum1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, MessageCommon1_Enum1_name)
			if v1 != nil {
				_vv := MessageCommon1_Enum1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum6 = vv
		case jsonKey == "f_duration1":
			var vv *durationpb.Duration
			if x.FDuration1 == nil {
				vv = new(durationpb.Duration)
			} else {
				vv = x.FDuration1
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FDuration1 = vv
		case jsonKey == "f_duration2":
			var vv *durationpb.Duration
			if x.FDuration2 == nil {
				vv = new(durationpb.Duration)
			} else {
				vv = x.FDuration2
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FDuration2 = vv
		case jsonKey == "f_timestamp1":
			var vv *timestamppb.Timestamp
			if x.FTimestamp1 == nil {
				vv = new(timestamppb.Timestamp)
			} else {
				vv = x.FTimestamp1
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FTimestamp1 = vv
		case jsonKey == "f_timestamp2":
			var vv *timestamppb.Timestamp
			if x.FTimestamp2 == nil {
				vv = new(timestamppb.Timestamp)
			} else {
				vv = x.FTimestamp2
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FTimestamp2 = vv
		case jsonKey == "f_any1":
			var vv *anypb.Any
			if x.FAny1 == nil {
				vv = new(anypb.Any)
			} else {
				vv = x.FAny1
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FAny1 = vv
		case jsonKey == "f_any2":
			var vv *anypb.Any
			if x.FAny2 == nil {
				vv = new(anypb.Any)
			} else {
				vv = x.FAny2
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FAny2 = vv
		case jsonKey == "f_message1":
			var vv *MessageOptional1
			if x.FMessage1 == nil {
				vv = new(MessageOptional1)
			} else {
				vv = x.FMessage1
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage1 = vv
		case jsonKey == "f_message2":
			var vv *MessageOptional1_Embed1
			if x.FMessage2 == nil {
				vv = new(MessageOptional1_Embed1)
			} else {
				vv = x.FMessage2
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage2 = vv
		case jsonKey == "f_message3":
			var vv *MessageOptional1_Embed1_Embed2
			if x.FMessage3 == nil {
				vv = new(MessageOptional1_Embed1_Embed2)
			} else {
				vv = x.FMessage3
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage3 = vv
		case jsonKey == "f_message4":
			var vv *pbexternal.Message1
			if x.FMessage4 == nil {
				vv = new(pbexternal.Message1)
			} else {
				vv = x.FMessage4
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage4 = vv
		case jsonKey == "f_message5":
			var vv *pbexternal.Message1_Embed1
			if x.FMessage5 == nil {
				vv = new(pbexternal.Message1_Embed1)
			} else {
				vv = x.FMessage5
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage5 = vv
		case jsonKey == "f_message6":
			var vv *pbexternal.Message1_Embed1_Embed2
			if x.FMessage6 == nil {
				vv = new(pbexternal.Message1_Embed1_Embed2)
			} else {
				vv = x.FMessage6
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage6 = vv
		case jsonKey == "f_message7":
			var vv *MessageCommon1
			if x.FMessage7 == nil {
				vv = new(MessageCommon1)
			} else {
				vv = x.FMessage7
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage7 = vv
		case jsonKey == "f_message8":
			var vv *MessageCommon1_Embed1
			if x.FMessage8 == nil {
				vv = new(MessageCommon1_Embed1)
			} else {
				vv = x.FMessage8
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage8 = vv
		case jsonKey == "f_message9":
			var vv *MessageCommon1_Embed1_Embed2
			if x.FMessage9 == nil {
				vv = new(MessageCommon1_Embed1_Embed2)
			} else {
				vv = x.FMessage9
			}
			ok, _err := decoder.ReadValueInterface(jsonKey, vv)
			if !ok { // The field is null
				vv = nil
			}
			if _err != nil {
				return _err
			}
			x.FMessage9 = vv
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
