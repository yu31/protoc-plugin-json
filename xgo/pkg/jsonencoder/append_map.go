package jsonencoder

func (enc *Encoder) AppendMapKeyString(v string) {
	enc.appendCommaSeparator()
	enc.appendString(v)
	enc.writeByte(':')
}

func (enc *Encoder) AppendMapKeyBool(v bool, quote bool) {
	enc.appendCommaSeparator()
	enc.appendBool(v, quote)
	enc.writeByte(':')
}

func (enc *Encoder) AppendMapKeyInt32(v int32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(int64(v), quote)
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyInt64(v int64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(v, quote)
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyUInt32(v uint32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(uint64(v), quote)
	enc.writeByte(':')
}
func (enc *Encoder) AppendMapKeyUInt64(v uint64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(v, quote)
	enc.writeByte(':')
}
