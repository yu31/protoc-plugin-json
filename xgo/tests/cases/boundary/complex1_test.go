package boundary

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

type Complex1 pbboundary.Complex1

var bComplex1 = []byte(
	`{ 
		"f_int32"     : 11      , 
		"unknown1": "unknown1",
		"r_int64"     : [ 21, 22, 23 ]     ,    
		"unknown2":"unknown2"         ,
		"f_message1"      : { "f_string1": "m11", "f_string2": "m12", "f_string3": "m13" }      ,   
		"unknown3"		:		"unknown3",     
		"m_int32"      : { "k31": 31, "k32": 32, "k33": 33 }    ,        
		"unknown4"	:"unknown4"			,     
		"f_int64": 51      ,    
		"unknown5": "unknown5"						
	}`,
)

func Test_Complex1_StdJSON(t *testing.T) {
	data := &Complex1{
		FInt32:    0,
		RInt64:    nil,
		FMessage1: nil,
		MInt32:    nil,
		FInt64:    nil,
	}
	err := json.Unmarshal(bComplex1, data)
	require.Nil(t, err, fmt.Sprintf("%v", err))

	require.Equal(t, int32(11), data.FInt32)
	require.Equal(t, []int64{21, 22, 23}, data.RInt64)

	require.NotNil(t, data.FMessage1)
	require.Equal(t, "m11", data.FMessage1.FString1)
	require.Equal(t, "m12", data.FMessage1.FString2)
	require.Equal(t, "m13", data.FMessage1.FString3)

	require.Equal(t, map[string]int32{"k31": 31, "k32": 32, "k33": 33}, data.MInt32)

	require.NotNil(t, data.FInt64)
	require.Equal(t, int64(51), *data.FInt64)
}

func Test_Complex1_Plugin(t *testing.T) {
	data := &pbboundary.Complex1{
		FInt32:    0,
		RInt64:    nil,
		FMessage1: nil,
		MInt32:    nil,
		FInt64:    nil,
	}
	err := json.Unmarshal(bComplex1, data)
	require.Nil(t, err, fmt.Sprintf("%v", err))

	require.Equal(t, int32(11), data.FInt32)
	require.Equal(t, []int64{21, 22, 23}, data.RInt64)

	require.NotNil(t, data.FMessage1)
	require.Equal(t, "m11", data.FMessage1.FString1)
	require.Equal(t, "m12", data.FMessage1.FString2)
	require.Equal(t, "m13", data.FMessage1.FString3)

	require.Equal(t, map[string]int32{"k31": 31, "k32": 32, "k33": 33}, data.MInt32)

	require.NotNil(t, data.FInt64)
	require.Equal(t, int64(51), *data.FInt64)
}
