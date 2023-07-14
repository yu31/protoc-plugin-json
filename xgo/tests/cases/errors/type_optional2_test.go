package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

func Test_Error_TypeOptional2(t *testing.T) {
	data := &pbref.TypeOptional2{
		FString1: nil,
		FBytes1:  nil,
	}

	t.Run("String", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{ "f_string1": sssss }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ "f_string1": false }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ "f_string1": true }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{ "f_string1": 11111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			bb := []byte(`{ "f_string1": 0.1111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
	})

	t.Run("Bytes", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": Ynl0ZXM= }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": "Ynl0ZXM" }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": 11111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": true }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": false }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case6", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": 0.111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
	})
}
