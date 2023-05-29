package jsondecoder

import "fmt"

func (dec *Decoder) ReadPointerString(jsonKey string) (*string, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseString(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type string", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerBool(jsonKey string) (*bool, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseBool(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type bool", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerInt32(jsonKey string) (*int32, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseInt32(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type int32", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerInt64(jsonKey string) (*int64, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseInt64(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type int64", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerUint32(jsonKey string) (*uint32, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseUint32(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type uint32", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerUint64(jsonKey string) (*uint64, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseUint64(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type uint64", string(item), jsonKey)
	}
	return &vv, nil
}

func (dec *Decoder) ReadPointerFloat32(jsonKey string) (*float32, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseFloat32(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type float32", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerFloat64(jsonKey string) (*float64, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseFloat64(item)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s into field %s of type float64", string(item), jsonKey)
	}
	return &vv, nil
}

func (dec *Decoder) ReadPointerEnumString(jsonKey string, em map[string]int32) (*int32, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, err := parseEnumString(item, em)
	if err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(item), jsonKey)
	}
	return &vv, nil
}
func (dec *Decoder) ReadPointerEnumNumber(jsonKey string, em map[int32]string) (*int32, error) {
	item := dec.ReadItem()
	if item[0] == 'n' { // 'n' means null
		return nil, nil
	}
	vv, _err := parseEnumNumber(item, em)
	if _err != nil {
		return nil, fmt.Errorf("json: cannot unmarshal %s as number into field %s of type Enum1", string(item), jsonKey)
	}
	return &vv, nil
}
