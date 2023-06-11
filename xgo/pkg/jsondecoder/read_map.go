package jsondecoder

import (
	"fmt"
)

func (dec *Decoder) ReadMapKeyString(jsonKey string) (vv string, err error) {
	_ = jsonKey

	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if vv, err = bytesToString(key); err != nil {
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyBool(jsonKey string) (vv bool, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if key, err = dec.unquoteKey(key); err != nil {
		return
	}
	if vv, err = bytesToBool(key); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type string", string(key), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyInt32(jsonKey string) (vv int32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if key, err = dec.unquoteKey(key); err != nil {
		return
	}
	if vv, err = bytesToInt32(key); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type int32", string(key), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyInt64(jsonKey string) (vv int64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if key, err = dec.unquoteKey(key); err != nil {
		return
	}
	if vv, err = bytesToInt64(key); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type int64", string(key), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyUint32(jsonKey string) (vv uint32, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if key, err = dec.unquoteKey(key); err != nil {
		return
	}
	if vv, err = bytesToUint32(key); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type uint32", string(key), jsonKey)
		return
	}
	return
}
func (dec *Decoder) ReadMapKeyUint64(jsonKey string) (vv uint64, err error) {
	var key []byte
	if key, err = dec.readObjectKey(); err != nil {
		return
	}
	if key, err = dec.unquoteKey(key); err != nil {
		return
	}
	if vv, err = bytesToUint64(key); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as map key into field %s of type uint64", string(key), jsonKey)
		return
	}
	return
}
