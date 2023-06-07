package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) ReadValueString(jsonKey string) (vv string, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueBool(jsonKey string) (vv bool, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToBool(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInt32(jsonKey string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToInt32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInt64(jsonKey string) (vv int64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueUint32(jsonKey string) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToUint32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueUint64(jsonKey string) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToUint64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(value), jsonKey)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadValueFloat32(jsonKey string) (vv float32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueFloat64(jsonKey string) (vv float64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueEnumNumber(jsonKey string, em map[int32]string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if vv, err = parseEnumNumber(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadValueBytes(jsonKey string) (vv []byte, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	if vv, err = bytesToBytes(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type []byte", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInterface(jsonKey string, initFN func() interface{}) (err error) {
	_ = jsonKey

	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil
	}
	if err = parseInterface(value, initFN); err != nil {
		return
	}
	return
}
