package jsondecoder

import (
	"errors"
	"unsafe"
)

// CheckJSONIsNULL check whether is null for the next item.
// Used to object, map, array.

// ReadItem read current item value.
func (dec *Decoder) ReadItem() []byte {
	// start index.
	start := dec.off - 1
	switch dec.OpCode {
	case ScanLiteralBegin:
		dec.RescanLiteral()
	case ScanArrayBegin, ScanObjectBegin:
		dec.Skip()
		dec.ScanNext()
	default:
		panic(PhasePanicMsg)
	}
	// end index.
	end := dec.off - 1
	return dec.data[start:end]
}

// CheckJSONBegin check whether the beginning of json content are valid.
// returns (isNull, error)
func (dec *Decoder) CheckJSONBegin() (bool, error) {
	// Skip the space charsets.
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanLiteralBegin {
		// Expect the next value is null.
		if item := dec.ReadItem(); item[0] != 'n' {
			return true, errors.New("json: invalid format of json content that you provides")
		}
		return true, nil
	}
	// The json content are not null. expected the first non-space charset is object begin(`{`)
	if dec.OpCode != ScanObjectBegin {
		// FIXME: return a error instead of panic ?
		panic(PhasePanicMsg)
	}
	return false, nil
}

// ReadObjectKey Read key of object or map.
func (dec *Decoder) ReadObjectKey() string {
	item := dec.ReadItem()
	key, ok := UnquoteBytes(item)
	if !ok {
		panic(PhasePanicMsg)
	}
	return *(*string)(unsafe.Pointer(&key))
}

func (dec *Decoder) ReadObjectKeyBefore() (stop bool) {
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanObjectEnd {
		// object closing charters(`}`) - can only happen on first iteration.
		return true
	}
	if dec.OpCode != ScanLiteralBegin {
		panic(PhasePanicMsg)
	}
	return
}

func (dec *Decoder) ReadObjectValueBefore() {
	if dec.OpCode == ScanSkipSpace {
		dec.ScanWhile(ScanSkipSpace)
	}
	if dec.OpCode != ScanObjectKey {
		panic(PhasePanicMsg)
	}
	dec.ScanWhile(ScanSkipSpace)
}

func (dec *Decoder) ReadObjectValueAfter() (stop bool) {
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
	return
}

func (dec *Decoder) ReadArrayValueBefore() (stop bool) {
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanArrayEnd {
		return true
	}
	return
}

func (dec *Decoder) AfterReadArrayValueAfter() (stop bool) {
	// Next token must be , or ].
	if dec.OpCode == ScanSkipSpace {
		dec.ScanWhile(ScanSkipSpace)
	}
	if dec.OpCode == ScanArrayEnd {
		return true
	}
	if dec.OpCode != ScanArrayValue {
		panic(PhasePanicMsg)
	}
	return
}
