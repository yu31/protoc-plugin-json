package base

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbbase"
)

type CopyMap3 pbbase.TypeMap3

var seedMap3 = &pbbase.TypeMap3{
	FString1:  map[bool]string{true: "kv11", false: "kv12"},
	FInt32:    map[bool]int32{true: 1011, false: 1012},
	FInt64:    map[bool]int64{true: 1111, false: 1112},
	FUint32:   map[bool]uint32{true: 1211, false: 1212},
	FUint64:   map[bool]uint64{true: 1311, false: 1312},
	FSint32:   map[bool]int32{true: 1411, false: 1412},
	FSint64:   map[bool]int64{true: 1511, false: 1512},
	FSfixed32: map[bool]int32{true: 1611, false: 1612},
	FSfixed64: map[bool]int64{true: 1711, false: 1712},
	FFixed32:  map[bool]uint32{true: 1811, false: 1812},
	FFixed64:  map[bool]uint64{true: 1911, false: 1912},
	FBool:     map[bool]bool{true: true, false: false},
}

func Test_TypeMap3_Assert_Type(t *testing.T) {
	data := &pbbase.TypeMap3{}
	_, ok1 := interface{}(data).(json.Marshaler)
	require.True(t, ok1)
	_, ok2 := interface{}(data).(json.Unmarshaler)
	require.True(t, ok2)
}

func Test_TypeMap3_Assert_Copy(t *testing.T) {
	dataCopy := &CopyMap3{}
	_, ok1 := interface{}(dataCopy).(json.Marshaler)
	require.False(t, ok1)
	_, ok2 := interface{}(dataCopy).(json.Unmarshaler)
	require.False(t, ok2)
}

func Test_TypeMap3_Basic(t *testing.T) {
	var (
		err error
		b1  []byte
	)
	t.Run("Marshal", func(t *testing.T) {
		b1, err = seedMap3.MarshalJSON()
		require.Nil(t, err)
	})

	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pbbase.TypeMap3{}
		require.NotEqual(t, seedMap3, dataNew)
		err = dataNew.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, seedMap3, dataNew)
	})
}

// Test cases the field is empty.
func Test_TypeMap3_Empty(t *testing.T) {
	dataEmtpy := &pbbase.TypeMap3{}
	dataCopy := &CopyMap3{}

	var (
		err error
		b1  []byte
		b2  []byte
	)

	t.Run("Marshal", func(t *testing.T) {
		b1, err = dataEmtpy.MarshalJSON()
		require.Nil(t, err)
		b2, err = json.Marshal(dataCopy)
		require.Nil(t, err)
	})

	t.Run("UnmarshalPlugin", func(t *testing.T) {
		// use content that get by MarshalJSON
		data1 := &pbbase.TypeMap3{}
		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, dataEmtpy, data1)

		// use content that get by json.Marshal
		data2 := &pbbase.TypeMap3{}
		err = data2.UnmarshalJSON(b2)
		require.Nil(t, err)
		require.Equal(t, dataEmtpy, data2)
	})
	t.Run("UnmarshalStandard", func(t *testing.T) {
		// use content that get by MarshalJSON
		data1 := &CopyMap3{}
		err = json.Unmarshal(b1, data1)
		require.Nil(t, err)
		require.Equal(t, dataCopy, data1)

		// use content that get by json.Marshal
		data2 := &CopyMap3{}
		err = json.Unmarshal(b2, data2)
		require.Nil(t, err)
		require.Equal(t, dataCopy, data2)
	})
}

func Test_TypeMap3_NULL(t *testing.T) {
	data1 := &pbbase.TypeMap3{}
	data2 := &CopyMap3{}

	var err error
	err = data1.UnmarshalJSON([]byte("null"))
	require.Nil(t, err)

	err = json.Unmarshal([]byte("null"), data2)
	require.Nil(t, err)
}
