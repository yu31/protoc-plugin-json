package errors

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbref"
)

func Test_Error_WKTPlain1(t *testing.T) {
	data := &pbref.WKTPlain1{}

	t.Run("Enum", func(t *testing.T) {
		t.Run("Number", func(t *testing.T) {
			bb := []byte(`{ "f_enum_number1": "1" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("NumberString", func(t *testing.T) {
			bb := []byte(`{ "f_enum_number_string1": 1 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("String", func(t *testing.T) {
			bb := []byte(`{ "f_enum_string1": Two }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Any", func(t *testing.T) {
		t.Run("Object", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_any_native1": "111" }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_any_native1": } }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
		t.Run("Proto", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_any_proto1": "111" }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_any_proto1": } }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
	})

	t.Run("Duration", func(t *testing.T) {
		t.Run("Object", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_duration_native1": "111" }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_duration_native1": { }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
		t.Run("String", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_duration_string1": "111" }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_duration_string1": 111 }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("3", func(t *testing.T) {
				bb := []byte(`{ "f_duration_string1": true }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})

		t.Run("Nanosecond", func(t *testing.T) {
			bb := []byte(`{ "f_duration_nanosecond1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("NanosecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_nanosecond_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Microsecond", func(t *testing.T) {
			bb := []byte(`{ "f_duration_microsecond1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("MicrosecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_microsecond_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Millisecond", func(t *testing.T) {
			bb := []byte(`{ "f_duration_millisecond1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("MillisecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_millisecond_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Second", func(t *testing.T) {
			bb := []byte(`{ "f_duration_second1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("SecondString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_second_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Minute", func(t *testing.T) {
			bb := []byte(`{ "f_duration_minute1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("MinuteString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_minute_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("Hour", func(t *testing.T) {
			bb := []byte(`{ "f_duration_hour1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("HourString", func(t *testing.T) {
			bb := []byte(`{ "f_duration_hour_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})

	t.Run("Timestamp", func(t *testing.T) {
		t.Run("Object", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_native1": "111" }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_native1": { }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})
		t.Run("TimeLayout", func(t *testing.T) {
			t.Run("1", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_time_layout1": "111" }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
			t.Run("2", func(t *testing.T) {
				bb := []byte(`{ "f_timestamp_time_layout1": 111 }`)
				err := data.UnmarshalJSON(bb)
				require.NotNil(t, err)
			})
		})

		t.Run("UnixNano", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_nano1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixNanoString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_nano_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("UnixMicro", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_micro1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixMicroString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_micro_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("UnixMilli", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_milli1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixMilliString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_milli_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})

		t.Run("UnixSec", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_sec1": "111" }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
		t.Run("UnixSecString", func(t *testing.T) {
			bb := []byte(`{ "f_timestamp_unix_sec_string1": 111 }`)
			err := data.UnmarshalJSON(bb)
			require.NotNil(t, err)
		})
	})
}
