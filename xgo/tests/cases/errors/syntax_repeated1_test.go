package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

func Test_Syntax_TypeRepeated1(t *testing.T) {
	data := &pbref.TypeRepeated1{}

	t.Run("InvalidArrayBegin", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": 1111 ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": null ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{"f_int32a": nv ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("InvalidArrayEnd", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111, }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 ] } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{"f_int32a": null } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("InvalidArraySeparator", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 ; 2222 ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 | 2222 ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111, 2222, ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 2222 ] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
