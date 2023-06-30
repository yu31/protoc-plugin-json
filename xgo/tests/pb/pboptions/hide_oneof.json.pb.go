// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/options/hide_oneof.proto

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

// MarshalJSON implements interface json.Marshaler for proto message HideOneOf1 in file tests/proto/cases/options/hide_oneof.proto
func (x *HideOneOf1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(768)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	switch ov := x.OneType1.(type) {
	case *HideOneOf1_FInt32:
		encoder.AppendObjectKey("t_int32")
		encoder.AppendLiteralInt32(ov.FInt32, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType2.(type) {
	case *HideOneOf1_FInt64:
		encoder.AppendObjectKey("t_int64")
		encoder.AppendLiteralInt64(ov.FInt64, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType3.(type) {
	case *HideOneOf1_FUint32:
		encoder.AppendObjectKey("t_uint32")
		encoder.AppendLiteralUint32(ov.FUint32, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType4.(type) {
	case *HideOneOf1_FUint64:
		encoder.AppendObjectKey("t_uint64")
		encoder.AppendLiteralUint64(ov.FUint64, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType5.(type) {
	case *HideOneOf1_FSint32:
		encoder.AppendObjectKey("t_sint32")
		encoder.AppendLiteralInt32(ov.FSint32, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType6.(type) {
	case *HideOneOf1_FSint64:
		encoder.AppendObjectKey("t_sint64")
		encoder.AppendLiteralInt64(ov.FSint64, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType7.(type) {
	case *HideOneOf1_FSfixed32:
		encoder.AppendObjectKey("t_sfixed32")
		encoder.AppendLiteralInt32(ov.FSfixed32, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType8.(type) {
	case *HideOneOf1_FSfixed64:
		encoder.AppendObjectKey("t_sfixed64")
		encoder.AppendLiteralInt64(ov.FSfixed64, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType9.(type) {
	case *HideOneOf1_FFixed32:
		encoder.AppendObjectKey("t_fixed32")
		encoder.AppendLiteralUint32(ov.FFixed32, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType10.(type) {
	case *HideOneOf1_FFixed64:
		encoder.AppendObjectKey("t_fixed64")
		encoder.AppendLiteralUint64(ov.FFixed64, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType11.(type) {
	case *HideOneOf1_FFloat:
		encoder.AppendObjectKey("t_float")
		encoder.AppendLiteralFloat32(ov.FFloat, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType12.(type) {
	case *HideOneOf1_FDouble:
		encoder.AppendObjectKey("t_double")
		encoder.AppendLiteralFloat64(ov.FDouble, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType13.(type) {
	case *HideOneOf1_FBool1:
		encoder.AppendObjectKey("t_bool1")
		encoder.AppendLiteralBool(ov.FBool1, false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType14.(type) {
	case *HideOneOf1_FString1:
		encoder.AppendObjectKey("t_string1")
		encoder.AppendLiteralString(ov.FString1)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType15.(type) {
	case *HideOneOf1_FBytes1:
		encoder.AppendObjectKey("t_bytes1")
		encoder.AppendLiteralBytes(ov.FBytes1)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType16.(type) {
	case *HideOneOf1_FEnum1:
		encoder.AppendObjectKey("t_enum1")
		encoder.AppendLiteralInt32(int32(ov.FEnum1.Number()), false)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType17.(type) {
	case *HideOneOf1_FMessage1:
		encoder.AppendObjectKey("t_message1")
		if err = encoder.AppendLiteralInterface(ov.FMessage1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType18.(type) {
	case *HideOneOf1_FAny1:
		encoder.AppendObjectKey("t_any1")
		if err = encoder.AppendLiteralInterface(ov.FAny1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType19.(type) {
	case *HideOneOf1_FDuration1:
		encoder.AppendObjectKey("t_duration1")
		if err = encoder.AppendLiteralInterface(ov.FDuration1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneType20.(type) {
	case *HideOneOf1_FTimestamp1:
		encoder.AppendObjectKey("t_timestamp1")
		if err = encoder.AppendLiteralInterface(ov.FTimestamp1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message HideOneOf1 in file tests/proto/cases/options/hide_oneof.proto
func (x *HideOneOf1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*HideOneOf1) is nil")
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
		case "t_int32":
			if oneOfIsFill_OneType1 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType1 = true

			var ok bool
			var ot *HideOneOf1_FInt32
			if ot, ok = x.OneType1.(*HideOneOf1_FInt32); !ok {
				ot = new(HideOneOf1_FInt32)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			ot.FInt32 = vv
			x.OneType1 = ot
		case "t_int64":
			if oneOfIsFill_OneType2 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType2 = true

			var ok bool
			var ot *HideOneOf1_FInt64
			if ot, ok = x.OneType2.(*HideOneOf1_FInt64); !ok {
				ot = new(HideOneOf1_FInt64)
			}
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			ot.FInt64 = vv
			x.OneType2 = ot
		case "t_uint32":
			if oneOfIsFill_OneType3 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType3 = true

			var ok bool
			var ot *HideOneOf1_FUint32
			if ot, ok = x.OneType3.(*HideOneOf1_FUint32); !ok {
				ot = new(HideOneOf1_FUint32)
			}
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			ot.FUint32 = vv
			x.OneType3 = ot
		case "t_uint64":
			if oneOfIsFill_OneType4 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType4 = true

			var ok bool
			var ot *HideOneOf1_FUint64
			if ot, ok = x.OneType4.(*HideOneOf1_FUint64); !ok {
				ot = new(HideOneOf1_FUint64)
			}
			var vv uint64
			if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
				return err
			}
			ot.FUint64 = vv
			x.OneType4 = ot
		case "t_sint32":
			if oneOfIsFill_OneType5 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType5 = true

			var ok bool
			var ot *HideOneOf1_FSint32
			if ot, ok = x.OneType5.(*HideOneOf1_FSint32); !ok {
				ot = new(HideOneOf1_FSint32)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			ot.FSint32 = vv
			x.OneType5 = ot
		case "t_sint64":
			if oneOfIsFill_OneType6 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType6 = true

			var ok bool
			var ot *HideOneOf1_FSint64
			if ot, ok = x.OneType6.(*HideOneOf1_FSint64); !ok {
				ot = new(HideOneOf1_FSint64)
			}
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			ot.FSint64 = vv
			x.OneType6 = ot
		case "t_sfixed32":
			if oneOfIsFill_OneType7 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType7 = true

			var ok bool
			var ot *HideOneOf1_FSfixed32
			if ot, ok = x.OneType7.(*HideOneOf1_FSfixed32); !ok {
				ot = new(HideOneOf1_FSfixed32)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			ot.FSfixed32 = vv
			x.OneType7 = ot
		case "t_sfixed64":
			if oneOfIsFill_OneType8 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType8 = true

			var ok bool
			var ot *HideOneOf1_FSfixed64
			if ot, ok = x.OneType8.(*HideOneOf1_FSfixed64); !ok {
				ot = new(HideOneOf1_FSfixed64)
			}
			var vv int64
			if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
				return err
			}
			ot.FSfixed64 = vv
			x.OneType8 = ot
		case "t_fixed32":
			if oneOfIsFill_OneType9 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType9 = true

			var ok bool
			var ot *HideOneOf1_FFixed32
			if ot, ok = x.OneType9.(*HideOneOf1_FFixed32); !ok {
				ot = new(HideOneOf1_FFixed32)
			}
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			ot.FFixed32 = vv
			x.OneType9 = ot
		case "t_fixed64":
			if oneOfIsFill_OneType10 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType10 = true

			var ok bool
			var ot *HideOneOf1_FFixed64
			if ot, ok = x.OneType10.(*HideOneOf1_FFixed64); !ok {
				ot = new(HideOneOf1_FFixed64)
			}
			var vv uint64
			if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
				return err
			}
			ot.FFixed64 = vv
			x.OneType10 = ot
		case "t_float":
			if oneOfIsFill_OneType11 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType11 = true

			var ok bool
			var ot *HideOneOf1_FFloat
			if ot, ok = x.OneType11.(*HideOneOf1_FFloat); !ok {
				ot = new(HideOneOf1_FFloat)
			}
			var vv float32
			if vv, err = decoder.ReadLiteralFloat32(jsonKey, false); err != nil {
				return err
			}
			ot.FFloat = vv
			x.OneType11 = ot
		case "t_double":
			if oneOfIsFill_OneType12 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType12 = true

			var ok bool
			var ot *HideOneOf1_FDouble
			if ot, ok = x.OneType12.(*HideOneOf1_FDouble); !ok {
				ot = new(HideOneOf1_FDouble)
			}
			var vv float64
			if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
				return err
			}
			ot.FDouble = vv
			x.OneType12 = ot
		case "t_bool1":
			if oneOfIsFill_OneType13 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType13 = true

			var ok bool
			var ot *HideOneOf1_FBool1
			if ot, ok = x.OneType13.(*HideOneOf1_FBool1); !ok {
				ot = new(HideOneOf1_FBool1)
			}
			var vv bool
			if vv, err = decoder.ReadLiteralBool(jsonKey, false); err != nil {
				return err
			}
			ot.FBool1 = vv
			x.OneType13 = ot
		case "t_string1":
			if oneOfIsFill_OneType14 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType14 = true

			var ok bool
			var ot *HideOneOf1_FString1
			if ot, ok = x.OneType14.(*HideOneOf1_FString1); !ok {
				ot = new(HideOneOf1_FString1)
			}
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			ot.FString1 = vv
			x.OneType14 = ot
		case "t_bytes1":
			if oneOfIsFill_OneType15 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType15 = true

			var ok bool
			var ot *HideOneOf1_FBytes1
			if ot, ok = x.OneType15.(*HideOneOf1_FBytes1); !ok {
				ot = new(HideOneOf1_FBytes1)
			}
			var vv []byte
			if vv, err = decoder.ReadLiteralBytes(jsonKey); err != nil {
				return err
			}
			ot.FBytes1 = vv
			x.OneType15 = ot
		case "t_enum1":
			if oneOfIsFill_OneType16 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType16 = true

			var ok bool
			var ot *HideOneOf1_FEnum1
			if ot, ok = x.OneType16.(*HideOneOf1_FEnum1); !ok {
				ot = new(HideOneOf1_FEnum1)
			}
			var vv pbexternal.Enum1
			var v1 int32
			if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
				return err
			}
			vv = pbexternal.Enum1(v1)
			ot.FEnum1 = vv
			x.OneType16 = ot
		case "t_message1":
			if oneOfIsFill_OneType17 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType17 = true

			var ok bool
			var ot *HideOneOf1_FMessage1
			if ot, ok = x.OneType17.(*HideOneOf1_FMessage1); !ok {
				ot = new(HideOneOf1_FMessage1)
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
		case "t_any1":
			if oneOfIsFill_OneType18 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType18 = true

			var ok bool
			var ot *HideOneOf1_FAny1
			if ot, ok = x.OneType18.(*HideOneOf1_FAny1); !ok {
				ot = new(HideOneOf1_FAny1)
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
		case "t_duration1":
			if oneOfIsFill_OneType19 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType19 = true

			var ok bool
			var ot *HideOneOf1_FDuration1
			if ot, ok = x.OneType19.(*HideOneOf1_FDuration1); !ok {
				ot = new(HideOneOf1_FDuration1)
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
		case "t_timestamp1":
			if oneOfIsFill_OneType20 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType20 = true

			var ok bool
			var ot *HideOneOf1_FTimestamp1
			if ot, ok = x.OneType20.(*HideOneOf1_FTimestamp1); !ok {
				ot = new(HideOneOf1_FTimestamp1)
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
	return nil
}
