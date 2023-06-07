package jsondecoder

import "fmt"

func (dec *Decoder) ReadPointerString(jsonKey string) (vv *string, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv string
	if _vv, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerBool(jsonKey string) (vv *bool, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv bool
	if _vv, err = bytesToBool(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerInt32(jsonKey string) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv int32
	if _vv, err = bytesToInt32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerInt64(jsonKey string) (vv *int64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv int64
	if _vv, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerUint32(jsonKey string) (vv *uint32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv uint32
	if _vv, err = bytesToUint32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerUint64(jsonKey string) (vv *uint64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv uint64
	if _vv, err = bytesToUint64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}

func (dec *Decoder) ReadPointerFloat32(jsonKey string) (vv *float32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv float32
	if _vv, err = bytesToFloat32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerFloat64(jsonKey string) (vv *float64, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv float64
	if _vv, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}

func (dec *Decoder) ReadPointerEnumString(jsonKey string, em map[string]int32) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv int32
	if _vv, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
func (dec *Decoder) ReadPointerEnumNumber(jsonKey string, em map[int32]string) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readObjectValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var _vv int32
	if _vv, err = parseEnumNumber(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum1", string(value), jsonKey)
		return
	}
	vv = &_vv
	return
}
