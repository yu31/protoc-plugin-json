package options

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_CustomKey_TypeOneOf1(t *testing.T) {
	seed1 := &pboptions.TypeOneOf1{
		OneType1:  &pboptions.TypeOneOf1_FInt32{FInt32: 111},
		OneType2:  &pboptions.TypeOneOf1_FInt64{FInt64: 121},
		OneType3:  &pboptions.TypeOneOf1_FUint32{FUint32: 131},
		OneType4:  &pboptions.TypeOneOf1_FUint64{FUint64: 141},
		OneType5:  &pboptions.TypeOneOf1_FSint32{FSint32: 151},
		OneType6:  &pboptions.TypeOneOf1_FSint64{FSint64: 161},
		OneType7:  &pboptions.TypeOneOf1_FSfixed32{FSfixed32: 171},
		OneType8:  &pboptions.TypeOneOf1_FSfixed64{FSfixed64: 181},
		OneType9:  &pboptions.TypeOneOf1_FFixed32{FFixed32: 191},
		OneType10: &pboptions.TypeOneOf1_FFixed64{FFixed64: 201},
		OneType11: &pboptions.TypeOneOf1_FFloat{FFloat: 211.111},
		OneType12: &pboptions.TypeOneOf1_FDouble{FDouble: 221.111},
		OneType13: &pboptions.TypeOneOf1_FBool1{FBool1: true},
		OneType14: &pboptions.TypeOneOf1_FString1{FString1: "ss101"},
		OneType15: &pboptions.TypeOneOf1_FBytes1{FBytes1: []byte("bb101")},
		OneType16: &pboptions.TypeOneOf1_FEnum1{FEnum1: 1},
		OneType17: &pboptions.TypeOneOf1_FMessage1{FMessage1: &pbexternal.Message1{
			FString1: "ms101",
			FString2: "ms102",
			FString3: "ms103",
		}},
		OneType18: &pboptions.TypeOneOf1_FAny1{
			FAny1: utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		},
		OneType19: &pboptions.TypeOneOf1_FDuration1{
			FDuration1: &durationpb.Duration{Seconds: 3600, Nanos: 0},
		},
		OneType20: &pboptions.TypeOneOf1_FTimestamp1{
			FTimestamp1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
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
		dataNew := &pboptions.TypePlain1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)
		/*
			{
				"t_one1":{"t_int32":111},
				"t_one2":{"t_int64":121},
				"t_one3":{"t_uint32":131},
				"t_one4":{"t_uint64":141},
				"t_one5":{"t_sint32":151},
				"t_one6":{"t_sint64":161},
				"t_one7":{"t_sfixed32":171},
				"t_one8":{"t_sfixed64":181},
				"t_one9":{"t_fixed32":191},
				"t_one10":{"t_fixed64":201},
				"t_one11":{"t_float":211.111},
				"t_one12":{"t_double":221.111},
				"t_one13":{"t_bool1":true},
				"t_one14":{"t_string1":"ss101"},
				"t_one15":{"t_bytes1":"YmIxMDE="},
				"t_one16":{"t_enum1":1},
				"t_one17":{"t_message1":{"f_string1":"ms101","f_string2":"ms102","f_string3":"ms103"}},
				"t_one18":{"t_any1":{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"}},
				"t_one19":{"t_duration1":{"seconds":3600}},
				"t_one20":{"t_timestamp1":{"seconds":1686416585}}
			}
		*/
		t.Run("t_one1", func(t *testing.T) {
			tOne1 := data["t_one1"].(map[string]interface{})
			require.Equal(t, float64(111), tOne1["t_int32"].(float64))
		})
		t.Run("t_one2", func(t *testing.T) {
			tOne1 := data["t_one2"].(map[string]interface{})
			require.Equal(t, float64(121), tOne1["t_int64"].(float64))
		})
		t.Run("t_one3", func(t *testing.T) {
			tOne1 := data["t_one3"].(map[string]interface{})
			require.Equal(t, float64(131), tOne1["t_uint32"].(float64))
		})
		t.Run("t_one4", func(t *testing.T) {
			tOne1 := data["t_one4"].(map[string]interface{})
			require.Equal(t, float64(141), tOne1["t_uint64"].(float64))
		})
		t.Run("t_one5", func(t *testing.T) {
			tOne1 := data["t_one5"].(map[string]interface{})
			require.Equal(t, float64(151), tOne1["t_sint32"].(float64))
		})
		t.Run("t_one6", func(t *testing.T) {
			tOne1 := data["t_one6"].(map[string]interface{})
			require.Equal(t, float64(161), tOne1["t_sint64"].(float64))
		})
		t.Run("t_one7", func(t *testing.T) {
			tOne1 := data["t_one7"].(map[string]interface{})
			require.Equal(t, float64(171), tOne1["t_sfixed32"].(float64))
		})
		t.Run("t_one8", func(t *testing.T) {
			tOne1 := data["t_one8"].(map[string]interface{})
			require.Equal(t, float64(181), tOne1["t_sfixed64"].(float64))
		})
		t.Run("t_one9", func(t *testing.T) {
			tOne1 := data["t_one9"].(map[string]interface{})
			require.Equal(t, float64(191), tOne1["t_fixed32"].(float64))
		})
		t.Run("t_one10", func(t *testing.T) {
			tOne1 := data["t_one10"].(map[string]interface{})
			require.Equal(t, float64(201), tOne1["t_fixed64"].(float64))
		})
		t.Run("t_one11", func(t *testing.T) {
			tOne1 := data["t_one11"].(map[string]interface{})
			require.Equal(t, float64(211.111), tOne1["t_float"].(float64))
		})
		t.Run("t_one12", func(t *testing.T) {
			tOne1 := data["t_one12"].(map[string]interface{})
			require.Equal(t, float64(221.111), tOne1["t_double"].(float64))
		})
		t.Run("t_one13", func(t *testing.T) {
			tOne1 := data["t_one13"].(map[string]interface{})
			require.Equal(t, true, tOne1["t_bool1"].(bool))
		})
		t.Run("t_one14", func(t *testing.T) {
			tOne1 := data["t_one14"].(map[string]interface{})
			require.Equal(t, "ss101", tOne1["t_string1"].(string))
		})
		t.Run("t_one15", func(t *testing.T) {
			tOne1 := data["t_one15"].(map[string]interface{})
			require.Equal(t, "YmIxMDE=", tOne1["t_bytes1"].(string))
		})
		t.Run("t_one16", func(t *testing.T) {
			tOne1 := data["t_one16"].(map[string]interface{})
			require.Equal(t, float64(1), tOne1["t_enum1"].(float64))
		})
		t.Run("t_one17", func(t *testing.T) {
			tOne17 := data["t_one17"].(map[string]interface{})
			tMessage1 := tOne17["t_message1"].(map[string]interface{})
			require.Equal(t, "ms101", tMessage1["f_string1"].(string))
			require.Equal(t, "ms102", tMessage1["f_string2"].(string))
			require.Equal(t, "ms103", tMessage1["f_string3"].(string))
		})
		t.Run("t_one18", func(t *testing.T) {
			tOne18 := data["t_one18"].(map[string]interface{})
			tAny1 := tOne18["t_any1"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", tAny1["type_url"].(string))
			require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", tAny1["value"].(string))
		})
		t.Run("t_one19", func(t *testing.T) {
			tOne19 := data["t_one19"].(map[string]interface{})
			tDuration1 := tOne19["t_duration1"].(map[string]interface{})
			require.Equal(t, float64(3600), tDuration1["seconds"].(float64))
		})
		t.Run("t_one20", func(t *testing.T) {
			tOne20 := data["t_one20"].(map[string]interface{})
			tTimestamp1 := tOne20["t_timestamp1"].(map[string]interface{})
			require.Equal(t, float64(1686416585), tTimestamp1["seconds"].(float64))
		})
	})
}
