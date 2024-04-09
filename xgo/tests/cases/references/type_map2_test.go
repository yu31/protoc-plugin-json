package references

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

var seedTypeMap2 = &pbref.TypeMap2{
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
	FSfixed64A: map[int64]int64{1811: 2811, 1812: 2812, 1813: 2812},
	FSfixed64B: map[int64]int64{1821: 2821, 1822: 2822, 1823: 2823},
	FFixed32A:  map[uint32]uint32{1911: 2911, 1912: 2912, 1913: 2913},
	FFixed32B:  map[uint32]uint32{1921: 2921, 1922: 2922, 1923: 2923},
	FFixed64A:  map[uint64]uint64{2011: 3011, 2012: 3012, 2013: 3013},
	FFixed64B:  map[uint64]uint64{2021: 3021, 2022: 3022, 2023: 3023},
	FFloat1:    map[string]float32{"2111": 2111.111, "2112": 2112.111, "2113": 2113.111},
	FFloat2:    map[string]float32{"2121": 2121.111, "2122": 2122.111, "2123": 2123.111},
	FDouble1:   map[string]float64{"2211": 2211.111, "2212": 2212.111, "2213": 2213.111},
	FDouble2:   map[string]float64{"2221": 2221.111, "2222": 2222.111, "2223": 2223.111},
	FBool1:     map[bool]bool{true: false, false: true},
	FBool2:     map[bool]bool{false: true, true: false},
}

func Test_Reference_Type_Map2_General(t *testing.T) {
	bb, err := seedTypeMap2.MarshalJSON()
	require.Nil(t, err)

	dataNew := &pbref.TypeMap2{}
	err = dataNew.UnmarshalJSON(bb)
	require.Nil(t, err, fmt.Sprintf("%v", err))

	require.Equal(t, seedTypeMap2, dataNew)
}
