package jsonencoder

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("L0", func(t *testing.T) {
		bufLen := 0
		enc := New(bufLen)
		require.NotNil(t, enc)
		require.Equal(t, 0, len(enc.buf))
		require.Equal(t, bufLen, cap(enc.buf))
	})
	t.Run("L8", func(t *testing.T) {
		bufLen := 8
		enc := New(bufLen)
		require.NotNil(t, enc)
		require.Equal(t, 0, len(enc.buf))
		require.Equal(t, bufLen, cap(enc.buf))
	})
}

func TestEncoder_Build1(t *testing.T) {
	enc := New(0)
	enc.AppendObjectBegin()
	enc.AppendJSONKey("sk1")
	enc.AppendValueString("sv1")
	enc.AppendJSONKey("ik64")
	enc.AppendValueInt64(64)
	enc.AppendJSONKey("ikb")
	enc.AppendValueBool(true)
	enc.AppendObjectEnd()

	b := enc.Bytes()
	require.NotNil(t, b)

	fmt.Println(string(b))

	t.Run("map", func(t *testing.T) {
		m := make(map[string]interface{})
		err := json.Unmarshal(b, &m)
		require.Nil(t, err)

		require.Equal(t, "sv1", m["sk1"])
		require.Equal(t, float64(64), m["ik64"])
		require.Equal(t, true, m["ikb"])
	})

	t.Run("struct", func(t *testing.T) {
		tt := struct {
			SK1  string `json:"sk1"`
			IK64 int64  `json:"ik64"`
			IKB  bool   `json:"ikb"`
		}{}

		err := json.Unmarshal(b, &tt)
		require.Nil(t, err)
		require.Equal(t, "sv1", tt.SK1)
		require.Equal(t, int64(64), tt.IK64)
		require.Equal(t, true, tt.IKB)
	})

}