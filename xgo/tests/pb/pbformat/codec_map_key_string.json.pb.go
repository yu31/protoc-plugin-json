// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/format/codec_map_key_string.proto

package pbformat

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	pbexternal "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message CodecMapKeyString in file tests/proto/cases/format/codec_map_key_string.proto
func (x *CodecMapKeyString) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(4864)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapStrI32(enc, "f_int32_key_none_val_unset", x.FInt32KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrI32(enc, "f_int32_key_none_val_numeric", x.FInt32KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrI32(enc, "f_int32_key_none_val_string", x.FInt32KeyNoneValString, false, true)
	jsonencoder.AppendMapStrI64(enc, "f_int64_key_none_val_unset", x.FInt64KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrI64(enc, "f_int64_key_none_val_numeric", x.FInt64KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrI64(enc, "f_int64_key_none_val_string", x.FInt64KeyNoneValString, false, true)
	jsonencoder.AppendMapStrU32(enc, "f_uint32_key_none_val_unset", x.FUint32KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrU32(enc, "f_uint32_key_none_val_numeric", x.FUint32KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrU32(enc, "f_uint32_key_none_val_string", x.FUint32KeyNoneValString, false, true)
	jsonencoder.AppendMapStrU64(enc, "f_uint64_key_none_val_unset", x.FUint64KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrU64(enc, "f_uint64_key_none_val_numeric", x.FUint64KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrU64(enc, "f_uint64_key_none_val_string", x.FUint64KeyNoneValString, false, true)
	jsonencoder.AppendMapStrI32(enc, "f_sint32_key_none_val_unset", x.FSint32KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrI32(enc, "f_sint32_key_none_val_numeric", x.FSint32KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrI32(enc, "f_sint32_key_none_val_string", x.FSint32KeyNoneValString, false, true)
	jsonencoder.AppendMapStrI64(enc, "f_sint64_key_none_val_unset", x.FSint64KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrI64(enc, "f_sint64_key_none_val_numeric", x.FSint64KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrI64(enc, "f_sint64_key_none_val_string", x.FSint64KeyNoneValString, false, true)
	jsonencoder.AppendMapStrI32(enc, "f_sfixed32_key_none_val_unset", x.FSfixed32KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrI32(enc, "f_sfixed32_key_none_val_numeric", x.FSfixed32KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrI32(enc, "f_sfixed32_key_none_val_string", x.FSfixed32KeyNoneValString, false, true)
	jsonencoder.AppendMapStrI64(enc, "f_sfixed64_key_none_val_unset", x.FSfixed64KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrI64(enc, "f_sfixed64_key_none_val_numeric", x.FSfixed64KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrI64(enc, "f_sfixed64_key_none_val_string", x.FSfixed64KeyNoneValString, false, true)
	jsonencoder.AppendMapStrU32(enc, "f_fixed32_key_none_val_unset", x.FFixed32KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrU32(enc, "f_fixed32_key_none_val_numeric", x.FFixed32KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrU32(enc, "f_fixed32_key_none_val_string", x.FFixed32KeyNoneValString, false, true)
	jsonencoder.AppendMapStrU64(enc, "f_fixed64_key_none_val_unset", x.FFixed64KeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrU64(enc, "f_fixed64_key_none_val_numeric", x.FFixed64KeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrU64(enc, "f_fixed64_key_none_val_string", x.FFixed64KeyNoneValString, false, true)
	jsonencoder.AppendMapStrF32(enc, "f_float_key_none_val_unset", x.FFloatKeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrF32(enc, "f_float_key_none_val_numeric", x.FFloatKeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrF32(enc, "f_float_key_none_val_string", x.FFloatKeyNoneValString, false, true)
	jsonencoder.AppendMapStrF64(enc, "f_double_key_none_val_unset", x.FDoubleKeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrF64(enc, "f_double_key_none_val_numeric", x.FDoubleKeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrF64(enc, "f_double_key_none_val_string", x.FDoubleKeyNoneValString, false, true)
	jsonencoder.AppendMapStrBool(enc, "f_bool_key_none_val_unset", x.FBoolKeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrBool(enc, "f_bool_key_none_val_bool", x.FBoolKeyNoneValBool, false, false)
	jsonencoder.AppendMapStrBool(enc, "f_bool_key_none_val_string", x.FBoolKeyNoneValString, false, true)
	jsonencoder.AppendMapStrStr(enc, "f_string_key_none_val_none", x.FStringKeyNoneValNone, false)
	if err := jsonencoder.AppendMapStrBytes(enc, "f_bytes_key_none_val_none", x.FBytesKeyNoneValNone, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendMapStrEnumNum(enc, "f_enum_key_none_val_unset", x.FEnumKeyNoneValUnset, false, false)
	jsonencoder.AppendMapStrEnumNum(enc, "f_enum_key_none_val_numeric", x.FEnumKeyNoneValNumeric, false, false)
	jsonencoder.AppendMapStrEnumNum(enc, "f_enum_key_none_val_numeric_string", x.FEnumKeyNoneValNumericString, false, true)
	jsonencoder.AppendMapStrEnumStr(enc, "f_enum_key_none_val_enum_string", x.FEnumKeyNoneValEnumString, false)
	if err := jsonencoder.AppendMapStrMessage(enc, "f_message_key_none_val_none", x.FMessageKeyNoneValNone, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendMapStrWKTDurObject(enc, "f_duration_key_none_val_unset", x.FDurationKeyNoneValUnset, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendMapStrWKTDurObject(enc, "f_duration_key_none_val_object", x.FDurationKeyNoneValObject, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendMapStrWKTDurTimeStr(enc, "f_duration_key_none_val_time_string", x.FDurationKeyNoneValTimeString, false)
	jsonencoder.AppendMapStrWKTDurNano(enc, "f_duration_key_none_val_nanosecond", x.FDurationKeyNoneValNanosecond, false, false)
	jsonencoder.AppendMapStrWKTDurNano(enc, "f_duration_key_none_val_nanosecond_string", x.FDurationKeyNoneValNanosecondString, false, true)
	jsonencoder.AppendMapStrWKTDurMicro(enc, "f_duration_key_none_val_microsecond", x.FDurationKeyNoneValMicrosecond, false, false)
	jsonencoder.AppendMapStrWKTDurMicro(enc, "f_duration_key_none_val_microsecond_string", x.FDurationKeyNoneValMicrosecondString, false, true)
	jsonencoder.AppendMapStrWKTDurMilli(enc, "f_duration_key_none_val_millisecond", x.FDurationKeyNoneValMillisecond, false, false)
	jsonencoder.AppendMapStrWKTDurMilli(enc, "f_duration_key_none_val_millisecond_string", x.FDurationKeyNoneValMillisecondString, false, true)
	jsonencoder.AppendMapStrWKTDurSecond(enc, "f_duration_key_none_val_second", x.FDurationKeyNoneValSecond, false, false)
	jsonencoder.AppendMapStrWKTDurSecond(enc, "f_duration_key_none_val_second_string", x.FDurationKeyNoneValSecondString, false, true)
	jsonencoder.AppendMapStrWKTDurMinute(enc, "f_duration_key_none_val_minute", x.FDurationKeyNoneValMinute, false, false)
	jsonencoder.AppendMapStrWKTDurMinute(enc, "f_duration_key_none_val_minute_string", x.FDurationKeyNoneValMinuteString, false, true)
	jsonencoder.AppendMapStrWKTDurHour(enc, "f_duration_key_none_val_hour", x.FDurationKeyNoneValHour, false, false)
	jsonencoder.AppendMapStrWKTDurHour(enc, "f_duration_key_none_val_hour_string", x.FDurationKeyNoneValHourString, false, true)
	if err := jsonencoder.AppendMapStrWKTTsObject(enc, "f_timestamp_key_none_val_unset", x.FTimestampKeyNoneValUnset, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendMapStrWKTTsObject(enc, "f_timestamp_key_none_val_object", x.FTimestampKeyNoneValObject, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendMapStrWKTTsLayout(enc, "f_timestamp_key_none_val_time_layout", x.FTimestampKeyNoneValTimeLayout, false, "2006-01-02T15:04:05.999999999Z07:00")
	jsonencoder.AppendMapStrWKTTsUnixNano(enc, "f_timestamp_key_none_val_unix_nano", x.FTimestampKeyNoneValUnixNano, false, false)
	jsonencoder.AppendMapStrWKTTsUnixNano(enc, "f_timestamp_key_none_val_unix_nano_string", x.FTimestampKeyNoneValUnixNanoString, false, true)
	jsonencoder.AppendMapStrWKTTsUnixMicro(enc, "f_timestamp_key_none_val_unix_micro", x.FTimestampKeyNoneValUnixMicro, false, false)
	jsonencoder.AppendMapStrWKTTsUnixMicro(enc, "f_timestamp_key_none_val_unix_micro_string", x.FTimestampKeyNoneValUnixMicroString, false, true)
	jsonencoder.AppendMapStrWKTTsUnixMilli(enc, "f_timestamp_key_none_val_unix_milli", x.FTimestampKeyNoneValUnixMilli, false, false)
	jsonencoder.AppendMapStrWKTTsUnixMilli(enc, "f_timestamp_key_none_val_unix_milli_string", x.FTimestampKeyNoneValUnixMilliString, false, true)
	jsonencoder.AppendMapStrWKTTsUnixSec(enc, "f_timestamp_key_none_val_unix_sec", x.FTimestampKeyNoneValUnixSec, false, false)
	jsonencoder.AppendMapStrWKTTsUnixSec(enc, "f_timestamp_key_none_val_unix_sec_string", x.FTimestampKeyNoneValUnixSecString, false, true)
	if err := jsonencoder.AppendMapStrWKTAnyObject(enc, "f_any_key_none_val_unset", x.FAnyKeyNoneValUnset, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendMapStrWKTAnyObject(enc, "f_any_key_none_val_object", x.FAnyKeyNoneValObject, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendMapStrWKTAnyProto(enc, "f_any_key_none_val_proto", x.FAnyKeyNoneValProto, false); err != nil {
		return nil, err
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message CodecMapKeyString in file tests/proto/cases/format/codec_map_key_string.proto
func (x *CodecMapKeyString) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbformat.(*CodecMapKeyString) is nil")
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
		case "f_int32_key_none_val_unset":
			if x.FInt32KeyNoneValUnset, err = jsondecoder.ReadMapStrI32(dec, x.FInt32KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_int32_key_none_val_numeric":
			if x.FInt32KeyNoneValNumeric, err = jsondecoder.ReadMapStrI32(dec, x.FInt32KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_int32_key_none_val_string":
			if x.FInt32KeyNoneValString, err = jsondecoder.ReadMapStrI32(dec, x.FInt32KeyNoneValString, true); err != nil {
				return err
			}
		case "f_int64_key_none_val_unset":
			if x.FInt64KeyNoneValUnset, err = jsondecoder.ReadMapStrI64(dec, x.FInt64KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_int64_key_none_val_numeric":
			if x.FInt64KeyNoneValNumeric, err = jsondecoder.ReadMapStrI64(dec, x.FInt64KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_int64_key_none_val_string":
			if x.FInt64KeyNoneValString, err = jsondecoder.ReadMapStrI64(dec, x.FInt64KeyNoneValString, true); err != nil {
				return err
			}
		case "f_uint32_key_none_val_unset":
			if x.FUint32KeyNoneValUnset, err = jsondecoder.ReadMapStrU32(dec, x.FUint32KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_uint32_key_none_val_numeric":
			if x.FUint32KeyNoneValNumeric, err = jsondecoder.ReadMapStrU32(dec, x.FUint32KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_uint32_key_none_val_string":
			if x.FUint32KeyNoneValString, err = jsondecoder.ReadMapStrU32(dec, x.FUint32KeyNoneValString, true); err != nil {
				return err
			}
		case "f_uint64_key_none_val_unset":
			if x.FUint64KeyNoneValUnset, err = jsondecoder.ReadMapStrU64(dec, x.FUint64KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_uint64_key_none_val_numeric":
			if x.FUint64KeyNoneValNumeric, err = jsondecoder.ReadMapStrU64(dec, x.FUint64KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_uint64_key_none_val_string":
			if x.FUint64KeyNoneValString, err = jsondecoder.ReadMapStrU64(dec, x.FUint64KeyNoneValString, true); err != nil {
				return err
			}
		case "f_sint32_key_none_val_unset":
			if x.FSint32KeyNoneValUnset, err = jsondecoder.ReadMapStrI32(dec, x.FSint32KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_sint32_key_none_val_numeric":
			if x.FSint32KeyNoneValNumeric, err = jsondecoder.ReadMapStrI32(dec, x.FSint32KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_sint32_key_none_val_string":
			if x.FSint32KeyNoneValString, err = jsondecoder.ReadMapStrI32(dec, x.FSint32KeyNoneValString, true); err != nil {
				return err
			}
		case "f_sint64_key_none_val_unset":
			if x.FSint64KeyNoneValUnset, err = jsondecoder.ReadMapStrI64(dec, x.FSint64KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_sint64_key_none_val_numeric":
			if x.FSint64KeyNoneValNumeric, err = jsondecoder.ReadMapStrI64(dec, x.FSint64KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_sint64_key_none_val_string":
			if x.FSint64KeyNoneValString, err = jsondecoder.ReadMapStrI64(dec, x.FSint64KeyNoneValString, true); err != nil {
				return err
			}
		case "f_sfixed32_key_none_val_unset":
			if x.FSfixed32KeyNoneValUnset, err = jsondecoder.ReadMapStrI32(dec, x.FSfixed32KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_sfixed32_key_none_val_numeric":
			if x.FSfixed32KeyNoneValNumeric, err = jsondecoder.ReadMapStrI32(dec, x.FSfixed32KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_sfixed32_key_none_val_string":
			if x.FSfixed32KeyNoneValString, err = jsondecoder.ReadMapStrI32(dec, x.FSfixed32KeyNoneValString, true); err != nil {
				return err
			}
		case "f_sfixed64_key_none_val_unset":
			if x.FSfixed64KeyNoneValUnset, err = jsondecoder.ReadMapStrI64(dec, x.FSfixed64KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_sfixed64_key_none_val_numeric":
			if x.FSfixed64KeyNoneValNumeric, err = jsondecoder.ReadMapStrI64(dec, x.FSfixed64KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_sfixed64_key_none_val_string":
			if x.FSfixed64KeyNoneValString, err = jsondecoder.ReadMapStrI64(dec, x.FSfixed64KeyNoneValString, true); err != nil {
				return err
			}
		case "f_fixed32_key_none_val_unset":
			if x.FFixed32KeyNoneValUnset, err = jsondecoder.ReadMapStrU32(dec, x.FFixed32KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_fixed32_key_none_val_numeric":
			if x.FFixed32KeyNoneValNumeric, err = jsondecoder.ReadMapStrU32(dec, x.FFixed32KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_fixed32_key_none_val_string":
			if x.FFixed32KeyNoneValString, err = jsondecoder.ReadMapStrU32(dec, x.FFixed32KeyNoneValString, true); err != nil {
				return err
			}
		case "f_fixed64_key_none_val_unset":
			if x.FFixed64KeyNoneValUnset, err = jsondecoder.ReadMapStrU64(dec, x.FFixed64KeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_fixed64_key_none_val_numeric":
			if x.FFixed64KeyNoneValNumeric, err = jsondecoder.ReadMapStrU64(dec, x.FFixed64KeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_fixed64_key_none_val_string":
			if x.FFixed64KeyNoneValString, err = jsondecoder.ReadMapStrU64(dec, x.FFixed64KeyNoneValString, true); err != nil {
				return err
			}
		case "f_float_key_none_val_unset":
			if x.FFloatKeyNoneValUnset, err = jsondecoder.ReadMapStrF32(dec, x.FFloatKeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_float_key_none_val_numeric":
			if x.FFloatKeyNoneValNumeric, err = jsondecoder.ReadMapStrF32(dec, x.FFloatKeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_float_key_none_val_string":
			if x.FFloatKeyNoneValString, err = jsondecoder.ReadMapStrF32(dec, x.FFloatKeyNoneValString, true); err != nil {
				return err
			}
		case "f_double_key_none_val_unset":
			if x.FDoubleKeyNoneValUnset, err = jsondecoder.ReadMapStrF64(dec, x.FDoubleKeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_double_key_none_val_numeric":
			if x.FDoubleKeyNoneValNumeric, err = jsondecoder.ReadMapStrF64(dec, x.FDoubleKeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_double_key_none_val_string":
			if x.FDoubleKeyNoneValString, err = jsondecoder.ReadMapStrF64(dec, x.FDoubleKeyNoneValString, true); err != nil {
				return err
			}
		case "f_bool_key_none_val_unset":
			if x.FBoolKeyNoneValUnset, err = jsondecoder.ReadMapStrBool(dec, x.FBoolKeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_bool_key_none_val_bool":
			if x.FBoolKeyNoneValBool, err = jsondecoder.ReadMapStrBool(dec, x.FBoolKeyNoneValBool, false); err != nil {
				return err
			}
		case "f_bool_key_none_val_string":
			if x.FBoolKeyNoneValString, err = jsondecoder.ReadMapStrBool(dec, x.FBoolKeyNoneValString, true); err != nil {
				return err
			}
		case "f_string_key_none_val_none":
			if x.FStringKeyNoneValNone, err = jsondecoder.ReadMapStrStr(dec, x.FStringKeyNoneValNone); err != nil {
				return err
			}
		case "f_bytes_key_none_val_none":
			if x.FBytesKeyNoneValNone, err = jsondecoder.ReadMapStrBytes(dec, x.FBytesKeyNoneValNone); err != nil {
				return err
			}
		case "f_enum_key_none_val_unset":
			if x.FEnumKeyNoneValUnset, err = jsondecoder.ReadMapStrEnumNum(dec, x.FEnumKeyNoneValUnset, false); err != nil {
				return err
			}
		case "f_enum_key_none_val_numeric":
			if x.FEnumKeyNoneValNumeric, err = jsondecoder.ReadMapStrEnumNum(dec, x.FEnumKeyNoneValNumeric, false); err != nil {
				return err
			}
		case "f_enum_key_none_val_numeric_string":
			if x.FEnumKeyNoneValNumericString, err = jsondecoder.ReadMapStrEnumNum(dec, x.FEnumKeyNoneValNumericString, true); err != nil {
				return err
			}
		case "f_enum_key_none_val_enum_string":
			if x.FEnumKeyNoneValEnumString, err = jsondecoder.ReadMapStrEnumStr(dec, x.FEnumKeyNoneValEnumString, pbexternal.EnumNum1_value); err != nil {
				return err
			}
		case "f_message_key_none_val_none":
			if x.FMessageKeyNoneValNone, err = jsondecoder.ReadMapStrMessage(dec, x.FMessageKeyNoneValNone); err != nil {
				return err
			}
		case "f_duration_key_none_val_unset":
			if x.FDurationKeyNoneValUnset, err = jsondecoder.ReadMapStrWKTDurObject(dec, x.FDurationKeyNoneValUnset); err != nil {
				return err
			}
		case "f_duration_key_none_val_object":
			if x.FDurationKeyNoneValObject, err = jsondecoder.ReadMapStrWKTDurObject(dec, x.FDurationKeyNoneValObject); err != nil {
				return err
			}
		case "f_duration_key_none_val_time_string":
			if x.FDurationKeyNoneValTimeString, err = jsondecoder.ReadMapStrWKTDurTimeStr(dec, x.FDurationKeyNoneValTimeString); err != nil {
				return err
			}
		case "f_duration_key_none_val_nanosecond":
			if x.FDurationKeyNoneValNanosecond, err = jsondecoder.ReadMapStrWKTDurNano(dec, x.FDurationKeyNoneValNanosecond, false); err != nil {
				return err
			}
		case "f_duration_key_none_val_nanosecond_string":
			if x.FDurationKeyNoneValNanosecondString, err = jsondecoder.ReadMapStrWKTDurNano(dec, x.FDurationKeyNoneValNanosecondString, true); err != nil {
				return err
			}
		case "f_duration_key_none_val_microsecond":
			if x.FDurationKeyNoneValMicrosecond, err = jsondecoder.ReadMapStrWKTDurMicro(dec, x.FDurationKeyNoneValMicrosecond, false); err != nil {
				return err
			}
		case "f_duration_key_none_val_microsecond_string":
			if x.FDurationKeyNoneValMicrosecondString, err = jsondecoder.ReadMapStrWKTDurMicro(dec, x.FDurationKeyNoneValMicrosecondString, true); err != nil {
				return err
			}
		case "f_duration_key_none_val_millisecond":
			if x.FDurationKeyNoneValMillisecond, err = jsondecoder.ReadMapStrWKTDurMilli(dec, x.FDurationKeyNoneValMillisecond, false); err != nil {
				return err
			}
		case "f_duration_key_none_val_millisecond_string":
			if x.FDurationKeyNoneValMillisecondString, err = jsondecoder.ReadMapStrWKTDurMilli(dec, x.FDurationKeyNoneValMillisecondString, true); err != nil {
				return err
			}
		case "f_duration_key_none_val_second":
			if x.FDurationKeyNoneValSecond, err = jsondecoder.ReadMapStrWKTDurSecond(dec, x.FDurationKeyNoneValSecond, false); err != nil {
				return err
			}
		case "f_duration_key_none_val_second_string":
			if x.FDurationKeyNoneValSecondString, err = jsondecoder.ReadMapStrWKTDurSecond(dec, x.FDurationKeyNoneValSecondString, true); err != nil {
				return err
			}
		case "f_duration_key_none_val_minute":
			if x.FDurationKeyNoneValMinute, err = jsondecoder.ReadMapStrWKTDurMinute(dec, x.FDurationKeyNoneValMinute, false); err != nil {
				return err
			}
		case "f_duration_key_none_val_minute_string":
			if x.FDurationKeyNoneValMinuteString, err = jsondecoder.ReadMapStrWKTDurMinute(dec, x.FDurationKeyNoneValMinuteString, true); err != nil {
				return err
			}
		case "f_duration_key_none_val_hour":
			if x.FDurationKeyNoneValHour, err = jsondecoder.ReadMapStrWKTDurHour(dec, x.FDurationKeyNoneValHour, false); err != nil {
				return err
			}
		case "f_duration_key_none_val_hour_string":
			if x.FDurationKeyNoneValHourString, err = jsondecoder.ReadMapStrWKTDurHour(dec, x.FDurationKeyNoneValHourString, true); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unset":
			if x.FTimestampKeyNoneValUnset, err = jsondecoder.ReadMapStrWKTTsObject(dec, x.FTimestampKeyNoneValUnset); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_object":
			if x.FTimestampKeyNoneValObject, err = jsondecoder.ReadMapStrWKTTsObject(dec, x.FTimestampKeyNoneValObject); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_time_layout":
			if x.FTimestampKeyNoneValTimeLayout, err = jsondecoder.ReadMapStrWKTTsLayout(dec, x.FTimestampKeyNoneValTimeLayout, "2006-01-02T15:04:05.999999999Z07:00"); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_nano":
			if x.FTimestampKeyNoneValUnixNano, err = jsondecoder.ReadMapStrWKTTsUnixNano(dec, x.FTimestampKeyNoneValUnixNano, false); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_nano_string":
			if x.FTimestampKeyNoneValUnixNanoString, err = jsondecoder.ReadMapStrWKTTsUnixNano(dec, x.FTimestampKeyNoneValUnixNanoString, true); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_micro":
			if x.FTimestampKeyNoneValUnixMicro, err = jsondecoder.ReadMapStrWKTTsUnixMicro(dec, x.FTimestampKeyNoneValUnixMicro, false); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_micro_string":
			if x.FTimestampKeyNoneValUnixMicroString, err = jsondecoder.ReadMapStrWKTTsUnixMicro(dec, x.FTimestampKeyNoneValUnixMicroString, true); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_milli":
			if x.FTimestampKeyNoneValUnixMilli, err = jsondecoder.ReadMapStrWKTTsUnixMilli(dec, x.FTimestampKeyNoneValUnixMilli, false); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_milli_string":
			if x.FTimestampKeyNoneValUnixMilliString, err = jsondecoder.ReadMapStrWKTTsUnixMilli(dec, x.FTimestampKeyNoneValUnixMilliString, true); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_sec":
			if x.FTimestampKeyNoneValUnixSec, err = jsondecoder.ReadMapStrWKTTsUnixSec(dec, x.FTimestampKeyNoneValUnixSec, false); err != nil {
				return err
			}
		case "f_timestamp_key_none_val_unix_sec_string":
			if x.FTimestampKeyNoneValUnixSecString, err = jsondecoder.ReadMapStrWKTTsUnixSec(dec, x.FTimestampKeyNoneValUnixSecString, true); err != nil {
				return err
			}
		case "f_any_key_none_val_unset":
			if x.FAnyKeyNoneValUnset, err = jsondecoder.ReadMapStrWKTAnyObject(dec, x.FAnyKeyNoneValUnset); err != nil {
				return err
			}
		case "f_any_key_none_val_object":
			if x.FAnyKeyNoneValObject, err = jsondecoder.ReadMapStrWKTAnyObject(dec, x.FAnyKeyNoneValObject); err != nil {
				return err
			}
		case "f_any_key_none_val_proto":
			if x.FAnyKeyNoneValProto, err = jsondecoder.ReadMapStrWKTAnyProto(dec, x.FAnyKeyNoneValProto); err != nil {
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
