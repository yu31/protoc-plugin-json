package bases

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbbases"

	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

type CopyMap2 pbbases.TypeMap2

var seedMap2 = &pbbases.TypeMap2{
	FString1:  map[string]string{"ks11": "kv11"},
	FInt32:    map[int32]int32{1010: 1011},
	FInt64:    map[int64]int64{1110: 1111},
	FUint32:   map[uint32]uint32{1210: 1211},
	FUint64:   map[uint64]uint64{1310: 1311},
	FSint32:   map[int32]int32{1410: 1411},
	FSint64:   map[int64]int64{1510: 1511},
	FSfixed32: map[int32]int32{1610: 1611},
	FSfixed64: map[int64]int64{1710: 1711},
	FFixed32:  map[uint32]uint32{1810: 1811},
	FFixed64:  map[uint64]uint64{1910: 1911},
	//FBool1:    map[bool]bool{true: true},
}

func Test_TypeMap2_Assert_Type(t *testing.T) {
	data := &pbbases.TypeMap2{}
	_, ok1 := interface{}(data).(json.Marshaler)
	require.True(t, ok1)
	_, ok2 := interface{}(data).(json.Unmarshaler)
	require.True(t, ok2)
}

func Test_TypeMap2_Assert_Copy(t *testing.T) {
	dataCopy := &CopyMap2{}
	_, ok1 := interface{}(dataCopy).(json.Marshaler)
	require.False(t, ok1)
	_, ok2 := interface{}(dataCopy).(json.Unmarshaler)
	require.False(t, ok2)
}

func Test_TypeMap2_General1(t *testing.T) {
	var (
		err error
		b1  []byte
	)
	t.Run("marshal", func(t *testing.T) {
		b1, err = seedMap2.MarshalJSON()
		require.Nil(t, err)
	})
	fmt.Println(string(b1))
	t.Run("unmarshal", func(t *testing.T) {
		dataNew := &pbbases.TypeMap2{}
		require.NotEqual(t, seedMap2, dataNew)
		err = dataNew.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, seedMap2, dataNew)
	})
}

// Test compatible with standard JSON with follow cases:
//   1. The json content that get by MarshalJSON that code generated by plugin can be used to standard json.Unmarshal.
//   2. The json content that get by standard json.Marshal can be used to UnmarshalJSON that code generated by plugin.
func Test_TypeMap2_Compatible(t *testing.T) {
	var (
		err error
		b1  []byte
		b2  []byte
		b3  []byte
	)
	_ = b3

	// Get json content by MarshalJSON.
	b1, err = seedMap2.MarshalJSON()
	require.Nil(t, err)

	dataCopy := &CopyMap2{}
	// Test the json content is correctly.
	// and same as copy the value to dataCopy.
	err = json.Unmarshal(b1, dataCopy)
	require.Nil(t, err)

	//Get json context by json.Marshal
	b2, err = json.Marshal(dataCopy)
	require.Nil(t, err)

	// Compare the json content.
	require.Equal(t, b1, b2)

	// Test the UnmarshalJSON
	dataNew := &pbbases.TypeMap2{}
	err = dataNew.UnmarshalJSON(b2)
	require.Nil(t, err)
	require.Equal(t, seedMap2, dataNew)

	// proto
	b3, err = utils.PMarshal.Marshal(seedMap2)
	require.Nil(t, err)

	fmt.Println("=========================================== PLUGIN =============================================")
	fmt.Println(string(b1))
	fmt.Println("=========================================== STANDARD ===========================================")
	fmt.Println(string(b2))
	fmt.Println("=========================================== PROTO ==============================================")
	fmt.Println(string(b3))
	fmt.Println("================================================================================================")
}

// Test cases the field is empty.
func Test_TypeMap2_Empty(t *testing.T) {
	dataEmtpy := &pbbases.TypeMap2{}
	dataCopy := &CopyMap2{}

	var (
		err error
		b1  []byte
		b2  []byte
	)

	t.Run("marshal", func(t *testing.T) {
		b1, err = dataEmtpy.MarshalJSON()
		require.Nil(t, err)
		b2, err = json.Marshal(dataCopy)
		require.Nil(t, err)
	})

	t.Run("unmarshal-plugin", func(t *testing.T) {
		// use content that get by MarshalJSON
		data1 := &pbbases.TypeMap2{}
		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, dataEmtpy, data1)

		// use content that get by json.Marshal
		data2 := &pbbases.TypeMap2{}
		err = data2.UnmarshalJSON(b2)
		require.Nil(t, err)
		require.Equal(t, dataEmtpy, data2)
	})
	t.Run("unmarshal-standard", func(t *testing.T) {
		// use content that get by MarshalJSON
		data1 := &CopyMap2{}
		err = json.Unmarshal(b1, data1)
		require.Nil(t, err)
		require.Equal(t, dataCopy, data1)

		// use content that get by json.Marshal
		data2 := &CopyMap2{}
		err = json.Unmarshal(b2, data2)
		require.Nil(t, err)
		require.Equal(t, dataCopy, data2)
	})
}

func Test_TypeMap2_NULL(t *testing.T) {
	data1 := &pbbases.TypeMap2{}
	data2 := &CopyMap2{}

	var err error
	err = data1.UnmarshalJSON([]byte("null"))
	require.Nil(t, err)

	err = json.Unmarshal([]byte("null"), data2)
	require.Nil(t, err)
}