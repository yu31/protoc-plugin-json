package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
)

func Test_Error_TypeMap2Key(t *testing.T) {
	data := &pbformat.TypeMap2{}

	t.Run("Int32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int32b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Int64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_int64a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int64b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_uint32a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint32b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_uint64a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint64b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sint32a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint32b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sint64a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint64b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64a": { "1111": 1211 } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64b": { 1112: "1212" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			bb := []byte(`{ "f_bool1": { "true": true } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_bool2": { true: "true" } }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
