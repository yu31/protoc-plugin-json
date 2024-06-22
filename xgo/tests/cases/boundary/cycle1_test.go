package boundary

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

func Test_InlineMessageCycle1(t *testing.T) {
	seed1 := &pbboundary.InlineMessageCycle1{
		FString1: "s101",
		FString2: "s102",
		FMessage1: &pbboundary.InlineMessageCycle1_EmbedMessage1{
			FString3: "s103",
			FString4: "s104",
			FMessage2: &pbboundary.InlineMessageCycle1{
				FString1:  "s105",
				FString2:  "s106",
				FMessage1: nil,
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
		dataNew := &pbboundary.InlineMessageCycle1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})
}
