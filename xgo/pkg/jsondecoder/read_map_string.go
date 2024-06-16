package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReadMapStrI32 read the next items from JSON contents as map key string and value int32.
func ReadMapStrI32(dec *Decoder, val map[string]int32,
	unquoteVal bool) (vv map[string]int32, err error) {
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
		vv = make(map[string]int32, 0)
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
		var mk string
		var mv int32
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrI64 read the next items from JSON contents as map key string and value int64.
func ReadMapStrI64(dec *Decoder, val map[string]int64,
	unquoteVal bool) (vv map[string]int64, err error) {
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
		vv = make(map[string]int64, 0)
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
		var mk string
		var mv int64
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrU32 read the next items from JSON contents as map key string and value uint32.
func ReadMapStrU32(dec *Decoder, val map[string]uint32,
	unquoteVal bool) (vv map[string]uint32, err error) {
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
		vv = make(map[string]uint32, 0)
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
		var mk string
		var mv uint32
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrU64 read the next items from JSON contents as map key string and value uint64.
func ReadMapStrU64(dec *Decoder, val map[string]uint64,
	unquoteVal bool) (vv map[string]uint64, err error) {
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
		vv = make(map[string]uint64, 0)
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
		var mk string
		var mv uint64
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrF32 read the next items from JSON contents as map key string and value float32.
func ReadMapStrF32(dec *Decoder, val map[string]float32,
	unquoteVal bool) (vv map[string]float32, err error) {
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
		vv = make(map[string]float32, 0)
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
		var mk string
		var mv float32
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrF64 read the next items from JSON contents as map key string and value uint64.
func ReadMapStrF64(dec *Decoder, val map[string]float64,
	unquoteVal bool) (vv map[string]float64, err error) {
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
		vv = make(map[string]float64, 0)
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
		var mk string
		var mv float64
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrBool read the next items from JSON contents as map key string and value bool.
func ReadMapStrBool(dec *Decoder, val map[string]bool,
	unquoteVal bool) (vv map[string]bool, err error) {
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
		vv = make(map[string]bool, 0)
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
		var mk string
		var mv bool
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrStr read the next items from JSON contents as map key string and value string.
func ReadMapStrStr(dec *Decoder, val map[string]string) (vv map[string]string, err error) {
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
		vv = make(map[string]string, 0)
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
		var mk string
		var mv string
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrBytes read the next items from JSON contents as map key string and value bytes.
func ReadMapStrBytes(dec *Decoder, val map[string][]byte) (vv map[string][]byte, err error) {
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
		vv = make(map[string][]byte, 0)
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
		var mk string
		var mv []byte
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrEnumNum read the next items from JSON contents as map key string and value enum with codec number.
func ReadMapStrEnumNum[T protoreflect.Enum](dec *Decoder, val map[string]T,
	unquoteVal bool) (vv map[string]T, err error) {
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
		vv = make(map[string]T, 0)
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
		var mk string
		var mv T
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrEnumStr read the next items from JSON contents as map key string and value enum with codec string.
func ReadMapStrEnumStr[T protoreflect.Enum](dec *Decoder, val map[string]T) (vv map[string]T, err error) {
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
		vv = make(map[string]T, 0)
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
		var mk string
		var mv T
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrMessage read the next items from JSON contents as map key string and value message.
//func ReadMapStrMessage[T PtrGoAny[U], U any](dec *Decoder,  val map[string]T) (vv map[string]T, err error) {
func ReadMapStrMessage[U any, T *U](dec *Decoder, val map[string]T,
) (vv map[string]T, err error) {
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
		vv = make(map[string]T, 0)
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
		var mk string
		var mv T
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTAnyObject read the next items from JSON contents as map key string and value WKT any with codec proto.
func ReadMapStrWKTAnyObject(dec *Decoder, val map[string]*anypb.Any,
) (vv map[string]*anypb.Any, err error) {
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
		vv = make(map[string]*anypb.Any, 0)
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
		var mk string
		var mv *anypb.Any
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTAnyProto read the next items from JSON contents as map key string and value WKT any with codec proto.
func ReadMapStrWKTAnyProto(dec *Decoder, val map[string]*anypb.Any,
) (vv map[string]*anypb.Any, err error) {
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
		vv = make(map[string]*anypb.Any, 0)
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
		var mk string
		var mv *anypb.Any
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurObject read the next items from JSON contents as map key string and value WKT duration with codec string.
func ReadMapStrWKTDurObject(dec *Decoder, val map[string]*durationpb.Duration,
) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurTimeStr read the next items from JSON contents as map key string and value WKT duration with codec string.
func ReadMapStrWKTDurTimeStr(dec *Decoder, val map[string]*durationpb.Duration,
) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurNano read the next items from JSON contents as map key string and value WKT duration with codec nanosecond.
func ReadMapStrWKTDurNano(dec *Decoder, val map[string]*durationpb.Duration,
	unquoteVal bool) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurMicro read the next items from JSON contents as map key string and value WKT duration with codec microsecond.
func ReadMapStrWKTDurMicro(dec *Decoder, val map[string]*durationpb.Duration,
	unquoteVal bool) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurMilli read the next items from JSON contents as map key string and value WKT duration with codec millisecond.
func ReadMapStrWKTDurMilli(dec *Decoder, val map[string]*durationpb.Duration,
	unquoteVal bool) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurSecond read the next items from JSON contents as map key string and value WKT duration with codec seconds.
func ReadMapStrWKTDurSecond(dec *Decoder, val map[string]*durationpb.Duration,
	unquoteVal bool) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurMinute read the next items from JSON contents as map key string and value WKT duration with codec minutes.
func ReadMapStrWKTDurMinute(dec *Decoder, val map[string]*durationpb.Duration,
	unquoteVal bool) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTDurHour read the next items from JSON contents as map key string and value WKT duration with codec hours.
func ReadMapStrWKTDurHour(dec *Decoder, val map[string]*durationpb.Duration,
	unquoteVal bool) (vv map[string]*durationpb.Duration, err error) {
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
		vv = make(map[string]*durationpb.Duration, 0)
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
		var mk string
		var mv *durationpb.Duration
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTTsObject read the next items from JSON contents as map key string and value WKT timestamp with codec time layout.
func ReadMapStrWKTTsObject(dec *Decoder, val map[string]*timestamppb.Timestamp,
) (vv map[string]*timestamppb.Timestamp, err error) {
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
		vv = make(map[string]*timestamppb.Timestamp, 0)
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
		var mk string
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTTsLayout read the next items from JSON contents as map key string and value WKT timestamp with codec time layout.
func ReadMapStrWKTTsLayout(dec *Decoder, val map[string]*timestamppb.Timestamp,
	layout string) (vv map[string]*timestamppb.Timestamp, err error) {
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
		vv = make(map[string]*timestamppb.Timestamp, 0)
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
		var mk string
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTTsUnixNano read the next items from JSON contents as map key string and value WKT timestamp with codec unix nano.
func ReadMapStrWKTTsUnixNano(dec *Decoder, val map[string]*timestamppb.Timestamp,
	unquoteVal bool) (vv map[string]*timestamppb.Timestamp, err error) {
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
		vv = make(map[string]*timestamppb.Timestamp, 0)
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
		var mk string
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTTsUnixMicro read the next items from JSON contents as map key string and value WKT timestamp with codec unix micro.
func ReadMapStrWKTTsUnixMicro(dec *Decoder, val map[string]*timestamppb.Timestamp,
	unquoteVal bool) (vv map[string]*timestamppb.Timestamp, err error) {
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
		vv = make(map[string]*timestamppb.Timestamp, 0)
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
		var mk string
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTTsUnixMilli read the next items from JSON contents as map key string and value WKT timestamp with codec unix milli.
func ReadMapStrWKTTsUnixMilli(dec *Decoder, val map[string]*timestamppb.Timestamp,
	unquoteVal bool) (vv map[string]*timestamppb.Timestamp, err error) {
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
		vv = make(map[string]*timestamppb.Timestamp, 0)
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
		var mk string
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyStr(); err != nil {
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

// ReadMapStrWKTTsUnixSec read the next items from JSON contents as map key string and value WKT timestamp with codec unix sec.
func ReadMapStrWKTTsUnixSec(dec *Decoder, val map[string]*timestamppb.Timestamp,
	unquoteVal bool) (vv map[string]*timestamppb.Timestamp, err error) {
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
		vv = make(map[string]*timestamppb.Timestamp, 0)
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
		var mk string
		var mv *timestamppb.Timestamp
		if mk, err = dec.readMapKeyStr(); err != nil {
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
