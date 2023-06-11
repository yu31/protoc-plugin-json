package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbtypes"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

type CopyOneOf1 pbtypes.TypeOneOf1
type CopyOneOfHide1 pbtypes.TypeOneOfHide1

func Test_TypeOneOf1_Assert_Type(t *testing.T) {
	data := &pbtypes.TypeOneOf1{}
	_, ok1 := interface{}(data).(json.Marshaler)
	require.True(t, ok1)
	_, ok2 := interface{}(data).(json.Unmarshaler)
	require.True(t, ok2)
}

func Test_TypeOneOf1_Assert_Copy(t *testing.T) {
	dataCopy := &CopyOneOf1{}
	_, ok1 := interface{}(dataCopy).(json.Marshaler)
	require.False(t, ok1)
	_, ok2 := interface{}(dataCopy).(json.Unmarshaler)
	require.False(t, ok2)
}

func Test_TypeOneOf1_General1(t *testing.T) {
	data1 := &pbtypes.TypeOneOf1{
		OneType01:  &pbtypes.TypeOneOf1_FString1{FString1: "s1"},
		One_Type02: &pbtypes.TypeOneOf1_FInt32{FInt32: 10},
		OneType03:  &pbtypes.TypeOneOf1_FUint32{FUint32: 12},
		One_Type04: &pbtypes.TypeOneOf1_FSint32{FSint32: 14},
		OneType05:  &pbtypes.TypeOneOf1_FSfixed32{FSfixed32: 16},
		OneType06:  &pbtypes.TypeOneOf1_FFixed32{FFixed32: 18},
		Onetype07:  &pbtypes.TypeOneOf1_FFloat{FFloat: 11.11},
		Onetype08:  &pbtypes.TypeOneOf1_FBool1{FBool1: true},
		OneType09:  &pbtypes.TypeOneOf1_FEnum1{FEnum1: 2},
		OneType10:  &pbtypes.TypeOneOf1_FDuration1{FDuration1: &durationpb.Duration{Seconds: 100, Nanos: 101}},
		OneType11:  &pbtypes.TypeOneOf1_FTimestamp1{FTimestamp1: &timestamppb.Timestamp{Seconds: 200, Nanos: 201}},
		OneType12:  &pbtypes.TypeOneOf1_FAny1{FAny1: utils.MustNewAny(&pbexternal.Message1{FString1: "any11", FString2: "any12", FString3: "any13"})},
		OneType13:  &pbtypes.TypeOneOf1_FMessage1{FMessage1: &pbtypes.MessageOneOf1{FString1: "1101", FString2: "1102", FString3: "1103"}},
	}
	data2 := &CopyOneOf1{
		OneType01:  &pbtypes.TypeOneOf1_FString1{FString1: "s1"},
		One_Type02: &pbtypes.TypeOneOf1_FInt32{FInt32: 10},
		OneType03:  &pbtypes.TypeOneOf1_FUint32{FUint32: 12},
		One_Type04: &pbtypes.TypeOneOf1_FSint32{FSint32: 14},
		OneType05:  &pbtypes.TypeOneOf1_FSfixed32{FSfixed32: 16},
		OneType06:  &pbtypes.TypeOneOf1_FFixed32{FFixed32: 18},
		Onetype07:  &pbtypes.TypeOneOf1_FFloat{FFloat: 11.11},
		Onetype08:  &pbtypes.TypeOneOf1_FBool1{FBool1: true},
		OneType09:  &pbtypes.TypeOneOf1_FEnum1{FEnum1: 2},
		OneType10:  &pbtypes.TypeOneOf1_FDuration1{FDuration1: &durationpb.Duration{Seconds: 100, Nanos: 101}},
		OneType11:  &pbtypes.TypeOneOf1_FTimestamp1{FTimestamp1: &timestamppb.Timestamp{Seconds: 200, Nanos: 201}},
		OneType12:  &pbtypes.TypeOneOf1_FAny1{FAny1: utils.MustNewAny(&pbexternal.Message1{FString1: "any11", FString2: "any12", FString3: "any13"})},
		OneType13:  &pbtypes.TypeOneOf1_FMessage1{FMessage1: &pbtypes.MessageOneOf1{FString1: "1101", FString2: "1102", FString3: "1103"}},
	}

	var (
		err error
		b1  []byte
		b2  []byte
		b3  []byte
	)
	_ = b2
	_ = b3

	t.Run("marshal", func(t *testing.T) {
		b1, err = data1.MarshalJSON()
		require.Nil(t, err)
	})
	t.Run("unmarshal", func(t *testing.T) {
		dataNew := &pbtypes.TypeOneOf1{}
		err = dataNew.UnmarshalJSON(b1)
		require.Nil(t, err)
	})

	t.Run("standard", func(t *testing.T) {
		b2, err = json.Marshal(data2)
		require.Nil(t, err)
	})
	t.Run("proto", func(t *testing.T) {
		b3, err = utils.PMarshal.Marshal(data1)
		require.Nil(t, err)
	})

	fmt.Println("=========================================== PLUGIN =============================================")
	fmt.Println(string(b1))
	fmt.Println("=========================================== STANDARD ===========================================")
	fmt.Println(string(b2))
	fmt.Println("=========================================== PROTO ==============================================")
	fmt.Println(string(b3))
	fmt.Println("================================================================================================")
}
func Test_TypeOneOf1_General2(t *testing.T) {
	data1 := &pbtypes.TypeOneOf1{
		OneType01:  &pbtypes.TypeOneOf1_FString2{FString2: "s1"},
		One_Type02: &pbtypes.TypeOneOf1_FInt64{FInt64: 11},
		OneType03:  &pbtypes.TypeOneOf1_FUint64{FUint64: 13},
		One_Type04: &pbtypes.TypeOneOf1_FSint64{FSint64: 15},
		OneType05:  &pbtypes.TypeOneOf1_FSfixed64{FSfixed64: 17},
		OneType06:  &pbtypes.TypeOneOf1_FFixed64{FFixed64: 19},
		Onetype07:  &pbtypes.TypeOneOf1_FDouble{FDouble: 12.12},
		Onetype08:  &pbtypes.TypeOneOf1_FBytes1{FBytes1: []byte("bytes1")},
		OneType09:  &pbtypes.TypeOneOf1_FEnum2{FEnum2: 2},
		OneType10:  &pbtypes.TypeOneOf1_FDuration2{FDuration2: &durationpb.Duration{Seconds: 0, Nanos: 0}},
		OneType11:  &pbtypes.TypeOneOf1_FTimestamp2{FTimestamp2: &timestamppb.Timestamp{Seconds: 0, Nanos: 0}},
		OneType12:  &pbtypes.TypeOneOf1_FAny2{FAny2: utils.MustNewAny(&pbtypes.MessageOneOf1{FString1: "any21", FString2: "any22", FString3: "any23"})},
		OneType13:  &pbtypes.TypeOneOf1_FMessage2{FMessage2: &pbtypes.MessageOneOf1_Embed1{FString1: "1201", FString2: "1202", FString3: "1203"}},
	}

	var (
		err error
		b1  []byte
	)

	t.Run("marshal", func(t *testing.T) {
		b1, err = data1.MarshalJSON()
		require.Nil(t, err)
	})
	t.Run("unmarshal", func(t *testing.T) {
		dataNew := &pbtypes.TypeOneOf1{}
		err = dataNew.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, data1, dataNew)
	})
}

func Test_TypeOneOfHide1(t *testing.T) {
	data1 := &pbtypes.TypeOneOfHide1{
		OneType01: &pbtypes.TypeOneOfHide1_FString1{FString1: "s1"},
		OneType02: &pbtypes.TypeOneOfHide1_FMessage2{FMessage2: &pbtypes.MessageOneOf1{FString1: "1201", FString2: "1202", FString3: "1203"}},
		OneType03: &pbtypes.TypeOneOfHide1_FEnum2{FEnum2: 2},
		OneType04: &pbtypes.TypeOneOfHide1_FDuration2{FDuration2: &durationpb.Duration{Seconds: 0, Nanos: 0}},
	}
	data2 := &CopyOneOfHide1{
		OneType01: &pbtypes.TypeOneOfHide1_FString1{FString1: "s1"},
	}

	var (
		err error
		b1  []byte
		b2  []byte
		b3  []byte
	)
	_ = b2
	_ = b3

	t.Run("marshal", func(t *testing.T) {
		b1, err = data1.MarshalJSON()
		require.Nil(t, err)
	})
	t.Run("unmarshal", func(t *testing.T) {
		dataNew := &pbtypes.TypeOneOf1{}
		err = dataNew.UnmarshalJSON(b1)
		require.Nil(t, err)
	})

	t.Run("standard", func(t *testing.T) {
		b2, err = json.Marshal(data2)
		require.Nil(t, err)
	})
	t.Run("proto", func(t *testing.T) {
		b3, err = utils.PMarshal.Marshal(data1)
		require.Nil(t, err)
		//require.Equal(t, b1, b3)
	})

	fmt.Println("=========================================== PLUGIN =============================================")
	fmt.Println(string(b1))
	fmt.Println("=========================================== STANDARD ===========================================")
	fmt.Println(string(b2))
	fmt.Println("=========================================== PROTO ==============================================")
	fmt.Println(string(b3))
	fmt.Println("================================================================================================")
}

func Test_TypeOneOf1_NULL(t *testing.T) {
	data1 := &pbtypes.TypeOneOf1{}
	data2 := &CopyOneOf1{}

	var err error
	err = data1.UnmarshalJSON([]byte("null"))
	require.Nil(t, err)

	err = json.Unmarshal([]byte("null"), data2)
	require.Nil(t, err)
}
