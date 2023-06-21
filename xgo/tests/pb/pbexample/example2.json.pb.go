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
	encoder := jsonencoder.New(520)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendLiteralString(x.FString1)
	switch ov := x.OnetType1.(type) {
	case *Example2_FString2:
		encoder.AppendObjectKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_string2")
		encoder.AppendLiteralString(ov.FString2)
		encoder.AppendObjectEnd()
	case *Example2_FMessage3:
		encoder.AppendObjectKey("OnetType1")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_message3")
		if err = encoder.AppendLiteralInterface(ov.FMessage3); err != nil {
			return nil, err
		}
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("OnetType1")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OnetType2.(type) {
	case *Example2_FString4:
		encoder.AppendObjectKey("f_string4")
		encoder.AppendLiteralString(ov.FString4)
	case *Example2_FString5:
		encoder.AppendObjectKey("f_string5")
		encoder.AppendLiteralString(ov.FString5)
	default:
		_ = ov // to avoids unused panics
	}
	if x.FEnum1 != 0 {
		encoder.AppendObjectKey("f_enum1")
		encoder.AppendLiteralInt32(int32(x.FEnum1.Number()), false)
	}
	if x.FEnum2 != 0 {
		encoder.AppendObjectKey("f_enum2")
		encoder.AppendLiteralString(x.FEnum2.String())
	}
	if x.FEnum3 != 0 {
		encoder.AppendObjectKey("f_enum3")
		encoder.AppendLiteralInt32(int32(x.FEnum3.Number()), true)
	}
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
	if x.FEnum7 != nil {
		encoder.AppendObjectKey("f_enum7")
		if x.FEnum7 != nil {
			encoder.AppendLiteralInt32(int32(x.FEnum7.Number()), true)
		} else {
			encoder.AppendLiteralNULL()
		}
	}
	encoder.AppendObjectKey("f_duration1")
	if err = encoder.AppendLiteralInterface(x.FDuration1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_timestamp1")
	encoder.AppendLiteralInt64(x.FTimestamp1.AsTime().Unix(), false)
	encoder.AppendObjectKey("f_message11")
	if err = encoder.AppendLiteralInterface(x.FMessage11); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_any1")
	if err = encoder.AppendLiteralInterface(x.FAny1); err != nil {
		return nil, err
	}
	encoder.AppendObjectKey("f_int32")
	encoder.AppendLiteralInt32(x.FInt32, false)
	encoder.AppendObjectKey("r_int32")
	if x.RInt32 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt32 {
			encoder.AppendLiteralString(ri)
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
	encoder.AppendObjectKey("m_int32")
	if x.MInt32 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MInt32 {
			encoder.AppendMapKeyInt32(mk, false)
			encoder.AppendLiteralInt32(mv, false)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendLiteralNULL()
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
					var ot *Example2_FString2
					if ot, ok = x.OnetType1.(*Example2_FString2); !ok {
						ot = new(Example2_FString2)
					}
					var vv string
					if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
						return err
					}
					ot.FString2 = vv
					x.OnetType1 = ot
				case "f_message3":
					if oneOfIsFill_OnetType1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OnetType1 = true

					var ok bool
					var ot *Example2_FMessage3
					if ot, ok = x.OnetType1.(*Example2_FMessage3); !ok {
						ot = new(Example2_FMessage3)
					}
					var vv *Empty2
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FMessage3 != nil {
							vv = ot.FMessage3
						} else {
							vv = new(Empty2)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FMessage3 = vv
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
			var ot *Example2_FString4
			if ot, ok = x.OnetType2.(*Example2_FString4); !ok {
				ot = new(Example2_FString4)
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
			var ot *Example2_FString5
			if ot, ok = x.OnetType2.(*Example2_FString5); !ok {
				ot = new(Example2_FString5)
			}
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			ot.FString5 = vv
			x.OnetType2 = ot
		case "f_enum1":
			var vv Enum2
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = Enum2(v1)
			x.FEnum1 = vv
		case "f_enum2":
			var vv Enum2
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumString(jsonKey, Enum2_value); err != nil {
				return err
			}
			vv = Enum2(v1)
			x.FEnum2 = vv
		case "f_enum3":
			var vv Enum2
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, true); err != nil {
				return err
			}
			vv = Enum2(v1)
			x.FEnum3 = vv
		case "f_enum5":
			var vv *Enum2
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumNumber(jsonKey, false); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum2(*v1)
				vv = &v2
			}
			x.FEnum5 = vv
		case "f_enum6":
			var vv *Enum2
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumString(jsonKey, Enum2_value); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum2(*v1)
				vv = &v2
			}
			x.FEnum6 = vv
		case "f_enum7":
			var vv *Enum2
			var v1 *int32
			if v1, err = decoder.ReadPointerEnumNumber(jsonKey, true); err != nil {
				return err
			}
			if v1 != nil {
				v2 := Enum2(*v1)
				vv = &v2
			}
			x.FEnum7 = vv
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
			if x.FTimestamp1 != nil {
				vv = x.FTimestamp1
			} else {
				vv = new(timestamppb.Timestamp)
			}
			if err = decoder.ReadWKTTimestampByUnixSec(jsonKey, vv, false); err != nil {
				return err
			}
			x.FTimestamp1 = vv
		case "f_message11":
			var vv *Empty2
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if x.FMessage11 != nil {
					vv = x.FMessage11
				} else {
					vv = new(Empty2)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			x.FMessage11 = vv
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
		case "f_int32":
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			x.FInt32 = vv
		case "r_int32":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RInt32 = nil
				continue LOOP_SCAN
			}
			if x.RInt32 == nil {
				x.RInt32 = make([]string, 0)
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
				if i >= len(x.RInt32) {
					x.RInt32 = append(x.RInt32, vv)
				}
				if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
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
				x.RMessage = make([]*Empty2, 0)
			}
			i := 0
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv *Empty2
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
						vv = new(Empty2)
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
				x.REnum = make([]Enum2, 0)
			}
			i := 0
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv Enum2
				if i >= len(x.REnum) {
					x.REnum = append(x.REnum, vv)
				}
				var v1 int32
				if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
					return err
				}
				vv = Enum2(v1)
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
		case "m_message1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MMessage1 = nil
				continue LOOP_SCAN
			}
			if x.MMessage1 == nil {
				x.MMessage1 = make(map[string]*Empty2)
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
				var vv *Empty2
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.MMessage1[mk] != nil {
						vv = x.MMessage1[mk]
					} else {
						vv = new(Empty2)
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
				x.MEnum = make(map[string]Enum2)
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
				var vv Enum2
				var v1 int32
				if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
					return err
				}
				vv = Enum2(v1)
				x.MEnum[mk] = vv
			}
		case "m_int32":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MInt32 = nil
				continue LOOP_SCAN
			}
			if x.MInt32 == nil {
				x.MInt32 = make(map[int32]int32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int32
				if mk, err = decoder.ReadMapKeyInt32(jsonKey, false); err != nil {
					return err
				}
				var vv int32
				if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
					return err
				}
				x.MInt32[mk] = vv
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
