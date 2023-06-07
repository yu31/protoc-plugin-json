package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) afterReadArrayElem() (stop bool) {
	// Next token must be , or ].
	if dec.OpCode == ScanSkipSpace {
		dec.ScanWhile(ScanSkipSpace)
	}
	if dec.OpCode == ScanArrayEnd {
		return true
	}
	if dec.OpCode != ScanArrayElem {
		panic(PhasePanicMsg)
	}

	// ReadArrayElemBefore
	// check before read next elements.
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanArrayEnd {
		return true
	}
	return
}

func (dec *Decoder) ReadArrayElemString(jsonKey string) (string, bool, error) {
	item := dec.readItem()
	vv, err := parseString(item)
	if err != nil {
		return "", false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type string", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemBool(jsonKey string) (bool, bool, error) {
	item := dec.readItem()
	vv, err := parseBool(item)
	if err != nil {
		return false, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type bool", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemInt32(jsonKey string) (int32, bool, error) {
	item := dec.readItem()
	vv, err := parseInt32(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type int32", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemInt64(jsonKey string) (int64, bool, error) {
	item := dec.readItem()
	vv, err := parseInt64(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type int64", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemUint32(jsonKey string) (uint32, bool, error) {
	item := dec.readItem()
	vv, err := parseUint32(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type uint32", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemUint64(jsonKey string) (uint64, bool, error) {
	item := dec.readItem()
	vv, err := parseUint64(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type uint64", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemFloat32(jsonKey string) (float32, bool, error) {
	item := dec.readItem()
	vv, err := parseFloat32(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type float32", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemFloat64(jsonKey string) (float64, bool, error) {
	item := dec.readItem()
	vv, err := parseFloat64(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type float64", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemEnumString(jsonKey string, em map[string]int32) (int32, bool, error) {
	item := dec.readItem()
	vv, err := parseEnumString(item, em)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as string as array element into field %s of type Enum", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemEnumNumber(jsonKey string, em map[int32]string) (int32, bool, error) {
	item := dec.readItem()
	vv, _err := parseEnumNumber(item, em)
	if _err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as number as array element into field %s of type Enum", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}
func (dec *Decoder) ReadArrayElemBytes(jsonKey string) ([]byte, bool, error) {
	item := dec.readItem()
	if item[0] == 'n' { // 'n' means null
		return nil, false, nil
	}
	vv, err := parseBytes(item)
	if err != nil {
		return nil, false, fmt.Errorf("json: cannot unmarshal %s as array element into field %s of type []byte", string(item), jsonKey)
	}
	noMore := dec.afterReadArrayElem()
	return vv, noMore, nil
}

// ReadArrayElemInterface returns (isNULL, bool, error)
func (dec *Decoder) ReadArrayElemInterface(jsonKey string, initFN func() interface{}) (bool, error) {
	err := dec.ReadValueInterface(jsonKey, initFN)
	if err != nil {
		return false, err
	}
	noMore := dec.afterReadArrayElem()
	return noMore, nil
}
