package jsondecoder

import "fmt"

func (dec *Decoder) ReadPointerString(jsonKey string) (vv *string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 string
	if v1, err = bytesToString(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerBool(jsonKey string) (vv *bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 bool
	if v1, err = bytesToBool(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerInt32(jsonKey string) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = bytesToInt32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerInt64(jsonKey string) (vv *int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int64
	if v1, err = bytesToInt64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerUint32(jsonKey string) (vv *uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 uint32
	if v1, err = bytesToUint32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerUint64(jsonKey string) (vv *uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 uint64
	if v1, err = bytesToUint64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}

func (dec *Decoder) ReadPointerFloat32(jsonKey string) (vv *float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 float32
	if v1, err = bytesToFloat32(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerFloat64(jsonKey string) (vv *float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 float64
	if v1, err = bytesToFloat64(value); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}

func (dec *Decoder) ReadPointerEnumString(jsonKey string, em map[string]int32) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
func (dec *Decoder) ReadPointerEnumNumber(jsonKey string, em map[int32]string) (vv *int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	var v1 int32
	if v1, err = parseEnumNumber(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum1", string(value), jsonKey)
		return
	}
	vv = &v1
	return
}
