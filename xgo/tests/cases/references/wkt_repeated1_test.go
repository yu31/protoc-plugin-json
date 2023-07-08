package references

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"

	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedWKTRepeated1 = &pbref.WKTRepeated1{
	FEnumNumber1:       []pbref.Enum1{0, 2, 6},
	FEnumNumberString1: []pbref.Enum1{0, 2, 6},
	FEnumString1:       []pbref.Enum1{3, 5, 7},
	FAnyNative1: []*any.Any{
		utils.MustNewAny(&pbref.Message1{FString1: "s111", FString2: "s112", FString3: "s113"}),
		utils.MustNewAny(&pbref.Message1{FString1: "s121", FString2: "s122", FString3: "s123"}),
		utils.MustNewAny(&pbref.Message1{FString1: "s131", FString2: "s132", FString3: "s133"}),
	},
	FAnyProto1: []*any.Any{
		utils.MustNewAny(&pbref.Message1{FString1: "s311", FString2: "s312", FString3: "s313"}),
		utils.MustNewAny(&pbref.Message1{FString1: "s321", FString2: "s322", FString3: "s323"}),
		utils.MustNewAny(&pbref.Message1{FString1: "s331", FString2: "s332", FString3: "s333"}),
	},
	FDurationNative1: []*durationpb.Duration{
		{Seconds: 111, Nanos: 0}, {Seconds: 112, Nanos: 0}, {Seconds: 113, Nanos: 0},
	},
	FDurationString1: []*durationpb.Duration{
		{Seconds: 121, Nanos: 0}, {Seconds: 122, Nanos: 0}, {Seconds: 123, Nanos: 0},
	},
	FDurationNanosecond1: []*durationpb.Duration{
		{Seconds: 131, Nanos: 0}, {Seconds: 132, Nanos: 0}, {Seconds: 133, Nanos: 0},
	},
	FDurationNanosecondString1: []*durationpb.Duration{
		{Seconds: 131, Nanos: 0}, {Seconds: 132, Nanos: 0}, {Seconds: 133, Nanos: 0},
	},
	FDurationMicrosecond1: []*durationpb.Duration{
		{Seconds: 141, Nanos: 0}, {Seconds: 142, Nanos: 0}, {Seconds: 143, Nanos: 0},
	},
	FDurationMicrosecondString1: []*durationpb.Duration{
		{Seconds: 141, Nanos: 0}, {Seconds: 142, Nanos: 0}, {Seconds: 143, Nanos: 0},
	},
	FDurationMillisecond1: []*durationpb.Duration{
		{Seconds: 151, Nanos: 0}, {Seconds: 152, Nanos: 0}, {Seconds: 153, Nanos: 0},
	},
	FDurationMillisecondString1: []*durationpb.Duration{
		{Seconds: 151, Nanos: 0}, {Seconds: 152, Nanos: 0}, {Seconds: 153, Nanos: 0},
	},
	FDurationSecond1: []*durationpb.Duration{
		{Seconds: 161, Nanos: 0}, {Seconds: 162, Nanos: 0}, {Seconds: 163, Nanos: 0},
	},
	FDurationSecondString1: []*durationpb.Duration{
		{Seconds: 161, Nanos: 0}, {Seconds: 162, Nanos: 0}, {Seconds: 163, Nanos: 0},
	},
	FDurationMinute1: []*durationpb.Duration{
		{Seconds: 360, Nanos: 0}, {Seconds: 720, Nanos: 0}, {Seconds: 1080, Nanos: 0},
	},
	FDurationMinuteString1: []*durationpb.Duration{
		{Seconds: 360, Nanos: 0}, {Seconds: 720, Nanos: 0}, {Seconds: 1080, Nanos: 0},
	},
	FDurationHour1: []*durationpb.Duration{
		{Seconds: 3600, Nanos: 0}, {Seconds: 7200, Nanos: 0}, {Seconds: 10800, Nanos: 0},
	},
	FDurationHourString1: []*durationpb.Duration{
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
	FTimestampUnixNanoString1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMicro1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMicroString1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMilli1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMilliString1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixSec1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixSecString1: []*timestamppb.Timestamp{
		{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
	},
}

func Test_Reference_WKT_Repeated1_General(t *testing.T) {
	bb, err := seedWKTRepeated1.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbref.WKTRepeated1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedWKTRepeated1, dataNew)
}

func Test_Reference_WKT_Repeated1_Check(t *testing.T) {
	bb, err := seedWKTRepeated1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("Enum", func(t *testing.T) {
		/*
			"f_enum_number1":[0,2,6],"
			"f_enum_number_string1":["0","2","6"],
			"f_enum_string1":["Two","Three","Five"],
		*/
		enumNumber1 := data["f_enum_number1"].([]interface{})
		require.Equal(t, float64(0), enumNumber1[0].(float64))
		require.Equal(t, float64(2), enumNumber1[1].(float64))
		require.Equal(t, float64(6), enumNumber1[2].(float64))

		enumNumberString1 := data["f_enum_number_string1"].([]interface{})
		require.Equal(t, "0", enumNumberString1[0].(string))
		require.Equal(t, "2", enumNumberString1[1].(string))
		require.Equal(t, "6", enumNumberString1[2].(string))

		enumString1 := data["f_enum_string1"].([]interface{})
		require.Equal(t, "Two", enumString1[0].(string))
		require.Equal(t, "Three", enumString1[1].(string))
		require.Equal(t, "Five", enumString1[2].(string))
	})

	t.Run("Any", func(t *testing.T) {
		/*
			"f_any_native1":[
			{"type_url":"type.googleapis.com/private.Message1","value":"CgRzMTExEgRzMTEyGgRzMTEz"},
			{"type_url":"type.googleapis.com/private.Message1","value":"CgRzMTIxEgRzMTIyGgRzMTIz"},
			{"type_url":"type.googleapis.com/private.Message1","value":"CgRzMTMxEgRzMTMyGgRzMTMz"}
			],
			"f_any_proto1":[
			{"@type":"type.googleapis.com/private.Message1","f_string1":"s311","f_string2":"s312","f_string3":"s313"},
			{"@type":"type.googleapis.com/private.Message1","f_string1":"s321","f_string2":"s322","f_string3":"s323"},
			{"@type":"type.googleapis.com/private.Message1","f_string1":"s331","f_string2":"s332","f_string3":"s333"}
			],
		*/

		// map[string]interface{}
		anyNative1 := data["f_any_native1"].([]interface{})
		require.Equal(t, "type.googleapis.com/private.Message1", anyNative1[0].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTExEgRzMTEyGgRzMTEz", anyNative1[0].(map[string]interface{})["value"].(string))

		require.Equal(t, "type.googleapis.com/private.Message1", anyNative1[1].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTIxEgRzMTIyGgRzMTIz", anyNative1[1].(map[string]interface{})["value"].(string))

		require.Equal(t, "type.googleapis.com/private.Message1", anyNative1[2].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTMxEgRzMTMyGgRzMTMz", anyNative1[2].(map[string]interface{})["value"].(string))

		anyProto1 := data["f_any_proto1"].([]interface{})
		require.Equal(t, "type.googleapis.com/private.Message1", anyProto1[0].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s311", anyProto1[0].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s312", anyProto1[0].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s313", anyProto1[0].(map[string]interface{})["f_string3"].(string))

		require.Equal(t, "type.googleapis.com/private.Message1", anyProto1[1].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s321", anyProto1[1].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s322", anyProto1[1].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s323", anyProto1[1].(map[string]interface{})["f_string3"].(string))

		require.Equal(t, "type.googleapis.com/private.Message1", anyProto1[2].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s331", anyProto1[2].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s332", anyProto1[2].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s333", anyProto1[2].(map[string]interface{})["f_string3"].(string))
	})

	t.Run("Duration", func(t *testing.T) {
		/*
			"f_duration_native1":[{"seconds":111},{"seconds":112},{"seconds":113}],
			"f_duration_string1":["2m1s","2m2s","2m3s"],
			"f_duration_nanosecond1":[131000000000,132000000000,133000000000],
			"f_duration_nanosecond_string1":["131000000000","132000000000","133000000000"],
			"f_duration_microsecond1":[141000000,142000000,143000000],
			"f_duration_microsecond_string1":["141000000","142000000","143000000"],
			"f_duration_millisecond1":[151000,152000,153000],
			"f_duration_millisecond_string1":["151000","152000","153000"],
			"f_duration_second1":[161,162,163],
			"f_duration_second_string1":["161","162","163"],
			"f_duration_minute1":[6,12,18],
			"f_duration_minute_string1":["6","12","18"],
			"f_duration_hour1":[1,2,3],
			"f_duration_hour_string1":["1","2","3"],
		*/

		anyDurationNative1 := data["f_duration_native1"].([]interface{})
		require.Equal(t, float64(111), anyDurationNative1[0].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(112), anyDurationNative1[1].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(113), anyDurationNative1[2].(map[string]interface{})["seconds"].(float64))

		anyDurationString1 := data["f_duration_string1"].([]interface{})
		require.Equal(t, "2m1s", anyDurationString1[0].(string))
		require.Equal(t, "2m2s", anyDurationString1[1].(string))
		require.Equal(t, "2m3s", anyDurationString1[2].(string))

		anyDurationNanosecond1 := data["f_duration_nanosecond1"].([]interface{})
		require.Equal(t, float64(131000000000), anyDurationNanosecond1[0].(float64))
		require.Equal(t, float64(132000000000), anyDurationNanosecond1[1].(float64))
		require.Equal(t, float64(133000000000), anyDurationNanosecond1[2].(float64))

		anyDurationNanosecondString1 := data["f_duration_nanosecond_string1"].([]interface{})
		require.Equal(t, "131000000000", anyDurationNanosecondString1[0].(string))
		require.Equal(t, "132000000000", anyDurationNanosecondString1[1].(string))
		require.Equal(t, "133000000000", anyDurationNanosecondString1[2].(string))

		anyDurationMicrosecond1 := data["f_duration_microsecond1"].([]interface{})
		require.Equal(t, float64(141000000), anyDurationMicrosecond1[0].(float64))
		require.Equal(t, float64(142000000), anyDurationMicrosecond1[1].(float64))
		require.Equal(t, float64(143000000), anyDurationMicrosecond1[2].(float64))

		anyDurationMicrosecondString1 := data["f_duration_microsecond_string1"].([]interface{})
		require.Equal(t, "141000000", anyDurationMicrosecondString1[0].(string))
		require.Equal(t, "142000000", anyDurationMicrosecondString1[1].(string))
		require.Equal(t, "143000000", anyDurationMicrosecondString1[2].(string))

		anyDurationMillisecond1 := data["f_duration_millisecond1"].([]interface{})
		require.Equal(t, float64(151000), anyDurationMillisecond1[0].(float64))
		require.Equal(t, float64(152000), anyDurationMillisecond1[1].(float64))
		require.Equal(t, float64(153000), anyDurationMillisecond1[2].(float64))

		anyDurationMillisecondString1 := data["f_duration_millisecond_string1"].([]interface{})
		require.Equal(t, "151000", anyDurationMillisecondString1[0].(string))
		require.Equal(t, "152000", anyDurationMillisecondString1[1].(string))
		require.Equal(t, "153000", anyDurationMillisecondString1[2].(string))

		anyDurationSecond1 := data["f_duration_second1"].([]interface{})
		require.Equal(t, float64(161), anyDurationSecond1[0].(float64))
		require.Equal(t, float64(162), anyDurationSecond1[1].(float64))
		require.Equal(t, float64(163), anyDurationSecond1[2].(float64))

		anyDurationSecondString1 := data["f_duration_second_string1"].([]interface{})
		require.Equal(t, "161", anyDurationSecondString1[0].(string))
		require.Equal(t, "162", anyDurationSecondString1[1].(string))
		require.Equal(t, "163", anyDurationSecondString1[2].(string))

		anyDurationMinute1 := data["f_duration_minute1"].([]interface{})
		require.Equal(t, float64(6), anyDurationMinute1[0].(float64))
		require.Equal(t, float64(12), anyDurationMinute1[1].(float64))
		require.Equal(t, float64(18), anyDurationMinute1[2].(float64))

		anyDurationMinuteString1 := data["f_duration_minute_string1"].([]interface{})
		require.Equal(t, "6", anyDurationMinuteString1[0].(string))
		require.Equal(t, "12", anyDurationMinuteString1[1].(string))
		require.Equal(t, "18", anyDurationMinuteString1[2].(string))

		anyDurationHour1 := data["f_duration_hour1"].([]interface{})
		require.Equal(t, float64(1), anyDurationHour1[0].(float64))
		require.Equal(t, float64(2), anyDurationHour1[1].(float64))
		require.Equal(t, float64(3), anyDurationHour1[2].(float64))

		anyDurationHourString1 := data["f_duration_hour_string1"].([]interface{})
		require.Equal(t, "1", anyDurationHourString1[0].(string))
		require.Equal(t, "2", anyDurationHourString1[1].(string))
		require.Equal(t, "3", anyDurationHourString1[2].(string))
	})

	t.Run("Timestamp", func(t *testing.T) {
		/*
			"f_timestamp_native1":[{"seconds":1686416585},{"seconds":1686416585}],
			"f_timestamp_time_layout1":["2023-06-10T17:03:05Z","2023-06-10T17:03:05Z"],
			"f_timestamp_unix_nano1":[1686416585000000000,1686416585000000000],
			"f_timestamp_unix_nano_string1":["1686416585000000000","1686416585000000000"],
			"f_timestamp_unix_micro1":[1686416585000000,1686416585000000],
			"f_timestamp_unix_micro_string1":["1686416585000000","1686416585000000"],
			"f_timestamp_unix_milli1":[1686416585000,1686416585000],
			"f_timestamp_unix_milli_string1":["1686416585000","1686416585000"],
			"f_timestamp_unix_sec1":[1686416585,1686416585],
			"f_timestamp_unix_sec_string1":["1686416585","1686416585"]
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

		anyTimestampUnixNanoString1 := data["f_timestamp_unix_nano_string1"].([]interface{})
		require.Equal(t, "1686416585000000000", anyTimestampUnixNanoString1[0].(string))
		require.Equal(t, "1686416585000000000", anyTimestampUnixNanoString1[1].(string))

		anyTimestampUnixMicro1 := data["f_timestamp_unix_micro1"].([]interface{})
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1[0].(float64))
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1[1].(float64))

		anyTimestampUnixMicroString1 := data["f_timestamp_unix_micro_string1"].([]interface{})
		require.Equal(t, "1686416585000000", anyTimestampUnixMicroString1[0].(string))
		require.Equal(t, "1686416585000000", anyTimestampUnixMicroString1[1].(string))

		anyTimestampUnixMilli1 := data["f_timestamp_unix_milli1"].([]interface{})
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1[0].(float64))
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1[1].(float64))

		anyTimestampUnixMilliString1 := data["f_timestamp_unix_milli_string1"].([]interface{})
		require.Equal(t, "1686416585000", anyTimestampUnixMilliString1[0].(string))
		require.Equal(t, "1686416585000", anyTimestampUnixMilliString1[1].(string))

		anyTimestampUnixSec1 := data["f_timestamp_unix_sec1"].([]interface{})
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1[0].(float64))
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1[1].(float64))

		anyTimestampUnixSecString1 := data["f_timestamp_unix_sec_string1"].([]interface{})
		require.Equal(t, "1686416585", anyTimestampUnixSecString1[0].(string))
		require.Equal(t, "1686416585", anyTimestampUnixSecString1[1].(string))
	})
}
