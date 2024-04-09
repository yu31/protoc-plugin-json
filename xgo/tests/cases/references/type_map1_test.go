package references

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

var seedTypeMap1 = &pbref.TypeMap1{
	FInt32A:    map[int32]int32{1111: 2111, 1112: 2112, 1113: 2113},
	FInt32B:    map[int32]int32{1121: 2121, 1122: 2122, 1123: 2123},
	FInt64A:    map[int64]int64{1211: 2211, 1212: 2212, 1213: 2213},
	FInt64B:    map[int64]int64{1221: 2221, 1222: 2222, 1223: 2223},
	FUint32A:   map[uint32]uint32{1311: 2311, 1312: 2312, 1313: 2313},
	FUint32B:   map[uint32]uint32{1321: 2321, 1322: 2322, 1323: 2323},
	FUint64A:   map[uint64]uint64{1411: 2411, 1412: 2412, 1413: 2413},
	FUint64B:   map[uint64]uint64{1421: 2421, 1422: 2422, 1423: 2423},
	FSint32A:   map[int32]int32{1511: 2511, 1512: 2512, 1513: 2513},
	FSint32B:   map[int32]int32{1521: 2521, 1522: 2522, 1523: 2523},
	FSint64A:   map[int64]int64{1611: 2611, 1612: 2612, 1613: 2613},
	FSint64B:   map[int64]int64{1621: 2621, 1622: 2622, 1623: 2623},
	FSfixed32A: map[int32]int32{1711: 2711, 1712: 2712, 1713: 2713},
	FSfixed32B: map[int32]int32{1721: 2721, 1722: 2722, 1723: 2723},
	FSfixed64A: map[int64]int64{1811: 2811, 1812: 2812, 1813: 2813},
	FSfixed64B: map[int64]int64{1821: 2821, 1822: 2822, 1823: 2823},
	FFixed32A:  map[uint32]uint32{1911: 2911, 1912: 2912, 1913: 2913},
	FFixed32B:  map[uint32]uint32{1921: 2921, 1922: 2922, 1923: 2923},
	FFixed64A:  map[uint64]uint64{2011: 3011, 2012: 3012, 2013: 3013},
	FFixed64B:  map[uint64]uint64{2021: 3021, 2022: 3022, 2023: 3023},
	FFloat1:    map[string]float32{"2111": 2111.111, "2112": 2112.111, "2113": 2113.111},
	FFloat2:    map[string]float32{"2121": 2121.111, "2122": 2122.111, "2123": 2123.111},
	FDouble1:   map[string]float64{"2211": 2211.111, "2212": 2212.111, "2213": 2213.111},
	FDouble2:   map[string]float64{"2221": 2221.111, "2222": 2222.111, "2223": 2223.111},
}

func Test_Reference_Type_Map1_General(t *testing.T) {
	bb, err := seedTypeMap1.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbref.TypeMap1{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, seedTypeMap1, dataNew)
}

func Test_Reference_Type_Map1_Check(t *testing.T) {
	bb, err := seedTypeMap1.MarshalJSON()
	require.Nil(t, err)

	data := map[string]interface{}{}
	err = json.Unmarshal(bb, &data)
	require.Nil(t, err)

	t.Run("int32", func(t *testing.T) {
		/*
			"f_int32a":{"1111":2111,"1112":2112,"1113":2113},
			"f_int32b":{"1122":"2122","1123":"2123","1121":"2121"},
		*/
		a := data["f_int32a"].(map[string]interface{})
		require.Equal(t, float64(2111), a["1111"].(float64))
		require.Equal(t, float64(2112), a["1112"].(float64))
		require.Equal(t, float64(2113), a["1113"].(float64))

		b := data["f_int32b"].(map[string]interface{})
		require.Equal(t, "2121", b["1121"].(string))
		require.Equal(t, "2122", b["1122"].(string))
		require.Equal(t, "2123", b["1123"].(string))
	})
	t.Run("int64", func(t *testing.T) {
		/*
			"f_int64a":{"1212":2212,"1213":2213,"1211":2211},
			"f_int64b":{"1223":"2223","1221":"2221","1222":"2222"},
		*/
		a := data["f_int64a"].(map[string]interface{})
		require.Equal(t, float64(2211), a["1211"].(float64))
		require.Equal(t, float64(2212), a["1212"].(float64))
		require.Equal(t, float64(2213), a["1213"].(float64))

		b := data["f_int64b"].(map[string]interface{})
		require.Equal(t, "2221", b["1221"].(string))
		require.Equal(t, "2222", b["1222"].(string))
		require.Equal(t, "2223", b["1223"].(string))
	})
	t.Run("uint32", func(t *testing.T) {
		/*
			"f_uint32a":{"1311":2311,"1312":2312,"1313":2313},
			"f_uint32b":{"1323":"2323","1321":"2321","1322":"2322"},
		*/
		a := data["f_uint32a"].(map[string]interface{})
		require.Equal(t, float64(2311), a["1311"].(float64))
		require.Equal(t, float64(2312), a["1312"].(float64))
		require.Equal(t, float64(2313), a["1313"].(float64))

		b := data["f_uint32b"].(map[string]interface{})
		require.Equal(t, "2321", b["1321"].(string))
		require.Equal(t, "2322", b["1322"].(string))
		require.Equal(t, "2323", b["1323"].(string))
	})
	t.Run("uint64", func(t *testing.T) {
		/*
			"f_uint64a":{"1411":2411,"1412":2412,"1413":2413},
			"f_uint64b":{"1422":"2422","1423":"2423","1421":"2421"},
		*/
		a := data["f_uint64a"].(map[string]interface{})
		require.Equal(t, float64(2411), a["1411"].(float64))
		require.Equal(t, float64(2412), a["1412"].(float64))
		require.Equal(t, float64(2413), a["1413"].(float64))

		b := data["f_uint64b"].(map[string]interface{})
		require.Equal(t, "2421", b["1421"].(string))
		require.Equal(t, "2422", b["1422"].(string))
		require.Equal(t, "2423", b["1423"].(string))
	})
	t.Run("sint32", func(t *testing.T) {
		/*
			"f_sint32a":{"1511":2511,"1512":2512,"1513":2513},
			"f_sint32b":{"1521":"2521","1522":"2522","1523":"2523"},
		*/
		a := data["f_sint32a"].(map[string]interface{})
		require.Equal(t, float64(2511), a["1511"].(float64))
		require.Equal(t, float64(2512), a["1512"].(float64))
		require.Equal(t, float64(2513), a["1513"].(float64))

		b := data["f_sint32b"].(map[string]interface{})
		require.Equal(t, "2521", b["1521"].(string))
		require.Equal(t, "2522", b["1522"].(string))
		require.Equal(t, "2523", b["1523"].(string))
	})
	t.Run("sint64", func(t *testing.T) {
		/*
			"f_sint64a":{"1613":2613,"1611":2611,"1612":2612},
			"f_sint64b":{"1621":"2621","1622":"2622","1623":"2623"},
		*/
		a := data["f_sint64a"].(map[string]interface{})
		require.Equal(t, float64(2611), a["1611"].(float64))
		require.Equal(t, float64(2612), a["1612"].(float64))
		require.Equal(t, float64(2613), a["1613"].(float64))

		b := data["f_sint64b"].(map[string]interface{})
		require.Equal(t, "2621", b["1621"].(string))
		require.Equal(t, "2622", b["1622"].(string))
		require.Equal(t, "2623", b["1623"].(string))
	})
	t.Run("sfixed32", func(t *testing.T) {
		/*
			"f_sfixed32a":{"1711":2711,"1712":2712,"1713":2713},
			"f_sfixed32b":{"1721":"2721","1722":"2722","1723":"2723"},
		*/
		a := data["f_sfixed32a"].(map[string]interface{})
		require.Equal(t, float64(2711), a["1711"].(float64))
		require.Equal(t, float64(2712), a["1712"].(float64))
		require.Equal(t, float64(2713), a["1713"].(float64))

		b := data["f_sfixed32b"].(map[string]interface{})
		require.Equal(t, "2721", b["1721"].(string))
		require.Equal(t, "2722", b["1722"].(string))
		require.Equal(t, "2723", b["1723"].(string))
	})
	t.Run("sfixed64", func(t *testing.T) {
		/*
			"f_sfixed64a":{"1811":2811,"1812":2812,"1813":2813},
			"f_sfixed64b":{"1821":"2821","1822":"2822","1823":"2823"},
		*/
		a := data["f_sfixed64a"].(map[string]interface{})
		require.Equal(t, float64(2811), a["1811"].(float64))
		require.Equal(t, float64(2812), a["1812"].(float64))
		require.Equal(t, float64(2813), a["1813"].(float64))

		b := data["f_sfixed64b"].(map[string]interface{})
		require.Equal(t, "2821", b["1821"].(string))
		require.Equal(t, "2822", b["1822"].(string))
		require.Equal(t, "2823", b["1823"].(string))
	})
	t.Run("fixed32", func(t *testing.T) {
		/*
			"f_fixed32a":{"1911":2911,"1912":2912,"1913":2913},
			"f_fixed32b":{"1921":"2921","1922":"2922","1923":"2923"},
		*/
		a := data["f_fixed32a"].(map[string]interface{})
		require.Equal(t, float64(2911), a["1911"].(float64))
		require.Equal(t, float64(2912), a["1912"].(float64))
		require.Equal(t, float64(2913), a["1913"].(float64))

		b := data["f_fixed32b"].(map[string]interface{})
		require.Equal(t, "2921", b["1921"].(string))
		require.Equal(t, "2922", b["1922"].(string))
		require.Equal(t, "2923", b["1923"].(string))
	})
	t.Run("fixed64", func(t *testing.T) {
		/*
			"f_fixed64a":{"2011":3011,"2012":3012,"2013":3013},
			"f_fixed64b":{"2023":"3023","2021":"3021","2022":"3022"},
		*/
		a := data["f_fixed64a"].(map[string]interface{})
		require.Equal(t, float64(3011), a["2011"].(float64))
		require.Equal(t, float64(3012), a["2012"].(float64))
		require.Equal(t, float64(3013), a["2013"].(float64))

		b := data["f_fixed64b"].(map[string]interface{})
		require.Equal(t, "3021", b["2021"].(string))
		require.Equal(t, "3022", b["2022"].(string))
		require.Equal(t, "3023", b["2023"].(string))
	})
	t.Run("float", func(t *testing.T) {
		/*

			"f_float1":{"2113":2113.111,"2111":2111.111,"2112":2112.111},
			"f_float2":{"2121":"2121.111","2122":"2122.111","2123":"2123.111"},
		*/
		a := data["f_float1"].(map[string]interface{})
		require.Equal(t, float64(2111.111), a["2111"].(float64))
		require.Equal(t, float64(2112.111), a["2112"].(float64))
		require.Equal(t, float64(2113.111), a["2113"].(float64))

		b := data["f_float2"].(map[string]interface{})
		require.Equal(t, "2121.111", b["2121"].(string))
		require.Equal(t, "2122.111", b["2122"].(string))
		require.Equal(t, "2123.111", b["2123"].(string))
	})
	t.Run("double", func(t *testing.T) {
		/*
			"f_double1":{"2211":2211.111,"2212":2212.111,"2213":2213.111},
			"f_double2":{"2221":"2221.111","2222":"2222.111","2223":"2223.111"}
		*/
		a := data["f_double1"].(map[string]interface{})
		require.Equal(t, float64(2211.111), a["2211"].(float64))
		require.Equal(t, float64(2212.111), a["2212"].(float64))
		require.Equal(t, float64(2213.111), a["2213"].(float64))

		b := data["f_double2"].(map[string]interface{})
		require.Equal(t, "2221.111", b["2221"].(string))
		require.Equal(t, "2222.111", b["2222"].(string))
		require.Equal(t, "2223.111", b["2223"].(string))
	})
}
