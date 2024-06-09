package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pberrors"
)

func Test_InvalidSyntaxRepeated(t *testing.T) {
	seed1 := &pberrors.InvalidSyntaxRepeated{}

	t.Run("ArrayBegin", func(t *testing.T) {
		t.Run("LossArrayBegin1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": 1111 ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossArrayBegin2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": nv ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusArrayEnd", func(t *testing.T) {
			bb := []byte(`{"f_int32a": null ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("ArrayEnd", func(t *testing.T) {
		t.Run("LossArrayEnd1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossArrayEnd2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111, }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONEnd1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 ] } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONEnd2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": null } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("ArraySeparator", func(t *testing.T) {
		t.Run("InvalidArraySeparator1", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 ; 2222 ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("InvalidArraySeparator2", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 | 2222 ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("InvalidArraySeparator3", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111 2222 ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusArraySeparator", func(t *testing.T) {
			bb := []byte(`{"f_int32a": [ 1111, 2222, ] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
