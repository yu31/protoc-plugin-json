// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/base/type_optional.proto

package pbbase

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	_ "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message MessageOptional1 in file tests/proto/cases/base/type_optional.proto
func (x *MessageOptional1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(72)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageOptional1 in file tests/proto/cases/base/type_optional.proto
func (x *MessageOptional1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbase.(*MessageOptional1) is nil")
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
		case "f_string1":
			if x.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/base/type_optional.proto
func (x *MessageOptional1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(72)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/base/type_optional.proto
func (x *MessageOptional1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbase.(*MessageOptional1_Embed1) is nil")
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
		case "f_string1":
			if x.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/base/type_optional.proto
func (x *MessageOptional1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(72)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/base/type_optional.proto
func (x *MessageOptional1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbase.(*MessageOptional1_Embed1_Embed2) is nil")
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
		case "f_string1":
			if x.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message TypeOptional1 in file tests/proto/cases/base/type_optional.proto
func (x *TypeOptional1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(800)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendPtrStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendPtrStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendPtrI32(enc, "f_int32", x.FInt32, false, false)
	jsonencoder.AppendPtrI64(enc, "f_int64", x.FInt64, false, false)
	jsonencoder.AppendPtrU32(enc, "f_uint32", x.FUint32, false, false)
	jsonencoder.AppendPtrU64(enc, "f_uint64", x.FUint64, false, false)
	jsonencoder.AppendPtrI32(enc, "f_sint32", x.FSint32, false, false)
	jsonencoder.AppendPtrI64(enc, "f_sint64", x.FSint64, false, false)
	jsonencoder.AppendPtrI32(enc, "f_sfixed32", x.FSfixed32, false, false)
	jsonencoder.AppendPtrI64(enc, "f_sfixed64", x.FSfixed64, false, false)
	jsonencoder.AppendPtrU32(enc, "f_fixed32", x.FFixed32, false, false)
	jsonencoder.AppendPtrU64(enc, "f_fixed64", x.FFixed64, false, false)
	jsonencoder.AppendPtrF32(enc, "f_float", x.FFloat, false, false)
	jsonencoder.AppendPtrF64(enc, "f_double", x.FDouble, false, false)
	jsonencoder.AppendPtrBool(enc, "f_bool1", x.FBool1, false, false)
	if err := jsonencoder.AppendValBytes(enc, "f_bytes1", x.FBytes1, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendPtrEnumNum(enc, "f_enum1", x.FEnum1, false, false)
	jsonencoder.AppendPtrEnumNum(enc, "f_enum2", x.FEnum2, false, false)
	jsonencoder.AppendPtrEnumNum(enc, "f_enum3", x.FEnum3, false, false)
	jsonencoder.AppendPtrEnumNum(enc, "f_enum4", x.FEnum4, false, false)
	jsonencoder.AppendPtrEnumNum(enc, "f_enum5", x.FEnum5, false, false)
	jsonencoder.AppendPtrEnumNum(enc, "f_enum6", x.FEnum6, false, false)
	if err := jsonencoder.AppendValWKTDurObject(enc, "f_duration1", x.FDuration1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValWKTDurObject(enc, "f_duration2", x.FDuration2, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValWKTTsObject(enc, "f_timestamp1", x.FTimestamp1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValWKTTsObject(enc, "f_timestamp2", x.FTimestamp2, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValWKTAnyObject(enc, "f_any1", x.FAny1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValWKTAnyObject(enc, "f_any2", x.FAny2, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message1", x.FMessage1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message2", x.FMessage2, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message3", x.FMessage3, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message4", x.FMessage4, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message5", x.FMessage5, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message6", x.FMessage6, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message7", x.FMessage7, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message8", x.FMessage8, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "f_message9", x.FMessage9, false); err != nil {
		return nil, err
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOptional1 in file tests/proto/cases/base/type_optional.proto
func (x *TypeOptional1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbbase.(*TypeOptional1) is nil")
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
		case "f_string1":
			if x.FString1, err = jsondecoder.ReadPtrStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadPtrStr(dec); err != nil {
				return err
			}
		case "f_int32":
			if x.FInt32, err = jsondecoder.ReadPtrI32(dec, false); err != nil {
				return err
			}
		case "f_int64":
			if x.FInt64, err = jsondecoder.ReadPtrI64(dec, false); err != nil {
				return err
			}
		case "f_uint32":
			if x.FUint32, err = jsondecoder.ReadPtrU32(dec, false); err != nil {
				return err
			}
		case "f_uint64":
			if x.FUint64, err = jsondecoder.ReadPtrU64(dec, false); err != nil {
				return err
			}
		case "f_sint32":
			if x.FSint32, err = jsondecoder.ReadPtrI32(dec, false); err != nil {
				return err
			}
		case "f_sint64":
			if x.FSint64, err = jsondecoder.ReadPtrI64(dec, false); err != nil {
				return err
			}
		case "f_sfixed32":
			if x.FSfixed32, err = jsondecoder.ReadPtrI32(dec, false); err != nil {
				return err
			}
		case "f_sfixed64":
			if x.FSfixed64, err = jsondecoder.ReadPtrI64(dec, false); err != nil {
				return err
			}
		case "f_fixed32":
			if x.FFixed32, err = jsondecoder.ReadPtrU32(dec, false); err != nil {
				return err
			}
		case "f_fixed64":
			if x.FFixed64, err = jsondecoder.ReadPtrU64(dec, false); err != nil {
				return err
			}
		case "f_float":
			if x.FFloat, err = jsondecoder.ReadPtrF32(dec, false); err != nil {
				return err
			}
		case "f_double":
			if x.FDouble, err = jsondecoder.ReadPtrF64(dec, false); err != nil {
				return err
			}
		case "f_bool1":
			if x.FBool1, err = jsondecoder.ReadPtrBool(dec, false); err != nil {
				return err
			}
		case "f_bytes1":
			if x.FBytes1, err = jsondecoder.ReadValBytes(dec); err != nil {
				return err
			}
		case "f_enum1":
			if x.FEnum1, err = jsondecoder.ReadPtrEnumNum(dec, x.FEnum1, false); err != nil {
				return err
			}
		case "f_enum2":
			if x.FEnum2, err = jsondecoder.ReadPtrEnumNum(dec, x.FEnum2, false); err != nil {
				return err
			}
		case "f_enum3":
			if x.FEnum3, err = jsondecoder.ReadPtrEnumNum(dec, x.FEnum3, false); err != nil {
				return err
			}
		case "f_enum4":
			if x.FEnum4, err = jsondecoder.ReadPtrEnumNum(dec, x.FEnum4, false); err != nil {
				return err
			}
		case "f_enum5":
			if x.FEnum5, err = jsondecoder.ReadPtrEnumNum(dec, x.FEnum5, false); err != nil {
				return err
			}
		case "f_enum6":
			if x.FEnum6, err = jsondecoder.ReadPtrEnumNum(dec, x.FEnum6, false); err != nil {
				return err
			}
		case "f_duration1":
			if x.FDuration1, err = jsondecoder.ReadValWKTDurObject(dec, x.FDuration1); err != nil {
				return err
			}
		case "f_duration2":
			if x.FDuration2, err = jsondecoder.ReadValWKTDurObject(dec, x.FDuration2); err != nil {
				return err
			}
		case "f_timestamp1":
			if x.FTimestamp1, err = jsondecoder.ReadValWKTTsObject(dec, x.FTimestamp1); err != nil {
				return err
			}
		case "f_timestamp2":
			if x.FTimestamp2, err = jsondecoder.ReadValWKTTsObject(dec, x.FTimestamp2); err != nil {
				return err
			}
		case "f_any1":
			if x.FAny1, err = jsondecoder.ReadValWKTAnyObject(dec, x.FAny1); err != nil {
				return err
			}
		case "f_any2":
			if x.FAny2, err = jsondecoder.ReadValWKTAnyObject(dec, x.FAny2); err != nil {
				return err
			}
		case "f_message1":
			if x.FMessage1, err = jsondecoder.ReadValMessage(dec, x.FMessage1); err != nil {
				return err
			}
		case "f_message2":
			if x.FMessage2, err = jsondecoder.ReadValMessage(dec, x.FMessage2); err != nil {
				return err
			}
		case "f_message3":
			if x.FMessage3, err = jsondecoder.ReadValMessage(dec, x.FMessage3); err != nil {
				return err
			}
		case "f_message4":
			if x.FMessage4, err = jsondecoder.ReadValMessage(dec, x.FMessage4); err != nil {
				return err
			}
		case "f_message5":
			if x.FMessage5, err = jsondecoder.ReadValMessage(dec, x.FMessage5); err != nil {
				return err
			}
		case "f_message6":
			if x.FMessage6, err = jsondecoder.ReadValMessage(dec, x.FMessage6); err != nil {
				return err
			}
		case "f_message7":
			if x.FMessage7, err = jsondecoder.ReadValMessage(dec, x.FMessage7); err != nil {
				return err
			}
		case "f_message8":
			if x.FMessage8, err = jsondecoder.ReadValMessage(dec, x.FMessage8); err != nil {
				return err
			}
		case "f_message9":
			if x.FMessage9, err = jsondecoder.ReadValMessage(dec, x.FMessage9); err != nil {
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