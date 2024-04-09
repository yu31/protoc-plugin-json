package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReadMapI64I32 read the next items from JSON contents as map key int64 and value int32.
func ReadMapI64I32(dec *Decoder, val map[int64]int32,
	unquoteKey, unquoteVal bool) (vv map[int64]int32, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]int32, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv int32
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValI32(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64I64 read the next items from JSON contents as map key int64 and value int64.
func ReadMapI64I64(dec *Decoder, val map[int64]int64,
	unquoteKey, unquoteVal bool) (vv map[int64]int64, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]int64, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv int64
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValI64(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	// Truncate the slice to consistent with standard library json
	return vv, nil
}

// ReadMapI64U32 read the next items from JSON contents as map key int64 and value uint32.
func ReadMapI64U32(dec *Decoder, val map[int64]uint32,
	unquoteKey, unquoteVal bool) (vv map[int64]uint32, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]uint32, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv uint32
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValU32(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64U64 read the next items from JSON contents as map key int64 and value uint64.
func ReadMapI64U64(dec *Decoder, val map[int64]uint64,
	unquoteKey, unquoteVal bool) (vv map[int64]uint64, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]uint64, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv uint64
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValU64(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64F32 read the next items from JSON contents as map key int64 and value float32.
func ReadMapI64F32(dec *Decoder, val map[int64]float32,
	unquoteKey, unquoteVal bool) (vv map[int64]float32, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]float32, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv float32
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValF32(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64F64 read the next items from JSON contents as map key int64 and value uint64.
func ReadMapI64F64(dec *Decoder, val map[int64]float64,
	unquoteKey, unquoteVal bool) (vv map[int64]float64, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]float64, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv float64
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValF64(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64Bool read the next items from JSON contents as map key int64 and value bool.
func ReadMapI64Bool(dec *Decoder, val map[int64]bool,
	unquoteKey, unquoteVal bool) (vv map[int64]bool, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]bool, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv bool
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValBool(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64Str read the next items from JSON contents as map key int64 and value string.
func ReadMapI64Str(dec *Decoder, val map[int64]string,
	unquoteKey bool) (vv map[int64]string, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]string, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv string
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValStr(); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64Bytes read the next items from JSON contents as map key int64 and value bytes.
func ReadMapI64Bytes(dec *Decoder, val map[int64][]byte,
	unquoteKey bool) (vv map[int64][]byte, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64][]byte, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv []byte
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = dec.readValBytes(); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64EnumNum read the next items from JSON contents as map key int64 and value enum with codec number.
// FIXME: optimized the codes.
func ReadMapI64EnumNum[T protoreflect.Enum](dec *Decoder, val map[int64]T,
	unquoteKey, unquoteVal bool) (vv map[int64]T, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}

	if val == nil {
		vv = make(map[int64]T, 0)
	} else {
		vv = val
	}

	var tt T
	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var v2 int32
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if v2, err = dec.readValEnumNum(unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return
		}
		mv := tt.Type().New(protoreflect.EnumNumber(v2)).(T)
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64EnumStr read the next items from JSON contents as map key int64 and value enum with codec string.
// FIXME: optimized the codes.
func ReadMapI64EnumStr[T protoreflect.Enum](dec *Decoder, val map[int64]T,
	unquoteKey bool, em map[string]int32) (vv map[int64]T, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}

	if val == nil {
		vv = make(map[int64]T, 0)
	} else {
		vv = val
	}

	var tt T
	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var v2 int32
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if v2, err = dec.readValEnumStr(em); err != nil {
			return
		}
		mv := tt.Type().New(protoreflect.EnumNumber(v2)).(T)
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64Message read the next items from JSON contents as map key int64 and value message.
//func ReadMapI64Message[T PtrGoAny[U], U any](dec *Decoder,  val map[int64]T) (vv map[int64]T, err error) {
func ReadMapI64Message[U any, T *U](dec *Decoder, val map[int64]T,
	unquoteKey bool) (vv map[int64]T, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]T, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv T
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if mv = vv[mk]; mv == nil {
				mv = new(U)
			}
			if err = dec.readValInterface(mv); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTAnyObject read the next items from JSON contents as map key int64 and value WKT any with codec proto.
func ReadMapI64WKTAnyObject(dec *Decoder, val map[int64]*anypb.Any,
	unquoteKey bool) (vv map[int64]*anypb.Any, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*anypb.Any, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *anypb.Any
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if mv = vv[mk]; mv == nil {
				mv = new(anypb.Any)
			}
			if err = dec.readValInterface(mv); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTAnyProto read the next items from JSON contents as map key int64 and value WKT any with codec proto.
func ReadMapI64WKTAnyProto(dec *Decoder, val map[int64]*anypb.Any,
	unquoteKey bool) (vv map[int64]*anypb.Any, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*anypb.Any, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *anypb.Any
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if mv = vv[mk]; mv == nil {
				mv = new(anypb.Any)
			}
			if err = dec.readValWKTAnyProto(mv); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurObject read the next items from JSON contents as map key int64 and value WKT duration with codec string.
func ReadMapI64WKTDurObject(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if mv = vv[mk]; mv == nil {
				mv = new(durationpb.Duration)
			}
			if err = dec.readValInterface(mv); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurTimeStr read the next items from JSON contents as map key int64 and value WKT duration with codec string.
func ReadMapI64WKTDurTimeStr(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurTimeStr(mv); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurNano read the next items from JSON contents as map key int64 and value WKT duration with codec nanosecond.
func ReadMapI64WKTDurNano(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurNano(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurMicro read the next items from JSON contents as map key int64 and value WKT duration with codec microsecond.
func ReadMapI64WKTDurMicro(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurMicro(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurMilli read the next items from JSON contents as map key int64 and value WKT duration with codec millisecond.
func ReadMapI64WKTDurMilli(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurMilli(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurSecond read the next items from JSON contents as map key int64 and value WKT duration with codec seconds.
func ReadMapI64WKTDurSecond(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurSecond(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurMinute read the next items from JSON contents as map key int64 and value WKT duration with codec minutes.
func ReadMapI64WKTDurMinute(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurMinute(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTDurHour read the next items from JSON contents as map key int64 and value WKT duration with codec hours.
func ReadMapI64WKTDurHour(dec *Decoder, val map[int64]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int64]*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurHour(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTTsObject read the next items from JSON contents as map key int64 and value WKT timestamp with codec time layout.
func ReadMapI64WKTTsObject(dec *Decoder, val map[int64]*timestamppb.Timestamp,
	unquoteKey bool) (vv map[int64]*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if mv = vv[mk]; mv == nil {
				mv = new(timestamppb.Timestamp)
			}
			if err = dec.readValInterface(mv); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTTsLayout read the next items from JSON contents as map key int64 and value WKT timestamp with codec time layout.
func ReadMapI64WKTTsLayout(dec *Decoder, val map[int64]*timestamppb.Timestamp,
	unquoteKey bool, layout string) (vv map[int64]*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsLayout(mv, layout); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTTsUnixNano read the next items from JSON contents as map key int64 and value WKT timestamp with codec unix nano.
func ReadMapI64WKTTsUnixNano(dec *Decoder, val map[int64]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int64]*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixNano(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTTsUnixMicro read the next items from JSON contents as map key int64 and value WKT timestamp with codec unix micro.
func ReadMapI64WKTTsUnixMicro(dec *Decoder, val map[int64]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int64]*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixMicro(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTTsUnixMilli read the next items from JSON contents as map key int64 and value WKT timestamp with codec unix milli.
func ReadMapI64WKTTsUnixMilli(dec *Decoder, val map[int64]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int64]*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixMilli(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI64WKTTsUnixSec read the next items from JSON contents as map key int64 and value WKT timestamp with codec unix sec.
func ReadMapI64WKTTsUnixSec(dec *Decoder, val map[int64]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int64]*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadMap(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make(map[int64]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	for { // Loop to read the map kv
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var mk int64
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI64(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv = vv[mk]; mv == nil {
			mv = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixSec(mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		vv[mk] = mv
	}
	return vv, nil
}
