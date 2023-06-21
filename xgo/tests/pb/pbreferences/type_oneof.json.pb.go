// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/references/type_oneof.proto

package pbreferences

import (
	errors "errors"
	fmt "fmt"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeOneOf1 in file tests/proto/cases/references/type_oneof.proto
func (x *TypeOneOf1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(896)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	switch ov := x.OneInt32.(type) {
	case *TypeOneOf1_FInt32A:
		encoder.AppendObjectKey("f_int32a")
		encoder.AppendLiteralInt32(ov.FInt32A, false)
	case *TypeOneOf1_FInt32B:
		encoder.AppendObjectKey("f_int32b")
		encoder.AppendLiteralInt32(ov.FInt32B, true)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneInt64.(type) {
	case *TypeOneOf1_FInt64A:
		encoder.AppendObjectKey("one_int64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_int64a")
		encoder.AppendLiteralInt64(ov.FInt64A, false)
		encoder.AppendObjectEnd()
	case *TypeOneOf1_FInt64B:
		encoder.AppendObjectKey("one_int64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_int64b")
		encoder.AppendLiteralInt64(ov.FInt64B, true)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_int64")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneUint32.(type) {
	case *TypeOneOf1_FUint32A:
		encoder.AppendObjectKey("f_uint32a")
		encoder.AppendLiteralUint32(ov.FUint32A, false)
	case *TypeOneOf1_FUint32B:
		encoder.AppendObjectKey("f_uint32b")
		encoder.AppendLiteralUint32(ov.FUint32B, true)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneUint64.(type) {
	case *TypeOneOf1_FUint64A:
		encoder.AppendObjectKey("one_uint64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_uint64a")
		encoder.AppendLiteralUint64(ov.FUint64A, false)
		encoder.AppendObjectEnd()
	case *TypeOneOf1_FUint64B:
		encoder.AppendObjectKey("one_uint64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_uint64b")
		encoder.AppendLiteralUint64(ov.FUint64B, true)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_uint64")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneSInt32.(type) {
	case *TypeOneOf1_FSint32A:
		encoder.AppendObjectKey("f_sint32a")
		encoder.AppendLiteralInt32(ov.FSint32A, false)
	case *TypeOneOf1_FSint32B:
		encoder.AppendObjectKey("f_sint32b")
		encoder.AppendLiteralInt32(ov.FSint32B, true)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneSInt64.(type) {
	case *TypeOneOf1_FSint64A:
		encoder.AppendObjectKey("one_sint64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_sint64a")
		encoder.AppendLiteralInt64(ov.FSint64A, false)
		encoder.AppendObjectEnd()
	case *TypeOneOf1_FSint64B:
		encoder.AppendObjectKey("one_sint64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_sint64b")
		encoder.AppendLiteralInt64(ov.FSint64B, true)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_sint64")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneSFixed32.(type) {
	case *TypeOneOf1_FSfixed32A:
		encoder.AppendObjectKey("f_sfixed32a")
		encoder.AppendLiteralInt32(ov.FSfixed32A, false)
	case *TypeOneOf1_FSfixed32B:
		encoder.AppendObjectKey("f_sfixed32b")
		encoder.AppendLiteralInt32(ov.FSfixed32B, true)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneSFixed64.(type) {
	case *TypeOneOf1_FSfixed64A:
		encoder.AppendObjectKey("one_sfixed64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_sfixed64a")
		encoder.AppendLiteralInt64(ov.FSfixed64A, false)
		encoder.AppendObjectEnd()
	case *TypeOneOf1_FSfixed64B:
		encoder.AppendObjectKey("one_sfixed64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_sfixed64b")
		encoder.AppendLiteralInt64(ov.FSfixed64B, true)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_sfixed64")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneFixed32.(type) {
	case *TypeOneOf1_FFixed32A:
		encoder.AppendObjectKey("f_fixed32a")
		encoder.AppendLiteralUint32(ov.FFixed32A, false)
	case *TypeOneOf1_FFixed32B:
		encoder.AppendObjectKey("f_fixed32b")
		encoder.AppendLiteralUint32(ov.FFixed32B, true)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneFixed64.(type) {
	case *TypeOneOf1_FFixed64A:
		encoder.AppendObjectKey("one_fixed64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_fixed64a")
		encoder.AppendLiteralUint64(ov.FFixed64A, false)
		encoder.AppendObjectEnd()
	case *TypeOneOf1_FFixed64B:
		encoder.AppendObjectKey("one_fixed64")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_fixed64b")
		encoder.AppendLiteralUint64(ov.FFixed64B, true)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_fixed64")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneFloat.(type) {
	case *TypeOneOf1_FFloat1:
		encoder.AppendObjectKey("f_float1")
		encoder.AppendLiteralFloat32(ov.FFloat1, false)
	case *TypeOneOf1_FFloat2:
		encoder.AppendObjectKey("f_float2")
		encoder.AppendLiteralFloat32(ov.FFloat2, true)
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneDouble.(type) {
	case *TypeOneOf1_FDouble1:
		encoder.AppendObjectKey("one_double")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_double1")
		encoder.AppendLiteralFloat64(ov.FDouble1, false)
		encoder.AppendObjectEnd()
	case *TypeOneOf1_FDouble2:
		encoder.AppendObjectKey("one_double")
		encoder.AppendObjectBegin()
		encoder.AppendObjectKey("f_double2")
		encoder.AppendLiteralFloat64(ov.FDouble2, true)
		encoder.AppendObjectEnd()
	case nil:
		encoder.AppendObjectKey("one_double")
		encoder.AppendLiteralNULL()
	default:
		_ = ov // to avoids unused panics
	}
	switch ov := x.OneBool.(type) {
	case *TypeOneOf1_FBool1:
		encoder.AppendObjectKey("f_bool1")
		encoder.AppendLiteralBool(ov.FBool1, false)
	case *TypeOneOf1_FBool2:
		encoder.AppendObjectKey("f_bool2")
		encoder.AppendLiteralBool(ov.FBool2, true)
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOneOf1 in file tests/proto/cases/references/type_oneof.proto
func (x *TypeOneOf1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbreferences.(*TypeOneOf1) is nil")
	}
	var (
		oneOfIsFill_OneInt32    bool
		oneOfIsFill_OneInt64    bool
		oneOfIsFill_OneUint32   bool
		oneOfIsFill_OneUint64   bool
		oneOfIsFill_OneSInt32   bool
		oneOfIsFill_OneSInt64   bool
		oneOfIsFill_OneSFixed32 bool
		oneOfIsFill_OneSFixed64 bool
		oneOfIsFill_OneFixed32  bool
		oneOfIsFill_OneFixed64  bool
		oneOfIsFill_OneFloat    bool
		oneOfIsFill_OneDouble   bool
		oneOfIsFill_OneBool     bool
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
		case "f_int32a":
			if oneOfIsFill_OneInt32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneInt32 = true

			var ok bool
			var ot *TypeOneOf1_FInt32A
			if ot, ok = x.OneInt32.(*TypeOneOf1_FInt32A); !ok {
				ot = new(TypeOneOf1_FInt32A)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			ot.FInt32A = vv
			x.OneInt32 = ot
		case "f_int32b":
			if oneOfIsFill_OneInt32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneInt32 = true

			var ok bool
			var ot *TypeOneOf1_FInt32B
			if ot, ok = x.OneInt32.(*TypeOneOf1_FInt32B); !ok {
				ot = new(TypeOneOf1_FInt32B)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, true); err != nil {
				return err
			}
			ot.FInt32B = vv
			x.OneInt32 = ot
		case "one_int64":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneInt64 = nil
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
				case "f_int64a":
					if oneOfIsFill_OneInt64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneInt64 = true

					var ok bool
					var ot *TypeOneOf1_FInt64A
					if ot, ok = x.OneInt64.(*TypeOneOf1_FInt64A); !ok {
						ot = new(TypeOneOf1_FInt64A)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
						return err
					}
					ot.FInt64A = vv
					x.OneInt64 = ot
				case "f_int64b":
					if oneOfIsFill_OneInt64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneInt64 = true

					var ok bool
					var ot *TypeOneOf1_FInt64B
					if ot, ok = x.OneInt64.(*TypeOneOf1_FInt64B); !ok {
						ot = new(TypeOneOf1_FInt64B)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, true); err != nil {
						return err
					}
					ot.FInt64B = vv
					x.OneInt64 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_uint32a":
			if oneOfIsFill_OneUint32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneUint32 = true

			var ok bool
			var ot *TypeOneOf1_FUint32A
			if ot, ok = x.OneUint32.(*TypeOneOf1_FUint32A); !ok {
				ot = new(TypeOneOf1_FUint32A)
			}
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			ot.FUint32A = vv
			x.OneUint32 = ot
		case "f_uint32b":
			if oneOfIsFill_OneUint32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneUint32 = true

			var ok bool
			var ot *TypeOneOf1_FUint32B
			if ot, ok = x.OneUint32.(*TypeOneOf1_FUint32B); !ok {
				ot = new(TypeOneOf1_FUint32B)
			}
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, true); err != nil {
				return err
			}
			ot.FUint32B = vv
			x.OneUint32 = ot
		case "one_uint64":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneUint64 = nil
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
				case "f_uint64a":
					if oneOfIsFill_OneUint64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneUint64 = true

					var ok bool
					var ot *TypeOneOf1_FUint64A
					if ot, ok = x.OneUint64.(*TypeOneOf1_FUint64A); !ok {
						ot = new(TypeOneOf1_FUint64A)
					}
					var vv uint64
					if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
						return err
					}
					ot.FUint64A = vv
					x.OneUint64 = ot
				case "f_uint64b":
					if oneOfIsFill_OneUint64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneUint64 = true

					var ok bool
					var ot *TypeOneOf1_FUint64B
					if ot, ok = x.OneUint64.(*TypeOneOf1_FUint64B); !ok {
						ot = new(TypeOneOf1_FUint64B)
					}
					var vv uint64
					if vv, err = decoder.ReadLiteralUint64(jsonKey, true); err != nil {
						return err
					}
					ot.FUint64B = vv
					x.OneUint64 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_sint32a":
			if oneOfIsFill_OneSInt32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneSInt32 = true

			var ok bool
			var ot *TypeOneOf1_FSint32A
			if ot, ok = x.OneSInt32.(*TypeOneOf1_FSint32A); !ok {
				ot = new(TypeOneOf1_FSint32A)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			ot.FSint32A = vv
			x.OneSInt32 = ot
		case "f_sint32b":
			if oneOfIsFill_OneSInt32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneSInt32 = true

			var ok bool
			var ot *TypeOneOf1_FSint32B
			if ot, ok = x.OneSInt32.(*TypeOneOf1_FSint32B); !ok {
				ot = new(TypeOneOf1_FSint32B)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, true); err != nil {
				return err
			}
			ot.FSint32B = vv
			x.OneSInt32 = ot
		case "one_sint64":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneSInt64 = nil
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
				case "f_sint64a":
					if oneOfIsFill_OneSInt64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneSInt64 = true

					var ok bool
					var ot *TypeOneOf1_FSint64A
					if ot, ok = x.OneSInt64.(*TypeOneOf1_FSint64A); !ok {
						ot = new(TypeOneOf1_FSint64A)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
						return err
					}
					ot.FSint64A = vv
					x.OneSInt64 = ot
				case "f_sint64b":
					if oneOfIsFill_OneSInt64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneSInt64 = true

					var ok bool
					var ot *TypeOneOf1_FSint64B
					if ot, ok = x.OneSInt64.(*TypeOneOf1_FSint64B); !ok {
						ot = new(TypeOneOf1_FSint64B)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, true); err != nil {
						return err
					}
					ot.FSint64B = vv
					x.OneSInt64 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_sfixed32a":
			if oneOfIsFill_OneSFixed32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneSFixed32 = true

			var ok bool
			var ot *TypeOneOf1_FSfixed32A
			if ot, ok = x.OneSFixed32.(*TypeOneOf1_FSfixed32A); !ok {
				ot = new(TypeOneOf1_FSfixed32A)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
				return err
			}
			ot.FSfixed32A = vv
			x.OneSFixed32 = ot
		case "f_sfixed32b":
			if oneOfIsFill_OneSFixed32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneSFixed32 = true

			var ok bool
			var ot *TypeOneOf1_FSfixed32B
			if ot, ok = x.OneSFixed32.(*TypeOneOf1_FSfixed32B); !ok {
				ot = new(TypeOneOf1_FSfixed32B)
			}
			var vv int32
			if vv, err = decoder.ReadLiteralInt32(jsonKey, true); err != nil {
				return err
			}
			ot.FSfixed32B = vv
			x.OneSFixed32 = ot
		case "one_sfixed64":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneSFixed64 = nil
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
				case "f_sfixed64a":
					if oneOfIsFill_OneSFixed64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneSFixed64 = true

					var ok bool
					var ot *TypeOneOf1_FSfixed64A
					if ot, ok = x.OneSFixed64.(*TypeOneOf1_FSfixed64A); !ok {
						ot = new(TypeOneOf1_FSfixed64A)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
						return err
					}
					ot.FSfixed64A = vv
					x.OneSFixed64 = ot
				case "f_sfixed64b":
					if oneOfIsFill_OneSFixed64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneSFixed64 = true

					var ok bool
					var ot *TypeOneOf1_FSfixed64B
					if ot, ok = x.OneSFixed64.(*TypeOneOf1_FSfixed64B); !ok {
						ot = new(TypeOneOf1_FSfixed64B)
					}
					var vv int64
					if vv, err = decoder.ReadLiteralInt64(jsonKey, true); err != nil {
						return err
					}
					ot.FSfixed64B = vv
					x.OneSFixed64 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_fixed32a":
			if oneOfIsFill_OneFixed32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneFixed32 = true

			var ok bool
			var ot *TypeOneOf1_FFixed32A
			if ot, ok = x.OneFixed32.(*TypeOneOf1_FFixed32A); !ok {
				ot = new(TypeOneOf1_FFixed32A)
			}
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
				return err
			}
			ot.FFixed32A = vv
			x.OneFixed32 = ot
		case "f_fixed32b":
			if oneOfIsFill_OneFixed32 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneFixed32 = true

			var ok bool
			var ot *TypeOneOf1_FFixed32B
			if ot, ok = x.OneFixed32.(*TypeOneOf1_FFixed32B); !ok {
				ot = new(TypeOneOf1_FFixed32B)
			}
			var vv uint32
			if vv, err = decoder.ReadLiteralUint32(jsonKey, true); err != nil {
				return err
			}
			ot.FFixed32B = vv
			x.OneFixed32 = ot
		case "one_fixed64":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneFixed64 = nil
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
				case "f_fixed64a":
					if oneOfIsFill_OneFixed64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneFixed64 = true

					var ok bool
					var ot *TypeOneOf1_FFixed64A
					if ot, ok = x.OneFixed64.(*TypeOneOf1_FFixed64A); !ok {
						ot = new(TypeOneOf1_FFixed64A)
					}
					var vv uint64
					if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
						return err
					}
					ot.FFixed64A = vv
					x.OneFixed64 = ot
				case "f_fixed64b":
					if oneOfIsFill_OneFixed64 {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneFixed64 = true

					var ok bool
					var ot *TypeOneOf1_FFixed64B
					if ot, ok = x.OneFixed64.(*TypeOneOf1_FFixed64B); !ok {
						ot = new(TypeOneOf1_FFixed64B)
					}
					var vv uint64
					if vv, err = decoder.ReadLiteralUint64(jsonKey, true); err != nil {
						return err
					}
					ot.FFixed64B = vv
					x.OneFixed64 = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_float1":
			if oneOfIsFill_OneFloat {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneFloat = true

			var ok bool
			var ot *TypeOneOf1_FFloat1
			if ot, ok = x.OneFloat.(*TypeOneOf1_FFloat1); !ok {
				ot = new(TypeOneOf1_FFloat1)
			}
			var vv float32
			if vv, err = decoder.ReadLiteralFloat32(jsonKey, false); err != nil {
				return err
			}
			ot.FFloat1 = vv
			x.OneFloat = ot
		case "f_float2":
			if oneOfIsFill_OneFloat {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneFloat = true

			var ok bool
			var ot *TypeOneOf1_FFloat2
			if ot, ok = x.OneFloat.(*TypeOneOf1_FFloat2); !ok {
				ot = new(TypeOneOf1_FFloat2)
			}
			var vv float32
			if vv, err = decoder.ReadLiteralFloat32(jsonKey, true); err != nil {
				return err
			}
			ot.FFloat2 = vv
			x.OneFloat = ot
		case "one_double":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.OneDouble = nil
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
				case "f_double1":
					if oneOfIsFill_OneDouble {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneDouble = true

					var ok bool
					var ot *TypeOneOf1_FDouble1
					if ot, ok = x.OneDouble.(*TypeOneOf1_FDouble1); !ok {
						ot = new(TypeOneOf1_FDouble1)
					}
					var vv float64
					if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
						return err
					}
					ot.FDouble1 = vv
					x.OneDouble = ot
				case "f_double2":
					if oneOfIsFill_OneDouble {
						return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
					}
					oneOfIsFill_OneDouble = true

					var ok bool
					var ot *TypeOneOf1_FDouble2
					if ot, ok = x.OneDouble.(*TypeOneOf1_FDouble2); !ok {
						ot = new(TypeOneOf1_FDouble2)
					}
					var vv float64
					if vv, err = decoder.ReadLiteralFloat64(jsonKey, true); err != nil {
						return err
					}
					ot.FDouble2 = vv
					x.OneDouble = ot
				default:
					if err = decoder.DiscardValue(jsonKey); err != nil {
						return err
					}
				} // end switch
			}
		case "f_bool1":
			if oneOfIsFill_OneBool {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneBool = true

			var ok bool
			var ot *TypeOneOf1_FBool1
			if ot, ok = x.OneBool.(*TypeOneOf1_FBool1); !ok {
				ot = new(TypeOneOf1_FBool1)
			}
			var vv bool
			if vv, err = decoder.ReadLiteralBool(jsonKey, false); err != nil {
				return err
			}
			ot.FBool1 = vv
			x.OneBool = ot
		case "f_bool2":
			if oneOfIsFill_OneBool {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneBool = true

			var ok bool
			var ot *TypeOneOf1_FBool2
			if ot, ok = x.OneBool.(*TypeOneOf1_FBool2); !ok {
				ot = new(TypeOneOf1_FBool2)
			}
			var vv bool
			if vv, err = decoder.ReadLiteralBool(jsonKey, true); err != nil {
				return err
			}
			ot.FBool2 = vv
			x.OneBool = ot
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}