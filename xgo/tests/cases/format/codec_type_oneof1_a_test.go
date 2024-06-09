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

func Test_CodecTypeOneOf1_A(t *testing.T) {
	seed1 := &pbformat.CodecTypeOneOf1{
		OneInt32:    &pbformat.CodecTypeOneOf1_FInt32Numeric{FInt32Numeric: 1111},
		OneInt64:    &pbformat.CodecTypeOneOf1_FInt64Numeric{FInt64Numeric: 1211},
		OneUint32:   &pbformat.CodecTypeOneOf1_FUint32Numeric{FUint32Numeric: 1311},
		OneUint64:   &pbformat.CodecTypeOneOf1_FUint64Numeric{FUint64Numeric: 1411},
		OneSInt32:   &pbformat.CodecTypeOneOf1_FSint32Numeric{FSint32Numeric: 1511},
		OneSInt64:   &pbformat.CodecTypeOneOf1_FSint64Numeric{FSint64Numeric: 1611},
		OneSFixed32: &pbformat.CodecTypeOneOf1_FSfixed32Numeric{FSfixed32Numeric: 1711},
		OneSFixed64: &pbformat.CodecTypeOneOf1_FSfixed64Numeric{FSfixed64Numeric: 1811},
		OneFixed32:  &pbformat.CodecTypeOneOf1_FFixed32Numeric{FFixed32Numeric: 1911},
		OneFixed64:  &pbformat.CodecTypeOneOf1_FFixed64Numeric{FFixed64Numeric: 2011},
		OneFloat:    &pbformat.CodecTypeOneOf1_FFloatNumeric{FFloatNumeric: 2111.111},
		OneDouble:   &pbformat.CodecTypeOneOf1_FDoubleNumeric{FDoubleNumeric: 2211.111},
		OneBool:     &pbformat.CodecTypeOneOf1_FBoolBool{FBoolBool: true},
		OneEnum1:    &pbformat.CodecTypeOneOf1_FEnumNumeric{FEnumNumeric: 2},
		OneEnum2:    &pbformat.CodecTypeOneOf1_FEnumNumericString1{FEnumNumericString1: 2},
		OneAny1: &pbformat.CodecTypeOneOf1_FAnyNative{
			FAnyNative: utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		},
		OneDuration1: &pbformat.CodecTypeOneOf1_FDurationNative{
			FDurationNative: &durationpb.Duration{Seconds: 111, Nanos: 0},
		},
		OneDuration2: &pbformat.CodecTypeOneOf1_FDurationNanosecond{
			FDurationNanosecond: &durationpb.Duration{Seconds: 113, Nanos: 0},
		},
		OneDuration3: &pbformat.CodecTypeOneOf1_FDurationMillisecond{
			FDurationMillisecond: &durationpb.Duration{Seconds: 115, Nanos: 0},
		},
		OneDuration4: &pbformat.CodecTypeOneOf1_FDurationMinute{
			FDurationMinute: &durationpb.Duration{Seconds: 117, Nanos: 0},
		},
		OneDuration5: &pbformat.CodecTypeOneOf1_FDurationNanosecondString{
			FDurationNanosecondString: &durationpb.Duration{Seconds: 113, Nanos: 0},
		},
		OneDuration6: &pbformat.CodecTypeOneOf1_FDurationMillisecondString{
			FDurationMillisecondString: &durationpb.Duration{Seconds: 115, Nanos: 0},
		},
		OneDuration7: &pbformat.CodecTypeOneOf1_FDurationMinuteString{
			FDurationMinuteString: &durationpb.Duration{Seconds: 117, Nanos: 0},
		},
		OneTimestamp1: &pbformat.CodecTypeOneOf1_FTimestampNative{
			FTimestampNative: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp2: &pbformat.CodecTypeOneOf1_FTimestampUnixNano{
			FTimestampUnixNano: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp3: &pbformat.CodecTypeOneOf1_FTimestampUnixMilli{
			FTimestampUnixMilli: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp4: &pbformat.CodecTypeOneOf1_FTimestampUnixNanoString{
			FTimestampUnixNanoString: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
		},
		OneTimestamp5: &pbformat.CodecTypeOneOf1_FTimestampUnixMilliString{
			FTimestampUnixMilliString: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
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
			require.Equal(t, float64(1111), data["f_int32_numeric"].(float64))
		})
		t.Run("OneInt64", func(t *testing.T) {
			oneInt64 := data["one_int64"].(map[string]interface{})
			require.Equal(t, float64(1211), oneInt64["f_int64_numeric"].(float64))
		})
		t.Run("OneUint32", func(t *testing.T) {
			require.Equal(t, float64(1311), data["f_uint32_numeric"].(float64))
		})
		t.Run("OneUint64", func(t *testing.T) {
			oneUint64 := data["one_uint64"].(map[string]interface{})
			require.Equal(t, float64(1411), oneUint64["f_uint64_numeric"].(float64))
		})
		t.Run("OneSInt32", func(t *testing.T) {
			require.Equal(t, float64(1511), data["f_sint32_numeric"].(float64))
		})
		t.Run("OneSInt64", func(t *testing.T) {
			oneSInt64 := data["one_sint64"].(map[string]interface{})
			require.Equal(t, float64(1611), oneSInt64["f_sint64_numeric"].(float64))
		})
		t.Run("OneSFixed32", func(t *testing.T) {
			require.Equal(t, float64(1711), data["f_sfixed32_numeric"].(float64))
		})
		t.Run("OneSFixed64", func(t *testing.T) {
			oneSFixed64 := data["one_sfixed64"].(map[string]interface{})
			require.Equal(t, float64(1811), oneSFixed64["f_sfixed64_numeric"].(float64))
		})
		t.Run("OneFixed32", func(t *testing.T) {
			require.Equal(t, float64(1911), data["f_fixed32_numeric"].(float64))
		})
		t.Run("OneFixed64", func(t *testing.T) {
			oneFixed64 := data["one_fixed64"].(map[string]interface{})
			require.Equal(t, float64(2011), oneFixed64["f_fixed64_numeric"].(float64))
		})

		t.Run("OneFloat", func(t *testing.T) {
			require.Equal(t, float64(2111.111), data["f_float_numeric"].(float64))
		})
		t.Run("OneDouble", func(t *testing.T) {
			oneDouble := data["one_double"].(map[string]interface{})
			require.Equal(t, float64(2211.111), oneDouble["f_double_numeric"].(float64))
		})

		t.Run("OneBool", func(t *testing.T) {
			require.Equal(t, true, data["f_bool_bool"].(bool))
		})

		t.Run("OneEnum1", func(t *testing.T) {
			oneEnum1 := data["one_enum1"].(map[string]interface{})
			enumNumber1 := oneEnum1["f_enum_numeric"].(float64)
			require.Equal(t, float64(2), enumNumber1)
		})
		t.Run("OneEnum2", func(t *testing.T) {
			oneEnum2 := data["f_enum_numeric_string1"].(string)
			require.Equal(t, "2", oneEnum2)
		})

		t.Run("OneAny1", func(t *testing.T) {
			anyNative1 := data["f_any_native"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", anyNative1["type_url"].(string))
			require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", anyNative1["value"].(string))
		})

		t.Run("OneDuration1", func(t *testing.T) {
			oneDuration1 := data["one_duration1"].(map[string]interface{})
			durationNative1 := oneDuration1["f_duration_native"].(map[string]interface{})
			require.Equal(t, float64(111), durationNative1["seconds"].(float64))
		})
		t.Run("OneDuration2", func(t *testing.T) {
			durationNanosecond1 := data["f_duration_nanosecond"].(float64)
			require.Equal(t, float64(113000000000), durationNanosecond1)
		})
		t.Run("OneDuration3", func(t *testing.T) {
			oneDuration3 := data["one_duration3"].(map[string]interface{})
			durationMillisecond1 := oneDuration3["f_duration_millisecond"].(float64)
			require.Equal(t, float64(115000), durationMillisecond1)
		})
		t.Run("OneDuration4", func(t *testing.T) {
			durationMinute1 := data["f_duration_minute"].(float64)
			require.Equal(t, float64(1.95), durationMinute1)
		})
		t.Run("OneDuration5", func(t *testing.T) {
			oneDuration5 := data["one_duration5"].(map[string]interface{})
			durationNanosecondString1 := oneDuration5["f_duration_nanosecond_string"].(string)
			require.Equal(t, "113000000000", durationNanosecondString1)
		})
		t.Run("OneDuration6", func(t *testing.T) {
			durationMillisecondString1 := data["f_duration_millisecond_string"].(string)
			require.Equal(t, "115000", durationMillisecondString1)
		})
		t.Run("OneDuration7", func(t *testing.T) {
			oneDuration7 := data["one_duration7"].(map[string]interface{})
			durationMinuteString1 := oneDuration7["f_duration_minute_string"].(string)
			require.Equal(t, "1.95", durationMinuteString1)
		})

		t.Run("OneTimestamp1", func(t *testing.T) {
			oneTimestamp1 := data["one_timestamp1"].(map[string]interface{})
			timestampNative1 := oneTimestamp1["f_timestamp_native"].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestampNative1["seconds"].(float64))
		})
		t.Run("OneTimestamp2", func(t *testing.T) {
			timestampUnixNano1 := data["f_timestamp_unix_nano"].(float64)
			require.Equal(t, float64(1686416585000000000), timestampUnixNano1)
		})
		t.Run("OneTimestamp3", func(t *testing.T) {
			oneTimestamp3 := data["one_timestamp3"].(map[string]interface{})
			timestampUnixMilli1 := oneTimestamp3["f_timestamp_unix_milli"].(float64)
			require.Equal(t, float64(1686416585000), timestampUnixMilli1)
		})
		t.Run("OneTimestamp4", func(t *testing.T) {
			timestampUnixNanoString1 := data["f_timestamp_unix_nano_string"].(string)
			require.Equal(t, "1686416585000000000", timestampUnixNanoString1)
		})
		t.Run("OneTimestamp5", func(t *testing.T) {
			oneTimestamp3 := data["one_timestamp5"].(map[string]interface{})
			timestampUnixMilliString1 := oneTimestamp3["f_timestamp_unix_milli_string"].(string)
			require.Equal(t, "1686416585000", timestampUnixMilliString1)
		})
	})
}
