package jsondecoder

import (
	"encoding/json"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func (dec *Decoder) readValI32(unquote bool) (vv int32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToInt32(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValI64(unquote bool) (vv int64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToInt64(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValU32(unquote bool) (vv uint32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToUint32(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValU64(unquote bool) (vv uint64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToUint64(unquote, value); err != nil {
		return
	}
	return vv, nil
}

func (dec *Decoder) readValF32(unquote bool) (vv float32, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToFloat32(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValF64(unquote bool) (vv float64, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToFloat64(unquote, value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValBool(unquote bool) (vv bool, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToBool(unquote, value); err != nil {
		return
	}
	return vv, nil
}

func (dec *Decoder) readValStr() (vv string, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if vv, err = convertToString(value); err != nil {
		return
	}
	return vv, nil
}
func (dec *Decoder) readValBytes() (vv []byte, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	if value[0] == 'n' { // 'n' means null
		return nil, nil
	}
	if vv, err = convertToBytes(value); err != nil {
		return nil, err
	}
	return vv, nil
}
func (dec *Decoder) readValInterface(vv interface{}) (err error) {
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
		//err = errorWrap(dec, err)
		return
	}
	return
}

func readValEnumNum[T protoreflect.Enum](dec *Decoder, val T, unquote bool) (vv T, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var v1 int32
	if v1, err = convertToEnum(unquote, value); err != nil {
		return
	}
	vv = val.Type().New(protoreflect.EnumNumber(v1)).(T)
	return vv, nil
}
func readValEnumStr[T protoreflect.Enum](dec *Decoder, val T) (vv T, err error) {
	var value []byte
	if value, err = dec.readLiteralValue(); err != nil {
		return
	}
	var s1 string
	if s1, err = bytesToStringUnsafe(value); err != nil {
		return
	}
	//enumNumber := protoimpl.X.EnumDescriptorOf(val).Values().ByName(protoreflect.Name(s1)).Number()
	enumNumber := val.Descriptor().Values().ByName(protoreflect.Name(s1)).Number()
	vv = val.Type().New(enumNumber).(T)
	return
}
