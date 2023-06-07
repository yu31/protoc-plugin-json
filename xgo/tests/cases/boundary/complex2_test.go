package boundary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

//type Complex2 pbboundary.Complex2

var bComplex2 = []byte(
	`
	{ 
		"f_string" : "s101" ,   
	"r_int64": [ 11  , 12   , 13 ]   ,     
	"level1"     :  {
		"level2":      {
			"f_string": "s102"   ,    
			"level3"      : {
				"f_int32":21 ,
				"r_int64":[31, 32, 33   ]  ,
				"m_int32": { "k1"    : 41, "k2"   : 42    , "k3"   :43}  ,      
				"p_int64":      51   
			}   ,
			"r_int64":     [    61, 62, 63]
		},
		"f_string": "s103"       ,
		"r_string": [  "s201"    , "s202"    , "s203"    ]       
	},
	"r_level1": null, 
	"r_level2":[
		{},
		{"f_string": "es103"},
		{"f_string": "es203"}
	]
}
`,
)

func checkComplex2(t *testing.T, data *pbboundary.Complex2) {
	require.Equal(t, "s101", data.FString)
	require.Equal(t, []int64{11, 12, 13}, data.RInt64)

	level1 := data.Level1
	require.NotNil(t, level1)
	require.Equal(t, "s103", level1.FString)
	require.Equal(t, []string{"s201", "s202", "s203"}, level1.RString)

	level2 := level1.Level2
	require.NotNil(t, level2)
	require.Equal(t, "s102", level2.FString)
	require.Equal(t, []int64{61, 62, 63}, level2.RInt64)

	level3 := level2.Level3
	require.NotNil(t, level3)
	require.Equal(t, int32(21), level3.FInt32)
	require.Equal(t, []int64{31, 32, 33}, level3.RInt64)
	require.Equal(t, map[string]int32{"k1": 41, "k2": 42, "k3": 43}, level3.MInt32)
	require.NotNil(t, level3.PInt64)
	require.Equal(t, int64(51), *level3.PInt64)

	rLevel1 := data.RLevel1
	require.Nil(t, rLevel1)

	rLevel2 := data.RLevel2
	require.Equal(t, 3, len(rLevel2))
	for _, rl := range rLevel2 {
		require.NotNil(t, rl)
		require.Nil(t, rl.Level2)
	}
	require.Equal(t, "", rLevel2[0].FString)
	require.Equal(t, "es103", rLevel2[1].FString)
	require.Equal(t, "es203", rLevel2[2].FString)
}

func Test_Complex2_Plugin1(t *testing.T) {
	data := &pbboundary.Complex2{
		FString: "",
		RInt64:  nil, // []int64
		Level1: &pbboundary.Complex2_Level1{
			Level2: &pbboundary.Complex2_Level2{
				FString: "",
				Level3: &pbboundary.Complex2_Level3{
					FInt32: 0,   // int32
					RInt64: nil, // []int64
					MInt32: nil, // map[string]int32
					PInt64: nil, // *int64
				},
				RInt64: nil, // []int64
			},
			FString: "",
			RString: nil, // []string
		},
		RLevel1: []*pbboundary.Complex2_Level1{},
		RLevel2: nil,
	}

	err := data.UnmarshalJSON(bComplex2)
	require.Nil(t, err, fmt.Sprintf("%v", err))
	checkComplex2(t, data)
}

func Test_Complex2_Plugin2(t *testing.T) {
	data := &pbboundary.Complex2{
		FString: "",
		RInt64:  nil,
		Level1:  nil,
		RLevel1: nil,
		RLevel2: nil,
	}

	err := data.UnmarshalJSON(bComplex2)
	require.Nil(t, err, fmt.Sprintf("%v", err))
	checkComplex2(t, data)
}
