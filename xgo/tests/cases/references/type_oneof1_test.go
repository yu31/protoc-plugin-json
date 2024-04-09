package references

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

var seedTypeOneOf1 = &pbref.TypeOneOf1{
	OneInt32:    &pbref.TypeOneOf1_FInt32A{FInt32A: 1111},
	OneInt64:    &pbref.TypeOneOf1_FInt64A{FInt64A: 1211},
	OneUint32:   &pbref.TypeOneOf1_FUint32A{FUint32A: 1311},
	OneUint64:   &pbref.TypeOneOf1_FUint64A{FUint64A: 1411},
	OneSInt32:   &pbref.TypeOneOf1_FSint32A{FSint32A: 1511},
	OneSInt64:   &pbref.TypeOneOf1_FSint64A{FSint64A: 1611},
	OneSFixed32: &pbref.TypeOneOf1_FSfixed32A{FSfixed32A: 1711},
	OneSFixed64: &pbref.TypeOneOf1_FSfixed64A{FSfixed64A: 1811},
	OneFixed32:  &pbref.TypeOneOf1_FFixed32A{FFixed32A: 1911},
	OneFixed64:  &pbref.TypeOneOf1_FFixed64A{FFixed64A: 2011},
	OneFloat:    &pbref.TypeOneOf1_FFloat1{FFloat1: 2111.111},
	OneDouble:   &pbref.TypeOneOf1_FDouble1{FDouble1: 2211.111},
	OneBool:     &pbref.TypeOneOf1_FBool1{FBool1: true},
}
var seedTypeOneOf2 = &pbref.TypeOneOf1{
	OneInt32:    &pbref.TypeOneOf1_FInt32B{FInt32B: 1112},
	OneInt64:    &pbref.TypeOneOf1_FInt64B{FInt64B: 1212},
	OneUint32:   &pbref.TypeOneOf1_FUint32B{FUint32B: 1312},
	OneUint64:   &pbref.TypeOneOf1_FUint64B{FUint64B: 1412},
	OneSInt32:   &pbref.TypeOneOf1_FSint32B{FSint32B: 1512},
	OneSInt64:   &pbref.TypeOneOf1_FSint64B{FSint64B: 1612},
	OneSFixed32: &pbref.TypeOneOf1_FSfixed32B{FSfixed32B: 1712},
	OneSFixed64: &pbref.TypeOneOf1_FSfixed64B{FSfixed64B: 1812},
	OneFixed32:  &pbref.TypeOneOf1_FFixed32B{FFixed32B: 1912},
	OneFixed64:  &pbref.TypeOneOf1_FFixed64B{FFixed64B: 2012},
	OneFloat:    &pbref.TypeOneOf1_FFloat2{FFloat2: 2112.111},
	OneDouble:   &pbref.TypeOneOf1_FDouble2{FDouble2: 2212.111},
	OneBool:     &pbref.TypeOneOf1_FBool2{FBool2: false},
}

func Test_Reference_Type_OneOf1_General1(t *testing.T) {
	bb, err := seedTypeOneOf1.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbref.TypeOneOf1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedTypeOneOf1, dataNew)
}

func Test_Reference_Type_OneOf1_General2(t *testing.T) {
	bb, err := seedTypeOneOf2.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbref.TypeOneOf1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedTypeOneOf2, dataNew)
}

func Test_Reference_Type_OneOf1_Check1(t *testing.T) {
	bb, err := seedTypeOneOf1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)
	/*
		{
			"f_int32a":1111,
			"one_int64":{"f_int64a":1211},
			"f_uint32a":1311,
			"one_uint64":{"f_uint64a":1411},
			"f_sint32a":1511,
			"one_sint64":{"f_sint64a":1611},
			"f_sfixed32a":1711,
			"one_sfixed64":{"f_sfixed64a":1811},
			"f_fixed32a":1911,
			"one_fixed64":{"f_fixed64a":2011},
			"f_float1":2111.111,
			"one_double":{"f_double1":2211.111},
			"f_bool1":true
		}
	*/
	t.Run("OneInt32", func(t *testing.T) {
		require.Equal(t, float64(1111), data["f_int32a"].(float64))
	})
	t.Run("OneInt64", func(t *testing.T) {
		oneInt64 := data["one_int64"].(map[string]interface{})
		require.Equal(t, float64(1211), oneInt64["f_int64a"].(float64))
	})
	t.Run("OneUint32", func(t *testing.T) {
		require.Equal(t, float64(1311), data["f_uint32a"].(float64))
	})
	t.Run("OneUint64", func(t *testing.T) {
		oneUint64 := data["one_uint64"].(map[string]interface{})
		require.Equal(t, float64(1411), oneUint64["f_uint64a"].(float64))
	})
	t.Run("OneSInt32", func(t *testing.T) {
		require.Equal(t, float64(1511), data["f_sint32a"].(float64))
	})
	t.Run("OneSInt64", func(t *testing.T) {
		oneSInt64 := data["one_sint64"].(map[string]interface{})
		require.Equal(t, float64(1611), oneSInt64["f_sint64a"].(float64))
	})
	t.Run("OneSFixed32", func(t *testing.T) {
		require.Equal(t, float64(1711), data["f_sfixed32a"].(float64))
	})
	t.Run("OneSFixed64", func(t *testing.T) {
		oneSFixed64 := data["one_sfixed64"].(map[string]interface{})
		require.Equal(t, float64(1811), oneSFixed64["f_sfixed64a"].(float64))
	})
	t.Run("OneFixed32", func(t *testing.T) {
		require.Equal(t, float64(1911), data["f_fixed32a"].(float64))
	})
	t.Run("OneFixed64", func(t *testing.T) {
		oneFixed64 := data["one_fixed64"].(map[string]interface{})
		require.Equal(t, float64(2011), oneFixed64["f_fixed64a"].(float64))
	})

	t.Run("OneFloat", func(t *testing.T) {
		require.Equal(t, float64(2111.111), data["f_float1"].(float64))
	})
	t.Run("OneDouble", func(t *testing.T) {
		oneDouble := data["one_double"].(map[string]interface{})
		require.Equal(t, float64(2211.111), oneDouble["f_double1"].(float64))
	})

	t.Run("OneBool", func(t *testing.T) {
		require.Equal(t, true, data["f_bool1"].(bool))
	})
}

func Test_Reference_Type_OneOf1_Check2(t *testing.T) {
	bb, err := seedTypeOneOf2.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)
	/*
		{
			"f_int32b":"1112",
			"one_int64":{"f_int64b":"1212"},
			"f_uint32b":"1312",
			"one_uint64":{"f_uint64b":"1412"},
			"f_sint32b":"1512",
			"one_sint64":{"f_sint64b":"1612"},
			"f_sfixed32b":"1712",
			"one_sfixed64":{"f_sfixed64b":"1812"},
			"f_fixed32b":"1912",
			"one_fixed64":{"f_fixed64b":"2012"},
			"f_float2":"2112.111",
			"one_double":{"f_double2":"2212.111"},
			"f_bool2":"false"
		}
	*/
	t.Run("OneInt32", func(t *testing.T) {
		require.Equal(t, "1112", data["f_int32b"].(string))
	})
	t.Run("OneInt64", func(t *testing.T) {
		oneInt64 := data["one_int64"].(map[string]interface{})
		require.Equal(t, "1212", oneInt64["f_int64b"].(string))
	})
	t.Run("OneUint32", func(t *testing.T) {
		require.Equal(t, "1312", data["f_uint32b"].(string))
	})
	t.Run("OneUint64", func(t *testing.T) {
		oneUint64 := data["one_uint64"].(map[string]interface{})
		require.Equal(t, "1412", oneUint64["f_uint64b"].(string))
	})
	t.Run("OneSInt32", func(t *testing.T) {
		require.Equal(t, "1512", data["f_sint32b"].(string))
	})
	t.Run("OneSInt64", func(t *testing.T) {
		oneSInt64 := data["one_sint64"].(map[string]interface{})
		require.Equal(t, "1612", oneSInt64["f_sint64b"].(string))
	})
	t.Run("OneSFixed32", func(t *testing.T) {
		require.Equal(t, "1712", data["f_sfixed32b"].(string))
	})
	t.Run("OneSFixed64", func(t *testing.T) {
		oneSFixed64 := data["one_sfixed64"].(map[string]interface{})
		require.Equal(t, "1812", oneSFixed64["f_sfixed64b"].(string))
	})
	t.Run("OneFixed32", func(t *testing.T) {
		require.Equal(t, "1912", data["f_fixed32b"].(string))
	})
	t.Run("OneFixed64", func(t *testing.T) {
		oneFixed64 := data["one_fixed64"].(map[string]interface{})
		require.Equal(t, "2012", oneFixed64["f_fixed64b"].(string))
	})

	t.Run("OneFloat", func(t *testing.T) {
		require.Equal(t, "2112.111", data["f_float2"].(string))
	})
	t.Run("OneDouble", func(t *testing.T) {
		oneDouble := data["one_double"].(map[string]interface{})
		require.Equal(t, "2212.111", oneDouble["f_double2"].(string))
	})

	t.Run("OneBool", func(t *testing.T) {
		require.Equal(t, "false", data["f_bool2"].(string))
	})
}
