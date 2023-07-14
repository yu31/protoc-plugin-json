package jsondecoder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Read_JSON_IsNULL(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		bb := []byte(` null  }  `)
		decoder, err := New(bb)
		require.Nil(t, err)

		_, err = decoder.BeforeScanJSON()
		require.NotNil(t, err)
	})
	t.Run("case2", func(t *testing.T) {
		bb := []byte(` null  xxxx  `)
		decoder, err := New(bb)
		require.Nil(t, err)

		_, err = decoder.BeforeScanJSON()
		require.NotNil(t, err)
	})
}
