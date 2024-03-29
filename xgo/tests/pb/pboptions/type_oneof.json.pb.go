// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/options/type_oneof.proto

package pboptions

import (
	errors "errors"
	fmt "fmt"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	pbexternal "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeOneOf1 in file tests/proto/cases/options/type_oneof.proto
func (x *TypeOneOf1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(768)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	switch ov := x.OneType1.(type) {
	case *TypeOneOf1_FInt32:
		encoder.AppendObjectKey("t_one1")
		encoder.AppendObjectBegin()
		if ov.FInt32 != 0 {
			encoder.AppendObjectKey("t_int32")
			encoder.AppendLiteralInt32(ov.FInt32, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType2.(type) {
	case *TypeOneOf1_FInt64:
		encoder.AppendObjectKey("t_one2")
		encoder.AppendObjectBegin()
		if ov.FInt64 != 0 {
			encoder.AppendObjectKey("t_int64")
			encoder.AppendLiteralInt64(ov.FInt64, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType3.(type) {
	case *TypeOneOf1_FUint32:
		encoder.AppendObjectKey("t_one3")
		encoder.AppendObjectBegin()
		if ov.FUint32 != 0 {
			encoder.AppendObjectKey("t_uint32")
			encoder.AppendLiteralUint32(ov.FUint32, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType4.(type) {
	case *TypeOneOf1_FUint64:
		encoder.AppendObjectKey("t_one4")
		encoder.AppendObjectBegin()
		if ov.FUint64 != 0 {
			encoder.AppendObjectKey("t_uint64")
			encoder.AppendLiteralUint64(ov.FUint64, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType5.(type) {
	case *TypeOneOf1_FSint32:
		encoder.AppendObjectKey("t_one5")
		encoder.AppendObjectBegin()
		if ov.FSint32 != 0 {
			encoder.AppendObjectKey("t_sint32")
			encoder.AppendLiteralInt32(ov.FSint32, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType6.(type) {
	case *TypeOneOf1_FSint64:
		encoder.AppendObjectKey("t_one6")
		encoder.AppendObjectBegin()
		if ov.FSint64 != 0 {
			encoder.AppendObjectKey("t_sint64")
			encoder.AppendLiteralInt64(ov.FSint64, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType7.(type) {
	case *TypeOneOf1_FSfixed32:
		encoder.AppendObjectKey("t_one7")
		encoder.AppendObjectBegin()
		if ov.FSfixed32 != 0 {
			encoder.AppendObjectKey("t_sfixed32")
			encoder.AppendLiteralInt32(ov.FSfixed32, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType8.(type) {
	case *TypeOneOf1_FSfixed64:
		encoder.AppendObjectKey("t_one8")
		encoder.AppendObjectBegin()
		if ov.FSfixed64 != 0 {
			encoder.AppendObjectKey("t_sfixed64")
			encoder.AppendLiteralInt64(ov.FSfixed64, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType9.(type) {
	case *TypeOneOf1_FFixed32:
		encoder.AppendObjectKey("t_one9")
		encoder.AppendObjectBegin()
		if ov.FFixed32 != 0 {
			encoder.AppendObjectKey("t_fixed32")
			encoder.AppendLiteralUint32(ov.FFixed32, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType10.(type) {
	case *TypeOneOf1_FFixed64:
		encoder.AppendObjectKey("t_one10")
		encoder.AppendObjectBegin()
		if ov.FFixed64 != 0 {
			encoder.AppendObjectKey("t_fixed64")
			encoder.AppendLiteralUint64(ov.FFixed64, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType11.(type) {
	case *TypeOneOf1_FFloat:
		encoder.AppendObjectKey("t_one11")
		encoder.AppendObjectBegin()
		if ov.FFloat != 0 {
			encoder.AppendObjectKey("t_float")
			encoder.AppendLiteralFloat32(ov.FFloat, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType12.(type) {
	case *TypeOneOf1_FDouble:
		encoder.AppendObjectKey("t_one12")
		encoder.AppendObjectBegin()
		if ov.FDouble != 0 {
			encoder.AppendObjectKey("t_double")
			encoder.AppendLiteralFloat64(ov.FDouble, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType13.(type) {
	case *TypeOneOf1_FBool1:
		encoder.AppendObjectKey("t_one13")
		encoder.AppendObjectBegin()
		if ov.FBool1 {
			encoder.AppendObjectKey("t_bool1")
			encoder.AppendLiteralBool(ov.FBool1, false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType14.(type) {
	case *TypeOneOf1_FString1:
		encoder.AppendObjectKey("t_one14")
		encoder.AppendObjectBegin()
		if ov.FString1 != "" {
			encoder.AppendObjectKey("t_string1")
			encoder.AppendLiteralString(ov.FString1)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType15.(type) {
	case *TypeOneOf1_FBytes1:
		encoder.AppendObjectKey("t_one15")
		encoder.AppendObjectBegin()
		if len(ov.FBytes1) != 0 {
			encoder.AppendObjectKey("t_bytes1")
			encoder.AppendLiteralBytes(ov.FBytes1)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType16.(type) {
	case *TypeOneOf1_FEnum1:
		encoder.AppendObjectKey("t_one16")
		encoder.AppendObjectBegin()
		if ov.FEnum1 != 0 {
			encoder.AppendObjectKey("t_enum1")
			encoder.AppendLiteralInt32(int32(ov.FEnum1.Number()), false)
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType17.(type) {
	case *TypeOneOf1_FMessage1:
		encoder.AppendObjectKey("t_one17")
		encoder.AppendObjectBegin()
		if ov.FMessage1 != nil {
			encoder.AppendObjectKey("t_message1")
			if err = encoder.AppendLiteralInterface(ov.FMessage1); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType18.(type) {
	case *TypeOneOf1_FAny1:
		encoder.AppendObjectKey("t_one18")
		encoder.AppendObjectBegin()
		if ov.FAny1 != nil {
			encoder.AppendObjectKey("t_any1")
			if err = encoder.AppendLiteralInterface(ov.FAny1); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType19.(type) {
	case *TypeOneOf1_FDuration1:
		encoder.AppendObjectKey("t_one19")
		encoder.AppendObjectBegin()
		if ov.FDuration1 != nil {
			encoder.AppendObjectKey("t_duration1")
			if err = encoder.AppendLiteralInterface(ov.FDuration1); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType20.(type) {
	case *TypeOneOf1_FTimestamp1:
		encoder.AppendObjectKey("t_one20")
		encoder.AppendObjectBegin()
		if ov.FTimestamp1 != nil {
			encoder.AppendObjectKey("t_timestamp1")
			if err = encoder.AppendLiteralInterface(ov.FTimestamp1); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOneOf1 in file tests/proto/cases/options/type_oneof.proto
func (x *TypeOneOf1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*TypeOneOf1) is nil")
	}
	var (
		oneOfIsFill_OneType1  bool
		oneOfIsFill_OneType2  bool
		oneOfIsFill_OneType3  bool
		oneOfIsFill_OneType4  bool
		oneOfIsFill_OneType5  bool
		oneOfIsFill_OneType6  bool
		oneOfIsFill_OneType7  bool
		oneOfIsFill_OneType8  bool
		oneOfIsFill_OneType9  bool
		oneOfIsFill_OneType10 bool
		oneOfIsFill_OneType11 bool
		oneOfIsFill_OneType12 bool
		oneOfIsFill_OneType13 bool
		oneOfIsFill_OneType14 bool
		oneOfIsFill_OneType15 bool
		oneOfIsFill_OneType16 bool
		oneOfIsFill_OneType17 bool
		oneOfIsFill_OneType18 bool
		oneOfIsFill_OneType19 bool
		oneOfIsFill_OneType20 bool
	)
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
		case "t_one1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType1 = nil
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
				case "t_int32":
					if oneOfIsFill_OneType1 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType1 = true

					var ok bool
					var ot *TypeOneOf1_FInt32
					if ot, ok = x.OneType1.(*TypeOneOf1_FInt32); !ok {
						ot = new(TypeOneOf1_FInt32)
					}
					var vv int32
					if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
						return err
					}
					ot.FInt32 = vv
					x.OneType1 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one2":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType2 = nil
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
				case "t_int64":
					if oneOfIsFill_OneType2 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType2 = true

					var ok bool
					var ot *TypeOneOf1_FInt64
					if ot, ok = x.OneType2.(*TypeOneOf1_FInt64); !ok {
						ot = new(TypeOneOf1_FInt64)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
						return err
					}
					ot.FInt64 = vv
					x.OneType2 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one3":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType3 = nil
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
				case "t_uint32":
					if oneOfIsFill_OneType3 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType3 = true

					var ok bool
					var ot *TypeOneOf1_FUint32
					if ot, ok = x.OneType3.(*TypeOneOf1_FUint32); !ok {
						ot = new(TypeOneOf1_FUint32)
					}
					var vv uint32
					if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
						return err
					}
					ot.FUint32 = vv
					x.OneType3 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one4":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType4 = nil
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
				case "t_uint64":
					if oneOfIsFill_OneType4 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType4 = true

					var ok bool
					var ot *TypeOneOf1_FUint64
					if ot, ok = x.OneType4.(*TypeOneOf1_FUint64); !ok {
						ot = new(TypeOneOf1_FUint64)
					}
					var vv uint64
					if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
						return err
					}
					ot.FUint64 = vv
					x.OneType4 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one5":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType5 = nil
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
				case "t_sint32":
					if oneOfIsFill_OneType5 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType5 = true

					var ok bool
					var ot *TypeOneOf1_FSint32
					if ot, ok = x.OneType5.(*TypeOneOf1_FSint32); !ok {
						ot = new(TypeOneOf1_FSint32)
					}
					var vv int32
					if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
						return err
					}
					ot.FSint32 = vv
					x.OneType5 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one6":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType6 = nil
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
				case "t_sint64":
					if oneOfIsFill_OneType6 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType6 = true

					var ok bool
					var ot *TypeOneOf1_FSint64
					if ot, ok = x.OneType6.(*TypeOneOf1_FSint64); !ok {
						ot = new(TypeOneOf1_FSint64)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
						return err
					}
					ot.FSint64 = vv
					x.OneType6 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one7":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType7 = nil
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
				case "t_sfixed32":
					if oneOfIsFill_OneType7 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType7 = true

					var ok bool
					var ot *TypeOneOf1_FSfixed32
					if ot, ok = x.OneType7.(*TypeOneOf1_FSfixed32); !ok {
						ot = new(TypeOneOf1_FSfixed32)
					}
					var vv int32
					if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
						return err
					}
					ot.FSfixed32 = vv
					x.OneType7 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one8":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType8 = nil
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
				case "t_sfixed64":
					if oneOfIsFill_OneType8 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType8 = true

					var ok bool
					var ot *TypeOneOf1_FSfixed64
					if ot, ok = x.OneType8.(*TypeOneOf1_FSfixed64); !ok {
						ot = new(TypeOneOf1_FSfixed64)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
						return err
					}
					ot.FSfixed64 = vv
					x.OneType8 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one9":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType9 = nil
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
				case "t_fixed32":
					if oneOfIsFill_OneType9 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType9 = true

					var ok bool
					var ot *TypeOneOf1_FFixed32
					if ot, ok = x.OneType9.(*TypeOneOf1_FFixed32); !ok {
						ot = new(TypeOneOf1_FFixed32)
					}
					var vv uint32
					if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
						return err
					}
					ot.FFixed32 = vv
					x.OneType9 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one10":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType10 = nil
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
				case "t_fixed64":
					if oneOfIsFill_OneType10 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType10 = true

					var ok bool
					var ot *TypeOneOf1_FFixed64
					if ot, ok = x.OneType10.(*TypeOneOf1_FFixed64); !ok {
						ot = new(TypeOneOf1_FFixed64)
					}
					var vv uint64
					if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
						return err
					}
					ot.FFixed64 = vv
					x.OneType10 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one11":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType11 = nil
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
				case "t_float":
					if oneOfIsFill_OneType11 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType11 = true

					var ok bool
					var ot *TypeOneOf1_FFloat
					if ot, ok = x.OneType11.(*TypeOneOf1_FFloat); !ok {
						ot = new(TypeOneOf1_FFloat)
					}
					var vv float32
					if vv, err = decoder.ReadLiteralFloat32(jsonKey, false); err != nil {
						return err
					}
					ot.FFloat = vv
					x.OneType11 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one12":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType12 = nil
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
				case "t_double":
					if oneOfIsFill_OneType12 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType12 = true

					var ok bool
					var ot *TypeOneOf1_FDouble
					if ot, ok = x.OneType12.(*TypeOneOf1_FDouble); !ok {
						ot = new(TypeOneOf1_FDouble)
					}
					var vv float64
					if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
						return err
					}
					ot.FDouble = vv
					x.OneType12 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one13":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType13 = nil
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
				case "t_bool1":
					if oneOfIsFill_OneType13 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType13 = true

					var ok bool
					var ot *TypeOneOf1_FBool1
					if ot, ok = x.OneType13.(*TypeOneOf1_FBool1); !ok {
						ot = new(TypeOneOf1_FBool1)
					}
					var vv bool
					if vv, err = decoder.ReadLiteralBool(jsonKey, false); err != nil {
						return err
					}
					ot.FBool1 = vv
					x.OneType13 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one14":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType14 = nil
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
				case "t_string1":
					if oneOfIsFill_OneType14 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType14 = true

					var ok bool
					var ot *TypeOneOf1_FString1
					if ot, ok = x.OneType14.(*TypeOneOf1_FString1); !ok {
						ot = new(TypeOneOf1_FString1)
					}
					var vv string
					if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
						return err
					}
					ot.FString1 = vv
					x.OneType14 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one15":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType15 = nil
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
				case "t_bytes1":
					if oneOfIsFill_OneType15 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType15 = true

					var ok bool
					var ot *TypeOneOf1_FBytes1
					if ot, ok = x.OneType15.(*TypeOneOf1_FBytes1); !ok {
						ot = new(TypeOneOf1_FBytes1)
					}
					var vv []byte
					if vv, err = decoder.ReadLiteralBytes(jsonKey); err != nil {
						return err
					}
					ot.FBytes1 = vv
					x.OneType15 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one16":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType16 = nil
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
				case "t_enum1":
					if oneOfIsFill_OneType16 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType16 = true

					var ok bool
					var ot *TypeOneOf1_FEnum1
					if ot, ok = x.OneType16.(*TypeOneOf1_FEnum1); !ok {
						ot = new(TypeOneOf1_FEnum1)
					}
					var vv pbexternal.Enum1
					var v1 int32
					if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
						return err
					}
					vv = pbexternal.Enum1(v1)
					ot.FEnum1 = vv
					x.OneType16 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one17":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType17 = nil
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
				case "t_message1":
					if oneOfIsFill_OneType17 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType17 = true

					var ok bool
					var ot *TypeOneOf1_FMessage1
					if ot, ok = x.OneType17.(*TypeOneOf1_FMessage1); !ok {
						ot = new(TypeOneOf1_FMessage1)
					}
					var vv *pbexternal.Message1
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FMessage1 != nil {
							vv = ot.FMessage1
						} else {
							vv = new(pbexternal.Message1)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FMessage1 = vv
					x.OneType17 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one18":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType18 = nil
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
				case "t_any1":
					if oneOfIsFill_OneType18 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType18 = true

					var ok bool
					var ot *TypeOneOf1_FAny1
					if ot, ok = x.OneType18.(*TypeOneOf1_FAny1); !ok {
						ot = new(TypeOneOf1_FAny1)
					}
					var vv *anypb.Any
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FAny1 != nil {
							vv = ot.FAny1
						} else {
							vv = new(anypb.Any)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FAny1 = vv
					x.OneType18 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one19":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType19 = nil
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
				case "t_duration1":
					if oneOfIsFill_OneType19 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType19 = true

					var ok bool
					var ot *TypeOneOf1_FDuration1
					if ot, ok = x.OneType19.(*TypeOneOf1_FDuration1); !ok {
						ot = new(TypeOneOf1_FDuration1)
					}
					var vv *durationpb.Duration
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FDuration1 != nil {
							vv = ot.FDuration1
						} else {
							vv = new(durationpb.Duration)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FDuration1 = vv
					x.OneType19 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "t_one20":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneType20 = nil
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
				case "t_timestamp1":
					if oneOfIsFill_OneType20 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneType20 = true

					var ok bool
					var ot *TypeOneOf1_FTimestamp1
					if ot, ok = x.OneType20.(*TypeOneOf1_FTimestamp1); !ok {
						ot = new(TypeOneOf1_FTimestamp1)
					}
					var vv *timestamppb.Timestamp
					if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
						return err
					}
					if !isNULL {
						if ot.FTimestamp1 != nil {
							vv = ot.FTimestamp1
						} else {
							vv = new(timestamppb.Timestamp)
						}
						if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
							return err
						}
					}
					ot.FTimestamp1 = vv
					x.OneType20 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
