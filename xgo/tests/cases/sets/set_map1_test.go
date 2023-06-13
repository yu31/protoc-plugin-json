package sets

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbsets"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedMap1 = &pbsets.TypeSetMap1{
	FEnumNumber1: map[string]pbsets.Enum1{
		"ke11": 0, "ke12": 2, "ke13": 6,
	},
	FEnumString1: map[string]pbsets.Enum1{
		"ke21": 3, "ke22": 5, "ke23": 7,
	},
	FAnyNative1: map[string]*any.Any{
		"ka11": utils.MustNewAny(&pbsets.Message1{FString1: "s111", FString2: "s112", FString3: "s113"}),
		"ka12": utils.MustNewAny(&pbsets.Message1{FString1: "s121", FString2: "s122", FString3: "s123"}),
		"ka13": utils.MustNewAny(&pbsets.Message1{FString1: "s131", FString2: "s132", FString3: "s133"}),
	},
	FAnyProto1: map[string]*any.Any{
		"ka21": utils.MustNewAny(&pbsets.Message1{FString1: "s311", FString2: "s312", FString3: "s313"}),
		"ka22": utils.MustNewAny(&pbsets.Message1{FString1: "s321", FString2: "s322", FString3: "s323"}),
		"ka23": utils.MustNewAny(&pbsets.Message1{FString1: "s331", FString2: "s332", FString3: "s333"}),
	},
	FDurationNative1: map[string]*duration.Duration{
		"kd11": {Seconds: 111, Nanos: 0}, "kd12": {Seconds: 112, Nanos: 0}, "kd13": {Seconds: 113, Nanos: 0},
	},
	FDurationString1: map[string]*duration.Duration{
		"kd21": {Seconds: 121, Nanos: 0}, "kd22": {Seconds: 122, Nanos: 0}, "kd23": {Seconds: 123, Nanos: 0},
	},
	FDurationNanoseconds1: map[string]*duration.Duration{
		"kd31": {Seconds: 131, Nanos: 0}, "kd32": {Seconds: 132, Nanos: 0}, "kd33": {Seconds: 133, Nanos: 0},
	},
	FDurationMicroseconds1: map[string]*duration.Duration{
		"kd41": {Seconds: 141, Nanos: 0}, "kd42": {Seconds: 142, Nanos: 0}, "kd43": {Seconds: 143, Nanos: 0},
	},
	FDurationMilliseconds1: map[string]*duration.Duration{
		"kd51": {Seconds: 151, Nanos: 0}, "kd52": {Seconds: 152, Nanos: 0}, "kd53": {Seconds: 153, Nanos: 0},
	},
	FDurationSeconds1: map[string]*duration.Duration{
		"kd61": {Seconds: 161, Nanos: 0}, "kd62": {Seconds: 162, Nanos: 0}, "kd63": {Seconds: 163, Nanos: 0},
	},
	FDurationMinutes1: map[string]*duration.Duration{
		"kd71": {Seconds: 360, Nanos: 0}, "kd72": {Seconds: 720, Nanos: 0}, "kd73": {Seconds: 1080, Nanos: 0},
	},
	FDurationHours1: map[string]*duration.Duration{
		"kd81": {Seconds: 3600, Nanos: 0}, "kd82": {Seconds: 7200, Nanos: 0}, "kd83": {Seconds: 10800, Nanos: 0},
	},
	FTimestampNative1: map[string]*timestamp.Timestamp{
		"kt11": {Seconds: 1686416585, Nanos: 0}, "kt12": {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampTimeLayout1: map[string]*timestamp.Timestamp{
		"kt21": {Seconds: 1686416585, Nanos: 0}, "kt22": {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixNano1: map[string]*timestamp.Timestamp{
		"kt31": {Seconds: 1686416585, Nanos: 0}, "kt32": {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMicro1: map[string]*timestamp.Timestamp{
		"kt41": {Seconds: 1686416585, Nanos: 0}, "kt42": {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixMilli1: map[string]*timestamp.Timestamp{
		"kt51": {Seconds: 1686416585, Nanos: 0}, "kt52": {Seconds: 1686416585, Nanos: 0},
	},
	FTimestampUnixSec1: map[string]*timestamp.Timestamp{
		"kt61": {Seconds: 1686416585, Nanos: 0}, "kt62": {Seconds: 1686416585, Nanos: 0},
	},
}

func Test_Set_Map1_General(t *testing.T) {
	bb, err := seedMap1.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbsets.TypeSetMap1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedMap1, dataNew)
}

func Test_Set_Map1_Check(t *testing.T) {
	bb, err := seedMap1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("Enum", func(t *testing.T) {
		/*
			"f_enum_number1":{"ke13":6,"ke11":0,"ke12":2},
			"f_enum_string1":{"ke21":"Two","ke22":"Three","ke23":"Five"},
		*/
		enumNumber1 := data["f_enum_number1"].(map[string]interface{})
		require.Equal(t, float64(0), enumNumber1["ke11"].(float64))
		require.Equal(t, float64(2), enumNumber1["ke12"].(float64))
		require.Equal(t, float64(6), enumNumber1["ke13"].(float64))

		enumString1 := data["f_enum_string1"].(map[string]interface{})
		require.Equal(t, "Two", enumString1["ke21"].(string))
		require.Equal(t, "Three", enumString1["ke22"].(string))
		require.Equal(t, "Five", enumString1["ke23"].(string))
	})

	t.Run("Any", func(t *testing.T) {
		/*
			"f_any_native1":{
				"ka11":{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTExEgRzMTEyGgRzMTEz"},
				"ka12":{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTIxEgRzMTIyGgRzMTIz"},
				"ka13":{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTMxEgRzMTMyGgRzMTMz"}
			},
			"f_any_proto1":{
				"ka21":{"@type":"type.googleapis.com/set_common.Message1", "f_string1":"s311", "f_string2":"s312", "f_string3":"s313"},
				"ka22":{"@type":"type.googleapis.com/set_common.Message1", "f_string1":"s321", "f_string2":"s322", "f_string3":"s323"},
				"ka23":{"@type":"type.googleapis.com/set_common.Message1", "f_string1":"s331", "f_string2":"s332", "f_string3":"s333"}
			},
		*/

		anyNative1 := data["f_any_native1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1["ka11"].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTExEgRzMTEyGgRzMTEz", anyNative1["ka11"].(map[string]interface{})["value"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1["ka12"].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTIxEgRzMTIyGgRzMTIz", anyNative1["ka12"].(map[string]interface{})["value"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1["ka13"].(map[string]interface{})["type_url"].(string))
		require.Equal(t, "CgRzMTMxEgRzMTMyGgRzMTMz", anyNative1["ka13"].(map[string]interface{})["value"].(string))

		anyProto1 := data["f_any_proto1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1["ka21"].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s311", anyProto1["ka21"].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s312", anyProto1["ka21"].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s313", anyProto1["ka21"].(map[string]interface{})["f_string3"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1["ka22"].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s321", anyProto1["ka22"].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s322", anyProto1["ka22"].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s323", anyProto1["ka22"].(map[string]interface{})["f_string3"].(string))

		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1["ka23"].(map[string]interface{})["@type"].(string))
		require.Equal(t, "s331", anyProto1["ka23"].(map[string]interface{})["f_string1"].(string))
		require.Equal(t, "s332", anyProto1["ka23"].(map[string]interface{})["f_string2"].(string))
		require.Equal(t, "s333", anyProto1["ka23"].(map[string]interface{})["f_string3"].(string))
	})

	t.Run("Duration", func(t *testing.T) {
		/*
			"f_duration_native1":{"kd11":{"seconds":111},"kd12":{"seconds":112},"kd13":{"seconds":113}},
			"f_duration_string1":{"kd21":"2m1s","kd22":"2m2s","kd23":"2m3s"},
			"f_duration_nanoseconds1":{"kd31":131000000000,"kd32":132000000000,"kd33":133000000000},
			"f_duration_microseconds1":{"kd41":141000000,"kd42":142000000,"kd43":143000000},
			"f_duration_milliseconds1":{"kd51":151000,"kd52":152000,"kd53":153000},
			"f_duration_seconds1":{"kd61":161,"kd62":162,"kd63":163},"f_duration_minutes1":{"kd73":18,"kd71":6,"kd72":12},
			"f_duration_hours1":{"kd81":1,"kd82":2,"kd83":3},
		*/

		anyDurationNative1 := data["f_duration_native1"].(map[string]interface{})
		require.Equal(t, float64(111), anyDurationNative1["kd11"].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(112), anyDurationNative1["kd12"].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(113), anyDurationNative1["kd13"].(map[string]interface{})["seconds"].(float64))

		anyDurationString1 := data["f_duration_string1"].(map[string]interface{})
		require.Equal(t, "2m1s", anyDurationString1["kd21"].(string))
		require.Equal(t, "2m2s", anyDurationString1["kd22"].(string))
		require.Equal(t, "2m3s", anyDurationString1["kd23"].(string))

		anyDurationNanoseconds1 := data["f_duration_nanoseconds1"].(map[string]interface{})
		require.Equal(t, float64(131000000000), anyDurationNanoseconds1["kd31"].(float64))
		require.Equal(t, float64(132000000000), anyDurationNanoseconds1["kd32"].(float64))
		require.Equal(t, float64(133000000000), anyDurationNanoseconds1["kd33"].(float64))

		anyDurationMicroseconds1 := data["f_duration_microseconds1"].(map[string]interface{})
		require.Equal(t, float64(141000000), anyDurationMicroseconds1["kd41"].(float64))
		require.Equal(t, float64(142000000), anyDurationMicroseconds1["kd42"].(float64))
		require.Equal(t, float64(143000000), anyDurationMicroseconds1["kd43"].(float64))

		anyDurationMilliseconds1 := data["f_duration_milliseconds1"].(map[string]interface{})
		require.Equal(t, float64(151000), anyDurationMilliseconds1["kd51"].(float64))
		require.Equal(t, float64(152000), anyDurationMilliseconds1["kd52"].(float64))
		require.Equal(t, float64(153000), anyDurationMilliseconds1["kd53"].(float64))

		anyDurationSeconds1 := data["f_duration_seconds1"].(map[string]interface{})
		require.Equal(t, float64(161), anyDurationSeconds1["kd61"].(float64))
		require.Equal(t, float64(162), anyDurationSeconds1["kd62"].(float64))
		require.Equal(t, float64(163), anyDurationSeconds1["kd63"].(float64))

		anyDurationMinutes1 := data["f_duration_minutes1"].(map[string]interface{})
		require.Equal(t, float64(6), anyDurationMinutes1["kd71"].(float64))
		require.Equal(t, float64(12), anyDurationMinutes1["kd72"].(float64))
		require.Equal(t, float64(18), anyDurationMinutes1["kd73"].(float64))

		anyDurationHours1 := data["f_duration_hours1"].(map[string]interface{})
		require.Equal(t, float64(1), anyDurationHours1["kd81"].(float64))
		require.Equal(t, float64(2), anyDurationHours1["kd82"].(float64))
		require.Equal(t, float64(3), anyDurationHours1["kd83"].(float64))
	})

	t.Run("Timestamp", func(t *testing.T) {
		/*
			"f_timestamp_native1":{"kt11":{"seconds":1686416585},"kt12":{"seconds":1686416585}},
			"f_timestamp_time_layout1":{"kt21":"2023-06-10T17:03:05Z","kt22":"2023-06-10T17:03:05Z"},
			"f_timestamp_unix_nano1":{"kt31":1686416585000000000,"kt32":1686416585000000000},
			"f_timestamp_unix_micro1":{"kt42":1686416585000000,"kt41":1686416585000000},
			"f_timestamp_unix_milli1":{"kt51":1686416585000,"kt52":1686416585000},
			"f_timestamp_unix_sec1":{"kt62":1686416585,"kt61":1686416585}
		*/

		anyTimestampNative1 := data["f_timestamp_native1"].(map[string]interface{})
		require.Equal(t, float64(1686416585), anyTimestampNative1["kt11"].(map[string]interface{})["seconds"].(float64))
		require.Equal(t, float64(1686416585), anyTimestampNative1["kt12"].(map[string]interface{})["seconds"].(float64))

		anyTimestampTimeLayout1 := data["f_timestamp_time_layout1"].(map[string]interface{})
		require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout1["kt21"].(string))
		require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout1["kt22"].(string))

		anyTimestampUnixNano1 := data["f_timestamp_unix_nano1"].(map[string]interface{})
		require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano1["kt31"].(float64))
		require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano1["kt32"].(float64))

		anyTimestampUnixMicro1 := data["f_timestamp_unix_micro1"].(map[string]interface{})
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1["kt41"].(float64))
		require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro1["kt42"].(float64))

		anyTimestampUnixMilli1 := data["f_timestamp_unix_milli1"].(map[string]interface{})
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1["kt51"].(float64))
		require.Equal(t, float64(1686416585000), anyTimestampUnixMilli1["kt52"].(float64))

		anyTimestampUnixSec1 := data["f_timestamp_unix_sec1"].(map[string]interface{})
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1["kt61"].(float64))
		require.Equal(t, float64(1686416585), anyTimestampUnixSec1["kt62"].(float64))
	})
}
