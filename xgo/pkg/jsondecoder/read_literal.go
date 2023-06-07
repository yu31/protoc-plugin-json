package jsondecoder

import (
	"encoding/json"
	"fmt"
)

func (dec *Decoder) ReadValueString(jsonKey string) (string, error) {
	item := dec.readItem()
	vv, err := parseString(item)
	if err != nil {
		return "", fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueBool(jsonKey string) (bool, error) {
	item := dec.readItem()
	vv, err := parseBool(item)
	if err != nil {
		return false, fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInt32(jsonKey string) (int32, error) {
	item := dec.readItem()
	vv, err := parseInt32(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueInt64(jsonKey string) (int64, error) {
	item := dec.readItem()
	vv, err := parseInt64(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueUint32(jsonKey string) (uint32, error) {
	item := dec.readItem()
	vv, err := parseUint32(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueUint64(jsonKey string) (uint64, error) {
	item := dec.readItem()
	vv, err := parseUint64(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(item), jsonKey)
	}
	return vv, nil
}

func (dec *Decoder) ReadValueFloat32(jsonKey string) (float32, error) {
	item := dec.readItem()
	vv, err := parseFloat32(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueFloat64(jsonKey string) (float64, error) {
	item := dec.readItem()
	vv, err := parseFloat64(item)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueEnumString(jsonKey string, em map[string]int32) (int32, error) {
	item := dec.readItem()
	vv, err := parseEnumString(item, em)
	if err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueEnumNumber(jsonKey string, em map[int32]string) (int32, error) {
	item := dec.readItem()
	vv, _err := parseEnumNumber(item, em)
	if _err != nil {
		return 0, fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum", string(item), jsonKey)
	}
	return vv, nil
}
func (dec *Decoder) ReadValueBytes(jsonKey string) ([]byte, error) {
	item := dec.readItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseBytes(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type []byte", string(item), jsonKey)
	}
	return vv, nil
}

// ReadValueInterface returns (isNULL, error)
func (dec *Decoder) ReadValueInterface(jsonKey string, initFN func() interface{}) error {
	_ = jsonKey
	item := dec.readItem()
	if item[0] == 'n' { // 'n' means null
		return nil
	}
	i := initFN()
	if um, ok := i.(json.Unmarshaler); ok {
		return um.UnmarshalJSON(item)
	}
	return json.Unmarshal(item, i)
}
