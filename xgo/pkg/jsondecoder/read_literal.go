package jsondecoder

import (
	"encoding/json"
	"fmt"
)

func (dec *Decoder) ReadLiteralString(jsonKey string) (vv string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralBool(jsonKey string) (vv bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToBool(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralInt32(jsonKey string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToInt32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralInt64(jsonKey string) (vv int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralUint32(jsonKey string) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToUint32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralUint64(jsonKey string) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToUint64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(value), jsonKey)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadLiteralFloat32(jsonKey string) (vv float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralFloat64(jsonKey string) (vv float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralEnumNumber(jsonKey string, em map[int32]string) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = parseEnumNumber(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadLiteralBytes(jsonKey string) (vv []byte, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
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

func (dec *Decoder) ReadLiteralInterface(jsonKey string, vv interface{}) (err error) {
	assertInterface(vv)

	_ = jsonKey
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
		return
	}
	return
}
