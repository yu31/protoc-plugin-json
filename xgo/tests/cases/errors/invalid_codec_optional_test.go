package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pberrors"
)

func Test_InvalidCodecOptional(t *testing.T) {
	seed1 := &pberrors.InvalidCodecOptional{}

	t.Run("Int32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_int32_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int32_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Int64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_int64_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_int64_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_uint32_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint32_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Uint64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_uint64_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_uint64_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sint32_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint32_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SInt64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sint64_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sint64_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed32_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("SFixed64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_sfixed64_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed32", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed32_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Fixed64", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_fixed64_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Float", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_float_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_float_string": 1112 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Double", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_double_numeric": "1111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_double_string": 1112 }`)
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
			bb := []byte(`{ "f_string_none": sssss }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": false }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": true }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": 11111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			bb := []byte(`{ "f_string_none": 0.1111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Bytes", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": Ynl0ZXM= }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case2", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": "Ynl0ZXM" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case3", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": 11111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case4", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": true }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case5", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": false }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("case6", func(t *testing.T) {
			// correct: "bytes"("Ynl0ZXM=")
			bb := []byte(`{ "f_bytes_none": 0.111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Enum", func(t *testing.T) {
		t.Run("Numeric", func(t *testing.T) {
			bb := []byte(`{ "f_enum_numeric": "1" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("NumericString", func(t *testing.T) {
			bb := []byte(`{ "f_enum_numeric_string": 1 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_enum_string": Two }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Any", func(t *testing.T) {
		t.Run("Object", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_any_native": "111" }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_any_native": } }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
		t.Run("Proto", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_any_proto": "111" }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_any_proto": } }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
	})

	t.Run("Duration", func(t *testing.T) {
		t.Run("Object", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_duration_native": "111" }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_duration_native": { }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
		t.Run("String", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_duration_string": "111" }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_duration_string": 111 }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("3", func(t *testing.T) {
				bb := []byte(`{ "f_duration_string": true }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})

		t.Run("Nanosecond", func(t *testing.T) {
			bb := []byte(`{ "f_duration_nanosecond": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("NanosecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_nanosecond_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Microsecond", func(t *testing.T) {
			bb := []byte(`{ "f_duration_microsecond": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("MicrosecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_microsecond_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Millisecond", func(t *testing.T) {
			bb := []byte(`{ "f_duration_millisecond": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("MillisecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_millisecond_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Second", func(t *testing.T) {
			bb := []byte(`{ "f_duration_second": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_second_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Minute", func(t *testing.T) {
			bb := []byte(`{ "f_duration_minute": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("MinuteString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_minute_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Hour", func(t *testing.T) {
			bb := []byte(`{ "f_duration_hour": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("HourString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_hour_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Timestamp", func(t *testing.T) {
		t.Run("Object", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_native": "111" }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_native": { }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
		t.Run("TimeLayout", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_time_layout": "111" }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_time_layout": 111 }`)
				err := seed1.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})

		t.Run("UnixNano", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_nano": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixNanoString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_nano_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("UnixMicro", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_micro": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixMicroString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_micro_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("UnixMilli", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_milli": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixMilliString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_milli_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("UnixSec", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_sec": "111" }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixSecString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_sec_string": 111 }`)
			err := seed1.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
