package jsondecoder

func (dec *Decoder) unquoteString(b []byte) (t []byte, err error) {
	var ok bool
	if t, ok = unquoteBytes(b); !ok {
		err = &SyntaxError{
			reason: "invalid string of JSON input",
			Offset: dec.offset,
		}
		return nil, err
	}
	return
}

// BeforeReadJSON only used to check if the JSON is NULL.
func (dec *Decoder) BeforeReadJSON() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		return
	}
	return
}

// BeforeScanNext only used to before read JSON key that in top level.
func (dec *Decoder) BeforeScanNext() (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadNext(); err != nil {
		return
	}
	return
}

// ReadJSONKey used to read JSON key that in top level.
// The `jsonKey` is unsafe after unmarshal end.
func (dec *Decoder) ReadJSONKey() (jsonKey string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if jsonKey, err = bytesToStringUnsafe(key); err != nil {
		return
	}
	return
}

func (dec *Decoder) BeforeReadObject(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		return
	}
	return
}
func (dec *Decoder) BeforeReadMap(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		return
	}
	return
}
func (dec *Decoder) BeforeReadArray(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		return
	}
	return
}

func (dec *Decoder) BeforeReadNext(jsonKey string) (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadNext(); err != nil {
		return
	}
	if isEnd {
		if err = dec.scanNext(); err != nil {
			return false, err
		}
	}
	return
}

// ReadObjectKey is wrappers for readObjectKey.
// The `objKey` is unsafe after unmarshal end.
func (dec *Decoder) ReadObjectKey(jsonKey string) (objKey string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if objKey, err = bytesToStringUnsafe(key); err != nil {
		return
	}
	return
}

// NextLiteralIsNULL only used to check if the next value that type of pointer value is NULL.
func (dec *Decoder) NextLiteralIsNULL(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.nextLiteralIsNULL(); err != nil {
		return
	}
	if isNULL {
		if _, err = dec.readLiteralValue(); err != nil { // Discard the value.
			return
		}
	}
	return
}

// DiscardValue used to discard the next value.
func (dec *Decoder) DiscardValue(jsonKey string) (err error) {
	if _, err = dec.readLiteralValue(); err != nil {
		return
	}
	return err
}
