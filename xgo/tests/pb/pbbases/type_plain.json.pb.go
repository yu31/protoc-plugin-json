// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/bases/type_plain.proto

package pbbases

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

// MarshalJSON implements interface json.Marshaler for proto message MessagePlain1 in file tests/proto/cases/bases/type_plain.proto
func (x *MessagePlain1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

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

// UnmarshalJSON implements json.Unmarshaler for proto message MessagePlain1 in file tests/proto/cases/bases/type_plain.proto
func (x *MessagePlain1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbases.(*MessagePlain1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/bases/type_plain.proto
func (x *MessagePlain1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/bases/type_plain.proto
func (x *MessagePlain1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbases.(*MessagePlain1_Embed1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/bases/type_plain.proto
func (x *MessagePlain1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

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

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/bases/type_plain.proto
func (x *MessagePlain1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbases.(*MessagePlain1_Embed1_Embed2) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message TypePlain1 in file tests/proto/cases/bases/type_plain.proto
func (x *TypePlain1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(800)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)
	encoder.AppendObjectKey("f_int32")
	encoder.AppendLiteralInt32(x.FInt32, false)
	encoder.AppendObjectKey("f_int64")
	encoder.AppendLiteralInt64(x.FInt64, false)
	encoder.AppendObjectKey("f_uint32")
	encoder.AppendLiteralUint32(x.FUint32, false)
	encoder.AppendObjectKey("f_uint64")
	encoder.AppendLiteralUint64(x.FUint64, false)
	encoder.AppendObjectKey("f_sint32")
	encoder.AppendLiteralInt32(x.FSint32, false)
	encoder.AppendObjectKey("f_sint64")
	encoder.AppendLiteralInt64(x.FSint64, false)
	encoder.AppendObjectKey("f_sfixed32")
	encoder.AppendLiteralInt32(x.FSfixed32, false)
	encoder.AppendObjectKey("f_sfixed64")
	encoder.AppendLiteralInt64(x.FSfixed64, false)
	encoder.AppendObjectKey("f_fixed32")
	encoder.AppendLiteralUint32(x.FFixed32, false)
	encoder.AppendObjectKey("f_fixed64")
	encoder.AppendLiteralUint64(x.FFixed64, false)
	encoder.AppendObjectKey("f_float")
	encoder.AppendLiteralFloat32(x.FFloat, false)
	encoder.AppendObjectKey("f_double")
	encoder.AppendLiteralFloat64(x.FDouble, false)
	encoder.AppendObjectKey("f_bool1")
	encoder.AppendLiteralBool(x.FBool1, false)
	encoder.AppendObjectKey("f_bytes1")
	encoder.AppendLiteralBytes(x.FBytes1)
	encoder.AppendObjectKey("f_enum1")
	encoder.AppendLiteralInt32(int32(x.FEnum1.Number()), false)
	encoder.AppendObjectKey("f_enum2")
	encoder.AppendLiteralInt32(int32(x.FEnum2.Number()), false)
	encoder.AppendObjectKey("f_enum3")
	encoder.AppendLiteralInt32(int32(x.FEnum3.Number()), false)
	encoder.AppendObjectKey("f_enum4")
	encoder.AppendLiteralInt32(int32(x.FEnum4.Number()), false)
	encoder.AppendObjectKey("f_enum5")
	encoder.AppendLiteralInt32(int32(x.FEnum5.Number()), false)
	encoder.AppendObjectKey("f_enum6")
	encoder.AppendLiteralInt32(int32(x.FEnum6.Number()), false)
	encoder.AppendObjectKey("f_duration1")
	if err = encoder.AppendLiteralInterface(x.FDuration1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_duration2")
	if err = encoder.AppendLiteralInterface(x.FDuration2); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_timestamp1")
	if err = encoder.AppendLiteralInterface(x.FTimestamp1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_timestamp2")
	if err = encoder.AppendLiteralInterface(x.FTimestamp2); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_any1")
	if err = encoder.AppendLiteralInterface(x.FAny1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_any2")
	if err = encoder.AppendLiteralInterface(x.FAny2); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message1")
	if err = encoder.AppendLiteralInterface(x.FMessage1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message2")
	if err = encoder.AppendLiteralInterface(x.FMessage2); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message3")
	if err = encoder.AppendLiteralInterface(x.FMessage3); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message4")
	if err = encoder.AppendLiteralInterface(x.FMessage4); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message5")
	if err = encoder.AppendLiteralInterface(x.FMessage5); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message6")
	if err = encoder.AppendLiteralInterface(x.FMessage6); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message7")
	if err = encoder.AppendLiteralInterface(x.FMessage7); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message8")
	if err = encoder.AppendLiteralInterface(x.FMessage8); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message9")
	if err = encoder.AppendLiteralInterface(x.FMessage9); err != nil {
		return nil, err
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypePlain1 in file tests/proto/cases/bases/type_plain.proto
func (x *TypePlain1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbases.(*TypePlain1) is nil")
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
		case "f_int32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FInt32 = vv
		case "f_int64":
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			x.FInt64 = vv
		case "f_uint32":
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			x.FUint32 = vv
		case "f_uint64":
			var vv uint64
			if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
				return err
			}
			x.FUint64 = vv
		case "f_sint32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FSint32 = vv
		case "f_sint64":
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			x.FSint64 = vv
		case "f_sfixed32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FSfixed32 = vv
		case "f_sfixed64":
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			x.FSfixed64 = vv
		case "f_fixed32":
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			x.FFixed32 = vv
		case "f_fixed64":
			var vv uint64
			if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
				return err
			}
			x.FFixed64 = vv
		case "f_float":
			var vv float32
			if vv, err = decoder.ReadLiteralFloat32(jsonKey, false); err != nil {
				return err
			}
			x.FFloat = vv
		case "f_double":
			var vv float64
			if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
				return err
			}
			x.FDouble = vv
		case "f_bool1":
			var vv bool
			if vv, err = decoder.ReadLiteralBool(jsonKey, false); err != nil {
				return err
			}
			x.FBool1 = vv
		case "f_bytes1":
			var vv []byte
			if vv, err = decoder.ReadLiteralBytes(jsonKey); err != nil {
				return err
			}
			x.FBytes1 = vv
		case "f_enum1":
			var vv EnumPlain1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = EnumPlain1(v1)
			x.FEnum1 = vv
		case "f_enum2":
			var vv pbexternal.Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = pbexternal.Enum1(v1)
			x.FEnum2 = vv
		case "f_enum3":
			var vv pbexternal.Embed_Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = pbexternal.Embed_Enum1(v1)
			x.FEnum3 = vv
		case "f_enum4":
			var vv pbexternal.Embed_Message_Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = pbexternal.Embed_Message_Enum1(v1)
			x.FEnum4 = vv
		case "f_enum5":
			var vv EnumCommon1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = EnumCommon1(v1)
			x.FEnum5 = vv
		case "f_enum6":
			var vv MessageCommon1_Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = MessageCommon1_Enum1(v1)
			x.FEnum6 = vv
		case "f_duration1":
			var vv *durationpb.Duration
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FDuration1 != nil {
					vv = x.FDuration1
				} else {
					vv = new(durationpb.Duration)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FDuration1 = vv
		case "f_duration2":
			var vv *durationpb.Duration
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FDuration2 != nil {
					vv = x.FDuration2
				} else {
					vv = new(durationpb.Duration)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FDuration2 = vv
		case "f_timestamp1":
			var vv *timestamppb.Timestamp
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FTimestamp1 != nil {
					vv = x.FTimestamp1
				} else {
					vv = new(timestamppb.Timestamp)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FTimestamp1 = vv
		case "f_timestamp2":
			var vv *timestamppb.Timestamp
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FTimestamp2 != nil {
					vv = x.FTimestamp2
				} else {
					vv = new(timestamppb.Timestamp)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FTimestamp2 = vv
		case "f_any1":
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FAny1 != nil {
					vv = x.FAny1
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FAny1 = vv
		case "f_any2":
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FAny2 != nil {
					vv = x.FAny2
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FAny2 = vv
		case "f_message1":
			var vv *MessagePlain1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage1 != nil {
					vv = x.FMessage1
				} else {
					vv = new(MessagePlain1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage1 = vv
		case "f_message2":
			var vv *MessagePlain1_Embed1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage2 != nil {
					vv = x.FMessage2
				} else {
					vv = new(MessagePlain1_Embed1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage2 = vv
		case "f_message3":
			var vv *MessagePlain1_Embed1_Embed2
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage3 != nil {
					vv = x.FMessage3
				} else {
					vv = new(MessagePlain1_Embed1_Embed2)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage3 = vv
		case "f_message4":
			var vv *pbexternal.Message1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage4 != nil {
					vv = x.FMessage4
				} else {
					vv = new(pbexternal.Message1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage4 = vv
		case "f_message5":
			var vv *pbexternal.Message1_Embed1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage5 != nil {
					vv = x.FMessage5
				} else {
					vv = new(pbexternal.Message1_Embed1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage5 = vv
		case "f_message6":
			var vv *pbexternal.Message1_Embed1_Embed2
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage6 != nil {
					vv = x.FMessage6
				} else {
					vv = new(pbexternal.Message1_Embed1_Embed2)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage6 = vv
		case "f_message7":
			var vv *MessageCommon1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage7 != nil {
					vv = x.FMessage7
				} else {
					vv = new(MessageCommon1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage7 = vv
		case "f_message8":
			var vv *MessageCommon1_Embed1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage8 != nil {
					vv = x.FMessage8
				} else {
					vv = new(MessageCommon1_Embed1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage8 = vv
		case "f_message9":
			var vv *MessageCommon1_Embed1_Embed2
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage9 != nil {
					vv = x.FMessage9
				} else {
					vv = new(MessageCommon1_Embed1_Embed2)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage9 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
