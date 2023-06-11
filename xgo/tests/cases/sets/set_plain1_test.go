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
		FAnyNative2: utils.MustNewAny(&pbsets.Message1{
			FString1: "s201",
			FString2: "s202",
			FString3: "s203",
		}),
		FAnyProto1: utils.MustNewAny(&pbsets.Message1{
			FString1: "s301",
			FString2: "s302",
			FString3: "s303",
		}),
		FAnyProto2: utils.MustNewAny(&pbsets.Message1{
			FString1: "s401",
			FString2: "s402",
			FString3: "s403",
		}),
		FDurationNative1:       &durationpb.Duration{Seconds: 111, Nanos: 0},
		FDurationNative2:       &durationpb.Duration{Seconds: 211, Nanos: 0},
		FDurationString1:       &durationpb.Duration{Seconds: 112, Nanos: 0},
		FDurationString2:       &durationpb.Duration{Seconds: 212, Nanos: 0},
		FDurationNanoseconds1:  &durationpb.Duration{Seconds: 113, Nanos: 0},
		FDurationNanoseconds2:  &durationpb.Duration{Seconds: 213, Nanos: 0},
		FDurationMicroseconds1: &durationpb.Duration{Seconds: 114, Nanos: 0},
		FDurationMicroseconds2: &durationpb.Duration{Seconds: 214, Nanos: 0},
		FDurationMilliseconds1: &durationpb.Duration{Seconds: 115, Nanos: 0},
		FDurationMilliseconds2: &durationpb.Duration{Seconds: 215, Nanos: 0},
		FDurationSeconds1:      &durationpb.Duration{Seconds: 116, Nanos: 0},
		FDurationSeconds2:      &durationpb.Duration{Seconds: 216, Nanos: 0},
		FDurationMinutes1:      &durationpb.Duration{Seconds: 117, Nanos: 0},
		FDurationMinutes2:      &durationpb.Duration{Seconds: 217, Nanos: 0},
		FDurationHours1:        &durationpb.Duration{Seconds: 7200, Nanos: 0},
		FDurationHours2:        &durationpb.Duration{Seconds: 7200, Nanos: 0},
		FTimestampNative1:      &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampNative2:      &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampTimeLayout1:  &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampTimeLayout2:  &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixNano1:    &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixNano2:    &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixMicro1:   &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixMicro2:   &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixMilli1:   &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixMilli2:   &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixSec1:     &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestampUnixSec2:     &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
	}

	bb, err := data1.MarshalJSON()
	require.Nil(t, err)
	fmt.Println(string(bb))

	data2 := &pbsets.TypeSetPlain1{}
	err = data2.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, data1, data2)
}
