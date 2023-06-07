package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) afterReadMapValue() (stop bool) {
	// After read value, Next token must be , or }.
	if dec.OpCode == ScanSkipSpace {
		dec.ScanWhile(ScanSkipSpace)
	}
	if dec.OpCode == ScanObjectEnd {
		return true
	}
	if dec.OpCode != ScanObjectValue {
		panic(PhasePanicMsg)
	}

	// ReadObjectKeyBefore
	// Check before read next map key.
	// Before read key, Read opening " of string key or closing }.
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanObjectEnd {
		// object closing charset(`}`) - can only happen on first iteration.
		return true
	}
	if dec.OpCode != ScanLiteralBegin {
		panic(PhasePanicMsg)
	}
	return
}

func (dec *Decoder) ReadMapValueString(jsonKey string) (string, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseString(item)
	if err != nil {
		return "", false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type string", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueBool(jsonKey string) (bool, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseBool(item)
	if err != nil {
		return false, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type bool", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueInt32(jsonKey string) (int32, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseInt32(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type int32", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueInt64(jsonKey string) (int64, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseInt64(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type int64", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueUint32(jsonKey string) (uint32, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseUint32(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type uint32", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueUint64(jsonKey string) (uint64, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseUint64(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type uint64", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}

func (dec *Decoder) ReadMapValueFloat32(jsonKey string) (float32, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseFloat32(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type float32", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueFloat64(jsonKey string) (float64, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseFloat64(item)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type float64", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueEnumString(jsonKey string, em map[string]int32) (int32, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, err := parseEnumString(item, em)
	if err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as string as map value into field %s of type Enum", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueEnumNumber(jsonKey string, em map[int32]string) (int32, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	vv, _err := parseEnumNumber(item, em)
	if _err != nil {
		return 0, false, fmt.Errorf("json: cannot unmarshal %s as number as map value into field %s of type Enum", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}
func (dec *Decoder) ReadMapValueBytes(jsonKey string) ([]byte, bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	item := dec.readItem()
	if item[0] == 'n' { // 'n' means null
		return nil, false, nil
	}
	vv, err := parseBytes(item)
	if err != nil {
		return nil, false, fmt.Errorf("json: cannot unmarshal %s as map value into field %s of type []byte", string(item), jsonKey)
	}
	noMore := dec.afterReadMapValue()
	return vv, noMore, nil
}

// ReadMapValueInterface returns (isNULL,bool, error)
func (dec *Decoder) ReadMapValueInterface(jsonKey string, initFN func() interface{}) (bool, error) {
	//dec.ReadObjectValueBefore() // TODO: review it.
	err := dec.ReadValueInterface(jsonKey, initFN)
	if err != nil {
		return false, err
	}
	noMore := dec.afterReadMapValue()
	return noMore, nil
}
