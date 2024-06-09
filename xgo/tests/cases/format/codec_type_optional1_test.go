package format

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_CodecTypeOptional1(t *testing.T) {
	seed1 := &pbformat.CodecTypeOptional1{
		FInt32Numeric:              utils.PointerInt32(111),
		FInt32String:               utils.PointerInt32(1112),
		FInt64Numeric:              utils.PointerInt64(121),
		FInt64String:               utils.PointerInt64(1212),
		FUint32Numeric:             utils.PointerUint32(131),
		FUint32String:              utils.PointerUint32(1312),
		FUint64Numeric:             utils.PointerUint64(141),
		FUint64String:              utils.PointerUint64(1412),
		FSint32Numeric:             utils.PointerInt32(151),
		FSint32String:              utils.PointerInt32(1512),
		FSint64Numeric:             utils.PointerInt64(161),
		FSint64String:              utils.PointerInt64(1612),
		FSfixed32Numeric:           utils.PointerInt32(171),
		FSfixed32String:            utils.PointerInt32(1712),
		FSfixed64Numeric:           utils.PointerInt64(181),
		FSfixed64String:            utils.PointerInt64(1812),
		FFixed32Numeric:            utils.PointerUint32(191),
		FFixed32String:             utils.PointerUint32(1912),
		FFixed64Numeric:            utils.PointerUint64(201),
		FFixed64String:             utils.PointerUint64(2012),
		FFloatNumeric:              utils.PointerFloat32(2111.11),
		FFloatString:               utils.PointerFloat32(2111.112),
		FDoubleNumeric:             utils.PointerFloat64(2211.11),
		FDoubleString:              utils.PointerFloat64(2211.112),
		FBoolBool:                  utils.PointerBool(true),
		FBoolString:                utils.PointerBool(false),
		FEnumNumeric:               pbexternal.EnumNum1(2).Enum(),
		FEnumNumericString:         pbexternal.EnumNum1(7).Enum(),
		FEnumString:                pbexternal.EnumNum1(5).Enum(),
		FAnyNative:                 utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		FAnyProto:                  utils.MustNewAny(&pbexternal.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
		FDurationNative:            &durationpb.Duration{Seconds: 111, Nanos: 0},
		FDurationString:            &durationpb.Duration{Seconds: 112, Nanos: 0},
		FDurationNanosecond:        &durationpb.Duration{Seconds: 113, Nanos: 0},
		FDurationNanosecondString:  &durationpb.Duration{Seconds: 114, Nanos: 0},
		FDurationMicrosecond:       &durationpb.Duration{Seconds: 115, Nanos: 0},
		FDurationMicrosecondString: &durationpb.Duration{Seconds: 116, Nanos: 0},
		FDurationMillisecond:       &durationpb.Duration{Seconds: 117, Nanos: 0},
		FDurationMillisecondString: &durationpb.Duration{Seconds: 118, Nanos: 0},
		FDurationSecond:            &durationpb.Duration{Seconds: 119, Nanos: 0},
		FDurationSecondString:      &durationpb.Duration{Seconds: 120, Nanos: 0},
		FDurationMinute:            &durationpb.Duration{Seconds: 240, Nanos: 0},
		FDurationMinuteString:      &durationpb.Duration{Seconds: 480, Nanos: 0},
		FDurationHour:              &durationpb.Duration{Seconds: 3600, Nanos: 0},
		FDurationHourString:        &durationpb.Duration{Seconds: 7200, Nanos: 0},
		FTimestampNative:           &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampTimeLayout:       &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixNano:         &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixNanoString:   &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixMicro:        &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixMicroString:  &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixMilli:        &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixMilliString:  &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixSec:          &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		FTimestampUnixSecString:    &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	}

	var (
		err error
		bb  []byte
	)

	t.Run("Marshal", func(t *testing.T) {
		bb, err = seed1.MarshalJSON()
		require.Nil(t, err)
	})

	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pbformat.CodecTypeOptional1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})

	t.Run("Check", func(t *testing.T) {
		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("Int32", func(t *testing.T) {
			require.Equal(t, float64(111), data["f_int32_numeric"].(float64))
			require.Equal(t, "1112", data["f_int32_string"].(string))
		})
		t.Run("Int64", func(t *testing.T) {
			require.Equal(t, float64(121), data["f_int64_numeric"].(float64))
			require.Equal(t, "1212", data["f_int64_string"].(string))
		})
		t.Run("UInt32", func(t *testing.T) {
			require.Equal(t, float64(131), data["f_uint32_numeric"].(float64))
			require.Equal(t, "1312", data["f_uint32_string"].(string))
		})
		t.Run("UInt64", func(t *testing.T) {
			require.Equal(t, float64(141), data["f_uint64_numeric"].(float64))
			require.Equal(t, "1412", data["f_uint64_string"].(string))
		})
		t.Run("SInt32", func(t *testing.T) {
			require.Equal(t, float64(151), data["f_sint32_numeric"].(float64))
			require.Equal(t, "1512", data["f_sint32_string"].(string))
		})
		t.Run("SInt64", func(t *testing.T) {
			require.Equal(t, float64(161), data["f_sint64_numeric"].(float64))
			require.Equal(t, "1612", data["f_sint64_string"].(string))
		})
		t.Run("SFixed32", func(t *testing.T) {
			require.Equal(t, float64(171), data["f_sfixed32_numeric"].(float64))
			require.Equal(t, "1712", data["f_sfixed32_string"].(string))
		})
		t.Run("SFixed64", func(t *testing.T) {
			require.Equal(t, float64(181), data["f_sfixed64_numeric"].(float64))
			require.Equal(t, "1812", data["f_sfixed64_string"].(string))
		})
		t.Run("Fixed32", func(t *testing.T) {
			require.Equal(t, float64(191), data["f_fixed32_numeric"].(float64))
			require.Equal(t, "1912", data["f_fixed32_string"].(string))
		})
		t.Run("Fixed64", func(t *testing.T) {
			require.Equal(t, float64(201), data["f_fixed64_numeric"].(float64))
			require.Equal(t, "2012", data["f_fixed64_string"].(string))
		})
		t.Run("Float", func(t *testing.T) {
			require.Equal(t, float64(2111.11), data["f_float_numeric"].(float64))
			require.Equal(t, "2111.112", data["f_float_string"].(string))
		})
		t.Run("Double", func(t *testing.T) {
			require.Equal(t, float64(2211.11), data["f_double_numeric"].(float64))
			require.Equal(t, "2211.112", data["f_double_string"].(string))
		})
		t.Run("Bool", func(t *testing.T) {
			require.Equal(t, true, data["f_bool_bool"].(bool))
			require.Equal(t, "false", data["f_bool_string"].(string))
		})

		t.Run("Enum", func(t *testing.T) {
			enumNumeric := data["f_enum_numeric"].(float64)
			require.Equal(t, float64(2), enumNumeric)

			enumNumericString := data["f_enum_numeric_string"].(string)
			require.Equal(t, "7", enumNumericString)

			enumString := data["f_enum_string"].(string)
			require.Equal(t, "Three", enumString)
		})

		t.Run("Any", func(t *testing.T) {
			anyNative := data["f_any_native"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", anyNative["type_url"].(string))
			require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", anyNative["value"].(string))

			anyProto := data["f_any_proto"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", anyProto["@type"].(string))
			require.Equal(t, "s301", anyProto["f_string1"].(string))
			require.Equal(t, "s302", anyProto["f_string2"].(string))
			require.Equal(t, "s303", anyProto["f_string3"].(string))
		})

		t.Run("Duration", func(t *testing.T) {
			anyDurationNative := data["f_duration_native"].(map[string]interface{})
			require.Equal(t, float64(111), anyDurationNative["seconds"].(float64))

			anyDurationString := data["f_duration_string"].(string)
			require.Equal(t, "1m52s", anyDurationString)

			anyDurationNanosecond := data["f_duration_nanosecond"].(float64)
			require.Equal(t, float64(113000000000), anyDurationNanosecond)

			anyDurationNanosecondString := data["f_duration_nanosecond_string"].(string)
			require.Equal(t, "114000000000", anyDurationNanosecondString)

			anyDurationMicrosecond := data["f_duration_microsecond"].(float64)
			require.Equal(t, float64(115000000), anyDurationMicrosecond)

			anyDurationMicrosecondString := data["f_duration_microsecond_string"].(string)
			require.Equal(t, "116000000", anyDurationMicrosecondString)

			anyDurationMillisecond := data["f_duration_millisecond"].(float64)
			require.Equal(t, float64(117000), anyDurationMillisecond)

			anyDurationMillisecondString := data["f_duration_millisecond_string"].(string)
			require.Equal(t, "118000", anyDurationMillisecondString)

			anyDurationSecond := data["f_duration_second"].(float64)
			require.Equal(t, float64(119), anyDurationSecond)

			anyDurationSecondString := data["f_duration_second_string"].(string)
			require.Equal(t, "120", anyDurationSecondString)

			anyDurationMinute := data["f_duration_minute"].(float64)
			require.Equal(t, float64(4), anyDurationMinute)

			anyDurationMinuteString := data["f_duration_minute_string"].(string)
			require.Equal(t, "8", anyDurationMinuteString)

			anyDurationHour := data["f_duration_hour"].(float64)
			require.Equal(t, float64(1), anyDurationHour)

			anyDurationHourString := data["f_duration_hour_string"].(string)
			require.Equal(t, "2", anyDurationHourString)
		})

		t.Run("Timestamp", func(t *testing.T) {
			anyTimestampNative := data["f_timestamp_native"].(map[string]interface{})
			require.Equal(t, float64(1686416585), anyTimestampNative["seconds"].(float64))

			anyTimestampTimeLayout := data["f_timestamp_time_layout"].(string)
			require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout)

			anyTimestampUnixNano := data["f_timestamp_unix_nano"].(float64)
			require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano)

			anyTimestampUnixNanoString := data["f_timestamp_unix_nano_string"].(string)
			require.Equal(t, "1686416585000000000", anyTimestampUnixNanoString)

			anyTimestampUnixMicro := data["f_timestamp_unix_micro"].(float64)
			require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro)

			anyTimestampUnixMicroString := data["f_timestamp_unix_micro_string"].(string)
			require.Equal(t, "1686416585000000", anyTimestampUnixMicroString)

			anyTimestampUnixMilli := data["f_timestamp_unix_milli"].(float64)
			require.Equal(t, float64(1686416585000), anyTimestampUnixMilli)

			anyTimestampUnixMilliString := data["f_timestamp_unix_milli_string"].(string)
			require.Equal(t, "1686416585000", anyTimestampUnixMilliString)

			anyTimestampUnixSec := data["f_timestamp_unix_sec"].(float64)
			require.Equal(t, float64(1686416585), anyTimestampUnixSec)

			anyTimestampUnixSecString := data["f_timestamp_unix_sec_string"].(string)
			require.Equal(t, "1686416585", anyTimestampUnixSecString)
		})
	})

}
