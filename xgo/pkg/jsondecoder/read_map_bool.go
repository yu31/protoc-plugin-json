package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReadMapBoolI32 read the next items from JSON contents as map key bool and value int32.
func ReadMapBoolI32(dec *Decoder, val map[bool]int32,
	unquoteKey, unquoteVal bool) (vv map[bool]int32, err error) {
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
		vv = make(map[bool]int32, 0)
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
		var mk bool
		var mv int32
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolI64 read the next items from JSON contents as map key bool and value int64.
func ReadMapBoolI64(dec *Decoder, val map[bool]int64,
	unquoteKey, unquoteVal bool) (vv map[bool]int64, err error) {
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
		vv = make(map[bool]int64, 0)
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
		var mk bool
		var mv int64
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolU32 read the next items from JSON contents as map key bool and value uint32.
func ReadMapBoolU32(dec *Decoder, val map[bool]uint32,
	unquoteKey, unquoteVal bool) (vv map[bool]uint32, err error) {
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
		vv = make(map[bool]uint32, 0)
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
		var mk bool
		var mv uint32
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolU64 read the next items from JSON contents as map key bool and value uint64.
func ReadMapBoolU64(dec *Decoder, val map[bool]uint64,
	unquoteKey, unquoteVal bool) (vv map[bool]uint64, err error) {
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
		vv = make(map[bool]uint64, 0)
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
		var mk bool
		var mv uint64
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolF32 read the next items from JSON contents as map key bool and value float32.
func ReadMapBoolF32(dec *Decoder, val map[bool]float32,
	unquoteKey, unquoteVal bool) (vv map[bool]float32, err error) {
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
		vv = make(map[bool]float32, 0)
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
		var mk bool
		var mv float32
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolF64 read the next items from JSON contents as map key bool and value uint64.
func ReadMapBoolF64(dec *Decoder, val map[bool]float64,
	unquoteKey, unquoteVal bool) (vv map[bool]float64, err error) {
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
		vv = make(map[bool]float64, 0)
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
		var mk bool
		var mv float64
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolBool read the next items from JSON contents as map key bool and value bool.
func ReadMapBoolBool(dec *Decoder, val map[bool]bool,
	unquoteKey, unquoteVal bool) (vv map[bool]bool, err error) {
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
		vv = make(map[bool]bool, 0)
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
		var mk bool
		var mv bool
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolStr read the next items from JSON contents as map key bool and value string.
func ReadMapBoolStr(dec *Decoder, val map[bool]string,
	unquoteKey bool) (vv map[bool]string, err error) {
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
		vv = make(map[bool]string, 0)
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
		var mk bool
		var mv string
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolBytes read the next items from JSON contents as map key bool and value bytes.
func ReadMapBoolBytes(dec *Decoder, val map[bool][]byte,
	unquoteKey bool) (vv map[bool][]byte, err error) {
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
		vv = make(map[bool][]byte, 0)
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
		var mk bool
		var mv []byte
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolEnumNum read the next items from JSON contents as map key bool and value enum with codec number.
func ReadMapBoolEnumNum[T protoreflect.Enum](dec *Decoder, val map[bool]T,
	unquoteKey, unquoteVal bool) (vv map[bool]T, err error) {
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
		vv = make(map[bool]T, 0)
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
		var mk bool
		var mv T
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolEnumStr read the next items from JSON contents as map key bool and value enum with codec string.
func ReadMapBoolEnumStr[T protoreflect.Enum](dec *Decoder, val map[bool]T,
	unquoteKey bool) (vv map[bool]T, err error) {
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
		vv = make(map[bool]T, 0)
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
		var mk bool
		var mv T
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolMessage read the next items from JSON contents as map key bool and value message.
// func ReadMapBoolMessage[T PtrGoAny[U], U any](dec *Decoder,  val map[bool]T) (vv map[bool]T, err error) {
func ReadMapBoolMessage[U any, T *U](dec *Decoder, val map[bool]T,
	unquoteKey bool) (vv map[bool]T, err error) {
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
		vv = make(map[bool]T, 0)
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
		var mk bool
		var mv T
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTAnyObject read the next items from JSON contents as map key bool and value WKT any with codec proto.
func ReadMapBoolWKTAnyObject(dec *Decoder, val map[bool]*anypb.Any,
	unquoteKey bool) (vv map[bool]*anypb.Any, err error) {
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
		vv = make(map[bool]*anypb.Any, 0)
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
		var mk bool
		var mv *anypb.Any
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTAnyProto read the next items from JSON contents as map key bool and value WKT any with codec proto.
func ReadMapBoolWKTAnyProto(dec *Decoder, val map[bool]*anypb.Any,
	unquoteKey bool) (vv map[bool]*anypb.Any, err error) {
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
		vv = make(map[bool]*anypb.Any, 0)
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
		var mk bool
		var mv *anypb.Any
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurObject read the next items from JSON contents as map key bool and value WKT duration with codec string.
func ReadMapBoolWKTDurObject(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurTimeStr read the next items from JSON contents as map key bool and value WKT duration with codec string.
func ReadMapBoolWKTDurTimeStr(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurNano read the next items from JSON contents as map key bool and value WKT duration with codec nanosecond.
func ReadMapBoolWKTDurNano(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurMicro read the next items from JSON contents as map key bool and value WKT duration with codec microsecond.
func ReadMapBoolWKTDurMicro(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurMilli read the next items from JSON contents as map key bool and value WKT duration with codec millisecond.
func ReadMapBoolWKTDurMilli(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurSecond read the next items from JSON contents as map key bool and value WKT duration with codec seconds.
func ReadMapBoolWKTDurSecond(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurMinute read the next items from JSON contents as map key bool and value WKT duration with codec minutes.
func ReadMapBoolWKTDurMinute(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTDurHour read the next items from JSON contents as map key bool and value WKT duration with codec hours.
func ReadMapBoolWKTDurHour(dec *Decoder, val map[bool]*durationpb.Duration,
	unquoteKey, unquoteVal bool) (vv map[bool]*durationpb.Duration, err error) {
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
		vv = make(map[bool]*durationpb.Duration, 0)
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
		var mk bool
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTTsObject read the next items from JSON contents as map key bool and value WKT timestamp with codec time layout.
func ReadMapBoolWKTTsObject(dec *Decoder, val map[bool]*timestamppb.Timestamp,
	unquoteKey bool) (vv map[bool]*timestamppb.Timestamp, err error) {
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
		vv = make(map[bool]*timestamppb.Timestamp, 0)
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
		var mk bool
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTTsLayout read the next items from JSON contents as map key bool and value WKT timestamp with codec time layout.
func ReadMapBoolWKTTsLayout(dec *Decoder, val map[bool]*timestamppb.Timestamp,
	unquoteKey bool, layout string) (vv map[bool]*timestamppb.Timestamp, err error) {
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
		vv = make(map[bool]*timestamppb.Timestamp, 0)
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
		var mk bool
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTTsUnixNano read the next items from JSON contents as map key bool and value WKT timestamp with codec unix nano.
func ReadMapBoolWKTTsUnixNano(dec *Decoder, val map[bool]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[bool]*timestamppb.Timestamp, err error) {
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
		vv = make(map[bool]*timestamppb.Timestamp, 0)
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
		var mk bool
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTTsUnixMicro read the next items from JSON contents as map key bool and value WKT timestamp with codec unix micro.
func ReadMapBoolWKTTsUnixMicro(dec *Decoder, val map[bool]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[bool]*timestamppb.Timestamp, err error) {
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
		vv = make(map[bool]*timestamppb.Timestamp, 0)
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
		var mk bool
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTTsUnixMilli read the next items from JSON contents as map key bool and value WKT timestamp with codec unix milli.
func ReadMapBoolWKTTsUnixMilli(dec *Decoder, val map[bool]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[bool]*timestamppb.Timestamp, err error) {
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
		vv = make(map[bool]*timestamppb.Timestamp, 0)
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
		var mk bool
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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

// ReadMapBoolWKTTsUnixSec read the next items from JSON contents as map key bool and value WKT timestamp with codec unix sec.
func ReadMapBoolWKTTsUnixSec(dec *Decoder, val map[bool]*timestamppb.Timestamp,
	unquoteKey, unquoteVal bool) (vv map[bool]*timestamppb.Timestamp, err error) {
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
		vv = make(map[bool]*timestamppb.Timestamp, 0)
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
		var mk bool
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyBool(unquoteKey); err != nil {
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
