package jsondecoder

import "google.golang.org/protobuf/reflect/protoreflect"

func (dec *Decoder) readPtrI32(unquote bool) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = convertToInt32(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) readPtrI64(unquote bool) (vv *int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int64
	if v1, err = convertToInt64(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) readPtrU32(unquote bool) (vv *uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 uint32
	if v1, err = convertToUint32(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) readPtrU64(unquote bool) (vv *uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 uint64
	if v1, err = convertToUint64(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}

func (dec *Decoder) readPtrF32(unquote bool) (vv *float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 float32
	if v1, err = convertToFloat32(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) readPtrF64(unquote bool) (vv *float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 float64
	if v1, err = convertToFloat64(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) readPtrBool(unquote bool) (vv *bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}

	var v1 bool
	if v1, err = convertToBool(unquote, value); err != nil {
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) readPtrStr() (vv *string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 string
	if v1, err = convertToString(value); err != nil {
		return
	}
	vv = &v1
	return
}
func readPtrEnumNum[T protoreflect.Enum](dec *Decoder, val *T, unquote bool) (vv *T, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = convertToEnum(unquote, value); err != nil {
		return
	}
	var tt T
	tt = tt.Type().New(protoreflect.EnumNumber(v1)).(T)
	vv = &tt
	return vv, nil
}
func readPtrEnumStr[T protoreflect.Enum](dec *Decoder, val *T) (vv *T, err error) {
	_ = val // Only used to confirm generic type.

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}

	var s1 string
	if s1, err = bytesToStringUnsafe(value); err != nil {
		return
	}

	var tt T
	//enumNumber := protoimpl.X.EnumDescriptorOf(val).Values().ByName(protoreflect.Name(s1)).Number()
	enumNumber := tt.Descriptor().Values().ByName(protoreflect.Name(s1)).Number()
	tt = tt.Type().New(enumNumber).(T)
	vv = &tt
	return vv, nil
}
