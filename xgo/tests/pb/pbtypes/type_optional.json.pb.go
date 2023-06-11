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

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeScanJSON(); err != nil {
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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string1":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
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

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeScanJSON(); err != nil {
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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string1":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
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

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeScanJSON(); err != nil {
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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string1":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
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
	if err = encoder.AppendValueInterface(x.FDuration1); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_duration2")
	if err = encoder.AppendValueInterface(x.FDuration2); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_timestamp1")
	if err = encoder.AppendValueInterface(x.FTimestamp1); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_timestamp2")
	if err = encoder.AppendValueInterface(x.FTimestamp2); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_any1")
	if err = encoder.AppendValueInterface(x.FAny1); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_any2")
	if err = encoder.AppendValueInterface(x.FAny2); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message1")
	if err = encoder.AppendValueInterface(x.FMessage1); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message2")
	if err = encoder.AppendValueInterface(x.FMessage2); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message3")
	if err = encoder.AppendValueInterface(x.FMessage3); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message4")
	if err = encoder.AppendValueInterface(x.FMessage4); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message5")
	if err = encoder.AppendValueInterface(x.FMessage5); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message6")
	if err = encoder.AppendValueInterface(x.FMessage6); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message7")
	if err = encoder.AppendValueInterface(x.FMessage7); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message8")
	if err = encoder.AppendValueInterface(x.FMessage8); err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message9")
	if err = encoder.AppendValueInterface(x.FMessage9); err != nil {
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

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeScanJSON(); err != nil {
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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string1":
			var vv *string
			vv, err = decoder.ReadPointerString(jsonKey)
			if err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv *string
			vv, err = decoder.ReadPointerString(jsonKey)
			if err != nil {
				return err
			}
			x.FString2 = vv
		case "f_int32":
			var vv *int32
			vv, err = decoder.ReadPointerInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FInt32 = vv
		case "f_int64":
			var vv *int64
			vv, err = decoder.ReadPointerInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FInt64 = vv
		case "f_uint32":
			var vv *uint32
			vv, err = decoder.ReadPointerUint32(jsonKey)
			if err != nil {
				return err
			}
			x.FUint32 = vv
		case "f_uint64":
			var vv *uint64
			vv, err = decoder.ReadPointerUint64(jsonKey)
			if err != nil {
				return err
			}
			x.FUint64 = vv
		case "f_sint32":
			var vv *int32
			vv, err = decoder.ReadPointerInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FSint32 = vv
		case "f_sint64":
			var vv *int64
			vv, err = decoder.ReadPointerInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FSint64 = vv
		case "f_sfixed32":
			var vv *int32
			vv, err = decoder.ReadPointerInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FSfixed32 = vv
		case "f_sfixed64":
			var vv *int64
			vv, err = decoder.ReadPointerInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FSfixed64 = vv
		case "f_fixed32":
			var vv *uint32
			vv, err = decoder.ReadPointerUint32(jsonKey)
			if err != nil {
				return err
			}
			x.FFixed32 = vv
		case "f_fixed64":
			var vv *uint64
			vv, err = decoder.ReadPointerUint64(jsonKey)
			if err != nil {
				return err
			}
			x.FFixed64 = vv
		case "f_float":
			var vv *float32
			vv, err = decoder.ReadPointerFloat32(jsonKey)
			if err != nil {
				return err
			}
			x.FFloat = vv
		case "f_double":
			var vv *float64
			vv, err = decoder.ReadPointerFloat64(jsonKey)
			if err != nil {
				return err
			}
			x.FDouble = vv
		case "f_bool1":
			var vv *bool
			vv, err = decoder.ReadPointerBool(jsonKey)
			if err != nil {
				return err
			}
			x.FBool1 = vv
		case "f_bytes1":
			var vv []byte
			vv, err = decoder.ReadValueBytes(jsonKey)
			if err != nil {
				return err
			}
			x.FBytes1 = vv
		case "f_enum1":
			var vv *EnumOptional1
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, EnumOptional1_name)
			if v1 != nil {
				v2 := EnumOptional1(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum1 = vv
		case "f_enum2":
			var vv *pbexternal.Enum1
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, pbexternal.Enum1_name)
			if v1 != nil {
				v2 := pbexternal.Enum1(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum2 = vv
		case "f_enum3":
			var vv *pbexternal.Embed_Enum1
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, pbexternal.Embed_Enum1_name)
			if v1 != nil {
				v2 := pbexternal.Embed_Enum1(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum3 = vv
		case "f_enum4":
			var vv *pbexternal.Embed_Message_Enum1
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, pbexternal.Embed_Message_Enum1_name)
			if v1 != nil {
				v2 := pbexternal.Embed_Message_Enum1(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum4 = vv
		case "f_enum5":
			var vv *EnumCommon1
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, EnumCommon1_name)
			if v1 != nil {
				v2 := EnumCommon1(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum5 = vv
		case "f_enum6":
			var vv *MessageCommon1_Enum1
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, MessageCommon1_Enum1_name)
			if v1 != nil {
				v2 := MessageCommon1_Enum1(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum6 = vv
		case "f_duration1":
			var vv *durationpb.Duration
			initFN := func() interface{} {
				if x.FDuration1 != nil {
					vv = x.FDuration1
				} else {
					vv = new(durationpb.Duration)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FDuration1 = vv
		case "f_duration2":
			var vv *durationpb.Duration
			initFN := func() interface{} {
				if x.FDuration2 != nil {
					vv = x.FDuration2
				} else {
					vv = new(durationpb.Duration)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FDuration2 = vv
		case "f_timestamp1":
			var vv *timestamppb.Timestamp
			initFN := func() interface{} {
				if x.FTimestamp1 != nil {
					vv = x.FTimestamp1
				} else {
					vv = new(timestamppb.Timestamp)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FTimestamp1 = vv
		case "f_timestamp2":
			var vv *timestamppb.Timestamp
			initFN := func() interface{} {
				if x.FTimestamp2 != nil {
					vv = x.FTimestamp2
				} else {
					vv = new(timestamppb.Timestamp)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FTimestamp2 = vv
		case "f_any1":
			var vv *anypb.Any
			initFN := func() interface{} {
				if x.FAny1 != nil {
					vv = x.FAny1
				} else {
					vv = new(anypb.Any)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FAny1 = vv
		case "f_any2":
			var vv *anypb.Any
			initFN := func() interface{} {
				if x.FAny2 != nil {
					vv = x.FAny2
				} else {
					vv = new(anypb.Any)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FAny2 = vv
		case "f_message1":
			var vv *MessageOptional1
			initFN := func() interface{} {
				if x.FMessage1 != nil {
					vv = x.FMessage1
				} else {
					vv = new(MessageOptional1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage1 = vv
		case "f_message2":
			var vv *MessageOptional1_Embed1
			initFN := func() interface{} {
				if x.FMessage2 != nil {
					vv = x.FMessage2
				} else {
					vv = new(MessageOptional1_Embed1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage2 = vv
		case "f_message3":
			var vv *MessageOptional1_Embed1_Embed2
			initFN := func() interface{} {
				if x.FMessage3 != nil {
					vv = x.FMessage3
				} else {
					vv = new(MessageOptional1_Embed1_Embed2)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage3 = vv
		case "f_message4":
			var vv *pbexternal.Message1
			initFN := func() interface{} {
				if x.FMessage4 != nil {
					vv = x.FMessage4
				} else {
					vv = new(pbexternal.Message1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage4 = vv
		case "f_message5":
			var vv *pbexternal.Message1_Embed1
			initFN := func() interface{} {
				if x.FMessage5 != nil {
					vv = x.FMessage5
				} else {
					vv = new(pbexternal.Message1_Embed1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage5 = vv
		case "f_message6":
			var vv *pbexternal.Message1_Embed1_Embed2
			initFN := func() interface{} {
				if x.FMessage6 != nil {
					vv = x.FMessage6
				} else {
					vv = new(pbexternal.Message1_Embed1_Embed2)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage6 = vv
		case "f_message7":
			var vv *MessageCommon1
			initFN := func() interface{} {
				if x.FMessage7 != nil {
					vv = x.FMessage7
				} else {
					vv = new(MessageCommon1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage7 = vv
		case "f_message8":
			var vv *MessageCommon1_Embed1
			initFN := func() interface{} {
				if x.FMessage8 != nil {
					vv = x.FMessage8
				} else {
					vv = new(MessageCommon1_Embed1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage8 = vv
		case "f_message9":
			var vv *MessageCommon1_Embed1_Embed2
			initFN := func() interface{} {
				if x.FMessage9 != nil {
					vv = x.FMessage9
				} else {
					vv = new(MessageCommon1_Embed1_Embed2)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage9 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}
