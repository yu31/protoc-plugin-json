// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/example/example1.proto

package pbexample

import (
	errors "errors"
	fmt "fmt"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message Empty1 in file tests/proto/example/example1.proto
func (x *Empty1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Empty1 in file tests/proto/example/example1.proto
func (x *Empty1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Empty1) is nil")
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Example1 in file tests/proto/example/example1.proto
func (x *Example1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(738)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	switch ov := x.OnetType1.(type) {
	case *Example1_FString2:
		encoder.AppendJSONKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendJSONKey("f_string2")
		encoder.AppendValueString(ov.FString2)
		encoder.AppendObjectEnd()
	case *Example1_FString3:
		encoder.AppendJSONKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendJSONKey("f_string3")
		encoder.AppendValueString(ov.FString3)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendJSONKey("OnetType1")
		encoder.AppendValueNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OnetType2.(type) {
	case *Example1_FString4:
		encoder.AppendJSONKey("f_string4")
		encoder.AppendValueString(ov.FString4)
	case *Example1_FString5:
		encoder.AppendJSONKey("f_string5")
		encoder.AppendValueString(ov.FString5)
	default:
		_ = ov // to avoids unused panics
	}
	encoder.AppendJSONKey("f_int32")
	encoder.AppendValueInt32(x.FInt32)
	encoder.AppendJSONKey("f_int64")
	encoder.AppendValueInt64(x.FInt64)
	encoder.AppendJSONKey("f_uint32")
	encoder.AppendValueUint32(x.FUint32)
	encoder.AppendJSONKey("f_uint64")
	encoder.AppendValueUint64(x.FUint64)
	encoder.AppendJSONKey("f_sint32")
	encoder.AppendValueInt32(x.FSint32)
	encoder.AppendJSONKey("f_sint64")
	encoder.AppendValueInt64(x.FSint64)
	encoder.AppendJSONKey("f_sfixed32")
	encoder.AppendValueInt32(x.FSfixed32)
	encoder.AppendJSONKey("f_sfixed64")
	encoder.AppendValueInt64(x.FSfixed64)
	encoder.AppendJSONKey("f_fixed32")
	encoder.AppendValueUint32(x.FFixed32)
	encoder.AppendJSONKey("f_fixed64")
	encoder.AppendValueUint64(x.FFixed64)
	encoder.AppendJSONKey("f_float")
	encoder.AppendValueFloat32(x.FFloat)
	encoder.AppendJSONKey("f_double")
	encoder.AppendValueFloat64(x.FDouble)
	encoder.AppendJSONKey("f_bool1")
	encoder.AppendValueBool(x.FBool1)
	encoder.AppendJSONKey("f_bytes1")
	encoder.AppendValueBytes(x.FBytes1)
	if x.FEnum1 != 0 {
		encoder.AppendJSONKey("f_enum1")
		encoder.AppendValueString(x.FEnum1.String())
	}
	encoder.AppendJSONKey("f_enum2")
	encoder.AppendValueInt32(int32(x.FEnum2.Number()))
	if x.FEnum5 != nil {
		encoder.AppendJSONKey("f_enum5")
		if x.FEnum5 != nil {
			encoder.AppendValueString(x.FEnum5.String())
		} else {
			encoder.AppendValueNULL()
		}
	}
	encoder.AppendJSONKey("f_enum6")
	if x.FEnum6 != nil {
		encoder.AppendValueInt32(int32(x.FEnum6.Number()))
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_duration1")
	err = encoder.AppendValueInterface(x.FDuration1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_timestamp1")
	err = encoder.AppendValueInterface(x.FTimestamp1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message11")
	err = encoder.AppendValueInterface(x.FMessage11)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message12")
	err = encoder.AppendValueInterface(x.FMessage12)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_message13")
	err = encoder.AppendValueInterface(x.FMessage13)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_any1")
	err = encoder.AppendValueInterface(x.FAny1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("r_string1")
	if x.RString1 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RString1 {
			encoder.AppendValueString(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("r_int32")
	if x.RInt32 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt32 {
			encoder.AppendValueInt32(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("r_message")
	if x.RMessage != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RMessage {
			err = encoder.AppendValueInterface(ri)
			if err != nil {
				return nil, err
			}
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("r_enum")
	if x.REnum != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.REnum {
			encoder.AppendValueInt32(int32(ri.Number()))
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("m_string1")
	if x.MString1 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MString1 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendValueString(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("m_string2")
	if x.MString2 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MString2 {
			encoder.AppendMapKeyInt32(mk)
			encoder.AppendValueString(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("m_message1")
	if x.MMessage1 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MMessage1 {
			encoder.AppendMapKeyString(mk)
			err = encoder.AppendValueInterface(mv)
			if err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("m_enum")
	if x.MEnum != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MEnum {
			encoder.AppendMapKeyString(mk)
			encoder.AppendValueInt32(int32(mv.Number()))
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Example1 in file tests/proto/example/example1.proto
func (x *Example1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Example1) is nil")
	}
	var oneOfIsFill_OnetType1 bool
	var oneOfIsFill_OnetType2 bool

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "OnetType1":
			if isNULL, err = decoder.CheckObjectBegin(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.OnetType1 = nil
			default:
			LOOP_ONEOF_OnetType1:
				for {
					oneofKey, stop := decoder.ReadJSONKey()
					if stop {
						break LOOP_ONEOF_OnetType1
					}
					switch {
					case oneofKey == "f_string2":
						var ot *Example1_FString2
						if _ot, _ok := x.OnetType1.(*Example1_FString2); _ok {
							ot = _ot
						} else {
							ot = new(Example1_FString2)
						}
						vv, _err := decoder.ReadValueString(jsonKey)
						if _err != nil {
							return _err
						}
						if oneOfIsFill_OnetType1 {
							return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
						}
						oneOfIsFill_OnetType1 = true
						ot.FString2 = vv
						x.OnetType1 = ot
					case oneofKey == "f_string3":
						var ot *Example1_FString3
						if _ot, _ok := x.OnetType1.(*Example1_FString3); _ok {
							ot = _ot
						} else {
							ot = new(Example1_FString3)
						}
						vv, _err := decoder.ReadValueString(jsonKey)
						if _err != nil {
							return _err
						}
						if oneOfIsFill_OnetType1 {
							return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
						}
						oneOfIsFill_OnetType1 = true
						ot.FString3 = vv
						x.OnetType1 = ot
					default:
						if err = decoder.Discard(); err != nil { // discard unknown field
							return err
						}
					}
					if decoder.ReadObjectValueAfter() { // After read object value
						break LOOP_ONEOF_OnetType1
					}
				}
				decoder.ScanNext()
			}
		case jsonKey == "f_string4":
			var ot *Example1_FString4
			if _ot, _ok := x.OnetType2.(*Example1_FString4); _ok {
				ot = _ot
			} else {
				ot = new(Example1_FString4)
			}
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			if oneOfIsFill_OnetType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OnetType2 = true
			ot.FString4 = vv
			x.OnetType2 = ot
		case jsonKey == "f_string5":
			var ot *Example1_FString5
			if _ot, _ok := x.OnetType2.(*Example1_FString5); _ok {
				ot = _ot
			} else {
				ot = new(Example1_FString5)
			}
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			if oneOfIsFill_OnetType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OnetType2 = true
			ot.FString5 = vv
			x.OnetType2 = ot
		case jsonKey == "f_int32":
			vv, _err := decoder.ReadValueInt32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FInt32 = vv
		case jsonKey == "f_int64":
			vv, _err := decoder.ReadValueInt64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FInt64 = vv
		case jsonKey == "f_uint32":
			vv, _err := decoder.ReadValueUint32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FUint32 = vv
		case jsonKey == "f_uint64":
			vv, _err := decoder.ReadValueUint64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FUint64 = vv
		case jsonKey == "f_sint32":
			vv, _err := decoder.ReadValueInt32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSint32 = vv
		case jsonKey == "f_sint64":
			vv, _err := decoder.ReadValueInt64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSint64 = vv
		case jsonKey == "f_sfixed32":
			vv, _err := decoder.ReadValueInt32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSfixed32 = vv
		case jsonKey == "f_sfixed64":
			vv, _err := decoder.ReadValueInt64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FSfixed64 = vv
		case jsonKey == "f_fixed32":
			vv, _err := decoder.ReadValueUint32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FFixed32 = vv
		case jsonKey == "f_fixed64":
			vv, _err := decoder.ReadValueUint64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FFixed64 = vv
		case jsonKey == "f_float":
			vv, _err := decoder.ReadValueFloat32(jsonKey)
			if _err != nil {
				return _err
			}
			x.FFloat = vv
		case jsonKey == "f_double":
			vv, _err := decoder.ReadValueFloat64(jsonKey)
			if _err != nil {
				return _err
			}
			x.FDouble = vv
		case jsonKey == "f_bool1":
			vv, _err := decoder.ReadValueBool(jsonKey)
			if _err != nil {
				return _err
			}
			x.FBool1 = vv
		case jsonKey == "f_bytes1":
			vv, _err := decoder.ReadValueBytes(jsonKey)
			if _err != nil {
				return _err
			}
			x.FBytes1 = vv
		case jsonKey == "f_enum1":
			v1, _err := decoder.ReadValueEnumString(jsonKey, Enum1_value)
			vv := Enum1(v1)
			if _err != nil {
				return _err
			}
			x.FEnum1 = vv
		case jsonKey == "f_enum2":
			v1, _err := decoder.ReadValueEnumNumber(jsonKey, Enum1_name)
			vv := Enum1(v1)
			if _err != nil {
				return _err
			}
			x.FEnum2 = vv
		case jsonKey == "f_enum5":
			var vv *Enum1
			v1, _err := decoder.ReadPointerEnumString(jsonKey, Enum1_value)
			if v1 != nil {
				_vv := Enum1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum5 = vv
		case jsonKey == "f_enum6":
			var vv *Enum1
			v1, _err := decoder.ReadPointerEnumNumber(jsonKey, Enum1_name)
			if v1 != nil {
				_vv := Enum1(*v1)
				vv = &_vv
			}
			if _err != nil {
				return _err
			}
			x.FEnum6 = vv
		case jsonKey == "f_duration1":
			var vv *durationpb.Duration
			initFN := func() interface{} {
				if x.FDuration1 != nil {
					vv = x.FDuration1
				} else {
					vv = new(durationpb.Duration)
				}
				return vv
			}
			_err := decoder.ReadValueInterface(jsonKey, initFN)
			if _err != nil {
				return _err
			}
			x.FDuration1 = vv
		case jsonKey == "f_timestamp1":
			var vv *timestamppb.Timestamp
			initFN := func() interface{} {
				if x.FTimestamp1 != nil {
					vv = x.FTimestamp1
				} else {
					vv = new(timestamppb.Timestamp)
				}
				return vv
			}
			_err := decoder.ReadValueInterface(jsonKey, initFN)
			if _err != nil {
				return _err
			}
			x.FTimestamp1 = vv
		case jsonKey == "f_message11":
			var vv *Message1
			initFN := func() interface{} {
				if x.FMessage11 != nil {
					vv = x.FMessage11
				} else {
					vv = new(Message1)
				}
				return vv
			}
			_err := decoder.ReadValueInterface(jsonKey, initFN)
			if _err != nil {
				return _err
			}
			x.FMessage11 = vv
		case jsonKey == "f_message12":
			var vv *Message1_Embed1
			initFN := func() interface{} {
				if x.FMessage12 != nil {
					vv = x.FMessage12
				} else {
					vv = new(Message1_Embed1)
				}
				return vv
			}
			_err := decoder.ReadValueInterface(jsonKey, initFN)
			if _err != nil {
				return _err
			}
			x.FMessage12 = vv
		case jsonKey == "f_message13":
			var vv *Message1_Embed1_Embed2
			initFN := func() interface{} {
				if x.FMessage13 != nil {
					vv = x.FMessage13
				} else {
					vv = new(Message1_Embed1_Embed2)
				}
				return vv
			}
			_err := decoder.ReadValueInterface(jsonKey, initFN)
			if _err != nil {
				return _err
			}
			x.FMessage13 = vv
		case jsonKey == "f_any1":
			var vv *anypb.Any
			initFN := func() interface{} {
				if x.FAny1 != nil {
					vv = x.FAny1
				} else {
					vv = new(anypb.Any)
				}
				return vv
			}
			_err := decoder.ReadValueInterface(jsonKey, initFN)
			if _err != nil {
				return _err
			}
			x.FAny1 = vv
		case jsonKey == "r_string1":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadArray(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.RString1 = nil
			case isEmpty:
				if x.RString1 == nil {
					x.RString1 = make([]string, 0)
				} else {
					x.RString1 = x.RString1[:0]
				}
			default:
				if x.RString1 == nil {
					x.RString1 = make([]string, 0)
				}
				i := 0
				length := len(x.RString1)
			LOOP_LIST_r_string1:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					vv, noMore, _err := decoder.ReadArrayElemString(jsonKey)
					if _err != nil {
						return _err
					}
					if i < length {
						x.RString1[i] = vv
					} else {
						x.RString1 = append(x.RString1, vv)
					}
					i++
					if noMore { // After read array value.
						break LOOP_LIST_r_string1
					}
				}
				if i < length {
					// truncate the slice to Consistent with standard library json.
					x.RString1 = x.RString1[:i]
				}
				decoder.ScanNext()
			}
		case jsonKey == "r_int32":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadArray(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.RInt32 = nil
			case isEmpty:
				if x.RInt32 == nil {
					x.RInt32 = make([]int32, 0)
				} else {
					x.RInt32 = x.RInt32[:0]
				}
			default:
				if x.RInt32 == nil {
					x.RInt32 = make([]int32, 0)
				}
				i := 0
				length := len(x.RInt32)
			LOOP_LIST_r_int32:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					vv, noMore, _err := decoder.ReadArrayElemInt32(jsonKey)
					if _err != nil {
						return _err
					}
					if i < length {
						x.RInt32[i] = vv
					} else {
						x.RInt32 = append(x.RInt32, vv)
					}
					i++
					if noMore { // After read array value.
						break LOOP_LIST_r_int32
					}
				}
				if i < length {
					// truncate the slice to Consistent with standard library json.
					x.RInt32 = x.RInt32[:i]
				}
				decoder.ScanNext()
			}
		case jsonKey == "r_message":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadArray(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.RMessage = nil
			case isEmpty:
				if x.RMessage == nil {
					x.RMessage = make([]*Message1, 0)
				} else {
					x.RMessage = x.RMessage[:0]
				}
			default:
				if x.RMessage == nil {
					x.RMessage = make([]*Message1, 0)
				}
				i := 0
				length := len(x.RMessage)
			LOOP_LIST_r_message:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					var vv *Message1
					initFN := func() interface{} {
						if i < length {
							vv = x.RMessage[i]
						}
						if vv == nil {
							vv = new(Message1)
						}
						return vv
					}
					noMore, _err := decoder.ReadArrayElemInterface(jsonKey, initFN)
					if _err != nil {
						return _err
					}
					if i < length {
						x.RMessage[i] = vv
					} else {
						x.RMessage = append(x.RMessage, vv)
					}
					i++
					if noMore { // After read array value.
						break LOOP_LIST_r_message
					}
				}
				if i < length {
					// truncate the slice to Consistent with standard library json.
					x.RMessage = x.RMessage[:i]
				}
				decoder.ScanNext()
			}
		case jsonKey == "r_enum":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadArray(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.REnum = nil
			case isEmpty:
				if x.REnum == nil {
					x.REnum = make([]Enum1, 0)
				} else {
					x.REnum = x.REnum[:0]
				}
			default:
				if x.REnum == nil {
					x.REnum = make([]Enum1, 0)
				}
				i := 0
				length := len(x.REnum)
			LOOP_LIST_r_enum:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					v1, noMore, _err := decoder.ReadArrayElemEnumNumber(jsonKey, Enum1_name)
					vv := Enum1(v1)
					if _err != nil {
						return _err
					}
					if i < length {
						x.REnum[i] = vv
					} else {
						x.REnum = append(x.REnum, vv)
					}
					i++
					if noMore { // After read array value.
						break LOOP_LIST_r_enum
					}
				}
				if i < length {
					// truncate the slice to Consistent with standard library json.
					x.REnum = x.REnum[:i]
				}
				decoder.ScanNext()
			}
		case jsonKey == "m_string1":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadMap(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.MString1 = nil
			case isEmpty:
				// do nothing
			default:
				if x.MString1 == nil {
					x.MString1 = make(map[string]string)
				}
			LOOP_MAP_m_string1:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					mapKey, _err := decoder.ReadMapKeyString(jsonKey)
					if _err != nil {
						return _err
					}
					vv, noMore, _err := decoder.ReadMapValueString(jsonKey)
					if _err != nil {
						return _err
					}
					x.MString1[mapKey] = vv
					if noMore {
						break LOOP_MAP_m_string1
					}
				}
				decoder.ScanNext()
			}
		case jsonKey == "m_string2":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadMap(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.MString2 = nil
			case isEmpty:
				// do nothing
			default:
				if x.MString2 == nil {
					x.MString2 = make(map[int32]string)
				}
			LOOP_MAP_m_string2:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					mapKey, _err := decoder.ReadMapKeyInt32(jsonKey)
					if _err != nil {
						return _err
					}
					vv, noMore, _err := decoder.ReadMapValueString(jsonKey)
					if _err != nil {
						return _err
					}
					x.MString2[mapKey] = vv
					if noMore {
						break LOOP_MAP_m_string2
					}
				}
				decoder.ScanNext()
			}
		case jsonKey == "m_message1":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadMap(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.MMessage1 = nil
			case isEmpty:
				// do nothing
			default:
				if x.MMessage1 == nil {
					x.MMessage1 = make(map[string]*Message1)
				}
			LOOP_MAP_m_message1:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					mapKey, _err := decoder.ReadMapKeyString(jsonKey)
					if _err != nil {
						return _err
					}
					var vv *Message1
					initFN := func() interface{} {
						vv = x.MMessage1[mapKey]
						if vv == nil {
							vv = new(Message1)
						}
						return vv
					}
					noMore, _err := decoder.ReadMapValueInterface(jsonKey, initFN)
					if _err != nil {
						return _err
					}
					x.MMessage1[mapKey] = vv
					if noMore {
						break LOOP_MAP_m_message1
					}
				}
				decoder.ScanNext()
			}
		case jsonKey == "m_enum":
			var isEmpty bool
			if isNULL, isEmpty, err = decoder.CheckBeforeReadMap(jsonKey); err != nil {
				return err
			}
			switch {
			case isNULL:
				x.MEnum = nil
			case isEmpty:
				// do nothing
			default:
				if x.MEnum == nil {
					x.MEnum = make(map[string]Enum1)
				}
			LOOP_MAP_m_enum:
				for {
					if err = decoder.ScanError(); err != nil {
						return err
					}
					mapKey, _err := decoder.ReadMapKeyString(jsonKey)
					if _err != nil {
						return _err
					}
					v1, noMore, _err := decoder.ReadMapValueEnumNumber(jsonKey, Enum1_name)
					vv := Enum1(v1)
					if _err != nil {
						return _err
					}
					x.MEnum[mapKey] = vv
					if noMore {
						break LOOP_MAP_m_enum
					}
				}
				decoder.ScanNext()
			}
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Message1 in file tests/proto/example/example1.proto
func (x *Message1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(46)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Message1 in file tests/proto/example/example1.proto
func (x *Message1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Message1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/example/example1.proto
func (x *Message1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(46)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/example/example1.proto
func (x *Message1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Message1_Embed1) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/example/example1.proto
func (x *Message1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(46)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/example/example1.proto
func (x *Message1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Message1_Embed1_Embed2) is nil")
	}

	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.CheckJSONBegin(); err != nil || isNULL {
		return err
	}
	// Loop to scan object.
LOOP_OBJECT:
	for {
		if err = decoder.ScanError(); err != nil {
			return err
		}
		jsonKey, stop := decoder.ReadJSONKey()
		if stop {
			break LOOP_OBJECT
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString1 = vv
		case jsonKey == "f_string2":
			vv, _err := decoder.ReadValueString(jsonKey)
			if _err != nil {
				return _err
			}
			x.FString2 = vv
		default:
			if err = decoder.Discard(); err != nil { // discard unknown field
				return err
			}
		}
		if decoder.ReadObjectValueAfter() { // After read object value
			break LOOP_OBJECT
		}
	}
	// Check error in decoder
	if err = decoder.ScanError(); err != nil {
		return err
	}
	return nil
}