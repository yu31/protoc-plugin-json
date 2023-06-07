package jsondecoder

import (
	"fmt"
	"strconv"
)

func (dec *Decoder) ReadMapKeyString(jsonKey string) (string, error) {
	item := dec.readObjectKey()
	return item, nil
	//vv, err := parseString(item)
	//if err != nil {
	//	return "", fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(item), jsonKey)
	//}
	//return vv, nil
}
func (dec *Decoder) ReadMapKeyBool(jsonKey string) (bool, error) {
	item := dec.readObjectKey()
	vv, err := strconv.ParseBool(item)
	if err != nil {
		return false, fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadMapKeyInt32(jsonKey string) (int32, error) {
	item := dec.readObjectKey()
	vv, err := strconv.ParseInt(item, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(item), jsonKey)
	}
	return int32(vv), nil
}
func (dec *Decoder) ReadMapKeyInt64(jsonKey string) (int64, error) {
	item := dec.readObjectKey()
	vv, err := strconv.ParseInt(item, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadMapKeyUint32(jsonKey string) (uint32, error) {
	item := dec.readObjectKey()
	vv, err := strconv.ParseUint(item, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(item), jsonKey)
	}
	return uint32(vv), nil
}
func (dec *Decoder) ReadMapKeyUint64(jsonKey string) (uint64, error) {
	item := dec.readObjectKey()
	vv, err := strconv.ParseUint(item, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(item), jsonKey)
	}
	return vv, nil
}
