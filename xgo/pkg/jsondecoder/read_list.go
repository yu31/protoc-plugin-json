package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ReadListI32 read the next items from JSON contents as a list of int32.
func ReadListI32(dec *Decoder, val []int32, unquote bool) (vv []int32, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]int32, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 int32
		if v1, err = dec.readValI32(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListI64 read the next items from JSON contents as a list of int64.
func ReadListI64(dec *Decoder, val []int64, unquote bool) (vv []int64, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]int64, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 int64
		if v1, err = dec.readValI64(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListU32 read the next items from JSON contents as a list of uint32.
func ReadListU32(dec *Decoder, val []uint32, unquote bool) (vv []uint32, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]uint32, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 uint32
		if v1, err = dec.readValU32(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListU64 read the next items from JSON contents as a list of uint64.
func ReadListU64(dec *Decoder, val []uint64, unquote bool) (vv []uint64, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]uint64, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 uint64
		if v1, err = dec.readValU64(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListF32 read the next items from JSON contents as a list of float32.
func ReadListF32(dec *Decoder, val []float32, unquote bool) (vv []float32, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]float32, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 float32
		if v1, err = dec.readValF32(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListF64 read the next items from JSON contents as a list of uint64.
func ReadListF64(dec *Decoder, val []float64, unquote bool) (vv []float64, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]float64, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 float64
		if v1, err = dec.readValF64(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListBool read the next items from JSON contents as a list of bool.
func ReadListBool(dec *Decoder, val []bool, unquote bool) (vv []bool, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]bool, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 bool
		if v1, err = dec.readValBool(unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListStr read the next items from JSON contents as a list of string.
func ReadListStr(dec *Decoder, val []string) (vv []string, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]string, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 string
		if v1, err = dec.readValStr(); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListBytes read the next items from JSON contents as a list of bytes.
func ReadListBytes(dec *Decoder, val [][]byte) (vv [][]byte, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([][]byte, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 []byte
		if v1, err = dec.readValBytes(); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListEnumNum read the next items from JSON contents as a list of enum with codec number.
func ReadListEnumNum[T protoreflect.Enum](dec *Decoder, val []T, unquote bool) (vv []T, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}

	if val == nil {
		vv = make([]T, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		var v1 T
		if v1, err = readValEnumNum(dec, v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListEnumStr read the next items from JSON contents as a list of enum with codec string.
func ReadListEnumStr[T protoreflect.Enum](dec *Decoder, val []T) (vv []T, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}

	if val == nil {
		vv = make([]T, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}

		var v1 T
		if v1, err = readValEnumStr(dec, v1); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListMessage read the next items from JSON contents as a list of message.
// func ReadListMessage[T PtrGoAny[U], U any](dec *Decoder,  val []T) (vv []T, err error) {
func ReadListMessage[U any, T *U](dec *Decoder, val []T) (vv []T, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]T, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		var v1 T
		if !isNULL {
			if i < len(vv) {
				v1 = vv[i]
			}
			if v1 == nil {
				//var v2 U
				v1 = new(U)
			}
			if err = dec.readValInterface(v1); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTAnyObject read the next items from JSON contents as a list of WKT any with codec proto.
func ReadListWKTAnyObject(dec *Decoder, val []*anypb.Any) (vv []*anypb.Any, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*anypb.Any, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *anypb.Any
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if i < len(vv) {
				v1 = vv[i]
			}
			if v1 == nil {
				v1 = new(anypb.Any)
			}
			if err = dec.readValInterface(v1); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTAnyProto read the next items from JSON contents as a list of WKT any with codec proto.
func ReadListWKTAnyProto(dec *Decoder, val []*anypb.Any) (vv []*anypb.Any, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*anypb.Any, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *anypb.Any
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if i < len(vv) {
				v1 = vv[i]
			}
			if v1 == nil {
				v1 = new(anypb.Any)
			}
			if err = dec.readValWKTAnyProto(v1); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurObject read the next items from JSON contents as a list of WKT duration with codec string.
func ReadListWKTDurObject(dec *Decoder, val []*durationpb.Duration) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if i < len(vv) {
				v1 = vv[i]
			}
			if v1 == nil {
				v1 = new(durationpb.Duration)
			}

			if err = dec.readValInterface(v1); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurTimeStr read the next items from JSON contents as a list of WKT duration with codec string.
func ReadListWKTDurTimeStr(dec *Decoder, val []*durationpb.Duration) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurTimeStr(v1); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurNano read the next items from JSON contents as a list of WKT duration with codec nanosecond.
func ReadListWKTDurNano(dec *Decoder, val []*durationpb.Duration, unquote bool) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurNano(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurMicro read the next items from JSON contents as a list of WKT duration with codec microsecond.
func ReadListWKTDurMicro(dec *Decoder, val []*durationpb.Duration, unquote bool) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurMicro(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurMilli read the next items from JSON contents as a list of WKT duration with codec millisecond.
func ReadListWKTDurMilli(dec *Decoder, val []*durationpb.Duration, unquote bool) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurMilli(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurSecond read the next items from JSON contents as a list of WKT duration with codec seconds.
func ReadListWKTDurSecond(dec *Decoder, val []*durationpb.Duration, unquote bool) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurSecond(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurMinute read the next items from JSON contents as a list of WKT duration with codec minutes.
func ReadListWKTDurMinute(dec *Decoder, val []*durationpb.Duration, unquote bool) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurMinute(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTDurHour read the next items from JSON contents as a list of WKT duration with codec hours.
func ReadListWKTDurHour(dec *Decoder, val []*durationpb.Duration, unquote bool) (vv []*durationpb.Duration, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*durationpb.Duration, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *durationpb.Duration
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(durationpb.Duration)
		}
		if err = dec.readValWKTDurHour(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTTsObject read the next items from JSON contents as a list of WKT timestamp with codec time layout.
func ReadListWKTTsObject(dec *Decoder, val []*timestamppb.Timestamp) (vv []*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *timestamppb.Timestamp
		if isNULL, err = dec.NextLiteralIsNULL(); err != nil {
			return
		}
		if !isNULL {
			if i < len(vv) {
				v1 = vv[i]
			}
			if v1 == nil {
				v1 = new(timestamppb.Timestamp)
			}
			if err = dec.readValInterface(v1); err != nil {
				err = errorWrap(dec, err)
				return
			}
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTTsLayout read the next items from JSON contents as a list of WKT timestamp with codec time layout.
func ReadListWKTTsLayout(dec *Decoder, val []*timestamppb.Timestamp, layout string) (vv []*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *timestamppb.Timestamp
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsLayout(v1, layout); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTTsUnixNano read the next items from JSON contents as a list of WKT timestamp with codec unix nano.
func ReadListWKTTsUnixNano(dec *Decoder, val []*timestamppb.Timestamp, unquote bool) (vv []*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *timestamppb.Timestamp
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixNano(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTTsUnixMicro read the next items from JSON contents as a list of WKT timestamp with codec unix micro.
func ReadListWKTTsUnixMicro(dec *Decoder, val []*timestamppb.Timestamp, unquote bool) (vv []*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *timestamppb.Timestamp
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixMicro(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTTsUnixMilli read the next items from JSON contents as a list of WKT timestamp with codec unix milli.
func ReadListWKTTsUnixMilli(dec *Decoder, val []*timestamppb.Timestamp, unquote bool) (vv []*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *timestamppb.Timestamp
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixMilli(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}

// ReadListWKTTsUnixSec read the next items from JSON contents as a list of WKT timestamp with codec unix sec.
func ReadListWKTTsUnixSec(dec *Decoder, val []*timestamppb.Timestamp, unquote bool) (vv []*timestamppb.Timestamp, err error) {
	var (
		isEnd  bool
		isNULL bool
	)
	if isNULL, err = dec.BeforeReadList(); err != nil {
		return nil, err
	}
	if isNULL { // The value should be set to nil.
		return nil, nil
	}
	if val == nil {
		vv = make([]*timestamppb.Timestamp, 0)
	} else {
		vv = val
	}

	i := 0
	for ; ; i++ { // Loop to read the list items
		if isEnd, err = dec.BeforeReadNext(); err != nil {
			return
		}
		if isEnd {
			break
		}
		var v1 *timestamppb.Timestamp
		if i < len(vv) {
			v1 = vv[i]
		}
		if v1 == nil {
			v1 = new(timestamppb.Timestamp)
		}
		if err = dec.readValWKTTsUnixSec(v1, unquote); err != nil {
			err = errorWrap(dec, err)
			return
		}
		if i >= len(vv) {
			vv = append(vv, v1) // Set the value
		} else {
			vv[i] = v1 // Set the value
		}
	}
	// Truncate the slice to consistent with standard library json
	return vv[:i], nil
}
