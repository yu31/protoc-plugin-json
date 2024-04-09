package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ReadPtrI32 read the next items from JSON contents as value of int32.
func ReadPtrI32(dec *Decoder, unquote bool) (vv *int32, err error) {
	if vv, err = dec.readPtrI32(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrI64 read the next items from JSON contents as value of int64.
func ReadPtrI64(dec *Decoder, unquote bool) (vv *int64, err error) {
	if vv, err = dec.readPtrI64(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrU32 read the next items from JSON contents as value of uint32.
func ReadPtrU32(dec *Decoder, unquote bool) (vv *uint32, err error) {
	if vv, err = dec.readPtrU32(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrU64 read the next items from JSON contents as value of uint64.
func ReadPtrU64(dec *Decoder, unquote bool) (vv *uint64, err error) {
	if vv, err = dec.readPtrU64(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrF32 read the next items from JSON contents as value of float32.
func ReadPtrF32(dec *Decoder, unquote bool) (vv *float32, err error) {
	if vv, err = dec.readPtrF32(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrF64 read the next items from JSON contents as value of uint64.
func ReadPtrF64(dec *Decoder, unquote bool) (vv *float64, err error) {
	if vv, err = dec.readPtrF64(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrBool read the next items from JSON contents as value of bool.
func ReadPtrBool(dec *Decoder, unquote bool) (vv *bool, err error) {
	if vv, err = dec.readPtrBool(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrStr read the next items from JSON contents as value of string.
func ReadPtrStr(dec *Decoder) (vv *string, err error) {
	if vv, err = dec.readPtrStr(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrEnumNum read the next items from JSON contents as value of enum with codec number.
// FIXME: optimized the codes.
func ReadPtrEnumNum[T protoreflect.Enum](dec *Decoder, val *T, unquote bool) (vv *T, err error) {
	_ = val // Only used to confirm generic type.
	var v2 *int32
	if v2, err = dec.readPtrEnumNum(unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if v2 == nil {
		return nil, err
	}
	var tt T
	tt = tt.Type().New(protoreflect.EnumNumber(*v2)).(T)
	vv = &tt
	return vv, nil
}

// ReadPtrEnumStr read the next items from JSON contents as value of enum with codec string.
// FIXME: optimized the codes.
func ReadPtrEnumStr[T protoreflect.Enum](dec *Decoder, val *T, em map[string]int32) (vv *T, err error) {
	_ = val // Only used to confirm generic type.
	var v2 *int32
	if v2, err = dec.readPtrEnumStr(em); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if v2 == nil {
		return nil, err
	}
	var tt T
	tt = tt.Type().New(protoreflect.EnumNumber(*v2)).(T)
	vv = &tt
	return vv, nil
}
