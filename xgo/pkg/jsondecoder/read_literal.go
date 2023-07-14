package jsondecoder

import "encoding/json"

func (dec *Decoder) ReadLiteralInterface(jsonKey string, vv interface{}) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}

	if value[0] == 'n' { // 'n' means null
		return nil
	}
	if um, ok := vv.(json.Unmarshaler); ok {
		err = um.UnmarshalJSON(value)
	} else {
		err = json.Unmarshal(value, vv)
	}
	if err != nil {
		//err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}

func (dec *Decoder) ReadLiteralBytes(jsonKey string) (vv []byte, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	if vv, err = dec.convertToBytes(value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return nil, err
	}
	return vv, nil
}

func (dec *Decoder) ReadLiteralString(jsonKey string) (vv string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToString(value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralBool(jsonKey string, unquote bool) (vv bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToBool(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralInt32(jsonKey string, unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToInt32(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralInt64(jsonKey string, unquote bool) (vv int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToInt64(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralUint32(jsonKey string, unquote bool) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToUint32(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralUint64(jsonKey string, unquote bool) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToUint64(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadLiteralFloat32(jsonKey string, unquote bool) (vv float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToFloat32(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralFloat64(jsonKey string, unquote bool) (vv float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToFloat64(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralEnumNumber(jsonKey string, unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if vv, err = dec.convertToEnum(unquote, value); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}

	if vv, err = parseEnumString(value, em); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return vv, nil
}
