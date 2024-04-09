package jsondecoder

// readMapKeyStr read the key of map for type string.
func (dec *Decoder) readMapKeyStr() (vv string, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = convertToString(key); err != nil {
		return
	}
	return
}

// readMapKeyBool read the key of map for type bool.
func (dec *Decoder) readMapKeyBool(unquote bool) (vv bool, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = convertToBool(unquote, key); err != nil {
		return
	}
	return
}

// readMapKeyI32 read the key of map for type int32.
func (dec *Decoder) readMapKeyI32(unquote bool) (vv int32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = convertToInt32(unquote, key); err != nil {
		return
	}
	return
}

// readMapKeyI64 read the key of map for type int64.
func (dec *Decoder) readMapKeyI64(unquote bool) (vv int64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = convertToInt64(unquote, key); err != nil {
		return
	}
	return
}

// readMapKeyU32 read the key of map for type uint32.
func (dec *Decoder) readMapKeyU32(unquote bool) (vv uint32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = convertToUint32(unquote, key); err != nil {
		return
	}
	return
}

// readMapKeyU64 read the key of map for type uint64.
func (dec *Decoder) readMapKeyU64(unquote bool) (vv uint64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = convertToUint64(unquote, key); err != nil {
		return
	}
	return
}
