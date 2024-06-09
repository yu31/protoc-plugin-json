package base

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbbase"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

type CopyMap1 pbbase.TypeMap1

var seedMap1 = &pbbase.TypeMap1{
	FString1:    map[string]string{"ks11": "kv11"},
	FString2:    map[string]string{"ks21": "kv21"},
	FInt32:      map[string]int32{"ki101": 101},
	FInt64:      map[string]int64{"ki111": 111},
	FUint32:     map[string]uint32{"ki121": 121},
	FUint64:     map[string]uint64{"ki131": 131},
	FSint32:     map[string]int32{"ki141": 141},
	FSint64:     map[string]int64{"ki151": 151},
	FSfixed32:   map[string]int32{"ki161": 161},
	FSfixed64:   map[string]int64{"ki171": 171},
	FFixed32:    map[string]uint32{"ki181": 181},
	FFixed64:    map[string]uint64{"ki191": 191},
	FFloat:      map[string]float32{"kf11.11": 11.11},
	FDouble:     map[string]float64{"kf22.21": 22.21},
	FBool1:      map[string]bool{"kbl1": true},
	FBytes1:     map[string][]byte{"kby1": []byte("bytes1")},
	FEnum1:      map[string]pbbase.TypeMap1_EnumNum1{"ke11": 2},
	FEnum2:      map[string]pbexternal.EnumNum1{"ke21": 3},
	FEnum3:      map[string]pbexternal.Embed_EnumNum1{"ke31": 5},
	FEnum4:      map[string]pbexternal.Embed_Message_EnumNum1{"ke41": 6},
	FEnum5:      map[string]pbbase.EnumNum1{"ke51": 7},
	FEnum6:      map[string]pbbase.MessageCommon1_EnumNum1{"ke61": 9},
	FDuration1:  map[string]*durationpb.Duration{"kd11": {Seconds: 101, Nanos: 102}},
	FDuration2:  map[string]*durationpb.Duration{"kd21": {Seconds: 301, Nanos: 302}},
	FTimestamp1: map[string]*timestamppb.Timestamp{"kt11": {Seconds: 201, Nanos: 202}},
	FTimestamp2: map[string]*timestamppb.Timestamp{"kt21": {Seconds: 401, Nanos: 402}},
	FAny1: map[string]*anypb.Any{
		"ka11": utils.MustNewAny(&pbexternal.Message1{FString1: "any11", FString2: "any12", FString3: "any13"}),
	},
	FAny2: map[string]*anypb.Any{
		"ka13": utils.MustNewAny(&pbbase.MessageCommon1{FString1: "any11", FString2: "any12", FString3: "any13"}),
	},
	FMessage1: map[string]*pbbase.MessageMap1{
		"km11": {FString1: "1101", FString2: "1102", FString3: "1103"},
	},
	FMessage2: map[string]*pbbase.MessageMap1_Embed1{
		"km21": {FString1: "1201", FString2: "1202", FString3: "1203"},
	},
	FMessage3: map[string]*pbbase.MessageMap1_Embed1_Embed2{
		"km31": {FString1: "1301", FString2: "1302", FString3: "1303"},
	},
	FMessage4: map[string]*pbexternal.Message1{
		"km41": {FString1: "1401", FString2: "1402", FString3: "1403"},
	},
	FMessage5: map[string]*pbexternal.Message1_Embed1{
		"km51": {FString1: "1501", FString2: "1502", FString3: "1503"},
	},
	FMessage6: map[string]*pbexternal.Message1_Embed1_Embed2{
		"km61": {FString1: "1601", FString2: "1602", FString3: "1603"},
	},
	FMessage7: map[string]*pbbase.MessageCommon1{
		"km71": {FString1: "1701", FString2: "1702", FString3: "1703"},
	},
	FMessage8: map[string]*pbbase.MessageCommon1_Embed1{
		"km81": {FString1: "1801", FString2: "1802", FString3: "1803"},
	},
	FMessage9: map[string]*pbbase.MessageCommon1_Embed1_Embed2{
		"km91": {FString1: "1901", FString2: "1902", FString3: "1903"},
	},
}

func Test_TypeMap1_Assert_Type(t *testing.T) {
	data := &pbbase.TypeMap1{}
	_, ok1 := interface{}(data).(json.Marshaler)
	require.True(t, ok1)
	_, ok2 := interface{}(data).(json.Unmarshaler)
	require.True(t, ok2)
}

func Test_TypeMap1_Assert_Copy(t *testing.T) {
	dataCopy := &CopyMap1{}
	_, ok1 := interface{}(dataCopy).(json.Marshaler)
	require.False(t, ok1)
	_, ok2 := interface{}(dataCopy).(json.Unmarshaler)
	require.False(t, ok2)
}

func Test_TypeMap1_Basic(t *testing.T) {
	var (
		err error
		b1  []byte
	)
	t.Run("Marshal", func(t *testing.T) {
		b1, err = seedMap1.MarshalJSON()
		require.Nil(t, err)
	})
	t.Run("Unmarshal", func(t *testing.T) {
		dataNew := &pbbase.TypeMap1{}
		require.NotEqual(t, seedMap1, dataNew)
		err = dataNew.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, seedMap1, dataNew)
	})
}

// Test compatible with standard JSON with follow cases:
//   1. The json content that get by MarshalJSON that code generated by plugin can be used to standard json.Unmarshal.
//   2. The json content that get by standard json.Marshal can be used to UnmarshalJSON that code generated by plugin.
func Test_TypeMap1_Compatible(t *testing.T) {
	var (
		err error
		b1  []byte
		b2  []byte
		b3  []byte
	)
	_ = b3

	// Get json content by MarshalJSON.
	b1, err = seedMap1.MarshalJSON()
	require.Nil(t, err)

	dataCopy := &CopyMap1{}
	// Test the json content is correctly.
	// and same as copy the value to dataCopy.
	err = json.Unmarshal(b1, dataCopy)
	require.Nil(t, err)

	// Get json context by json.Marshal
	b2, err = json.Marshal(dataCopy)
	require.Nil(t, err)

	// Compare the json content.
	require.Equal(t, b1, b2)

	// Test the UnmarshalJSON
	dataNew := &pbbase.TypeMap1{}
	err = dataNew.UnmarshalJSON(b2)
	require.Nil(t, err)
	require.Equal(t, seedMap1, dataNew)

	// proto
	b3, err = utils.PMarshal.Marshal(seedMap1)
	require.Nil(t, err)
}

// Test cases the field is empty.
func Test_TypeMap1_Empty(t *testing.T) {
	dataEmtpy := &pbbase.TypeMap1{}
	dataCopy := &CopyMap1{}

	var (
		err error
		b1  []byte
		b2  []byte
	)

	t.Run("Marshal", func(t *testing.T) {
		b1, err = dataEmtpy.MarshalJSON()
		require.Nil(t, err)
		b2, err = json.Marshal(dataCopy)
		require.Nil(t, err)
	})

	t.Run("UnmarshalPlugin", func(t *testing.T) {
		// use content that get by MarshalJSON
		data1 := &pbbase.TypeMap1{}
		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, dataEmtpy, data1)

		// use content that get by json.Marshal
		data2 := &pbbase.TypeMap1{}
		err = data2.UnmarshalJSON(b2)
		require.Nil(t, err)
		require.Equal(t, dataEmtpy, data2)
	})
	t.Run("UnmarshalStandard", func(t *testing.T) {
		// use content that get by MarshalJSON
		data1 := &CopyMap1{}
		err = json.Unmarshal(b1, data1)
		require.Nil(t, err)
		require.Equal(t, dataCopy, data1)

		// use content that get by json.Marshal
		data2 := &CopyMap1{}
		err = json.Unmarshal(b2, data2)
		require.Nil(t, err)
		require.Equal(t, dataCopy, data2)
	})
}

func Test_TypeMap1_NULL(t *testing.T) {
	data1 := &pbbase.TypeMap1{}
	data2 := &CopyMap1{}

	var err error
	err = data1.UnmarshalJSON([]byte("null"))
	require.Nil(t, err)

	err = json.Unmarshal([]byte("null"), data2)
	require.Nil(t, err)
}
