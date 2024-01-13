package options

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
)

// Test cases for oneof multiplayer nesting
func Test_InlineOneOf2(t *testing.T) {
	data1 := &pboptions.InlineOneOf2{
		FString1: "v_string0",
		OneType01: &pboptions.InlineOneOf2_FMessage1{
			FMessage1: &pboptions.InlineOneOf2Message1{
				FString1: "v_string1",
				OneType01: &pboptions.InlineOneOf2Message1_FMessage1{
					FMessage1: &pboptions.InlineOneOf2Message2{
						FString1: "v_string2",
						OneType01: &pboptions.InlineOneOf2Message2_FMessage1{
							FMessage1: &pboptions.InlineOneOf2Message3{
								FString1: "v_string3",
								OneType01: &pboptions.InlineOneOf2Message3_FDouble{
									FDouble: 11.11,
								},
							},
						},
					},
				},
			},
		},
	}

	var (
		err error
		bb  []byte
	)

	t.Run("Marshal", func(t *testing.T) {
		bb, err = data1.MarshalJSON()
		require.Nil(t, err)
		fmt.Println(string(bb))
	})

	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pboptions.InlineOneOf2{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		data2 := &pboptions.InlineOneOf2{}
		err = data2.UnmarshalJSON(bb)
		require.Nil(t, err)

		require.Equal(t, "v_string0", data2.FString1)

		msg1, ok := data2.OneType01.(*pboptions.InlineOneOf2_FMessage1)
		require.True(t, ok)
		require.NotNil(t, msg1)
		require.NotNil(t, msg1.FMessage1)
		require.Equal(t, "v_string1", msg1.FMessage1.FString1)

		msg2, ok := msg1.FMessage1.OneType01.(*pboptions.InlineOneOf2Message1_FMessage1)
		require.True(t, ok)
		require.NotNil(t, msg2)
		require.NotNil(t, msg2.FMessage1)
		require.Equal(t, "v_string2", msg2.FMessage1.FString1)

		msg3, ok := msg2.FMessage1.OneType01.(*pboptions.InlineOneOf2Message2_FMessage1)
		require.True(t, ok)
		require.NotNil(t, msg3)
		require.NotNil(t, msg3.FMessage1)
		require.Equal(t, "v_string3", msg3.FMessage1.FString1)

		msg4, ok := msg3.FMessage1.OneType01.(*pboptions.InlineOneOf2Message3_FDouble)
		require.True(t, ok)
		require.NotNil(t, msg4)
		require.Equal(t, float64(11.11), msg4.FDouble)
	})
}
