package jsondecoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ReadPtrI32 read the next items from JSON contents as value of int32.
func ReadPtrI32(dec *Decoder, val *int32, unquote bool) (vv *int32, err error) {
	if vv, err = dec.readPtrI32(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrI64 read the next items from JSON contents as value of int64.
func ReadPtrI64(dec *Decoder, val *int64, unquote bool) (vv *int64, err error) {
	if vv, err = dec.readPtrI64(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrU32 read the next items from JSON contents as value of uint32.
func ReadPtrU32(dec *Decoder, val *uint32, unquote bool) (vv *uint32, err error) {
	if vv, err = dec.readPtrU32(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrU64 read the next items from JSON contents as value of uint64.
func ReadPtrU64(dec *Decoder, val *uint64, unquote bool) (vv *uint64, err error) {
	if vv, err = dec.readPtrU64(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrF32 read the next items from JSON contents as value of float32.
func ReadPtrF32(dec *Decoder, val *float32, unquote bool) (vv *float32, err error) {
	if vv, err = dec.readPtrF32(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrF64 read the next items from JSON contents as value of uint64.
func ReadPtrF64(dec *Decoder, val *float64, unquote bool) (vv *float64, err error) {
	if vv, err = dec.readPtrF64(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrBool read the next items from JSON contents as value of bool.
func ReadPtrBool(dec *Decoder, val *bool, unquote bool) (vv *bool, err error) {
	if vv, err = dec.readPtrBool(val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrStr read the next items from JSON contents as value of string.
func ReadPtrStr(dec *Decoder, val *string) (vv *string, err error) {
	if vv, err = dec.readPtrStr(val); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrEnumNum read the next items from JSON contents as value of enum with codec number.
func ReadPtrEnumNum[T protoreflect.Enum](dec *Decoder, val *T, unquote bool) (vv *T, err error) {
	if vv, err = readPtrEnumNum(dec, val, unquote); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return vv, nil
}

// ReadPtrEnumStr read the next items from JSON contents as value of enum with codec string.
func ReadPtrEnumStr[T protoreflect.Enum](dec *Decoder, val *T) (vv *T, err error) {
	if vv, err = readPtrEnumStr(dec, val); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return
}
