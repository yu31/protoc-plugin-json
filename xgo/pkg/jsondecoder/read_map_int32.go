package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReadMapI32I32 read the next items from JSON contents as map key int32 and value int32.
func ReadMapI32I32(dec *Decoder, val map[int32]int32,
	unquoteKey, unquoteVal bool) (vv map[int32]int32, err error) {
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
		vv = make(map[int32]int32, 0)
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
		var mk int32
		var mv int32
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32I64 read the next items from JSON contents as map key int32 and value int64.
func ReadMapI32I64(dec *Decoder, val map[int32]int64,
	unquoteKey, unquoteVal bool) (vv map[int32]int64, err error) {
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
		vv = make(map[int32]int64, 0)
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
		var mk int32
		var mv int64
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32U32 read the next items from JSON contents as map key int32 and value uint32.
func ReadMapI32U32(dec *Decoder, val map[int32]uint32,
	unquoteKey, unquoteVal bool) (vv map[int32]uint32, err error) {
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
		vv = make(map[int32]uint32, 0)
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
		var mk int32
		var mv uint32
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32U64 read the next items from JSON contents as map key int32 and value uint64.
func ReadMapI32U64(dec *Decoder, val map[int32]uint64,
	unquoteKey, unquoteVal bool) (vv map[int32]uint64, err error) {
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
		vv = make(map[int32]uint64, 0)
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
		var mk int32
		var mv uint64
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32F32 read the next items from JSON contents as map key int32 and value float32.
func ReadMapI32F32(dec *Decoder, val map[int32]float32,
	unquoteKey, unquoteVal bool) (vv map[int32]float32, err error) {
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
		vv = make(map[int32]float32, 0)
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
		var mk int32
		var mv float32
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32F64 read the next items from JSON contents as map key int32 and value uint64.
func ReadMapI32F64(dec *Decoder, val map[int32]float64,
	unquoteKey, unquoteVal bool) (vv map[int32]float64, err error) {
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
		vv = make(map[int32]float64, 0)
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
		var mk int32
		var mv float64
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32Bool read the next items from JSON contents as map key int32 and value bool.
func ReadMapI32Bool(dec *Decoder, val map[int32]bool,
	unquoteKey, unquoteVal bool) (vv map[int32]bool, err error) {
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
		vv = make(map[int32]bool, 0)
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
		var mk int32
		var mv bool
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32Str read the next items from JSON contents as map key int32 and value string.
func ReadMapI32Str(dec *Decoder, val map[int32]string,
	unquoteKey bool) (vv map[int32]string, err error) {
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
		vv = make(map[int32]string, 0)
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
		var mk int32
		var mv string
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32Bytes read the next items from JSON contents as map key int32 and value bytes.
func ReadMapI32Bytes(dec *Decoder, val map[int32][]byte,
	unquoteKey bool) (vv map[int32][]byte, err error) {
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
		vv = make(map[int32][]byte, 0)
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
		var mk int32
		var mv []byte
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32EnumNum read the next items from JSON contents as map key int32 and value enum with codec number.
func ReadMapI32EnumNum[T protoreflect.Enum](dec *Decoder, val map[int32]T,
	unquoteKey, unquoteVal bool) (vv map[int32]T, err error) {
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
		vv = make(map[int32]T, 0)
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
		var mk int32
		var mv T
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = readValEnumNum(dec, mv, unquoteVal); err != nil {
			err = errorWrap(dec, err)
			return
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI32EnumStr read the next items from JSON contents as map key int32 and value enum with codec string.
func ReadMapI32EnumStr[T protoreflect.Enum](dec *Decoder, val map[int32]T,
	unquoteKey bool) (vv map[int32]T, err error) {
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
		vv = make(map[int32]T, 0)
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
		var mk int32
		var mv T
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
			err = errorWrap(dec, err)
			return nil, err
		}
		if mv, err = readValEnumStr(dec, mv); err != nil {
			err = errorWrap(dec, err)
			return
		}
		vv[mk] = mv
	}
	return vv, nil
}

// ReadMapI32Message read the next items from JSON contents as map key int32 and value message.
//func ReadMapI32Message[T PtrGoAny[U], U any](dec *Decoder,  val map[int32]T) (vv map[int32]T, err error) {
func ReadMapI32Message[U any, T *U](dec *Decoder, val map[int32]T,
	unquoteKey bool) (vv map[int32]T, err error) {
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
		vv = make(map[int32]T, 0)
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
		var mk int32
		var mv T
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTAnyObject read the next items from JSON contents as map key int32 and value WKT any with codec proto.
func ReadMapI32WKTAnyObject(dec *Decoder, val map[int32]*anypb.Any,
	unquoteKey bool) (vv map[int32]*anypb.Any, err error) {
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
		vv = make(map[int32]*anypb.Any, 0)
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
		var mk int32
		var mv *anypb.Any
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTAnyProto read the next items from JSON contents as map key int32 and value WKT any with codec proto.
func ReadMapI32WKTAnyProto(dec *Decoder, val map[int32]*anypb.Any,
	unquoteKey bool) (vv map[int32]*anypb.Any, err error) {
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
		vv = make(map[int32]*anypb.Any, 0)
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
		var mk int32
		var mv *anypb.Any
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurObject read the next items from JSON contents as map key int32 and value WKT duration with codec string.
func ReadMapI32WKTDurObject(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurTimeStr read the next items from JSON contents as map key int32 and value WKT duration with codec string.
func ReadMapI32WKTDurTimeStr(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurNano read the next items from JSON contents as map key int32 and value WKT duration with codec nanosecond.
func ReadMapI32WKTDurNano(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurMicro read the next items from JSON contents as map key int32 and value WKT duration with codec microsecond.
func ReadMapI32WKTDurMicro(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurMilli read the next items from JSON contents as map key int32 and value WKT duration with codec millisecond.
func ReadMapI32WKTDurMilli(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurSecond read the next items from JSON contents as map key int32 and value WKT duration with codec seconds.
func ReadMapI32WKTDurSecond(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurMinute read the next items from JSON contents as map key int32 and value WKT duration with codec minutes.
func ReadMapI32WKTDurMinute(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTDurHour read the next items from JSON contents as map key int32 and value WKT duration with codec hours.
func ReadMapI32WKTDurHour(dec *Decoder, val map[int32]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[int32]*durationpb.Duration, err error) {
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
		vv = make(map[int32]*durationpb.Duration, 0)
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
		var mk int32
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTTsObject read the next items from JSON contents as map key int32 and value WKT timestamp with codec time layout.
func ReadMapI32WKTTsObject(dec *Decoder, val map[int32]*timestamppb.Timestamp,
	unquoteKey bool) (vv map[int32]*timestamppb.Timestamp, err error) {
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
		vv = make(map[int32]*timestamppb.Timestamp, 0)
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
		var mk int32
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTTsLayout read the next items from JSON contents as map key int32 and value WKT timestamp with codec time layout.
func ReadMapI32WKTTsLayout(dec *Decoder, val map[int32]*timestamppb.Timestamp,
	unquoteKey bool, layout string) (vv map[int32]*timestamppb.Timestamp, err error) {
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
		vv = make(map[int32]*timestamppb.Timestamp, 0)
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
		var mk int32
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTTsUnixNano read the next items from JSON contents as map key int32 and value WKT timestamp with codec unix nano.
func ReadMapI32WKTTsUnixNano(dec *Decoder, val map[int32]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int32]*timestamppb.Timestamp, err error) {
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
		vv = make(map[int32]*timestamppb.Timestamp, 0)
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
		var mk int32
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTTsUnixMicro read the next items from JSON contents as map key int32 and value WKT timestamp with codec unix micro.
func ReadMapI32WKTTsUnixMicro(dec *Decoder, val map[int32]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int32]*timestamppb.Timestamp, err error) {
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
		vv = make(map[int32]*timestamppb.Timestamp, 0)
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
		var mk int32
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTTsUnixMilli read the next items from JSON contents as map key int32 and value WKT timestamp with codec unix milli.
func ReadMapI32WKTTsUnixMilli(dec *Decoder, val map[int32]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int32]*timestamppb.Timestamp, err error) {
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
		vv = make(map[int32]*timestamppb.Timestamp, 0)
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
		var mk int32
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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

// ReadMapI32WKTTsUnixSec read the next items from JSON contents as map key int32 and value WKT timestamp with codec unix sec.
func ReadMapI32WKTTsUnixSec(dec *Decoder, val map[int32]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[int32]*timestamppb.Timestamp, err error) {
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
		vv = make(map[int32]*timestamppb.Timestamp, 0)
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
		var mk int32
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyI32(unquoteKey); err != nil {
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
