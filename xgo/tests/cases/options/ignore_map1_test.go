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

func Test_IgnoreMap1_FieldNotNil(t *testing.T) {
	seed1 := &pboptions.IgnoreMap1{
		FInt32:    map[int32]int32{111: 110, 112: 111, 113: 112},
		FInt64:    map[int64]int64{121: 120, 122: 121, 123: 122},
		FUint32:   map[uint32]uint32{131: 130, 132: 131, 133: 132},
		FUint64:   map[uint64]uint64{141: 140, 142: 141, 143: 142},
		FSint32:   map[int32]int32{151: 150, 152: 151, 153: 152},
		FSint64:   map[int64]int64{161: 160, 162: 161, 163: 162},
		FSfixed32: map[int32]int32{171: 170, 172: 171, 173: 172},
		FSfixed64: map[int64]int64{181: 180, 182: 181, 183: 182},
		FFixed32:  map[uint32]uint32{191: 190, 192: 191, 193: 192},
		FFixed64:  map[uint64]uint64{201: 200, 202: 201, 203: 202},
		FFloat:    map[string]float32{"f1": 211.111, "f2": 212.111, "f3": 213.111},
		FDouble:   map[string]float64{"d1": 221.111, "d2": 222.111, "d3": 223.111},
		FBool1:    map[bool]bool{true: false, false: true},
		FString1:  map[string]string{"sk101": "sv101", "sk102": "sv102", "sk103": "sv103"},
		FBytes1:   map[string][]byte{"bk101": []byte("bb101"), "bk102": []byte("bb102"), "bk103": []byte("bb103")},
		FEnum1:    map[string]pbexternal.EnumNum1{"ek1": 1, "ek2": 3, "ek3": 5},
		FMessage1: map[string]*pbexternal.Message1{
			"mk101": {FString1: "ms101", FString2: "ms102", FString3: "ms103"},
			"mk201": {FString1: "ms201", FString2: "ms202", FString3: "ms203"},
			"mk301": {FString1: "ms301", FString2: "ms302", FString3: "ms303"},
		},
		FAny1: map[string]*anypb.Any{
			"ak101": utils.MustNewAny(&pbexternal.Message1{FString1: "s101", FString2: "s102", FString3: "s103"}),
			"ak201": utils.MustNewAny(&pbexternal.Message1{FString1: "s201", FString2: "s202", FString3: "s203"}),
			"ak301": utils.MustNewAny(&pbexternal.Message1{FString1: "s301", FString2: "s302", FString3: "s303"}),
		},
		FDuration1: map[string]*durationpb.Duration{
			"dk101": {Seconds: 1800, Nanos: 0},
			"dk201": {Seconds: 3600, Nanos: 0},
			"dk301": {Seconds: 7200, Nanos: 0},
		},
		FTimestamp1: map[string]*timestamppb.Timestamp{
			"tk101": {Seconds: 1686416585, Nanos: 0},
			"tk201": {Seconds: 1686416585, Nanos: 0},
			"tk301": {Seconds: 1686416585, Nanos: 0},
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

func Test_IgnoreMap1_FieldIsNil(t *testing.T) {
	seed1 := &pboptions.IgnoreMap1{
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
