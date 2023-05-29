package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) ReadArrayElemString(jsonKey string) (string, error) {
	item := dec.ReadItem()
	vv, err := parseString(item)
	if err != nil {
		return "", fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type string", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemBool(jsonKey string) (bool, error) {
	item := dec.ReadItem()
	vv, err := parseBool(item)
	if err != nil {
		return false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type bool", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemInt32(jsonKey string) (int32, error) {
	item := dec.ReadItem()
	vv, err := parseInt32(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type int32", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemInt64(jsonKey string) (int64, error) {
	item := dec.ReadItem()
	vv, err := parseInt64(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type int64", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemUint32(jsonKey string) (uint32, error) {
	item := dec.ReadItem()
	vv, err := parseUint32(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type uint32", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemUint64(jsonKey string) (uint64, error) {
	item := dec.ReadItem()
	vv, err := parseUint64(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type uint64", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemFloat32(jsonKey string) (float32, error) {
	item := dec.ReadItem()
	vv, err := parseFloat32(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type float32", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemFloat64(jsonKey string) (float64, error) {
	item := dec.ReadItem()
	vv, err := parseFloat64(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type float64", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemEnumString(jsonKey string, em map[string]int32) (int32, error) {
	item := dec.ReadItem()
	vv, err := parseEnumString(item, em)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as string as array element into field %s of type Enum", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemEnumNumber(jsonKey string, em map[int32]string) (int32, error) {
	item := dec.ReadItem()
	vv, _err := parseEnumNumber(item, em)
	if _err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as number as array element into field %s of type Enum", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadArrayElemBytes(jsonKey string) ([]byte, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseBytes(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type []byte", string(item), jsonKey)
	}
	return vv, nil
}

// ReadArrayElemInterface returns (notNull, error)
func (dec *Decoder) ReadArrayElemInterface(jsonKey string, i interface{}) (bool, error) {
	return dec.ReadValueInterface(jsonKey, i)
}
