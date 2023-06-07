package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) ReadMapValueString(jsonKey string) (vv string, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type string", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueBool(jsonKey string) (vv bool, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToBool(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type bool", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueInt32(jsonKey string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToInt32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type int32", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueInt64(jsonKey string) (vv int64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type int64", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueUint32(jsonKey string) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToUint32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type uint32", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueUint64(jsonKey string) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToUint64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type uint64", string(value), jsonKey)
		return
	}
	return
}

func (dec *Decoder) ReadMapValueFloat32(jsonKey string) (vv float32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type float32", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueFloat64(jsonKey string) (vv float64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type float64", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string as map value into field %s of type Enum", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueEnumNumber(jsonKey string, em map[int32]string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = parseEnumNumber(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number as map value into field %s of type Enum", string(value), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapValueBytes(jsonKey string) (vv []byte, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] != 'n' { // 'n' means null
		if vv, err = bytesToBytes(value); err != nil {
			err = fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type []byte", string(value), jsonKey)
			return
		}
	}
	return
}

// ReadMapValueInterface returns (isNULL,bool, error)
func (dec *Decoder) ReadMapValueInterface(jsonKey string, initFN func() interface{}) (err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] != 'n' { // 'n' means null
		if err = parseInterface(value, initFN); err != nil {
			return
		}
	}
	return
}
