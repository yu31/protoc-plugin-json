package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pberrors"
)

func Test_InvalidSyntaxJSON(t *testing.T) {
	seed1 := &pberrors.InvalidSyntaxJSON{}

	t.Run("JSONBegin", func(t *testing.T) {
		t.Run("LossJSONStartBrackets1", func(t *testing.T) {
			bb := []byte(`"f_int32a": 1111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("InvalidLocationOfFieldSeparator", func(t *testing.T) {
			bb := []byte(`{ , "f_int32a": 1111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONBegin", func(t *testing.T) {
			bb := []byte(`{ { "f_int32a": 1111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossJSONKey", func(t *testing.T) {
			bb := []byte(`{ null `)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("JSONEnd", func(t *testing.T) {
		t.Run("LossJSONEnd1", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossJSONEnd2", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111, `)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusFieldSeparator", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111 , }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONEnd1", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SurplusJSONEnd2", func(t *testing.T) {
			bb := []byte(` null }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("FieldSeparator", func(t *testing.T) {
		t.Run("InvalidFieldSeparator", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": 1111; "f_int32b": "1112" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossFieldColon", func(t *testing.T) {
			bb := []byte(`{ "f_int32a"; 1111, "f_int32b": "1112" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("JSONKey", func(t *testing.T) {
		t.Run("LossDoubleQuotationMarkBegin1", func(t *testing.T) {
			bb := []byte(`{ f_int32a": 1111 , }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossDoubleQuotationMarkBegin2", func(t *testing.T) {
			bb := []byte(`{ t_int32a": 1111 , }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossDoubleQuotationMarkBegin3", func(t *testing.T) {
			bb := []byte(`{ t_int32a: 1111 , }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
