package boundary

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbboundary"
)

func Test_InlineMessageCircular3(t *testing.T) {
	seed1 := &pbboundary.InlineMessageCircular3{
		FString1: "s101",
		FString2: "s102",
		FMessage1: &pbboundary.InlineMessageCircular3_EmbedMessage1{
			FString3: "s103",
			FString4: "s104",
			FMessage2: &pbboundary.InlineMessageCircular3_EmbedMessage2{
				FString5: "s105",
				FString6: "s106",
				FMessage3: &pbboundary.InlineMessageCircular3_EmbedMessage1{
					FString3:  "s107",
					FString4:  "s108",
					FMessage2: nil,
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
		dataNew := &pbboundary.InlineMessageCircular3{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
		require.Equal(t, seed1, dataNew)
	})
}
