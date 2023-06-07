// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/example/example2.proto

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

// MarshalJSON implements interface json.Marshaler for proto message Empty2 in file tests/proto/example/example2.proto
func (x *Empty2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	return []byte("{}"), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Empty2 in file tests/proto/example/example2.proto
func (x *Empty2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Empty2) is nil")
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Example2 in file tests/proto/example/example2.proto
func (x *Example2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(380)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	switch ov := x.OnetType1.(type) {
	case *Example2_FString2:
		encoder.AppendJSONKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendJSONKey("f_string2")
		encoder.AppendValueString(ov.FString2)
		encoder.AppendObjectEnd()
	case *Example2_FMessage3:
		encoder.AppendJSONKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendJSONKey("f_message3")
		err = encoder.AppendValueInterface(ov.FMessage3)
		if err != nil {
			return nil, err
		}
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendJSONKey("OnetType1")
		encoder.AppendValueNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OnetType2.(type) {
	case *Example2_FString4:
		encoder.AppendJSONKey("f_string4")
		encoder.AppendValueString(ov.FString4)
	case *Example2_FString5:
		encoder.AppendJSONKey("f_string5")
		encoder.AppendValueString(ov.FString5)
	default:
		_ = ov // to avoids unused panics
	}
	encoder.AppendJSONKey("f_int32")
	encoder.AppendValueInt32(x.FInt32)
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

// UnmarshalJSON implements json.Unmarshaler for proto message Example2 in file tests/proto/example/example2.proto
func (x *Example2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbexample.(*Example2) is nil")
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
	if isNULL, err = decoder.BeforeScanJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
	// Loop to scan object.
LOOP_SCAN:
	for {
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch { // match the JSON KEY
		case jsonKey == "f_string1":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString1 = vv
		case jsonKey == "OnetType1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OnetType1 = nil
				continue LOOP_SCAN
			}
		LOOP_ONEOF_OnetType1:
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_ONEOF_OnetType1
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch {
				case oneofKey == "f_string2":
					var ok bool
					var ot *Example2_FString2
					if ot, ok = x.OnetType1.(*Example2_FString2); !ok {
						ot = new(Example2_FString2)
					}
					var vv string
					vv, err = decoder.ReadValueString(jsonKey)
					if err != nil {
						return err
					}
					if oneOfIsFill_OnetType1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OnetType1 = true
					ot.FString2 = vv
					x.OnetType1 = ot
				case oneofKey == "f_message3":
					var ok bool
					var ot *Example2_FMessage3
					if ot, ok = x.OnetType1.(*Example2_FMessage3); !ok {
						ot = new(Example2_FMessage3)
					}
					var vv *Empty2
					initFN := func() interface{} {
						if ot.FMessage3 != nil {
							vv = ot.FMessage3
						} else {
							vv = new(Empty2)
						}
						return vv
					}
					err = decoder.ReadValueInterface(jsonKey, initFN)
					if err != nil {
						return err
					}
					if oneOfIsFill_OnetType1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OnetType1 = true
					ot.FMessage3 = vv
					x.OnetType1 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				}
			}
		case jsonKey == "f_string4":
			var ok bool
			var ot *Example2_FString4
			if ot, ok = x.OnetType2.(*Example2_FString4); !ok {
				ot = new(Example2_FString4)
			}
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			if oneOfIsFill_OnetType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OnetType2 = true
			ot.FString4 = vv
			x.OnetType2 = ot
		case jsonKey == "f_string5":
			var ok bool
			var ot *Example2_FString5
			if ot, ok = x.OnetType2.(*Example2_FString5); !ok {
				ot = new(Example2_FString5)
			}
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			if oneOfIsFill_OnetType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OnetType2 = true
			ot.FString5 = vv
			x.OnetType2 = ot
		case jsonKey == "f_int32":
			var vv int32
			vv, err = decoder.ReadValueInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FInt32 = vv
		case jsonKey == "f_enum1":
			var vv Enum2
			var v1 int32
			v1, err = decoder.ReadValueEnumString(jsonKey, Enum2_value)
			vv = Enum2(v1)
			if err != nil {
				return err
			}
			x.FEnum1 = vv
		case jsonKey == "f_enum2":
			var vv Enum2
			var v1 int32
			v1, err = decoder.ReadValueEnumNumber(jsonKey, Enum2_name)
			vv = Enum2(v1)
			if err != nil {
				return err
			}
			x.FEnum2 = vv
		case jsonKey == "f_enum5":
			var vv *Enum2
			var v1 *int32
			v1, err = decoder.ReadPointerEnumString(jsonKey, Enum2_value)
			if v1 != nil {
				v2 := Enum2(*v1)
				vv = &v2
			}
			if err != nil {
				return err
			}
			x.FEnum5 = vv
		case jsonKey == "f_enum6":
			var vv *Enum2
			var v1 *int32
			v1, err = decoder.ReadPointerEnumNumber(jsonKey, Enum2_name)
			if v1 != nil {
				v2 := Enum2(*v1)
				vv = &v2
			}
			if err != nil {
				return err
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
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
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
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FTimestamp1 = vv
		case jsonKey == "f_message11":
			var vv *Empty2
			initFN := func() interface{} {
				if x.FMessage11 != nil {
					vv = x.FMessage11
				} else {
					vv = new(Empty2)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage11 = vv
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
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FAny1 = vv
		case jsonKey == "r_string1":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RString1 = nil
				continue LOOP_SCAN
			}
			if x.RString1 == nil {
				x.RString1 = make([]string, 0)
			}
			i := 0
			length := len(x.RString1)
		LOOP_LIST_r_string1:
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_LIST_r_string1
				}
				var vv string
				vv, err = decoder.ReadArrayElemString(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RString1[i] = vv
				} else {
					x.RString1 = append(x.RString1, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RString1 = x.RString1[:i]
			}
		case jsonKey == "r_message":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RMessage = nil
				continue LOOP_SCAN
			}
			if x.RMessage == nil {
				x.RMessage = make([]*Empty2, 0)
			}
			i := 0
			length := len(x.RMessage)
		LOOP_LIST_r_message:
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_LIST_r_message
				}
				var vv *Empty2
				initFN := func() interface{} {
					if i < length {
						vv = x.RMessage[i]
					}
					if vv == nil {
						vv = new(Empty2)
					}
					return vv
				}
				err = decoder.ReadArrayElemInterface(jsonKey, initFN)
				if err != nil {
					return err
				}
				if i < length {
					x.RMessage[i] = vv
				} else {
					x.RMessage = append(x.RMessage, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RMessage = x.RMessage[:i]
			}
		case jsonKey == "r_enum":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.REnum = nil
				continue LOOP_SCAN
			}
			if x.REnum == nil {
				x.REnum = make([]Enum2, 0)
			}
			i := 0
			length := len(x.REnum)
		LOOP_LIST_r_enum:
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_LIST_r_enum
				}
				var vv Enum2
				var v1 int32
				v1, err = decoder.ReadArrayElemEnumNumber(jsonKey, Enum2_name)
				vv = Enum2(v1)
				if err != nil {
					return err
				}
				if i < length {
					x.REnum[i] = vv
				} else {
					x.REnum = append(x.REnum, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.REnum = x.REnum[:i]
			}
		case jsonKey == "m_string1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MString1 = nil
				continue LOOP_SCAN
			}
			if x.MString1 == nil {
				x.MString1 = make(map[string]string)
			}
		LOOP_MAP_m_string1:
			for {
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_MAP_m_string1
				}
				var mapKey string
				mapKey, err = decoder.ReadMapKeyString(jsonKey)
				if err != nil {
					return err
				}
				var mapVal string
				mapVal, err = decoder.ReadMapValueString(jsonKey)
				if err != nil {
					return err
				}
				x.MString1[mapKey] = mapVal
			}
		case jsonKey == "m_message1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MMessage1 = nil
				continue LOOP_SCAN
			}
			if x.MMessage1 == nil {
				x.MMessage1 = make(map[string]*Empty2)
			}
		LOOP_MAP_m_message1:
			for {
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_MAP_m_message1
				}
				var mapKey string
				mapKey, err = decoder.ReadMapKeyString(jsonKey)
				if err != nil {
					return err
				}
				var mapVal *Empty2
				initFN := func() interface{} {
					mapVal = x.MMessage1[mapKey]
					if mapVal == nil {
						mapVal = new(Empty2)
					}
					return mapVal
				}
				err = decoder.ReadMapValueInterface(jsonKey, initFN)
				if err != nil {
					return err
				}
				x.MMessage1[mapKey] = mapVal
			}
		case jsonKey == "m_enum":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MEnum = nil
				continue LOOP_SCAN
			}
			if x.MEnum == nil {
				x.MEnum = make(map[string]Enum2)
			}
		LOOP_MAP_m_enum:
			for {
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break LOOP_MAP_m_enum
				}
				var mapKey string
				mapKey, err = decoder.ReadMapKeyString(jsonKey)
				if err != nil {
					return err
				}
				var mapVal Enum2
				var v1 int32
				v1, err = decoder.ReadMapValueEnumNumber(jsonKey, Enum2_name)
				mapVal = Enum2(v1)
				if err != nil {
					return err
				}
				x.MEnum[mapKey] = mapVal
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}
