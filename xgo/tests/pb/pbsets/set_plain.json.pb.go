// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/sets/set_plain.proto

package pbsets

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeSetPlain1 in file tests/proto/cases/sets/set_plain.proto
func (x *TypeSetPlain1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(704)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_any_native1")
	if err = encoder.AppendLiteralInterface(x.FAnyNative1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_any_proto1")
	if err = encoder.AppendWKTAnyByProto(x.FAnyProto1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_duration_native1")
	if err = encoder.AppendLiteralInterface(x.FDurationNative1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_duration_string1")
	encoder.AppendLiteralString(x.FDurationString1.AsDuration().String())
	encoder.AppendObjectKey("f_duration_nanoseconds1")
	encoder.AppendLiteralInt64(x.FDurationNanoseconds1.AsDuration().Nanoseconds())
	encoder.AppendObjectKey("f_duration_microseconds1")
	encoder.AppendLiteralInt64(x.FDurationMicroseconds1.AsDuration().Microseconds())
	encoder.AppendObjectKey("f_duration_milliseconds1")
	encoder.AppendLiteralInt64(x.FDurationMilliseconds1.AsDuration().Milliseconds())
	encoder.AppendObjectKey("f_duration_seconds1")
	encoder.AppendLiteralFloat64(x.FDurationSeconds1.AsDuration().Seconds())
	encoder.AppendObjectKey("f_duration_minutes1")
	encoder.AppendLiteralFloat64(x.FDurationMinutes1.AsDuration().Minutes())
	encoder.AppendObjectKey("f_duration_hours1")
	encoder.AppendLiteralFloat64(x.FDurationHours1.AsDuration().Hours())
	encoder.AppendObjectKey("f_timestamp_native1")
	if err = encoder.AppendLiteralInterface(x.FTimestampNative1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_timestamp_time_layout1")
	encoder.AppendLiteralString(x.FTimestampTimeLayout1.AsTime().Format("2006-01-02T15:04:05Z07:00"))
	encoder.AppendObjectKey("f_timestamp_unix_nano1")
	encoder.AppendLiteralInt64(x.FTimestampUnixNano1.AsTime().UnixNano())
	encoder.AppendObjectKey("f_timestamp_unix_micro1")
	encoder.AppendLiteralInt64(x.FTimestampUnixMicro1.AsTime().UnixMicro())
	encoder.AppendObjectKey("f_timestamp_unix_milli1")
	encoder.AppendLiteralInt64(x.FTimestampUnixMilli1.AsTime().UnixMilli())
	encoder.AppendObjectKey("f_timestamp_unix_sec1")
	encoder.AppendLiteralInt64(x.FTimestampUnixSec1.AsTime().Unix())

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeSetPlain1 in file tests/proto/cases/sets/set_plain.proto
func (x *TypeSetPlain1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbsets.(*TypeSetPlain1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_any_native1":
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FAnyNative1 != nil {
					vv = x.FAnyNative1
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FAnyNative1 = vv
		case "f_any_proto1":
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FAnyProto1 != nil {
					vv = x.FAnyProto1
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadWKTAnyByProto(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FAnyProto1 = vv
		case "f_duration_native1":
			var vv *durationpb.Duration
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FDurationNative1 != nil {
					vv = x.FDurationNative1
				} else {
					vv = new(durationpb.Duration)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FDurationNative1 = vv
		case "f_duration_string1":
			var vv *durationpb.Duration
			if x.FDurationString1 != nil {
				vv = x.FDurationString1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByString(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationString1 = vv
		case "f_duration_nanoseconds1":
			var vv *durationpb.Duration
			if x.FDurationNanoseconds1 != nil {
				vv = x.FDurationNanoseconds1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByNanoseconds(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationNanoseconds1 = vv
		case "f_duration_microseconds1":
			var vv *durationpb.Duration
			if x.FDurationMicroseconds1 != nil {
				vv = x.FDurationMicroseconds1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByMicroseconds(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationMicroseconds1 = vv
		case "f_duration_milliseconds1":
			var vv *durationpb.Duration
			if x.FDurationMilliseconds1 != nil {
				vv = x.FDurationMilliseconds1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByMilliseconds(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationMilliseconds1 = vv
		case "f_duration_seconds1":
			var vv *durationpb.Duration
			if x.FDurationSeconds1 != nil {
				vv = x.FDurationSeconds1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationBySeconds(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationSeconds1 = vv
		case "f_duration_minutes1":
			var vv *durationpb.Duration
			if x.FDurationMinutes1 != nil {
				vv = x.FDurationMinutes1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByMinutes(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationMinutes1 = vv
		case "f_duration_hours1":
			var vv *durationpb.Duration
			if x.FDurationHours1 != nil {
				vv = x.FDurationHours1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByHours(jsonKey, vv); err != nil {
				return err
			}
			x.FDurationHours1 = vv
		case "f_timestamp_native1":
			var vv *timestamppb.Timestamp
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FTimestampNative1 != nil {
					vv = x.FTimestampNative1
				} else {
					vv = new(timestamppb.Timestamp)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FTimestampNative1 = vv
		case "f_timestamp_time_layout1":
			var vv *timestamppb.Timestamp
			if x.FTimestampTimeLayout1 != nil {
				vv = x.FTimestampTimeLayout1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByString(jsonKey, vv, "2006-01-02T15:04:05Z07:00"); err != nil {
				return err
			}
			x.FTimestampTimeLayout1 = vv
		case "f_timestamp_unix_nano1":
			var vv *timestamppb.Timestamp
			if x.FTimestampUnixNano1 != nil {
				vv = x.FTimestampUnixNano1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixNano(jsonKey, vv); err != nil {
				return err
			}
			x.FTimestampUnixNano1 = vv
		case "f_timestamp_unix_micro1":
			var vv *timestamppb.Timestamp
			if x.FTimestampUnixMicro1 != nil {
				vv = x.FTimestampUnixMicro1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixMicro(jsonKey, vv); err != nil {
				return err
			}
			x.FTimestampUnixMicro1 = vv
		case "f_timestamp_unix_milli1":
			var vv *timestamppb.Timestamp
			if x.FTimestampUnixMilli1 != nil {
				vv = x.FTimestampUnixMilli1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixMilli(jsonKey, vv); err != nil {
				return err
			}
			x.FTimestampUnixMilli1 = vv
		case "f_timestamp_unix_sec1":
			var vv *timestamppb.Timestamp
			if x.FTimestampUnixSec1 != nil {
				vv = x.FTimestampUnixSec1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixSec(jsonKey, vv); err != nil {
				return err
			}
			x.FTimestampUnixSec1 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
