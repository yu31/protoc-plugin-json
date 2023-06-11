package advanced

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_WKT_Any1(t *testing.T) {
	m1 := &pbexternal.Message1{
		FString1: "s101",
		FString2: "s102",
		FString3: "s103",
	}

	t.Run("proto", func(t *testing.T) {
		any1 := utils.MustNewAny(m1)

		bb, err := utils.PMarshal.Marshal(any1)
		require.Nil(t, err)

		any2 := &anypb.Any{}
		err = utils.PUnmarshal.Unmarshal(bb, any2)
		require.Nil(t, err)

		m2 := &pbexternal.Message1{}
		err = any2.UnmarshalTo(m2)
		require.Nil(t, err)

		require.Equal(t, m1.FString1, m2.FString1)
		require.Equal(t, m1.FString2, m2.FString2)
		require.Equal(t, m1.FString3, m2.FString3)

		fmt.Println("-------------------------------")
		fmt.Println(string(bb))
		fmt.Println(m2)
		fmt.Println("-------------------------------")
	})

	t.Run("standard", func(t *testing.T) {
		any1 := utils.MustNewAny(m1)
		bb, err := json.Marshal(any1)
		require.Nil(t, err)

		any2 := &anypb.Any{}
		err = json.Unmarshal(bb, any2)
		require.Nil(t, err)

		m2 := &pbexternal.Message1{}
		err = any2.UnmarshalTo(m2)
		require.Nil(t, err)

		require.Equal(t, m1.FString1, m2.FString1)
		require.Equal(t, m1.FString2, m2.FString2)
		require.Equal(t, m1.FString3, m2.FString3)

		fmt.Println("-------------------------------")
		fmt.Println(string(bb))
		fmt.Println(m2)
		fmt.Println("-------------------------------")
	})
}

func Test_WKT_Duration1(t *testing.T) {
	d1 := &durationpb.Duration{
		Seconds: 100,
		Nanos:   100,
	}

	dd := d1.AsDuration()

	fmt.Println("--------------------")

	fmt.Println("Duration.Nanoseconds:", dd.Nanoseconds())
	fmt.Println("Duration.Microseconds:", dd.Microseconds())
	fmt.Println("Duration.Milliseconds:", dd.Milliseconds())
	fmt.Println("Duration.Seconds:", dd.Seconds())
	fmt.Println("Duration.Minutes:", dd.Minutes())
	fmt.Println("Duration.Hours:", dd.Hours())
	fmt.Println("Duration.String:", dd.String())

	fmt.Println("--------------------")
	fmt.Println(time.ParseDuration(dd.String()))
	fmt.Println(time.Duration(dd.Seconds() * float64(time.Second)).Seconds())
	fmt.Println("-----------------")
}

func Test_WKT_Timestamp(t *testing.T) {
	ts1 := timestamppb.New(time.Now())

	tt1 := ts1.AsTime()

	fmt.Println("--------------------")
	fmt.Println("Timestamp.String:", tt1.Format(time.RFC3339Nano))
	fmt.Println("Timestamp.Unix:", tt1.Unix())
	fmt.Println("Timestamp.UnixMilli:", tt1.UnixMilli())
	fmt.Println("Timestamp.UnixMicro:", tt1.UnixMicro())
	fmt.Println("Timestamp.UnixNano:", tt1.UnixNano())
	fmt.Println("--------------------")
}
