package boundary

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

func Test_InlineMessageCycle2(t *testing.T) {
	seed1 := &pbboundary.InlineMessageCycle2{
		FString1: "s101",
		FString2: "s102",
		FMessage1: &pbboundary.InlineMessageCycle2_EmbedMessage1{
			FString3: "s103",
			FString4: "s104",
			FMessage2: &pbboundary.InlineMessageCycle2_EmbedMessage2{
				FString5: "s105",
				FString6: "s106",
				FMessage3: &pbboundary.InlineMessageCycle2{
					FString1:  "s107",
					FString2:  "s108",
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
		dataNew := &pbboundary.InlineMessageCycle2{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})
}
