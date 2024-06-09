package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pberrors"
)

func Test_InvalidSyntaxMap(t *testing.T) {
	seed1 := &pberrors.InvalidSyntaxMap{}

	t.Run("MapBegin", func(t *testing.T) {
		t.Run("LossMapBegin1", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": "11111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossMapBegin2", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": nv }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossMapBegin3", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": "11111": 22222 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("MapEnd", func(t *testing.T) {
		t.Run("LossMapColonAndValue1", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": { "11111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossMapColonAndValue2", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": { "11111", }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossJSONEnd1", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": { "11111": 22222 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossJSONEnd2", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": null `)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONEnd1", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": null } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONEnd2", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": { "11111": 22222 } } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("MapSeparator", func(t *testing.T) {
		t.Run("InvalidColon", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": { "11111"; 22222 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossMapColon", func(t *testing.T) {
			bb := []byte(`{"f_int32_key_string_val_numeric": { "11111" 22222 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
