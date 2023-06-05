package jsondecoder

import (
	"unsafe"
)

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

// ReadObjectKey Read key of object or map.
func (dec *Decoder) ReadObjectKey() string {
	item := dec.ReadItem()
	key, ok := UnquoteBytes(item)
	if !ok {
		panic(PhasePanicMsg)
	}
	return *(*string)(unsafe.Pointer(&key))
}

// ReadObjectKeyBefore used to check if the object is read end.
func (dec *Decoder) ReadObjectKeyBefore() (stop bool) {
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

func (dec *Decoder) ReadArrayElemBefore() (stop bool) {
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanArrayEnd {
		return true
	}
	return
}

func (dec *Decoder) ReadArrayElemAfter() (stop bool) {
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

	return dec.ReadArrayElemBefore()

	//// check before read next elements.
	//dec.ScanWhile(ScanSkipSpace)
	//if dec.OpCode == ScanArrayEnd {
	//	return true
	//}
	//return
}

func (dec *Decoder) ReadMapValueAfter() (stop bool) {
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

	return dec.ReadObjectKeyBefore()

	//// Check before read next map key.
	//// Before read key, Read opening " of string key or closing }.
	//dec.ScanWhile(ScanSkipSpace)
	//if dec.OpCode == ScanObjectEnd {
	//	// object closing charset(`}`) - can only happen on first iteration.
	//	return true
	//}
	//if dec.OpCode != ScanLiteralBegin {
	//	panic(PhasePanicMsg)
	//}
	//return
}
