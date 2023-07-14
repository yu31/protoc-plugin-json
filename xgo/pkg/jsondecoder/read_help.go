package jsondecoder

// BeforeScanJSON only used to check if the JSON is NULL before loop scan.
func (dec *Decoder) BeforeScanJSON() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap("", dec.offset, err)
		return
	}
	if !isNULL {
		return
	}
	// Check the end of JSON input.
	if err = dec.scanWhile(scanEnd); err != nil {
		err = errorWrap("", dec.offset, err)
		return false, err
	}
	return
}

// BeforeScanNext only used to before read JSON key that in top level.
func (dec *Decoder) BeforeScanNext() (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadNext(); err != nil {
		err = errorWrap("", dec.offset, err)
		return
	}
	if !isEnd {
		return
	}
	// Check the end of JSON input.
	if err = dec.scanWhile(scanEnd); err != nil {
		err = errorWrap("", dec.offset, err)
		return false, err
	}
	return
}

// ReadJSONKey used to read JSON key that in top level.
// The `jsonKey` is unsafe after unmarshal end.
func (dec *Decoder) ReadJSONKey() (jsonKey string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap("", dec.offset, err)
		return
	}
	if jsonKey, err = bytesToStringUnsafe(key); err != nil {
		// TODO: optimize the error message.
		err = errorWrap("", dec.offset, err)
		return
	}
	return
}

func (dec *Decoder) BeforeReadObject(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) BeforeReadMap(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}
func (dec *Decoder) BeforeReadArray(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}

func (dec *Decoder) BeforeReadNext(jsonKey string) (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadNext(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if !isEnd {
		return
	}
	if err = dec.scanNext(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return false, err
	}
	return
}

// ReadObjectKey is wrappers for readObjectKey.
// The `objKey` is unsafe after unmarshal end.
func (dec *Decoder) ReadObjectKey(jsonKey string) (objKey string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if objKey, err = bytesToStringUnsafe(key); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}

// NextLiteralIsNULL only used to check if the next value that type of pointer value is NULL.
func (dec *Decoder) NextLiteralIsNULL(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.nextLiteralIsNULL(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	if !isNULL {
		return
	}
	if _, err = dec.readLiteralValue(); err != nil { // Discard the value.
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return
}

// DiscardValue used to discard the next value.
func (dec *Decoder) DiscardValue(jsonKey string) (err error) {
	if _, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(jsonKey, dec.offset, err)
		return
	}
	return err
}
