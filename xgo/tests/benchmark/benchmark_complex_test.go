package benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbbenchmark"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

type CopyComplex pbbenchmark.BenchModelComplex

var seedComplex = &pbbenchmark.BenchModelComplex{
	FString1:  "s11111",
	FString2:  "s11112",
	FString3:  "s11113",
	FString4:  "s11114",
	FString5:  "s11115",
	FString6:  "s11116",
	FString7:  "s11117",
	FString8:  "s11118",
	FString9:  "s11119",
	FInt32:    11111,
	FInt64:    11112,
	FUint32:   11113,
	FUint64:   11114,
	FSint32:   11115,
	FSint64:   11116,
	FSfixed32: 11117,
	FSfixed64: 11118,
	FFixed32:  11119,
	FFixed64:  11120,
	FFloat:    111.111,
	FDouble:   112.112,
	FBool1:    true,
	FBool2:    true,
	FBool3:    true,
	FBytes1:   []byte("b11111"),
	FBytes2:   []byte("b11112"),
	FBytes3:   []byte("b11113"),
	FEnum1:    pbbenchmark.Enum1(2),
	FEnum2:    pbbenchmark.Enum1(3),
	FEnum3:    pbbenchmark.Enum1(5),
	FEnum4:    pbbenchmark.Enum1(7),
	FEnum5:    pbbenchmark.Enum1(9),
	FEnum6:    pbbenchmark.Enum1(11),
	FEnum7:    pbbenchmark.Enum1(12),
	FMessage1: &pbbenchmark.Message1{FString1: "ms1111", FString2: "ms1112", FString3: "ms1113"},
	FMessage2: &pbbenchmark.Message1{FString1: "ms1121", FString2: "ms1122", FString3: "ms1123"},
	FMessage3: &pbbenchmark.Message1{FString1: "ms1131", FString2: "ms1132", FString3: "ms1133"},
	FMessage4: &pbbenchmark.Message1{FString1: "ms1141", FString2: "ms1142", FString3: "ms1143"},
	FMessage5: &pbbenchmark.Message1{FString1: "ms1151", FString2: "ms1152", FString3: "ms1153"},
	FMessage6: &pbbenchmark.Message1{FString1: "ms1161", FString2: "ms1162", FString3: "ms1163"},
	FMessage7: &pbbenchmark.Message1{FString1: "ms1171", FString2: "ms1172", FString3: "ms1173"},
	PString1:  utils.PointerString("s11211"),
	PString2:  utils.PointerString("s11212"),
	PString3:  utils.PointerString("s11213"),
	PString4:  utils.PointerString("s11214"),
	PString5:  utils.PointerString("s11215"),
	PString6:  utils.PointerString("s11216"),
	PString7:  utils.PointerString("s11217"),
	PString8:  utils.PointerString("s11218"),
	PString9:  utils.PointerString("s11219"),
	PInt32:    utils.PointerInt32(11211),
	PInt64:    utils.PointerInt64(11212),
	PUint32:   utils.PointerUint32(11213),
	PUint64:   utils.PointerUint64(11214),
	PSint32:   utils.PointerInt32(11215),
	PSint64:   utils.PointerInt64(11216),
	PSfixed32: utils.PointerInt32(11217),
	PSfixed64: utils.PointerInt64(11218),
	PFixed32:  utils.PointerUint32(11219),
	PFixed64:  utils.PointerUint64(11220),
	PFloat:    utils.PointerFloat32(211.211),
	PDouble:   utils.PointerFloat64(212.212),
	PBool1:    utils.PointerBool(true),
	PBool2:    utils.PointerBool(true),
	PBool3:    utils.PointerBool(true),
	PBytes1:   []byte("b11211"),
	PBytes2:   []byte("b11212"),
	PBytes3:   []byte("b11213"),
	PEnum1:    pbbenchmark.Enum1(2).Enum(),
	PEnum2:    pbbenchmark.Enum1(3).Enum(),
	PEnum3:    pbbenchmark.Enum1(5).Enum(),
	PEnum4:    pbbenchmark.Enum1(7).Enum(),
	PEnum5:    pbbenchmark.Enum1(9).Enum(),
	PEnum6:    pbbenchmark.Enum1(11).Enum(),
	PEnum7:    pbbenchmark.Enum1(12).Enum(),
	PMessage1: &pbbenchmark.Message1{FString1: "ms1211", FString2: "ms1212", FString3: "ms1213"},
	PMessage2: &pbbenchmark.Message1{FString1: "ms1221", FString2: "ms1222", FString3: "ms1223"},
	PMessage3: &pbbenchmark.Message1{FString1: "ms1231", FString2: "ms1232", FString3: "ms1233"},
	PMessage4: &pbbenchmark.Message1{FString1: "ms1241", FString2: "ms1242", FString3: "ms1243"},
	PMessage5: &pbbenchmark.Message1{FString1: "ms1251", FString2: "ms1252", FString3: "ms1253"},
	PMessage6: &pbbenchmark.Message1{FString1: "ms1261", FString2: "ms1262", FString3: "ms1263"},
	PMessage7: &pbbenchmark.Message1{FString1: "ms1271", FString2: "ms1272", FString3: "ms1273"},
	RString1:  []string{"s11311", "s11312", "s11313", "s11314", "s11315", "s11316", "s11317"},
	RString2:  []string{"s11321", "s11322", "s11323", "s11324", "s11325", "s11326", "s11327"},
	RString3:  []string{"s11331", "s11332", "s11333", "s11334", "s11335", "s11336", "s11337"},
	RString4:  []string{"s11341", "s11342", "s11343", "s11344", "s11345", "s11346", "s11347"},
	RString5:  []string{"s11351", "s11352", "s11353", "s11354", "s11355", "s11356", "s11357"},
	RString6:  []string{"s11361", "s11362", "s11363", "s11364", "s11365", "s11366", "s11367"},
	RString7:  []string{"s11371", "s11372", "s11373", "s11374", "s11375", "s11376", "s11377"},
	RString8:  []string{"s11381", "s11382", "s11383", "s11384", "s11385", "s11386", "s11387"},
	RString9:  []string{"s11391", "s11392", "s11393", "s11394", "s11395", "s11396", "s11397"},
	RInt32:    []int32{11311, 11312, 11313, 11314, 11315, 11316, 11317},
	RInt64:    []int64{11321, 11322, 11323, 11324, 11325, 11326, 11327},
	RUint32:   []uint32{11331, 11332, 11333, 11334, 11335, 11336, 11337},
	RUint64:   []uint64{11341, 11342, 11343, 11344, 11345, 11346, 11347},
	RSint32:   []int32{11351, 11352, 11353, 11354, 11355, 11356, 11357},
	RSint64:   []int64{11361, 11362, 11363, 11364, 11365, 11366, 11367},
	RSfixed32: []int32{11371, 11372, 11373, 11374, 11375, 11376, 11377},
	RSfixed64: []int64{11381, 11382, 11383, 11384, 11385, 11386, 11387},
	RFixed32:  []uint32{11391, 11392, 11393, 11394, 11395, 11396, 11397},
	RFixed64:  []uint64{11401, 11402, 11403, 11404, 11405, 11406, 11407},
	RFloat:    []float32{311.311, 311.312, 311.313, 311.314, 311.315, 311.316, 311.317},
	RDouble:   []float64{411.411, 411.412, 411.413, 411.414, 411.415, 411.416, 411.417},
	RBool1:    []bool{true, true, true, false, true, true, true},
	RBool2:    []bool{true, true, true, false, true, true, true},
	RBool3:    []bool{true, true, true, false, true, true, true},
	RBytes1:   [][]byte{[]byte("b11311"), []byte("b11312"), []byte("b11313"), []byte("b11314"), []byte("b11315")},
	RBytes2:   [][]byte{[]byte("b11411"), []byte("b11412"), []byte("b11413"), []byte("b11414"), []byte("b11415")},
	RBytes3:   [][]byte{[]byte("b11511"), []byte("b11512"), []byte("b11513"), []byte("b11514"), []byte("b11515")},
	REnum1:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum2:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum3:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum4:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum5:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum6:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum7:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	RMessage1: []*pbbenchmark.Message1{
		{FString1: "ms1311", FString2: "ms1312", FString3: "ms1313"},
		{FString1: "ms1411", FString2: "ms1412", FString3: "ms1413"},
		{FString1: "ms1511", FString2: "ms1412", FString3: "ms1513"},
	},
	RMessage2: []*pbbenchmark.Message1{
		{FString1: "ms1321", FString2: "ms1322", FString3: "ms1323"},
		{FString1: "ms1421", FString2: "ms1422", FString3: "ms1423"},
		{FString1: "ms1521", FString2: "ms1422", FString3: "ms1523"},
	},
	RMessage3: []*pbbenchmark.Message1{
		{FString1: "ms1331", FString2: "ms1332", FString3: "ms1333"},
		{FString1: "ms1431", FString2: "ms1432", FString3: "ms1433"},
		{FString1: "ms1531", FString2: "ms1432", FString3: "ms1533"},
	},
	RMessage4: []*pbbenchmark.Message1{
		{FString1: "ms1341", FString2: "ms1342", FString3: "ms1343"},
		{FString1: "ms1441", FString2: "ms1442", FString3: "ms1443"},
		{FString1: "ms1541", FString2: "ms1442", FString3: "ms1543"},
	},
	RMessage5: []*pbbenchmark.Message1{
		{FString1: "ms1351", FString2: "ms1352", FString3: "ms1353"},
		{FString1: "ms1451", FString2: "ms1452", FString3: "ms1453"},
		{FString1: "ms1551", FString2: "ms1452", FString3: "ms1553"},
	},
	RMessage6: []*pbbenchmark.Message1{
		{FString1: "ms1361", FString2: "ms1362", FString3: "ms1363"},
		{FString1: "ms1461", FString2: "ms1462", FString3: "ms1463"},
		{FString1: "ms1561", FString2: "ms1462", FString3: "ms1563"},
	},
	RMessage7: []*pbbenchmark.Message1{
		{FString1: "ms1371", FString2: "ms1372", FString3: "ms1373"},
		{FString1: "ms1471", FString2: "ms1472", FString3: "ms1473"},
		{FString1: "ms1571", FString2: "ms1472", FString3: "ms1573"},
	},
	MString1: map[string]string{
		"mk111": "mv111",
		"mk112": "mv112",
		"mk113": "mv113",
		"mk114": "mv114",
		"mk115": "mv115",
		"mk116": "mv116",
		"mk117": "mv117",
	},
	MString2: map[string]string{
		"mk121": "mv121",
		"mk122": "mv122",
		"mk123": "mv123",
		"mk124": "mv124",
		"mk125": "mv125",
		"mk126": "mv126",
		"mk127": "mv127",
	},
	MString3: map[string]string{
		"mk131": "mv131",
		"mk132": "mv132",
		"mk133": "mv133",
		"mk134": "mv134",
		"mk135": "mv135",
		"mk136": "mv136",
		"mk137": "mv137",
	},
	MString4: map[string]string{
		"mk141": "mv141",
		"mk142": "mv142",
		"mk143": "mv143",
		"mk144": "mv144",
		"mk145": "mv145",
		"mk146": "mv146",
		"mk147": "mv147",
	},
	MString5: map[string]string{
		"mk151": "mv151",
		"mk152": "mv152",
		"mk153": "mv153",
		"mk154": "mv154",
		"mk155": "mv155",
		"mk156": "mv156",
		"mk157": "mv157",
	},
	MString6: map[string]string{
		"mk161": "mv161",
		"mk162": "mv162",
		"mk163": "mv163",
		"mk164": "mv164",
		"mk165": "mv165",
		"mk166": "mv166",
		"mk167": "mv167",
	},
	MString7: map[string]string{
		"mk171": "mv171",
		"mk172": "mv172",
		"mk173": "mv173",
		"mk174": "mv174",
		"mk175": "mv175",
		"mk176": "mv176",
		"mk177": "mv177",
	},
	MString8: map[string]string{
		"mk181": "mv181",
		"mk182": "mv182",
		"mk183": "mv183",
		"mk184": "mv184",
		"mk185": "mv185",
		"mk186": "mv186",
		"mk187": "mv187",
	},
	MString9: map[string]string{
		"mk191": "mv191",
		"mk192": "mv192",
		"mk193": "mv193",
		"mk194": "mv194",
		"mk195": "mv195",
		"mk196": "mv196",
		"mk197": "mv197",
	},
	MInt32: map[string]int32{
		"mi111": 111,
		"mi112": 112,
		"mi113": 113,
		"mi114": 114,
		"mi115": 115,
		"mi116": 116,
		"mi117": 117,
	},
	MInt64: map[string]int64{
		"mi121": 121,
		"mi122": 122,
		"mi123": 123,
		"mi124": 124,
		"mi125": 125,
		"mi126": 126,
		"mi127": 127,
	},
	MUint32: map[string]uint32{
		"mi131": 131,
		"mi132": 132,
		"mi133": 133,
		"mi134": 134,
		"mi135": 135,
		"mi136": 136,
		"mi137": 137,
	},
	MUint64: map[string]uint64{
		"mi141": 141,
		"mi142": 142,
		"mi143": 143,
		"mi144": 144,
		"mi145": 145,
		"mi146": 146,
		"mi147": 147,
	},
	MSint32: map[string]int32{
		"mi151": 151,
		"mi152": 152,
		"mi153": 153,
		"mi154": 154,
		"mi155": 155,
		"mi156": 156,
		"mi157": 157,
	},
	MSint64: map[string]int64{
		"mi161": 161,
		"mi162": 162,
		"mi163": 163,
		"mi164": 164,
		"mi165": 165,
		"mi166": 166,
		"mi167": 167,
	},
	MSfixed32: map[string]int32{
		"mi171": 171,
		"mi172": 172,
		"mi173": 173,
		"mi174": 174,
		"mi175": 175,
		"mi176": 176,
		"mi177": 177,
	},
	MSfixed64: map[string]int64{
		"mi181": 181,
		"mi182": 182,
		"mi183": 183,
		"mi184": 184,
		"mi185": 185,
		"mi186": 186,
		"mi187": 187,
	},
	MFixed32: map[string]uint32{
		"mi191": 191,
		"mi192": 192,
		"mi193": 193,
		"mi194": 194,
		"mi195": 195,
		"mi196": 196,
		"mi197": 197,
	},
	MFixed64: map[string]uint64{
		"mi211": 211,
		"mi212": 212,
		"mi213": 213,
		"mi214": 214,
		"mi215": 215,
		"mi216": 216,
		"mi217": 217,
	},
	MFloat: map[string]float32{
		"mf111.111": 111.111,
		"mf112.112": 112.112,
		"mf113.113": 113.113,
		"mf114.114": 114.114,
		"mf115.115": 115.115,
		"mf116.116": 116.116,
		"mf117.117": 117.117,
	},
	MDouble: map[string]float64{
		"mf211.211": 211.211,
		"mf212.212": 212.212,
		"mf213.213": 213.213,
		"mf214.214": 214.214,
		"mf215.215": 215.215,
		"mf216.216": 216.216,
		"mf217.217": 217.217,
	},
	MBool1: map[string]bool{
		"mb111": true,
		"mb112": true,
		"mb113": true,
		"mb114": true,
		"mb115": true,
		"mb116": true,
		"mb117": true,
	},
	MBool2: map[string]bool{
		"mb211": true,
		"mb212": true,
		"mb213": true,
		"mb214": true,
		"mb215": true,
		"mb216": true,
		"mb217": true,
	},
	MBool3: map[string]bool{
		"mb311": true,
		"mb312": true,
		"mb313": true,
		"mb314": true,
		"mb315": true,
		"mb316": true,
		"mb317": true,
	},
	MBytes1: map[string][]byte{
		"mb111": []byte("bv111"),
		"mb112": []byte("bv112"),
		"mb113": []byte("bv113"),
		"mb114": []byte("bv114"),
		"mb115": []byte("bv115"),
		"mb116": []byte("bv116"),
		"mb117": []byte("bv117"),
	},
	MBytes2: map[string][]byte{
		"mb211": []byte("bv211"),
		"mb212": []byte("bv212"),
		"mb213": []byte("bv213"),
		"mb214": []byte("bv214"),
		"mb215": []byte("bv215"),
		"mb216": []byte("bv216"),
		"mb217": []byte("bv217"),
	},
	MBytes3: map[string][]byte{
		"mb311": []byte("bv311"),
		"mb312": []byte("bv312"),
		"mb313": []byte("bv313"),
		"mb314": []byte("bv314"),
		"mb315": []byte("bv315"),
		"mb316": []byte("bv316"),
		"mb317": []byte("bv317"),
	},
	MEnum1: map[string]pbbenchmark.Enum1{
		"me111": 1,
		"me112": 2,
		"me113": 3,
		"me114": 4,
		"me115": 5,
		"me116": 6,
		"me117": 7,
	},
	MEnum2: map[string]pbbenchmark.Enum1{
		"me121": 1,
		"me122": 2,
		"me123": 3,
		"me124": 4,
		"me125": 5,
		"me126": 6,
		"me127": 7,
	},
	MEnum3: map[string]pbbenchmark.Enum1{
		"me131": 1,
		"me132": 2,
		"me133": 3,
		"me134": 4,
		"me135": 5,
		"me136": 6,
		"me137": 7,
	},
	MEnum4: map[string]pbbenchmark.Enum1{
		"me141": 1,
		"me142": 2,
		"me143": 3,
		"me144": 4,
		"me145": 5,
		"me146": 6,
		"me147": 7,
	},
	MEnum5: map[string]pbbenchmark.Enum1{
		"me151": 1,
		"me152": 2,
		"me153": 3,
		"me154": 4,
		"me155": 5,
		"me156": 6,
		"me157": 7,
	},
	MEnum6: map[string]pbbenchmark.Enum1{
		"me161": 1,
		"me162": 2,
		"me163": 3,
		"me164": 4,
		"me165": 5,
		"me166": 6,
		"me167": 7,
	},
	MEnum7: map[string]pbbenchmark.Enum1{
		"me171": 1,
		"me172": 2,
		"me173": 3,
		"me174": 4,
		"me175": 5,
		"me176": 6,
		"me177": 7,
	},
	MMessage1: map[string]*pbbenchmark.Message1{
		"mmk111": {FString1: "mmv111", FString2: "mmv112", FString3: "mmv113"},
		"mmk112": {FString1: "mmv121", FString2: "mmv122", FString3: "mmv123"},
		"mmk113": {FString1: "mmv131", FString2: "mmv132", FString3: "mmv133"},
		"mmk114": {FString1: "mmv141", FString2: "mmv142", FString3: "mmv143"},
		"mmk115": {FString1: "mmv151", FString2: "mmv152", FString3: "mmv153"},
	},
	MMessage2: map[string]*pbbenchmark.Message1{
		"mmk211": {FString1: "mmv211", FString2: "mmv212", FString3: "mmv213"},
		"mmk212": {FString1: "mmv221", FString2: "mmv222", FString3: "mmv223"},
		"mmk213": {FString1: "mmv231", FString2: "mmv232", FString3: "mmv233"},
		"mmk214": {FString1: "mmv241", FString2: "mmv242", FString3: "mmv243"},
		"mmk215": {FString1: "mmv251", FString2: "mmv252", FString3: "mmv253"},
	},
	MMessage3: map[string]*pbbenchmark.Message1{
		"mmk311": {FString1: "mmv311", FString2: "mmv312", FString3: "mmv313"},
		"mmk312": {FString1: "mmv321", FString2: "mmv322", FString3: "mmv323"},
		"mmk313": {FString1: "mmv331", FString2: "mmv332", FString3: "mmv333"},
		"mmk314": {FString1: "mmv341", FString2: "mmv342", FString3: "mmv343"},
		"mmk315": {FString1: "mmv351", FString2: "mmv352", FString3: "mmv353"},
	},
	MMessage4: map[string]*pbbenchmark.Message1{
		"mmk411": {FString1: "mmv411", FString2: "mmv412", FString3: "mmv413"},
		"mmk412": {FString1: "mmv421", FString2: "mmv422", FString3: "mmv423"},
		"mmk413": {FString1: "mmv431", FString2: "mmv432", FString3: "mmv433"},
		"mmk414": {FString1: "mmv441", FString2: "mmv442", FString3: "mmv443"},
		"mmk415": {FString1: "mmv451", FString2: "mmv452", FString3: "mmv453"},
	},
	MMessage5: map[string]*pbbenchmark.Message1{
		"mmk511": {FString1: "mmv511", FString2: "mmv512", FString3: "mmv513"},
		"mmk512": {FString1: "mmv521", FString2: "mmv522", FString3: "mmv523"},
		"mmk513": {FString1: "mmv531", FString2: "mmv532", FString3: "mmv533"},
		"mmk514": {FString1: "mmv541", FString2: "mmv542", FString3: "mmv543"},
		"mmk515": {FString1: "mmv551", FString2: "mmv552", FString3: "mmv553"},
	},
	MMessage6: map[string]*pbbenchmark.Message1{
		"mmk611": {FString1: "mmv611", FString2: "mmv612", FString3: "mmv613"},
		"mmk612": {FString1: "mmv621", FString2: "mmv622", FString3: "mmv623"},
		"mmk613": {FString1: "mmv631", FString2: "mmv632", FString3: "mmv633"},
		"mmk614": {FString1: "mmv641", FString2: "mmv642", FString3: "mmv643"},
		"mmk615": {FString1: "mmv651", FString2: "mmv652", FString3: "mmv653"},
	},
	MMessage7: map[string]*pbbenchmark.Message1{
		"mmk711": {FString1: "mmv711", FString2: "mmv712", FString3: "mmv713"},
		"mmk712": {FString1: "mmv721", FString2: "mmv722", FString3: "mmv723"},
		"mmk713": {FString1: "mmv731", FString2: "mmv732", FString3: "mmv733"},
		"mmk714": {FString1: "mmv741", FString2: "mmv742", FString3: "mmv743"},
		"mmk715": {FString1: "mmv751", FString2: "mmv752", FString3: "mmv753"},
	},
}

var copyComplex = &CopyComplex{
	FString1:  "s11111",
	FString2:  "s11112",
	FString3:  "s11113",
	FString4:  "s11114",
	FString5:  "s11115",
	FString6:  "s11116",
	FString7:  "s11117",
	FString8:  "s11118",
	FString9:  "s11119",
	FInt32:    11111,
	FInt64:    11112,
	FUint32:   11113,
	FUint64:   11114,
	FSint32:   11115,
	FSint64:   11116,
	FSfixed32: 11117,
	FSfixed64: 11118,
	FFixed32:  11119,
	FFixed64:  11120,
	FFloat:    111.111,
	FDouble:   112.112,
	FBool1:    true,
	FBool2:    true,
	FBool3:    true,
	FBytes1:   []byte("b11111"),
	FBytes2:   []byte("b11112"),
	FBytes3:   []byte("b11113"),
	FEnum1:    pbbenchmark.Enum1(2),
	FEnum2:    pbbenchmark.Enum1(3),
	FEnum3:    pbbenchmark.Enum1(5),
	FEnum4:    pbbenchmark.Enum1(7),
	FEnum5:    pbbenchmark.Enum1(9),
	FEnum6:    pbbenchmark.Enum1(11),
	FEnum7:    pbbenchmark.Enum1(12),
	FMessage1: &pbbenchmark.Message1{FString1: "ms1111", FString2: "ms1112", FString3: "ms1113"},
	FMessage2: &pbbenchmark.Message1{FString1: "ms1121", FString2: "ms1122", FString3: "ms1123"},
	FMessage3: &pbbenchmark.Message1{FString1: "ms1131", FString2: "ms1132", FString3: "ms1133"},
	FMessage4: &pbbenchmark.Message1{FString1: "ms1141", FString2: "ms1142", FString3: "ms1143"},
	FMessage5: &pbbenchmark.Message1{FString1: "ms1151", FString2: "ms1152", FString3: "ms1153"},
	FMessage6: &pbbenchmark.Message1{FString1: "ms1161", FString2: "ms1162", FString3: "ms1163"},
	FMessage7: &pbbenchmark.Message1{FString1: "ms1171", FString2: "ms1172", FString3: "ms1173"},
	PString1:  utils.PointerString("s11211"),
	PString2:  utils.PointerString("s11212"),
	PString3:  utils.PointerString("s11213"),
	PString4:  utils.PointerString("s11214"),
	PString5:  utils.PointerString("s11215"),
	PString6:  utils.PointerString("s11216"),
	PString7:  utils.PointerString("s11217"),
	PString8:  utils.PointerString("s11218"),
	PString9:  utils.PointerString("s11219"),
	PInt32:    utils.PointerInt32(11211),
	PInt64:    utils.PointerInt64(11212),
	PUint32:   utils.PointerUint32(11213),
	PUint64:   utils.PointerUint64(11214),
	PSint32:   utils.PointerInt32(11215),
	PSint64:   utils.PointerInt64(11216),
	PSfixed32: utils.PointerInt32(11217),
	PSfixed64: utils.PointerInt64(11218),
	PFixed32:  utils.PointerUint32(11219),
	PFixed64:  utils.PointerUint64(11220),
	PFloat:    utils.PointerFloat32(211.211),
	PDouble:   utils.PointerFloat64(212.212),
	PBool1:    utils.PointerBool(true),
	PBool2:    utils.PointerBool(true),
	PBool3:    utils.PointerBool(true),
	PBytes1:   []byte("b11211"),
	PBytes2:   []byte("b11212"),
	PBytes3:   []byte("b11213"),
	PEnum1:    pbbenchmark.Enum1(2).Enum(),
	PEnum2:    pbbenchmark.Enum1(3).Enum(),
	PEnum3:    pbbenchmark.Enum1(5).Enum(),
	PEnum4:    pbbenchmark.Enum1(7).Enum(),
	PEnum5:    pbbenchmark.Enum1(9).Enum(),
	PEnum6:    pbbenchmark.Enum1(11).Enum(),
	PEnum7:    pbbenchmark.Enum1(12).Enum(),
	PMessage1: &pbbenchmark.Message1{FString1: "ms1211", FString2: "ms1212", FString3: "ms1213"},
	PMessage2: &pbbenchmark.Message1{FString1: "ms1221", FString2: "ms1222", FString3: "ms1223"},
	PMessage3: &pbbenchmark.Message1{FString1: "ms1231", FString2: "ms1232", FString3: "ms1233"},
	PMessage4: &pbbenchmark.Message1{FString1: "ms1241", FString2: "ms1242", FString3: "ms1243"},
	PMessage5: &pbbenchmark.Message1{FString1: "ms1251", FString2: "ms1252", FString3: "ms1253"},
	PMessage6: &pbbenchmark.Message1{FString1: "ms1261", FString2: "ms1262", FString3: "ms1263"},
	PMessage7: &pbbenchmark.Message1{FString1: "ms1271", FString2: "ms1272", FString3: "ms1273"},
	RString1:  []string{"s11311", "s11312", "s11313", "s11314", "s11315", "s11316", "s11317"},
	RString2:  []string{"s11321", "s11322", "s11323", "s11324", "s11325", "s11326", "s11327"},
	RString3:  []string{"s11331", "s11332", "s11333", "s11334", "s11335", "s11336", "s11337"},
	RString4:  []string{"s11341", "s11342", "s11343", "s11344", "s11345", "s11346", "s11347"},
	RString5:  []string{"s11351", "s11352", "s11353", "s11354", "s11355", "s11356", "s11357"},
	RString6:  []string{"s11361", "s11362", "s11363", "s11364", "s11365", "s11366", "s11367"},
	RString7:  []string{"s11371", "s11372", "s11373", "s11374", "s11375", "s11376", "s11377"},
	RString8:  []string{"s11381", "s11382", "s11383", "s11384", "s11385", "s11386", "s11387"},
	RString9:  []string{"s11391", "s11392", "s11393", "s11394", "s11395", "s11396", "s11397"},
	RInt32:    []int32{11311, 11312, 11313, 11314, 11315, 11316, 11317},
	RInt64:    []int64{11321, 11322, 11323, 11324, 11325, 11326, 11327},
	RUint32:   []uint32{11331, 11332, 11333, 11334, 11335, 11336, 11337},
	RUint64:   []uint64{11341, 11342, 11343, 11344, 11345, 11346, 11347},
	RSint32:   []int32{11351, 11352, 11353, 11354, 11355, 11356, 11357},
	RSint64:   []int64{11361, 11362, 11363, 11364, 11365, 11366, 11367},
	RSfixed32: []int32{11371, 11372, 11373, 11374, 11375, 11376, 11377},
	RSfixed64: []int64{11381, 11382, 11383, 11384, 11385, 11386, 11387},
	RFixed32:  []uint32{11391, 11392, 11393, 11394, 11395, 11396, 11397},
	RFixed64:  []uint64{11401, 11402, 11403, 11404, 11405, 11406, 11407},
	RFloat:    []float32{311.311, 311.312, 311.313, 311.314, 311.315, 311.316, 311.317},
	RDouble:   []float64{411.411, 411.412, 411.413, 411.414, 411.415, 411.416, 411.417},
	RBool1:    []bool{true, true, true, false, true, true, true},
	RBool2:    []bool{true, true, true, false, true, true, true},
	RBool3:    []bool{true, true, true, false, true, true, true},
	RBytes1:   [][]byte{[]byte("b11311"), []byte("b11312"), []byte("b11313"), []byte("b11314"), []byte("b11315")},
	RBytes2:   [][]byte{[]byte("b11411"), []byte("b11412"), []byte("b11413"), []byte("b11414"), []byte("b11415")},
	RBytes3:   [][]byte{[]byte("b11511"), []byte("b11512"), []byte("b11513"), []byte("b11514"), []byte("b11515")},
	REnum1:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum2:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum3:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum4:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum5:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum6:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	REnum7:    []pbbenchmark.Enum1{1, 2, 3, 4, 5, 6, 7},
	RMessage1: []*pbbenchmark.Message1{
		{FString1: "ms1311", FString2: "ms1312", FString3: "ms1313"},
		{FString1: "ms1411", FString2: "ms1412", FString3: "ms1413"},
		{FString1: "ms1511", FString2: "ms1412", FString3: "ms1513"},
	},
	RMessage2: []*pbbenchmark.Message1{
		{FString1: "ms1321", FString2: "ms1322", FString3: "ms1323"},
		{FString1: "ms1421", FString2: "ms1422", FString3: "ms1423"},
		{FString1: "ms1521", FString2: "ms1422", FString3: "ms1523"},
	},
	RMessage3: []*pbbenchmark.Message1{
		{FString1: "ms1331", FString2: "ms1332", FString3: "ms1333"},
		{FString1: "ms1431", FString2: "ms1432", FString3: "ms1433"},
		{FString1: "ms1531", FString2: "ms1432", FString3: "ms1533"},
	},
	RMessage4: []*pbbenchmark.Message1{
		{FString1: "ms1341", FString2: "ms1342", FString3: "ms1343"},
		{FString1: "ms1441", FString2: "ms1442", FString3: "ms1443"},
		{FString1: "ms1541", FString2: "ms1442", FString3: "ms1543"},
	},
	RMessage5: []*pbbenchmark.Message1{
		{FString1: "ms1351", FString2: "ms1352", FString3: "ms1353"},
		{FString1: "ms1451", FString2: "ms1452", FString3: "ms1453"},
		{FString1: "ms1551", FString2: "ms1452", FString3: "ms1553"},
	},
	RMessage6: []*pbbenchmark.Message1{
		{FString1: "ms1361", FString2: "ms1362", FString3: "ms1363"},
		{FString1: "ms1461", FString2: "ms1462", FString3: "ms1463"},
		{FString1: "ms1561", FString2: "ms1462", FString3: "ms1563"},
	},
	RMessage7: []*pbbenchmark.Message1{
		{FString1: "ms1371", FString2: "ms1372", FString3: "ms1373"},
		{FString1: "ms1471", FString2: "ms1472", FString3: "ms1473"},
		{FString1: "ms1571", FString2: "ms1472", FString3: "ms1573"},
	},
	MString1: map[string]string{
		"mk111": "mv111",
		"mk112": "mv112",
		"mk113": "mv113",
		"mk114": "mv114",
		"mk115": "mv115",
		"mk116": "mv116",
		"mk117": "mv117",
	},
	MString2: map[string]string{
		"mk121": "mv121",
		"mk122": "mv122",
		"mk123": "mv123",
		"mk124": "mv124",
		"mk125": "mv125",
		"mk126": "mv126",
		"mk127": "mv127",
	},
	MString3: map[string]string{
		"mk131": "mv131",
		"mk132": "mv132",
		"mk133": "mv133",
		"mk134": "mv134",
		"mk135": "mv135",
		"mk136": "mv136",
		"mk137": "mv137",
	},
	MString4: map[string]string{
		"mk141": "mv141",
		"mk142": "mv142",
		"mk143": "mv143",
		"mk144": "mv144",
		"mk145": "mv145",
		"mk146": "mv146",
		"mk147": "mv147",
	},
	MString5: map[string]string{
		"mk151": "mv151",
		"mk152": "mv152",
		"mk153": "mv153",
		"mk154": "mv154",
		"mk155": "mv155",
		"mk156": "mv156",
		"mk157": "mv157",
	},
	MString6: map[string]string{
		"mk161": "mv161",
		"mk162": "mv162",
		"mk163": "mv163",
		"mk164": "mv164",
		"mk165": "mv165",
		"mk166": "mv166",
		"mk167": "mv167",
	},
	MString7: map[string]string{
		"mk171": "mv171",
		"mk172": "mv172",
		"mk173": "mv173",
		"mk174": "mv174",
		"mk175": "mv175",
		"mk176": "mv176",
		"mk177": "mv177",
	},
	MString8: map[string]string{
		"mk181": "mv181",
		"mk182": "mv182",
		"mk183": "mv183",
		"mk184": "mv184",
		"mk185": "mv185",
		"mk186": "mv186",
		"mk187": "mv187",
	},
	MString9: map[string]string{
		"mk191": "mv191",
		"mk192": "mv192",
		"mk193": "mv193",
		"mk194": "mv194",
		"mk195": "mv195",
		"mk196": "mv196",
		"mk197": "mv197",
	},
	MInt32: map[string]int32{
		"mi111": 111,
		"mi112": 112,
		"mi113": 113,
		"mi114": 114,
		"mi115": 115,
		"mi116": 116,
		"mi117": 117,
	},
	MInt64: map[string]int64{
		"mi121": 121,
		"mi122": 122,
		"mi123": 123,
		"mi124": 124,
		"mi125": 125,
		"mi126": 126,
		"mi127": 127,
	},
	MUint32: map[string]uint32{
		"mi131": 131,
		"mi132": 132,
		"mi133": 133,
		"mi134": 134,
		"mi135": 135,
		"mi136": 136,
		"mi137": 137,
	},
	MUint64: map[string]uint64{
		"mi141": 141,
		"mi142": 142,
		"mi143": 143,
		"mi144": 144,
		"mi145": 145,
		"mi146": 146,
		"mi147": 147,
	},
	MSint32: map[string]int32{
		"mi151": 151,
		"mi152": 152,
		"mi153": 153,
		"mi154": 154,
		"mi155": 155,
		"mi156": 156,
		"mi157": 157,
	},
	MSint64: map[string]int64{
		"mi161": 161,
		"mi162": 162,
		"mi163": 163,
		"mi164": 164,
		"mi165": 165,
		"mi166": 166,
		"mi167": 167,
	},
	MSfixed32: map[string]int32{
		"mi171": 171,
		"mi172": 172,
		"mi173": 173,
		"mi174": 174,
		"mi175": 175,
		"mi176": 176,
		"mi177": 177,
	},
	MSfixed64: map[string]int64{
		"mi181": 181,
		"mi182": 182,
		"mi183": 183,
		"mi184": 184,
		"mi185": 185,
		"mi186": 186,
		"mi187": 187,
	},
	MFixed32: map[string]uint32{
		"mi191": 191,
		"mi192": 192,
		"mi193": 193,
		"mi194": 194,
		"mi195": 195,
		"mi196": 196,
		"mi197": 197,
	},
	MFixed64: map[string]uint64{
		"mi211": 211,
		"mi212": 212,
		"mi213": 213,
		"mi214": 214,
		"mi215": 215,
		"mi216": 216,
		"mi217": 217,
	},
	MFloat: map[string]float32{
		"mf111.111": 111.111,
		"mf112.112": 112.112,
		"mf113.113": 113.113,
		"mf114.114": 114.114,
		"mf115.115": 115.115,
		"mf116.116": 116.116,
		"mf117.117": 117.117,
	},
	MDouble: map[string]float64{
		"mf211.211": 211.211,
		"mf212.212": 212.212,
		"mf213.213": 213.213,
		"mf214.214": 214.214,
		"mf215.215": 215.215,
		"mf216.216": 216.216,
		"mf217.217": 217.217,
	},
	MBool1: map[string]bool{
		"mb111": true,
		"mb112": true,
		"mb113": true,
		"mb114": true,
		"mb115": true,
		"mb116": true,
		"mb117": true,
	},
	MBool2: map[string]bool{
		"mb211": true,
		"mb212": true,
		"mb213": true,
		"mb214": true,
		"mb215": true,
		"mb216": true,
		"mb217": true,
	},
	MBool3: map[string]bool{
		"mb311": true,
		"mb312": true,
		"mb313": true,
		"mb314": true,
		"mb315": true,
		"mb316": true,
		"mb317": true,
	},
	MBytes1: map[string][]byte{
		"mb111": []byte("bv111"),
		"mb112": []byte("bv112"),
		"mb113": []byte("bv113"),
		"mb114": []byte("bv114"),
		"mb115": []byte("bv115"),
		"mb116": []byte("bv116"),
		"mb117": []byte("bv117"),
	},
	MBytes2: map[string][]byte{
		"mb211": []byte("bv211"),
		"mb212": []byte("bv212"),
		"mb213": []byte("bv213"),
		"mb214": []byte("bv214"),
		"mb215": []byte("bv215"),
		"mb216": []byte("bv216"),
		"mb217": []byte("bv217"),
	},
	MBytes3: map[string][]byte{
		"mb311": []byte("bv311"),
		"mb312": []byte("bv312"),
		"mb313": []byte("bv313"),
		"mb314": []byte("bv314"),
		"mb315": []byte("bv315"),
		"mb316": []byte("bv316"),
		"mb317": []byte("bv317"),
	},
	MEnum1: map[string]pbbenchmark.Enum1{
		"me111": 1,
		"me112": 2,
		"me113": 3,
		"me114": 4,
		"me115": 5,
		"me116": 6,
		"me117": 7,
	},
	MEnum2: map[string]pbbenchmark.Enum1{
		"me121": 1,
		"me122": 2,
		"me123": 3,
		"me124": 4,
		"me125": 5,
		"me126": 6,
		"me127": 7,
	},
	MEnum3: map[string]pbbenchmark.Enum1{
		"me131": 1,
		"me132": 2,
		"me133": 3,
		"me134": 4,
		"me135": 5,
		"me136": 6,
		"me137": 7,
	},
	MEnum4: map[string]pbbenchmark.Enum1{
		"me141": 1,
		"me142": 2,
		"me143": 3,
		"me144": 4,
		"me145": 5,
		"me146": 6,
		"me147": 7,
	},
	MEnum5: map[string]pbbenchmark.Enum1{
		"me151": 1,
		"me152": 2,
		"me153": 3,
		"me154": 4,
		"me155": 5,
		"me156": 6,
		"me157": 7,
	},
	MEnum6: map[string]pbbenchmark.Enum1{
		"me161": 1,
		"me162": 2,
		"me163": 3,
		"me164": 4,
		"me165": 5,
		"me166": 6,
		"me167": 7,
	},
	MEnum7: map[string]pbbenchmark.Enum1{
		"me171": 1,
		"me172": 2,
		"me173": 3,
		"me174": 4,
		"me175": 5,
		"me176": 6,
		"me177": 7,
	},
	MMessage1: map[string]*pbbenchmark.Message1{
		"mmk111": {FString1: "mmv111", FString2: "mmv112", FString3: "mmv113"},
		"mmk112": {FString1: "mmv121", FString2: "mmv122", FString3: "mmv123"},
		"mmk113": {FString1: "mmv131", FString2: "mmv132", FString3: "mmv133"},
		"mmk114": {FString1: "mmv141", FString2: "mmv142", FString3: "mmv143"},
		"mmk115": {FString1: "mmv151", FString2: "mmv152", FString3: "mmv153"},
	},
	MMessage2: map[string]*pbbenchmark.Message1{
		"mmk211": {FString1: "mmv211", FString2: "mmv212", FString3: "mmv213"},
		"mmk212": {FString1: "mmv221", FString2: "mmv222", FString3: "mmv223"},
		"mmk213": {FString1: "mmv231", FString2: "mmv232", FString3: "mmv233"},
		"mmk214": {FString1: "mmv241", FString2: "mmv242", FString3: "mmv243"},
		"mmk215": {FString1: "mmv251", FString2: "mmv252", FString3: "mmv253"},
	},
	MMessage3: map[string]*pbbenchmark.Message1{
		"mmk311": {FString1: "mmv311", FString2: "mmv312", FString3: "mmv313"},
		"mmk312": {FString1: "mmv321", FString2: "mmv322", FString3: "mmv323"},
		"mmk313": {FString1: "mmv331", FString2: "mmv332", FString3: "mmv333"},
		"mmk314": {FString1: "mmv341", FString2: "mmv342", FString3: "mmv343"},
		"mmk315": {FString1: "mmv351", FString2: "mmv352", FString3: "mmv353"},
	},
	MMessage4: map[string]*pbbenchmark.Message1{
		"mmk411": {FString1: "mmv411", FString2: "mmv412", FString3: "mmv413"},
		"mmk412": {FString1: "mmv421", FString2: "mmv422", FString3: "mmv423"},
		"mmk413": {FString1: "mmv431", FString2: "mmv432", FString3: "mmv433"},
		"mmk414": {FString1: "mmv441", FString2: "mmv442", FString3: "mmv443"},
		"mmk415": {FString1: "mmv451", FString2: "mmv452", FString3: "mmv453"},
	},
	MMessage5: map[string]*pbbenchmark.Message1{
		"mmk511": {FString1: "mmv511", FString2: "mmv512", FString3: "mmv513"},
		"mmk512": {FString1: "mmv521", FString2: "mmv522", FString3: "mmv523"},
		"mmk513": {FString1: "mmv531", FString2: "mmv532", FString3: "mmv533"},
		"mmk514": {FString1: "mmv541", FString2: "mmv542", FString3: "mmv543"},
		"mmk515": {FString1: "mmv551", FString2: "mmv552", FString3: "mmv553"},
	},
	MMessage6: map[string]*pbbenchmark.Message1{
		"mmk611": {FString1: "mmv611", FString2: "mmv612", FString3: "mmv613"},
		"mmk612": {FString1: "mmv621", FString2: "mmv622", FString3: "mmv623"},
		"mmk613": {FString1: "mmv631", FString2: "mmv632", FString3: "mmv633"},
		"mmk614": {FString1: "mmv641", FString2: "mmv642", FString3: "mmv643"},
		"mmk615": {FString1: "mmv651", FString2: "mmv652", FString3: "mmv653"},
	},
	MMessage7: map[string]*pbbenchmark.Message1{
		"mmk711": {FString1: "mmv711", FString2: "mmv712", FString3: "mmv713"},
		"mmk712": {FString1: "mmv721", FString2: "mmv722", FString3: "mmv723"},
		"mmk713": {FString1: "mmv731", FString2: "mmv732", FString3: "mmv733"},
		"mmk714": {FString1: "mmv741", FString2: "mmv742", FString3: "mmv743"},
		"mmk715": {FString1: "mmv751", FString2: "mmv752", FString3: "mmv753"},
	},
}

func Benchmark_Complex_Marshal_Plugin(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := seedComplex.MarshalJSON()
			if err != nil {
				b.Fatal("complex marshal by plugin error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Complex_Marshal_Proto(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := utils.PMarshal.Marshal(seedComplex)
			if err != nil {
				b.Fatal("complex marshal by proto error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Complex_Marshal_StdJSON(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := json.Marshal(copyComplex)
			if err != nil {
				b.Fatal("complex marshal by standard json error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Complex_Marshal_JSONIter(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb, err := _json.Marshal(&copyComplex)
			if err != nil {
				b.Fatal("complex marshal by json iterator error:", err)
			}
			_ = bb
		}
	})
}

func Benchmark_Complex_Unmarshal_Plugin(b *testing.B) {
	bb, err := seedComplex.MarshalJSON()
	if err != nil {
		b.Fatal("complex marshal by plugin error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &pbbenchmark.BenchModelComplex{}
			err = dataNew.UnmarshalJSON(bb)
			if err != nil {
				b.Fatal("complex unmarshal by plugin error:", err)
			}
		}
	})
}

func Benchmark_Complex_Unmarshal_Proto(b *testing.B) {
	bb, err := utils.PMarshal.Marshal(seedComplex)
	if err != nil {
		b.Fatal("complex marshal by proto error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &pbbenchmark.BenchModelComplex{}
			err = utils.PUnmarshal.Unmarshal(bb, dataNew)
			if err != nil {
				b.Fatal("complex unmarshal by proto error:", err)
			}
		}
	})
}

func Benchmark_Complex_Unmarshal_StdJSON(b *testing.B) {
	bb, err := json.Marshal(copyComplex)
	if err != nil {
		b.Fatal("complex marshal by standard json error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &CopyComplex{}
			err = json.Unmarshal(bb, dataNew)
			if err != nil {
				b.Fatal("complex unmarshal by standard json error:", err)
			}
		}
	})
}

func Benchmark_Complex_Unmarshal_JSONIter(b *testing.B) {
	_json := jsoniter.ConfigCompatibleWithStandardLibrary

	bb, err := _json.Marshal(&copyComplex)
	if err != nil {
		b.Fatal("complex marshal by json iterator error:", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dataNew := &CopyComplex{}
			err = _json.Unmarshal(bb, dataNew)
			if err != nil {
				b.Fatal("complex unmarshal by json iterator error:", err)
			}
		}
	})
}
