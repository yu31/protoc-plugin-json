package boundary

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

func Test_InlineMessageCircular4(t *testing.T) {
	seed1 := &pbboundary.InlineMessageCircular4{
		FString1: "s101",
		FString2: "s102",
		FMessage1: &pbboundary.InlineMessageCircular4_EmbedMessage1{
			FString3: "s103",
			FString4: "s104",
			OneType2: &pbboundary.InlineMessageCircular4_EmbedMessage1_FMessage2{
				FMessage2: &pbboundary.InlineMessageCircular4{
					FString1:  "s105",
					FString2:  "s106",
					FMessage1: nil,
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
		dataNew := &pbboundary.InlineMessageCircular4{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})
}
