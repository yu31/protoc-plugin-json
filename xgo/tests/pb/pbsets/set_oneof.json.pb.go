// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/sets/set_oneof.proto

package pbsets

import (
	errors "errors"
	fmt "fmt"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeSetOneOf1 in file tests/proto/cases/sets/set_oneof.proto
func (x *TypeSetOneOf1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(1032)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	switch ov := x.OneEnum1.(type) {
	case *TypeSetOneOf1_FEnumNumber1:
		encoder.AppendObjectKey("one_enum1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_enum_number1")
		encoder.AppendLiteralInt32(int32(ov.FEnumNumber1.Number()))
		encoder.AppendObjectEnd()
	case *TypeSetOneOf1_FEnumString1:
		encoder.AppendObjectKey("one_enum1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_enum_string1")
		encoder.AppendLiteralString(ov.FEnumString1.String())
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_enum1")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneAny1.(type) {
	case *TypeSetOneOf1_FAnyNative1:
		encoder.AppendObjectKey("f_any_native1")
		if err = encoder.AppendLiteralInterface(ov.FAnyNative1); err != nil {
			return nil, err
		}
	case *TypeSetOneOf1_FAnyProto1:
		encoder.AppendObjectKey("f_any_proto1")
		if err = encoder.AppendWKTAnyByProto(ov.FAnyProto1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneDuration1.(type) {
	case *TypeSetOneOf1_FDurationNative1:
		encoder.AppendObjectKey("one_duration1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_duration_native1")
		if err = encoder.AppendLiteralInterface(ov.FDurationNative1); err != nil {
			return nil, err
		}
		encoder.AppendObjectEnd()
	case *TypeSetOneOf1_FDurationString1:
		encoder.AppendObjectKey("one_duration1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_duration_string1")
		encoder.AppendLiteralString(ov.FDurationString1.AsDuration().String())
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_duration1")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneDuration2.(type) {
	case *TypeSetOneOf1_FDurationNanoseconds1:
		encoder.AppendObjectKey("f_duration_nanoseconds1")
		encoder.AppendLiteralInt64(ov.FDurationNanoseconds1.AsDuration().Nanoseconds())
	case *TypeSetOneOf1_FDurationMicroseconds1:
		encoder.AppendObjectKey("f_duration_microseconds1")
		encoder.AppendLiteralInt64(ov.FDurationMicroseconds1.AsDuration().Microseconds())
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneDuration3.(type) {
	case *TypeSetOneOf1_FDurationMilliseconds1:
		encoder.AppendObjectKey("one_duration3")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_duration_milliseconds1")
		encoder.AppendLiteralInt64(ov.FDurationMilliseconds1.AsDuration().Milliseconds())
		encoder.AppendObjectEnd()
	case *TypeSetOneOf1_FDurationSeconds1:
		encoder.AppendObjectKey("one_duration3")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_duration_seconds1")
		encoder.AppendLiteralFloat64(ov.FDurationSeconds1.AsDuration().Seconds())
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_duration3")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneDuration4.(type) {
	case *TypeSetOneOf1_FDurationMinutes1:
		encoder.AppendObjectKey("f_duration_minutes1")
		encoder.AppendLiteralFloat64(ov.FDurationMinutes1.AsDuration().Minutes())
	case *TypeSetOneOf1_FDurationHours1:
		encoder.AppendObjectKey("f_duration_hours1")
		encoder.AppendLiteralFloat64(ov.FDurationHours1.AsDuration().Hours())
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneTimestamp1.(type) {
	case *TypeSetOneOf1_FTimestampNative1:
		encoder.AppendObjectKey("one_timestamp1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_timestamp_native1")
		if err = encoder.AppendLiteralInterface(ov.FTimestampNative1); err != nil {
			return nil, err
		}
		encoder.AppendObjectEnd()
	case *TypeSetOneOf1_FTimestampTimeLayout1:
		encoder.AppendObjectKey("one_timestamp1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_timestamp_time_layout1")
		encoder.AppendLiteralString(ov.FTimestampTimeLayout1.AsTime().Format("2006-01-02T15:04:05Z07:00"))
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_timestamp1")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneTimestamp2.(type) {
	case *TypeSetOneOf1_FTimestampUnixNano1:
		encoder.AppendObjectKey("f_timestamp_unix_nano1")
		encoder.AppendLiteralInt64(ov.FTimestampUnixNano1.AsTime().UnixNano())
	case *TypeSetOneOf1_FTimestampUnixMicro1:
		encoder.AppendObjectKey("f_timestamp_unix_micro1")
		encoder.AppendLiteralInt64(ov.FTimestampUnixMicro1.AsTime().UnixMicro())
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneTimestamp3.(type) {
	case *TypeSetOneOf1_FTimestampUnixMilli1:
		encoder.AppendObjectKey("one_timestamp3")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_timestamp_unix_milli1")
		encoder.AppendLiteralInt64(ov.FTimestampUnixMilli1.AsTime().UnixMilli())
		encoder.AppendObjectEnd()
	case *TypeSetOneOf1_FTimestampUnixSec1:
		encoder.AppendObjectKey("one_timestamp3")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_timestamp_unix_sec1")
		encoder.AppendLiteralInt64(ov.FTimestampUnixSec1.AsTime().Unix())
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_timestamp3")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeSetOneOf1 in file tests/proto/cases/sets/set_oneof.proto
func (x *TypeSetOneOf1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbsets.(*TypeSetOneOf1) is nil")
	}
	var (
		oneOfIsFill_OneEnum1      bool
		oneOfIsFill_OneAny1       bool
		oneOfIsFill_OneDuration1  bool
		oneOfIsFill_OneDuration2  bool
		oneOfIsFill_OneDuration3  bool
		oneOfIsFill_OneDuration4  bool
		oneOfIsFill_OneTimestamp1 bool
		oneOfIsFill_OneTimestamp2 bool
		oneOfIsFill_OneTimestamp3 bool
	)
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
		case "one_enum1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneEnum1 = nil
				continue LOOP_SCAN
			}
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch oneofKey { // match oneof key
				case "f_enum_number1":
					if oneOfIsFill_OneEnum1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneEnum1 = true

					var ok bool
					var ot *TypeSetOneOf1_FEnumNumber1
					if ot, ok = x.OneEnum1.(*TypeSetOneOf1_FEnumNumber1); !ok {
						ot = new(TypeSetOneOf1_FEnumNumber1)
					}
					var vv Enum1
					var v1 int32
					if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, Enum1_name); err != nil {
						return err
					}
					vv = Enum1(v1)
					ot.FEnumNumber1 = vv
					x.OneEnum1 = ot
				case "f_enum_string1":
					if oneOfIsFill_OneEnum1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneEnum1 = true

					var ok bool
					var ot *TypeSetOneOf1_FEnumString1
					if ot, ok = x.OneEnum1.(*TypeSetOneOf1_FEnumString1); !ok {
						ot = new(TypeSetOneOf1_FEnumString1)
					}
					var vv Enum1
					var v1 int32
					if v1, err = decoder.ReadLiteralEnumString(jsonKey, Enum1_value); err != nil {
						return err
					}
					vv = Enum1(v1)
					ot.FEnumString1 = vv
					x.OneEnum1 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_any_native1":
			if oneOfIsFill_OneAny1 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneAny1 = true

			var ok bool
			var ot *TypeSetOneOf1_FAnyNative1
			if ot, ok = x.OneAny1.(*TypeSetOneOf1_FAnyNative1); !ok {
				ot = new(TypeSetOneOf1_FAnyNative1)
			}
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if ot.FAnyNative1 != nil {
					vv = ot.FAnyNative1
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			ot.FAnyNative1 = vv
			x.OneAny1 = ot
		case "f_any_proto1":
			if oneOfIsFill_OneAny1 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneAny1 = true

			var ok bool
			var ot *TypeSetOneOf1_FAnyProto1
			if ot, ok = x.OneAny1.(*TypeSetOneOf1_FAnyProto1); !ok {
				ot = new(TypeSetOneOf1_FAnyProto1)
			}
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if ot.FAnyProto1 != nil {
					vv = ot.FAnyProto1
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadWKTAnyByProto(jsonKey, vv); err != nil {
					return err
				}
			}
			ot.FAnyProto1 = vv
			x.OneAny1 = ot
		case "one_duration1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneDuration1 = nil
				continue LOOP_SCAN
			}
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch oneofKey { // match oneof key
				case "f_duration_native1":
					if oneOfIsFill_OneDuration1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneDuration1 = true

					var ok bool
					var ot *TypeSetOneOf1_FDurationNative1
					if ot, ok = x.OneDuration1.(*TypeSetOneOf1_FDurationNative1); !ok {
						ot = new(TypeSetOneOf1_FDurationNative1)
					}
					var vv *durationpb.Duration
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FDurationNative1 != nil {
							vv = ot.FDurationNative1
						} else {
							vv = new(durationpb.Duration)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FDurationNative1 = vv
					x.OneDuration1 = ot
				case "f_duration_string1":
					if oneOfIsFill_OneDuration1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneDuration1 = true

					var ok bool
					var ot *TypeSetOneOf1_FDurationString1
					if ot, ok = x.OneDuration1.(*TypeSetOneOf1_FDurationString1); !ok {
						ot = new(TypeSetOneOf1_FDurationString1)
					}
					var vv *durationpb.Duration
					if ot.FDurationString1 != nil {
						vv = ot.FDurationString1
					} else {
						vv = new(durationpb.Duration)
					}
					if err = decoder.ReadWKTDurationByString(jsonKey, vv); err != nil {
						return err
					}
					ot.FDurationString1 = vv
					x.OneDuration1 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_duration_nanoseconds1":
			if oneOfIsFill_OneDuration2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneDuration2 = true

			var ok bool
			var ot *TypeSetOneOf1_FDurationNanoseconds1
			if ot, ok = x.OneDuration2.(*TypeSetOneOf1_FDurationNanoseconds1); !ok {
				ot = new(TypeSetOneOf1_FDurationNanoseconds1)
			}
			var vv *durationpb.Duration
			if ot.FDurationNanoseconds1 != nil {
				vv = ot.FDurationNanoseconds1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByNanoseconds(jsonKey, vv); err != nil {
				return err
			}
			ot.FDurationNanoseconds1 = vv
			x.OneDuration2 = ot
		case "f_duration_microseconds1":
			if oneOfIsFill_OneDuration2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneDuration2 = true

			var ok bool
			var ot *TypeSetOneOf1_FDurationMicroseconds1
			if ot, ok = x.OneDuration2.(*TypeSetOneOf1_FDurationMicroseconds1); !ok {
				ot = new(TypeSetOneOf1_FDurationMicroseconds1)
			}
			var vv *durationpb.Duration
			if ot.FDurationMicroseconds1 != nil {
				vv = ot.FDurationMicroseconds1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByMicroseconds(jsonKey, vv); err != nil {
				return err
			}
			ot.FDurationMicroseconds1 = vv
			x.OneDuration2 = ot
		case "one_duration3":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneDuration3 = nil
				continue LOOP_SCAN
			}
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch oneofKey { // match oneof key
				case "f_duration_milliseconds1":
					if oneOfIsFill_OneDuration3 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneDuration3 = true

					var ok bool
					var ot *TypeSetOneOf1_FDurationMilliseconds1
					if ot, ok = x.OneDuration3.(*TypeSetOneOf1_FDurationMilliseconds1); !ok {
						ot = new(TypeSetOneOf1_FDurationMilliseconds1)
					}
					var vv *durationpb.Duration
					if ot.FDurationMilliseconds1 != nil {
						vv = ot.FDurationMilliseconds1
					} else {
						vv = new(durationpb.Duration)
					}
					if err = decoder.ReadWKTDurationByMilliseconds(jsonKey, vv); err != nil {
						return err
					}
					ot.FDurationMilliseconds1 = vv
					x.OneDuration3 = ot
				case "f_duration_seconds1":
					if oneOfIsFill_OneDuration3 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneDuration3 = true

					var ok bool
					var ot *TypeSetOneOf1_FDurationSeconds1
					if ot, ok = x.OneDuration3.(*TypeSetOneOf1_FDurationSeconds1); !ok {
						ot = new(TypeSetOneOf1_FDurationSeconds1)
					}
					var vv *durationpb.Duration
					if ot.FDurationSeconds1 != nil {
						vv = ot.FDurationSeconds1
					} else {
						vv = new(durationpb.Duration)
					}
					if err = decoder.ReadWKTDurationBySeconds(jsonKey, vv); err != nil {
						return err
					}
					ot.FDurationSeconds1 = vv
					x.OneDuration3 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_duration_minutes1":
			if oneOfIsFill_OneDuration4 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneDuration4 = true

			var ok bool
			var ot *TypeSetOneOf1_FDurationMinutes1
			if ot, ok = x.OneDuration4.(*TypeSetOneOf1_FDurationMinutes1); !ok {
				ot = new(TypeSetOneOf1_FDurationMinutes1)
			}
			var vv *durationpb.Duration
			if ot.FDurationMinutes1 != nil {
				vv = ot.FDurationMinutes1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByMinutes(jsonKey, vv); err != nil {
				return err
			}
			ot.FDurationMinutes1 = vv
			x.OneDuration4 = ot
		case "f_duration_hours1":
			if oneOfIsFill_OneDuration4 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneDuration4 = true

			var ok bool
			var ot *TypeSetOneOf1_FDurationHours1
			if ot, ok = x.OneDuration4.(*TypeSetOneOf1_FDurationHours1); !ok {
				ot = new(TypeSetOneOf1_FDurationHours1)
			}
			var vv *durationpb.Duration
			if ot.FDurationHours1 != nil {
				vv = ot.FDurationHours1
			} else {
				vv = new(durationpb.Duration)
			}
			if err = decoder.ReadWKTDurationByHours(jsonKey, vv); err != nil {
				return err
			}
			ot.FDurationHours1 = vv
			x.OneDuration4 = ot
		case "one_timestamp1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneTimestamp1 = nil
				continue LOOP_SCAN
			}
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch oneofKey { // match oneof key
				case "f_timestamp_native1":
					if oneOfIsFill_OneTimestamp1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneTimestamp1 = true

					var ok bool
					var ot *TypeSetOneOf1_FTimestampNative1
					if ot, ok = x.OneTimestamp1.(*TypeSetOneOf1_FTimestampNative1); !ok {
						ot = new(TypeSetOneOf1_FTimestampNative1)
					}
					var vv *timestamppb.Timestamp
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FTimestampNative1 != nil {
							vv = ot.FTimestampNative1
						} else {
							vv = new(timestamppb.Timestamp)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FTimestampNative1 = vv
					x.OneTimestamp1 = ot
				case "f_timestamp_time_layout1":
					if oneOfIsFill_OneTimestamp1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneTimestamp1 = true

					var ok bool
					var ot *TypeSetOneOf1_FTimestampTimeLayout1
					if ot, ok = x.OneTimestamp1.(*TypeSetOneOf1_FTimestampTimeLayout1); !ok {
						ot = new(TypeSetOneOf1_FTimestampTimeLayout1)
					}
					var vv *timestamppb.Timestamp
					if ot.FTimestampTimeLayout1 != nil {
						vv = ot.FTimestampTimeLayout1
					} else {
						vv = new(timestamppb.Timestamp)
					}
					if err = decoder.ReadWKTTimestampByString(jsonKey, vv, "2006-01-02T15:04:05Z07:00"); err != nil {
						return err
					}
					ot.FTimestampTimeLayout1 = vv
					x.OneTimestamp1 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_timestamp_unix_nano1":
			if oneOfIsFill_OneTimestamp2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneTimestamp2 = true

			var ok bool
			var ot *TypeSetOneOf1_FTimestampUnixNano1
			if ot, ok = x.OneTimestamp2.(*TypeSetOneOf1_FTimestampUnixNano1); !ok {
				ot = new(TypeSetOneOf1_FTimestampUnixNano1)
			}
			var vv *timestamppb.Timestamp
			if ot.FTimestampUnixNano1 != nil {
				vv = ot.FTimestampUnixNano1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixNano(jsonKey, vv); err != nil {
				return err
			}
			ot.FTimestampUnixNano1 = vv
			x.OneTimestamp2 = ot
		case "f_timestamp_unix_micro1":
			if oneOfIsFill_OneTimestamp2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneTimestamp2 = true

			var ok bool
			var ot *TypeSetOneOf1_FTimestampUnixMicro1
			if ot, ok = x.OneTimestamp2.(*TypeSetOneOf1_FTimestampUnixMicro1); !ok {
				ot = new(TypeSetOneOf1_FTimestampUnixMicro1)
			}
			var vv *timestamppb.Timestamp
			if ot.FTimestampUnixMicro1 != nil {
				vv = ot.FTimestampUnixMicro1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixMicro(jsonKey, vv); err != nil {
				return err
			}
			ot.FTimestampUnixMicro1 = vv
			x.OneTimestamp2 = ot
		case "one_timestamp3":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneTimestamp3 = nil
				continue LOOP_SCAN
			}
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch oneofKey { // match oneof key
				case "f_timestamp_unix_milli1":
					if oneOfIsFill_OneTimestamp3 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneTimestamp3 = true

					var ok bool
					var ot *TypeSetOneOf1_FTimestampUnixMilli1
					if ot, ok = x.OneTimestamp3.(*TypeSetOneOf1_FTimestampUnixMilli1); !ok {
						ot = new(TypeSetOneOf1_FTimestampUnixMilli1)
					}
					var vv *timestamppb.Timestamp
					if ot.FTimestampUnixMilli1 != nil {
						vv = ot.FTimestampUnixMilli1
					} else {
						vv = new(timestamppb.Timestamp)
					}
					if err = decoder.ReadWKTTimestampByUnixMilli(jsonKey, vv); err != nil {
						return err
					}
					ot.FTimestampUnixMilli1 = vv
					x.OneTimestamp3 = ot
				case "f_timestamp_unix_sec1":
					if oneOfIsFill_OneTimestamp3 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneTimestamp3 = true

					var ok bool
					var ot *TypeSetOneOf1_FTimestampUnixSec1
					if ot, ok = x.OneTimestamp3.(*TypeSetOneOf1_FTimestampUnixSec1); !ok {
						ot = new(TypeSetOneOf1_FTimestampUnixSec1)
					}
					var vv *timestamppb.Timestamp
					if ot.FTimestampUnixSec1 != nil {
						vv = ot.FTimestampUnixSec1
					} else {
						vv = new(timestamppb.Timestamp)
					}
					if err = decoder.ReadWKTTimestampByUnixSec(jsonKey, vv); err != nil {
						return err
					}
					ot.FTimestampUnixSec1 = vv
					x.OneTimestamp3 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
