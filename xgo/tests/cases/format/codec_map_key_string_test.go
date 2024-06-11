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

func Test_CodecMapKeyString(t *testing.T) {
	seed1 := &pbformat.CodecMapKeyString{
		FInt32KeyNoneValUnset:        map[string]int32{"1110111": 2110111, "1110112": 2110112, "1110113": 2110113},
		FInt32KeyNoneValNumeric:      map[string]int32{"1110211": 2110211, "1110212": 2110212, "1110213": 2110213},
		FInt32KeyNoneValString:       map[string]int32{"1110311": 2110311, "1110312": 2110312, "1110313": 2110313},
		FInt64KeyNoneValUnset:        map[string]int64{"1120111": 2120111, "1120112": 2120112, "1120113": 2120113},
		FInt64KeyNoneValNumeric:      map[string]int64{"1120211": 2120211, "1120212": 2120212, "1120213": 2120213},
		FInt64KeyNoneValString:       map[string]int64{"1120311": 2120311, "1120312": 2120312, "1120313": 2120313},
		FUint32KeyNoneValUnset:       map[string]uint32{"1130111": 2130111, "1130112": 2130112, "1130113": 2130113},
		FUint32KeyNoneValNumeric:     map[string]uint32{"1130211": 2130211, "1130212": 2130212, "1130213": 2130213},
		FUint32KeyNoneValString:      map[string]uint32{"1130311": 2130311, "1130312": 2130312, "1130313": 2130313},
		FUint64KeyNoneValUnset:       map[string]uint64{"1140111": 2140111, "1140112": 2140112, "1140113": 2140113},
		FUint64KeyNoneValNumeric:     map[string]uint64{"1140211": 2140211, "1140212": 2140212, "1140213": 2140213},
		FUint64KeyNoneValString:      map[string]uint64{"1140311": 2140311, "1140312": 2140312, "1140313": 2140313},
		FSint32KeyNoneValUnset:       map[string]int32{"1150111": 2150111, "1150112": 2150112, "1150113": 2150113},
		FSint32KeyNoneValNumeric:     map[string]int32{"1150211": 2150211, "1150212": 2150212, "1150213": 2150213},
		FSint32KeyNoneValString:      map[string]int32{"1150311": 2150311, "1150312": 2150312, "1150313": 2150313},
		FSint64KeyNoneValUnset:       map[string]int64{"1160111": 2160111, "1160112": 2160112, "1160113": 2160113},
		FSint64KeyNoneValNumeric:     map[string]int64{"1160211": 2160211, "1160212": 2160212, "1160213": 2160213},
		FSint64KeyNoneValString:      map[string]int64{"1160311": 2160311, "1160312": 2160312, "1160313": 2160313},
		FSfixed32KeyNoneValUnset:     map[string]int32{"1170111": 2170111, "1170112": 2170112, "1170113": 2170113},
		FSfixed32KeyNoneValNumeric:   map[string]int32{"1170211": 2170211, "1170212": 2170212, "1170213": 2170213},
		FSfixed32KeyNoneValString:    map[string]int32{"1170311": 2170311, "1170312": 2170312, "1170313": 2170313},
		FSfixed64KeyNoneValUnset:     map[string]int64{"1180111": 2180111, "1180112": 2180112, "1180113": 2180113},
		FSfixed64KeyNoneValNumeric:   map[string]int64{"1180211": 2180211, "1180212": 2180212, "1180213": 2180213},
		FSfixed64KeyNoneValString:    map[string]int64{"1180311": 2180311, "1180312": 2180312, "1180313": 2180313},
		FFixed32KeyNoneValUnset:      map[string]uint32{"1190111": 2190111, "1190112": 2190112, "1190113": 2190113},
		FFixed32KeyNoneValNumeric:    map[string]uint32{"1190211": 2190211, "1190212": 2190212, "1190213": 2190213},
		FFixed32KeyNoneValString:     map[string]uint32{"1190311": 2190311, "1190312": 2190312, "1190313": 2190313},
		FFixed64KeyNoneValUnset:      map[string]uint64{"1210111": 2210111, "1210112": 2210112, "1210113": 2210113},
		FFixed64KeyNoneValNumeric:    map[string]uint64{"1210211": 2210211, "1210212": 2210212, "1210213": 2210213},
		FFixed64KeyNoneValString:     map[string]uint64{"1210311": 2210311, "1210312": 2210312, "1210313": 2210313},
		FFloatKeyNoneValUnset:        map[string]float32{"1220111": 2220.111, "1220112": 2220.112, "1220113": 2220.113},
		FFloatKeyNoneValNumeric:      map[string]float32{"1220211": 2220.211, "1220212": 2220.212, "1220213": 2220.213},
		FFloatKeyNoneValString:       map[string]float32{"1220311": 2220.311, "1220312": 2220.312, "1220313": 2220.313},
		FDoubleKeyNoneValUnset:       map[string]float64{"1230111": 2230111.111, "1230112": 2230112.111, "1230113": 2230113.111},
		FDoubleKeyNoneValNumeric:     map[string]float64{"1230211": 2230211.111, "1230212": 2230212.111, "1230213": 2230213.111},
		FDoubleKeyNoneValString:      map[string]float64{"1230311": 2230311.111, "1230312": 2230312.111, "1230313": 2230313.111},
		FBoolKeyNoneValUnset:         map[string]bool{"1240111": true, "1240112": true, "1240113": true},
		FBoolKeyNoneValBool:          map[string]bool{"1240211": true, "1240212": false, "1240213": true},
		FBoolKeyNoneValString:        map[string]bool{"1240311": false, "1240312": true, "1240313": false},
		FStringKeyNoneValNone:        map[string]string{"1250111": "s101", "1250112": "s102", "1250113": "s103"},
		FBytesKeyNoneValNone:         map[string][]byte{"1260111": []byte("b101"), "1260112": []byte("b102"), "1260113": []byte("b103")},
		FEnumKeyNoneValUnset:         map[string]pbexternal.EnumNum1{"1270111": 0, "1270112": 2, "1270113": 3},
		FEnumKeyNoneValNumeric:       map[string]pbexternal.EnumNum1{"1270211": 2, "1270212": 3, "1270213": 5},
		FEnumKeyNoneValNumericString: map[string]pbexternal.EnumNum1{"1270311": 3, "1270312": 5, "1270313": 6},
		FEnumKeyNoneValEnumString:    map[string]pbexternal.EnumNum1{"1270411": 5, "1270412": 6, "1270413": 7},
		FMessageKeyNoneValNone: map[string]*pbexternal.Message1{
			"1280111": {FString1: "m101", FString2: "m102", FString3: "m103"},
			"1280112": {FString1: "m201", FString2: "m202", FString3: "m203"},
			"1280113": {FString1: "m301", FString2: "m302", FString3: "m303"},
		},
		FDurationKeyNoneValUnset: map[string]*durationpb.Duration{
			"1290111": {Seconds: 110111, Nanos: 111},
			"1290112": {Seconds: 110112, Nanos: 112},
			"1290113": {Seconds: 110113, Nanos: 113},
		},
		FDurationKeyNoneValObject: map[string]*durationpb.Duration{
			"1290121": {Seconds: 110121, Nanos: 121},
			"1290122": {Seconds: 110122, Nanos: 122},
			"1290123": {Seconds: 110123, Nanos: 123},
		},
		FDurationKeyNoneValTimeString: map[string]*durationpb.Duration{
			"1290131": {Seconds: 110131, Nanos: 131},
			"1290132": {Seconds: 110132, Nanos: 132},
			"1290133": {Seconds: 110133, Nanos: 133},
		},
		FDurationKeyNoneValNanosecond: map[string]*durationpb.Duration{
			"1290141": {Seconds: 110141, Nanos: 141},
			"1290142": {Seconds: 110142, Nanos: 142},
			"1290143": {Seconds: 110143, Nanos: 143},
		},
		FDurationKeyNoneValNanosecondString: map[string]*durationpb.Duration{
			"1290151": {Seconds: 110151, Nanos: 151},
			"1290152": {Seconds: 110152, Nanos: 152},
			"1290153": {Seconds: 110153, Nanos: 153},
		},
		FDurationKeyNoneValMicrosecond: map[string]*durationpb.Duration{
			"1290161": {Seconds: 110161, Nanos: 161000},
			"1290162": {Seconds: 110162, Nanos: 162000},
			"1290163": {Seconds: 110163, Nanos: 163000},
		},
		FDurationKeyNoneValMicrosecondString: map[string]*durationpb.Duration{
			"1290171": {Seconds: 110171, Nanos: 171000},
			"1290172": {Seconds: 110172, Nanos: 172000},
			"1290173": {Seconds: 110173, Nanos: 173000},
		},
		FDurationKeyNoneValMillisecond: map[string]*durationpb.Duration{
			"1290181": {Seconds: 110181, Nanos: 181000000},
			"1290182": {Seconds: 110182, Nanos: 182000000},
			"1290183": {Seconds: 110183, Nanos: 183000000},
		},
		FDurationKeyNoneValMillisecondString: map[string]*durationpb.Duration{
			"1290191": {Seconds: 110191, Nanos: 191000000},
			"1290192": {Seconds: 110192, Nanos: 192000000},
			"1290193": {Seconds: 110193, Nanos: 193000000},
		},
		FDurationKeyNoneValSecond: map[string]*durationpb.Duration{
			"1290211": {Seconds: 110211, Nanos: 0},
			"1290212": {Seconds: 110212, Nanos: 0},
			"1290213": {Seconds: 110213, Nanos: 0},
		},
		FDurationKeyNoneValSecondString: map[string]*durationpb.Duration{
			"1290221": {Seconds: 110221, Nanos: 0},
			"1290222": {Seconds: 110222, Nanos: 0},
			"1290223": {Seconds: 110223, Nanos: 0},
		},
		FDurationKeyNoneValMinute: map[string]*durationpb.Duration{
			"1290231": {Seconds: 110231, Nanos: 0},
			"1290232": {Seconds: 110232, Nanos: 0},
			"1290233": {Seconds: 110233, Nanos: 0},
		},
		FDurationKeyNoneValMinuteString: map[string]*durationpb.Duration{
			"1290241": {Seconds: 110241, Nanos: 0},
			"1290242": {Seconds: 110242, Nanos: 0},
			"1290243": {Seconds: 110243, Nanos: 0},
		},
		FDurationKeyNoneValHour: map[string]*durationpb.Duration{
			"1290251": {Seconds: 110251, Nanos: 0},
			"1290252": {Seconds: 110252, Nanos: 0},
			"1290253": {Seconds: 110253, Nanos: 0},
		},
		FDurationKeyNoneValHourString: map[string]*durationpb.Duration{
			"1290261": {Seconds: 110261, Nanos: 0},
			"1290262": {Seconds: 110262, Nanos: 0},
			"1290263": {Seconds: 110263, Nanos: 0},
		},
		FTimestampKeyNoneValUnset: map[string]*timestamppb.Timestamp{
			"1310111": {Seconds: 1718217911, Nanos: 111},
			"1310112": {Seconds: 1718217912, Nanos: 112},
			"1310113": {Seconds: 1718217913, Nanos: 113},
		},
		FTimestampKeyNoneValObject: map[string]*timestamppb.Timestamp{
			"1310121": {Seconds: 1718217921, Nanos: 121},
			"1310122": {Seconds: 1718217922, Nanos: 122},
			"1310123": {Seconds: 1718217923, Nanos: 123},
		},
		FTimestampKeyNoneValTimeLayout: map[string]*timestamppb.Timestamp{
			"1310131": {Seconds: 1718217931, Nanos: 131},
			"1310132": {Seconds: 1718217932, Nanos: 132},
			"1310133": {Seconds: 1718217933, Nanos: 133},
		},
		FTimestampKeyNoneValUnixNano: map[string]*timestamppb.Timestamp{
			"1310141": {Seconds: 1718217941, Nanos: 141},
			"1310142": {Seconds: 1718217942, Nanos: 142},
			"1310143": {Seconds: 1718217943, Nanos: 143},
		},
		FTimestampKeyNoneValUnixNanoString: map[string]*timestamppb.Timestamp{
			"1310151": {Seconds: 1718217951, Nanos: 151},
			"1310152": {Seconds: 1718217952, Nanos: 152},
			"1310153": {Seconds: 1718217953, Nanos: 153},
		},
		FTimestampKeyNoneValUnixMicro: map[string]*timestamppb.Timestamp{
			"1310161": {Seconds: 1718217961, Nanos: 161000},
			"1310162": {Seconds: 1718217962, Nanos: 162000},
			"1310163": {Seconds: 1718217963, Nanos: 163000},
		},
		FTimestampKeyNoneValUnixMicroString: map[string]*timestamppb.Timestamp{
			"1310171": {Seconds: 1718217971, Nanos: 171000},
			"1310172": {Seconds: 1718217972, Nanos: 172000},
			"1310173": {Seconds: 1718217973, Nanos: 173000},
		},
		FTimestampKeyNoneValUnixMilli: map[string]*timestamppb.Timestamp{
			"1310181": {Seconds: 1718217981, Nanos: 181000000},
			"1310182": {Seconds: 1718217982, Nanos: 182000000},
			"1310183": {Seconds: 1718217983, Nanos: 183000000},
		},
		FTimestampKeyNoneValUnixMilliString: map[string]*timestamppb.Timestamp{
			"1310191": {Seconds: 1718217991, Nanos: 191000000},
			"1310192": {Seconds: 1718217992, Nanos: 192000000},
			"1310193": {Seconds: 1718217993, Nanos: 193000000},
		},
		FTimestampKeyNoneValUnixSec: map[string]*timestamppb.Timestamp{
			"1310211": {Seconds: 1718334211, Nanos: 0},
			"1310212": {Seconds: 1718334212, Nanos: 0},
			"1310213": {Seconds: 1718334213, Nanos: 0},
		},
		FTimestampKeyNoneValUnixSecString: map[string]*timestamppb.Timestamp{
			"1310221": {Seconds: 1718334221, Nanos: 0},
			"1310222": {Seconds: 1718334222, Nanos: 0},
			"1310223": {Seconds: 1718334223, Nanos: 0},
		},
		FAnyKeyNoneValUnset: map[string]*anypb.Any{
			"1320111": utils.MustNewAny(&pbexternal.Message1{FString1: "a111", FString2: "a112", FString3: "a113"}),
			"1320121": utils.MustNewAny(&pbexternal.Message1{FString1: "a121", FString2: "a122", FString3: "a123"}),
			"1320131": utils.MustNewAny(&pbexternal.Message1{FString1: "a131", FString2: "a132", FString3: "a133"}),
		},
		FAnyKeyNoneValObject: map[string]*anypb.Any{
			"1320211": utils.MustNewAny(&pbexternal.Message1{FString1: "a211", FString2: "a212", FString3: "a213"}),
			"1320221": utils.MustNewAny(&pbexternal.Message1{FString1: "a221", FString2: "a222", FString3: "a223"}),
			"1320231": utils.MustNewAny(&pbexternal.Message1{FString1: "a231", FString2: "a232", FString3: "a233"}),
		},
		FAnyKeyNoneValProto: map[string]*anypb.Any{
			"1320311": utils.MustNewAny(&pbexternal.Message1{FString1: "a311", FString2: "a312", FString3: "a313"}),
			"1320321": utils.MustNewAny(&pbexternal.Message1{FString1: "a321", FString2: "a322", FString3: "a323"}),
			"1320331": utils.MustNewAny(&pbexternal.Message1{FString1: "a331", FString2: "a332", FString3: "a333"}),
		},
	}

	t.Run("MarshalAndUnmarshal", func(t *testing.T) {
		bb, err := seed1.MarshalJSON()
		require.Nil(t, err)

		dataNew := &pbformat.CodecMapKeyString{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})

	t.Run("Check", func(t *testing.T) {
		bb, err := seed1.MarshalJSON()
		require.Nil(t, err)

		data := make(map[string]interface{})
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("Int32", func(t *testing.T) {
			t.Run("f_int32_key_none_val_unset", func(t *testing.T) {
				kv := data["f_int32_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2110111), kv["1110111"])
				require.Equal(t, float64(2110112), kv["1110112"])
				require.Equal(t, float64(2110113), kv["1110113"])
			})
			t.Run("f_int32_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_int32_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2110211), kv["1110211"])
				require.Equal(t, float64(2110212), kv["1110212"])
				require.Equal(t, float64(2110213), kv["1110213"])
			})
			t.Run("f_int32_key_none_val_string", func(t *testing.T) {
				kv := data["f_int32_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2110311", kv["1110311"])
				require.Equal(t, "2110312", kv["1110312"])
				require.Equal(t, "2110313", kv["1110313"])
			})
		})
		t.Run("Int64", func(t *testing.T) {
			t.Run("f_int64_key_none_val_unset", func(t *testing.T) {
				kv := data["f_int64_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2120111), kv["1120111"])
				require.Equal(t, float64(2120112), kv["1120112"])
				require.Equal(t, float64(2120113), kv["1120113"])
			})
			t.Run("f_int64_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_int64_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2120211), kv["1120211"])
				require.Equal(t, float64(2120212), kv["1120212"])
				require.Equal(t, float64(2120213), kv["1120213"])
			})
			t.Run("f_int64_key_none_val_string", func(t *testing.T) {
				kv := data["f_int64_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2120311", kv["1120311"])
				require.Equal(t, "2120312", kv["1120312"])
				require.Equal(t, "2120313", kv["1120313"])
			})
		})
		t.Run("UInt32", func(t *testing.T) {
			t.Run("f_uint32_key_none_val_unset", func(t *testing.T) {
				kv := data["f_uint32_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2130111), kv["1130111"])
				require.Equal(t, float64(2130112), kv["1130112"])
				require.Equal(t, float64(2130113), kv["1130113"])
			})
			t.Run("f_uint32_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_uint32_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2130211), kv["1130211"])
				require.Equal(t, float64(2130212), kv["1130212"])
				require.Equal(t, float64(2130213), kv["1130213"])
			})
			t.Run("f_uint32_key_none_val_string", func(t *testing.T) {
				kv := data["f_uint32_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2130311", kv["1130311"])
				require.Equal(t, "2130312", kv["1130312"])
				require.Equal(t, "2130313", kv["1130313"])
			})
		})
		t.Run("UInt64", func(t *testing.T) {
			t.Run("f_uint64_key_none_val_unset", func(t *testing.T) {
				kv := data["f_uint64_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2140111), kv["1140111"])
				require.Equal(t, float64(2140112), kv["1140112"])
				require.Equal(t, float64(2140113), kv["1140113"])
			})
			t.Run("f_uint64_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_uint64_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2140211), kv["1140211"])
				require.Equal(t, float64(2140212), kv["1140212"])
				require.Equal(t, float64(2140213), kv["1140213"])
			})
			t.Run("f_uint64_key_none_val_string", func(t *testing.T) {
				kv := data["f_uint64_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2140311", kv["1140311"])
				require.Equal(t, "2140312", kv["1140312"])
				require.Equal(t, "2140313", kv["1140313"])
			})
		})
		t.Run("SInt32", func(t *testing.T) {
			t.Run("f_sint32_key_none_val_unset", func(t *testing.T) {
				kv := data["f_sint32_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2150111), kv["1150111"])
				require.Equal(t, float64(2150112), kv["1150112"])
				require.Equal(t, float64(2150113), kv["1150113"])
			})
			t.Run("f_sint32_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_sint32_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2150211), kv["1150211"])
				require.Equal(t, float64(2150212), kv["1150212"])
				require.Equal(t, float64(2150213), kv["1150213"])
			})
			t.Run("f_sint32_key_none_val_string", func(t *testing.T) {
				kv := data["f_sint32_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2150311", kv["1150311"])
				require.Equal(t, "2150312", kv["1150312"])
				require.Equal(t, "2150313", kv["1150313"])
			})
		})
		t.Run("SInt64", func(t *testing.T) {
			t.Run("f_sint64_key_none_val_unset", func(t *testing.T) {
				kv := data["f_sint64_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2160111), kv["1160111"])
				require.Equal(t, float64(2160112), kv["1160112"])
				require.Equal(t, float64(2160113), kv["1160113"])
			})
			t.Run("f_sint64_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_sint64_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2160211), kv["1160211"])
				require.Equal(t, float64(2160212), kv["1160212"])
				require.Equal(t, float64(2160213), kv["1160213"])
			})
			t.Run("f_sint64_key_none_val_string", func(t *testing.T) {
				kv := data["f_sint64_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2160311", kv["1160311"])
				require.Equal(t, "2160312", kv["1160312"])
				require.Equal(t, "2160313", kv["1160313"])
			})
		})
		t.Run("SFixed32", func(t *testing.T) {
			t.Run("f_sfixed32_key_none_val_unset", func(t *testing.T) {
				kv := data["f_sfixed32_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2170111), kv["1170111"])
				require.Equal(t, float64(2170112), kv["1170112"])
				require.Equal(t, float64(2170113), kv["1170113"])
			})
			t.Run("f_sfixed32_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_sfixed32_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2170211), kv["1170211"])
				require.Equal(t, float64(2170212), kv["1170212"])
				require.Equal(t, float64(2170213), kv["1170213"])
			})
			t.Run("f_sfixed32_key_none_val_string", func(t *testing.T) {
				kv := data["f_sfixed32_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2170311", kv["1170311"])
				require.Equal(t, "2170312", kv["1170312"])
				require.Equal(t, "2170313", kv["1170313"])
			})
		})
		t.Run("SFixed64", func(t *testing.T) {
			t.Run("f_sfixed64_key_none_val_unset", func(t *testing.T) {
				kv := data["f_sfixed64_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2180111), kv["1180111"])
				require.Equal(t, float64(2180112), kv["1180112"])
				require.Equal(t, float64(2180113), kv["1180113"])
			})
			t.Run("f_sfixed64_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_sfixed64_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2180211), kv["1180211"])
				require.Equal(t, float64(2180212), kv["1180212"])
				require.Equal(t, float64(2180213), kv["1180213"])
			})
			t.Run("f_sfixed64_key_none_val_string", func(t *testing.T) {
				kv := data["f_sfixed64_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2180311", kv["1180311"])
				require.Equal(t, "2180312", kv["1180312"])
				require.Equal(t, "2180313", kv["1180313"])
			})
		})
		t.Run("Fixed32", func(t *testing.T) {
			t.Run("f_fixed32_key_none_val_unset", func(t *testing.T) {
				kv := data["f_fixed32_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2190111), kv["1190111"])
				require.Equal(t, float64(2190112), kv["1190112"])
				require.Equal(t, float64(2190113), kv["1190113"])
			})
			t.Run("f_fixed32_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_fixed32_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2190211), kv["1190211"])
				require.Equal(t, float64(2190212), kv["1190212"])
				require.Equal(t, float64(2190213), kv["1190213"])
			})
			t.Run("f_fixed32_key_none_val_string", func(t *testing.T) {
				kv := data["f_fixed32_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2190311", kv["1190311"])
				require.Equal(t, "2190312", kv["1190312"])
				require.Equal(t, "2190313", kv["1190313"])
			})
		})
		t.Run("Fixed64", func(t *testing.T) {
			t.Run("f_fixed64_key_none_val_unset", func(t *testing.T) {
				kv := data["f_fixed64_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2210111), kv["1210111"])
				require.Equal(t, float64(2210112), kv["1210112"])
				require.Equal(t, float64(2210113), kv["1210113"])
			})
			t.Run("f_fixed64_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_fixed64_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2210211), kv["1210211"])
				require.Equal(t, float64(2210212), kv["1210212"])
				require.Equal(t, float64(2210213), kv["1210213"])
			})
			t.Run("f_fixed64_key_none_val_string", func(t *testing.T) {
				kv := data["f_fixed64_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2210311", kv["1210311"])
				require.Equal(t, "2210312", kv["1210312"])
				require.Equal(t, "2210313", kv["1210313"])
			})
		})
		t.Run("Float", func(t *testing.T) {
			t.Run("f_float_key_none_val_unset", func(t *testing.T) {
				kv := data["f_float_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2220.111), kv["1220111"])
				require.Equal(t, float64(2220.112), kv["1220112"])
				require.Equal(t, float64(2220.113), kv["1220113"])
			})
			t.Run("f_float_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_float_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2220.211), kv["1220211"])
				require.Equal(t, float64(2220.212), kv["1220212"])
				require.Equal(t, float64(2220.213), kv["1220213"])
			})
			t.Run("f_float_key_none_val_string", func(t *testing.T) {
				kv := data["f_float_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2220.311", kv["1220311"])
				require.Equal(t, "2220.312", kv["1220312"])
				require.Equal(t, "2220.313", kv["1220313"])
			})
		})
		t.Run("Double", func(t *testing.T) {
			t.Run("f_double_key_none_val_unset", func(t *testing.T) {
				kv := data["f_double_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2230111.111), kv["1230111"])
				require.Equal(t, float64(2230112.111), kv["1230112"])
				require.Equal(t, float64(2230113.111), kv["1230113"])
			})
			t.Run("f_double_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_double_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2230211.111), kv["1230211"])
				require.Equal(t, float64(2230212.111), kv["1230212"])
				require.Equal(t, float64(2230213.111), kv["1230213"])
			})
			t.Run("f_double_key_none_val_string", func(t *testing.T) {
				kv := data["f_double_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "2230311.111", kv["1230311"])
				require.Equal(t, "2230312.111", kv["1230312"])
				require.Equal(t, "2230313.111", kv["1230313"])
			})
		})
		t.Run("Bool", func(t *testing.T) {
			t.Run("f_bool_key_none_val_unset", func(t *testing.T) {
				kv := data["f_bool_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, true, kv["1240111"])
				require.Equal(t, true, kv["1240112"])
				require.Equal(t, true, kv["1240113"])
			})
			t.Run("f_bool_key_none_val_bool", func(t *testing.T) {
				kv := data["f_bool_key_none_val_bool"].(map[string]interface{})
				require.Equal(t, true, kv["1240211"])
				require.Equal(t, false, kv["1240212"])
				require.Equal(t, true, kv["1240213"])
			})
			t.Run("f_bool_key_none_val_string", func(t *testing.T) {
				kv := data["f_bool_key_none_val_string"].(map[string]interface{})
				require.Equal(t, "false", kv["1240311"])
				require.Equal(t, "true", kv["1240312"])
				require.Equal(t, "false", kv["1240313"])
			})
		})
		t.Run("String", func(t *testing.T) {
			t.Run("f_string_key_none_val_none", func(t *testing.T) {
				kv := data["f_string_key_none_val_none"].(map[string]interface{})
				require.Equal(t, "s101", kv["1250111"])
				require.Equal(t, "s102", kv["1250112"])
				require.Equal(t, "s103", kv["1250113"])
			})
		})
		t.Run("Bytes", func(t *testing.T) {
			t.Run("f_bytes_key_none_val_none", func(t *testing.T) {
				kv := data["f_bytes_key_none_val_none"].(map[string]interface{})
				require.Equal(t, "b101", utils.ConvertJSONBytesToString(kv["1260111"].(string)))
				require.Equal(t, "b102", utils.ConvertJSONBytesToString(kv["1260112"].(string)))
				require.Equal(t, "b103", utils.ConvertJSONBytesToString(kv["1260113"].(string)))
			})
		})
		t.Run("Enum", func(t *testing.T) {
			t.Run("f_enum_key_none_val_unset", func(t *testing.T) {
				kv := data["f_enum_key_none_val_unset"].(map[string]interface{})
				require.Equal(t, float64(0), kv["1270111"])
				require.Equal(t, float64(2), kv["1270112"])
				require.Equal(t, float64(3), kv["1270113"])
			})
			t.Run("f_enum_key_none_val_numeric", func(t *testing.T) {
				kv := data["f_enum_key_none_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2), kv["1270211"])
				require.Equal(t, float64(3), kv["1270212"])
				require.Equal(t, float64(5), kv["1270213"])
			})
			t.Run("f_enum_key_none_val_numeric_string", func(t *testing.T) {
				kv := data["f_enum_key_none_val_numeric_string"].(map[string]interface{})
				require.Equal(t, "3", kv["1270311"])
				require.Equal(t, "5", kv["1270312"])
				require.Equal(t, "6", kv["1270313"])
			})
			t.Run("f_enum_key_none_val_enum_string", func(t *testing.T) {
				kv := data["f_enum_key_none_val_enum_string"].(map[string]interface{})
				require.Equal(t, "Three", kv["1270411"])
				require.Equal(t, "Four", kv["1270412"])
				require.Equal(t, "Five", kv["1270413"])
			})
		})
		t.Run("Message", func(t *testing.T) {
			t.Run("f_message_key_none_val_none", func(t *testing.T) {
				kv := data["f_message_key_none_val_none"].(map[string]interface{})
				v1 := kv["1280111"].(map[string]interface{})
				require.Equal(t, "m101", v1["f_string1"])
				require.Equal(t, "m102", v1["f_string2"])
				require.Equal(t, "m103", v1["f_string3"])
				v2 := kv["1280112"].(map[string]interface{})
				require.Equal(t, "m201", v2["f_string1"])
				require.Equal(t, "m202", v2["f_string2"])
				require.Equal(t, "m203", v2["f_string3"])
				v3 := kv["1280113"].(map[string]interface{})
				require.Equal(t, "m301", v3["f_string1"])
				require.Equal(t, "m302", v3["f_string2"])
				require.Equal(t, "m303", v3["f_string3"])
			})
		})
		t.Run("Duration", func(t *testing.T) {
			t.Run("f_duration_key_none_val_unset", func(t *testing.T) {
				kv := data["f_duration_key_none_val_unset"].(map[string]interface{})
				v1 := kv["1290111"].(map[string]interface{})
				require.Equal(t, float64(110111), v1["seconds"])
				require.Equal(t, float64(111), v1["nanos"])
				v2 := kv["1290112"].(map[string]interface{})
				require.Equal(t, float64(110112), v2["seconds"])
				require.Equal(t, float64(112), v2["nanos"])
				v3 := kv["1290113"].(map[string]interface{})
				require.Equal(t, float64(110113), v3["seconds"])
				require.Equal(t, float64(113), v3["nanos"])
			})
			t.Run("f_duration_key_none_val_object", func(t *testing.T) {
				kv := data["f_duration_key_none_val_object"].(map[string]interface{})
				v1 := kv["1290121"].(map[string]interface{})
				require.Equal(t, float64(110121), v1["seconds"])
				require.Equal(t, float64(121), v1["nanos"])
				v2 := kv["1290122"].(map[string]interface{})
				require.Equal(t, float64(110122), v2["seconds"])
				require.Equal(t, float64(122), v2["nanos"])
				v3 := kv["1290123"].(map[string]interface{})
				require.Equal(t, float64(110123), v3["seconds"])
				require.Equal(t, float64(123), v3["nanos"])
			})
			t.Run("f_duration_key_none_val_time_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_time_string"].(map[string]interface{})
				require.Equal(t, "30h35m31.000000131s", kv["1290131"])
				require.Equal(t, "30h35m32.000000132s", kv["1290132"])
				require.Equal(t, "30h35m33.000000133s", kv["1290133"])
			})
			t.Run("f_duration_key_none_val_nanosecond", func(t *testing.T) {
				kv := data["f_duration_key_none_val_nanosecond"].(map[string]interface{})
				require.Equal(t, float64(110141000000141), kv["1290141"])
				require.Equal(t, float64(110142000000142), kv["1290142"])
				require.Equal(t, float64(110143000000143), kv["1290143"])
			})
			t.Run("f_duration_key_none_val_nanosecond_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_nanosecond_string"].(map[string]interface{})
				require.Equal(t, "110151000000151", kv["1290151"])
				require.Equal(t, "110152000000152", kv["1290152"])
				require.Equal(t, "110153000000153", kv["1290153"])
			})
			t.Run("f_duration_key_none_val_microsecond", func(t *testing.T) {
				kv := data["f_duration_key_none_val_microsecond"].(map[string]interface{})
				require.Equal(t, float64(110161000161), kv["1290161"])
				require.Equal(t, float64(110162000162), kv["1290162"])
				require.Equal(t, float64(110163000163), kv["1290163"])
			})
			t.Run("f_duration_key_none_val_microsecond_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_microsecond_string"].(map[string]interface{})
				require.Equal(t, "110171000171", kv["1290171"])
				require.Equal(t, "110172000172", kv["1290172"])
				require.Equal(t, "110173000173", kv["1290173"])
			})
			t.Run("f_duration_key_none_val_millisecond", func(t *testing.T) {
				kv := data["f_duration_key_none_val_millisecond"].(map[string]interface{})
				require.Equal(t, float64(110181181), kv["1290181"])
				require.Equal(t, float64(110182182), kv["1290182"])
				require.Equal(t, float64(110183183), kv["1290183"])
			})
			t.Run("f_duration_key_none_val_millisecond_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_millisecond_string"].(map[string]interface{})
				require.Equal(t, "110191191", kv["1290191"])
				require.Equal(t, "110192192", kv["1290192"])
				require.Equal(t, "110193193", kv["1290193"])
			})
			t.Run("f_duration_key_none_val_second", func(t *testing.T) {
				kv := data["f_duration_key_none_val_second"].(map[string]interface{})
				require.Equal(t, float64(110211), kv["1290211"])
				require.Equal(t, float64(110212), kv["1290212"])
				require.Equal(t, float64(110213), kv["1290213"])
			})
			t.Run("f_duration_key_none_val_second_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_second_string"].(map[string]interface{})
				require.Equal(t, "110221", kv["1290221"])
				require.Equal(t, "110222", kv["1290222"])
				require.Equal(t, "110223", kv["1290223"])
			})
			t.Run("f_duration_key_none_val_minute", func(t *testing.T) {
				kv := data["f_duration_key_none_val_minute"].(map[string]interface{})
				require.Equal(t, float64(1837.1833333333334), kv["1290231"])
				require.Equal(t, float64(1837.2), kv["1290232"])
				require.Equal(t, float64(1837.2166666666667), kv["1290233"])
			})
			t.Run("f_duration_key_none_val_minute_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_minute_string"].(map[string]interface{})
				require.Equal(t, "1837.35", kv["1290241"])
				require.Equal(t, "1837.3666666666666", kv["1290242"])
				require.Equal(t, "1837.3833333333334", kv["1290243"])
			})
			t.Run("f_duration_key_none_val_hour", func(t *testing.T) {
				kv := data["f_duration_key_none_val_hour"].(map[string]interface{})
				require.Equal(t, float64(30.62527777777778), kv["1290251"])
				require.Equal(t, float64(30.625555555555554), kv["1290252"])
				require.Equal(t, float64(30.625833333333333), kv["1290253"])
			})
			t.Run("f_duration_key_none_val_hour_string", func(t *testing.T) {
				kv := data["f_duration_key_none_val_hour_string"].(map[string]interface{})
				require.Equal(t, "30.628055555555555", kv["1290261"])
				require.Equal(t, "30.628333333333334", kv["1290262"])
				require.Equal(t, "30.628611111111113", kv["1290263"])
			})
		})
		t.Run("Timestamp", func(t *testing.T) {
			t.Run("f_timestamp_key_none_val_unset", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unset"].(map[string]interface{})
				v1 := kv["1310111"].(map[string]interface{})
				require.Equal(t, float64(1718217911), v1["seconds"])
				require.Equal(t, float64(111), v1["nanos"])
				v2 := kv["1310112"].(map[string]interface{})
				require.Equal(t, float64(1718217912), v2["seconds"])
				require.Equal(t, float64(112), v2["nanos"])
				v3 := kv["1310113"].(map[string]interface{})
				require.Equal(t, float64(1718217913), v3["seconds"])
				require.Equal(t, float64(113), v3["nanos"])
			})
			t.Run("f_timestamp_key_none_val_object", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_object"].(map[string]interface{})
				v1 := kv["1310121"].(map[string]interface{})
				require.Equal(t, float64(1718217921), v1["seconds"])
				require.Equal(t, float64(121), v1["nanos"])
				v2 := kv["1310122"].(map[string]interface{})
				require.Equal(t, float64(1718217922), v2["seconds"])
				require.Equal(t, float64(122), v2["nanos"])
				v3 := kv["1310123"].(map[string]interface{})
				require.Equal(t, float64(1718217923), v3["seconds"])
				require.Equal(t, float64(123), v3["nanos"])
			})
			t.Run("f_timestamp_key_none_val_time_layout", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_time_layout"].(map[string]interface{})
				require.Equal(t, "2024-06-12T18:45:31.000000131Z", kv["1310131"])
				require.Equal(t, "2024-06-12T18:45:32.000000132Z", kv["1310132"])
				require.Equal(t, "2024-06-12T18:45:33.000000133Z", kv["1310133"])
			})
			t.Run("f_timestamp_key_none_val_unix_nano", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_nano"].(map[string]interface{})
				require.Equal(t, float64(1718217941000000141), kv["1310141"])
				require.Equal(t, float64(1718217942000000142), kv["1310142"])
				require.Equal(t, float64(1718217943000000143), kv["1310143"])
			})
			t.Run("f_timestamp_key_none_val_unix_nano_string", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_nano_string"].(map[string]interface{})
				require.Equal(t, "1718217951000000151", kv["1310151"])
				require.Equal(t, "1718217952000000152", kv["1310152"])
				require.Equal(t, "1718217953000000153", kv["1310153"])
			})
			t.Run("f_timestamp_key_none_val_unix_micro", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_micro"].(map[string]interface{})
				require.Equal(t, float64(1718217961000161), kv["1310161"])
				require.Equal(t, float64(1718217962000162), kv["1310162"])
				require.Equal(t, float64(1718217963000163), kv["1310163"])
			})
			t.Run("f_timestamp_key_none_val_unix_micro_string", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_micro_string"].(map[string]interface{})
				require.Equal(t, "1718217971000171", kv["1310171"])
				require.Equal(t, "1718217972000172", kv["1310172"])
				require.Equal(t, "1718217973000173", kv["1310173"])
			})
			t.Run("f_timestamp_key_none_val_unix_milli", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_milli"].(map[string]interface{})
				require.Equal(t, float64(1718217981181), kv["1310181"])
				require.Equal(t, float64(1718217982182), kv["1310182"])
				require.Equal(t, float64(1718217983183), kv["1310183"])
			})
			t.Run("f_timestamp_key_none_val_unix_milli_string", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_milli_string"].(map[string]interface{})
				require.Equal(t, "1718217991191", kv["1310191"])
				require.Equal(t, "1718217992192", kv["1310192"])
				require.Equal(t, "1718217993193", kv["1310193"])
			})
			t.Run("f_timestamp_key_none_val_unix_sec", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_sec"].(map[string]interface{})
				require.Equal(t, float64(1718334211), kv["1310211"])
				require.Equal(t, float64(1718334212), kv["1310212"])
				require.Equal(t, float64(1718334213), kv["1310213"])
			})
			t.Run("f_timestamp_key_none_val_unix_sec_string", func(t *testing.T) {
				kv := data["f_timestamp_key_none_val_unix_sec_string"].(map[string]interface{})
				require.Equal(t, "1718334221", kv["1310221"])
				require.Equal(t, "1718334222", kv["1310222"])
				require.Equal(t, "1718334223", kv["1310223"])
			})
		})
		t.Run("Any", func(t *testing.T) {
			t.Run("f_any_key_none_val_unset", func(t *testing.T) {
				kv := data["f_any_key_none_val_unset"].(map[string]interface{})
				v1 := kv["1320111"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["type_url"])
				require.Equal(t, "CgRhMTExEgRhMTEyGgRhMTEz", v1["value"])
				v2 := kv["1320121"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["type_url"])
				require.Equal(t, "CgRhMTIxEgRhMTIyGgRhMTIz", v2["value"])
				v3 := kv["1320131"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v3["type_url"])
				require.Equal(t, "CgRhMTMxEgRhMTMyGgRhMTMz", v3["value"])
			})
			t.Run("f_any_key_none_val_object", func(t *testing.T) {
				kv := data["f_any_key_none_val_object"].(map[string]interface{})
				v1 := kv["1320211"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["type_url"])
				require.Equal(t, "CgRhMjExEgRhMjEyGgRhMjEz", v1["value"])
				v2 := kv["1320221"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["type_url"])
				require.Equal(t, "CgRhMjIxEgRhMjIyGgRhMjIz", v2["value"])
				v3 := kv["1320231"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v3["type_url"])
				require.Equal(t, "CgRhMjMxEgRhMjMyGgRhMjMz", v3["value"])
			})
			t.Run("f_any_key_none_val_proto", func(t *testing.T) {
				kv := data["f_any_key_none_val_proto"].(map[string]interface{})
				v1 := kv["1320311"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["@type"])
				require.Equal(t, "a311", v1["f_string1"])
				require.Equal(t, "a312", v1["f_string2"])
				require.Equal(t, "a313", v1["f_string3"])
				v2 := kv["1320321"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["@type"])
				require.Equal(t, "a321", v2["f_string1"])
				require.Equal(t, "a322", v2["f_string2"])
				require.Equal(t, "a323", v2["f_string3"])
				v3 := kv["1320331"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v3["@type"])
				require.Equal(t, "a331", v3["f_string1"])
				require.Equal(t, "a332", v3["f_string2"])
				require.Equal(t, "a333", v3["f_string3"])
			})
		})
	})
}
