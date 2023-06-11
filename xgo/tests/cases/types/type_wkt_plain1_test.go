package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbtypes"
	"github.com/yu31/protoc-plugin-json/xgo/tests/utils"
)

func Test_TypeWKTPlain1(t *testing.T) {
	// 2023-06-11 23:02:47.794758
	ts := int64(1686495767)
	data1 := &pbtypes.TypeWKTPlain1{
		FAny1: utils.MustNewAny(&pbtypes.MessageWKT1{
			FString1: "s101",
			FString2: "s102",
			FString3: "s103",
		}),
		FAny2: utils.MustNewAny(&pbtypes.MessageWKT1{
			FString1: "s201",
			FString2: "s202",
			FString3: "s203",
		}),
		FAny3: utils.MustNewAny(&pbtypes.MessageWKT1{
			FString1: "s301",
			FString2: "s302",
			FString3: "s303",
		}),
		FAny4: utils.MustNewAny(&pbtypes.MessageWKT1{
			FString1: "s401",
			FString2: "s402",
			FString3: "s403",
		}),
		FDuration11:  &durationpb.Duration{Seconds: 111, Nanos: 0},
		FDuration21:  &durationpb.Duration{Seconds: 211, Nanos: 0},
		FDuration12:  &durationpb.Duration{Seconds: 112, Nanos: 0},
		FDuration22:  &durationpb.Duration{Seconds: 212, Nanos: 0},
		FDuration13:  &durationpb.Duration{Seconds: 113, Nanos: 0},
		FDuration23:  &durationpb.Duration{Seconds: 213, Nanos: 0},
		FDuration14:  &durationpb.Duration{Seconds: 114, Nanos: 0},
		FDuration24:  &durationpb.Duration{Seconds: 214, Nanos: 0},
		FDuration15:  &durationpb.Duration{Seconds: 115, Nanos: 0},
		FDuration25:  &durationpb.Duration{Seconds: 215, Nanos: 0},
		FDuration16:  &durationpb.Duration{Seconds: 116, Nanos: 0},
		FDuration26:  &durationpb.Duration{Seconds: 216, Nanos: 0},
		FDuration17:  &durationpb.Duration{Seconds: 117, Nanos: 0},
		FDuration27:  &durationpb.Duration{Seconds: 217, Nanos: 0},
		FDuration18:  &durationpb.Duration{Seconds: 7200, Nanos: 0},
		FDuration28:  &durationpb.Duration{Seconds: 7200, Nanos: 0},
		FTimestamp31: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp41: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp32: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp42: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp33: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp43: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp34: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp44: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp35: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp45: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp36: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
		FTimestamp46: &timestamppb.Timestamp{Seconds: ts, Nanos: 0},
	}

	bb, err := data1.MarshalJSON()
	require.Nil(t, err)
	fmt.Println(string(bb))

	data2 := &pbtypes.TypeWKTPlain1{}
	err = data2.UnmarshalJSON(bb)
	require.Nil(t, err)

	require.Equal(t, data1, data2)
}
