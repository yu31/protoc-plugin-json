package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbformat"
)

func Test_Error_TypeRepeated1(t *testing.T) {
	data := &pbformat.TypeRepeated1{}

	t.Run("Int32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_int32a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int32b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Int64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_int64a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int64b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_uint32a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint32b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_uint64a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint64b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sint32a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint32b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sint64a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint64b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed32", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed64", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64a": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64b": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Float", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_float1": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_float2": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Double", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_double1": ["1111"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_double2": [1112] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			bb := []byte(`{ "f_bool1": "true" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_bool2": true }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}

func Test_Error_TypeRepeated2(t *testing.T) {
	data := &pbformat.TypeRepeated2{}

	t.Run("String", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{ "f_string1": [sssss] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ "f_string1": [false] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ "f_string1": [true] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{ "f_string1": [11111] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			bb := []byte(`{ "f_string1": [0.1111] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bytes", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": [Ynl0ZXM=] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": ["Ynl0ZXM"] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": [11111] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": [true] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": [false] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case6", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes1": [0.111] }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
