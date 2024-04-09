package jsonencoder

// AppendObjectBegin append a beginning of curly braces into json contents.
// It can be used for begin of JSON, map and struct.
func (enc *Encoder) AppendObjectBegin() {
	enc.writeByte('{')
}

// AppendObjectEnd append an ending of curly braces into json contents.
// It can be used for end of JSON, map and struct.
func (enc *Encoder) AppendObjectEnd() {
	enc.writeByte('}')
}

// AppendListBegin append a beginning of square brackets into json contents.
// It can be used for begin of array and slice.
func (enc *Encoder) AppendListBegin() {
	enc.writeByte('[')
}

// AppendListEnd append an ending of square brackets into json contents.
// It can be used for end of array and slice.
func (enc *Encoder) AppendListEnd() {
	enc.writeByte(']')
}

// AppendObjectKey append an object key and colon delimiter into JSON contents.
func (enc *Encoder) AppendObjectKey(k string) {
	enc.appendCommaSeparator()
	enc.appendString(k)
	enc.writeByte(':')
}

// AppendValNULL append a value `null` to JSON contents.
func (enc *Encoder) AppendValNULL() {
	enc.appendCommaSeparator()
	enc.writeString("null")
}
