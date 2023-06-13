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
var seedOneOf1 = &pbsets.TypeSetOneOf1{
	OneEnum1: &pbsets.TypeSetOneOf1_FEnumNumber1{FEnumNumber1: 2},
	OneAny1: &pbsets.TypeSetOneOf1_FAnyNative1{
		FAnyNative1: utils.MustNewAny(&pbsets.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
	},
	OneDuration1: &pbsets.TypeSetOneOf1_FDurationNative1{
		FDurationNative1: &durationpb.Duration{Seconds: 111, Nanos: 0},
	},
	OneDuration2: &pbsets.TypeSetOneOf1_FDurationNanoseconds1{
		FDurationNanoseconds1: &durationpb.Duration{Seconds: 113, Nanos: 0},
	},
	OneDuration3: &pbsets.TypeSetOneOf1_FDurationMilliseconds1{
		FDurationMilliseconds1: &durationpb.Duration{Seconds: 115, Nanos: 0},
	},
	OneDuration4: &pbsets.TypeSetOneOf1_FDurationMinutes1{
		FDurationMinutes1: &durationpb.Duration{Seconds: 117, Nanos: 0},
	},
	OneTimestamp1: &pbsets.TypeSetOneOf1_FTimestampNative1{
		FTimestampNative1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp2: &pbsets.TypeSetOneOf1_FTimestampUnixNano1{
		FTimestampUnixNano1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp3: &pbsets.TypeSetOneOf1_FTimestampUnixMilli1{
		FTimestampUnixMilli1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
}

// timestamp: 1686416585 -> 2023-06-11 01:03:05
var seedOneOf2 = &pbsets.TypeSetOneOf1{
	OneEnum1: &pbsets.TypeSetOneOf1_FEnumString1{FEnumString1: 5},
	OneAny1: &pbsets.TypeSetOneOf1_FAnyProto1{
		FAnyProto1: utils.MustNewAny(&pbsets.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
	},
	OneDuration1: &pbsets.TypeSetOneOf1_FDurationString1{
		FDurationString1: &durationpb.Duration{Seconds: 112, Nanos: 0},
	},
	OneDuration2: &pbsets.TypeSetOneOf1_FDurationMicroseconds1{
		FDurationMicroseconds1: &durationpb.Duration{Seconds: 114, Nanos: 0},
	},
	OneDuration3: &pbsets.TypeSetOneOf1_FDurationSeconds1{
		FDurationSeconds1: &durationpb.Duration{Seconds: 116, Nanos: 0},
	},
	OneDuration4: &pbsets.TypeSetOneOf1_FDurationHours1{
		FDurationHours1: &durationpb.Duration{Seconds: 7200, Nanos: 0},
	},
	OneTimestamp1: &pbsets.TypeSetOneOf1_FTimestampTimeLayout1{
		FTimestampTimeLayout1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp2: &pbsets.TypeSetOneOf1_FTimestampUnixMicro1{
		FTimestampUnixMicro1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
	OneTimestamp3: &pbsets.TypeSetOneOf1_FTimestampUnixSec1{
		FTimestampUnixSec1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
	},
}

func Test_Set_OneOf1_General1(t *testing.T) {
	bb, err := seedOneOf1.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbsets.TypeSetOneOf1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedOneOf1, dataNew)
}

func Test_Set_OneOf1_General2(t *testing.T) {
	bb, err := seedOneOf2.MarshalJSON()
	require.Nil(t, err)

	fmt.Println(string(bb))

	dataNew := &pbsets.TypeSetOneOf1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedOneOf2, dataNew)
}

func Test_Set_OneOf1_Check1(t *testing.T) {
	bb, err := seedOneOf1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)
	/*
		{
			"one_enum1":{"f_enum_number1":2},
			"f_any_native1":{"type_url":"type.googleapis.com/set_common.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
			"one_duration1":{"f_duration_native1":{"seconds":111}},
			"f_duration_nanoseconds1":113000000000,
			"one_duration3":{"f_duration_milliseconds1":115000},
			"f_duration_minutes1":1.95,
			"one_timestamp1":{"f_timestamp_native1":{"seconds":1686416585}},
			"f_timestamp_unix_nano1":1686416585000000000,
			"one_timestamp3":{"f_timestamp_unix_milli1":1686416585000}
		}
	*/
	t.Run("OneEnum1", func(t *testing.T) {
		oneEnum1 := data["one_enum1"].(map[string]interface{})
		enumNumber1 := oneEnum1["f_enum_number1"].(float64)
		require.Equal(t, float64(2), enumNumber1)
	})
	t.Run("OneAny1", func(t *testing.T) {
		anyNative1 := data["f_any_native1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyNative1["type_url"].(string))
		require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", anyNative1["value"].(string))
	})

	t.Run("OneDuration1", func(t *testing.T) {
		oneDuration1 := data["one_duration1"].(map[string]interface{})
		durationNative1 := oneDuration1["f_duration_native1"].(map[string]interface{})
		require.Equal(t, float64(111), durationNative1["seconds"].(float64))
	})
	t.Run("OneDuration2", func(t *testing.T) {
		durationNanoseconds1 := data["f_duration_nanoseconds1"].(float64)
		require.Equal(t, float64(113000000000), durationNanoseconds1)
	})
	t.Run("OneDuration3", func(t *testing.T) {
		oneDuration3 := data["one_duration3"].(map[string]interface{})
		durationMilliseconds1 := oneDuration3["f_duration_milliseconds1"].(float64)
		require.Equal(t, float64(115000), durationMilliseconds1)
	})
	t.Run("OneDuration4", func(t *testing.T) {
		durationMinutes1 := data["f_duration_minutes1"].(float64)
		require.Equal(t, float64(1.95), durationMinutes1)
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
}

func Test_Set_OneOf1_Check2(t *testing.T) {
	bb, err := seedOneOf2.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)
	/*
		{
			"one_enum1":{"f_enum_string1":"Three"},
			"f_any_proto1":{"@type":"type.googleapis.com/set_common.Message1", "f_string1":"s301", "f_string2":"s302", "f_string3":"s303"},
			"one_duration1":{"f_duration_string1":"1m52s"},
			"f_duration_microseconds1":114000000,
			"one_duration3":{"f_duration_seconds1":116},
			"f_duration_hours1":2,
			"one_timestamp1":{"f_timestamp_time_layout1":"2023-06-10T17:03:05Z"},
			"f_timestamp_unix_micro1":1686416585000000,
			"one_timestamp3":{"f_timestamp_unix_sec1":1686416585}
		}
	*/
	t.Run("OneEnum1", func(t *testing.T) {
		oneEnum1 := data["one_enum1"].(map[string]interface{})
		enumString1 := oneEnum1["f_enum_string1"].(string)
		require.Equal(t, "Three", enumString1)
	})
	t.Run("OneAny1", func(t *testing.T) {
		anyProto1 := data["f_any_proto1"].(map[string]interface{})
		require.Equal(t, "type.googleapis.com/set_common.Message1", anyProto1["@type"].(string))
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
		durationMicroseconds1 := data["f_duration_microseconds1"].(float64)
		require.Equal(t, float64(114000000), durationMicroseconds1)
	})
	t.Run("OneDuration3", func(t *testing.T) {
		oneDuration3 := data["one_duration3"].(map[string]interface{})
		durationSeconds1 := oneDuration3["f_duration_seconds1"].(float64)
		require.Equal(t, float64(116), durationSeconds1)
	})
	t.Run("OneDuration4", func(t *testing.T) {
		durationHours1 := data["f_duration_hours1"].(float64)
		require.Equal(t, float64(2), durationHours1)
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
}
