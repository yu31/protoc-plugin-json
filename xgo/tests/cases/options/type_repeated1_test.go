package options

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_CustomKey_TypeRepeated1(t *testing.T) {
	seed1 := &pboptions.TypeRepeated1{
		FInt32:    []int32{111, 112, 113},
		FInt64:    []int64{121, 122, 123},
		FUint32:   []uint32{131, 132, 133},
		FUint64:   []uint64{141, 142, 143},
		FSint32:   []int32{151, 152, 153},
		FSint64:   []int64{161, 162, 163},
		FSfixed32: []int32{171, 172, 173},
		FSfixed64: []int64{181, 182, 183},
		FFixed32:  []uint32{191, 192, 193},
		FFixed64:  []uint64{201, 202, 203},
		FFloat:    []float32{211.111, 212.111, 213.111},
		FDouble:   []float64{221.111, 222.111, 223.111},
		FBool1:    []bool{true, false},
		FString1:  []string{"ss101", "ss102", "ss103"},
		FBytes1:   [][]byte{[]byte("bb101"), []byte("bb102"), []byte("bb103")},
		FEnum1:    []pbexternal.EnumNum1{1, 3, 5},
		FMessage1: []*pbexternal.Message1{
			{FString1: "ms101", FString2: "ms102", FString3: "ms103"},
			{FString1: "ms201", FString2: "ms202", FString3: "ms203"},
			{FString1: "ms301", FString2: "ms302", FString3: "ms303"},
		},
		FAny1: []*anypb.Any{
			utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s201", FString2: "s202", FString3: "s203"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
		},
		FDuration1: []*durationpb.Duration{
			{Seconds: 1800, Nanos: 0},
			{Seconds: 3600, Nanos: 0},
			{Seconds: 7200, Nanos: 0},
		},
		FTimestamp1: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0},
			{Seconds: 1686416585, Nanos: 0},
			{Seconds: 1686416585, Nanos: 0},
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
		dataNew := &pboptions.TypeRepeated1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		data := map[string]interface{}{}
		err = json.Unmarshal(bb, &data)
		require.Nil(t, err)

		/*
			{
				"t_int32":[111,112,113],
				"t_int64":[121,122,123],
				"t_uint32":[131,132,133],
				"t_uint64":[141,142,143],
				"t_sint32":[151,152,153],
				"t_sint64":[161,162,163],
				"t_sfixed32":[171,172,173],
				"t_sfixed64":[181,182,183],
				"t_fixed32":[191,192,193],
				"t_fixed64":[201,202,203],
				"t_float":[211.111,212.111,213.111],
				"t_double":[221.111,222.111,223.111],
				"t_bool1":[true,false],
				"t_string1":["ss101","ss102","ss103"],
				"t_bytes1":["YmIxMDE=","YmIxMDI=","YmIxMDM="],
				"t_enum1":[1,3,5],
				"t_message1":[
					{"f_string1":"ms101","f_string2":"ms102","f_string3":"ms103"},
					{"f_string1":"ms201","f_string2":"ms202","f_string3":"ms203"},
					{"f_string1":"ms301","f_string2":"ms302","f_string3":"ms303"}
				],
				"t_any1":[
					{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
					{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMjAxEgRzMjAyGgRzMjAz"},
					{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMzAxEgRzMzAyGgRzMzAz"}
				],
				"t_duration1":[{"seconds":1800},{"seconds":3600},{"seconds":7200}],
				"t_timestamp1":[{"seconds":1686416585},{"seconds":1686416585},{"seconds":1686416585}]
			}
		*/
		t.Run("int32", func(t *testing.T) {
			tInt32 := data["t_int32"].([]interface{})
			require.Equal(t, float64(111), tInt32[0].(float64))
			require.Equal(t, float64(112), tInt32[1].(float64))
			require.Equal(t, float64(113), tInt32[2].(float64))
		})
		t.Run("int64", func(t *testing.T) {
			tInt64 := data["t_int64"].([]interface{})
			require.Equal(t, float64(121), tInt64[0].(float64))
			require.Equal(t, float64(122), tInt64[1].(float64))
			require.Equal(t, float64(123), tInt64[2].(float64))
		})
		t.Run("uint32", func(t *testing.T) {
			tUint32 := data["t_uint32"].([]interface{})
			require.Equal(t, float64(131), tUint32[0].(float64))
			require.Equal(t, float64(132), tUint32[1].(float64))
			require.Equal(t, float64(133), tUint32[2].(float64))
		})
		t.Run("uint64", func(t *testing.T) {
			tUint64 := data["t_uint64"].([]interface{})
			require.Equal(t, float64(141), tUint64[0].(float64))
			require.Equal(t, float64(142), tUint64[1].(float64))
			require.Equal(t, float64(143), tUint64[2].(float64))
		})
		t.Run("sint32", func(t *testing.T) {
			tSInt32 := data["t_sint32"].([]interface{})
			require.Equal(t, float64(151), tSInt32[0].(float64))
			require.Equal(t, float64(152), tSInt32[1].(float64))
			require.Equal(t, float64(153), tSInt32[2].(float64))
		})
		t.Run("sint64", func(t *testing.T) {
			tSInt64 := data["t_sint64"].([]interface{})
			require.Equal(t, float64(161), tSInt64[0].(float64))
			require.Equal(t, float64(162), tSInt64[1].(float64))
			require.Equal(t, float64(163), tSInt64[2].(float64))
		})
		t.Run("sfixed32", func(t *testing.T) {
			tSFixed32 := data["t_sfixed32"].([]interface{})
			require.Equal(t, float64(171), tSFixed32[0].(float64))
			require.Equal(t, float64(172), tSFixed32[1].(float64))
			require.Equal(t, float64(173), tSFixed32[2].(float64))
		})
		t.Run("sfixed64", func(t *testing.T) {
			tSFixed64 := data["t_sfixed64"].([]interface{})
			require.Equal(t, float64(181), tSFixed64[0].(float64))
			require.Equal(t, float64(182), tSFixed64[1].(float64))
			require.Equal(t, float64(183), tSFixed64[2].(float64))
		})
		t.Run("fixed32", func(t *testing.T) {
			tFixed32 := data["t_fixed32"].([]interface{})
			require.Equal(t, float64(191), tFixed32[0].(float64))
			require.Equal(t, float64(192), tFixed32[1].(float64))
			require.Equal(t, float64(193), tFixed32[2].(float64))
		})
		t.Run("fixed64", func(t *testing.T) {
			tFixed64 := data["t_fixed64"].([]interface{})
			require.Equal(t, float64(201), tFixed64[0].(float64))
			require.Equal(t, float64(202), tFixed64[1].(float64))
			require.Equal(t, float64(203), tFixed64[2].(float64))
		})
		t.Run("float", func(t *testing.T) {
			tFloat := data["t_float"].([]interface{})
			require.Equal(t, float64(211.111), tFloat[0].(float64))
			require.Equal(t, float64(212.111), tFloat[1].(float64))
			require.Equal(t, float64(213.111), tFloat[2].(float64))
		})
		t.Run("double", func(t *testing.T) {
			tDouble := data["t_double"].([]interface{})
			require.Equal(t, float64(221.111), tDouble[0].(float64))
			require.Equal(t, float64(222.111), tDouble[1].(float64))
			require.Equal(t, float64(223.111), tDouble[2].(float64))
		})
		t.Run("bool1", func(t *testing.T) {
			tBool1 := data["t_bool1"].([]interface{})
			require.Equal(t, true, tBool1[0].(bool))
			require.Equal(t, false, tBool1[1].(bool))
		})
		t.Run("string1", func(t *testing.T) {
			tString1 := data["t_string1"].([]interface{})
			require.Equal(t, "ss101", tString1[0].(string))
			require.Equal(t, "ss102", tString1[1].(string))
			require.Equal(t, "ss103", tString1[2].(string))
		})
		t.Run("bytes1", func(t *testing.T) {
			tBytes1 := data["t_bytes1"].([]interface{})
			require.Equal(t, "YmIxMDE=", tBytes1[0].(string))
			require.Equal(t, "YmIxMDI=", tBytes1[1].(string))
			require.Equal(t, "YmIxMDM=", tBytes1[2].(string))
		})
		t.Run("enum1", func(t *testing.T) {
			tEnum1 := data["t_enum1"].([]interface{})
			require.Equal(t, float64(1), tEnum1[0].(float64))
			require.Equal(t, float64(3), tEnum1[1].(float64))
			require.Equal(t, float64(5), tEnum1[2].(float64))
		})
		t.Run("message1", func(t *testing.T) {
			tMessage1 := data["t_message1"].([]interface{})

			message1 := tMessage1[0].(map[string]interface{})
			require.Equal(t, "ms101", message1["f_string1"].(string))
			require.Equal(t, "ms102", message1["f_string2"].(string))
			require.Equal(t, "ms103", message1["f_string3"].(string))

			message2 := tMessage1[1].(map[string]interface{})
			require.Equal(t, "ms201", message2["f_string1"].(string))
			require.Equal(t, "ms202", message2["f_string2"].(string))
			require.Equal(t, "ms203", message2["f_string3"].(string))

			message3 := tMessage1[2].(map[string]interface{})
			require.Equal(t, "ms301", message3["f_string1"].(string))
			require.Equal(t, "ms302", message3["f_string2"].(string))
			require.Equal(t, "ms303", message3["f_string3"].(string))
		})
		t.Run("any1", func(t *testing.T) {
			tAny1 := data["t_any1"].([]interface{})

			any1 := tAny1[0].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", any1["type_url"].(string))
			require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", any1["value"].(string))

			any2 := tAny1[1].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", any2["type_url"].(string))
			require.Equal(t, "CgRzMjAxEgRzMjAyGgRzMjAz", any2["value"].(string))

			any3 := tAny1[2].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", any3["type_url"].(string))
			require.Equal(t, "CgRzMzAxEgRzMzAyGgRzMzAz", any3["value"].(string))

		})
		t.Run("duration1", func(t *testing.T) {
			tDuration1 := data["t_duration1"].([]interface{})

			duration1 := tDuration1[0].(map[string]interface{})
			require.Equal(t, float64(1800), duration1["seconds"].(float64))

			duration2 := tDuration1[1].(map[string]interface{})
			require.Equal(t, float64(3600), duration2["seconds"].(float64))

			duration3 := tDuration1[2].(map[string]interface{})
			require.Equal(t, float64(7200), duration3["seconds"].(float64))
		})
		t.Run("timestamp1", func(t *testing.T) {
			tTimestamp1 := data["t_timestamp1"].([]interface{})

			timestamp1 := tTimestamp1[0].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestamp1["seconds"].(float64))

			timestamp2 := tTimestamp1[1].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestamp2["seconds"].(float64))

			timestamp3 := tTimestamp1[2].(map[string]interface{})
			require.Equal(t, float64(1686416585), timestamp3["seconds"].(float64))
		})
	})
}

func Test_Omitempty_TypeRepeated1(t *testing.T) {
	seed1 := &pboptions.TypeRepeated1{
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
		dataNew := &pboptions.TypeRepeated1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		require.Equal(t, "{}", string(bb))
	})
}
