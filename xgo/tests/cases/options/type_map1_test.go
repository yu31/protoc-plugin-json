package options

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	anypb "google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_CustomKey_TypeMap1(t *testing.T) {
	seed1 := &pboptions.TypeMap1{
		FInt32:    map[int32]int32{111: 110, 112: 111, 113: 112},
		FInt64:    map[int64]int64{121: 120, 122: 121, 123: 122},
		FUint32:   map[uint32]uint32{131: 130, 132: 131, 133: 132},
		FUint64:   map[uint64]uint64{141: 140, 142: 141, 143: 142},
		FSint32:   map[int32]int32{151: 150, 152: 151, 153: 152},
		FSint64:   map[int64]int64{161: 160, 162: 161, 163: 162},
		FSfixed32: map[int32]int32{171: 170, 172: 171, 173: 172},
		FSfixed64: map[int64]int64{181: 180, 182: 181, 183: 182},
		FFixed32:  map[uint32]uint32{191: 190, 192: 191, 193: 192},
		FFixed64:  map[uint64]uint64{201: 200, 202: 201, 203: 202},
		FFloat:    map[string]float32{"f1": 211.111, "f2": 212.111, "f3": 213.111},
		FDouble:   map[string]float64{"d1": 221.111, "d2": 222.111, "d3": 223.111},
		FBool1:    map[bool]bool{true: false, false: true},
		FString1:  map[string]string{"sk101": "sv101", "sk102": "sv102", "sk103": "sv103"},
		FBytes1:   map[string][]byte{"bk101": []byte("bb101"), "bk102": []byte("bb102"), "bk103": []byte("bb103")},
		FEnum1:    map[string]pbexternal.Enum1{"ek1": 1, "ek2": 3, "ek3": 5},
		FMessage1: map[string]*pbexternal.Message1{
			"mk101": {FString1: "ms101", FString2: "ms102", FString3: "ms103"},
			"mk201": {FString1: "ms201", FString2: "ms202", FString3: "ms203"},
			"mk301": {FString1: "ms301", FString2: "ms302", FString3: "ms303"},
		},
		FAny1: map[string]*anypb.Any{
			"ak101": utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
			"ak201": utils.MustNewAny(&pbexternal.Message1{FString1: "s201", FString2: "s202", FString3: "s203"}),
			"ak301": utils.MustNewAny(&pbexternal.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
		},
		FDuration1: map[string]*durationpb.Duration{
			"dk101": {Seconds: 1800, Nanos: 0},
			"dk201": {Seconds: 3600, Nanos: 0},
			"dk301": {Seconds: 7200, Nanos: 0},
		},
		FTimestamp1: map[string]*timestamppb.Timestamp{
			"tk101": {Seconds: 1686416585, Nanos: 0},
			"tk201": {Seconds: 1686416585, Nanos: 0},
			"tk301": {Seconds: 1686416585, Nanos: 0},
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
		dataNew := &pboptions.TypeMap1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)
		/*
			{
				"t_int32":{"113":112,"111":110,"112":111},
				"t_int64":{"123":122,"121":120,"122":121},
				"t_uint32":{"131":130,"132":131,"133":132},
				"t_uint64":{"141":140,"142":141,"143":142},
				"t_sint32":{"152":151,"153":152,"151":150},
				"t_sint64":{"161":160,"162":161,"163":162},
				"t_sfixed32":{"171":170,"172":171,"173":172},
				"t_sfixed64":{"181":180,"182":181,"183":182},
				"t_fixed32":{"191":190,"192":191,"193":192},
				"t_fixed64":{"202":201,"203":202,"201":200},
				"t_float":{"f1":211.111,"f2":212.111,"f3":213.111},
				"t_double":{"d2":222.111,"d3":223.111,"d1":221.111},
				"t_bool1":{"true":false,"false":true},
				"t_string1":{"sk102":"sv102","sk103":"sv103","sk101":"sv101"},
				"t_bytes1":{"bk101":"YmIxMDE=","bk102":"YmIxMDI=","bk103":"YmIxMDM="},
				"t_enum1":{"ek1":1,"ek2":3,"ek3":5},
				"t_message1":{
					"mk301":{"f_string1":"ms301","f_string2":"ms302","f_string3":"ms303"},
					"mk101":{"f_string1":"ms101","f_string2":"ms102","f_string3":"ms103"},
					"mk201":{"f_string1":"ms201","f_string2":"ms202","f_string3":"ms203"}
				},
				"t_any1":{
					"ak101":{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
					"ak201":{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMjAxEgRzMjAyGgRzMjAz"},
					"ak301":{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMzAxEgRzMjAyGgRzMzAz"}
				},
				"t_duration1":{"dk301":{"seconds":7200},"dk101":{"seconds":1800},"dk201":{"seconds":3600}},
				"t_timestamp1":{"tk101":{"seconds":1686416585},"tk201":{"seconds":1686416585},"tk301":{"seconds":1686416585}}}
		*/

		t.Run("int32", func(t *testing.T) {
			tInt32 := data["t_int32"].(map[string]interface{})
			require.Equal(t, float64(110), tInt32["111"].(float64))
			require.Equal(t, float64(111), tInt32["112"].(float64))
			require.Equal(t, float64(112), tInt32["113"].(float64))
		})
		t.Run("int64", func(t *testing.T) {
			tInt64 := data["t_int64"].(map[string]interface{})
			require.Equal(t, float64(120), tInt64["121"].(float64))
			require.Equal(t, float64(121), tInt64["122"].(float64))
			require.Equal(t, float64(122), tInt64["123"].(float64))
		})
		t.Run("uint32", func(t *testing.T) {
			tUint32 := data["t_uint32"].(map[string]interface{})
			require.Equal(t, float64(130), tUint32["131"].(float64))
			require.Equal(t, float64(131), tUint32["132"].(float64))
			require.Equal(t, float64(132), tUint32["133"].(float64))
		})
		t.Run("uint64", func(t *testing.T) {
			tUint64 := data["t_uint64"].(map[string]interface{})
			require.Equal(t, float64(140), tUint64["141"].(float64))
			require.Equal(t, float64(141), tUint64["142"].(float64))
			require.Equal(t, float64(142), tUint64["143"].(float64))
		})
		t.Run("sint32", func(t *testing.T) {
			tSInt32 := data["t_sint32"].(map[string]interface{})
			require.Equal(t, float64(150), tSInt32["151"].(float64))
			require.Equal(t, float64(151), tSInt32["152"].(float64))
			require.Equal(t, float64(152), tSInt32["153"].(float64))
		})
		t.Run("sint64", func(t *testing.T) {
			tSInt64 := data["t_sint64"].(map[string]interface{})
			require.Equal(t, float64(160), tSInt64["161"].(float64))
			require.Equal(t, float64(161), tSInt64["162"].(float64))
			require.Equal(t, float64(162), tSInt64["163"].(float64))
		})
		t.Run("sfixed32", func(t *testing.T) {
			tSFixed32 := data["t_sfixed32"].(map[string]interface{})
			require.Equal(t, float64(170), tSFixed32["171"].(float64))
			require.Equal(t, float64(171), tSFixed32["172"].(float64))
			require.Equal(t, float64(172), tSFixed32["173"].(float64))
		})
		t.Run("sfixed64", func(t *testing.T) {
			tSFixed64 := data["t_sfixed64"].(map[string]interface{})
			require.Equal(t, float64(180), tSFixed64["181"].(float64))
			require.Equal(t, float64(181), tSFixed64["182"].(float64))
			require.Equal(t, float64(182), tSFixed64["183"].(float64))
		})
		t.Run("fixed32", func(t *testing.T) {
			tFixed32 := data["t_fixed32"].(map[string]interface{})
			require.Equal(t, float64(190), tFixed32["191"].(float64))
			require.Equal(t, float64(191), tFixed32["192"].(float64))
			require.Equal(t, float64(192), tFixed32["193"].(float64))
		})
		t.Run("fixed64", func(t *testing.T) {
			tFixed64 := data["t_fixed64"].(map[string]interface{})
			require.Equal(t, float64(200), tFixed64["201"].(float64))
			require.Equal(t, float64(201), tFixed64["202"].(float64))
			require.Equal(t, float64(202), tFixed64["203"].(float64))
		})
		t.Run("float", func(t *testing.T) {
			tFloat := data["t_float"].(map[string]interface{})
			require.Equal(t, float64(211.111), tFloat["f1"].(float64))
			require.Equal(t, float64(212.111), tFloat["f2"].(float64))
			require.Equal(t, float64(213.111), tFloat["f3"].(float64))
		})
		t.Run("double", func(t *testing.T) {
			tDouble := data["t_double"].(map[string]interface{})
			require.Equal(t, float64(221.111), tDouble["d1"].(float64))
			require.Equal(t, float64(222.111), tDouble["d2"].(float64))
			require.Equal(t, float64(223.111), tDouble["d3"].(float64))
		})
		t.Run("bool1", func(t *testing.T) {
			tBool1 := data["t_bool1"].(map[string]interface{})
			require.Equal(t, true, tBool1["false"].(bool))
			require.Equal(t, false, tBool1["true"].(bool))
		})
		t.Run("string1", func(t *testing.T) {
			tString1 := data["t_string1"].(map[string]interface{})
			require.Equal(t, "sv101", tString1["sk101"].(string))
			require.Equal(t, "sv102", tString1["sk102"].(string))
			require.Equal(t, "sv103", tString1["sk103"].(string))
		})
		t.Run("bytes1", func(t *testing.T) {
			tBytes1 := data["t_bytes1"].(map[string]interface{})
			require.Equal(t, "YmIxMDE=", tBytes1["bk101"].(string))
			require.Equal(t, "YmIxMDI=", tBytes1["bk102"].(string))
			require.Equal(t, "YmIxMDM=", tBytes1["bk103"].(string))
		})
		t.Run("enum1", func(t *testing.T) {
			tEnum1 := data["t_enum1"].(map[string]interface{})
			require.Equal(t, float64(1), tEnum1["ek1"].(float64))
			require.Equal(t, float64(3), tEnum1["ek2"].(float64))
			require.Equal(t, float64(5), tEnum1["ek3"].(float64))
		})
		t.Run("message1", func(t *testing.T) {
			tMessage1 := data["t_message1"].(map[string]interface{})

			message1 := tMessage1["mk101"].(map[string]interface{})
			require.Equal(t, "ms101", message1["f_string1"].(string))
			require.Equal(t, "ms102", message1["f_string2"].(string))
			require.Equal(t, "ms103", message1["f_string3"].(string))

			message2 := tMessage1["mk201"].(map[string]interface{})
			require.Equal(t, "ms201", message2["f_string1"].(string))
			require.Equal(t, "ms202", message2["f_string2"].(string))
			require.Equal(t, "ms203", message2["f_string3"].(string))

			message3 := tMessage1["mk301"].(map[string]interface{})
			require.Equal(t, "ms301", message3["f_string1"].(string))
			require.Equal(t, "ms302", message3["f_string2"].(string))
			require.Equal(t, "ms303", message3["f_string3"].(string))
		})
		t.Run("any1", func(t *testing.T) {
			tAny1 := data["t_any1"].(map[string]interface{})

			any1 := tAny1["ak101"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", any1["type_url"].(string))
			require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", any1["value"].(string))

			any2 := tAny1["ak201"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", any2["type_url"].(string))
			require.Equal(t, "CgRzMjAxEgRzMjAyGgRzMjAz", any2["value"].(string))

			any3 := tAny1["ak301"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", any3["type_url"].(string))
			require.Equal(t, "CgRzMzAxEgRzMzAyGgRzMzAz", any3["value"].(string))

		})
		t.Run("duration1", func(t *testing.T) {
			tDuration1 := data["t_duration1"].(map[string]interface{})

			duration1 := tDuration1["dk101"].(map[string]interface{})
			require.Equal(t, float64(1800), duration1["seconds"].(float64))

			duration2 := tDuration1["dk201"].(map[string]interface{})
			require.Equal(t, float64(3600), duration2["seconds"].(float64))

			duration3 := tDuration1["dk301"].(map[string]interface{})
			require.Equal(t, float64(7200), duration3["seconds"].(float64))
		})
		t.Run("timestamp1", func(t *testing.T) {
			tTimestamp1 := data["t_timestamp1"].(map[string]interface{})

			timestamp1 := tTimestamp1["tk101"].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestamp1["seconds"].(float64))

			timestamp2 := tTimestamp1["tk201"].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestamp2["seconds"].(float64))

			timestamp3 := tTimestamp1["tk301"].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestamp3["seconds"].(float64))
		})
	})
}

func Test_Omitempty_TypeMap1(t *testing.T) {
	seed1 := &pboptions.TypeMap1{
		FInt32:      nil,
		FInt64:      nil,
		FUint32:     nil,
		FUint64:     nil,
		FSint32:     nil,
		FSint64:     nil,
		FSfixed32:   nil,
		FSfixed64:   nil,
		FFixed32:    nil,
		FFixed64:    nil,
		FFloat:      nil,
		FDouble:     nil,
		FBool1:      nil,
		FString1:    nil,
		FBytes1:     nil,
		FEnum1:      nil,
		FMessage1:   nil,
		FAny1:       nil,
		FDuration1:  nil,
		FTimestamp1: nil,
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
		dataNew := &pboptions.TypeMap1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		require.Equal(t, "{}", string(bb))
	})
}
