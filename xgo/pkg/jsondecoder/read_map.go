package jsondecoder

func (dec *Decoder) ReadMapKeyString(jsonKey string) (vv string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToString(key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyBool(jsonKey string, unquote bool) (vv bool, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToBool(unquote, key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyInt32(jsonKey string, unquote bool) (vv int32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToInt32(unquote, key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyInt64(jsonKey string, unquote bool) (vv int64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToInt64(unquote, key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyUint32(jsonKey string, unquote bool) (vv uint32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToUint32(unquote, key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyUint64(jsonKey string, unquote bool) (vv uint64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToUint64(unquote, key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
