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

func Test_CodecMapKeyUInt64(t *testing.T) {
	seed1 := &pbformat.CodecMapKeyUInt64{
		FInt32KeyUnsetValUnset:          map[uint64]int32{1110111: 2110111, 1110112: 2110112, 1110113: 2110113},
		FInt32KeyNumericValNumeric:      map[uint64]int32{1110211: 2110211, 1110212: 2110212, 1110213: 2110213},
		FInt32KeyNumericValString:       map[uint64]int32{1110311: 2110311, 1110312: 2110312, 1110313: 2110313},
		FInt32KeyStringValNumeric:       map[uint64]int32{1110411: 2110411, 1110412: 2110412, 1110413: 2110413},
		FInt32KeyStringValString:        map[uint64]int32{1110511: 2110511, 1110512: 2110512, 1110513: 2110513},
		FInt64KeyUnsetValUnset:          map[uint64]int64{1120111: 2120111, 1120112: 2120112, 1120113: 2120113},
		FInt64KeyNumericValNumeric:      map[uint64]int64{1120211: 2120211, 1120212: 2120212, 1120213: 2120213},
		FInt64KeyNumericValString:       map[uint64]int64{1120311: 2120311, 1120312: 2120312, 1120313: 2120313},
		FInt64KeyStringValNumeric:       map[uint64]int64{1120411: 2120411, 1120412: 2120412, 1120413: 2120413},
		FInt64KeyStringValString:        map[uint64]int64{1120511: 2120511, 1120512: 2120512, 1120513: 2120513},
		FUint32KeyUnsetValUnset:         map[uint64]uint32{1130111: 2130111, 1130112: 2130112, 1130113: 2130113},
		FUint32KeyNumericValNumeric:     map[uint64]uint32{1130211: 2130211, 1130212: 2130212, 1130213: 2130213},
		FUint32KeyNumericValString:      map[uint64]uint32{1130311: 2130311, 1130312: 2130312, 1130313: 2130313},
		FUint32KeyStringValNumeric:      map[uint64]uint32{1130411: 2130411, 1130412: 2130412, 1130413: 2130413},
		FUint32KeyStringValString:       map[uint64]uint32{1130511: 2130511, 1130512: 2130512, 1130513: 2130513},
		FUint64KeyUnsetValUnset:         map[uint64]uint64{1140111: 2140111, 1140112: 2140112, 1140113: 2140113},
		FUint64KeyNumericValNumeric:     map[uint64]uint64{1140211: 2140211, 1140212: 2140212, 1140213: 2140213},
		FUint64KeyNumericValString:      map[uint64]uint64{1140311: 2140311, 1140312: 2140312, 1140313: 2140313},
		FUint64KeyStringValNumeric:      map[uint64]uint64{1140411: 2140411, 1140412: 2140412, 1140413: 2140413},
		FUint64KeyStringValString:       map[uint64]uint64{1140511: 2140511, 1140512: 2140512, 1140513: 2140513},
		FSint32KeyUnsetValUnset:         map[uint64]int32{1150111: 2150111, 1150112: 2150112, 1150113: 2150113},
		FSint32KeyNumericValNumeric:     map[uint64]int32{1150211: 2150211, 1150212: 2150212, 1150213: 2150213},
		FSint32KeyNumericValString:      map[uint64]int32{1150311: 2150311, 1150312: 2150312, 1150313: 2150313},
		FSint32KeyStringValNumeric:      map[uint64]int32{1150411: 2150411, 1150412: 2150412, 1150413: 2150413},
		FSint32KeyStringValString:       map[uint64]int32{1150511: 2150511, 1150512: 2150512, 1150513: 2150513},
		FSint64KeyUnsetValUnset:         map[uint64]int64{1160111: 2160111, 1160112: 2160112, 1160113: 2160113},
		FSint64KeyNumericValNumeric:     map[uint64]int64{1160211: 2160211, 1160212: 2160212, 1160213: 2160213},
		FSint64KeyNumericValString:      map[uint64]int64{1160311: 2160311, 1160312: 2160312, 1160313: 2160313},
		FSint64KeyStringValNumeric:      map[uint64]int64{1160411: 2160411, 1160412: 2160412, 1160413: 2160413},
		FSint64KeyStringValString:       map[uint64]int64{1160511: 2160511, 1160512: 2160512, 1160513: 2160513},
		FSfixed32KeyUnsetValUnset:       map[uint64]int32{1170111: 2170111, 1170112: 2170112, 1170113: 2170113},
		FSfixed32KeyNumericValNumeric:   map[uint64]int32{1170211: 2170211, 1170212: 2170212, 1170213: 2170213},
		FSfixed32KeyNumericValString:    map[uint64]int32{1170311: 2170311, 1170312: 2170312, 1170313: 2170313},
		FSfixed32KeyStringValNumeric:    map[uint64]int32{1170411: 2170411, 1170412: 2170412, 1170413: 2170413},
		FSfixed32KeyStringValString:     map[uint64]int32{1170511: 2170511, 1170512: 2170512, 1170513: 2170513},
		FSfixed64KeyUnsetValUnset:       map[uint64]int64{1180111: 2180111, 1180112: 2180112, 1180113: 2180113},
		FSfixed64KeyNumericValNumeric:   map[uint64]int64{1180211: 2180211, 1180212: 2180212, 1180213: 2180213},
		FSfixed64KeyNumericValString:    map[uint64]int64{1180311: 2180311, 1180312: 2180312, 1180313: 2180313},
		FSfixed64KeyStringValNumeric:    map[uint64]int64{1180411: 2180411, 1180412: 2180412, 1180413: 2180413},
		FSfixed64KeyStringValString:     map[uint64]int64{1180511: 2180511, 1180512: 2180512, 1180513: 2180513},
		FFixed32KeyUnsetValUnset:        map[uint64]uint32{1190111: 2190111, 1190112: 2190112, 1190113: 2190113},
		FFixed32KeyNumericValNumeric:    map[uint64]uint32{1190211: 2190211, 1190212: 2190212, 1190213: 2190213},
		FFixed32KeyNumericValString:     map[uint64]uint32{1190311: 2190311, 1190312: 2190312, 1190313: 2190313},
		FFixed32KeyStringValNumeric:     map[uint64]uint32{1190411: 2190411, 1190412: 2190412, 1190413: 2190413},
		FFixed32KeyStringValString:      map[uint64]uint32{1190511: 2190511, 1190512: 2190512, 1190513: 2190513},
		FFixed64KeyUnsetValUnset:        map[uint64]uint64{1210111: 2210111, 1210112: 2210112, 1210113: 2210113},
		FFixed64KeyNumericValNumeric:    map[uint64]uint64{1210211: 2210211, 1210212: 2210212, 1210213: 2210213},
		FFixed64KeyNumericValString:     map[uint64]uint64{1210311: 2210311, 1210312: 2210312, 1210313: 2210313},
		FFixed64KeyStringValNumeric:     map[uint64]uint64{1210411: 2210411, 1210412: 2210412, 1210413: 2210413},
		FFixed64KeyStringValString:      map[uint64]uint64{1210511: 2210511, 1210512: 2210512, 1210513: 2210513},
		FFloatKeyUnsetValUnset:          map[uint64]float32{1220111: 2220.111, 1220112: 2220.112, 1220113: 2220.113},
		FFloatKeyNumericValNumeric:      map[uint64]float32{1220211: 2220.211, 1220212: 2220.212, 1220213: 2220.213},
		FFloatKeyNumericValString:       map[uint64]float32{1220311: 2220.311, 1220312: 2220.312, 1220313: 2220.313},
		FFloatKeyStringValNumeric:       map[uint64]float32{1220411: 2220.411, 1220412: 2220.412, 1220413: 2220.413},
		FFloatKeyStringValString:        map[uint64]float32{1220511: 2220.511, 1220512: 2220.512, 1220513: 2220.513},
		FDoubleKeyUnsetValUnset:         map[uint64]float64{1230111: 2230111.111, 1230112: 2230112.111, 1230113: 2230113.111},
		FDoubleKeyNumericValNumeric:     map[uint64]float64{1230211: 2230211.111, 1230212: 2230212.111, 1230213: 2230213.111},
		FDoubleKeyNumericValString:      map[uint64]float64{1230311: 2230311.111, 1230312: 2230312.111, 1230313: 2230313.111},
		FDoubleKeyStringValNumeric:      map[uint64]float64{1230411: 2230411.111, 1230412: 2230412.111, 1230413: 2230413.111},
		FDoubleKeyStringValString:       map[uint64]float64{1230511: 2230511.111, 1230512: 2230512.111, 1230513: 2230513.111},
		FBoolKeyUnsetValUnset:           map[uint64]bool{1240111: true, 1240112: true, 1240113: true},
		FBoolKeyNumericValBool:          map[uint64]bool{1240211: true, 1240212: false, 1240213: true},
		FBoolKeyNumericValString:        map[uint64]bool{1240311: false, 1240312: true, 1240313: false},
		FBoolKeyStringValBool:           map[uint64]bool{1240411: true, 1240412: false, 1240413: true},
		FBoolKeyStringValString:         map[uint64]bool{1240511: false, 1240512: true, 1240513: false},
		FStringKeyUnsetValNone:          map[uint64]string{1250111: "s101", 1250112: "s102", 1250113: "s103"},
		FStringKeyNumericValNone:        map[uint64]string{1250211: "s201", 1250212: "s202", 1250213: "s203"},
		FStringKeyStringValNone:         map[uint64]string{1250311: "s301", 1250312: "s302", 1250313: "s303"},
		FBytesKeyUnsetValNone:           map[uint64][]byte{1260111: []byte("b101"), 1260112: []byte("b102"), 1260113: []byte("b103")},
		FBytesKeyNumericValNone:         map[uint64][]byte{1260211: []byte("b201"), 1260212: []byte("b202"), 1260213: []byte("b203")},
		FBytesKeyStringValNone:          map[uint64][]byte{1260311: []byte("b301"), 1260312: []byte("b302"), 1260313: []byte("b303")},
		FEnumKeyUnsetValUnset:           map[uint64]pbexternal.EnumNum1{1270111: 0, 1270112: 2, 1270113: 3},
		FEnumKeyNumericValNumeric:       map[uint64]pbexternal.EnumNum1{1270211: 2, 1270212: 3, 1270213: 5},
		FEnumKeyNumericValNumericString: map[uint64]pbexternal.EnumNum1{1270311: 3, 1270312: 5, 1270313: 6},
		FEnumKeyNumericValEnumString:    map[uint64]pbexternal.EnumNum1{1270411: 5, 1270412: 6, 1270413: 7},
		FEnumKeyStringValNumeric:        map[uint64]pbexternal.EnumNum1{1270511: 6, 1270512: 7, 1270513: 9},
		FEnumKeyStringValNumericString:  map[uint64]pbexternal.EnumNum1{1270611: 7, 1270612: 9, 1270613: 11},
		FEnumKeyStringValEnumString:     map[uint64]pbexternal.EnumNum1{1270711: 9, 1270712: 11, 1270713: 12},
		FMessageKeyUnsetValNone: map[uint64]*pbexternal.Message1{
			1280111: {FString1: "m101", FString2: "m102", FString3: "m103"},
			1280112: {FString1: "m201", FString2: "m202", FString3: "m203"},
			1280113: {FString1: "m301", FString2: "m302", FString3: "m303"},
		},
		FMessageKeyNumericValNone: map[uint64]*pbexternal.Message1{
			1280211: {FString1: "m111", FString2: "m112", FString3: "m113"},
			1280212: {FString1: "m211", FString2: "m212", FString3: "m213"},
			1280213: {FString1: "m311", FString2: "m312", FString3: "m313"},
		},
		FMessageKeyStringValNone: map[uint64]*pbexternal.Message1{
			1280311: {FString1: "m121", FString2: "m122", FString3: "m123"},
			1280312: {FString1: "m221", FString2: "m222", FString3: "m223"},
			1280313: {FString1: "m321", FString2: "m322", FString3: "m323"},
		},
		FDurationKeyUnsetValUnset: map[uint64]*durationpb.Duration{
			1290111: {Seconds: 110111, Nanos: 111},
			1290112: {Seconds: 110112, Nanos: 112},
			1290113: {Seconds: 110113, Nanos: 113},
		},
		FDurationKeyNumericValObject: map[uint64]*durationpb.Duration{
			1290121: {Seconds: 110121, Nanos: 121},
			1290122: {Seconds: 110122, Nanos: 122},
			1290123: {Seconds: 110123, Nanos: 123},
		},
		FDurationKeyNumericValTimeString: map[uint64]*durationpb.Duration{
			1290131: {Seconds: 110131, Nanos: 131},
			1290132: {Seconds: 110132, Nanos: 132},
			1290133: {Seconds: 110133, Nanos: 133},
		},
		FDurationKeyNumericValNanosecond: map[uint64]*durationpb.Duration{
			1290141: {Seconds: 110141, Nanos: 141},
			1290142: {Seconds: 110142, Nanos: 142},
			1290143: {Seconds: 110143, Nanos: 143},
		},
		FDurationKeyNumericValNanosecondString: map[uint64]*durationpb.Duration{
			1290151: {Seconds: 110151, Nanos: 151},
			1290152: {Seconds: 110152, Nanos: 152},
			1290153: {Seconds: 110153, Nanos: 153},
		},
		FDurationKeyNumericValMicrosecond: map[uint64]*durationpb.Duration{
			1290161: {Seconds: 110161, Nanos: 161000},
			1290162: {Seconds: 110162, Nanos: 162000},
			1290163: {Seconds: 110163, Nanos: 163000},
		},
		FDurationKeyNumericValMicrosecondString: map[uint64]*durationpb.Duration{
			1290171: {Seconds: 110171, Nanos: 171000},
			1290172: {Seconds: 110172, Nanos: 172000},
			1290173: {Seconds: 110173, Nanos: 173000},
		},
		FDurationKeyNumericValMillisecond: map[uint64]*durationpb.Duration{
			1290181: {Seconds: 110181, Nanos: 181000000},
			1290182: {Seconds: 110182, Nanos: 182000000},
			1290183: {Seconds: 110183, Nanos: 183000000},
		},
		FDurationKeyNumericValMillisecondString: map[uint64]*durationpb.Duration{
			1290191: {Seconds: 110191, Nanos: 191000000},
			1290192: {Seconds: 110192, Nanos: 192000000},
			1290193: {Seconds: 110193, Nanos: 193000000},
		},
		FDurationKeyNumericValSecond: map[uint64]*durationpb.Duration{
			1290211: {Seconds: 110211, Nanos: 0},
			1290212: {Seconds: 110212, Nanos: 0},
			1290213: {Seconds: 110213, Nanos: 0},
		},
		FDurationKeyNumericValSecondString: map[uint64]*durationpb.Duration{
			1290221: {Seconds: 110221, Nanos: 0},
			1290222: {Seconds: 110222, Nanos: 0},
			1290223: {Seconds: 110223, Nanos: 0},
		},
		FDurationKeyNumericValMinute: map[uint64]*durationpb.Duration{
			1290231: {Seconds: 110231, Nanos: 0},
			1290232: {Seconds: 110232, Nanos: 0},
			1290233: {Seconds: 110233, Nanos: 0},
		},
		FDurationKeyNumericValMinuteString: map[uint64]*durationpb.Duration{
			1290241: {Seconds: 110241, Nanos: 0},
			1290242: {Seconds: 110242, Nanos: 0},
			1290243: {Seconds: 110243, Nanos: 0},
		},
		FDurationKeyNumericValHour: map[uint64]*durationpb.Duration{
			1290251: {Seconds: 110251, Nanos: 0},
			1290252: {Seconds: 110252, Nanos: 0},
			1290253: {Seconds: 110253, Nanos: 0},
		},
		FDurationKeyNumericValHourString: map[uint64]*durationpb.Duration{
			1290261: {Seconds: 110261, Nanos: 0},
			1290262: {Seconds: 110262, Nanos: 0},
			1290263: {Seconds: 110263, Nanos: 0},
		},
		FDurationKeyStringValUnset: map[uint64]*durationpb.Duration{
			1290271: {Seconds: 110271, Nanos: 271},
			1290272: {Seconds: 110272, Nanos: 272},
			1290273: {Seconds: 110273, Nanos: 273},
		},
		FDurationKeyStringValObject: map[uint64]*durationpb.Duration{
			1290281: {Seconds: 110281, Nanos: 281},
			1290282: {Seconds: 110282, Nanos: 282},
			1290283: {Seconds: 110283, Nanos: 283},
		},
		FDurationKeyStringValTimeString: map[uint64]*durationpb.Duration{
			1290291: {Seconds: 110291, Nanos: 291},
			1290292: {Seconds: 110292, Nanos: 292},
			1290293: {Seconds: 110293, Nanos: 293},
		},
		FDurationKeyStringValNanosecond: map[uint64]*durationpb.Duration{
			1290311: {Seconds: 110311, Nanos: 311},
			1290312: {Seconds: 110312, Nanos: 312},
			1290313: {Seconds: 110313, Nanos: 313},
		},
		FDurationKeyStringValNanosecondString: map[uint64]*durationpb.Duration{
			1290321: {Seconds: 110321, Nanos: 321},
			1290322: {Seconds: 110322, Nanos: 322},
			1290323: {Seconds: 110323, Nanos: 323},
		},
		FDurationKeyStringValMicrosecond: map[uint64]*durationpb.Duration{
			1290331: {Seconds: 110331, Nanos: 331000},
			1290332: {Seconds: 110332, Nanos: 332000},
			1290333: {Seconds: 110333, Nanos: 333000},
		},
		FDurationKeyStringValMicrosecondString: map[uint64]*durationpb.Duration{
			1290341: {Seconds: 110341, Nanos: 341000},
			1290342: {Seconds: 110342, Nanos: 342000},
			1290343: {Seconds: 110343, Nanos: 343000},
		},
		FDurationKeyStringValMillisecond: map[uint64]*durationpb.Duration{
			1290351: {Seconds: 110351, Nanos: 351000000},
			1290352: {Seconds: 110352, Nanos: 352000000},
			1290353: {Seconds: 110353, Nanos: 353000000},
		},
		FDurationKeyStringValMillisecondString: map[uint64]*durationpb.Duration{
			1290361: {Seconds: 110361, Nanos: 361000000},
			1290362: {Seconds: 110362, Nanos: 362000000},
			1290363: {Seconds: 110363, Nanos: 363000000},
		},
		FDurationKeyStringValSecond: map[uint64]*durationpb.Duration{
			1290371: {Seconds: 110371, Nanos: 0},
			1290372: {Seconds: 110372, Nanos: 0},
			1290373: {Seconds: 110373, Nanos: 0},
		},
		FDurationKeyStringValSecondString: map[uint64]*durationpb.Duration{
			1290381: {Seconds: 110381, Nanos: 0},
			1290382: {Seconds: 110382, Nanos: 0},
			1290383: {Seconds: 110383, Nanos: 0},
		},
		FDurationKeyStringValMinute: map[uint64]*durationpb.Duration{
			1290391: {Seconds: 110391, Nanos: 0},
			1290392: {Seconds: 110392, Nanos: 0},
			1290393: {Seconds: 110393, Nanos: 0},
		},
		FDurationKeyStringValMinuteString: map[uint64]*durationpb.Duration{
			1290411: {Seconds: 110411, Nanos: 0},
			1290412: {Seconds: 110412, Nanos: 0},
			1290413: {Seconds: 110413, Nanos: 0},
		},
		FDurationKeyStringValHour: map[uint64]*durationpb.Duration{
			1290421: {Seconds: 110421, Nanos: 0},
			1290422: {Seconds: 110422, Nanos: 0},
			1290423: {Seconds: 110423, Nanos: 0},
		},
		FDurationKeyStringValHourString: map[uint64]*durationpb.Duration{
			1290431: {Seconds: 110431, Nanos: 0},
			1290432: {Seconds: 110432, Nanos: 0},
			1290433: {Seconds: 110433, Nanos: 0},
		},
		FTimestampKeyUnsetValUnset: map[uint64]*timestamppb.Timestamp{
			1310111: {Seconds: 1718217911, Nanos: 111},
			1310112: {Seconds: 1718217912, Nanos: 112},
			1310113: {Seconds: 1718217913, Nanos: 113},
		},
		FTimestampKeyNumericValObject: map[uint64]*timestamppb.Timestamp{
			1310121: {Seconds: 1718217921, Nanos: 121},
			1310122: {Seconds: 1718217922, Nanos: 122},
			1310123: {Seconds: 1718217923, Nanos: 123},
		},
		FTimestampKeyNumericValTimeLayout: map[uint64]*timestamppb.Timestamp{
			1310131: {Seconds: 1718217931, Nanos: 131},
			1310132: {Seconds: 1718217932, Nanos: 132},
			1310133: {Seconds: 1718217933, Nanos: 133},
		},
		FTimestampKeyNumericValUnixNano: map[uint64]*timestamppb.Timestamp{
			1310141: {Seconds: 1718217941, Nanos: 141},
			1310142: {Seconds: 1718217942, Nanos: 142},
			1310143: {Seconds: 1718217943, Nanos: 143},
		},
		FTimestampKeyNumericValUnixNanoString: map[uint64]*timestamppb.Timestamp{
			1310151: {Seconds: 1718217951, Nanos: 151},
			1310152: {Seconds: 1718217952, Nanos: 152},
			1310153: {Seconds: 1718217953, Nanos: 153},
		},
		FTimestampKeyNumericValUnixMicro: map[uint64]*timestamppb.Timestamp{
			1310161: {Seconds: 1718217961, Nanos: 161000},
			1310162: {Seconds: 1718217962, Nanos: 162000},
			1310163: {Seconds: 1718217963, Nanos: 163000},
		},
		FTimestampKeyNumericValUnixMicroString: map[uint64]*timestamppb.Timestamp{
			1310171: {Seconds: 1718217971, Nanos: 171000},
			1310172: {Seconds: 1718217972, Nanos: 172000},
			1310173: {Seconds: 1718217973, Nanos: 173000},
		},
		FTimestampKeyNumericValUnixMilli: map[uint64]*timestamppb.Timestamp{
			1310181: {Seconds: 1718217981, Nanos: 181000000},
			1310182: {Seconds: 1718217982, Nanos: 182000000},
			1310183: {Seconds: 1718217983, Nanos: 183000000},
		},
		FTimestampKeyNumericValUnixMilliString: map[uint64]*timestamppb.Timestamp{
			1310191: {Seconds: 1718217991, Nanos: 191000000},
			1310192: {Seconds: 1718217992, Nanos: 192000000},
			1310193: {Seconds: 1718217993, Nanos: 193000000},
		},
		FTimestampKeyNumericValUnixSec: map[uint64]*timestamppb.Timestamp{
			1310211: {Seconds: 1718334211, Nanos: 0},
			1310212: {Seconds: 1718334212, Nanos: 0},
			1310213: {Seconds: 1718334213, Nanos: 0},
		},
		FTimestampKeyNumericValUnixSecString: map[uint64]*timestamppb.Timestamp{
			1310221: {Seconds: 1718334221, Nanos: 0},
			1310222: {Seconds: 1718334222, Nanos: 0},
			1310223: {Seconds: 1718334223, Nanos: 0},
		},
		FTimestampKeyStringValObject: map[uint64]*timestamppb.Timestamp{
			1310231: {Seconds: 1718334231, Nanos: 231},
			1310232: {Seconds: 1718334232, Nanos: 232},
			1310233: {Seconds: 1718334233, Nanos: 233},
		},
		FTimestampKeyStringValTimeLayout: map[uint64]*timestamppb.Timestamp{
			1310241: {Seconds: 1718334241, Nanos: 241},
			1310242: {Seconds: 1718334242, Nanos: 242},
			1310243: {Seconds: 1718334243, Nanos: 243},
		},
		FTimestampKeyStringValUnixNano: map[uint64]*timestamppb.Timestamp{
			1310251: {Seconds: 1718334251, Nanos: 251},
			1310252: {Seconds: 1718334252, Nanos: 252},
			1310253: {Seconds: 1718334253, Nanos: 253},
		},
		FTimestampKeyStringValUnixNanoString: map[uint64]*timestamppb.Timestamp{
			1310261: {Seconds: 1718334261, Nanos: 261},
			1310262: {Seconds: 1718334262, Nanos: 262},
			1310263: {Seconds: 1718334263, Nanos: 263},
		},
		FTimestampKeyStringValUnixMicro: map[uint64]*timestamppb.Timestamp{
			1310271: {Seconds: 1718334271, Nanos: 271000},
			1310272: {Seconds: 1718334272, Nanos: 272000},
			1310273: {Seconds: 1718334273, Nanos: 273000},
		},
		FTimestampKeyStringValUnixMicroString: map[uint64]*timestamppb.Timestamp{
			1310281: {Seconds: 1718334281, Nanos: 281000},
			1310282: {Seconds: 1718334282, Nanos: 282000},
			1310283: {Seconds: 1718334283, Nanos: 283000},
		},
		FTimestampKeyStringValUnixMilli: map[uint64]*timestamppb.Timestamp{
			1310311: {Seconds: 1718334311, Nanos: 311000000},
			1310312: {Seconds: 1718334312, Nanos: 312000000},
			1310313: {Seconds: 1718334313, Nanos: 313000000},
		},
		FTimestampKeyStringValUnixMilliString: map[uint64]*timestamppb.Timestamp{
			1310321: {Seconds: 1718334321, Nanos: 321000000},
			1310322: {Seconds: 1718334322, Nanos: 322000000},
			1310323: {Seconds: 1718334323, Nanos: 323000000},
		},
		FTimestampKeyStringValUnixSec: map[uint64]*timestamppb.Timestamp{
			1310331: {Seconds: 1718334331, Nanos: 0},
			1310332: {Seconds: 1718334332, Nanos: 0},
			1310333: {Seconds: 1718334333, Nanos: 0},
		},
		FTimestampKeyStringValUnixSecString: map[uint64]*timestamppb.Timestamp{
			1310341: {Seconds: 1718334341, Nanos: 0},
			1310342: {Seconds: 1718334342, Nanos: 0},
			1310343: {Seconds: 1718334343, Nanos: 0},
		},
		FAnyKeyUnsetValUnset: map[uint64]*anypb.Any{
			1320111: utils.MustNewAny(&pbexternal.Message1{FString1: "a111", FString2: "a112", FString3: "a113"}),
			1320121: utils.MustNewAny(&pbexternal.Message1{FString1: "a121", FString2: "a122", FString3: "a123"}),
			1320131: utils.MustNewAny(&pbexternal.Message1{FString1: "a131", FString2: "a132", FString3: "a133"}),
		},
		FAnyKeyNumericValObject: map[uint64]*anypb.Any{
			1320211: utils.MustNewAny(&pbexternal.Message1{FString1: "a211", FString2: "a212", FString3: "a213"}),
			1320221: utils.MustNewAny(&pbexternal.Message1{FString1: "a221", FString2: "a222", FString3: "a223"}),
			1320231: utils.MustNewAny(&pbexternal.Message1{FString1: "a231", FString2: "a232", FString3: "a233"}),
		},
		FAnyKeyNumericValProto: map[uint64]*anypb.Any{
			1320311: utils.MustNewAny(&pbexternal.Message1{FString1: "a311", FString2: "a312", FString3: "a313"}),
			1320321: utils.MustNewAny(&pbexternal.Message1{FString1: "a321", FString2: "a322", FString3: "a323"}),
			1320331: utils.MustNewAny(&pbexternal.Message1{FString1: "a331", FString2: "a332", FString3: "a333"}),
		},
		FAnyKeyStringValObject: map[uint64]*anypb.Any{
			1320411: utils.MustNewAny(&pbexternal.Message1{FString1: "a411", FString2: "a412", FString3: "a413"}),
			1320421: utils.MustNewAny(&pbexternal.Message1{FString1: "a421", FString2: "a422", FString3: "a423"}),
			1320431: utils.MustNewAny(&pbexternal.Message1{FString1: "a431", FString2: "a432", FString3: "a433"}),
		},
		FAnyKeyStringValProto: map[uint64]*anypb.Any{
			1320511: utils.MustNewAny(&pbexternal.Message1{FString1: "a511", FString2: "a512", FString3: "a513"}),
			1320521: utils.MustNewAny(&pbexternal.Message1{FString1: "a521", FString2: "a522", FString3: "a523"}),
			1320531: utils.MustNewAny(&pbexternal.Message1{FString1: "a531", FString2: "a532", FString3: "a533"}),
		},
	}

	t.Run("MarshalAndUnmarshal", func(t *testing.T) {
		bb, err := seed1.MarshalJSON()
		require.Nil(t, err)

		dataNew := &pbformat.CodecMapKeyUInt64{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})

	t.Run("Check", func(t *testing.T) {
		{ // make the field to nil which not support with standard JSON.
			seed1.FInt32KeyNumericValNumeric = nil
			seed1.FInt32KeyNumericValString = nil

			seed1.FInt64KeyNumericValNumeric = nil
			seed1.FInt64KeyNumericValString = nil

			seed1.FUint32KeyNumericValNumeric = nil
			seed1.FUint32KeyNumericValString = nil

			seed1.FUint64KeyNumericValNumeric = nil
			seed1.FUint64KeyNumericValString = nil

			seed1.FSint32KeyNumericValNumeric = nil
			seed1.FSint32KeyNumericValString = nil

			seed1.FSint64KeyNumericValNumeric = nil
			seed1.FSint64KeyNumericValString = nil

			seed1.FSfixed32KeyNumericValNumeric = nil
			seed1.FSfixed32KeyNumericValString = nil

			seed1.FSfixed64KeyNumericValNumeric = nil
			seed1.FSfixed64KeyNumericValString = nil

			seed1.FFixed32KeyNumericValNumeric = nil
			seed1.FFixed32KeyNumericValString = nil

			seed1.FFixed64KeyNumericValNumeric = nil
			seed1.FFixed64KeyNumericValString = nil

			seed1.FFloatKeyNumericValNumeric = nil
			seed1.FFloatKeyNumericValString = nil

			seed1.FDoubleKeyNumericValNumeric = nil
			seed1.FDoubleKeyNumericValString = nil

			seed1.FBoolKeyNumericValBool = nil
			seed1.FBoolKeyNumericValString = nil

			seed1.FStringKeyNumericValNone = nil

			seed1.FBytesKeyNumericValNone = nil

			seed1.FEnumKeyNumericValNumeric = nil
			seed1.FEnumKeyNumericValNumericString = nil
			seed1.FEnumKeyNumericValEnumString = nil

			seed1.FMessageKeyNumericValNone = nil

			seed1.FDurationKeyNumericValObject = nil
			seed1.FDurationKeyNumericValTimeString = nil
			seed1.FDurationKeyNumericValNanosecond = nil
			seed1.FDurationKeyNumericValNanosecondString = nil
			seed1.FDurationKeyNumericValMicrosecond = nil
			seed1.FDurationKeyNumericValMicrosecondString = nil
			seed1.FDurationKeyNumericValMillisecond = nil
			seed1.FDurationKeyNumericValMillisecondString = nil
			seed1.FDurationKeyNumericValSecond = nil
			seed1.FDurationKeyNumericValSecondString = nil
			seed1.FDurationKeyNumericValMinute = nil
			seed1.FDurationKeyNumericValMinuteString = nil
			seed1.FDurationKeyNumericValHour = nil
			seed1.FDurationKeyNumericValHourString = nil

			seed1.FTimestampKeyNumericValObject = nil
			seed1.FTimestampKeyNumericValTimeLayout = nil
			seed1.FTimestampKeyNumericValUnixNano = nil
			seed1.FTimestampKeyNumericValUnixNanoString = nil
			seed1.FTimestampKeyNumericValUnixMicro = nil
			seed1.FTimestampKeyNumericValUnixMicroString = nil
			seed1.FTimestampKeyNumericValUnixMilli = nil
			seed1.FTimestampKeyNumericValUnixMilliString = nil
			seed1.FTimestampKeyNumericValUnixSec = nil
			seed1.FTimestampKeyNumericValUnixSecString = nil

			seed1.FAnyKeyNumericValObject = nil
			seed1.FAnyKeyNumericValProto = nil
		}

		bb, err := seed1.MarshalJSON()
		require.Nil(t, err)

		data := make(map[string]interface{})
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		t.Run("Int32", func(t *testing.T) {
			t.Run("f_int32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_int32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2110111), kv["1110111"])
				require.Equal(t, float64(2110112), kv["1110112"])
				require.Equal(t, float64(2110113), kv["1110113"])
			})
			t.Run("f_int32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_int32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2110411), kv["1110411"])
				require.Equal(t, float64(2110412), kv["1110412"])
				require.Equal(t, float64(2110413), kv["1110413"])
			})
			t.Run("f_int32_key_string_val_string", func(t *testing.T) {
				kv := data["f_int32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2110511", kv["1110511"])
				require.Equal(t, "2110512", kv["1110512"])
				require.Equal(t, "2110513", kv["1110513"])
			})
		})
		t.Run("Int64", func(t *testing.T) {
			t.Run("f_int64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_int64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2120111), kv["1120111"])
				require.Equal(t, float64(2120112), kv["1120112"])
				require.Equal(t, float64(2120113), kv["1120113"])
			})
			t.Run("f_int64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_int64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2120411), kv["1120411"])
				require.Equal(t, float64(2120412), kv["1120412"])
				require.Equal(t, float64(2120413), kv["1120413"])
			})
			t.Run("f_int64_key_string_val_string", func(t *testing.T) {
				kv := data["f_int64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2120511", kv["1120511"])
				require.Equal(t, "2120512", kv["1120512"])
				require.Equal(t, "2120513", kv["1120513"])
			})
		})
		t.Run("UInt32", func(t *testing.T) {
			t.Run("f_uint32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_uint32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2130111), kv["1130111"])
				require.Equal(t, float64(2130112), kv["1130112"])
				require.Equal(t, float64(2130113), kv["1130113"])
			})
			t.Run("f_uint32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_uint32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2130411), kv["1130411"])
				require.Equal(t, float64(2130412), kv["1130412"])
				require.Equal(t, float64(2130413), kv["1130413"])
			})
			t.Run("f_uint32_key_string_val_string", func(t *testing.T) {
				kv := data["f_uint32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2130511", kv["1130511"])
				require.Equal(t, "2130512", kv["1130512"])
				require.Equal(t, "2130513", kv["1130513"])
			})
		})
		t.Run("UInt64", func(t *testing.T) {
			t.Run("f_uint64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_uint64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2140111), kv["1140111"])
				require.Equal(t, float64(2140112), kv["1140112"])
				require.Equal(t, float64(2140113), kv["1140113"])
			})
			t.Run("f_uint64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_uint64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2140411), kv["1140411"])
				require.Equal(t, float64(2140412), kv["1140412"])
				require.Equal(t, float64(2140413), kv["1140413"])
			})
			t.Run("f_uint64_key_string_val_string", func(t *testing.T) {
				kv := data["f_uint64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2140511", kv["1140511"])
				require.Equal(t, "2140512", kv["1140512"])
				require.Equal(t, "2140513", kv["1140513"])
			})
		})
		t.Run("SInt32", func(t *testing.T) {
			t.Run("f_sint32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sint32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2150111), kv["1150111"])
				require.Equal(t, float64(2150112), kv["1150112"])
				require.Equal(t, float64(2150113), kv["1150113"])
			})
			t.Run("f_sint32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sint32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2150411), kv["1150411"])
				require.Equal(t, float64(2150412), kv["1150412"])
				require.Equal(t, float64(2150413), kv["1150413"])
			})
			t.Run("f_sint32_key_string_val_string", func(t *testing.T) {
				kv := data["f_sint32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2150511", kv["1150511"])
				require.Equal(t, "2150512", kv["1150512"])
				require.Equal(t, "2150513", kv["1150513"])
			})
		})
		t.Run("SInt64", func(t *testing.T) {
			t.Run("f_sint64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sint64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2160111), kv["1160111"])
				require.Equal(t, float64(2160112), kv["1160112"])
				require.Equal(t, float64(2160113), kv["1160113"])
			})
			t.Run("f_sint64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sint64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2160411), kv["1160411"])
				require.Equal(t, float64(2160412), kv["1160412"])
				require.Equal(t, float64(2160413), kv["1160413"])
			})
			t.Run("f_sint64_key_string_val_string", func(t *testing.T) {
				kv := data["f_sint64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2160511", kv["1160511"])
				require.Equal(t, "2160512", kv["1160512"])
				require.Equal(t, "2160513", kv["1160513"])
			})
		})
		t.Run("SFixed32", func(t *testing.T) {
			t.Run("f_sfixed32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sfixed32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2170111), kv["1170111"])
				require.Equal(t, float64(2170112), kv["1170112"])
				require.Equal(t, float64(2170113), kv["1170113"])
			})
			t.Run("f_sfixed32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sfixed32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2170411), kv["1170411"])
				require.Equal(t, float64(2170412), kv["1170412"])
				require.Equal(t, float64(2170413), kv["1170413"])
			})
			t.Run("f_sfixed32_key_string_val_string", func(t *testing.T) {
				kv := data["f_sfixed32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2170511", kv["1170511"])
				require.Equal(t, "2170512", kv["1170512"])
				require.Equal(t, "2170513", kv["1170513"])
			})
		})
		t.Run("SFixed64", func(t *testing.T) {
			t.Run("f_sfixed64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_sfixed64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2180111), kv["1180111"])
				require.Equal(t, float64(2180112), kv["1180112"])
				require.Equal(t, float64(2180113), kv["1180113"])
			})
			t.Run("f_sfixed64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_sfixed64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2180411), kv["1180411"])
				require.Equal(t, float64(2180412), kv["1180412"])
				require.Equal(t, float64(2180413), kv["1180413"])
			})
			t.Run("f_sfixed64_key_string_val_string", func(t *testing.T) {
				kv := data["f_sfixed64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2180511", kv["1180511"])
				require.Equal(t, "2180512", kv["1180512"])
				require.Equal(t, "2180513", kv["1180513"])
			})
		})
		t.Run("Fixed32", func(t *testing.T) {
			t.Run("f_fixed32_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_fixed32_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2190111), kv["1190111"])
				require.Equal(t, float64(2190112), kv["1190112"])
				require.Equal(t, float64(2190113), kv["1190113"])
			})
			t.Run("f_fixed32_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_fixed32_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2190411), kv["1190411"])
				require.Equal(t, float64(2190412), kv["1190412"])
				require.Equal(t, float64(2190413), kv["1190413"])
			})
			t.Run("f_fixed32_key_string_val_string", func(t *testing.T) {
				kv := data["f_fixed32_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2190511", kv["1190511"])
				require.Equal(t, "2190512", kv["1190512"])
				require.Equal(t, "2190513", kv["1190513"])
			})
		})
		t.Run("Fixed64", func(t *testing.T) {
			t.Run("f_fixed64_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_fixed64_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2210111), kv["1210111"])
				require.Equal(t, float64(2210112), kv["1210112"])
				require.Equal(t, float64(2210113), kv["1210113"])
			})
			t.Run("f_fixed64_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_fixed64_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2210411), kv["1210411"])
				require.Equal(t, float64(2210412), kv["1210412"])
				require.Equal(t, float64(2210413), kv["1210413"])
			})
			t.Run("f_fixed64_key_string_val_string", func(t *testing.T) {
				kv := data["f_fixed64_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2210511", kv["1210511"])
				require.Equal(t, "2210512", kv["1210512"])
				require.Equal(t, "2210513", kv["1210513"])
			})
		})
		t.Run("Float", func(t *testing.T) {
			t.Run("f_float_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_float_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2220.111), kv["1220111"])
				require.Equal(t, float64(2220.112), kv["1220112"])
				require.Equal(t, float64(2220.113), kv["1220113"])
			})
			t.Run("f_float_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_float_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2220.411), kv["1220411"])
				require.Equal(t, float64(2220.412), kv["1220412"])
				require.Equal(t, float64(2220.413), kv["1220413"])
			})
			t.Run("f_float_key_string_val_string", func(t *testing.T) {
				kv := data["f_float_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2220.511", kv["1220511"])
				require.Equal(t, "2220.512", kv["1220512"])
				require.Equal(t, "2220.513", kv["1220513"])
			})
		})
		t.Run("Double", func(t *testing.T) {
			t.Run("f_double_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_double_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(2230111.111), kv["1230111"])
				require.Equal(t, float64(2230112.111), kv["1230112"])
				require.Equal(t, float64(2230113.111), kv["1230113"])
			})
			t.Run("f_double_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_double_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(2230411.111), kv["1230411"])
				require.Equal(t, float64(2230412.111), kv["1230412"])
				require.Equal(t, float64(2230413.111), kv["1230413"])
			})
			t.Run("f_double_key_string_val_string", func(t *testing.T) {
				kv := data["f_double_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "2230511.111", kv["1230511"])
				require.Equal(t, "2230512.111", kv["1230512"])
				require.Equal(t, "2230513.111", kv["1230513"])
			})
		})
		t.Run("Bool", func(t *testing.T) {
			t.Run("f_bool_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_bool_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, true, kv["1240111"])
				require.Equal(t, true, kv["1240112"])
				require.Equal(t, true, kv["1240113"])
			})
			t.Run("f_bool_key_string_val_bool", func(t *testing.T) {
				kv := data["f_bool_key_string_val_bool"].(map[string]interface{})
				require.Equal(t, true, kv["1240411"])
				require.Equal(t, false, kv["1240412"])
				require.Equal(t, true, kv["1240413"])
			})
			t.Run("f_bool_key_string_val_string", func(t *testing.T) {
				kv := data["f_bool_key_string_val_string"].(map[string]interface{})
				require.Equal(t, "false", kv["1240511"])
				require.Equal(t, "true", kv["1240512"])
				require.Equal(t, "false", kv["1240513"])
			})
		})
		t.Run("String", func(t *testing.T) {
			t.Run("f_string_key_unset_val_none", func(t *testing.T) {
				kv := data["f_string_key_unset_val_none"].(map[string]interface{})
				require.Equal(t, "s101", kv["1250111"])
				require.Equal(t, "s102", kv["1250112"])
				require.Equal(t, "s103", kv["1250113"])
			})
			t.Run("f_string_key_string_val_none", func(t *testing.T) {
				kv := data["f_string_key_string_val_none"].(map[string]interface{})
				require.Equal(t, "s301", kv["1250311"])
				require.Equal(t, "s302", kv["1250312"])
				require.Equal(t, "s303", kv["1250313"])
			})
		})
		t.Run("Bytes", func(t *testing.T) {
			t.Run("f_bytes_key_unset_val_none", func(t *testing.T) {
				kv := data["f_bytes_key_unset_val_none"].(map[string]interface{})
				require.Equal(t, "b101", utils.ConvertJSONBytesToString(kv["1260111"].(string)))
				require.Equal(t, "b102", utils.ConvertJSONBytesToString(kv["1260112"].(string)))
				require.Equal(t, "b103", utils.ConvertJSONBytesToString(kv["1260113"].(string)))
			})
			t.Run("f_bytes_key_bytes_val_none", func(t *testing.T) {
				kv := data["f_bytes_key_string_val_none"].(map[string]interface{})
				require.Equal(t, "b301", utils.ConvertJSONBytesToString(kv["1260311"].(string)))
				require.Equal(t, "b302", utils.ConvertJSONBytesToString(kv["1260312"].(string)))
				require.Equal(t, "b303", utils.ConvertJSONBytesToString(kv["1260313"].(string)))
			})
		})
		t.Run("Enum", func(t *testing.T) {
			t.Run("f_enum_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_enum_key_unset_val_unset"].(map[string]interface{})
				require.Equal(t, float64(0), kv["1270111"])
				require.Equal(t, float64(2), kv["1270112"])
				require.Equal(t, float64(3), kv["1270113"])
			})
			t.Run("f_enum_key_string_val_numeric", func(t *testing.T) {
				kv := data["f_enum_key_string_val_numeric"].(map[string]interface{})
				require.Equal(t, float64(6), kv["1270511"])
				require.Equal(t, float64(7), kv["1270512"])
				require.Equal(t, float64(9), kv["1270513"])
			})
			t.Run("f_enum_key_string_val_numeric_string", func(t *testing.T) {
				kv := data["f_enum_key_string_val_numeric_string"].(map[string]interface{})
				require.Equal(t, "7", kv["1270611"])
				require.Equal(t, "9", kv["1270612"])
				require.Equal(t, "11", kv["1270613"])
			})
			t.Run("f_enum_key_string_val_enum_string", func(t *testing.T) {
				kv := data["f_enum_key_string_val_enum_string"].(map[string]interface{})
				require.Equal(t, "Six", kv["1270711"])
				require.Equal(t, "Seven", kv["1270712"])
				require.Equal(t, "Eight", kv["1270713"])
			})
		})
		t.Run("Message", func(t *testing.T) {
			t.Run("f_message_key_unset_val_none", func(t *testing.T) {
				kv := data["f_message_key_unset_val_none"].(map[string]interface{})
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
			t.Run("f_message_key_string_val_none", func(t *testing.T) {
				kv := data["f_message_key_string_val_none"].(map[string]interface{})
				v1 := kv["1280311"].(map[string]interface{})
				require.Equal(t, "m121", v1["f_string1"])
				require.Equal(t, "m122", v1["f_string2"])
				require.Equal(t, "m123", v1["f_string3"])
				v2 := kv["1280312"].(map[string]interface{})
				require.Equal(t, "m221", v2["f_string1"])
				require.Equal(t, "m222", v2["f_string2"])
				require.Equal(t, "m223", v2["f_string3"])
				v3 := kv["1280313"].(map[string]interface{})
				require.Equal(t, "m321", v3["f_string1"])
				require.Equal(t, "m322", v3["f_string2"])
				require.Equal(t, "m323", v3["f_string3"])
			})
		})
		t.Run("Duration", func(t *testing.T) {
			t.Run("f_duration_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_duration_key_unset_val_unset"].(map[string]interface{})
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
			t.Run("f_duration_key_string_val_unset", func(t *testing.T) {
				kv := data["f_duration_key_string_val_unset"].(map[string]interface{})
				v1 := kv["1290271"].(map[string]interface{})
				require.Equal(t, float64(110271), v1["seconds"])
				require.Equal(t, float64(271), v1["nanos"])
				v2 := kv["1290272"].(map[string]interface{})
				require.Equal(t, float64(110272), v2["seconds"])
				require.Equal(t, float64(272), v2["nanos"])
				v3 := kv["1290273"].(map[string]interface{})
				require.Equal(t, float64(110273), v3["seconds"])
				require.Equal(t, float64(273), v3["nanos"])
			})
			t.Run("f_duration_key_string_val_object", func(t *testing.T) {
				kv := data["f_duration_key_string_val_object"].(map[string]interface{})
				v1 := kv["1290281"].(map[string]interface{})
				require.Equal(t, float64(110281), v1["seconds"])
				require.Equal(t, float64(281), v1["nanos"])
				v2 := kv["1290282"].(map[string]interface{})
				require.Equal(t, float64(110282), v2["seconds"])
				require.Equal(t, float64(282), v2["nanos"])
				v3 := kv["1290283"].(map[string]interface{})
				require.Equal(t, float64(110283), v3["seconds"])
				require.Equal(t, float64(283), v3["nanos"])
			})
			t.Run("f_duration_key_string_val_time_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_time_string"].(map[string]interface{})
				require.Equal(t, "30h38m11.000000291s", kv["1290291"])
				require.Equal(t, "30h38m12.000000292s", kv["1290292"])
				require.Equal(t, "30h38m13.000000293s", kv["1290293"])
			})
			t.Run("f_duration_key_string_val_nanosecond", func(t *testing.T) {
				kv := data["f_duration_key_string_val_nanosecond"].(map[string]interface{})
				require.Equal(t, float64(110311000000311), kv["1290311"])
				require.Equal(t, float64(110312000000312), kv["1290312"])
				require.Equal(t, float64(110313000000313), kv["1290313"])
			})
			t.Run("f_duration_key_string_val_nanosecond_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_nanosecond_string"].(map[string]interface{})
				require.Equal(t, "110321000000321", kv["1290321"])
				require.Equal(t, "110322000000322", kv["1290322"])
				require.Equal(t, "110323000000323", kv["1290323"])
			})
			t.Run("f_duration_key_string_val_microsecond", func(t *testing.T) {
				kv := data["f_duration_key_string_val_microsecond"].(map[string]interface{})
				require.Equal(t, float64(110331000331), kv["1290331"])
				require.Equal(t, float64(110332000332), kv["1290332"])
				require.Equal(t, float64(110333000333), kv["1290333"])
			})
			t.Run("f_duration_key_string_val_microsecond_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_microsecond_string"].(map[string]interface{})
				require.Equal(t, "110341000341", kv["1290341"])
				require.Equal(t, "110342000342", kv["1290342"])
				require.Equal(t, "110343000343", kv["1290343"])
			})
			t.Run("f_duration_key_string_val_millisecond", func(t *testing.T) {
				kv := data["f_duration_key_string_val_millisecond"].(map[string]interface{})
				require.Equal(t, float64(110351351), kv["1290351"])
				require.Equal(t, float64(110352352), kv["1290352"])
				require.Equal(t, float64(110353353), kv["1290353"])
			})
			t.Run("f_duration_key_string_val_millisecond_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_millisecond_string"].(map[string]interface{})
				require.Equal(t, "110361361", kv["1290361"])
				require.Equal(t, "110362362", kv["1290362"])
				require.Equal(t, "110363363", kv["1290363"])
			})
			t.Run("f_duration_key_string_val_second", func(t *testing.T) {
				kv := data["f_duration_key_string_val_second"].(map[string]interface{})
				require.Equal(t, float64(110371), kv["1290371"])
				require.Equal(t, float64(110372), kv["1290372"])
				require.Equal(t, float64(110373), kv["1290373"])
			})
			t.Run("f_duration_key_string_val_second_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_second_string"].(map[string]interface{})
				require.Equal(t, "110381", kv["1290381"])
				require.Equal(t, "110382", kv["1290382"])
				require.Equal(t, "110383", kv["1290383"])
			})
			t.Run("f_duration_key_string_val_minute", func(t *testing.T) {
				kv := data["f_duration_key_string_val_minute"].(map[string]interface{})
				require.Equal(t, float64(1839.85), kv["1290391"])
				require.Equal(t, float64(1839.8666666666666), kv["1290392"])
				require.Equal(t, float64(1839.8833333333334), kv["1290393"])
			})
			t.Run("f_duration_key_string_val_minute_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_minute_string"].(map[string]interface{})
				require.Equal(t, "1840.1833333333334", kv["1290411"])
				require.Equal(t, "1840.2", kv["1290412"])
				require.Equal(t, "1840.2166666666667", kv["1290413"])
			})
			t.Run("f_duration_key_string_val_hour", func(t *testing.T) {
				kv := data["f_duration_key_string_val_hour"].(map[string]interface{})
				require.Equal(t, float64(30.6725), kv["1290421"])
				require.Equal(t, float64(30.672777777777778), kv["1290422"])
				require.Equal(t, float64(30.673055555555557), kv["1290423"])
			})
			t.Run("f_duration_key_string_val_hour_string", func(t *testing.T) {
				kv := data["f_duration_key_string_val_hour_string"].(map[string]interface{})
				require.Equal(t, "30.67527777777778", kv["1290431"])
				require.Equal(t, "30.675555555555555", kv["1290432"])
				require.Equal(t, "30.675833333333333", kv["1290433"])
			})
		})
		t.Run("Timestamp", func(t *testing.T) {
			t.Run("f_timestamp_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_timestamp_key_unset_val_unset"].(map[string]interface{})
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
			t.Run("f_timestamp_key_string_val_object", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_object"].(map[string]interface{})
				v1 := kv["1310231"].(map[string]interface{})
				require.Equal(t, float64(1718334231), v1["seconds"])
				require.Equal(t, float64(231), v1["nanos"])
				v2 := kv["1310232"].(map[string]interface{})
				require.Equal(t, float64(1718334232), v2["seconds"])
				require.Equal(t, float64(232), v2["nanos"])
				v3 := kv["1310233"].(map[string]interface{})
				require.Equal(t, float64(1718334233), v3["seconds"])
				require.Equal(t, float64(233), v3["nanos"])
			})
			t.Run("f_timestamp_key_string_val_time_layout", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_time_layout"].(map[string]interface{})
				require.Equal(t, "2024-06-14T03:04:01.000000241Z", kv["1310241"])
				require.Equal(t, "2024-06-14T03:04:02.000000242Z", kv["1310242"])
				require.Equal(t, "2024-06-14T03:04:03.000000243Z", kv["1310243"])
			})
			t.Run("f_timestamp_key_string_val_unix_nano", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_nano"].(map[string]interface{})
				require.Equal(t, float64(1718334251000000251), kv["1310251"])
				require.Equal(t, float64(1718334252000000252), kv["1310252"])
				require.Equal(t, float64(1718334253000000253), kv["1310253"])
			})
			t.Run("f_timestamp_key_string_val_unix_nano_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_nano_string"].(map[string]interface{})
				require.Equal(t, "1718334261000000261", kv["1310261"])
				require.Equal(t, "1718334262000000262", kv["1310262"])
				require.Equal(t, "1718334263000000263", kv["1310263"])
			})
			t.Run("f_timestamp_key_string_val_unix_micro", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_micro"].(map[string]interface{})
				require.Equal(t, float64(1718334271000271), kv["1310271"])
				require.Equal(t, float64(1718334272000272), kv["1310272"])
				require.Equal(t, float64(1718334273000273), kv["1310273"])
			})
			t.Run("f_timestamp_key_string_val_unix_micro_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_micro_string"].(map[string]interface{})
				require.Equal(t, "1718334281000281", kv["1310281"])
				require.Equal(t, "1718334282000282", kv["1310282"])
				require.Equal(t, "1718334283000283", kv["1310283"])
			})
			t.Run("f_timestamp_key_string_val_unix_milli", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_milli"].(map[string]interface{})
				require.Equal(t, float64(1718334311311), kv["1310311"])
				require.Equal(t, float64(1718334312312), kv["1310312"])
				require.Equal(t, float64(1718334313313), kv["1310313"])
			})
			t.Run("f_timestamp_key_string_val_unix_milli_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_milli_string"].(map[string]interface{})
				require.Equal(t, "1718334321321", kv["1310321"])
				require.Equal(t, "1718334322322", kv["1310322"])
				require.Equal(t, "1718334323323", kv["1310323"])
			})
			t.Run("f_timestamp_key_string_val_unix_sec", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_sec"].(map[string]interface{})
				require.Equal(t, float64(1718334331), kv["1310331"])
				require.Equal(t, float64(1718334332), kv["1310332"])
				require.Equal(t, float64(1718334333), kv["1310333"])
			})
			t.Run("f_timestamp_key_string_val_unix_sec_string", func(t *testing.T) {
				kv := data["f_timestamp_key_string_val_unix_sec_string"].(map[string]interface{})
				require.Equal(t, "1718334341", kv["1310341"])
				require.Equal(t, "1718334342", kv["1310342"])
				require.Equal(t, "1718334343", kv["1310343"])
			})
		})
		t.Run("Any", func(t *testing.T) {
			t.Run("f_any_key_unset_val_unset", func(t *testing.T) {
				kv := data["f_any_key_unset_val_unset"].(map[string]interface{})
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
			t.Run("f_any_key_string_val_object", func(t *testing.T) {
				kv := data["f_any_key_string_val_object"].(map[string]interface{})
				v1 := kv["1320411"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["type_url"])
				require.Equal(t, "CgRhNDExEgRhNDEyGgRhNDEz", v1["value"])
				v2 := kv["1320421"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["type_url"])
				require.Equal(t, "CgRhNDIxEgRhNDIyGgRhNDIz", v2["value"])
				v3 := kv["1320431"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v3["type_url"])
				require.Equal(t, "CgRhNDMxEgRhNDMyGgRhNDMz", v3["value"])
			})
			t.Run("f_any_key_string_val_proto", func(t *testing.T) {
				kv := data["f_any_key_string_val_proto"].(map[string]interface{})
				v1 := kv["1320511"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v1["@type"])
				require.Equal(t, "a511", v1["f_string1"])
				require.Equal(t, "a512", v1["f_string2"])
				require.Equal(t, "a513", v1["f_string3"])
				v2 := kv["1320521"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v2["@type"])
				require.Equal(t, "a521", v2["f_string1"])
				require.Equal(t, "a522", v2["f_string2"])
				require.Equal(t, "a523", v2["f_string3"])
				v3 := kv["1320531"].(map[string]interface{})
				require.Equal(t, "type.googleapis.com/external.Message1", v3["@type"])
				require.Equal(t, "a531", v3["f_string1"])
				require.Equal(t, "a532", v3["f_string2"])
				require.Equal(t, "a533", v3["f_string3"])
			})
		})
	})
}
