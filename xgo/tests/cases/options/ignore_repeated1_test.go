package options

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pboptions"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_IgnoreRepeated1_1(t *testing.T) {
	seed1 := &pboptions.IgnoreRepeated1{
		FInt32:    []int32{111, 112, 113},
		FInt64:    []int64{121, 122, 123},
		FUint32:   []uint32{131, 132, 133},
		FUint64:   []uint64{141, 142, 143},
		FSint32:   []int32{151, 152, 153},
		FSint64:   []int64{161, 162, 163},
		FSfixed32: []int32{171, 172, 173},
		FSfixed64: []int64{181, 182, 183},
		FFixed32:  []uint32{191, 192, 193},
		FFixed64:  []uint64{201, 202, 203},
		FFloat:    []float32{211.111, 212.111, 213.111},
		FDouble:   []float64{221.111, 222.111, 223.111},
		FBool1:    []bool{true, false},
		FString1:  []string{"ss101", "ss102", "ss103"},
		FBytes1:   [][]byte{[]byte("bb101"), []byte("bb102"), []byte("bb103")},
		FEnum1:    []pbexternal.Enum1{1, 3, 5},
		FMessage1: []*pbexternal.Message1{
			{FString1: "ms101", FString2: "ms102", FString3: "ms103"},
			{FString1: "ms201", FString2: "ms202", FString3: "ms203"},
			{FString1: "ms301", FString2: "ms302", FString3: "ms303"},
		},
		FAny1: []*anypb.Any{
			utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s201", FString2: "s202", FString3: "s203"}),
			utils.MustNewAny(&pbexternal.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
		},
		FDuration1: []*durationpb.Duration{
			{Seconds: 1800, Nanos: 0},
			{Seconds: 3600, Nanos: 0},
			{Seconds: 7200, Nanos: 0},
		},
		FTimestamp1: []*timestamppb.Timestamp{
			{Seconds: 1686416585, Nanos: 0},
			{Seconds: 1686416585, Nanos: 0},
			{Seconds: 1686416585, Nanos: 0},
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
		dataNew := &pboptions.TypePlain1{}
		err = dataNew.UnmarshalJSON(bb)
		require.Nil(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		require.Equal(t, "{}", string(bb))
	})
}

func Test_IgnoreRepeated1_2(t *testing.T) {
	seed1 := &pboptions.IgnoreRepeated1{
		FInt32:      nil,
		FInt64:      nil,
		FUint32:     nil,
		FUint64:     nil,
		FSint32:     nil,
		FSint64:     nil,
		FSfixed32:   nil,
		FSfixed64:   nil,
		FFixed32:    nil,
		FFixed64:    nil,
		FFloat:      nil,
		FDouble:     nil,
		FBool1:      nil,
		FString1:    nil,
		FBytes1:     nil,
		FEnum1:      nil,
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
