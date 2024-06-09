package format

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_CodecTypeRepeated1(t *testing.T) {
	seed1 := &pbformat.CodecTypeRepeated1{
		FInt32Numeric:      []int32{1111, 1112, 1113},
		FInt32String:       []int32{1121, 1122, 1123},
		FInt64Numeric:      []int64{1211, 1212, 1213},
		FInt64String:       []int64{1221, 1222, 1223},
		FUint32Numeric:     []uint32{1311, 1312, 1313},
		FUint32String:      []uint32{1321, 1322, 1323},
		FUint64Numeric:     []uint64{1411, 1412, 1413},
		FUint64String:      []uint64{1421, 1422, 1423},
		FSint32Numeric:     []int32{1511, 1512, 1513},
		FSint32String:      []int32{1521, 1522, 1523},
		FSint64Numeric:     []int64{1611, 1612, 1613},
		FSint64String:      []int64{1621, 1622, 1623},
		FSfixed32Numeric:   []int32{1711, 1712, 1713},
		FSfixed32String:    []int32{1721, 1722, 1723},
		FSfixed64Numeric:   []int64{1811, 1812, 1813},
		FSfixed64String:    []int64{1821, 1822, 1823},
		FFixed32Numeric:    []uint32{1911, 1912, 1913},
		FFixed32String:     []uint32{1921, 1922, 1923},
		FFixed64Numeric:    []uint64{2011, 2012, 2013},
		FFixed64String:     []uint64{2021, 2022, 2023},
		FFloatNumeric:      []float32{2111.111, 2112.111, 2113.111},
		FFloatString:       []float32{2121.111, 2122.111, 2123.111},
		FDoubleNumeric:     []float64{2211.111, 2212.111, 2213.111},
		FDoubleString:      []float64{2221.111, 2222.111, 2223.111},
		FBoolBool:          []bool{true, false, true},
		FBoolString:        []bool{true, false, true},
		FEnumNumeric:       []pbexternal.EnumNum1{0, 2, 6},
		FEnumNumericString: []pbexternal.EnumNum1{0, 2, 6},
		FEnumString:        []pbexternal.EnumNum1{3, 5, 7},
		FAnyNative: []*any.Any{
			utils.MustNewAny(&pbexternal.Message1{FString1: "s111", FString2: "s112", FString3: "s113"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s121", FString2: "s122", FString3: "s123"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s131", FString2: "s132", FString3: "s133"}),
		},
		FAnyProto: []*any.Any{
			utils.MustNewAny(&pbexternal.Message1{FString1: "s311", FString2: "s312", FString3: "s313"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s321", FString2: "s322", FString3: "s323"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s331", FString2: "s332", FString3: "s333"}),
		},
		FDurationNative: []*durationpb.Duration{
			{Seconds: 111, Nanos: 0}, {Seconds: 112, Nanos: 0}, {Seconds: 113, Nanos: 0},
		},
		FDurationString: []*durationpb.Duration{
			{Seconds: 121, Nanos: 0}, {Seconds: 122, Nanos: 0}, {Seconds: 123, Nanos: 0},
		},
		FDurationNanosecond: []*durationpb.Duration{
			{Seconds: 131, Nanos: 0}, {Seconds: 132, Nanos: 0}, {Seconds: 133, Nanos: 0},
		},
		FDurationNanosecondString: []*durationpb.Duration{
			{Seconds: 131, Nanos: 0}, {Seconds: 132, Nanos: 0}, {Seconds: 133, Nanos: 0},
		},
		FDurationMicrosecond: []*durationpb.Duration{
			{Seconds: 141, Nanos: 0}, {Seconds: 142, Nanos: 0}, {Seconds: 143, Nanos: 0},
		},
		FDurationMicrosecondString: []*durationpb.Duration{
			{Seconds: 141, Nanos: 0}, {Seconds: 142, Nanos: 0}, {Seconds: 143, Nanos: 0},
		},
		FDurationMillisecond: []*durationpb.Duration{
			{Seconds: 151, Nanos: 0}, {Seconds: 152, Nanos: 0}, {Seconds: 153, Nanos: 0},
		},
		FDurationMillisecondString: []*durationpb.Duration{
			{Seconds: 151, Nanos: 0}, {Seconds: 152, Nanos: 0}, {Seconds: 153, Nanos: 0},
		},
		FDurationSecond: []*durationpb.Duration{
			{Seconds: 161, Nanos: 0}, {Seconds: 162, Nanos: 0}, {Seconds: 163, Nanos: 0},
		},
		FDurationSecondString: []*durationpb.Duration{
			{Seconds: 161, Nanos: 0}, {Seconds: 162, Nanos: 0}, {Seconds: 163, Nanos: 0},
		},
		FDurationMinute: []*durationpb.Duration{
			{Seconds: 360, Nanos: 0}, {Seconds: 720, Nanos: 0}, {Seconds: 1080, Nanos: 0},
		},
		FDurationMinuteString: []*durationpb.Duration{
			{Seconds: 360, Nanos: 0}, {Seconds: 720, Nanos: 0}, {Seconds: 1080, Nanos: 0},
		},
		FDurationHour: []*durationpb.Duration{
			{Seconds: 3600, Nanos: 0}, {Seconds: 7200, Nanos: 0}, {Seconds: 10800, Nanos: 0},
		},
		FDurationHourString: []*durationpb.Duration{
			{Seconds: 3600, Nanos: 0}, {Seconds: 7200, Nanos: 0}, {Seconds: 10800, Nanos: 0},
		},
		FTimestampNative: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampTimeLayout: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixNano: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixNanoString: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixMicro: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixMicroString: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixMilli: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixMilliString: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixSec: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
		},
		FTimestampUnixSecString: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0}, {Seconds: 1686416585, Nanos: 0},
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
		dataNew := &pbformat.CodecTypeRepeated1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})

	t.Run("Check", func(t *testing.T) {
		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("Int32", func(t *testing.T) {
			a := data["f_int32_numeric"].([]interface{})
			require.Equal(t, float64(1111), a[0].(float64))
			require.Equal(t, float64(1112), a[1].(float64))
			require.Equal(t, float64(1113), a[2].(float64))

			b := data["f_int32_string"].([]interface{})
			require.Equal(t, "1121", b[0].(string))
			require.Equal(t, "1122", b[1].(string))
			require.Equal(t, "1123", b[2].(string))
		})
		t.Run("Int64", func(t *testing.T) {
			a := data["f_int64_numeric"].([]interface{})
			require.Equal(t, float64(1211), a[0].(float64))
			require.Equal(t, float64(1212), a[1].(float64))
			require.Equal(t, float64(1213), a[2].(float64))

			b := data["f_int64_string"].([]interface{})
			require.Equal(t, "1221", b[0].(string))
			require.Equal(t, "1222", b[1].(string))
			require.Equal(t, "1223", b[2].(string))
		})
		t.Run("UInt32", func(t *testing.T) {
			a := data["f_uint32_numeric"].([]interface{})
			require.Equal(t, float64(1311), a[0].(float64))
			require.Equal(t, float64(1312), a[1].(float64))
			require.Equal(t, float64(1313), a[2].(float64))

			b := data["f_uint32_string"].([]interface{})
			require.Equal(t, "1321", b[0].(string))
			require.Equal(t, "1322", b[1].(string))
			require.Equal(t, "1323", b[2].(string))
		})
		t.Run("UInt64", func(t *testing.T) {
			a := data["f_uint64_numeric"].([]interface{})
			require.Equal(t, float64(1411), a[0].(float64))
			require.Equal(t, float64(1412), a[1].(float64))
			require.Equal(t, float64(1413), a[2].(float64))

			b := data["f_uint64_string"].([]interface{})
			require.Equal(t, "1421", b[0].(string))
			require.Equal(t, "1422", b[1].(string))
			require.Equal(t, "1423", b[2].(string))
		})
		t.Run("SInt32", func(t *testing.T) {
			a := data["f_sint32_numeric"].([]interface{})
			require.Equal(t, float64(1511), a[0].(float64))
			require.Equal(t, float64(1512), a[1].(float64))
			require.Equal(t, float64(1513), a[2].(float64))

			b := data["f_sint32_string"].([]interface{})
			require.Equal(t, "1521", b[0].(string))
			require.Equal(t, "1522", b[1].(string))
			require.Equal(t, "1523", b[2].(string))
		})
		t.Run("SInt64", func(t *testing.T) {
			a := data["f_sint64_numeric"].([]interface{})
			require.Equal(t, float64(1611), a[0].(float64))
			require.Equal(t, float64(1612), a[1].(float64))
			require.Equal(t, float64(1613), a[2].(float64))

			b := data["f_sint64_string"].([]interface{})
			require.Equal(t, "1621", b[0].(string))
			require.Equal(t, "1622", b[1].(string))
			require.Equal(t, "1623", b[2].(string))
		})
		t.Run("SFixed32", func(t *testing.T) {
			a := data["f_sfixed32_numeric"].([]interface{})
			require.Equal(t, float64(1711), a[0].(float64))
			require.Equal(t, float64(1712), a[1].(float64))
			require.Equal(t, float64(1713), a[2].(float64))

			b := data["f_sfixed32_string"].([]interface{})
			require.Equal(t, "1721", b[0].(string))
			require.Equal(t, "1722", b[1].(string))
			require.Equal(t, "1723", b[2].(string))
		})
		t.Run("SFixed64", func(t *testing.T) {
			a := data["f_sfixed64_numeric"].([]interface{})
			require.Equal(t, float64(1811), a[0].(float64))
			require.Equal(t, float64(1812), a[1].(float64))
			require.Equal(t, float64(1813), a[2].(float64))

			b := data["f_sfixed64_string"].([]interface{})
			require.Equal(t, "1821", b[0].(string))
			require.Equal(t, "1822", b[1].(string))
			require.Equal(t, "1823", b[2].(string))
		})
		t.Run("Fixed32", func(t *testing.T) {
			a := data["f_fixed32_numeric"].([]interface{})
			require.Equal(t, float64(1911), a[0].(float64))
			require.Equal(t, float64(1912), a[1].(float64))
			require.Equal(t, float64(1913), a[2].(float64))

			b := data["f_fixed32_string"].([]interface{})
			require.Equal(t, "1921", b[0].(string))
			require.Equal(t, "1922", b[1].(string))
			require.Equal(t, "1923", b[2].(string))
		})
		t.Run("Fixed64", func(t *testing.T) {
			a := data["f_fixed64_numeric"].([]interface{})
			require.Equal(t, float64(2011), a[0].(float64))
			require.Equal(t, float64(2012), a[1].(float64))
			require.Equal(t, float64(2013), a[2].(float64))

			b := data["f_fixed64_string"].([]interface{})
			require.Equal(t, "2021", b[0].(string))
			require.Equal(t, "2022", b[1].(string))
			require.Equal(t, "2023", b[2].(string))
		})
		t.Run("Float", func(t *testing.T) {
			a := data["f_float_numeric"].([]interface{})
			require.Equal(t, float64(2111.111), a[0].(float64))
			require.Equal(t, float64(2112.111), a[1].(float64))
			require.Equal(t, float64(2113.111), a[2].(float64))

			b := data["f_float_string"].([]interface{})
			require.Equal(t, "2121.111", b[0].(string))
			require.Equal(t, "2122.111", b[1].(string))
			require.Equal(t, "2123.111", b[2].(string))
		})
		t.Run("Double", func(t *testing.T) {
			a := data["f_double_numeric"].([]interface{})
			require.Equal(t, float64(2211.111), a[0].(float64))
			require.Equal(t, float64(2212.111), a[1].(float64))
			require.Equal(t, float64(2213.111), a[2].(float64))

			b := data["f_double_string"].([]interface{})
			require.Equal(t, "2221.111", b[0].(string))
			require.Equal(t, "2222.111", b[1].(string))
			require.Equal(t, "2223.111", b[2].(string))
		})
		t.Run("Bool", func(t *testing.T) {
			a := data["f_bool_bool"].([]interface{})
			require.Equal(t, true, a[0].(bool))
			require.Equal(t, false, a[1].(bool))
			require.Equal(t, true, a[2].(bool))

			b := data["f_bool_string"].([]interface{})
			require.Equal(t, "true", b[0].(string))
			require.Equal(t, "false", b[1].(string))
			require.Equal(t, "true", b[2].(string))
		})

		t.Run("Enum", func(t *testing.T) {
			enumNumeric := data["f_enum_numeric"].([]interface{})
			require.Equal(t, float64(0), enumNumeric[0].(float64))
			require.Equal(t, float64(2), enumNumeric[1].(float64))
			require.Equal(t, float64(6), enumNumeric[2].(float64))

			enumNumericString := data["f_enum_numeric_string"].([]interface{})
			require.Equal(t, "0", enumNumericString[0].(string))
			require.Equal(t, "2", enumNumericString[1].(string))
			require.Equal(t, "6", enumNumericString[2].(string))

			enumString := data["f_enum_string"].([]interface{})
			require.Equal(t, "Two", enumString[0].(string))
			require.Equal(t, "Three", enumString[1].(string))
			require.Equal(t, "Five", enumString[2].(string))
		})

		t.Run("Any", func(t *testing.T) {
			anyNative := data["f_any_native"].([]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", anyNative[0].(map[string]interface{})["type_url"].(string))
			require.Equal(t, "CgRzMTExEgRzMTEyGgRzMTEz", anyNative[0].(map[string]interface{})["value"].(string))

			require.Equal(t, "type.googleapis.com/external.Message1", anyNative[1].(map[string]interface{})["type_url"].(string))
			require.Equal(t, "CgRzMTIxEgRzMTIyGgRzMTIz", anyNative[1].(map[string]interface{})["value"].(string))

			require.Equal(t, "type.googleapis.com/external.Message1", anyNative[2].(map[string]interface{})["type_url"].(string))
			require.Equal(t, "CgRzMTMxEgRzMTMyGgRzMTMz", anyNative[2].(map[string]interface{})["value"].(string))

			anyProto := data["f_any_proto"].([]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", anyProto[0].(map[string]interface{})["@type"].(string))
			require.Equal(t, "s311", anyProto[0].(map[string]interface{})["f_string1"].(string))
			require.Equal(t, "s312", anyProto[0].(map[string]interface{})["f_string2"].(string))
			require.Equal(t, "s313", anyProto[0].(map[string]interface{})["f_string3"].(string))

			require.Equal(t, "type.googleapis.com/external.Message1", anyProto[1].(map[string]interface{})["@type"].(string))
			require.Equal(t, "s321", anyProto[1].(map[string]interface{})["f_string1"].(string))
			require.Equal(t, "s322", anyProto[1].(map[string]interface{})["f_string2"].(string))
			require.Equal(t, "s323", anyProto[1].(map[string]interface{})["f_string3"].(string))

			require.Equal(t, "type.googleapis.com/external.Message1", anyProto[2].(map[string]interface{})["@type"].(string))
			require.Equal(t, "s331", anyProto[2].(map[string]interface{})["f_string1"].(string))
			require.Equal(t, "s332", anyProto[2].(map[string]interface{})["f_string2"].(string))
			require.Equal(t, "s333", anyProto[2].(map[string]interface{})["f_string3"].(string))
		})

		t.Run("Duration", func(t *testing.T) {
			anyDurationNative := data["f_duration_native"].([]interface{})
			require.Equal(t, float64(111), anyDurationNative[0].(map[string]interface{})["seconds"].(float64))
			require.Equal(t, float64(112), anyDurationNative[1].(map[string]interface{})["seconds"].(float64))
			require.Equal(t, float64(113), anyDurationNative[2].(map[string]interface{})["seconds"].(float64))

			anyDurationString := data["f_duration_string"].([]interface{})
			require.Equal(t, "2m1s", anyDurationString[0].(string))
			require.Equal(t, "2m2s", anyDurationString[1].(string))
			require.Equal(t, "2m3s", anyDurationString[2].(string))

			anyDurationNanosecond := data["f_duration_nanosecond"].([]interface{})
			require.Equal(t, float64(131000000000), anyDurationNanosecond[0].(float64))
			require.Equal(t, float64(132000000000), anyDurationNanosecond[1].(float64))
			require.Equal(t, float64(133000000000), anyDurationNanosecond[2].(float64))

			anyDurationNanosecondString := data["f_duration_nanosecond_string"].([]interface{})
			require.Equal(t, "131000000000", anyDurationNanosecondString[0].(string))
			require.Equal(t, "132000000000", anyDurationNanosecondString[1].(string))
			require.Equal(t, "133000000000", anyDurationNanosecondString[2].(string))

			anyDurationMicrosecond := data["f_duration_microsecond"].([]interface{})
			require.Equal(t, float64(141000000), anyDurationMicrosecond[0].(float64))
			require.Equal(t, float64(142000000), anyDurationMicrosecond[1].(float64))
			require.Equal(t, float64(143000000), anyDurationMicrosecond[2].(float64))

			anyDurationMicrosecondString := data["f_duration_microsecond_string"].([]interface{})
			require.Equal(t, "141000000", anyDurationMicrosecondString[0].(string))
			require.Equal(t, "142000000", anyDurationMicrosecondString[1].(string))
			require.Equal(t, "143000000", anyDurationMicrosecondString[2].(string))

			anyDurationMillisecond := data["f_duration_millisecond"].([]interface{})
			require.Equal(t, float64(151000), anyDurationMillisecond[0].(float64))
			require.Equal(t, float64(152000), anyDurationMillisecond[1].(float64))
			require.Equal(t, float64(153000), anyDurationMillisecond[2].(float64))

			anyDurationMillisecondString := data["f_duration_millisecond_string"].([]interface{})
			require.Equal(t, "151000", anyDurationMillisecondString[0].(string))
			require.Equal(t, "152000", anyDurationMillisecondString[1].(string))
			require.Equal(t, "153000", anyDurationMillisecondString[2].(string))

			anyDurationSecond := data["f_duration_second"].([]interface{})
			require.Equal(t, float64(161), anyDurationSecond[0].(float64))
			require.Equal(t, float64(162), anyDurationSecond[1].(float64))
			require.Equal(t, float64(163), anyDurationSecond[2].(float64))

			anyDurationSecondString := data["f_duration_second_string"].([]interface{})
			require.Equal(t, "161", anyDurationSecondString[0].(string))
			require.Equal(t, "162", anyDurationSecondString[1].(string))
			require.Equal(t, "163", anyDurationSecondString[2].(string))

			anyDurationMinute := data["f_duration_minute"].([]interface{})
			require.Equal(t, float64(6), anyDurationMinute[0].(float64))
			require.Equal(t, float64(12), anyDurationMinute[1].(float64))
			require.Equal(t, float64(18), anyDurationMinute[2].(float64))

			anyDurationMinuteString := data["f_duration_minute_string"].([]interface{})
			require.Equal(t, "6", anyDurationMinuteString[0].(string))
			require.Equal(t, "12", anyDurationMinuteString[1].(string))
			require.Equal(t, "18", anyDurationMinuteString[2].(string))

			anyDurationHour := data["f_duration_hour"].([]interface{})
			require.Equal(t, float64(1), anyDurationHour[0].(float64))
			require.Equal(t, float64(2), anyDurationHour[1].(float64))
			require.Equal(t, float64(3), anyDurationHour[2].(float64))

			anyDurationHourString := data["f_duration_hour_string"].([]interface{})
			require.Equal(t, "1", anyDurationHourString[0].(string))
			require.Equal(t, "2", anyDurationHourString[1].(string))
			require.Equal(t, "3", anyDurationHourString[2].(string))
		})

		t.Run("Timestamp", func(t *testing.T) {
			anyTimestampNative := data["f_timestamp_native"].([]interface{})
			require.Equal(t, float64(1686416585), anyTimestampNative[0].(map[string]interface{})["seconds"].(float64))
			require.Equal(t, float64(1686416585), anyTimestampNative[1].(map[string]interface{})["seconds"].(float64))

			anyTimestampTimeLayout := data["f_timestamp_time_layout"].([]interface{})
			require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout[0].(string))
			require.Equal(t, "2023-06-10T17:03:05Z", anyTimestampTimeLayout[1].(string))

			anyTimestampUnixNano := data["f_timestamp_unix_nano"].([]interface{})
			require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano[0].(float64))
			require.Equal(t, float64(1686416585000000000), anyTimestampUnixNano[1].(float64))

			anyTimestampUnixNanoString := data["f_timestamp_unix_nano_string"].([]interface{})
			require.Equal(t, "1686416585000000000", anyTimestampUnixNanoString[0].(string))
			require.Equal(t, "1686416585000000000", anyTimestampUnixNanoString[1].(string))

			anyTimestampUnixMicro := data["f_timestamp_unix_micro"].([]interface{})
			require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro[0].(float64))
			require.Equal(t, float64(1686416585000000), anyTimestampUnixMicro[1].(float64))

			anyTimestampUnixMicroString := data["f_timestamp_unix_micro_string"].([]interface{})
			require.Equal(t, "1686416585000000", anyTimestampUnixMicroString[0].(string))
			require.Equal(t, "1686416585000000", anyTimestampUnixMicroString[1].(string))

			anyTimestampUnixMilli := data["f_timestamp_unix_milli"].([]interface{})
			require.Equal(t, float64(1686416585000), anyTimestampUnixMilli[0].(float64))
			require.Equal(t, float64(1686416585000), anyTimestampUnixMilli[1].(float64))

			anyTimestampUnixMilliString := data["f_timestamp_unix_milli_string"].([]interface{})
			require.Equal(t, "1686416585000", anyTimestampUnixMilliString[0].(string))
			require.Equal(t, "1686416585000", anyTimestampUnixMilliString[1].(string))

			anyTimestampUnixSec := data["f_timestamp_unix_sec"].([]interface{})
			require.Equal(t, float64(1686416585), anyTimestampUnixSec[0].(float64))
			require.Equal(t, float64(1686416585), anyTimestampUnixSec[1].(float64))

			anyTimestampUnixSecString := data["f_timestamp_unix_sec_string"].([]interface{})
			require.Equal(t, "1686416585", anyTimestampUnixSecString[0].(string))
			require.Equal(t, "1686416585", anyTimestampUnixSecString[1].(string))
		})
	})
}
