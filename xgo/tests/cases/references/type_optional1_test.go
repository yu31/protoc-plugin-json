package references

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"

	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

var seedTypeOptional1 = &pbref.TypeOptional1{
	FInt32A:    utils.PointerInt32(1111),
	FInt32B:    utils.PointerInt32(1112),
	FInt64A:    utils.PointerInt64(1211),
	FInt64B:    utils.PointerInt64(1212),
	FUint32A:   utils.PointerUint32(1311),
	FUint32B:   utils.PointerUint32(1312),
	FUint64A:   utils.PointerUint64(1411),
	FUint64B:   utils.PointerUint64(1412),
	FSint32A:   utils.PointerInt32(1511),
	FSint32B:   utils.PointerInt32(1512),
	FSint64A:   utils.PointerInt64(1611),
	FSint64B:   utils.PointerInt64(1612),
	FSfixed32A: utils.PointerInt32(1711),
	FSfixed32B: utils.PointerInt32(1712),
	FSfixed64A: utils.PointerInt64(1811),
	FSfixed64B: utils.PointerInt64(1812),
	FFixed32A:  utils.PointerUint32(1911),
	FFixed32B:  utils.PointerUint32(1912),
	FFixed64A:  utils.PointerUint64(2011),
	FFixed64B:  utils.PointerUint64(2012),
	FFloat1:    utils.PointerFloat32(2111.111),
	FFloat2:    utils.PointerFloat32(2111.112),
	FDouble1:   utils.PointerFloat64(2211.111),
	FDouble2:   utils.PointerFloat64(2211.112),
	FBool1:     utils.PointerBool(true),
	FBool2:     utils.PointerBool(false),
}

func Test_Reference_Type_Optional1_General(t *testing.T) {
	bb, err := seedTypeOptional1.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbref.TypeOptional1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedTypeOptional1, dataNew)
}

func Test_Reference_Type_Optional1_Check(t *testing.T) {
	bb, err := seedTypeOptional1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("int32", func(t *testing.T) {
		/*
			"f_int32a":1111,
			"f_int32b":"1112",
		*/
		require.Equal(t, float64(1111), data["f_int32a"].(float64))
		require.Equal(t, "1112", data["f_int32b"].(string))
	})
	t.Run("int64", func(t *testing.T) {
		/*
			"f_int64a":1211,
			"f_int64b":"1212",
		*/
		require.Equal(t, float64(1211), data["f_int64a"].(float64))
		require.Equal(t, "1212", data["f_int64b"].(string))
	})
	t.Run("uint32", func(t *testing.T) {
		/*
			"f_uint32a":1311,
			"f_uint32b":"1312",
		*/
		require.Equal(t, float64(1311), data["f_uint32a"].(float64))
		require.Equal(t, "1312", data["f_uint32b"].(string))
	})
	t.Run("uint64", func(t *testing.T) {
		/*
			"f_uint64a":1411,
			"f_uint64b":"1412",
		*/
		require.Equal(t, float64(1411), data["f_uint64a"].(float64))
		require.Equal(t, "1412", data["f_uint64b"].(string))
	})
	t.Run("sint32", func(t *testing.T) {
		/*
			"f_sint32a":1511,
			"f_sint32b":"1512",
		*/
		require.Equal(t, float64(1511), data["f_sint32a"].(float64))
		require.Equal(t, "1512", data["f_sint32b"].(string))
	})
	t.Run("sint64", func(t *testing.T) {
		/*
			"f_sint64a":1611,
			"f_sint64b":"1612",
		*/
		require.Equal(t, float64(1611), data["f_sint64a"].(float64))
		require.Equal(t, "1612", data["f_sint64b"].(string))
	})
	t.Run("sfixed32", func(t *testing.T) {
		/*
			"f_sfixed32a":1711,
			"f_sfixed32b":"1712",
		*/
		require.Equal(t, float64(1711), data["f_sfixed32a"].(float64))
		require.Equal(t, "1712", data["f_sfixed32b"].(string))
	})
	t.Run("sfixed64", func(t *testing.T) {
		/*
			"f_sfixed64a":1811,
			"f_sfixed64b":"1812",
		*/
		require.Equal(t, float64(1811), data["f_sfixed64a"].(float64))
		require.Equal(t, "1812", data["f_sfixed64b"].(string))
	})
	t.Run("fixed32", func(t *testing.T) {
		/*
			"f_fixed32a":1911,
			"f_fixed32b":"1912",
		*/
		require.Equal(t, float64(1911), data["f_fixed32a"].(float64))
		require.Equal(t, "1912", data["f_fixed32b"].(string))
	})
	t.Run("fixed64", func(t *testing.T) {
		/*
			"f_fixed64a":2011,
			"f_fixed64b":"2012",
		*/
		require.Equal(t, float64(2011), data["f_fixed64a"].(float64))
		require.Equal(t, "2012", data["f_fixed64b"].(string))
	})
	t.Run("float", func(t *testing.T) {
		/*
			"f_float1":2111.111,
			"f_float2":"2111.112",
		*/
		require.Equal(t, float64(2111.111), data["f_float1"].(float64))
		require.Equal(t, "2111.112", data["f_float2"].(string))
	})
	t.Run("double", func(t *testing.T) {
		/*
			"f_double1":2211.111,
			"f_double2":"2211.112",
		*/
		require.Equal(t, float64(2211.111), data["f_double1"].(float64))
		require.Equal(t, "2211.112", data["f_double2"].(string))
	})
	t.Run("bool", func(t *testing.T) {
		/*
			"f_bool1":true,
			"f_bool2":"false"
		*/
		require.Equal(t, true, data["f_bool1"].(bool))
		require.Equal(t, "false", data["f_bool2"].(string))
	})
}
