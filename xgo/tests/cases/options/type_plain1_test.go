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

func Test_CustomKey_TypePlain1(t *testing.T) {
	seed1 := &pboptions.TypePlain1{
		FInt32:    111,
		FInt64:    121,
		FUint32:   131,
		FUint64:   141,
		FSint32:   151,
		FSint64:   161,
		FSfixed32: 171,
		FSfixed64: 181,
		FFixed32:  191,
		FFixed64:  201,
		FFloat:    211.111,
		FDouble:   221.111,
		FBool1:    true,
		FString1:  "ss101",
		FBytes1:   []byte("bb101"),
		FEnum1:    1,
		FMessage1: &pbexternal.Message1{
			FString1: "ms101",
			FString2: "ms102",
			FString3: "ms103",
		},
		FAny1:       utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		FDuration1:  &durationpb.Duration{Seconds: 3600, Nanos: 0},
		FTimestamp1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
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
				"t_int32":111,
				"t_int64":121,
				"t_uint32":131,
				"t_uint64":141,
				"t_sint32":151,
				"t_sint64":161,
				"t_sfixed32":171,
				"t_sfixed64":181,
				"t_fixed32":191,
				"t_fixed64":201,
				"t_float":211.111,
				"t_double":221.111,
				"t_bool1":true,
				"t_string1":"ss101",
				"t_bytes1":"YmIxMDE=",
				"t_enum1":1,
				"t_message1":{"f_string1":"ms101","f_string2":"ms102","f_string3":"ms103"},
				"t_any1":{"type_url":"type.googleapis.com/external.Message1","value":"CgRzMTAxEgRzMTAyGgRzMTAz"},
				"t_duration1":{"seconds":3600},
				"t_timestamp1":{"seconds":1686416585}
			}
		*/
		t.Run("int32", func(t *testing.T) {
			require.Equal(t, float64(111), data["t_int32"].(float64))
		})
		t.Run("int64", func(t *testing.T) {
			require.Equal(t, float64(121), data["t_int64"].(float64))
		})
		t.Run("uint32", func(t *testing.T) {
			require.Equal(t, float64(131), data["t_uint32"].(float64))
		})
		t.Run("uint64", func(t *testing.T) {
			require.Equal(t, float64(141), data["t_uint64"].(float64))
		})
		t.Run("sint32", func(t *testing.T) {
			require.Equal(t, float64(151), data["t_sint32"].(float64))
		})
		t.Run("sint64", func(t *testing.T) {
			require.Equal(t, float64(161), data["t_sint64"].(float64))
		})
		t.Run("sfixed32", func(t *testing.T) {
			require.Equal(t, float64(171), data["t_sfixed32"].(float64))
		})
		t.Run("sfixed64", func(t *testing.T) {
			require.Equal(t, float64(181), data["t_sfixed64"].(float64))
		})
		t.Run("fixed32", func(t *testing.T) {
			require.Equal(t, float64(191), data["t_fixed32"].(float64))
		})
		t.Run("fixed64", func(t *testing.T) {
			require.Equal(t, float64(201), data["t_fixed64"].(float64))
		})
		t.Run("float", func(t *testing.T) {
			require.Equal(t, float64(211.111), data["t_float"].(float64))
		})
		t.Run("double", func(t *testing.T) {
			require.Equal(t, float64(221.111), data["t_double"].(float64))
		})
		t.Run("bool1", func(t *testing.T) {
			require.Equal(t, true, data["t_bool1"].(bool))
		})
		t.Run("string1", func(t *testing.T) {
			require.Equal(t, "ss101", data["t_string1"].(string))
		})
		t.Run("bytes1", func(t *testing.T) {
			require.Equal(t, "YmIxMDE=", data["t_bytes1"].(string))
		})
		t.Run("enum1", func(t *testing.T) {
			require.Equal(t, float64(1), data["t_enum1"].(float64))
		})
		t.Run("message1", func(t *testing.T) {
			tMessage1 := data["t_message1"].(map[string]interface{})
			require.Equal(t, "ms101", tMessage1["f_string1"].(string))
			require.Equal(t, "ms102", tMessage1["f_string2"].(string))
			require.Equal(t, "ms103", tMessage1["f_string3"].(string))
		})
		t.Run("any1", func(t *testing.T) {
			tAny1 := data["t_any1"].(map[string]interface{})
			require.Equal(t, "type.googleapis.com/external.Message1", tAny1["type_url"].(string))
			require.Equal(t, "CgRzMTAxEgRzMTAyGgRzMTAz", tAny1["value"].(string))
		})
		t.Run("duration1", func(t *testing.T) {
			tDuration1 := data["t_duration1"].(map[string]interface{})
			require.Equal(t, float64(3600), tDuration1["seconds"].(float64))
		})
		t.Run("timestamp1", func(t *testing.T) {
			tTimestamp1 := data["t_timestamp1"].(map[string]interface{})
			require.Equal(t, float64(1686416585), tTimestamp1["seconds"].(float64))
		})
	})
}

func Test_Omitempty_TypePlain1(t *testing.T) {
	seed1 := &pboptions.TypePlain1{
		FInt32:      0,
		FInt64:      0,
		FUint32:     0,
		FUint64:     0,
		FSint32:     0,
		FSint64:     0,
		FSfixed32:   0,
		FSfixed64:   0,
		FFixed32:    0,
		FFixed64:    0,
		FFloat:      0,
		FDouble:     0,
		FBool1:      false,
		FString1:    "",
		FBytes1:     nil,
		FEnum1:      0,
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
		dataNew := &pboptions.TypePlain1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		require.Equal(t, "{}", string(bb))
	})
}
