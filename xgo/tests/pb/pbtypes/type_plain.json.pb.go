// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/types/type_plain.proto

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

// MarshalJSON implements interface json.Marshaler for proto message MessagePlain1 in file tests/proto/cases/types/type_plain.proto
func (x *MessagePlain1) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message MessagePlain1 in file tests/proto/cases/types/type_plain.proto
func (x *MessagePlain1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessagePlain1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/types/type_plain.proto
func (x *MessagePlain1_Embed1) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/types/type_plain.proto
func (x *MessagePlain1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessagePlain1_Embed1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/types/type_plain.proto
func (x *MessagePlain1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/types/type_plain.proto
func (x *MessagePlain1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*MessagePlain1_Embed1_Embed2) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message TypePlain1 in file tests/proto/cases/types/type_plain.proto
func (x *TypePlain1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(794)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)
	encoder.AppendJSONKey("f_int32")
	encoder.AppendValueInt32(x.FInt32)
	encoder.AppendJSONKey("f_int64")
	encoder.AppendValueInt64(x.FInt64)
	encoder.AppendJSONKey("f_uint32")
	encoder.AppendValueUint32(x.FUint32)
	encoder.AppendJSONKey("f_uint64")
	encoder.AppendValueUint64(x.FUint64)
	encoder.AppendJSONKey("f_sint32")
	encoder.AppendValueInt32(x.FSint32)
	encoder.AppendJSONKey("f_sint64")
	encoder.AppendValueInt64(x.FSint64)
	encoder.AppendJSONKey("f_sfixed32")
	encoder.AppendValueInt32(x.FSfixed32)
	encoder.AppendJSONKey("f_sfixed64")
	encoder.AppendValueInt64(x.FSfixed64)
	encoder.AppendJSONKey("f_fixed32")
	encoder.AppendValueUint32(x.FFixed32)
	encoder.AppendJSONKey("f_fixed64")
	encoder.AppendValueUint64(x.FFixed64)
	encoder.AppendJSONKey("f_float")
	encoder.AppendValueFloat32(x.FFloat)
	encoder.AppendJSONKey("f_double")
	encoder.AppendValueFloat64(x.FDouble)
	encoder.AppendJSONKey("f_bool1")
	encoder.AppendValueBool(x.FBool1)
	encoder.AppendJSONKey("f_bytes1")
	encoder.AppendValueBytes(x.FBytes1)
	encoder.AppendJSONKey("f_enum1")
	encoder.AppendValueInt32(int32(x.FEnum1.Number()))
	encoder.AppendJSONKey("f_enum2")
	encoder.AppendValueInt32(int32(x.FEnum2.Number()))
	encoder.AppendJSONKey("f_enum3")
	encoder.AppendValueInt32(int32(x.FEnum3.Number()))
	encoder.AppendJSONKey("f_enum4")
	encoder.AppendValueInt32(int32(x.FEnum4.Number()))
	encoder.AppendJSONKey("f_enum5")
	encoder.AppendValueInt32(int32(x.FEnum5.Number()))
	encoder.AppendJSONKey("f_enum6")
	encoder.AppendValueInt32(int32(x.FEnum6.Number()))
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

// UnmarshalJSON implements json.Unmarshaler for proto message TypePlain1 in file tests/proto/cases/types/type_plain.proto
func (x *TypePlain1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbtypes.(*TypePlain1) is nil")
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
		case "f_int32":
			var vv int32
			vv, err = decoder.ReadValueInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FInt32 = vv
		case "f_int64":
			var vv int64
			vv, err = decoder.ReadValueInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FInt64 = vv
		case "f_uint32":
			var vv uint32
			vv, err = decoder.ReadValueUint32(jsonKey)
			if err != nil {
				return err
			}
			x.FUint32 = vv
		case "f_uint64":
			var vv uint64
			vv, err = decoder.ReadValueUint64(jsonKey)
			if err != nil {
				return err
			}
			x.FUint64 = vv
		case "f_sint32":
			var vv int32
			vv, err = decoder.ReadValueInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FSint32 = vv
		case "f_sint64":
			var vv int64
			vv, err = decoder.ReadValueInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FSint64 = vv
		case "f_sfixed32":
			var vv int32
			vv, err = decoder.ReadValueInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FSfixed32 = vv
		case "f_sfixed64":
			var vv int64
			vv, err = decoder.ReadValueInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FSfixed64 = vv
		case "f_fixed32":
			var vv uint32
			vv, err = decoder.ReadValueUint32(jsonKey)
			if err != nil {
				return err
			}
			x.FFixed32 = vv
		case "f_fixed64":
			var vv uint64
			vv, err = decoder.ReadValueUint64(jsonKey)
			if err != nil {
				return err
			}
			x.FFixed64 = vv
		case "f_float":
			var vv float32
			vv, err = decoder.ReadValueFloat32(jsonKey)
			if err != nil {
				return err
			}
			x.FFloat = vv
		case "f_double":
			var vv float64
			vv, err = decoder.ReadValueFloat64(jsonKey)
			if err != nil {
				return err
			}
			x.FDouble = vv
		case "f_bool1":
			var vv bool
			vv, err = decoder.ReadValueBool(jsonKey)
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
			var vv EnumPlain1
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, EnumPlain1_name)
			vv = EnumPlain1(v1)
			if err != nil {
				return err
			}
			x.FEnum1 = vv
		case "f_enum2":
			var vv pbexternal.Enum1
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, pbexternal.Enum1_name)
			vv = pbexternal.Enum1(v1)
			if err != nil {
				return err
			}
			x.FEnum2 = vv
		case "f_enum3":
			var vv pbexternal.Embed_Enum1
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, pbexternal.Embed_Enum1_name)
			vv = pbexternal.Embed_Enum1(v1)
			if err != nil {
				return err
			}
			x.FEnum3 = vv
		case "f_enum4":
			var vv pbexternal.Embed_Message_Enum1
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, pbexternal.Embed_Message_Enum1_name)
			vv = pbexternal.Embed_Message_Enum1(v1)
			if err != nil {
				return err
			}
			x.FEnum4 = vv
		case "f_enum5":
			var vv EnumCommon1
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, EnumCommon1_name)
			vv = EnumCommon1(v1)
			if err != nil {
				return err
			}
			x.FEnum5 = vv
		case "f_enum6":
			var vv MessageCommon1_Enum1
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, MessageCommon1_Enum1_name)
			vv = MessageCommon1_Enum1(v1)
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
			var vv *MessagePlain1
			initFN := func() interface{} {
				if x.FMessage1 != nil {
					vv = x.FMessage1
				} else {
					vv = new(MessagePlain1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage1 = vv
		case "f_message2":
			var vv *MessagePlain1_Embed1
			initFN := func() interface{} {
				if x.FMessage2 != nil {
					vv = x.FMessage2
				} else {
					vv = new(MessagePlain1_Embed1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage2 = vv
		case "f_message3":
			var vv *MessagePlain1_Embed1_Embed2
			initFN := func() interface{} {
				if x.FMessage3 != nil {
					vv = x.FMessage3
				} else {
					vv = new(MessagePlain1_Embed1_Embed2)
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
