package format

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
)

var seedTypeRepeated1 = &pbformat.TypeRepeated1{
	FInt32A:    []int32{1111, 1112, 1113},
	FInt32B:    []int32{1121, 1122, 1123},
	FInt64A:    []int64{1211, 1212, 1213},
	FInt64B:    []int64{1221, 1222, 1223},
	FUint32A:   []uint32{1311, 1312, 1313},
	FUint32B:   []uint32{1321, 1322, 1323},
	FUint64A:   []uint64{1411, 1412, 1413},
	FUint64B:   []uint64{1421, 1422, 1423},
	FSint32A:   []int32{1511, 1512, 1513},
	FSint32B:   []int32{1521, 1522, 1523},
	FSint64A:   []int64{1611, 1612, 1613},
	FSint64B:   []int64{1621, 1622, 1623},
	FSfixed32A: []int32{1711, 1712, 1713},
	FSfixed32B: []int32{1721, 1722, 1723},
	FSfixed64A: []int64{1811, 1812, 1813},
	FSfixed64B: []int64{1821, 1822, 1823},
	FFixed32A:  []uint32{1911, 1912, 1913},
	FFixed32B:  []uint32{1921, 1922, 1923},
	FFixed64A:  []uint64{2011, 2012, 2013},
	FFixed64B:  []uint64{2021, 2022, 2023},
	FFloat1:    []float32{2111.111, 2112.111, 2113.111},
	FFloat2:    []float32{2121.111, 2122.111, 2123.111},
	FDouble1:   []float64{2211.111, 2212.111, 2213.111},
	FDouble2:   []float64{2221.111, 2222.111, 2223.111},
	FBool1:     []bool{true, false, true},
	FBool2:     []bool{true, false, true},
}

func Test_Reference_Type_Repeated1_General(t *testing.T) {
	bb, err := seedTypeRepeated1.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbformat.TypeRepeated1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedTypeRepeated1, dataNew)
}

func Test_Reference_Type_Repeated1_Check(t *testing.T) {
	bb, err := seedTypeRepeated1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("int32", func(t *testing.T) {
		/*
			"f_int32a":[1111,1112,1113],
			"f_int32b":["1121","1122","1123"],
		*/
		a := data["f_int32a"].([]interface{})
		require.Equal(t, float64(1111), a[0].(float64))
		require.Equal(t, float64(1112), a[1].(float64))
		require.Equal(t, float64(1113), a[2].(float64))

		b := data["f_int32b"].([]interface{})
		require.Equal(t, "1121", b[0].(string))
		require.Equal(t, "1122", b[1].(string))
		require.Equal(t, "1123", b[2].(string))
	})
	t.Run("int64", func(t *testing.T) {
		/*
			"f_int64a":[1211,1212,1213],
			"f_int64b":["1221","1222","1223"],
		*/
		a := data["f_int64a"].([]interface{})
		require.Equal(t, float64(1211), a[0].(float64))
		require.Equal(t, float64(1212), a[1].(float64))
		require.Equal(t, float64(1213), a[2].(float64))

		b := data["f_int64b"].([]interface{})
		require.Equal(t, "1221", b[0].(string))
		require.Equal(t, "1222", b[1].(string))
		require.Equal(t, "1223", b[2].(string))
	})
	t.Run("uint32", func(t *testing.T) {
		/*
			"f_uint32a":[1311,1312,1313],
			"f_uint32b":["1321","1322","1323"],
		*/
		a := data["f_uint32a"].([]interface{})
		require.Equal(t, float64(1311), a[0].(float64))
		require.Equal(t, float64(1312), a[1].(float64))
		require.Equal(t, float64(1313), a[2].(float64))

		b := data["f_uint32b"].([]interface{})
		require.Equal(t, "1321", b[0].(string))
		require.Equal(t, "1322", b[1].(string))
		require.Equal(t, "1323", b[2].(string))
	})
	t.Run("uint64", func(t *testing.T) {
		/*
			"f_uint64a":[1411,1412,1413],
			"f_uint64b":["1421","1422","1423"],
		*/
		a := data["f_uint64a"].([]interface{})
		require.Equal(t, float64(1411), a[0].(float64))
		require.Equal(t, float64(1412), a[1].(float64))
		require.Equal(t, float64(1413), a[2].(float64))

		b := data["f_uint64b"].([]interface{})
		require.Equal(t, "1421", b[0].(string))
		require.Equal(t, "1422", b[1].(string))
		require.Equal(t, "1423", b[2].(string))
	})
	t.Run("sint32", func(t *testing.T) {
		/*
			"f_sint32a":[1511,1512,1513],
			"f_sint32b":["1521","1522","1523"],
		*/
		a := data["f_sint32a"].([]interface{})
		require.Equal(t, float64(1511), a[0].(float64))
		require.Equal(t, float64(1512), a[1].(float64))
		require.Equal(t, float64(1513), a[2].(float64))

		b := data["f_sint32b"].([]interface{})
		require.Equal(t, "1521", b[0].(string))
		require.Equal(t, "1522", b[1].(string))
		require.Equal(t, "1523", b[2].(string))
	})
	t.Run("sint64", func(t *testing.T) {
		/*
			"f_sint64a":[1611,1612,1613],
			"f_sint64b":["1621","1622","1623"],
		*/

		a := data["f_sint64a"].([]interface{})
		require.Equal(t, float64(1611), a[0].(float64))
		require.Equal(t, float64(1612), a[1].(float64))
		require.Equal(t, float64(1613), a[2].(float64))

		b := data["f_sint64b"].([]interface{})
		require.Equal(t, "1621", b[0].(string))
		require.Equal(t, "1622", b[1].(string))
		require.Equal(t, "1623", b[2].(string))
	})
	t.Run("sfixed32", func(t *testing.T) {
		/*
			"f_sfixed32a":[1711,1712,1713],
			"f_sfixed32b":["1721","1722","1723"],
		*/
		a := data["f_sfixed32a"].([]interface{})
		require.Equal(t, float64(1711), a[0].(float64))
		require.Equal(t, float64(1712), a[1].(float64))
		require.Equal(t, float64(1713), a[2].(float64))

		b := data["f_sfixed32b"].([]interface{})
		require.Equal(t, "1721", b[0].(string))
		require.Equal(t, "1722", b[1].(string))
		require.Equal(t, "1723", b[2].(string))
	})
	t.Run("sfixed64", func(t *testing.T) {
		/*
			"f_sfixed64a":[1811,1812,1813],
			"f_sfixed64b":["1821","1822","1823"],
		*/

		a := data["f_sfixed64a"].([]interface{})
		require.Equal(t, float64(1811), a[0].(float64))
		require.Equal(t, float64(1812), a[1].(float64))
		require.Equal(t, float64(1813), a[2].(float64))

		b := data["f_sfixed64b"].([]interface{})
		require.Equal(t, "1821", b[0].(string))
		require.Equal(t, "1822", b[1].(string))
		require.Equal(t, "1823", b[2].(string))
	})
	t.Run("fixed32", func(t *testing.T) {
		/*
			"f_fixed32a":[1911,1912,1913],
			"f_fixed32b":["1921","1922","1923"],
		*/

		a := data["f_fixed32a"].([]interface{})
		require.Equal(t, float64(1911), a[0].(float64))
		require.Equal(t, float64(1912), a[1].(float64))
		require.Equal(t, float64(1913), a[2].(float64))

		b := data["f_fixed32b"].([]interface{})
		require.Equal(t, "1921", b[0].(string))
		require.Equal(t, "1922", b[1].(string))
		require.Equal(t, "1923", b[2].(string))
	})
	t.Run("fixed64", func(t *testing.T) {
		/*
			"f_fixed64a":[2011,2012,2013],
			"f_fixed64b":["2021","2022","2023"],
		*/
		a := data["f_fixed64a"].([]interface{})
		require.Equal(t, float64(2011), a[0].(float64))
		require.Equal(t, float64(2012), a[1].(float64))
		require.Equal(t, float64(2013), a[2].(float64))

		b := data["f_fixed64b"].([]interface{})
		require.Equal(t, "2021", b[0].(string))
		require.Equal(t, "2022", b[1].(string))
		require.Equal(t, "2023", b[2].(string))
	})
	t.Run("float", func(t *testing.T) {
		/*
			"f_float1":[2111.111,2112.111,2113.111],
			"f_float2":["2121.111","2122.111","2123.111"],
		*/
		a := data["f_float1"].([]interface{})
		require.Equal(t, float64(2111.111), a[0].(float64))
		require.Equal(t, float64(2112.111), a[1].(float64))
		require.Equal(t, float64(2113.111), a[2].(float64))

		b := data["f_float2"].([]interface{})
		require.Equal(t, "2121.111", b[0].(string))
		require.Equal(t, "2122.111", b[1].(string))
		require.Equal(t, "2123.111", b[2].(string))
	})
	t.Run("double", func(t *testing.T) {
		/*
			"f_double1":[2211.111,2212.111,2213.111],
			"f_double2":["2221.111","2222.111","2223.111"],
		*/
		a := data["f_double1"].([]interface{})
		require.Equal(t, float64(2211.111), a[0].(float64))
		require.Equal(t, float64(2212.111), a[1].(float64))
		require.Equal(t, float64(2213.111), a[2].(float64))

		b := data["f_double2"].([]interface{})
		require.Equal(t, "2221.111", b[0].(string))
		require.Equal(t, "2222.111", b[1].(string))
		require.Equal(t, "2223.111", b[2].(string))
	})
	t.Run("bool", func(t *testing.T) {
		/*
			"f_bool1":[true,false,true],
			"f_bool2":["true","false","true"]
		*/
		a := data["f_bool1"].([]interface{})
		require.Equal(t, true, a[0].(bool))
		require.Equal(t, false, a[1].(bool))
		require.Equal(t, true, a[2].(bool))

		b := data["f_bool2"].([]interface{})
		require.Equal(t, "true", b[0].(string))
		require.Equal(t, "false", b[1].(string))
		require.Equal(t, "true", b[2].(string))
	})
}
