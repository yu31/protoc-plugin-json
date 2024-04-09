package jsonencoder

// appendMapKeyStr to append the key of map for type string.
func (enc *Encoder) appendMapKeyStr(v string) {
	enc.appendCommaSeparator()
	enc.appendString(v)
	enc.writeByte(':')
}

// appendMapKeyBool to append the key of map for type bool.
func (enc *Encoder) appendMapKeyBool(v bool, quote bool) {
	enc.appendCommaSeparator()
	enc.appendBool(v, quote)
	enc.writeByte(':')
}

// appendMapKeyI32 to append the key of map for type int32.
func (enc *Encoder) appendMapKeyI32(v int32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(int64(v), quote)
	enc.writeByte(':')
}

// appendMapKeyI64 to append the key of map for type int64.
func (enc *Encoder) appendMapKeyI64(v int64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(v, quote)
	enc.writeByte(':')
}

// appendMapKeyU32 to append the key of map for type uint32.
func (enc *Encoder) appendMapKeyU32(v uint32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(uint64(v), quote)
	enc.writeByte(':')
}

// appendMapKeyU64 to append the key of map for type uint64.
func (enc *Encoder) appendMapKeyU64(v uint64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(v, quote)
	enc.writeByte(':')
}
