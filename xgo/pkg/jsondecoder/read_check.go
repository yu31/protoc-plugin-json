package jsondecoder

import (
	"errors"
	"fmt"
)

// CheckJSONBegin check if the content is null or a valid JSON.
func (dec *Decoder) CheckJSONBegin() (isNULL bool, err error) {
	// Skip the space charsets.
	dec.ScanWhile(ScanSkipSpace)
	// Check if the json is null.
	if dec.OpCode == ScanLiteralBegin {
		// Expect the next value is null.
		if item := dec.ReadItem(); item[0] != 'n' {
			return true, errors.New("json: invalid format of json content that you provides")
		}
		return true, nil
	}
	// The json are not null. expected the first non-space charset is object begin(`{`)
	if dec.OpCode != ScanObjectBegin {
		return false, errors.New("json: invalid charsets beginning in JSON content")
	}
	return false, nil
}

// CheckBeforeReadArray check if the content is a valid array, and check if it is null or empty.
func (dec *Decoder) CheckBeforeReadArray(jsonKey string) (isNULL, isEmpty bool, err error) {
	// Check if the array is null.
	if dec.OpCode == ScanLiteralBegin {
		if item := dec.ReadItem(); item[0] != 'n' {
			err = fmt.Errorf("json: cannot unmarshal %s as array into field %s", string(item), jsonKey)
			return false, false, err
		}
		return true, false, nil
	}
	// The array is not null. expected the first non-space charset is object begin(`[`)
	if dec.OpCode != ScanArrayBegin {
		item := dec.ReadItem()
		err = fmt.Errorf("json: cannot unmarshal %s as array into field %s", string(item), jsonKey)
		return false, false, err
	}
	// Check if the array is empty.
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanArrayEnd {
		dec.ScanWhile(ScanSkipSpace)
		return false, true, nil
	}
	return false, false, nil
}

func (dec *Decoder) CheckBeforeReadMap(jsonKey string) (isNULL, isEmpty bool, err error) {
	// Check if the map is null.
	if dec.OpCode == ScanLiteralBegin {
		if item := dec.ReadItem(); item[0] != 'n' {
			err = fmt.Errorf("json: cannot unmarshal %s as map into field %s", string(item), jsonKey)
			return false, false, err
		}
		return true, false, nil
	}
	// The map is not null. expected the first non-space charset is object begin(`{`)
	if dec.OpCode != ScanObjectBegin {
		item := dec.ReadItem()
		err = fmt.Errorf("json: cannot unmarshal %s as map into field %s of type", string(item), jsonKey)
		return false, false, err
	}

	// read next.
	dec.ScanWhile(ScanSkipSpace)

	// Check is the map is empty.
	if dec.OpCode == ScanObjectEnd {
		//dec.ScanWhile(ScanSkipSpace)
		// object closing charset(`}`) - can only happen on first iteration.
		return false, true, nil
	}
	if dec.OpCode != ScanLiteralBegin {
		panic(PhasePanicMsg)
	}
	return false, false, nil
}

// CheckObjectBegin check if the content is null or a valid object.
func (dec *Decoder) CheckObjectBegin(jsonKey string) (isNULL bool, err error) {
	// Check if the object is null.
	if dec.OpCode == ScanLiteralBegin {
		if item := dec.ReadItem(); item[0] != 'n' {
			err = fmt.Errorf("json: cannot unmarshal %s as object into field %s", string(item), jsonKey)
			return false, err
		}
		return true, nil
	}
	// The object is not null. expected the first non-space charset is object begin(`[`)
	if dec.OpCode != ScanObjectBegin {
		item := dec.ReadItem()
		err = fmt.Errorf("json: cannot unmarshal %s as object into field %s of type", string(item), jsonKey)
		return false, err
	}
	return false, nil
}
