package jsondecoder

import "encoding/json"

func (dec *Decoder) readValI32(unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToInt32(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValI64(unquote bool) (vv int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToInt64(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValU32(unquote bool) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToUint32(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValU64(unquote bool) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToUint64(unquote, value); err != nil {
		return
	}
	return vv, nil
}

func (dec *Decoder) readValF32(unquote bool) (vv float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToFloat32(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValF64(unquote bool) (vv float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToFloat64(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValBool(unquote bool) (vv bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToBool(unquote, value); err != nil {
		return
	}
	return vv, nil
}

func (dec *Decoder) readValStr() (vv string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToString(value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValBytes() (vv []byte, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	if vv, err = convertToBytes(value); err != nil {
		return nil, err
	}
	return vv, nil
}
func (dec *Decoder) readValInterface(vv interface{}) (err error) {
	assertInterface(vv)

	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
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
		//err = errorWrap(dec, err)
		return
	}
	return
}

func (dec *Decoder) readValEnumNum(unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToEnum(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValEnumStr(em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}

	if vv, err = parseEnumString(value, em); err != nil {
		return
	}
	return vv, nil
}
