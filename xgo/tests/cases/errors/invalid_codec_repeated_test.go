package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pberrors"
)

func Test_InvalidCodecRepeated(t *testing.T) {
	seed1 := &pberrors.InvalidCodecRepeated{}

	t.Run("Int32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_int32_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int32_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Int64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_int64_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int64_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_uint32_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint32_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_uint64_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint64_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sint32_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint32_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sint64_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint64_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Float", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_float_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_float_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Double", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_double_numeric": ["1111"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_double_string": [1112] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("Bool", func(t *testing.T) {
			bb := []byte(`{ "f_bool_bool": "true" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_bool_string": true }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": [sssss] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": [false] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": [true] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": [11111] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": [0.1111] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bytes", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": [Ynl0ZXM=] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": ["Ynl0ZXM"] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": [11111] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": [true] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": [false] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case6", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": [0.111] }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
