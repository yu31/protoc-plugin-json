package jsondecoder

func (dec *Decoder) ReadMapKeyString(jsonKey string) (vv string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = bytesToString(key); err != nil {
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyBool(jsonKey string, unquote bool) (vv bool, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = dec.convertToBool(jsonKey, key, unquote); err != nil {
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyInt32(jsonKey string, unquote bool) (vv int32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = dec.convertToInt32(jsonKey, key, unquote); err != nil {
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyInt64(jsonKey string, unquote bool) (vv int64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = dec.convertToInt64(jsonKey, key, unquote); err != nil {
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyUint32(jsonKey string, unquote bool) (vv uint32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = dec.convertToUint32(jsonKey, key, unquote); err != nil {
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyUint64(jsonKey string, unquote bool) (vv uint64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = dec.convertToUint64(jsonKey, key, unquote); err != nil {
		return
	}
	return
}
