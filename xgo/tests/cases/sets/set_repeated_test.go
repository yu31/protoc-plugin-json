package sets

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbsets"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedRepeated1 = &pbsets.TypeSetRepeated1{
	FEnumNumber1: []pbsets.Enum1{0, 2, 6},
	FEnumString1: []pbsets.Enum1{3, 5, 7},
	FAnyNative1: []*any.Any{
		utils.MustNewAny(&pbsets.Message1{FString1: "s111", FString2: "s112", FString3: "s113"}),
		utils.MustNewAny(&pbsets.Message1{FString1: "s121", FString2: "s122", FString3: "s123"}),
		utils.MustNewAny(&pbsets.Message1{FString1: "s131", FString2: "s132", FString3: "s133"}),
	},
	FAnyProto1: []*any.Any{
		utils.MustNewAny(&pbsets.Message1{FString1: "s311", FString2: "s312", FString3: "s313"}),
		utils.MustNewAny(&pbsets.Message1{FString1: "s321", FString2: "s322", FString3: "s323"}),
		utils.MustNewAny(&pbsets.Message1{FString1: "s331", FString2: "s332", FString3: "s333"}),
	},
	FDurationNative1: []*durationpb.Duration{
		{Seconds: 111, Nanos: 0}, {Seconds: 112, Nanos: 0}, {Seconds: 113, Nanos: 0},
	},
	FDurationString1: []*durationpb.Duration{
		{Seconds: 121, Nanos: 0}, {Seconds: 122, Nanos: 0}, {Seconds: 123, Nanos: 0},
	},
	FDurationNanoseconds1: []*durationpb.Duration{
		{Seconds: 131, Nanos: 0}, {Seconds: 132, Nanos: 0}, {Seconds: 133, Nanos: 0},
	},
	FDurationMicroseconds1: []*durationpb.Duration{
		{Seconds: 141, Nanos: 0}, {Seconds: 142, Nanos: 0}, {Seconds: 143, Nanos: 0},
	},
	FDurationMilliseconds1: []*durationpb.Duration{
		{Seconds: 151, Nanos: 0}, {Seconds: 152, Nanos: 0}, {Seconds: 153, Nanos: 0},
	},
	FDurationSeconds1: []*durationpb.Duration{
		{Seconds: 161, Nanos: 0}, {Seconds: 162, Nanos: 0}, {Seconds: 163, Nanos: 0},
	},
	FDurationMinutes1: []*durationpb.Duration{
		{Seconds: 360, Nanos: 0}, {Seconds: 720, Nanos: 0}, {Seconds: 1080, Nanos: 0},
	},
	FDurationHours1: []*durationpb.Duration{
		{Seconds: 3600, Nanos: 0}, {Seconds: 7200, Nanos: 0}, {Seconds: 10800, Nanos: 0},
	},
	FTimestampNative1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampTimeLayout1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixNano1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMicro1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMilli1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixSec1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
}

func Test_Set_Repeated1_General(t *testing.T) {
	bb, err := seedRepeated1.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbsets.TypeSetRepeated1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedRepeated1, dataNew)
}

func Test_Set_Repeated1_Check(t *testing.T) {
	bb, err := seedRepeated1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("Enum", func(t *testing.T) {
		/*
			"f_enum_number1":[0,2,6],"
			f_enum_string1":["Two","Three","Five"],
		*/
		enumNumber1 := data["f_enum_number1"].([]interface{})
		require.Equal(t, float64(0), enumNumber1[0].(float64))
		require.Equal(t, float64(2), enumNumber1[1].(float64))
		require.Equal(t, float64(6), enumNumber1[2].(float64))

		enumString1 := data["f_enum_string1"].([]interface{})
		require.Equal(t, "Two", enumString1[0].(string))
		require.Equal(t, "Three", enumString1[1].(string))
		require.Equal(t, "Five", enumString1[2].(string))
	})

	t.Run("Any", func(t *testing.T) {
		/*
			"f_any_native1":[
			{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTExEgRzMTEyGgRzMTEz"},
			{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTIxEgRzMTIyGgRzMTIz"},
			{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTMxEgRzMTMyGgRzMTMz"}
			],
			"f_any_proto1":[
			{"@type":"type.googleapis.com/set_common.Message1","f_string1":"s311","f_string2":"s312","f_string3":"s313"},
			{"@type":"type.googleapis.com/set_common.Message1","f_string1":"s321","f_string2":"s322","f_string3":"s323"},
			{"@type":"type.googleapis.com/set_common.Message1","f_string1":"s331","f_string2":"s332","f_string3":"s333"}
			],
		*/

		// map[string]interface{}
		anyNative1 := data["f_any_native1"].([]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1[0].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTExEgRzMTEyGgRzMTEz", anyNative1[0].(map[string]interface{})["value"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1[1].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTIxEgRzMTIyGgRzMTIz", anyNative1[1].(map[string]interface{})["value"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1[2].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTMxEgRzMTMyGgRzMTMz", anyNative1[2].(map[string]interface{})["value"].(string))

		anyProto1 := data["f_any_proto1"].([]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1[0].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s311", anyProto1[0].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s312", anyProto1[0].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s313", anyProto1[0].(map[string]interface{})["f_string3"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1[1].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s321", anyProto1[1].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s322", anyProto1[1].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s323", anyProto1[1].(map[string]interface{})["f_string3"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1[2].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s331", anyProto1[2].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s332", anyProto1[2].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s333", anyProto1[2].(map[string]interface{})["f_string3"].(string))
	})

	t.Run("Duration", func(t *testing.T) {
		/*
			"f_duration_native1":[{"seconds":111},{"seconds":112},{"seconds":113}],
			"f_duration_string1":["2m1s","2m2s","2m3s"],
			"f_duration_nanoseconds1":[131000000000,132000000000,133000000000],
			"f_duration_microseconds1":[141000000,142000000,143000000],
			"f_duration_milliseconds1":[151000,152000,153000],
			"f_duration_seconds1":[161,162,163],
			"f_duration_minutes1":[6,12,18],
			"f_duration_hours1":[1,2,3],
		*/

		anyDurationNative1 := data["f_duration_native1"].([]interface{})
		require.Equal(t, float64(111), anyDurationNative1[0].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(112), anyDurationNative1[1].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(113), anyDurationNative1[2].(map[string]interface{})["seconds"].(float64))

		anyDurationString1 := data["f_duration_string1"].([]interface{})
		require.Equal(t, "2m1s", anyDurationString1[0].(string))
		require.Equal(t, "2m2s", anyDurationString1[1].(string))
		require.Equal(t, "2m3s", anyDurationString1[2].(string))

		anyDurationNanoseconds1 := data["f_duration_nanoseconds1"].([]interface{})
		require.Equal(t, float64(131000000000), anyDurationNanoseconds1[0].(float64))
		require.Equal(t, float64(132000000000), anyDurationNanoseconds1[1].(float64))
		require.Equal(t, float64(133000000000), anyDurationNanoseconds1[2].(float64))

		anyDurationMicroseconds1 := data["f_duration_microseconds1"].([]interface{})
		require.Equal(t, float64(141000000), anyDurationMicroseconds1[0].(float64))
		require.Equal(t, float64(142000000), anyDurationMicroseconds1[1].(float64))
		require.Equal(t, float64(143000000), anyDurationMicroseconds1[2].(float64))

		anyDurationMilliseconds1 := data["f_duration_milliseconds1"].([]interface{})
		require.Equal(t, float64(151000), anyDurationMilliseconds1[0].(float64))
		require.Equal(t, float64(152000), anyDurationMilliseconds1[1].(float64))
		require.Equal(t, float64(153000), anyDurationMilliseconds1[2].(float64))

		anyDurationSeconds1 := data["f_duration_seconds1"].([]interface{})
		require.Equal(t, float64(161), anyDurationSeconds1[0].(float64))
		require.Equal(t, float64(162), anyDurationSeconds1[1].(float64))
		require.Equal(t, float64(163), anyDurationSeconds1[2].(float64))

		anyDurationMinutes1 := data["f_duration_minutes1"].([]interface{})
		require.Equal(t, float64(6), anyDurationMinutes1[0].(float64))
		require.Equal(t, float64(12), anyDurationMinutes1[1].(float64))
		require.Equal(t, float64(18), anyDurationMinutes1[2].(float64))

		anyDurationHours1 := data["f_duration_hours1"].([]interface{})
		require.Equal(t, float64(1), anyDurationHours1[0].(float64))
		require.Equal(t, float64(2), anyDurationHours1[1].(float64))
		require.Equal(t, float64(3), anyDurationHours1[2].(float64))
	})

	t.Run("Timestamp", func(t *testing.T) {
		/*
			"f_timestamp_native1":[{"seconds":1686416585},{"seconds":1686416585}],
			"f_timestamp_time_layout1":["2023-06-10T17:03:05Z","2023-06-10T17:03:05Z"],
			"f_timestamp_unix_nano1":[1686416585000000000,1686416585000000000],
			"f_timestamp_unix_micro1":[1686416585000000,1686416585000000],
			"f_timestamp_unix_milli1":[1686416585000,1686416585000],
			"f_timestamp_unix_sec1":[1686416585,1686416585]
		*/

		anyTimestampNative1 := data["f_timestamp_native1"].([]interface{})
		require.Equal(t, float64(1686416585), anyTimestampNative1[0].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(1686416585), anyTimestampNative1[1].(map[string]interface{})["seconds"].(float64))

		anyTimestampTimeLayout1 := data["f_timestamp_time_layout1"].([]interface{})
		require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout1[0].(string))
		require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout1[1].(string))

		anyTimestampUnixNano1 := data["f_timestamp_unix_nano1"].([]interface{})
		require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano1[0].(float64))
		require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano1[1].(float64))

		anyTimestampUnixMicro1 := data["f_timestamp_unix_micro1"].([]interface{})
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1[0].(float64))
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1[1].(float64))

		anyTimestampUnixMilli1 := data["f_timestamp_unix_milli1"].([]interface{})
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1[0].(float64))
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1[1].(float64))

		anyTimestampUnixSec1 := data["f_timestamp_unix_sec1"].([]interface{})
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1[0].(float64))
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1[1].(float64))
	})
}
