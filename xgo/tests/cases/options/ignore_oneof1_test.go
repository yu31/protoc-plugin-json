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

func Test_IgnoreOneOf1_1(t *testing.T) {
	seed1 := &pboptions.IgnoreOneOf1{
		OneType1:  &pboptions.IgnoreOneOf1_FInt32{FInt32: 111},
		OneType2:  &pboptions.IgnoreOneOf1_FInt64{FInt64: 121},
		OneType3:  &pboptions.IgnoreOneOf1_FUint32{FUint32: 131},
		OneType4:  &pboptions.IgnoreOneOf1_FUint64{FUint64: 141},
		OneType5:  &pboptions.IgnoreOneOf1_FSint32{FSint32: 151},
		OneType6:  &pboptions.IgnoreOneOf1_FSint64{FSint64: 161},
		OneType7:  &pboptions.IgnoreOneOf1_FSfixed32{FSfixed32: 171},
		OneType8:  &pboptions.IgnoreOneOf1_FSfixed64{FSfixed64: 181},
		OneType9:  &pboptions.IgnoreOneOf1_FFixed32{FFixed32: 191},
		OneType10: &pboptions.IgnoreOneOf1_FFixed64{FFixed64: 201},
		OneType11: &pboptions.IgnoreOneOf1_FFloat{FFloat: 211.111},
		OneType12: &pboptions.IgnoreOneOf1_FDouble{FDouble: 221.111},
		OneType13: &pboptions.IgnoreOneOf1_FBool1{FBool1: true},
		OneType14: &pboptions.IgnoreOneOf1_FString1{FString1: "ss101"},
		OneType15: &pboptions.IgnoreOneOf1_FBytes1{FBytes1: []byte("bb101")},
		OneType16: &pboptions.IgnoreOneOf1_FEnum1{FEnum1: 1},
		OneType17: &pboptions.IgnoreOneOf1_FMessage1{FMessage1: &pbexternal.Message1{
			FString1: "ms101",
			FString2: "ms102",
			FString3: "ms103",
		}},
		OneType18: &pboptions.IgnoreOneOf1_FAny1{
			FAny1: utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		},
		OneType19: &pboptions.IgnoreOneOf1_FDuration1{
			FDuration1: &durationpb.Duration{Seconds: 3600, Nanos: 0},
		},
		OneType20: &pboptions.IgnoreOneOf1_FTimestamp1{
			FTimestamp1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
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

func Test_IgnoreOneOf1_2(t *testing.T) {
	seed1 := &pboptions.IgnoreOneOf1{
		OneType1:  nil,
		OneType2:  nil,
		OneType3:  nil,
		OneType4:  nil,
		OneType5:  nil,
		OneType6:  nil,
		OneType7:  nil,
		OneType8:  nil,
		OneType9:  nil,
		OneType10: nil,
		OneType11: nil,
		OneType12: nil,
		OneType13: nil,
		OneType14: nil,
		OneType15: nil,
		OneType16: nil,
		OneType17: nil,
		OneType18: nil,
		OneType19: nil,
		OneType20: nil,
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

func Test_IgnoreOneOf2_1(t *testing.T) {
	seed1 := &pboptions.IgnoreOneOf2{
		OneType1:  &pboptions.IgnoreOneOf2_FInt32{FInt32: 111},
		OneType2:  &pboptions.IgnoreOneOf2_FInt64{FInt64: 121},
		OneType3:  &pboptions.IgnoreOneOf2_FUint32{FUint32: 131},
		OneType4:  &pboptions.IgnoreOneOf2_FUint64{FUint64: 141},
		OneType5:  &pboptions.IgnoreOneOf2_FSint32{FSint32: 151},
		OneType6:  &pboptions.IgnoreOneOf2_FSint64{FSint64: 161},
		OneType7:  &pboptions.IgnoreOneOf2_FSfixed32{FSfixed32: 171},
		OneType8:  &pboptions.IgnoreOneOf2_FSfixed64{FSfixed64: 181},
		OneType9:  &pboptions.IgnoreOneOf2_FFixed32{FFixed32: 191},
		OneType10: &pboptions.IgnoreOneOf2_FFixed64{FFixed64: 201},
		OneType11: &pboptions.IgnoreOneOf2_FFloat{FFloat: 211.111},
		OneType12: &pboptions.IgnoreOneOf2_FDouble{FDouble: 221.111},
		OneType13: &pboptions.IgnoreOneOf2_FBool1{FBool1: true},
		OneType14: &pboptions.IgnoreOneOf2_FString1{FString1: "ss101"},
		OneType15: &pboptions.IgnoreOneOf2_FBytes1{FBytes1: []byte("bb101")},
		OneType16: &pboptions.IgnoreOneOf2_FEnum1{FEnum1: 1},
		OneType17: &pboptions.IgnoreOneOf2_FMessage1{FMessage1: &pbexternal.Message1{
			FString1: "ms101",
			FString2: "ms102",
			FString3: "ms103",
		}},
		OneType18: &pboptions.IgnoreOneOf2_FAny1{
			FAny1: utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
		},
		OneType19: &pboptions.IgnoreOneOf2_FDuration1{
			FDuration1: &durationpb.Duration{Seconds: 3600, Nanos: 0},
		},
		OneType20: &pboptions.IgnoreOneOf2_FTimestamp1{
			FTimestamp1: &timestamppb.Timestamp{Seconds: 1686416585, Nanos: 0},
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

func Test_IgnoreOneOf2_2(t *testing.T) {
	seed1 := &pboptions.IgnoreOneOf2{
		OneType1:  nil,
		OneType2:  nil,
		OneType3:  nil,
		OneType4:  nil,
		OneType5:  nil,
		OneType6:  nil,
		OneType7:  nil,
		OneType8:  nil,
		OneType9:  nil,
		OneType10: nil,
		OneType11: nil,
		OneType12: nil,
		OneType13: nil,
		OneType14: nil,
		OneType15: nil,
		OneType16: nil,
		OneType17: nil,
		OneType18: nil,
		OneType19: nil,
		OneType20: nil,
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
