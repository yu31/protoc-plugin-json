package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pberrors"
)

func Test_InvalidCodecMapKey(t *testing.T) {
	seed1 := &pberrors.InvalidCodecMapKey{}

	t.Run("Int32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_int32_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int32_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Int64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_int64_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int64_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_uint32_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint32_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_uint64_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint64_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sint32_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint32_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sint64_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint64_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64_key_numeric_val_numeric": { "1111": 1211 } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64_key_string_val_string": { 1112: "1212" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			bb := []byte(`{ "f_bool_key_bool_val_bool": { "true": true } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_bool_key_string_val_string": { true: "true" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("LossDoubleQuotationMarkBegin", func(t *testing.T) {
			bb := []byte(`{ "f_string_key_none_val_none1": { 1111": "v1" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("LossDoubleQuotationMarkEnd", func(t *testing.T) {
			bb := []byte(`{ "f_string_key_none_val_none1": { "1111: "v1" } }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
