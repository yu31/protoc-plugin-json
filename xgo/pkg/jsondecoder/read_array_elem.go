package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) ReadArrayElemString(jsonKey string) (vv string, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToString(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type string", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemBool(jsonKey string) (vv bool, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToBool(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type bool", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemInt32(jsonKey string) (vv int32, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToInt32(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type int32", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemInt64(jsonKey string) (vv int64, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToInt64(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type int64", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemUint32(jsonKey string) (vv uint32, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToUint32(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type uint32", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemUint64(jsonKey string) (vv uint64, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToUint64(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type uint64", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemFloat32(jsonKey string) (vv float32, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToFloat32(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type float32", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemFloat64(jsonKey string) (vv float64, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = bytesToFloat64(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type float64", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = parseEnumString(elem, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string as array element into field %s of type Enum", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemEnumNumber(jsonKey string, em map[int32]string) (vv int32, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if vv, err = parseEnumNumber(elem, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as number as array element into field %s of type Enum", string(elem), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemBytes(jsonKey string) (vv []byte, err error) {
	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if elem[0] == 'n' { // 'n' means null
		return nil, nil
	}
	if vv, err = bytesToBytes(elem); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type []byte", string(elem), jsonKey)
		return nil, err
	}
	return vv, nil
}

// ReadArrayElemInterface returns (isNULL, noMore bool, err error)
func (dec *Decoder) ReadArrayElemInterface(jsonKey string, initFN func() interface{}) (err error) {
	_ = jsonKey

	var elem []byte
	if elem, err = dec.readArrayElem(); err != nil {
		return
	}
	if elem[0] == 'n' { // 'n' means null
		return nil
	}
	if err = parseInterface(elem, initFN); err != nil {
		return
	}
	return nil
}
