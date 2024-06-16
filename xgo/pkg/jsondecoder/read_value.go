package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReadValI32 read the next items from JSON contents as value of int32.
func ReadValI32(dec *Decoder, unquote bool) (vv int32, err error) {
	if vv, err = dec.readValI32(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValI64 read the next items from JSON contents as value of int64.
func ReadValI64(dec *Decoder, unquote bool) (vv int64, err error) {
	if vv, err = dec.readValI64(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValU32 read the next items from JSON contents as value of uint32.
func ReadValU32(dec *Decoder, unquote bool) (vv uint32, err error) {
	if vv, err = dec.readValU32(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValU64 read the next items from JSON contents as value of uint64.
func ReadValU64(dec *Decoder, unquote bool) (vv uint64, err error) {
	if vv, err = dec.readValU64(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValF32 read the next items from JSON contents as value of float32.
func ReadValF32(dec *Decoder, unquote bool) (vv float32, err error) {
	if vv, err = dec.readValF32(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValF64 read the next items from JSON contents as value of uint64.
func ReadValF64(dec *Decoder, unquote bool) (vv float64, err error) {
	if vv, err = dec.readValF64(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValBool read the next items from JSON contents as value of bool.
func ReadValBool(dec *Decoder, unquote bool) (vv bool, err error) {
	if vv, err = dec.readValBool(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValStr read the next items from JSON contents as value of string.
func ReadValStr(dec *Decoder) (vv string, err error) {
	if vv, err = dec.readValStr(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValBytes read the next items from JSON contents as value of bytes.
func ReadValBytes(dec *Decoder) (vv []byte, err error) {
	if vv, err = dec.readValBytes(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValEnumNum read the next items from JSON contents as value of enum with codec number.
func ReadValEnumNum[T protoreflect.Enum](dec *Decoder, val T, unquote bool) (vv T, err error) {
	if vv, err = readValEnumNum(dec, val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadValEnumStr read the next items from JSON contents as value of enum with codec string.
func ReadValEnumStr[T protoreflect.Enum](dec *Decoder, val T) (vv T, err error) {
	if vv, err = readValEnumStr(dec, val); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return
}

// ReadValMessage read the next items from JSON contents as value of message.
//func ReadValMessage[T PtrGoAny[U], U any](dec *Decoder,  val []T) (vv T, err error) {
func ReadValMessage[U any, T *U](dec *Decoder, val T) (vv T, err error) {
	var isNULL bool
	if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
		return
	}
	if isNULL {
		return nil, nil
	}
	if vv = val; vv == nil {
		vv = new(U)
	}
	if err = dec.readValInterface(vv); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTAnyObject read the next items from JSON contents as value of WKT any with codec proto.
func ReadValWKTAnyObject(dec *Decoder, val *anypb.Any) (vv *anypb.Any, err error) {
	var isNULL bool
	if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
		return
	}
	if isNULL {
		return nil, nil
	}
	if vv = val; vv == nil {
		vv = new(anypb.Any)
	}
	if err = dec.readValInterface(vv); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTAnyProto read the next items from JSON contents as value of WKT any with codec proto.
func ReadValWKTAnyProto(dec *Decoder, val *anypb.Any) (vv *anypb.Any, err error) {
	var isNULL bool
	if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
		return
	}
	if isNULL {
		return nil, nil
	}
	if vv = val; vv == nil {
		vv = new(anypb.Any)
	}
	if err = dec.readValWKTAnyProto(vv); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurObject read the next items from JSON contents as value of WKT duration with codec string.
func ReadValWKTDurObject(dec *Decoder, val *durationpb.Duration) (vv *durationpb.Duration, err error) {
	var isNULL bool
	if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
		return
	}
	if isNULL {
		return nil, nil
	}
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValInterface(vv); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurTimeStr read the next items from JSON contents as value of WKT duration with codec string.
func ReadValWKTDurTimeStr(dec *Decoder, val *durationpb.Duration) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurTimeStr(vv); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurNano read the next items from JSON contents as value of WKT duration with codec nanosecond.
func ReadValWKTDurNano(dec *Decoder, val *durationpb.Duration,
	unquote bool) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurNano(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurMicro read the next items from JSON contents as value of WKT duration with codec microsecond.
func ReadValWKTDurMicro(dec *Decoder, val *durationpb.Duration,
	unquote bool) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurMicro(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurMilli read the next items from JSON contents as value of WKT duration with codec millisecond.
func ReadValWKTDurMilli(dec *Decoder, val *durationpb.Duration,
	unquote bool) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurMilli(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurSecond read the next items from JSON contents as value of WKT duration with codec seconds.
func ReadValWKTDurSecond(dec *Decoder, val *durationpb.Duration,
	unquote bool) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurSecond(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurMinute read the next items from JSON contents as value of WKT duration with codec minutes.
func ReadValWKTDurMinute(dec *Decoder, val *durationpb.Duration,
	unquote bool) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurMinute(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTDurHour read the next items from JSON contents as value of WKT duration with codec hours.
func ReadValWKTDurHour(dec *Decoder, val *durationpb.Duration,
	unquote bool) (vv *durationpb.Duration, err error) {
	if vv = val; vv == nil {
		vv = new(durationpb.Duration)
	}
	if err = dec.readValWKTDurHour(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTTsObject read the next items from JSON contents as value of WKT timestamp with codec time layout.
func ReadValWKTTsObject(dec *Decoder, val *timestamppb.Timestamp) (vv *timestamppb.Timestamp, err error) {
	var isNULL bool
	if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
		return
	}
	if isNULL {
		return nil, nil
	}
	if vv = val; vv == nil {
		vv = new(timestamppb.Timestamp)
	}
	if err = dec.readValInterface(vv); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTTsLayout read the next items from JSON contents as value of WKT timestamp with codec time layout.
func ReadValWKTTsLayout(dec *Decoder, val *timestamppb.Timestamp,
	layout string) (vv *timestamppb.Timestamp, err error) {
	if vv = val; vv == nil {
		vv = new(timestamppb.Timestamp)
	}
	if err = dec.readValWKTTsLayout(vv, layout); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTTsUnixNano read the next items from JSON contents as value of WKT timestamp with codec unix nano.
func ReadValWKTTsUnixNano(dec *Decoder, val *timestamppb.Timestamp,
	unquote bool) (vv *timestamppb.Timestamp, err error) {
	if vv = val; vv == nil {
		vv = new(timestamppb.Timestamp)
	}
	if err = dec.readValWKTTsUnixNano(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTTsUnixMicro read the next items from JSON contents as value of WKT timestamp with codec unix micro.
func ReadValWKTTsUnixMicro(dec *Decoder, val *timestamppb.Timestamp,
	unquote bool) (vv *timestamppb.Timestamp, err error) {
	if vv = val; vv == nil {
		vv = new(timestamppb.Timestamp)
	}
	if err = dec.readValWKTTsUnixMicro(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTTsUnixMilli read the next items from JSON contents as value of WKT timestamp with codec unix milli.
func ReadValWKTTsUnixMilli(dec *Decoder, val *timestamppb.Timestamp,
	unquote bool) (vv *timestamppb.Timestamp, err error) {
	if vv = val; vv == nil {
		vv = new(timestamppb.Timestamp)
	}
	if err = dec.readValWKTTsUnixMilli(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}

// ReadValWKTTsUnixSec read the next items from JSON contents as value of WKT timestamp with codec unix sec.
func ReadValWKTTsUnixSec(dec *Decoder, val *timestamppb.Timestamp,
	unquote bool) (vv *timestamppb.Timestamp, err error) {
	if vv = val; vv == nil {
		vv = new(timestamppb.Timestamp)
	}
	if err = dec.readValWKTTsUnixSec(vv, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, err
}
