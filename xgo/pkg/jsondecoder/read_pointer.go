package jsondecoder

func (dec *Decoder) ReadPointerString(jsonKey string) (vv *string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 string
	if v1, err = dec.convertToString(value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerBool(jsonKey string, unquote bool) (vv *bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}

	var v1 bool
	if v1, err = dec.convertToBool(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerInt32(jsonKey string, unquote bool) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = dec.convertToInt32(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerInt64(jsonKey string, unquote bool) (vv *int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int64
	if v1, err = dec.convertToInt64(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerUint32(jsonKey string, unquote bool) (vv *uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 uint32
	if v1, err = dec.convertToUint32(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerUint64(jsonKey string, unquote bool) (vv *uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 uint64
	if v1, err = dec.convertToUint64(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}

func (dec *Decoder) ReadPointerFloat32(jsonKey string, unquote bool) (vv *float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 float32
	if v1, err = dec.convertToFloat32(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerFloat64(jsonKey string, unquote bool) (vv *float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 float64
	if v1, err = dec.convertToFloat64(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerEnumNumber(jsonKey string, unquote bool) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = dec.convertToEnum(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}

	vv = &v1
	return
}
func (dec *Decoder) ReadPointerEnumString(jsonKey string, em map[string]int32) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}

	var v1 int32
	if v1, err = parseEnumString(value, em); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	vv = &v1
	return
}
