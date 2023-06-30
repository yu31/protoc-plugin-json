package options

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_IgnorePlain1_1(t *testing.T) {
	seed1 := &pboptions.IgnorePlain1{
		FInt32:    111,
		FInt64:    121,
		FUint32:   131,
		FUint64:   141,
		FSint32:   151,
		FSint64:   161,
		FSfixed32: 171,
		FSfixed64: 181,
		FFixed32:  191,
		FFixed64:  201,
		FFloat:    211.111,
		FDouble:   221.111,
		FBool1:    true,
		FString1:  "ss101",
		FBytes1:   []byte("bb101"),
		FEnum1:    1,
		FMessage1: &pbexternal.Message1{
			FString1: "ms101",
			FString2: "ms102",
			FString3: "ms103",
		},
		FAny1:       utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		FDuration1:  &durationpb.Duration{Seconds: 3600, Nanos: 0},
		FTimestamp1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
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
		dataNew := &pboptions.TypePlain1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		require.Equal(t, "{}", string(bb))
	})
}

func Test_IgnorePlain1_2(t *testing.T) {
	seed1 := &pboptions.IgnorePlain1{
		FInt32:      0,
		FInt64:      0,
		FUint32:     0,
		FUint64:     0,
		FSint32:     0,
		FSint64:     0,
		FSfixed32:   0,
		FSfixed64:   0,
		FFixed32:    0,
		FFixed64:    0,
		FFloat:      0,
		FDouble:     0,
		FBool1:      false,
		FString1:    "",
		FBytes1:     nil,
		FEnum1:      0,
		FMessage1:   nil,
		FAny1:       nil,
		FDuration1:  nil,
		FTimestamp1: nil,
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
		dataNew := &pboptions.TypePlain1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		require.Equal(t, "{}", string(bb))
	})
}
