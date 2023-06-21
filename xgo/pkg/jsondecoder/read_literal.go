package jsondecoder

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func (dec *Decoder) ReadLiteralBytes(jsonKey string) (vv []byte, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}

	b, ok := unquoteStdBytes(value)
	if !ok {
		// TODO: optimized the error message.
		err = fmt.Errorf("json: cannot unmarshal %s into field %s of type []byte", string(value), jsonKey)
		return
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(dst, b)
	if err != nil {
		// TODO: optimized the error message.
		return nil, err
	}
	vv = dst[:n]
	return vv, nil
}

func (dec *Decoder) ReadLiteralInterface(jsonKey string, vv interface{}) (err error) {
	assertInterface(vv)
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil
	}
	if um, ok := vv.(json.Unmarshaler); ok {
		err = um.UnmarshalJSON(value)
	} else {
		err = json.Unmarshal(value, vv)
	}
	if err != nil {
		return
	}
	return
}

func (dec *Decoder) ReadLiteralString(jsonKey string) (vv string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToString(jsonKey, value); err != nil {
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadLiteralBool(jsonKey string, unquote bool) (vv bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToBool(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralInt32(jsonKey string, unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToInt32(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralInt64(jsonKey string, unquote bool) (vv int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToInt64(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralUint32(jsonKey string, unquote bool) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToUint32(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralUint64(jsonKey string, unquote bool) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToUint64(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}

func (dec *Decoder) ReadLiteralFloat32(jsonKey string, unquote bool) (vv float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToFloat32(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralFloat64(jsonKey string, unquote bool) (vv float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToFloat64(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralEnumString(jsonKey string, em map[string]int32) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}

	if vv, err = parseEnumString(value, em); err != nil {
		err = fmt.Errorf("json: cannot unmarshal %s as string into field %s of type Enum", string(value), jsonKey)
		return
	}
	return vv, nil
}
func (dec *Decoder) ReadLiteralEnumNumber(jsonKey string, unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = dec.convertToEnum(jsonKey, value, unquote); err != nil {
		return
	}
	return vv, nil
}
