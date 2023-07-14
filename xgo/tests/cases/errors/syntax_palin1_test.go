package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

func Test_Syntax_TypePlain1(t *testing.T) {
	data := &pbref.TypePlain1{}

	t.Run("InvalidJSONBegin", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`"f_int32a": 1111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ , "f_int32a": 1111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ { "f_int32a": 1111 }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{ null `)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
	})

	t.Run("InvalidJSONEnd", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111, `)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111 , }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111 } }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			bb := []byte(` null }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
	})

	t.Run("InvalidFieldSeparator", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111; "f_int32b": "1112" }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ "f_int32a"; 1111, "f_int32b": "1112" }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
	})

	t.Run("InvalidJSONKey1", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{ f_int32a": 1111 , }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ t_int32a": 1111 , }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ t_int32a: 1111 , }`)
			err := data.UnmarshalJSON(bb)
			fmt.Println(err)
			require.NotNil(t, err)
		})
	})
}
