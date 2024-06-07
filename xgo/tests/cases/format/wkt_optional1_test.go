package format

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedWKTOptional1 = &pbformat.WKTOptional1{
	FEnumNumber1:                pbformat.Enum1(2).Enum(),
	FEnumNumberString1:          pbformat.Enum1(7).Enum(),
	FEnumString1:                pbformat.Enum1(5).Enum(),
	FAnyNative1:                 utils.MustNewAny(&pbformat.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
	FAnyProto1:                  utils.MustNewAny(&pbformat.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
	FDurationNative1:            &durationpb.Duration{Seconds: 111, Nanos: 0},
	FDurationString1:            &durationpb.Duration{Seconds: 112, Nanos: 0},
	FDurationNanosecond1:        &durationpb.Duration{Seconds: 113, Nanos: 0},
	FDurationNanosecondString1:  &durationpb.Duration{Seconds: 114, Nanos: 0},
	FDurationMicrosecond1:       &durationpb.Duration{Seconds: 115, Nanos: 0},
	FDurationMicrosecondString1: &durationpb.Duration{Seconds: 116, Nanos: 0},
	FDurationMillisecond1:       &durationpb.Duration{Seconds: 117, Nanos: 0},
	FDurationMillisecondString1: &durationpb.Duration{Seconds: 118, Nanos: 0},
	FDurationSecond1:            &durationpb.Duration{Seconds: 119, Nanos: 0},
	FDurationSecondString1:      &durationpb.Duration{Seconds: 120, Nanos: 0},
	FDurationMinute1:            &durationpb.Duration{Seconds: 240, Nanos: 0},
	FDurationMinuteString1:      &durationpb.Duration{Seconds: 480, Nanos: 0},
	FDurationHour1:              &durationpb.Duration{Seconds: 3600, Nanos: 0},
	FDurationHourString1:        &durationpb.Duration{Seconds: 7200, Nanos: 0},
	FTimestampNative1:           &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampTimeLayout1:       &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixNano1:         &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixNanoString1:   &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixMicro1:        &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixMicroString1:  &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixMilli1:        &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixMilliString1:  &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixSec1:          &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixSecString1:    &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
}

func Test_Reference_WKT_Optional1_General(t *testing.T) {
	bb, err := seedWKTOptional1.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbformat.WKTOptional1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedWKTOptional1, dataNew)
}

func Test_Reference_WKT_Optional1_Check(t *testing.T) {
	bb, err := seedWKTOptional1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("Enum", func(t *testing.T) {
		/*
			"f_enum_number1":2,
			"f_enum_number_string1":"7",
			"f_enum_string1":"Three",
		*/
		enumNumber1 := data["f_enum_number1"].(float64)
		require.Equal(t, float64(2), enumNumber1)

		enumNumberString1 := data["f_enum_number_string1"].(string)
		require.Equal(t, "7", enumNumberString1)

		enumString1 := data["f_enum_string1"].(string)
		require.Equal(t, "Three", enumString1)
	})

	t.Run("Any", func(t *testing.T) {
		/*
			"f_any_native1":{"type_url":"type.googleapis.com/private.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
			"f_any_proto1":{"@type":"type.googleapis.com/private.Message1","f_string1":"s301", "f_string2":"s302", "f_string3":"s303"},
		*/
		anyNative1 := data["f_any_native1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/private.Message1", anyNative1["type_url"].(string))
		require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", anyNative1["value"].(string))

		anyProto1 := data["f_any_proto1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/private.Message1", anyProto1["@type"].(string))
		require.Equal(t, "s301", anyProto1["f_string1"].(string))
		require.Equal(t, "s302", anyProto1["f_string2"].(string))
		require.Equal(t, "s303", anyProto1["f_string3"].(string))
	})

	t.Run("Duration", func(t *testing.T) {
		/*
			"f_duration_native1":{"seconds":111},
			"f_duration_string1":"1m52s",
			"f_duration_nanosecond1":113000000000,
			"f_duration_nanosecond_string1":"114000000000",
			"f_duration_microsecond1":115000000,
			"f_duration_microsecond_string1":"116000000",
			"f_duration_millisecond1":117000,
			"f_duration_millisecond_string1":"118000",
			"f_duration_second1":119,
			"f_duration_second_string1":"120",
			"f_duration_minute1":4,
			"f_duration_minute_string1":"8",
			"f_duration_hour1":1,
			"f_duration_hour_string1":"2",
		*/

		anyDurationNative1 := data["f_duration_native1"].(map[string]interface{})
		require.Equal(t, float64(111), anyDurationNative1["seconds"].(float64))

		anyDurationString1 := data["f_duration_string1"].(string)
		require.Equal(t, "1m52s", anyDurationString1)

		anyDurationNanosecond1 := data["f_duration_nanosecond1"].(float64)
		require.Equal(t, float64(113000000000), anyDurationNanosecond1)

		anyDurationNanosecondString1 := data["f_duration_nanosecond_string1"].(string)
		require.Equal(t, "114000000000", anyDurationNanosecondString1)

		anyDurationMicrosecond1 := data["f_duration_microsecond1"].(float64)
		require.Equal(t, float64(115000000), anyDurationMicrosecond1)

		anyDurationMicrosecondString1 := data["f_duration_microsecond_string1"].(string)
		require.Equal(t, "116000000", anyDurationMicrosecondString1)

		anyDurationMillisecond1 := data["f_duration_millisecond1"].(float64)
		require.Equal(t, float64(117000), anyDurationMillisecond1)

		anyDurationMillisecondString1 := data["f_duration_millisecond_string1"].(string)
		require.Equal(t, "118000", anyDurationMillisecondString1)

		anyDurationSecond1 := data["f_duration_second1"].(float64)
		require.Equal(t, float64(119), anyDurationSecond1)

		anyDurationSecondString1 := data["f_duration_second_string1"].(string)
		require.Equal(t, "120", anyDurationSecondString1)

		anyDurationMinute1 := data["f_duration_minute1"].(float64)
		require.Equal(t, float64(4), anyDurationMinute1)

		anyDurationMinuteString1 := data["f_duration_minute_string1"].(string)
		require.Equal(t, "8", anyDurationMinuteString1)

		anyDurationHour1 := data["f_duration_hour1"].(float64)
		require.Equal(t, float64(1), anyDurationHour1)

		anyDurationHourString1 := data["f_duration_hour_string1"].(string)
		require.Equal(t, "2", anyDurationHourString1)
	})

	t.Run("Timestamp", func(t *testing.T) {
		/*
			"f_timestamp_native1":{"seconds":1686416585},
			"f_timestamp_time_layout1":"2023-06-10T17:03:05Z",
			"f_timestamp_unix_nano1":1686416585000000000,
			"f_timestamp_unix_nano_string1":"1686416585000000000",
			"f_timestamp_unix_micro1":1686416585000000,
			"f_timestamp_unix_micro_string1":"1686416585000000",
			"f_timestamp_unix_milli1":1686416585000,
			"f_timestamp_unix_milli_string1":"1686416585000",
			"f_timestamp_unix_sec1":1686416585,
			"f_timestamp_unix_sec_string1":"1686416585"
		*/
		anyTimestampNative1 := data["f_timestamp_native1"].(map[string]interface{})
		require.Equal(t, float64(1686416585), anyTimestampNative1["seconds"].(float64))

		anyTimestampTimeLayout1 := data["f_timestamp_time_layout1"].(string)
		require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout1)

		anyTimestampUnixNano1 := data["f_timestamp_unix_nano1"].(float64)
		require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano1)

		anyTimestampUnixNanoString1 := data["f_timestamp_unix_nano_string1"].(string)
		require.Equal(t, "1686416585000000000", anyTimestampUnixNanoString1)

		anyTimestampUnixMicro1 := data["f_timestamp_unix_micro1"].(float64)
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1)

		anyTimestampUnixMicroString1 := data["f_timestamp_unix_micro_string1"].(string)
		require.Equal(t, "1686416585000000", anyTimestampUnixMicroString1)

		anyTimestampUnixMilli1 := data["f_timestamp_unix_milli1"].(float64)
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1)

		anyTimestampUnixMilliString1 := data["f_timestamp_unix_milli_string1"].(string)
		require.Equal(t, "1686416585000", anyTimestampUnixMilliString1)

		anyTimestampUnixSec1 := data["f_timestamp_unix_sec1"].(float64)
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1)

		anyTimestampUnixSecString1 := data["f_timestamp_unix_sec_string1"].(string)
		require.Equal(t, "1686416585", anyTimestampUnixSecString1)
	})
}
