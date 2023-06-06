package boundary

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

type RepeatedElem1 pbboundary.RepeatedElem1

func Test_RepeatedElem1_StdJSON(t *testing.T) {
	var err error

	// 1. The field is not empty and the content is null in JSON.
	t.Run("case1", func(t *testing.T) {
		bb := []byte(`       {     "r_string1"   :       null     }     `)
		data := &RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		err = json.Unmarshal(bb, data)
		require.Nil(t, err)
		require.Nil(t, data.RString1)
	})
	// 2. The field is not empty and the content is empty in JSON.
	t.Run("case2", func(t *testing.T) {
		bb := []byte(`       {    "r_string1"    : [      ]    }   `)
		data := &RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		p1 := reflect.ValueOf(data.RString1).Pointer()
		c1 := cap(data.RString1)

		err = json.Unmarshal(bb, data)
		require.Nil(t, err)

		require.NotNil(t, data.RString1)
		require.Equal(t, 0, len(data.RString1))

		p2 := reflect.ValueOf(data.RString1).Pointer()
		require.NotEqual(t, p1, p2)

		c2 := cap(data.RString1)
		require.NotEqual(t, c1, c2)
	})
	// 3. The length of field is less than the content in JSON.
	t.Run("case3", func(t *testing.T) {
		bb := []byte(`    {    "r_string1"        :  ["4",  "5",  "6",  "7"  ]    }  `)
		data := &RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		p1 := reflect.ValueOf(data.RString1).Pointer()

		err = json.Unmarshal(bb, data)
		require.Nil(t, err)

		require.NotNil(t, data.RString1)
		require.Equal(t, 4, len(data.RString1))

		p2 := reflect.ValueOf(data.RString1).Pointer()
		require.NotEqual(t, p1, p2)

		require.Equal(t, []string{"4", "5", "6", "7"}, data.RString1)
	})
	// 4. The length of field is greater than the content in JSON.
	t.Run("case4", func(t *testing.T) {
		bb := []byte(`    {    "r_string1":  ["4",  "5"   ]    }  `)
		data := &RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		p1 := reflect.ValueOf(data.RString1).Pointer()
		c1 := cap(data.RString1)

		err = json.Unmarshal(bb, data)
		require.Nil(t, err)

		require.NotNil(t, data.RString1)
		require.Equal(t, 2, len(data.RString1))

		p2 := reflect.ValueOf(data.RString1).Pointer()
		require.Equal(t, p1, p2)

		c2 := cap(data.RString1)
		require.Equal(t, c1, c2)

		require.Equal(t, []string{"4", "5"}, data.RString1)
	})
}

func Test_RepeatedElem1_Plugin(t *testing.T) {
	var err error

	// 1. The field is not empty and the content is null in JSON.
	t.Run("case1", func(t *testing.T) {
		bb := []byte(`       {     "r_string1"   :       null     }     `)
		data := &pbboundary.RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Nil(t, data.RString1)
	})
	// 2. The field is not empty and the content is empty in JSON.
	t.Run("case2", func(t *testing.T) {
		bb := []byte(`       {    "r_string1"    : [      ]    }   `)
		data := &pbboundary.RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		p1 := reflect.ValueOf(data.RString1).Pointer()
		c1 := cap(data.RString1)

		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.NotNil(t, data.RString1)
		require.Equal(t, 0, len(data.RString1))

		// Is different with standard json.
		p2 := reflect.ValueOf(data.RString1).Pointer()
		require.Equal(t, p1, p2)

		// Is different with standard json.
		c2 := cap(data.RString1)
		require.Equal(t, c1, c2)
	})
	//3. The length of field is less than the content in JSON.
	t.Run("case3", func(t *testing.T) {
		bb := []byte(`    {    "r_string1"     :  ["4",  "5",  "6",  "7"  ]    }  `)
		data := &pbboundary.RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		p1 := reflect.ValueOf(data.RString1).Pointer()

		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.NotNil(t, data.RString1)
		require.Equal(t, 4, len(data.RString1))

		p2 := reflect.ValueOf(data.RString1).Pointer()
		require.NotEqual(t, p1, p2)

		require.Equal(t, []string{"4", "5", "6", "7"}, data.RString1)
	})
	//4. The length of field is greater than the content in JSON.
	t.Run("case4", func(t *testing.T) {
		bb := []byte(`    {    "r_string1":  ["4",  "5"   ]    }  `)
		data := &pbboundary.RepeatedElem1{
			RString1: []string{"1", "2", "3"},
		}
		p1 := reflect.ValueOf(data.RString1).Pointer()
		c1 := cap(data.RString1)

		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.NotNil(t, data.RString1)
		require.Equal(t, 2, len(data.RString1))

		p2 := reflect.ValueOf(data.RString1).Pointer()
		require.Equal(t, p1, p2)

		c2 := cap(data.RString1)
		require.Equal(t, c1, c2)

		require.Equal(t, []string{"4", "5"}, data.RString1)
	})
}
