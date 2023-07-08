package references

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"

	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedWKTOneOf1 = &pbref.WKTOneOf1{
	OneEnum1: &pbref.WKTOneOf1_FEnumNumber1{FEnumNumber1: 2},
	OneEnum2: &pbref.WKTOneOf1_FEnumNumberString1{FEnumNumberString1: 2},
	OneAny1: &pbref.WKTOneOf1_FAnyNative1{
		FAnyNative1: utils.MustNewAny(&pbref.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
	},
	OneDuration1: &pbref.WKTOneOf1_FDurationNative1{
		FDurationNative1: &durationpb.Duration{Seconds: 111, Nanos: 0},
	},
	OneDuration2: &pbref.WKTOneOf1_FDurationNanosecond1{
		FDurationNanosecond1: &durationpb.Duration{Seconds: 113, Nanos: 0},
	},
	OneDuration3: &pbref.WKTOneOf1_FDurationMillisecond1{
		FDurationMillisecond1: &durationpb.Duration{Seconds: 115, Nanos: 0},
	},
	OneDuration4: &pbref.WKTOneOf1_FDurationMinute1{
		FDurationMinute1: &durationpb.Duration{Seconds: 117, Nanos: 0},
	},
	OneDuration5: &pbref.WKTOneOf1_FDurationNanosecondString1{
		FDurationNanosecondString1: &durationpb.Duration{Seconds: 113, Nanos: 0},
	},
	OneDuration6: &pbref.WKTOneOf1_FDurationMillisecondString1{
		FDurationMillisecondString1: &durationpb.Duration{Seconds: 115, Nanos: 0},
	},
	OneDuration7: &pbref.WKTOneOf1_FDurationMinuteString1{
		FDurationMinuteString1: &durationpb.Duration{Seconds: 117, Nanos: 0},
	},
	OneTimestamp1: &pbref.WKTOneOf1_FTimestampNative1{
		FTimestampNative1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp2: &pbref.WKTOneOf1_FTimestampUnixNano1{
		FTimestampUnixNano1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp3: &pbref.WKTOneOf1_FTimestampUnixMilli1{
		FTimestampUnixMilli1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp4: &pbref.WKTOneOf1_FTimestampUnixNanoString1{
		FTimestampUnixNanoString1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp5: &pbref.WKTOneOf1_FTimestampUnixMilliString1{
		FTimestampUnixMilliString1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
}

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedWKTOneOf2 = &pbref.WKTOneOf1{
	OneEnum1: &pbref.WKTOneOf1_FEnumString1{FEnumString1: 5},
	OneEnum2: &pbref.WKTOneOf1_FEnumNumberString2{FEnumNumberString2: 5},
	OneAny1: &pbref.WKTOneOf1_FAnyProto1{
		FAnyProto1: utils.MustNewAny(&pbref.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
	},
	OneDuration1: &pbref.WKTOneOf1_FDurationString1{
		FDurationString1: &durationpb.Duration{Seconds: 112, Nanos: 0},
	},
	OneDuration2: &pbref.WKTOneOf1_FDurationMicrosecond1{
		FDurationMicrosecond1: &durationpb.Duration{Seconds: 114, Nanos: 0},
	},
	OneDuration3: &pbref.WKTOneOf1_FDurationSecond1{
		FDurationSecond1: &durationpb.Duration{Seconds: 116, Nanos: 0},
	},
	OneDuration4: &pbref.WKTOneOf1_FDurationHour1{
		FDurationHour1: &durationpb.Duration{Seconds: 7200, Nanos: 0},
	},
	OneDuration5: &pbref.WKTOneOf1_FDurationMicrosecondString1{
		FDurationMicrosecondString1: &durationpb.Duration{Seconds: 114, Nanos: 0},
	},
	OneDuration6: &pbref.WKTOneOf1_FDurationSecondString1{
		FDurationSecondString1: &durationpb.Duration{Seconds: 116, Nanos: 0},
	},
	OneDuration7: &pbref.WKTOneOf1_FDurationHourString1{
		FDurationHourString1: &durationpb.Duration{Seconds: 7200, Nanos: 0},
	},
	OneTimestamp1: &pbref.WKTOneOf1_FTimestampTimeLayout1{
		FTimestampTimeLayout1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp2: &pbref.WKTOneOf1_FTimestampUnixMicro1{
		FTimestampUnixMicro1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp3: &pbref.WKTOneOf1_FTimestampUnixSec1{
		FTimestampUnixSec1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp4: &pbref.WKTOneOf1_FTimestampUnixMicroString1{
		FTimestampUnixMicroString1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp5: &pbref.WKTOneOf1_FTimestampUnixSecString1{
		FTimestampUnixSecString1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
}

func Test_Reference_WKT_OneOf1_General1(t *testing.T) {
	bb, err := seedWKTOneOf1.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbref.WKTOneOf1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedWKTOneOf1, dataNew)
}

func Test_Reference_WKT_OneOf1_General2(t *testing.T) {
	bb, err := seedWKTOneOf2.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbref.WKTOneOf1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedWKTOneOf2, dataNew)
}

func Test_Reference_WKT_OneOf1_Check1(t *testing.T) {
	bb, err := seedWKTOneOf1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)
	/*
		{
			"one_enum1":{"f_enum_number1":2},
			"f_enum_number_string1": "2",
			"f_any_native1":{"type_url":"type.googleapis.com/private.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
			"one_duration1":{"f_duration_native1":{"seconds":111}},
			"f_duration_nanosecond1":113000000000,
			"one_duration3":{"f_duration_millisecond1":115000},
			"f_duration_minute1":1.95,

			"one_duration5":{"f_duration_nanosecond_string1":"113000000000"},
			"f_duration_millisecond_string1":"115000",
			"one_duration7":{"f_duration_minute_string1":"1.95"},

			"one_timestamp1":{"f_timestamp_native1":{"seconds":1686416585}},
			"f_timestamp_unix_nano1":1686416585000000000,
			"one_timestamp3":{"f_timestamp_unix_milli1":1686416585000},
			"f_timestamp_unix_nano_string1":"1686416585000000000",
			"one_timestamp5":{"f_timestamp_unix_milli_string1": "1686416585000"},
		}
	*/
	t.Run("OneEnum1", func(t *testing.T) {
		oneEnum1 := data["one_enum1"].(map[string]interface{})
		enumNumber1 := oneEnum1["f_enum_number1"].(float64)
		require.Equal(t, float64(2), enumNumber1)
	})
	t.Run("OneEnum2", func(t *testing.T) {
		oneEnum2 := data["f_enum_number_string1"].(string)
		require.Equal(t, "2", oneEnum2)
	})

	t.Run("OneAny1", func(t *testing.T) {
		anyNative1 := data["f_any_native1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/private.Message1", anyNative1["type_url"].(string))
		require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", anyNative1["value"].(string))
	})

	t.Run("OneDuration1", func(t *testing.T) {
		oneDuration1 := data["one_duration1"].(map[string]interface{})
		durationNative1 := oneDuration1["f_duration_native1"].(map[string]interface{})
		require.Equal(t, float64(111), durationNative1["seconds"].(float64))
	})
	t.Run("OneDuration2", func(t *testing.T) {
		durationNanosecond1 := data["f_duration_nanosecond1"].(float64)
		require.Equal(t, float64(113000000000), durationNanosecond1)
	})
	t.Run("OneDuration3", func(t *testing.T) {
		oneDuration3 := data["one_duration3"].(map[string]interface{})
		durationMillisecond1 := oneDuration3["f_duration_millisecond1"].(float64)
		require.Equal(t, float64(115000), durationMillisecond1)
	})
	t.Run("OneDuration4", func(t *testing.T) {
		durationMinute1 := data["f_duration_minute1"].(float64)
		require.Equal(t, float64(1.95), durationMinute1)
	})
	t.Run("OneDuration5", func(t *testing.T) {
		oneDuration5 := data["one_duration5"].(map[string]interface{})
		durationNanosecondString1 := oneDuration5["f_duration_nanosecond_string1"].(string)
		require.Equal(t, "113000000000", durationNanosecondString1)
	})
	t.Run("OneDuration6", func(t *testing.T) {
		durationMillisecondString1 := data["f_duration_millisecond_string1"].(string)
		require.Equal(t, "115000", durationMillisecondString1)
	})
	t.Run("OneDuration7", func(t *testing.T) {
		oneDuration7 := data["one_duration7"].(map[string]interface{})
		durationMinuteString1 := oneDuration7["f_duration_minute_string1"].(string)
		require.Equal(t, "1.95", durationMinuteString1)
	})

	t.Run("OneTimestamp1", func(t *testing.T) {
		oneTimestamp1 := data["one_timestamp1"].(map[string]interface{})
		timestampNative1 := oneTimestamp1["f_timestamp_native1"].(map[string]interface{})
		require.Equal(t, float64(1686416585), timestampNative1["seconds"].(float64))
	})
	t.Run("OneTimestamp2", func(t *testing.T) {
		timestampUnixNano1 := data["f_timestamp_unix_nano1"].(float64)
		require.Equal(t, float64(1686416585000000000), timestampUnixNano1)
	})
	t.Run("OneTimestamp3", func(t *testing.T) {
		oneTimestamp3 := data["one_timestamp3"].(map[string]interface{})
		timestampUnixMilli1 := oneTimestamp3["f_timestamp_unix_milli1"].(float64)
		require.Equal(t, float64(1686416585000), timestampUnixMilli1)
	})
	t.Run("OneTimestamp4", func(t *testing.T) {
		timestampUnixNanoString1 := data["f_timestamp_unix_nano_string1"].(string)
		require.Equal(t, "1686416585000000000", timestampUnixNanoString1)
	})
	t.Run("OneTimestamp5", func(t *testing.T) {
		oneTimestamp3 := data["one_timestamp5"].(map[string]interface{})
		timestampUnixMilliString1 := oneTimestamp3["f_timestamp_unix_milli_string1"].(string)
		require.Equal(t, "1686416585000", timestampUnixMilliString1)
	})
}

func Test_Reference_WKT_OneOf1_Check2(t *testing.T) {
	bb, err := seedWKTOneOf2.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)
	/*
		{
			"one_enum1":{"f_enum_string1":"Three"},
			"f_enum_number_string2": "5",
			"f_any_proto1":{"@type":"type.googleapis.com/private.Message1", "f_string1":"s301", "f_string2":"s302", "f_string3":"s303"},
			"one_duration1":{"f_duration_string1":"1m52s"},
			"f_duration_microsecond1":114000000,
			"one_duration3":{"f_duration_second1":116},
			"f_duration_hour1":2,
			"one_duration5":{"f_duration_microsecond_string1":"114000000"},
			"f_duration_second_string1":"116",
			"one_duration7":{"f_duration_hour_string1":"2"},
			"one_timestamp1":{"f_timestamp_time_layout1":"2023-06-10T17:03:05Z"},
			"f_timestamp_unix_micro1":1686416585000000,
			"one_timestamp3":{"f_timestamp_unix_sec1":1686416585},
			"f_timestamp_unix_micro_string1":"1686416585000000",
			"one_timestamp5":{"f_timestamp_unix_sec_string1": "1686416585"},
		}
	*/
	t.Run("OneEnum1", func(t *testing.T) {
		oneEnum1 := data["one_enum1"].(map[string]interface{})
		enumString1 := oneEnum1["f_enum_string1"].(string)
		require.Equal(t, "Three", enumString1)
	})
	t.Run("OneEnum2", func(t *testing.T) {
		oneEnum2 := data["f_enum_number_string2"].(string)
		require.Equal(t, "5", oneEnum2)
	})

	t.Run("OneAny1", func(t *testing.T) {
		anyProto1 := data["f_any_proto1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/private.Message1", anyProto1["@type"].(string))
		require.Equal(t, "s301", anyProto1["f_string1"].(string))
		require.Equal(t, "s302", anyProto1["f_string2"].(string))
		require.Equal(t, "s303", anyProto1["f_string3"].(string))
	})

	t.Run("OneDuration1", func(t *testing.T) {
		oneDuration1 := data["one_duration1"].(map[string]interface{})
		durationString1 := oneDuration1["f_duration_string1"].(string)
		require.Equal(t, "1m52s", durationString1)
	})
	t.Run("OneDuration2", func(t *testing.T) {
		durationMicrosecond1 := data["f_duration_microsecond1"].(float64)
		require.Equal(t, float64(114000000), durationMicrosecond1)
	})
	t.Run("OneDuration3", func(t *testing.T) {
		oneDuration3 := data["one_duration3"].(map[string]interface{})
		durationSecond1 := oneDuration3["f_duration_second1"].(float64)
		require.Equal(t, float64(116), durationSecond1)
	})
	t.Run("OneDuration4", func(t *testing.T) {
		durationHour1 := data["f_duration_hour1"].(float64)
		require.Equal(t, float64(2), durationHour1)
	})
	t.Run("OneDuration5", func(t *testing.T) {
		oneDuration5 := data["one_duration5"].(map[string]interface{})
		durationMicrosecondString1 := oneDuration5["f_duration_microsecond_string1"].(string)
		require.Equal(t, "114000000", durationMicrosecondString1)
	})
	t.Run("OneDuration6", func(t *testing.T) {
		durationSecondString1 := data["f_duration_second_string1"].(string)
		require.Equal(t, "116", durationSecondString1)
	})
	t.Run("OneDuration7", func(t *testing.T) {
		oneDuration7 := data["one_duration7"].(map[string]interface{})
		durationHourString1 := oneDuration7["f_duration_hour_string1"].(string)
		require.Equal(t, "2", durationHourString1)
	})

	t.Run("OneTimestamp1", func(t *testing.T) {
		oneTimestamp1 := data["one_timestamp1"].(map[string]interface{})
		timestampTimeLayout1 := oneTimestamp1["f_timestamp_time_layout1"].(string)
		require.Equal(t, "2023-06-10T17:03:05Z", timestampTimeLayout1)
	})
	t.Run("OneTimestamp2", func(t *testing.T) {
		timestampUnixMicro1 := data["f_timestamp_unix_micro1"].(float64)
		require.Equal(t, float64(1686416585000000), timestampUnixMicro1)
	})
	t.Run("OneTimestamp3", func(t *testing.T) {
		oneTimestamp3 := data["one_timestamp3"].(map[string]interface{})
		timestampUnixSec1 := oneTimestamp3["f_timestamp_unix_sec1"].(float64)
		require.Equal(t, float64(1686416585), timestampUnixSec1)
	})

	t.Run("OneTimestamp4", func(t *testing.T) {
		timestampUnixMicroString1 := data["f_timestamp_unix_micro_string1"].(string)
		require.Equal(t, "1686416585000000", timestampUnixMicroString1)
	})
	t.Run("OneTimestamp5", func(t *testing.T) {
		oneTimestamp3 := data["one_timestamp5"].(map[string]interface{})
		timestampUnixSecString1 := oneTimestamp3["f_timestamp_unix_sec_string1"].(string)
		require.Equal(t, "1686416585", timestampUnixSecString1)
	})
}
