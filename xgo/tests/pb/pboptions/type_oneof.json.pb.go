// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/options/type_oneof.proto

package pboptions

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	_ "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeOneOf1 in file tests/proto/cases/options/type_oneof.proto
func (x *TypeOneOf1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(960)
	enc.AppendObjectBegin() // Add begin JSON identifier

	if x.OneType1 != nil {
		enc.AppendObjectKey("t_one1")
		enc.AppendObjectBegin()
		switch o1_1 := x.OneType1.(type) {
		case *TypeOneOf1_FInt32:
			jsonencoder.AppendValI32(enc, "t_int32", o1_1.FInt32, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType2 != nil {
		enc.AppendObjectKey("t_one2")
		enc.AppendObjectBegin()
		switch o1_3 := x.OneType2.(type) {
		case *TypeOneOf1_FInt64:
			jsonencoder.AppendValI64(enc, "t_int64", o1_3.FInt64, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType3 != nil {
		enc.AppendObjectKey("t_one3")
		enc.AppendObjectBegin()
		switch o1_5 := x.OneType3.(type) {
		case *TypeOneOf1_FUint32:
			jsonencoder.AppendValU32(enc, "t_uint32", o1_5.FUint32, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType4 != nil {
		enc.AppendObjectKey("t_one4")
		enc.AppendObjectBegin()
		switch o1_7 := x.OneType4.(type) {
		case *TypeOneOf1_FUint64:
			jsonencoder.AppendValU64(enc, "t_uint64", o1_7.FUint64, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType5 != nil {
		enc.AppendObjectKey("t_one5")
		enc.AppendObjectBegin()
		switch o1_9 := x.OneType5.(type) {
		case *TypeOneOf1_FSint32:
			jsonencoder.AppendValI32(enc, "t_sint32", o1_9.FSint32, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType6 != nil {
		enc.AppendObjectKey("t_one6")
		enc.AppendObjectBegin()
		switch o1_11 := x.OneType6.(type) {
		case *TypeOneOf1_FSint64:
			jsonencoder.AppendValI64(enc, "t_sint64", o1_11.FSint64, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType7 != nil {
		enc.AppendObjectKey("t_one7")
		enc.AppendObjectBegin()
		switch o1_13 := x.OneType7.(type) {
		case *TypeOneOf1_FSfixed32:
			jsonencoder.AppendValI32(enc, "t_sfixed32", o1_13.FSfixed32, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType8 != nil {
		enc.AppendObjectKey("t_one8")
		enc.AppendObjectBegin()
		switch o1_15 := x.OneType8.(type) {
		case *TypeOneOf1_FSfixed64:
			jsonencoder.AppendValI64(enc, "t_sfixed64", o1_15.FSfixed64, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType9 != nil {
		enc.AppendObjectKey("t_one9")
		enc.AppendObjectBegin()
		switch o1_17 := x.OneType9.(type) {
		case *TypeOneOf1_FFixed32:
			jsonencoder.AppendValU32(enc, "t_fixed32", o1_17.FFixed32, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType10 != nil {
		enc.AppendObjectKey("t_one10")
		enc.AppendObjectBegin()
		switch o1_19 := x.OneType10.(type) {
		case *TypeOneOf1_FFixed64:
			jsonencoder.AppendValU64(enc, "t_fixed64", o1_19.FFixed64, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType11 != nil {
		enc.AppendObjectKey("t_one11")
		enc.AppendObjectBegin()
		switch o1_21 := x.OneType11.(type) {
		case *TypeOneOf1_FFloat:
			jsonencoder.AppendValF32(enc, "t_float", o1_21.FFloat, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType12 != nil {
		enc.AppendObjectKey("t_one12")
		enc.AppendObjectBegin()
		switch o1_23 := x.OneType12.(type) {
		case *TypeOneOf1_FDouble:
			jsonencoder.AppendValF64(enc, "t_double", o1_23.FDouble, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType13 != nil {
		enc.AppendObjectKey("t_one13")
		enc.AppendObjectBegin()
		switch o1_25 := x.OneType13.(type) {
		case *TypeOneOf1_FBool1:
			jsonencoder.AppendValBool(enc, "t_bool1", o1_25.FBool1, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType14 != nil {
		enc.AppendObjectKey("t_one14")
		enc.AppendObjectBegin()
		switch o1_27 := x.OneType14.(type) {
		case *TypeOneOf1_FString1:
			jsonencoder.AppendValStr(enc, "t_string1", o1_27.FString1, true)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType15 != nil {
		enc.AppendObjectKey("t_one15")
		enc.AppendObjectBegin()
		switch o1_29 := x.OneType15.(type) {
		case *TypeOneOf1_FBytes1:
			if err := jsonencoder.AppendValBytes(enc, "t_bytes1", o1_29.FBytes1, true); err != nil {
				return nil, err
			}
		}
		enc.AppendObjectEnd()
	}
	if x.OneType16 != nil {
		enc.AppendObjectKey("t_one16")
		enc.AppendObjectBegin()
		switch o1_31 := x.OneType16.(type) {
		case *TypeOneOf1_FEnum1:
			jsonencoder.AppendValEnumNum(enc, "t_enum1", o1_31.FEnum1, true, false)
		}
		enc.AppendObjectEnd()
	}
	if x.OneType17 != nil {
		enc.AppendObjectKey("t_one17")
		enc.AppendObjectBegin()
		switch o1_33 := x.OneType17.(type) {
		case *TypeOneOf1_FMessage1:
			if err := jsonencoder.AppendValMessage(enc, "t_message1", o1_33.FMessage1, true); err != nil {
				return nil, err
			}
		}
		enc.AppendObjectEnd()
	}
	if x.OneType18 != nil {
		enc.AppendObjectKey("t_one18")
		enc.AppendObjectBegin()
		switch o1_35 := x.OneType18.(type) {
		case *TypeOneOf1_FAny1:
			if err := jsonencoder.AppendValWKTAnyObject(enc, "t_any1", o1_35.FAny1, true); err != nil {
				return nil, err
			}
		}
		enc.AppendObjectEnd()
	}
	if x.OneType19 != nil {
		enc.AppendObjectKey("t_one19")
		enc.AppendObjectBegin()
		switch o1_37 := x.OneType19.(type) {
		case *TypeOneOf1_FDuration1:
			if err := jsonencoder.AppendValWKTDurObject(enc, "t_duration1", o1_37.FDuration1, true); err != nil {
				return nil, err
			}
		}
		enc.AppendObjectEnd()
	}
	if x.OneType20 != nil {
		enc.AppendObjectKey("t_one20")
		enc.AppendObjectBegin()
		switch o1_39 := x.OneType20.(type) {
		case *TypeOneOf1_FTimestamp1:
			if err := jsonencoder.AppendValWKTTsObject(enc, "t_timestamp1", o1_39.FTimestamp1, true); err != nil {
				return nil, err
			}
		}
		enc.AppendObjectEnd()
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOneOf1 in file tests/proto/cases/options/type_oneof.proto
func (x *TypeOneOf1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*TypeOneOf1) is nil")
	}
	var (
		err    error
		isNULL bool
		dec    *jsondecoder.Decoder
	)
	if dec, err = jsondecoder.New(b); err != nil {
		return err
	}
	if isNULL, err = dec.BeforeScanJSON(); err != nil || isNULL {
		return err
	}
	// declares variables to report whether the oneof field is loaded.
	var (
		isLoad_o1_1  bool
		isLoad_o1_3  bool
		isLoad_o1_5  bool
		isLoad_o1_7  bool
		isLoad_o1_9  bool
		isLoad_o1_11 bool
		isLoad_o1_13 bool
		isLoad_o1_15 bool
		isLoad_o1_17 bool
		isLoad_o1_19 bool
		isLoad_o1_21 bool
		isLoad_o1_23 bool
		isLoad_o1_25 bool
		isLoad_o1_27 bool
		isLoad_o1_29 bool
		isLoad_o1_31 bool
		isLoad_o1_33 bool
		isLoad_o1_35 bool
		isLoad_o1_37 bool
		isLoad_o1_39 bool
	)

LOOP_SCAN:
	for { // Loop to read the JSON objects
		var (
			jsonKey string
			isEnd   bool
		)

		if isEnd, err = dec.BeforeScanNext(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}

		if jsonKey, err = dec.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the jsonKey
		case "t_one1":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType1 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_1 string
					if oneOfKey1_1, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_1 { // match the oneof key
					case "t_int32":
						if isLoad_o1_1 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_1 = true
						o1_1, ok := x.OneType1.(*TypeOneOf1_FInt32)
						if !ok {
							o1_1 = new(TypeOneOf1_FInt32)
							x.OneType1 = o1_1
						}
						if o1_1.FInt32, err = jsondecoder.ReadValI32(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one2":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType2 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_3 string
					if oneOfKey1_3, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_3 { // match the oneof key
					case "t_int64":
						if isLoad_o1_3 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_3 = true
						o1_3, ok := x.OneType2.(*TypeOneOf1_FInt64)
						if !ok {
							o1_3 = new(TypeOneOf1_FInt64)
							x.OneType2 = o1_3
						}
						if o1_3.FInt64, err = jsondecoder.ReadValI64(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one3":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType3 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_5 string
					if oneOfKey1_5, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_5 { // match the oneof key
					case "t_uint32":
						if isLoad_o1_5 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_5 = true
						o1_5, ok := x.OneType3.(*TypeOneOf1_FUint32)
						if !ok {
							o1_5 = new(TypeOneOf1_FUint32)
							x.OneType3 = o1_5
						}
						if o1_5.FUint32, err = jsondecoder.ReadValU32(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one4":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType4 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_7 string
					if oneOfKey1_7, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_7 { // match the oneof key
					case "t_uint64":
						if isLoad_o1_7 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_7 = true
						o1_7, ok := x.OneType4.(*TypeOneOf1_FUint64)
						if !ok {
							o1_7 = new(TypeOneOf1_FUint64)
							x.OneType4 = o1_7
						}
						if o1_7.FUint64, err = jsondecoder.ReadValU64(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one5":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType5 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_9 string
					if oneOfKey1_9, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_9 { // match the oneof key
					case "t_sint32":
						if isLoad_o1_9 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_9 = true
						o1_9, ok := x.OneType5.(*TypeOneOf1_FSint32)
						if !ok {
							o1_9 = new(TypeOneOf1_FSint32)
							x.OneType5 = o1_9
						}
						if o1_9.FSint32, err = jsondecoder.ReadValI32(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one6":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType6 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_11 string
					if oneOfKey1_11, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_11 { // match the oneof key
					case "t_sint64":
						if isLoad_o1_11 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_11 = true
						o1_11, ok := x.OneType6.(*TypeOneOf1_FSint64)
						if !ok {
							o1_11 = new(TypeOneOf1_FSint64)
							x.OneType6 = o1_11
						}
						if o1_11.FSint64, err = jsondecoder.ReadValI64(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one7":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType7 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_13 string
					if oneOfKey1_13, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_13 { // match the oneof key
					case "t_sfixed32":
						if isLoad_o1_13 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_13 = true
						o1_13, ok := x.OneType7.(*TypeOneOf1_FSfixed32)
						if !ok {
							o1_13 = new(TypeOneOf1_FSfixed32)
							x.OneType7 = o1_13
						}
						if o1_13.FSfixed32, err = jsondecoder.ReadValI32(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one8":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType8 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_15 string
					if oneOfKey1_15, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_15 { // match the oneof key
					case "t_sfixed64":
						if isLoad_o1_15 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_15 = true
						o1_15, ok := x.OneType8.(*TypeOneOf1_FSfixed64)
						if !ok {
							o1_15 = new(TypeOneOf1_FSfixed64)
							x.OneType8 = o1_15
						}
						if o1_15.FSfixed64, err = jsondecoder.ReadValI64(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one9":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType9 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_17 string
					if oneOfKey1_17, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_17 { // match the oneof key
					case "t_fixed32":
						if isLoad_o1_17 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_17 = true
						o1_17, ok := x.OneType9.(*TypeOneOf1_FFixed32)
						if !ok {
							o1_17 = new(TypeOneOf1_FFixed32)
							x.OneType9 = o1_17
						}
						if o1_17.FFixed32, err = jsondecoder.ReadValU32(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one10":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType10 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_19 string
					if oneOfKey1_19, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_19 { // match the oneof key
					case "t_fixed64":
						if isLoad_o1_19 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_19 = true
						o1_19, ok := x.OneType10.(*TypeOneOf1_FFixed64)
						if !ok {
							o1_19 = new(TypeOneOf1_FFixed64)
							x.OneType10 = o1_19
						}
						if o1_19.FFixed64, err = jsondecoder.ReadValU64(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one11":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType11 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_21 string
					if oneOfKey1_21, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_21 { // match the oneof key
					case "t_float":
						if isLoad_o1_21 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_21 = true
						o1_21, ok := x.OneType11.(*TypeOneOf1_FFloat)
						if !ok {
							o1_21 = new(TypeOneOf1_FFloat)
							x.OneType11 = o1_21
						}
						if o1_21.FFloat, err = jsondecoder.ReadValF32(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one12":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType12 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_23 string
					if oneOfKey1_23, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_23 { // match the oneof key
					case "t_double":
						if isLoad_o1_23 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_23 = true
						o1_23, ok := x.OneType12.(*TypeOneOf1_FDouble)
						if !ok {
							o1_23 = new(TypeOneOf1_FDouble)
							x.OneType12 = o1_23
						}
						if o1_23.FDouble, err = jsondecoder.ReadValF64(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one13":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType13 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_25 string
					if oneOfKey1_25, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_25 { // match the oneof key
					case "t_bool1":
						if isLoad_o1_25 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_25 = true
						o1_25, ok := x.OneType13.(*TypeOneOf1_FBool1)
						if !ok {
							o1_25 = new(TypeOneOf1_FBool1)
							x.OneType13 = o1_25
						}
						if o1_25.FBool1, err = jsondecoder.ReadValBool(dec, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one14":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType14 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_27 string
					if oneOfKey1_27, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_27 { // match the oneof key
					case "t_string1":
						if isLoad_o1_27 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_27 = true
						o1_27, ok := x.OneType14.(*TypeOneOf1_FString1)
						if !ok {
							o1_27 = new(TypeOneOf1_FString1)
							x.OneType14 = o1_27
						}
						if o1_27.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one15":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType15 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_29 string
					if oneOfKey1_29, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_29 { // match the oneof key
					case "t_bytes1":
						if isLoad_o1_29 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_29 = true
						o1_29, ok := x.OneType15.(*TypeOneOf1_FBytes1)
						if !ok {
							o1_29 = new(TypeOneOf1_FBytes1)
							x.OneType15 = o1_29
						}
						if o1_29.FBytes1, err = jsondecoder.ReadValBytes(dec); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one16":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType16 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_31 string
					if oneOfKey1_31, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_31 { // match the oneof key
					case "t_enum1":
						if isLoad_o1_31 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_31 = true
						o1_31, ok := x.OneType16.(*TypeOneOf1_FEnum1)
						if !ok {
							o1_31 = new(TypeOneOf1_FEnum1)
							x.OneType16 = o1_31
						}
						if o1_31.FEnum1, err = jsondecoder.ReadValEnumNum(dec, o1_31.FEnum1, false); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one17":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType17 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_33 string
					if oneOfKey1_33, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_33 { // match the oneof key
					case "t_message1":
						if isLoad_o1_33 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_33 = true
						o1_33, ok := x.OneType17.(*TypeOneOf1_FMessage1)
						if !ok {
							o1_33 = new(TypeOneOf1_FMessage1)
							x.OneType17 = o1_33
						}
						if o1_33.FMessage1, err = jsondecoder.ReadValMessage(dec, o1_33.FMessage1); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one18":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType18 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_35 string
					if oneOfKey1_35, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_35 { // match the oneof key
					case "t_any1":
						if isLoad_o1_35 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_35 = true
						o1_35, ok := x.OneType18.(*TypeOneOf1_FAny1)
						if !ok {
							o1_35 = new(TypeOneOf1_FAny1)
							x.OneType18 = o1_35
						}
						if o1_35.FAny1, err = jsondecoder.ReadValWKTAnyObject(dec, o1_35.FAny1); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one19":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType19 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_37 string
					if oneOfKey1_37, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_37 { // match the oneof key
					case "t_duration1":
						if isLoad_o1_37 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_37 = true
						o1_37, ok := x.OneType19.(*TypeOneOf1_FDuration1)
						if !ok {
							o1_37 = new(TypeOneOf1_FDuration1)
							x.OneType19 = o1_37
						}
						if o1_37.FDuration1, err = jsondecoder.ReadValWKTDurObject(dec, o1_37.FDuration1); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "t_one20":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType20 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_39 string
					if oneOfKey1_39, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_39 { // match the oneof key
					case "t_timestamp1":
						if isLoad_o1_39 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_39 = true
						o1_39, ok := x.OneType20.(*TypeOneOf1_FTimestamp1)
						if !ok {
							o1_39 = new(TypeOneOf1_FTimestamp1)
							x.OneType20 = o1_39
						}
						if o1_39.FTimestamp1, err = jsondecoder.ReadValWKTTsObject(dec, o1_39.FTimestamp1); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		default:
			if err = dec.DiscardValue(); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
