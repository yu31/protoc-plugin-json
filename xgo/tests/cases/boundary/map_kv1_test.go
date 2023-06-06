package boundary

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

type MapKV1 pbboundary.MapKV1

func Test_MapKV1_StdJSON(t *testing.T) {
	var err error

	// 1. The field is not empty and the content is null in JSON.
	t.Run("case1", func(t *testing.T) {
		bb := []byte(`       {     "m_string1"   :       null     }     `)
		data := &MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}
		err = json.Unmarshal(bb, data)
		require.Nil(t, err)
		require.Nil(t, data.MString1)
	})
	// 2. The field is not empty and the content is empty in JSON.
	t.Run("case2", func(t *testing.T) {
		bb := []byte(`       {     "m_string1"   :       {      }     }     `)
		data := &MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}
		p1 := reflect.ValueOf(data.MString1).Pointer()
		err = json.Unmarshal(bb, data)
		require.Nil(t, err)

		require.NotNil(t, data.MString1)
		require.Equal(t, 3, len(data.MString1))

		p2 := reflect.ValueOf(data.MString1).Pointer()
		require.Equal(t, p1, p2)
	})
	// 3. The length of field is less than the content in JSON.
	t.Run("case3", func(t *testing.T) {
		bb := []byte(`    {    "m_string1"        :  {"k1":  "v11",  "k2"   :  "v22",  "k4"   :"v4"  }    }  `)
		data := &MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}
		err = json.Unmarshal(bb, data)
		require.Nil(t, err)

		require.NotNil(t, data.MString1)
		require.Equal(t, 4, len(data.MString1))

		m1 := map[string]string{"k1": "v11", "k2": "v22", "k3": "v3", "k4": "v4"}
		require.Equal(t, m1, data.MString1)
	})
	// 4. The length of field is greater than the content in JSON.
	t.Run("case4", func(t *testing.T) {
		bb := []byte(`    {    "m_string1":  {"k4": "v4",  "k5": "v5"   }    }  `)
		data := &MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}

		err = json.Unmarshal(bb, data)
		require.Nil(t, err)

		require.NotNil(t, data.MString1)
		require.Equal(t, 5, len(data.MString1))

		m1 := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5"}
		require.Equal(t, m1, data.MString1)
	})
}

func Test_MapKV1_Plugin(t *testing.T) {
	var err error

	// 1. The field is not empty and the content is null in JSON.
	t.Run("case1", func(t *testing.T) {
		bb := []byte(`       {     "m_string1"   :       null     }     `)
		data := &pbboundary.MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}
		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Nil(t, data.MString1)
	})
	// 2. The field is not empty and the content is empty in JSON.
	t.Run("case2", func(t *testing.T) {
		bb := []byte(`       {     "m_string1"   :       {      }     }     `)
		data := &pbboundary.MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}
		p1 := reflect.ValueOf(data.MString1).Pointer()
		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.NotNil(t, data.MString1)
		require.Equal(t, 3, len(data.MString1))

		p2 := reflect.ValueOf(data.MString1).Pointer()
		require.Equal(t, p1, p2)
	})
	// 3. The length of field is less than the content in JSON.
	t.Run("case3", func(t *testing.T) {
		bb := []byte(`    {    "m_string1"        :  {"k1":  "v11",  "k2"   :  "v22",  "k4"   :"v4"  }    }  `)
		data := &pbboundary.MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}
		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.NotNil(t, data.MString1)
		require.Equal(t, 4, len(data.MString1))

		m1 := map[string]string{"k1": "v11", "k2": "v22", "k3": "v3", "k4": "v4"}
		require.Equal(t, m1, data.MString1)
	})
	// 4. The length of field is greater than the content in JSON.
	t.Run("case4", func(t *testing.T) {
		bb := []byte(`    {    "m_string1":  {"k4": "v4",  "k5": "v5"   }    }  `)
		data := &pbboundary.MapKV1{
			MString1: map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"},
		}

		err = data.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.NotNil(t, data.MString1)
		require.Equal(t, 5, len(data.MString1))

		m1 := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5"}
		require.Equal(t, m1, data.MString1)
	})
}
