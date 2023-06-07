package jsondecoder

import (
	"unsafe"
)

// readItem read current item value.
func (dec *Decoder) readItem() []byte {
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

func (dec *Decoder) Discard() error {
	_ = dec.readItem()
	return nil
}

func (dec *Decoder) ReadJSONKey() (key string, stop bool) {
	// Expected opening charsets(`"`) of string key or closing charsets(`}`)
	dec.ScanWhile(ScanSkipSpace)
	if dec.OpCode == ScanObjectEnd {
		// object closing charset(`}`) - can only happen on first iteration.
		return "", true
	}
	if dec.OpCode != ScanLiteralBegin {
		// FIXME: return a error?
		panic(PhasePanicMsg)
	}
	// Read the json key.
	key = dec.readObjectKey()
	return key, false
}

// readObjectKey Read key of object or map.
func (dec *Decoder) readObjectKey() string {
	item := dec.readItem()
	bb, ok := unquoteBytes(item)
	if !ok {
		// FIXME: return error instead of panic?
		panic(PhasePanicMsg)
	}

	// prepare to read object value.
	// TODO: review it.
	//dec.ReadObjectValueBefore()
	if dec.OpCode == ScanSkipSpace {
		dec.ScanWhile(ScanSkipSpace)
	}
	if dec.OpCode != ScanObjectKey {
		panic(PhasePanicMsg)
	}
	dec.ScanWhile(ScanSkipSpace)

	return *(*string)(unsafe.Pointer(&bb))
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
