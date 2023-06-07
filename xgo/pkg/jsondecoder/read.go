package jsondecoder

func (dec *Decoder) unquoteKey(b []byte) (t []byte, err error) {
	var ok bool
	if t, ok = unquoteBytes(b); !ok {
		err = &SyntaxError{
			reason: "invalid object key of JSON input",
			Offset: dec.offset,
		}
		return nil, err
	}
	return
}

func (dec *Decoder) BeforeScanJSON() (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadObject(); err != nil {
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
func (dec *Decoder) BeforeReadArray(jsonKey string) (isNULL bool, err error) {
	if isNULL, err = dec.beforeReadArray(); err != nil {
		return
	}
	return
}

func (dec *Decoder) BeforeReadJSONKey() (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadKey(); err != nil {
		return
	}
	return
}
func (dec *Decoder) BeforeReadObjectKey(jsonKey string) (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadKey(); err != nil {
		return
	}
	if !isEnd {
		return
	}
	if err = dec.scanNext(); err != nil {
		return false, err
	}
	return
}
func (dec *Decoder) BeforeReadArrayElem(jsonKey string) (isEnd bool, err error) {
	if isEnd, err = dec.beforeReadElem(); err != nil {
		return
	}
	if !isEnd {
		return
	}
	if err = dec.scanNext(); err != nil {
		return false, err
	}
	return
}

// ReadJSONKey get the key of JSON input. the `jsonKey` is unsafe after unmarshal end.
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

// ReadObjectKey get the key of JSON input. the `objKey` is unsafe after unmarshal end.
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

func (dec *Decoder) DiscardValue(jsonKey string) (err error) {
	if _, err = dec.readObjectValue(); err != nil {
		return
	}
	return err
}
