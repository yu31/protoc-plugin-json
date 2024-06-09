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

func Test_CodecTypeOneOf1_B(t *testing.T) {
	seed1 := &pbformat.CodecTypeOneOf1{
		OneInt32:    &pbformat.CodecTypeOneOf1_FInt32String{FInt32String: 1112},
		OneInt64:    &pbformat.CodecTypeOneOf1_FInt64String{FInt64String: 1212},
		OneUint32:   &pbformat.CodecTypeOneOf1_FUint32String{FUint32String: 1312},
		OneUint64:   &pbformat.CodecTypeOneOf1_FUint64String{FUint64String: 1412},
		OneSInt32:   &pbformat.CodecTypeOneOf1_FSint32String{FSint32String: 1512},
		OneSInt64:   &pbformat.CodecTypeOneOf1_FSint64String{FSint64String: 1612},
		OneSFixed32: &pbformat.CodecTypeOneOf1_FSfixed32String{FSfixed32String: 1712},
		OneSFixed64: &pbformat.CodecTypeOneOf1_FSfixed64String{FSfixed64String: 1812},
		OneFixed32:  &pbformat.CodecTypeOneOf1_FFixed32String{FFixed32String: 1912},
		OneFixed64:  &pbformat.CodecTypeOneOf1_FFixed64String{FFixed64String: 2012},
		OneFloat:    &pbformat.CodecTypeOneOf1_FFloatString{FFloatString: 2112.111},
		OneDouble:   &pbformat.CodecTypeOneOf1_FDoubleString{FDoubleString: 2212.111},
		OneBool:     &pbformat.CodecTypeOneOf1_FBoolString{FBoolString: false},
		OneEnum1:    &pbformat.CodecTypeOneOf1_FEnumString{FEnumString: 5},
		OneEnum2:    &pbformat.CodecTypeOneOf1_FEnumNumericString2{FEnumNumericString2: 5},
		OneAny1: &pbformat.CodecTypeOneOf1_FAnyProto{
			FAnyProto: utils.MustNewAny(&pbexternal.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
		},
		OneDuration1: &pbformat.CodecTypeOneOf1_FDurationString{
			FDurationString: &durationpb.Duration{Seconds: 112, Nanos: 0},
		},
		OneDuration2: &pbformat.CodecTypeOneOf1_FDurationMicrosecond{
			FDurationMicrosecond: &durationpb.Duration{Seconds: 114, Nanos: 0},
		},
		OneDuration3: &pbformat.CodecTypeOneOf1_FDurationSecond{
			FDurationSecond: &durationpb.Duration{Seconds: 116, Nanos: 0},
		},
		OneDuration4: &pbformat.CodecTypeOneOf1_FDurationHour{
			FDurationHour: &durationpb.Duration{Seconds: 7200, Nanos: 0},
		},
		OneDuration5: &pbformat.CodecTypeOneOf1_FDurationMicrosecondString{
			FDurationMicrosecondString: &durationpb.Duration{Seconds: 114, Nanos: 0},
		},
		OneDuration6: &pbformat.CodecTypeOneOf1_FDurationSecondString{
			FDurationSecondString: &durationpb.Duration{Seconds: 116, Nanos: 0},
		},
		OneDuration7: &pbformat.CodecTypeOneOf1_FDurationHourString{
			FDurationHourString: &durationpb.Duration{Seconds: 7200, Nanos: 0},
		},
		OneTimestamp1: &pbformat.CodecTypeOneOf1_FTimestampTimeLayout{
			FTimestampTimeLayout: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp2: &pbformat.CodecTypeOneOf1_FTimestampUnixMicro{
			FTimestampUnixMicro: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp3: &pbformat.CodecTypeOneOf1_FTimestampUnixSec{
			FTimestampUnixSec: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp4: &pbformat.CodecTypeOneOf1_FTimestampUnixMicroString{
			FTimestampUnixMicroString: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp5: &pbformat.CodecTypeOneOf1_FTimestampUnixSecString{
			FTimestampUnixSecString: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
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
		dataNew := &pbformat.CodecTypeOneOf1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})

	t.Run("Check", func(t *testing.T) {
		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("OneInt32", func(t *testing.T) {
			require.Equal(t, "1112", data["f_int32_string"].(string))
		})
		t.Run("OneInt64", func(t *testing.T) {
			oneInt64 := data["one_int64"].(map[string]interface{})
			require.Equal(t, "1212", oneInt64["f_int64_string"].(string))
		})
		t.Run("OneUint32", func(t *testing.T) {
			require.Equal(t, "1312", data["f_uint32_string"].(string))
		})
		t.Run("OneUint64", func(t *testing.T) {
			oneUint64 := data["one_uint64"].(map[string]interface{})
			require.Equal(t, "1412", oneUint64["f_uint64_string"].(string))
		})
		t.Run("OneSInt32", func(t *testing.T) {
			require.Equal(t, "1512", data["f_sint32_string"].(string))
		})
		t.Run("OneSInt64", func(t *testing.T) {
			oneSInt64 := data["one_sint64"].(map[string]interface{})
			require.Equal(t, "1612", oneSInt64["f_sint64_string"].(string))
		})
		t.Run("OneSFixed32", func(t *testing.T) {
			require.Equal(t, "1712", data["f_sfixed32_string"].(string))
		})
		t.Run("OneSFixed64", func(t *testing.T) {
			oneSFixed64 := data["one_sfixed64"].(map[string]interface{})
			require.Equal(t, "1812", oneSFixed64["f_sfixed64_string"].(string))
		})
		t.Run("OneFixed32", func(t *testing.T) {
			require.Equal(t, "1912", data["f_fixed32_string"].(string))
		})
		t.Run("OneFixed64", func(t *testing.T) {
			oneFixed64 := data["one_fixed64"].(map[string]interface{})
			require.Equal(t, "2012", oneFixed64["f_fixed64_string"].(string))
		})

		t.Run("OneFloat", func(t *testing.T) {
			require.Equal(t, "2112.111", data["f_float_string"].(string))
		})
		t.Run("OneDouble", func(t *testing.T) {
			oneDouble := data["one_double"].(map[string]interface{})
			require.Equal(t, "2212.111", oneDouble["f_double_string"].(string))
		})

		t.Run("OneBool", func(t *testing.T) {
			require.Equal(t, "false", data["f_bool_string"].(string))
		})

		t.Run("OneEnum1", func(t *testing.T) {
			oneEnum1 := data["one_enum1"].(map[string]interface{})
			enumString1 := oneEnum1["f_enum_string"].(string)
			require.Equal(t, "Three", enumString1)
		})
		t.Run("OneEnum2", func(t *testing.T) {
			oneEnum2 := data["f_enum_numeric_string2"].(string)
			require.Equal(t, "5", oneEnum2)
		})

		t.Run("OneAny1", func(t *testing.T) {
			anyProto1 := data["f_any_proto"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", anyProto1["@type"].(string))
			require.Equal(t, "s301", anyProto1["f_string1"].(string))
			require.Equal(t, "s302", anyProto1["f_string2"].(string))
			require.Equal(t, "s303", anyProto1["f_string3"].(string))
		})

		t.Run("OneDuration1", func(t *testing.T) {
			oneDuration1 := data["one_duration1"].(map[string]interface{})
			durationString1 := oneDuration1["f_duration_string"].(string)
			require.Equal(t, "1m52s", durationString1)
		})
		t.Run("OneDuration2", func(t *testing.T) {
			durationMicrosecond1 := data["f_duration_microsecond"].(float64)
			require.Equal(t, float64(114000000), durationMicrosecond1)
		})
		t.Run("OneDuration3", func(t *testing.T) {
			oneDuration3 := data["one_duration3"].(map[string]interface{})
			durationSecond1 := oneDuration3["f_duration_second"].(float64)
			require.Equal(t, float64(116), durationSecond1)
		})
		t.Run("OneDuration4", func(t *testing.T) {
			durationHour1 := data["f_duration_hour"].(float64)
			require.Equal(t, float64(2), durationHour1)
		})
		t.Run("OneDuration5", func(t *testing.T) {
			oneDuration5 := data["one_duration5"].(map[string]interface{})
			durationMicrosecondString1 := oneDuration5["f_duration_microsecond_string"].(string)
			require.Equal(t, "114000000", durationMicrosecondString1)
		})
		t.Run("OneDuration6", func(t *testing.T) {
			durationSecondString1 := data["f_duration_second_string"].(string)
			require.Equal(t, "116", durationSecondString1)
		})
		t.Run("OneDuration7", func(t *testing.T) {
			oneDuration7 := data["one_duration7"].(map[string]interface{})
			durationHourString1 := oneDuration7["f_duration_hour_string"].(string)
			require.Equal(t, "2", durationHourString1)
		})

		t.Run("OneTimestamp1", func(t *testing.T) {
			oneTimestamp1 := data["one_timestamp1"].(map[string]interface{})
			timestampTimeLayout1 := oneTimestamp1["f_timestamp_time_layout"].(string)
			require.Equal(t, "2023-06-10T17:03:05Z", timestampTimeLayout1)
		})
		t.Run("OneTimestamp2", func(t *testing.T) {
			timestampUnixMicro1 := data["f_timestamp_unix_micro"].(float64)
			require.Equal(t, float64(1686416585000000), timestampUnixMicro1)
		})
		t.Run("OneTimestamp3", func(t *testing.T) {
			oneTimestamp3 := data["one_timestamp3"].(map[string]interface{})
			timestampUnixSec1 := oneTimestamp3["f_timestamp_unix_sec"].(float64)
			require.Equal(t, float64(1686416585), timestampUnixSec1)
		})

		t.Run("OneTimestamp4", func(t *testing.T) {
			timestampUnixMicroString1 := data["f_timestamp_unix_micro_string"].(string)
			require.Equal(t, "1686416585000000", timestampUnixMicroString1)
		})
		t.Run("OneTimestamp5", func(t *testing.T) {
			oneTimestamp3 := data["one_timestamp5"].(map[string]interface{})
			timestampUnixSecString1 := oneTimestamp3["f_timestamp_unix_sec_string"].(string)
			require.Equal(t, "1686416585", timestampUnixSecString1)
		})
	})
}
