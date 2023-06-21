package jsonencoder

// AppendPointerString can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerString(v *string) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendString(*v)
	}
}

// AppendPointerBool can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerBool(v *bool, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendBool(*v, quote)
	}
}

// AppendPointerInt32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerInt32(v *int32, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendInt(int64(*v), quote)
	}
}

// AppendPointerInt64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerInt64(v *int64, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendInt(*v, quote)
	}
}

// AppendPointerUint32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerUint32(v *uint32, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendUint(uint64(*v), quote)
	}
}

// AppendPointerUint64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerUint64(v *uint64, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendUint(*v, quote)
	}
}

// AppendPointerFloat32 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerFloat32(v *float32, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendFloat(float64(*v), 32, quote)
	}
}

// AppendPointerFloat64 can be used for value of literal, slice, array and map.
func (enc *Encoder) AppendPointerFloat64(v *float64, quote bool) {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
	} else {
		enc.appendFloat(*v, 64, quote)
	}
}
