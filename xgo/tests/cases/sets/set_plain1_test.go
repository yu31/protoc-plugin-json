package sets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbsets"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_WKT_Plain1(t *testing.T) {
	// 2023-06-11 23:02:47.794758
	ts := int64(1686495767)
	data1 := &pbsets.TypeSetPlain1{
		FAnyNative1: utils.MustNewAny(&pbsets.Message1{
			FString1: "s101",
			FString2: "s102",
			FString3: "s103",
		}),
		FAnyProto1: utils.MustNewAny(&pbsets.Message1{
			FString1: "s301",
			FString2: "s302",
			FString3: "s303",
		}),
		FDurationNative1:       &durationpb.Duration{Seconds: 111, Nanos: 0},
		FDurationString1:       &durationpb.Duration{Seconds: 112, Nanos: 0},
		FDurationNanoseconds1:  &durationpb.Duration{Seconds: 113, Nanos: 0},
		FDurationMicroseconds1: &durationpb.Duration{Seconds: 114, Nanos: 0},
		FDurationMilliseconds1: &durationpb.Duration{Seconds: 115, Nanos: 0},
		FDurationSeconds1:      &durationpb.Duration{Seconds: 116, Nanos: 0},
		FDurationMinutes1:      &durationpb.Duration{Seconds: 117, Nanos: 0},
		FDurationHours1:        &durationpb.Duration{Seconds: 7200, Nanos: 0},
		FTimestampNative1:      &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampTimeLayout1:  &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixNano1:    &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixMicro1:   &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixMilli1:   &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixSec1:     &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
	}

	bb, err := data1.MarshalJSON()
	require.Nil(t, err)
	fmt.Println(string(bb))

	data2 := &pbsets.TypeSetPlain1{}
	err = data2.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, data1, data2)
}
