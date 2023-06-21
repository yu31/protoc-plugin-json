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
	encoder := jsonencoder.New(904)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	switch ov := x.OnetType1.(type) {
	case *Example1_FString2:
		encoder.AppendObjectKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_string2")
		encoder.AppendLiteralString(ov.FString2)
		encoder.AppendObjectEnd()
	case *Example1_FString3:
		encoder.AppendObjectKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_string3")
		encoder.AppendLiteralString(ov.FString3)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("OnetType1")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OnetType2.(type) {
	case *Example1_FString4:
		encoder.AppendObjectKey("f_string4")
		encoder.AppendLiteralString(ov.FString4)
	case *Example1_FString5:
		encoder.AppendObjectKey("f_string5")
		encoder.AppendLiteralString(ov.FString5)
	default:
		_ = ov // to avoids unused panics
	}
	encoder.AppendObjectKey("f_int32")
	encoder.AppendLiteralInt32(x.FInt32, false)
	encoder.AppendObjectKey("f_int64")
	encoder.AppendLiteralInt64(x.FInt64, false)
	encoder.AppendObjectKey("f_uint32")
	encoder.AppendLiteralUint32(x.FUint32, false)
	encoder.AppendObjectKey("f_uint64")
	encoder.AppendLiteralUint64(x.FUint64, false)
	encoder.AppendObjectKey("f_sint32")
	encoder.AppendLiteralInt32(x.FSint32, false)
	encoder.AppendObjectKey("f_sint64")
	encoder.AppendLiteralInt64(x.FSint64, false)
	encoder.AppendObjectKey("f_sfixed32")
	encoder.AppendLiteralInt32(x.FSfixed32, false)
	encoder.AppendObjectKey("f_sfixed64")
	encoder.AppendLiteralInt64(x.FSfixed64, false)
	encoder.AppendObjectKey("f_fixed32")
	encoder.AppendLiteralUint32(x.FFixed32, false)
	encoder.AppendObjectKey("f_fixed64")
	encoder.AppendLiteralUint64(x.FFixed64, false)
	encoder.AppendObjectKey("f_float")
	encoder.AppendLiteralFloat32(x.FFloat, false)
	encoder.AppendObjectKey("f_double")
	encoder.AppendLiteralFloat64(x.FDouble, false)
	encoder.AppendObjectKey("f_bool1")
	encoder.AppendLiteralBool(x.FBool1, false)
	encoder.AppendObjectKey("f_bytes1")
	encoder.AppendLiteralBytes(x.FBytes1)
	if x.FEnum1 != 0 {
		encoder.AppendObjectKey("f_enum1")
		encoder.AppendLiteralInt32(int32(x.FEnum1.Number()), false)
	}
	if x.FEnum2 != 0 {
		encoder.AppendObjectKey("f_enum2")
		encoder.AppendLiteralString(x.FEnum2.String())
	}
	encoder.AppendObjectKey("f_enum3")
	encoder.AppendLiteralInt32(int32(x.FEnum3.Number()), false)
	encoder.AppendObjectKey("f_enum4")
	encoder.AppendLiteralString(x.FEnum4.String())
	if x.FEnum5 != nil {
		encoder.AppendObjectKey("f_enum5")
		if x.FEnum5 != nil {
			encoder.AppendLiteralInt32(int32(x.FEnum5.Number()), false)
		} else {
			encoder.AppendLiteralNULL()
		}
	}
	if x.FEnum6 != nil {
		encoder.AppendObjectKey("f_enum6")
		if x.FEnum6 != nil {
			encoder.AppendLiteralString(x.FEnum6.String())
		} else {
			encoder.AppendLiteralNULL()
		}
	}
	encoder.AppendObjectKey("f_enum7")
	if x.FEnum7 != nil {
		encoder.AppendLiteralInt32(int32(x.FEnum7.Number()), false)
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("f_enum8")
	if x.FEnum8 != nil {
		encoder.AppendLiteralString(x.FEnum8.String())
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("f_duration1")
	if err = encoder.AppendLiteralInterface(x.FDuration1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_timestamp1")
	if err = encoder.AppendLiteralInterface(x.FTimestamp1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message11")
	if err = encoder.AppendLiteralInterface(x.FMessage11); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message12")
	if err = encoder.AppendLiteralInterface(x.FMessage12); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_message13")
	if err = encoder.AppendLiteralInterface(x.FMessage13); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_any1")
	if err = encoder.AppendLiteralInterface(x.FAny1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("r_string1")
	if x.RString1 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RString1 {
			encoder.AppendLiteralString(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("r_int32")
	if x.RInt32 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt32 {
			encoder.AppendLiteralInt32(ri, false)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("r_message")
	if x.RMessage != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RMessage {
			if err = encoder.AppendLiteralInterface(ri); err != nil {
				return nil, err
			}
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("r_enum")
	if x.REnum != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.REnum {
			encoder.AppendLiteralInt32(int32(ri.Number()), false)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("m_string1")
	if x.MString1 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MString1 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralString(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("m_string2")
	if x.MString2 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MString2 {
			encoder.AppendMapKeyInt32(mk, true)
			encoder.AppendLiteralString(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("m_message1")
	if x.MMessage1 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MMessage1 {
			encoder.AppendMapKeyString(mk)
			if err = encoder.AppendLiteralInterface(mv); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendLiteralNULL()
	}
	encoder.AppendObjectKey("m_enum")
	if x.MEnum != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MEnum {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralInt32(int32(mv.Number()), false)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendLiteralNULL()
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
	var (
		oneOfIsFill_OnetType1 bool
		oneOfIsFill_OnetType2 bool
	)
	var (
		err     error
		isNULL  bool
		decoder *jsondecoder.Decoder
	)
	if decoder, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "OnetType1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OnetType1 = nil
				continue LOOP_SCAN
			}
			for {
				var oneofKey string
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				if oneofKey, err = decoder.ReadObjectKey(jsonKey); err != nil {
					return err
				}
				switch oneofKey { // match oneof key
				case "f_string2":
					if oneOfIsFill_OnetType1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OnetType1 = true

					var ok bool
					var ot *Example1_FString2
					if ot, ok = x.OnetType1.(*Example1_FString2); !ok {
						ot = new(Example1_FString2)
					}
					var vv string
					if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
						return err
					}
					ot.FString2 = vv
					x.OnetType1 = ot
				case "f_string3":
					if oneOfIsFill_OnetType1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OnetType1 = true

					var ok bool
					var ot *Example1_FString3
					if ot, ok = x.OnetType1.(*Example1_FString3); !ok {
						ot = new(Example1_FString3)
					}
					var vv string
					if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
						return err
					}
					ot.FString3 = vv
					x.OnetType1 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_string4":
			if oneOfIsFill_OnetType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OnetType2 = true

			var ok bool
			var ot *Example1_FString4
			if ot, ok = x.OnetType2.(*Example1_FString4); !ok {
				ot = new(Example1_FString4)
			}
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			ot.FString4 = vv
			x.OnetType2 = ot
		case "f_string5":
			if oneOfIsFill_OnetType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OnetType2 = true

			var ok bool
			var ot *Example1_FString5
			if ot, ok = x.OnetType2.(*Example1_FString5); !ok {
				ot = new(Example1_FString5)
			}
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			ot.FString5 = vv
			x.OnetType2 = ot
		case "f_int32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FInt32 = vv
		case "f_int64":
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			x.FInt64 = vv
		case "f_uint32":
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			x.FUint32 = vv
		case "f_uint64":
			var vv uint64
			if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
				return err
			}
			x.FUint64 = vv
		case "f_sint32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FSint32 = vv
		case "f_sint64":
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			x.FSint64 = vv
		case "f_sfixed32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FSfixed32 = vv
		case "f_sfixed64":
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			x.FSfixed64 = vv
		case "f_fixed32":
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			x.FFixed32 = vv
		case "f_fixed64":
			var vv uint64
			if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
				return err
			}
			x.FFixed64 = vv
		case "f_float":
			var vv float32
			if vv, err = decoder.ReadLiteralFloat32(jsonKey, false); err != nil {
				return err
			}
			x.FFloat = vv
		case "f_double":
			var vv float64
			if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
				return err
			}
			x.FDouble = vv
		case "f_bool1":
			var vv bool
			if vv, err = decoder.ReadLiteralBool(jsonKey, false); err != nil {
				return err
			}
			x.FBool1 = vv
		case "f_bytes1":
			var vv []byte
			if vv, err = decoder.ReadLiteralBytes(jsonKey); err != nil {
				return err
			}
			x.FBytes1 = vv
		case "f_enum1":
			var vv Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = Enum1(v1)
			x.FEnum1 = vv
		case "f_enum2":
			var vv Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumString(jsonKey, Enum1_value); err != nil {
				return err
			}
			vv = Enum1(v1)
			x.FEnum2 = vv
		case "f_enum3":
			var vv Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = Enum1(v1)
			x.FEnum3 = vv
		case "f_enum4":
			var vv Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumString(jsonKey, Enum1_value); err != nil {
				return err
			}
			vv = Enum1(v1)
			x.FEnum4 = vv
		case "f_enum5":
			var vv *Enum1
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumNumber(jsonKey, false); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum1(*v1)
				vv = &v2
			}
			x.FEnum5 = vv
		case "f_enum6":
			var vv *Enum1
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumString(jsonKey, Enum1_value); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum1(*v1)
				vv = &v2
			}
			x.FEnum6 = vv
		case "f_enum7":
			var vv *Enum1
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumNumber(jsonKey, false); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum1(*v1)
				vv = &v2
			}
			x.FEnum7 = vv
		case "f_enum8":
			var vv *Enum1
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumString(jsonKey, Enum1_value); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum1(*v1)
				vv = &v2
			}
			x.FEnum8 = vv
		case "f_duration1":
			var vv *durationpb.Duration
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FDuration1 != nil {
					vv = x.FDuration1
				} else {
					vv = new(durationpb.Duration)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FDuration1 = vv
		case "f_timestamp1":
			var vv *timestamppb.Timestamp
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FTimestamp1 != nil {
					vv = x.FTimestamp1
				} else {
					vv = new(timestamppb.Timestamp)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FTimestamp1 = vv
		case "f_message11":
			var vv *Message1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage11 != nil {
					vv = x.FMessage11
				} else {
					vv = new(Message1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage11 = vv
		case "f_message12":
			var vv *Message1_Embed1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage12 != nil {
					vv = x.FMessage12
				} else {
					vv = new(Message1_Embed1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage12 = vv
		case "f_message13":
			var vv *Message1_Embed1_Embed2
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage13 != nil {
					vv = x.FMessage13
				} else {
					vv = new(Message1_Embed1_Embed2)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage13 = vv
		case "f_any1":
			var vv *anypb.Any
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FAny1 != nil {
					vv = x.FAny1
				} else {
					vv = new(anypb.Any)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FAny1 = vv
		case "r_string1":
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
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv string
				if i >= len(x.RString1) {
					x.RString1 = append(x.RString1, vv)
				}
				if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
					return err
				}
				x.RString1[i] = vv
				i++
			}
			// truncate the slice to consistent with standard library json.
			x.RString1 = x.RString1[:i]
		case "r_int32":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RInt32 = nil
				continue LOOP_SCAN
			}
			if x.RInt32 == nil {
				x.RInt32 = make([]int32, 0)
			}
			i := 0
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv int32
				if i >= len(x.RInt32) {
					x.RInt32 = append(x.RInt32, vv)
				}
				if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
					return err
				}
				x.RInt32[i] = vv
				i++
			}
			// truncate the slice to consistent with standard library json.
			x.RInt32 = x.RInt32[:i]
		case "r_message":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RMessage = nil
				continue LOOP_SCAN
			}
			if x.RMessage == nil {
				x.RMessage = make([]*Message1, 0)
			}
			i := 0
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv *Message1
				if i >= len(x.RMessage) {
					x.RMessage = append(x.RMessage, vv)
				}
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.RMessage[i] != nil {
						vv = x.RMessage[i]
					} else {
						vv = new(Message1)
					}
					if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
						return err
					}
				}
				x.RMessage[i] = vv
				i++
			}
			// truncate the slice to consistent with standard library json.
			x.RMessage = x.RMessage[:i]
		case "r_enum":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.REnum = nil
				continue LOOP_SCAN
			}
			if x.REnum == nil {
				x.REnum = make([]Enum1, 0)
			}
			i := 0
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv Enum1
				if i >= len(x.REnum) {
					x.REnum = append(x.REnum, vv)
				}
				var v1 int32
				if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
					return err
				}
				vv = Enum1(v1)
				x.REnum[i] = vv
				i++
			}
			// truncate the slice to consistent with standard library json.
			x.REnum = x.REnum[:i]
		case "m_string1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MString1 = nil
				continue LOOP_SCAN
			}
			if x.MString1 == nil {
				x.MString1 = make(map[string]string)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv string
				if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
					return err
				}
				x.MString1[mk] = vv
			}
		case "m_string2":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MString2 = nil
				continue LOOP_SCAN
			}
			if x.MString2 == nil {
				x.MString2 = make(map[int32]string)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int32
				if mk, err = decoder.ReadMapKeyInt32(jsonKey, true); err != nil {
					return err
				}
				var vv string
				if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
					return err
				}
				x.MString2[mk] = vv
			}
		case "m_message1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MMessage1 = nil
				continue LOOP_SCAN
			}
			if x.MMessage1 == nil {
				x.MMessage1 = make(map[string]*Message1)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv *Message1
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.MMessage1[mk] != nil {
						vv = x.MMessage1[mk]
					} else {
						vv = new(Message1)
					}
					if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
						return err
					}
				}
				x.MMessage1[mk] = vv
			}
		case "m_enum":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MEnum = nil
				continue LOOP_SCAN
			}
			if x.MEnum == nil {
				x.MEnum = make(map[string]Enum1)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv Enum1
				var v1 int32
				if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
					return err
				}
				vv = Enum1(v1)
				x.MEnum[mk] = vv
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Message1 in file tests/proto/example/example1.proto
func (x *Message1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(56)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)

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
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString2 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/example/example1.proto
func (x *Message1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(56)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)

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
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString2 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/example/example1.proto
func (x *Message1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(56)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	encoder.AppendObjectKey("f_string2")
	encoder.AppendLiteralString(x.FString2)

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
	if isNULL, err = decoder.BeforeReadJSON(); err != nil {
		return err
	}
	if isNULL {
		return nil
	}
LOOP_SCAN:
	for { // Loop to scan object.
		var (
			jsonKey string
			isEnd   bool
		)
		if isEnd, err = decoder.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "f_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString2 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
