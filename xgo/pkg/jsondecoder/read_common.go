package jsondecoder

// BeforeScanJSON only used to check if the JSON is NULL before loop scan.
func (dec *Decoder) BeforeScanJSON() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if !isNULL {
		return
	}
	// Check the end of JSON input.
	if err = dec.scanWhile(scanEnd); err != nil {
		err = errorWrap(dec, err)
		return false, err
	}
	return
}

// BeforeScanNext only used to before read JSON key that in top level.
func (dec *Decoder) BeforeScanNext() (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadNext(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if !isEnd {
		return
	}
	// Check the end of JSON input.
	if err = dec.scanWhile(scanEnd); err != nil {
		err = errorWrap(dec, err)
		return false, err
	}
	return
}

// ReadJSONKey used to read JSON key that in top level.
// Notice: The `jsonKey` is unsafe after unmarshal end.
func (dec *Decoder) ReadJSONKey() (jsonKey string, err error) {
	// reset jsonKey
	dec.jsonKey = ""
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if jsonKey, err = bytesToStringUnsafe(key); err != nil {
		// TODO: optimize the error message.
		err = errorWrap(dec, err)
		return
	}
	dec.jsonKey = jsonKey
	return
}

func (dec *Decoder) BeforeReadObject() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return
}
func (dec *Decoder) BeforeReadMap() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return
}
func (dec *Decoder) BeforeReadList() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return
}

func (dec *Decoder) BeforeReadNext() (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadNext(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if !isEnd {
		return
	}
	if err = dec.scanNext(); err != nil {
		err = errorWrap(dec, err)
		return false, err
	}
	return
}

// ReadObjectKey is wrappers for readObjectKey.
// The `objKey` is unsafe after unmarshal end.
func (dec *Decoder) ReadObjectKey() (objKey string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if objKey, err = bytesToStringUnsafe(key); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return
}

// NextLiteralIsNULL only used to check if the next value that type of pointer value is NULL.
func (dec *Decoder) NextLiteralIsNULL() (isNULL bool, err error) {
	if isNULL, err = dec.nextLiteralIsNULL(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	if !isNULL {
		return
	}
	if _, err = dec.readLiteralValue(); err != nil { // Discard the value.
		err = errorWrap(dec, err)
		return
	}
	return
}

// DiscardValue used to discard the next value.
func (dec *Decoder) DiscardValue() (err error) {
	if _, err = dec.readLiteralValue(); err != nil {
		err = errorWrap(dec, err)
		return
	}
	return err
}
