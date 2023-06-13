package sets

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbsets"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedOptional1 = &pbsets.TypeSetOptional1{
	FEnumNumber1:           pbsets.Enum1(2).Enum(),
	FEnumString1:           pbsets.Enum1(5).Enum(),
	FAnyNative1:            utils.MustNewAny(&pbsets.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
	FAnyProto1:             utils.MustNewAny(&pbsets.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
	FDurationNative1:       &durationpb.Duration{Seconds: 111, Nanos: 0},
	FDurationString1:       &durationpb.Duration{Seconds: 112, Nanos: 0},
	FDurationNanoseconds1:  &durationpb.Duration{Seconds: 113, Nanos: 0},
	FDurationMicroseconds1: &durationpb.Duration{Seconds: 114, Nanos: 0},
	FDurationMilliseconds1: &durationpb.Duration{Seconds: 115, Nanos: 0},
	FDurationSeconds1:      &durationpb.Duration{Seconds: 116, Nanos: 0},
	FDurationMinutes1:      &durationpb.Duration{Seconds: 117, Nanos: 0},
	FDurationHours1:        &durationpb.Duration{Seconds: 7200, Nanos: 0},
	FTimestampNative1:      &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampTimeLayout1:  &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixNano1:    &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixMicro1:   &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixMilli1:   &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	FTimestampUnixSec1:     &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
}

func Test_Set_Optional1_General(t *testing.T) {
	bb, err := seedOptional1.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbsets.TypeSetOptional1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedOptional1, dataNew)
}

func Test_Set_Optional1_Check(t *testing.T) {
	bb, err := seedOptional1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("Enum", func(t *testing.T) {
		/*
			"f_enum_number1":2,
			"f_enum_string1":"Three",
		*/
		enumNumber1 := data["f_enum_number1"].(float64)
		require.Equal(t, float64(2), enumNumber1)

		enumString1 := data["f_enum_string1"].(string)
		require.Equal(t, "Three", enumString1)
	})

	t.Run("Any", func(t *testing.T) {
		/*
			"f_any_native1":{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
			"f_any_proto1":{"@type":"type.googleapis.com/set_common.Message1","f_string1":"s301", "f_string2":"s302", "f_string3":"s303"},
		*/
		anyNative1 := data["f_any_native1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1["type_url"].(string))
		require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", anyNative1["value"].(string))

		anyProto1 := data["f_any_proto1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1["@type"].(string))
		require.Equal(t, "s301", anyProto1["f_string1"].(string))
		require.Equal(t, "s302", anyProto1["f_string2"].(string))
		require.Equal(t, "s303", anyProto1["f_string3"].(string))
	})

	t.Run("Duration", func(t *testing.T) {
		/*
			"f_duration_native1":{"seconds":111},
			"f_duration_string1":"1m52s",
			"f_duration_nanoseconds1":113000000000,
			"f_duration_microseconds1":114000000,
			"f_duration_milliseconds1":115000,
			"f_duration_seconds1":116,
			"f_duration_minutes1":1.95,
			"f_duration_hours1":2,
		*/

		anyDurationNative1 := data["f_duration_native1"].(map[string]interface{})
		require.Equal(t, float64(111), anyDurationNative1["seconds"].(float64))

		anyDurationString1 := data["f_duration_string1"].(string)
		require.Equal(t, "1m52s", anyDurationString1)

		anyDurationNanoseconds1 := data["f_duration_nanoseconds1"].(float64)
		require.Equal(t, float64(113000000000), anyDurationNanoseconds1)

		anyDurationMicroseconds1 := data["f_duration_microseconds1"].(float64)
		require.Equal(t, float64(114000000), anyDurationMicroseconds1)

		anyDurationMilliseconds1 := data["f_duration_milliseconds1"].(float64)
		require.Equal(t, float64(115000), anyDurationMilliseconds1)

		anyDurationSeconds1 := data["f_duration_seconds1"].(float64)
		require.Equal(t, float64(116), anyDurationSeconds1)

		anyDurationMinutes1 := data["f_duration_minutes1"].(float64)
		require.Equal(t, float64(1.95), anyDurationMinutes1)

		anyDurationHours1 := data["f_duration_hours1"].(float64)
		require.Equal(t, float64(2), anyDurationHours1)

		anyTimestampNative1 := data["f_timestamp_native1"].(map[string]interface{})
		require.Equal(t, float64(1686416585), anyTimestampNative1["seconds"].(float64))

		anyTimestampTimeLayout1 := data["f_timestamp_time_layout1"].(string)
		require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout1)
	})

	t.Run("Timestamp", func(t *testing.T) {
		/*
			"f_timestamp_native1":{"seconds":1686416585},
			"f_timestamp_time_layout1":"2023-06-10T17:03:05Z",
			"f_timestamp_unix_nano1":1686416585000000000,
			"f_timestamp_unix_micro1":1686416585000000,
			"f_timestamp_unix_milli1":1686416585000,
			"f_timestamp_unix_sec1":1686416585
		*/
		anyTimestampUnixNano1 := data["f_timestamp_unix_nano1"].(float64)
		require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano1)

		anyTimestampUnixMicro1 := data["f_timestamp_unix_micro1"].(float64)
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1)

		anyTimestampUnixMilli1 := data["f_timestamp_unix_milli1"].(float64)
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1)

		anyTimestampUnixSec1 := data["f_timestamp_unix_sec1"].(float64)
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1)
	})
}
