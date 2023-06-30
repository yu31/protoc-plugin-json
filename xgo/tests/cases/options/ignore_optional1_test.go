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

func Test_IgnoreOptional1_1(t *testing.T) {
	seed1 := &pboptions.IgnoreOptional1{
		FInt32:    utils.PointerInt32(111),
		FInt64:    utils.PointerInt64(121),
		FUint32:   utils.PointerUint32(131),
		FUint64:   utils.PointerUint64(141),
		FSint32:   utils.PointerInt32(151),
		FSint64:   utils.PointerInt64(161),
		FSfixed32: utils.PointerInt32(171),
		FSfixed64: utils.PointerInt64(181),
		FFixed32:  utils.PointerUint32(191),
		FFixed64:  utils.PointerUint64(201),
		FFloat:    utils.PointerFloat32(211.111),
		FDouble:   utils.PointerFloat64(221.111),
		FBool1:    utils.PointerBool(true),
		FString1:  utils.PointerString("ss101"),
		FBytes1:   []byte("bb101"),
		FEnum1:    pbexternal.Enum1(1).Enum(),
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

func Test_IgnoreOptional1_2(t *testing.T) {
	seed1 := &pboptions.IgnoreOptional1{
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
