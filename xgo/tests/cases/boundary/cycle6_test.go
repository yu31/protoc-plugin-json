package boundary

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

func Test_InlineMessageCycle6(t *testing.T) {
	seed1 := &pbboundary.InlineMessageCycle6{
		FString1: "s101",
		FString2: "s102",
		FMessage1: &pbboundary.InlineMessageCycle6_EmbedMessage1{
			FString3: "s103",
			FString4: "s104",
			OneType2: &pbboundary.InlineMessageCycle6_EmbedMessage1_FMessage2{
				FMessage2: &pbboundary.InlineMessageCycle6_EmbedMessage2{
					FString7: "s105",
					FString8: "s106",
					OneType3: &pbboundary.InlineMessageCycle6_EmbedMessage2_FMessage3{
						FMessage3: &pbboundary.InlineMessageCycle6{
							FString1:  "s107",
							FString2:  "s108",
							FMessage1: nil,
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
		bb, err = seed1.MarshalJSON()
		require.Nil(t, err)
	})

	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pbboundary.InlineMessageCycle6{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})
}
