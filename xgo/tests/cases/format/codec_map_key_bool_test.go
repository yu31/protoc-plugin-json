package format

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_CodecMapKeyBool(t *testing.T) {
	seed1 := &pbformat.CodecMapKeyBool{
		FInt32KeyUnsetValUnset:         map[bool]int32{true: 2110111, false: 2110112},
		FInt32KeyBoolValNumeric:        map[bool]int32{true: 2110211, false: 2110212},
		FInt32KeyBoolValString:         map[bool]int32{true: 2110311, false: 2110312},
		FInt32KeyStringValNumeric:      map[bool]int32{true: 2110411, false: 2110412},
		FInt32KeyStringValString:       map[bool]int32{true: 2110511, false: 2110512},
		FInt64KeyUnsetValUnset:         map[bool]int64{true: 2120111, false: 2120112},
		FInt64KeyBoolValNumeric:        map[bool]int64{true: 2120211, false: 2120212},
		FInt64KeyBoolValString:         map[bool]int64{true: 2120311, false: 2120312},
		FInt64KeyStringValNumeric:      map[bool]int64{true: 2120411, false: 2120412},
		FInt64KeyStringValString:       map[bool]int64{true: 2120511, false: 2120512},
		FUint32KeyUnsetValUnset:        map[bool]uint32{true: 2130111, false: 2130112},
		FUint32KeyBoolValNumeric:       map[bool]uint32{true: 2130211, false: 2130212},
		FUint32KeyBoolValString:        map[bool]uint32{true: 2130311, false: 2130312},
		FUint32KeyStringValNumeric:     map[bool]uint32{true: 2130411, false: 2130412},
		FUint32KeyStringValString:      map[bool]uint32{true: 2130511, false: 2130512},
		FUint64KeyUnsetValUnset:        map[bool]uint64{true: 2140111, false: 2140112},
		FUint64KeyBoolValNumeric:       map[bool]uint64{true: 2140211, false: 2140212},
		FUint64KeyBoolValString:        map[bool]uint64{true: 2140311, false: 2140312},
		FUint64KeyStringValNumeric:     map[bool]uint64{true: 2140411, false: 2140412},
		FUint64KeyStringValString:      map[bool]uint64{true: 2140511, false: 2140512},
		FSint32KeyUnsetValUnset:        map[bool]int32{true: 2150111, false: 2150112},
		FSint32KeyBoolValNumeric:       map[bool]int32{true: 2150211, false: 2150212},
		FSint32KeyBoolValString:        map[bool]int32{true: 2150311, false: 2150312},
		FSint32KeyStringValNumeric:     map[bool]int32{true: 2150411, false: 2150412},
		FSint32KeyStringValString:      map[bool]int32{true: 2150511, false: 2150512},
		FSint64KeyUnsetValUnset:        map[bool]int64{true: 2160111, false: 2160112},
		FSint64KeyBoolValNumeric:       map[bool]int64{true: 2160211, false: 2160212},
		FSint64KeyBoolValString:        map[bool]int64{true: 2160311, false: 2160312},
		FSint64KeyStringValNumeric:     map[bool]int64{true: 2160411, false: 2160412},
		FSint64KeyStringValString:      map[bool]int64{true: 2160511, false: 2160512},
		FSfixed32KeyUnsetValUnset:      map[bool]int32{true: 2170111, false: 2170112},
		FSfixed32KeyBoolValNumeric:     map[bool]int32{true: 2170211, false: 2170212},
		FSfixed32KeyBoolValString:      map[bool]int32{true: 2170311, false: 2170312},
		FSfixed32KeyStringValNumeric:   map[bool]int32{true: 2170411, false: 2170412},
		FSfixed32KeyStringValString:    map[bool]int32{true: 2170511, false: 2170512},
		FSfixed64KeyUnsetValUnset:      map[bool]int64{true: 2180111, false: 2180112},
		FSfixed64KeyBoolValNumeric:     map[bool]int64{true: 2180211, false: 2180212},
		FSfixed64KeyBoolValString:      map[bool]int64{true: 2180311, false: 2180312},
		FSfixed64KeyStringValNumeric:   map[bool]int64{true: 2180411, false: 2180412},
		FSfixed64KeyStringValString:    map[bool]int64{true: 2180511, false: 2180512},
		FFixed32KeyUnsetValUnset:       map[bool]uint32{true: 2190111, false: 2190112},
		FFixed32KeyBoolValNumeric:      map[bool]uint32{true: 2190211, false: 2190212},
		FFixed32KeyBoolValString:       map[bool]uint32{true: 2190311, false: 2190312},
		FFixed32KeyStringValNumeric:    map[bool]uint32{true: 2190411, false: 2190412},
		FFixed32KeyStringValString:     map[bool]uint32{true: 2190511, false: 2190512},
		FFixed64KeyUnsetValUnset:       map[bool]uint64{true: 2210111, false: 2210112},
		FFixed64KeyBoolValNumeric:      map[bool]uint64{true: 2210211, false: 2210212},
		FFixed64KeyBoolValString:       map[bool]uint64{true: 2210311, false: 2210312},
		FFixed64KeyStringValNumeric:    map[bool]uint64{true: 2210411, false: 2210412},
		FFixed64KeyStringValString:     map[bool]uint64{true: 2210511, false: 2210512},
		FFloatKeyUnsetValUnset:         map[bool]float32{true: 2220.111, false: 2220.112},
		FFloatKeyBoolValNumeric:        map[bool]float32{true: 2220.211, false: 2220.212},
		FFloatKeyBoolValString:         map[bool]float32{true: 2220.311, false: 2220.312},
		FFloatKeyStringValNumeric:      map[bool]float32{true: 2220.411, false: 2220.412},
		FFloatKeyStringValString:       map[bool]float32{true: 2220.511, false: 2220.512},
		FDoubleKeyUnsetValUnset:        map[bool]float64{true: 2230111.111, false: 2230112.111},
		FDoubleKeyBoolValNumeric:       map[bool]float64{true: 2230211.111, false: 2230212.111},
		FDoubleKeyBoolValString:        map[bool]float64{true: 2230311.111, false: 2230312.111},
		FDoubleKeyStringValNumeric:     map[bool]float64{true: 2230411.111, false: 2230412.111},
		FDoubleKeyStringValString:      map[bool]float64{true: 2230511.111, false: 2230512.111},
		FBoolKeyUnsetValUnset:          map[bool]bool{true: true, false: true},
		FBoolKeyBoolValBool:            map[bool]bool{true: true, false: false},
		FBoolKeyBoolValString:          map[bool]bool{true: false, false: true},
		FBoolKeyStringValBool:          map[bool]bool{true: true, false: false},
		FBoolKeyStringValString:        map[bool]bool{true: false, false: true},
		FStringKeyUnsetValNone:         map[bool]string{true: "s101", false: "s102"},
		FStringKeyBoolValNone:          map[bool]string{true: "s201", false: "s202"},
		FStringKeyStringValNone:        map[bool]string{true: "s301", false: "s302"},
		FBytesKeyUnsetValNone:          map[bool][]byte{true: []byte("b101"), false: []byte("b102")},
		FBytesKeyBoolValNone:           map[bool][]byte{true: []byte("b201"), false: []byte("b202")},
		FBytesKeyStringValNone:         map[bool][]byte{true: []byte("b301"), false: []byte("b302")},
		FEnumKeyUnsetValUnset:          map[bool]pbexternal.EnumNum1{true: 0, false: 2},
		FEnumKeyBoolValNumeric:         map[bool]pbexternal.EnumNum1{true: 2, false: 3},
		FEnumKeyBoolValNumericString:   map[bool]pbexternal.EnumNum1{true: 3, false: 5},
		FEnumKeyBoolValEnumString:      map[bool]pbexternal.EnumNum1{true: 5, false: 6},
		FEnumKeyStringValNumeric:       map[bool]pbexternal.EnumNum1{true: 6, false: 7},
		FEnumKeyStringValNumericString: map[bool]pbexternal.EnumNum1{true: 7, false: 9},
		FEnumKeyStringValEnumString:    map[bool]pbexternal.EnumNum1{true: 9, false: 11},
		FMessageKeyUnsetValNone: map[bool]*pbexternal.Message1{
			true:  {FString1: "m101", FString2: "m102", FString3: "m103"},
			false: {FString1: "m201", FString2: "m202", FString3: "m203"},
		},
		FMessageKeyBoolValNone: map[bool]*pbexternal.Message1{
			true:  {FString1: "m111", FString2: "m112", FString3: "m113"},
			false: {FString1: "m211", FString2: "m212", FString3: "m213"},
		},
		FMessageKeyStringValNone: map[bool]*pbexternal.Message1{
			true:  {FString1: "m121", FString2: "m122", FString3: "m123"},
			false: {FString1: "m221", FString2: "m222", FString3: "m223"},
		},
		FDurationKeyUnsetValUnset: map[bool]*durationpb.Duration{
			true:  {Seconds: 110111, Nanos: 111},
			false: {Seconds: 110112, Nanos: 112},
		},
		FDurationKeyBoolValObject: map[bool]*durationpb.Duration{
			true:  {Seconds: 110121, Nanos: 121},
			false: {Seconds: 110122, Nanos: 122},
		},
		FDurationKeyBoolValTimeString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110131, Nanos: 131},
			false: {Seconds: 110132, Nanos: 132},
		},
		FDurationKeyBoolValNanosecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110141, Nanos: 141},
			false: {Seconds: 110142, Nanos: 142},
		},
		FDurationKeyBoolValNanosecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110151, Nanos: 151},
			false: {Seconds: 110152, Nanos: 152},
		},
		FDurationKeyBoolValMicrosecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110161, Nanos: 161000},
			false: {Seconds: 110162, Nanos: 162000},
		},
		FDurationKeyBoolValMicrosecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110171, Nanos: 171000},
			false: {Seconds: 110172, Nanos: 172000},
		},
		FDurationKeyBoolValMillisecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110181, Nanos: 181000000},
			false: {Seconds: 110182, Nanos: 182000000},
		},
		FDurationKeyBoolValMillisecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110191, Nanos: 191000000},
			false: {Seconds: 110192, Nanos: 192000000},
		},
		FDurationKeyBoolValSecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110211, Nanos: 0},
			false: {Seconds: 110212, Nanos: 0},
		},
		FDurationKeyBoolValSecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110221, Nanos: 0},
			false: {Seconds: 110222, Nanos: 0},
		},
		FDurationKeyBoolValMinute: map[bool]*durationpb.Duration{
			true:  {Seconds: 110231, Nanos: 0},
			false: {Seconds: 110232, Nanos: 0},
		},
		FDurationKeyBoolValMinuteString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110241, Nanos: 0},
			false: {Seconds: 110242, Nanos: 0},
		},
		FDurationKeyBoolValHour: map[bool]*durationpb.Duration{
			true:  {Seconds: 110251, Nanos: 0},
			false: {Seconds: 110252, Nanos: 0},
		},
		FDurationKeyBoolValHourString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110261, Nanos: 0},
			false: {Seconds: 110262, Nanos: 0},
		},
		FDurationKeyStringValUnset: map[bool]*durationpb.Duration{
			true:  {Seconds: 110271, Nanos: 271},
			false: {Seconds: 110272, Nanos: 272},
		},
		FDurationKeyStringValObject: map[bool]*durationpb.Duration{
			true:  {Seconds: 110281, Nanos: 281},
			false: {Seconds: 110282, Nanos: 282},
		},
		FDurationKeyStringValTimeString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110291, Nanos: 291},
			false: {Seconds: 110292, Nanos: 292},
		},
		FDurationKeyStringValNanosecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110311, Nanos: 311},
			false: {Seconds: 110312, Nanos: 312},
		},
		FDurationKeyStringValNanosecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110321, Nanos: 321},
			false: {Seconds: 110322, Nanos: 322},
		},
		FDurationKeyStringValMicrosecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110331, Nanos: 331000},
			false: {Seconds: 110332, Nanos: 332000},
		},
		FDurationKeyStringValMicrosecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110341, Nanos: 341000},
			false: {Seconds: 110342, Nanos: 342000},
		},
		FDurationKeyStringValMillisecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110351, Nanos: 351000000},
			false: {Seconds: 110352, Nanos: 352000000},
		},
		FDurationKeyStringValMillisecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110361, Nanos: 361000000},
			false: {Seconds: 110362, Nanos: 362000000},
		},
		FDurationKeyStringValSecond: map[bool]*durationpb.Duration{
			true:  {Seconds: 110371, Nanos: 0},
			false: {Seconds: 110372, Nanos: 0},
		},
		FDurationKeyStringValSecondString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110381, Nanos: 0},
			false: {Seconds: 110382, Nanos: 0},
		},
		FDurationKeyStringValMinute: map[bool]*durationpb.Duration{
			true:  {Seconds: 110391, Nanos: 0},
			false: {Seconds: 110392, Nanos: 0},
		},
		FDurationKeyStringValMinuteString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110411, Nanos: 0},
			false: {Seconds: 110412, Nanos: 0},
		},
		FDurationKeyStringValHour: map[bool]*durationpb.Duration{
			true:  {Seconds: 110421, Nanos: 0},
			false: {Seconds: 110422, Nanos: 0},
		},
		FDurationKeyStringValHourString: map[bool]*durationpb.Duration{
			true:  {Seconds: 110431, Nanos: 0},
			false: {Seconds: 110432, Nanos: 0},
		},
		FTimestampKeyUnsetValUnset: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217911, Nanos: 111},
			false: {Seconds: 1718217912, Nanos: 112},
		},
		FTimestampKeyBoolValObject: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217921, Nanos: 121},
			false: {Seconds: 1718217922, Nanos: 122},
		},
		FTimestampKeyBoolValTimeLayout: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217931, Nanos: 131},
			false: {Seconds: 1718217932, Nanos: 132},
		},
		FTimestampKeyBoolValUnixNano: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217941, Nanos: 141},
			false: {Seconds: 1718217942, Nanos: 142},
		},
		FTimestampKeyBoolValUnixNanoString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217951, Nanos: 151},
			false: {Seconds: 1718217952, Nanos: 152},
		},
		FTimestampKeyBoolValUnixMicro: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217961, Nanos: 161000},
			false: {Seconds: 1718217962, Nanos: 162000},
		},
		FTimestampKeyBoolValUnixMicroString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217971, Nanos: 171000},
			false: {Seconds: 1718217972, Nanos: 172000},
		},
		FTimestampKeyBoolValUnixMilli: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217981, Nanos: 181000000},
			false: {Seconds: 1718217982, Nanos: 182000000},
		},
		FTimestampKeyBoolValUnixMilliString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718217991, Nanos: 191000000},
			false: {Seconds: 1718217992, Nanos: 192000000},
		},
		FTimestampKeyBoolValUnixSec: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334211, Nanos: 0},
			false: {Seconds: 1718334212, Nanos: 0},
		},
		FTimestampKeyBoolValUnixSecString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334221, Nanos: 0},
			false: {Seconds: 1718334222, Nanos: 0},
		},
		FTimestampKeyStringValObject: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334231, Nanos: 231},
			false: {Seconds: 1718334232, Nanos: 232},
		},
		FTimestampKeyStringValTimeLayout: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334241, Nanos: 241},
			false: {Seconds: 1718334242, Nanos: 242},
		},
		FTimestampKeyStringValUnixNano: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334251, Nanos: 251},
			false: {Seconds: 1718334252, Nanos: 252},
		},
		FTimestampKeyStringValUnixNanoString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334261, Nanos: 261},
			false: {Seconds: 1718334262, Nanos: 262},
		},
		FTimestampKeyStringValUnixMicro: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334271, Nanos: 271000},
			false: {Seconds: 1718334272, Nanos: 272000},
		},
		FTimestampKeyStringValUnixMicroString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334281, Nanos: 281000},
			false: {Seconds: 1718334282, Nanos: 282000},
		},
		FTimestampKeyStringValUnixMilli: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334311, Nanos: 311000000},
			false: {Seconds: 1718334312, Nanos: 312000000},
		},
		FTimestampKeyStringValUnixMilliString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334321, Nanos: 321000000},
			false: {Seconds: 1718334322, Nanos: 322000000},
		},
		FTimestampKeyStringValUnixSec: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334331, Nanos: 0},
			false: {Seconds: 1718334332, Nanos: 0},
		},
		FTimestampKeyStringValUnixSecString: map[bool]*timestamppb.Timestamp{
			true:  {Seconds: 1718334341, Nanos: 0},
			false: {Seconds: 1718334342, Nanos: 0},
		},
		FAnyKeyUnsetValUnset: map[bool]*anypb.Any{
			true:  utils.MustNewAny(&pbexternal.Message1{FString1: "a111", FString2: "a112", FString3: "a113"}),
			false: utils.MustNewAny(&pbexternal.Message1{FString1: "a121", FString2: "a122", FString3: "a123"}),
		},
		FAnyKeyBoolValObject: map[bool]*anypb.Any{
			true:  utils.MustNewAny(&pbexternal.Message1{FString1: "a211", FString2: "a212", FString3: "a213"}),
			false: utils.MustNewAny(&pbexternal.Message1{FString1: "a221", FString2: "a222", FString3: "a223"}),
		},
		FAnyKeyBoolValProto: map[bool]*anypb.Any{
			true:  utils.MustNewAny(&pbexternal.Message1{FString1: "a311", FString2: "a312", FString3: "a313"}),
			false: utils.MustNewAny(&pbexternal.Message1{FString1: "a321", FString2: "a322", FString3: "a323"}),
		},
		FAnyKeyStringValObject: map[bool]*anypb.Any{
			true:  utils.MustNewAny(&pbexternal.Message1{FString1: "a411", FString2: "a412", FString3: "a413"}),
			false: utils.MustNewAny(&pbexternal.Message1{FString1: "a421", FString2: "a422", FString3: "a423"}),
		},
		FAnyKeyStringValProto: map[bool]*anypb.Any{
			true:  utils.MustNewAny(&pbexternal.Message1{FString1: "a511", FString2: "a512", FString3: "a513"}),
			false: utils.MustNewAny(&pbexternal.Message1{FString1: "a521", FString2: "a522", FString3: "a523"}),
		},
	}

	t.Run("MarshalAndUnmarshal", func(t *testing.T) {
		bb, err := seed1.MarshalJSON()
		require.Nil(t, err)

		dataNew := &pbformat.CodecMapKeyBool{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})

	t.Run("Check", func(t *testing.T) {
		{ // make the field to nil which not support with standard JSON.
			seed1.FInt32KeyBoolValNumeric = nil
			seed1.FInt32KeyBoolValString = nil

			seed1.FInt64KeyBoolValNumeric = nil
			seed1.FInt64KeyBoolValString = nil

			seed1.FUint32KeyBoolValNumeric = nil
			seed1.FUint32KeyBoolValString = nil

			seed1.FUint64KeyBoolValNumeric = nil
			seed1.FUint64KeyBoolValString = nil

			seed1.FSint32KeyBoolValNumeric = nil
			seed1.FSint32KeyBoolValString = nil

			seed1.FSint64KeyBoolValNumeric = nil
			seed1.FSint64KeyBoolValString = nil

			seed1.FSfixed32KeyBoolValNumeric = nil
			seed1.FSfixed32KeyBoolValString = nil

			seed1.FSfixed64KeyBoolValNumeric = nil
			seed1.FSfixed64KeyBoolValString = nil

			seed1.FFixed32KeyBoolValNumeric = nil
			seed1.FFixed32KeyBoolValString = nil

			seed1.FFixed64KeyBoolValNumeric = nil
			seed1.FFixed64KeyBoolValString = nil

			seed1.FFloatKeyBoolValNumeric = nil
			seed1.FFloatKeyBoolValString = nil

			seed1.FDoubleKeyBoolValNumeric = nil
			seed1.FDoubleKeyBoolValString = nil

			seed1.FBoolKeyBoolValBool = nil
			seed1.FBoolKeyBoolValString = nil

			seed1.FStringKeyBoolValNone = nil

			seed1.FBytesKeyBoolValNone = nil

			seed1.FEnumKeyBoolValNumeric = nil
			seed1.FEnumKeyBoolValNumericString = nil
			seed1.FEnumKeyBoolValEnumString = nil

			seed1.FMessageKeyBoolValNone = nil

			seed1.FDurationKeyBoolValObject = nil
			seed1.FDurationKeyBoolValTimeString = nil
			seed1.FDurationKeyBoolValNanosecond = nil
			seed1.FDurationKeyBoolValNanosecondString = nil
			seed1.FDurationKeyBoolValMicrosecond = nil
			seed1.FDurationKeyBoolValMicrosecondString = nil
			seed1.FDurationKeyBoolValMillisecond = nil
			seed1.FDurationKeyBoolValMillisecondString = nil
			seed1.FDurationKeyBoolValSecond = nil
			seed1.FDurationKeyBoolValSecondString = nil
			seed1.FDurationKeyBoolValMinute = nil
			seed1.FDurationKeyBoolValMinuteString = nil
			seed1.FDurationKeyBoolValHour = nil
			seed1.FDurationKeyBoolValHourString = nil

			seed1.FTimestampKeyBoolValObject = nil
			seed1.FTimestampKeyBoolValTimeLayout = nil
			seed1.FTimestampKeyBoolValUnixNano = nil
			seed1.FTimestampKeyBoolValUnixNanoString = nil
			seed1.FTimestampKeyBoolValUnixMicro = nil
			seed1.FTimestampKeyBoolValUnixMicroString = nil
			seed1.FTimestampKeyBoolValUnixMilli = nil
			seed1.FTimestampKeyBoolValUnixMilliString = nil
			seed1.FTimestampKeyBoolValUnixSec = nil
			seed1.FTimestampKeyBoolValUnixSecString = nil

			seed1.FAnyKeyBoolValObject = nil
			seed1.FAnyKeyBoolValProto = nil
		}

		bb, err := seed1.MarshalJSON()
		require.Nil(t, err)

		data := make(map[string]interface{})
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("Int32", func(t *testing.T) {
			t.Run("f_int32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_int32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2110111), kv["true"])
				require.Equal(t, float64(2110112), kv["false"])
			})
			t.Run("f_int32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_int32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2110411), kv["true"])
				require.Equal(t, float64(2110412), kv["false"])
			})
			t.Run("f_int32_key_string_val_string", func(t *testing.T) {
				kv := data["f_int32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2110511", kv["true"])
				require.Equal(t, "2110512", kv["false"])
			})
		})
		t.Run("Int64", func(t *testing.T) {
			t.Run("f_int64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_int64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2120111), kv["true"])
				require.Equal(t, float64(2120112), kv["false"])
			})
			t.Run("f_int64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_int64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2120411), kv["true"])
				require.Equal(t, float64(2120412), kv["false"])
			})
			t.Run("f_int64_key_string_val_string", func(t *testing.T) {
				kv := data["f_int64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2120511", kv["true"])
				require.Equal(t, "2120512", kv["false"])
			})
		})
		t.Run("UInt32", func(t *testing.T) {
			t.Run("f_uint32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_uint32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2130111), kv["true"])
				require.Equal(t, float64(2130112), kv["false"])
			})
			t.Run("f_uint32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_uint32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2130411), kv["true"])
				require.Equal(t, float64(2130412), kv["false"])
			})
			t.Run("f_uint32_key_string_val_string", func(t *testing.T) {
				kv := data["f_uint32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2130511", kv["true"])
				require.Equal(t, "2130512", kv["false"])
			})
		})
		t.Run("UInt64", func(t *testing.T) {
			t.Run("f_uint64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_uint64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2140111), kv["true"])
				require.Equal(t, float64(2140112), kv["false"])
			})
			t.Run("f_uint64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_uint64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2140411), kv["true"])
				require.Equal(t, float64(2140412), kv["false"])
			})
			t.Run("f_uint64_key_string_val_string", func(t *testing.T) {
				kv := data["f_uint64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2140511", kv["true"])
				require.Equal(t, "2140512", kv["false"])
			})
		})
		t.Run("SInt32", func(t *testing.T) {
			t.Run("f_sint32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sint32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2150111), kv["true"])
				require.Equal(t, float64(2150112), kv["false"])
			})
			t.Run("f_sint32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sint32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2150411), kv["true"])
				require.Equal(t, float64(2150412), kv["false"])
			})
			t.Run("f_sint32_key_string_val_string", func(t *testing.T) {
				kv := data["f_sint32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2150511", kv["true"])
				require.Equal(t, "2150512", kv["false"])
			})
		})
		t.Run("SInt64", func(t *testing.T) {
			t.Run("f_sint64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sint64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2160111), kv["true"])
				require.Equal(t, float64(2160112), kv["false"])
			})
			t.Run("f_sint64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sint64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2160411), kv["true"])
				require.Equal(t, float64(2160412), kv["false"])
			})
			t.Run("f_sint64_key_string_val_string", func(t *testing.T) {
				kv := data["f_sint64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2160511", kv["true"])
				require.Equal(t, "2160512", kv["false"])
			})
		})
		t.Run("SFixed32", func(t *testing.T) {
			t.Run("f_sfixed32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sfixed32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2170111), kv["true"])
				require.Equal(t, float64(2170112), kv["false"])
			})
			t.Run("f_sfixed32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sfixed32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2170411), kv["true"])
				require.Equal(t, float64(2170412), kv["false"])
			})
			t.Run("f_sfixed32_key_string_val_string", func(t *testing.T) {
				kv := data["f_sfixed32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2170511", kv["true"])
				require.Equal(t, "2170512", kv["false"])
			})
		})
		t.Run("SFixed64", func(t *testing.T) {
			t.Run("f_sfixed64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sfixed64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2180111), kv["true"])
				require.Equal(t, float64(2180112), kv["false"])
			})
			t.Run("f_sfixed64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sfixed64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2180411), kv["true"])
				require.Equal(t, float64(2180412), kv["false"])
			})
			t.Run("f_sfixed64_key_string_val_string", func(t *testing.T) {
				kv := data["f_sfixed64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2180511", kv["true"])
				require.Equal(t, "2180512", kv["false"])
			})
		})
		t.Run("Fixed32", func(t *testing.T) {
			t.Run("f_fixed32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_fixed32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2190111), kv["true"])
				require.Equal(t, float64(2190112), kv["false"])
			})
			t.Run("f_fixed32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_fixed32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2190411), kv["true"])
				require.Equal(t, float64(2190412), kv["false"])
			})
			t.Run("f_fixed32_key_string_val_string", func(t *testing.T) {
				kv := data["f_fixed32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2190511", kv["true"])
				require.Equal(t, "2190512", kv["false"])
			})
		})
		t.Run("Fixed64", func(t *testing.T) {
			t.Run("f_fixed64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_fixed64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2210111), kv["true"])
				require.Equal(t, float64(2210112), kv["false"])
			})
			t.Run("f_fixed64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_fixed64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2210411), kv["true"])
				require.Equal(t, float64(2210412), kv["false"])
			})
			t.Run("f_fixed64_key_string_val_string", func(t *testing.T) {
				kv := data["f_fixed64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2210511", kv["true"])
				require.Equal(t, "2210512", kv["false"])
			})
		})
		t.Run("Float", func(t *testing.T) {
			t.Run("f_float_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_float_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2220.111), kv["true"])
				require.Equal(t, float64(2220.112), kv["false"])
			})
			t.Run("f_float_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_float_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2220.411), kv["true"])
				require.Equal(t, float64(2220.412), kv["false"])
			})
			t.Run("f_float_key_string_val_string", func(t *testing.T) {
				kv := data["f_float_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2220.511", kv["true"])
				require.Equal(t, "2220.512", kv["false"])
			})
		})
		t.Run("Double", func(t *testing.T) {
			t.Run("f_double_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_double_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2230111.111), kv["true"])
				require.Equal(t, float64(2230112.111), kv["false"])
			})
			t.Run("f_double_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_double_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2230411.111), kv["true"])
				require.Equal(t, float64(2230412.111), kv["false"])
			})
			t.Run("f_double_key_string_val_string", func(t *testing.T) {
				kv := data["f_double_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2230511.111", kv["true"])
				require.Equal(t, "2230512.111", kv["false"])
			})
		})
		t.Run("Bool", func(t *testing.T) {
			t.Run("f_bool_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_bool_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, true, kv["true"])
				require.Equal(t, true, kv["false"])
			})
			t.Run("f_bool_key_string_val_bool", func(t *testing.T) {
				kv := data["f_bool_key_string_val_bool"].(map[string]interface{})
				require.Equal(t, true, kv["true"])
				require.Equal(t, false, kv["false"])
			})
			t.Run("f_bool_key_string_val_string", func(t *testing.T) {
				kv := data["f_bool_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "false", kv["true"])
				require.Equal(t, "true", kv["false"])
			})
		})
		t.Run("String", func(t *testing.T) {
			t.Run("f_string_key_unset_val_none", func(t *testing.T) {
				kv := data["f_string_key_unset_val_none"].(map[string]interface{})
				require.Equal(t, "s101", kv["true"])
				require.Equal(t, "s102", kv["false"])
			})
			t.Run("f_string_key_string_val_none", func(t *testing.T) {
				kv := data["f_string_key_string_val_none"].(map[string]interface{})
				require.Equal(t, "s301", kv["true"])
				require.Equal(t, "s302", kv["false"])
			})
		})
		t.Run("Bytes", func(t *testing.T) {
			t.Run("f_bytes_key_unset_val_none", func(t *testing.T) {
				kv := data["f_bytes_key_unset_val_none"].(map[string]interface{})
				require.Equal(t, "b101", utils.ConvertJSONBytesToString(kv["true"].(string)))
				require.Equal(t, "b102", utils.ConvertJSONBytesToString(kv["false"].(string)))
			})
			t.Run("f_bytes_key_bytes_val_none", func(t *testing.T) {
				kv := data["f_bytes_key_string_val_none"].(map[string]interface{})
				require.Equal(t, "b301", utils.ConvertJSONBytesToString(kv["true"].(string)))
				require.Equal(t, "b302", utils.ConvertJSONBytesToString(kv["false"].(string)))
			})
		})
		t.Run("Enum", func(t *testing.T) {
			t.Run("f_enum_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_enum_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(0), kv["true"])
				require.Equal(t, float64(2), kv["false"])
			})
			t.Run("f_enum_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_enum_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(6), kv["true"])
				require.Equal(t, float64(7), kv["false"])
			})
			t.Run("f_enum_key_string_val_numeric_string", func(t *testing.T) {
				kv := data["f_enum_key_string_val_numeric_string"].(map[string]interface{})
				require.Equal(t, "7", kv["true"])
				require.Equal(t, "9", kv["false"])
			})
			t.Run("f_enum_key_string_val_enum_string", func(t *testing.T) {
				kv := data["f_enum_key_string_val_enum_string"].(map[string]interface{})
				require.Equal(t, "Six", kv["true"])
				require.Equal(t, "Seven", kv["false"])
			})
		})
		t.Run("Message", func(t *testing.T) {
			t.Run("f_message_key_unset_val_none", func(t *testing.T) {
				kv := data["f_message_key_unset_val_none"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, "m101", v1["f_string1"])
				require.Equal(t, "m102", v1["f_string2"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, "m201", v2["f_string1"])
				require.Equal(t, "m202", v2["f_string2"])
			})
			t.Run("f_message_key_string_val_none", func(t *testing.T) {
				kv := data["f_message_key_string_val_none"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, "m121", v1["f_string1"])
				require.Equal(t, "m122", v1["f_string2"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, "m221", v2["f_string1"])
				require.Equal(t, "m222", v2["f_string2"])
			})
		})
		t.Run("Duration", func(t *testing.T) {
			t.Run("f_duration_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_duration_key_unset_val_unset"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, float64(110111), v1["seconds"])
				require.Equal(t, float64(111), v1["nanos"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, float64(110112), v2["seconds"])
				require.Equal(t, float64(112), v2["nanos"])
			})
			t.Run("f_duration_key_string_val_unset", func(t *testing.T) {
				kv := data["f_duration_key_string_val_unset"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, float64(110271), v1["seconds"])
				require.Equal(t, float64(271), v1["nanos"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, float64(110272), v2["seconds"])
				require.Equal(t, float64(272), v2["nanos"])
			})
			t.Run("f_duration_key_string_val_object", func(t *testing.T) {
				kv := data["f_duration_key_string_val_object"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, float64(110281), v1["seconds"])
				require.Equal(t, float64(281), v1["nanos"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, float64(110282), v2["seconds"])
				require.Equal(t, float64(282), v2["nanos"])
			})
			t.Run("f_duration_key_string_val_time_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_time_string"].(map[string]interface{})
				require.Equal(t, "30h38m11.000000291s", kv["true"])
				require.Equal(t, "30h38m12.000000292s", kv["false"])
			})
			t.Run("f_duration_key_string_val_nanosecond", func(t *testing.T) {
				kv := data["f_duration_key_string_val_nanosecond"].(map[string]interface{})
				require.Equal(t, float64(110311000000311), kv["true"])
				require.Equal(t, float64(110312000000312), kv["false"])
			})
			t.Run("f_duration_key_string_val_nanosecond_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_nanosecond_string"].(map[string]interface{})
				require.Equal(t, "110321000000321", kv["true"])
				require.Equal(t, "110322000000322", kv["false"])
			})
			t.Run("f_duration_key_string_val_microsecond", func(t *testing.T) {
				kv := data["f_duration_key_string_val_microsecond"].(map[string]interface{})
				require.Equal(t, float64(110331000331), kv["true"])
				require.Equal(t, float64(110332000332), kv["false"])
			})
			t.Run("f_duration_key_string_val_microsecond_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_microsecond_string"].(map[string]interface{})
				require.Equal(t, "110341000341", kv["true"])
				require.Equal(t, "110342000342", kv["false"])
			})
			t.Run("f_duration_key_string_val_millisecond", func(t *testing.T) {
				kv := data["f_duration_key_string_val_millisecond"].(map[string]interface{})
				require.Equal(t, float64(110351351), kv["true"])
				require.Equal(t, float64(110352352), kv["false"])
			})
			t.Run("f_duration_key_string_val_millisecond_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_millisecond_string"].(map[string]interface{})
				require.Equal(t, "110361361", kv["true"])
				require.Equal(t, "110362362", kv["false"])
			})
			t.Run("f_duration_key_string_val_second", func(t *testing.T) {
				kv := data["f_duration_key_string_val_second"].(map[string]interface{})
				require.Equal(t, float64(110371), kv["true"])
				require.Equal(t, float64(110372), kv["false"])
			})
			t.Run("f_duration_key_string_val_second_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_second_string"].(map[string]interface{})
				require.Equal(t, "110381", kv["true"])
				require.Equal(t, "110382", kv["false"])
			})
			t.Run("f_duration_key_string_val_minute", func(t *testing.T) {
				kv := data["f_duration_key_string_val_minute"].(map[string]interface{})
				require.Equal(t, float64(1839.85), kv["true"])
				require.Equal(t, float64(1839.8666666666666), kv["false"])
			})
			t.Run("f_duration_key_string_val_minute_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_minute_string"].(map[string]interface{})
				require.Equal(t, "1840.1833333333334", kv["true"])
				require.Equal(t, "1840.2", kv["false"])
			})
			t.Run("f_duration_key_string_val_hour", func(t *testing.T) {
				kv := data["f_duration_key_string_val_hour"].(map[string]interface{})
				require.Equal(t, float64(30.6725), kv["true"])
				require.Equal(t, float64(30.672777777777778), kv["false"])
			})
			t.Run("f_duration_key_string_val_hour_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_hour_string"].(map[string]interface{})
				require.Equal(t, "30.67527777777778", kv["true"])
				require.Equal(t, "30.675555555555555", kv["false"])
			})
		})
		t.Run("Timestamp", func(t *testing.T) {
			t.Run("f_timestamp_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_timestamp_key_unset_val_unset"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, float64(1718217911), v1["seconds"])
				require.Equal(t, float64(111), v1["nanos"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, float64(1718217912), v2["seconds"])
				require.Equal(t, float64(112), v2["nanos"])
			})
			t.Run("f_timestamp_key_string_val_object", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_object"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, float64(1718334231), v1["seconds"])
				require.Equal(t, float64(231), v1["nanos"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, float64(1718334232), v2["seconds"])
				require.Equal(t, float64(232), v2["nanos"])
			})
			t.Run("f_timestamp_key_string_val_time_layout", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_time_layout"].(map[string]interface{})
				require.Equal(t, "2024-06-14T03:04:01.000000241Z", kv["true"])
				require.Equal(t, "2024-06-14T03:04:02.000000242Z", kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_nano", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_nano"].(map[string]interface{})
				require.Equal(t, float64(1718334251000000251), kv["true"])
				require.Equal(t, float64(1718334252000000252), kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_nano_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_nano_string"].(map[string]interface{})
				require.Equal(t, "1718334261000000261", kv["true"])
				require.Equal(t, "1718334262000000262", kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_micro", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_micro"].(map[string]interface{})
				require.Equal(t, float64(1718334271000271), kv["true"])
				require.Equal(t, float64(1718334272000272), kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_micro_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_micro_string"].(map[string]interface{})
				require.Equal(t, "1718334281000281", kv["true"])
				require.Equal(t, "1718334282000282", kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_milli", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_milli"].(map[string]interface{})
				require.Equal(t, float64(1718334311311), kv["true"])
				require.Equal(t, float64(1718334312312), kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_milli_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_milli_string"].(map[string]interface{})
				require.Equal(t, "1718334321321", kv["true"])
				require.Equal(t, "1718334322322", kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_sec", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_sec"].(map[string]interface{})
				require.Equal(t, float64(1718334331), kv["true"])
				require.Equal(t, float64(1718334332), kv["false"])
			})
			t.Run("f_timestamp_key_string_val_unix_sec_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_sec_string"].(map[string]interface{})
				require.Equal(t, "1718334341", kv["true"])
				require.Equal(t, "1718334342", kv["false"])
			})
		})
		t.Run("Any", func(t *testing.T) {
			t.Run("f_any_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_any_key_unset_val_unset"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["type_url"])
				require.Equal(t, "CgRhMTExEgRhMTEyGgRhMTEz", v1["value"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["type_url"])
				require.Equal(t, "CgRhMTIxEgRhMTIyGgRhMTIz", v2["value"])
			})
			t.Run("f_any_key_string_val_object", func(t *testing.T) {
				kv := data["f_any_key_string_val_object"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["type_url"])
				require.Equal(t, "CgRhNDExEgRhNDEyGgRhNDEz", v1["value"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["type_url"])
				require.Equal(t, "CgRhNDIxEgRhNDIyGgRhNDIz", v2["value"])
			})
			t.Run("f_any_key_string_val_proto", func(t *testing.T) {
				kv := data["f_any_key_string_val_proto"].(map[string]interface{})
				v1 := kv["true"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["@type"])
				require.Equal(t, "a511", v1["f_string1"])
				require.Equal(t, "a512", v1["f_string2"])
				require.Equal(t, "a513", v1["f_string3"])
				v2 := kv["false"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["@type"])
				require.Equal(t, "a521", v2["f_string1"])
				require.Equal(t, "a522", v2["f_string2"])
				require.Equal(t, "a523", v2["f_string3"])
			})
		})
	})
}
