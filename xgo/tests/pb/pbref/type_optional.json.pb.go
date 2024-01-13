// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/references/type_optional.proto

package pbref

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeOptional1 in file tests/proto/cases/references/type_optional.proto
func (x *TypeOptional1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(584)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_int32a")
	encoder.AppendPointerInt32(x.FInt32A, false)
	encoder.AppendObjectKey("f_int32b")
	encoder.AppendPointerInt32(x.FInt32B, true)
	encoder.AppendObjectKey("f_int64a")
	encoder.AppendPointerInt64(x.FInt64A, false)
	encoder.AppendObjectKey("f_int64b")
	encoder.AppendPointerInt64(x.FInt64B, true)
	encoder.AppendObjectKey("f_uint32a")
	encoder.AppendPointerUint32(x.FUint32A, false)
	encoder.AppendObjectKey("f_uint32b")
	encoder.AppendPointerUint32(x.FUint32B, true)
	encoder.AppendObjectKey("f_uint64a")
	encoder.AppendPointerUint64(x.FUint64A, false)
	encoder.AppendObjectKey("f_uint64b")
	encoder.AppendPointerUint64(x.FUint64B, true)
	encoder.AppendObjectKey("f_sint32a")
	encoder.AppendPointerInt32(x.FSint32A, false)
	encoder.AppendObjectKey("f_sint32b")
	encoder.AppendPointerInt32(x.FSint32B, true)
	encoder.AppendObjectKey("f_sint64a")
	encoder.AppendPointerInt64(x.FSint64A, false)
	encoder.AppendObjectKey("f_sint64b")
	encoder.AppendPointerInt64(x.FSint64B, true)
	encoder.AppendObjectKey("f_sfixed32a")
	encoder.AppendPointerInt32(x.FSfixed32A, false)
	encoder.AppendObjectKey("f_sfixed32b")
	encoder.AppendPointerInt32(x.FSfixed32B, true)
	encoder.AppendObjectKey("f_sfixed64a")
	encoder.AppendPointerInt64(x.FSfixed64A, false)
	encoder.AppendObjectKey("f_sfixed64b")
	encoder.AppendPointerInt64(x.FSfixed64B, true)
	encoder.AppendObjectKey("f_fixed32a")
	encoder.AppendPointerUint32(x.FFixed32A, false)
	encoder.AppendObjectKey("f_fixed32b")
	encoder.AppendPointerUint32(x.FFixed32B, true)
	encoder.AppendObjectKey("f_fixed64a")
	encoder.AppendPointerUint64(x.FFixed64A, false)
	encoder.AppendObjectKey("f_fixed64b")
	encoder.AppendPointerUint64(x.FFixed64B, true)
	encoder.AppendObjectKey("f_float1")
	encoder.AppendPointerFloat32(x.FFloat1, false)
	encoder.AppendObjectKey("f_float2")
	encoder.AppendPointerFloat32(x.FFloat2, true)
	encoder.AppendObjectKey("f_double1")
	encoder.AppendPointerFloat64(x.FDouble1, false)
	encoder.AppendObjectKey("f_double2")
	encoder.AppendPointerFloat64(x.FDouble2, true)
	encoder.AppendObjectKey("f_bool1")
	encoder.AppendPointerBool(x.FBool1, false)
	encoder.AppendObjectKey("f_bool2")
	encoder.AppendPointerBool(x.FBool2, true)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOptional1 in file tests/proto/cases/references/type_optional.proto
func (x *TypeOptional1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbref.(*TypeOptional1) is nil")
	}
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
		case "f_int32a":
			var vv *int32
			if vv, err = decoder.ReadPointerInt32(jsonKey, false); err != nil {
				return err
			}
			x.FInt32A = vv
		case "f_int32b":
			var vv *int32
			if vv, err = decoder.ReadPointerInt32(jsonKey, true); err != nil {
				return err
			}
			x.FInt32B = vv
		case "f_int64a":
			var vv *int64
			if vv, err = decoder.ReadPointerInt64(jsonKey, false); err != nil {
				return err
			}
			x.FInt64A = vv
		case "f_int64b":
			var vv *int64
			if vv, err = decoder.ReadPointerInt64(jsonKey, true); err != nil {
				return err
			}
			x.FInt64B = vv
		case "f_uint32a":
			var vv *uint32
			if vv, err = decoder.ReadPointerUint32(jsonKey, false); err != nil {
				return err
			}
			x.FUint32A = vv
		case "f_uint32b":
			var vv *uint32
			if vv, err = decoder.ReadPointerUint32(jsonKey, true); err != nil {
				return err
			}
			x.FUint32B = vv
		case "f_uint64a":
			var vv *uint64
			if vv, err = decoder.ReadPointerUint64(jsonKey, false); err != nil {
				return err
			}
			x.FUint64A = vv
		case "f_uint64b":
			var vv *uint64
			if vv, err = decoder.ReadPointerUint64(jsonKey, true); err != nil {
				return err
			}
			x.FUint64B = vv
		case "f_sint32a":
			var vv *int32
			if vv, err = decoder.ReadPointerInt32(jsonKey, false); err != nil {
				return err
			}
			x.FSint32A = vv
		case "f_sint32b":
			var vv *int32
			if vv, err = decoder.ReadPointerInt32(jsonKey, true); err != nil {
				return err
			}
			x.FSint32B = vv
		case "f_sint64a":
			var vv *int64
			if vv, err = decoder.ReadPointerInt64(jsonKey, false); err != nil {
				return err
			}
			x.FSint64A = vv
		case "f_sint64b":
			var vv *int64
			if vv, err = decoder.ReadPointerInt64(jsonKey, true); err != nil {
				return err
			}
			x.FSint64B = vv
		case "f_sfixed32a":
			var vv *int32
			if vv, err = decoder.ReadPointerInt32(jsonKey, false); err != nil {
				return err
			}
			x.FSfixed32A = vv
		case "f_sfixed32b":
			var vv *int32
			if vv, err = decoder.ReadPointerInt32(jsonKey, true); err != nil {
				return err
			}
			x.FSfixed32B = vv
		case "f_sfixed64a":
			var vv *int64
			if vv, err = decoder.ReadPointerInt64(jsonKey, false); err != nil {
				return err
			}
			x.FSfixed64A = vv
		case "f_sfixed64b":
			var vv *int64
			if vv, err = decoder.ReadPointerInt64(jsonKey, true); err != nil {
				return err
			}
			x.FSfixed64B = vv
		case "f_fixed32a":
			var vv *uint32
			if vv, err = decoder.ReadPointerUint32(jsonKey, false); err != nil {
				return err
			}
			x.FFixed32A = vv
		case "f_fixed32b":
			var vv *uint32
			if vv, err = decoder.ReadPointerUint32(jsonKey, true); err != nil {
				return err
			}
			x.FFixed32B = vv
		case "f_fixed64a":
			var vv *uint64
			if vv, err = decoder.ReadPointerUint64(jsonKey, false); err != nil {
				return err
			}
			x.FFixed64A = vv
		case "f_fixed64b":
			var vv *uint64
			if vv, err = decoder.ReadPointerUint64(jsonKey, true); err != nil {
				return err
			}
			x.FFixed64B = vv
		case "f_float1":
			var vv *float32
			if vv, err = decoder.ReadPointerFloat32(jsonKey, false); err != nil {
				return err
			}
			x.FFloat1 = vv
		case "f_float2":
			var vv *float32
			if vv, err = decoder.ReadPointerFloat32(jsonKey, true); err != nil {
				return err
			}
			x.FFloat2 = vv
		case "f_double1":
			var vv *float64
			if vv, err = decoder.ReadPointerFloat64(jsonKey, false); err != nil {
				return err
			}
			x.FDouble1 = vv
		case "f_double2":
			var vv *float64
			if vv, err = decoder.ReadPointerFloat64(jsonKey, true); err != nil {
				return err
			}
			x.FDouble2 = vv
		case "f_bool1":
			var vv *bool
			if vv, err = decoder.ReadPointerBool(jsonKey, false); err != nil {
				return err
			}
			x.FBool1 = vv
		case "f_bool2":
			var vv *bool
			if vv, err = decoder.ReadPointerBool(jsonKey, true); err != nil {
				return err
			}
			x.FBool2 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message TypeOptional2 in file tests/proto/cases/references/type_optional.proto
func (x *TypeOptional2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(48)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("f_string1")
	encoder.AppendPointerString(x.FString1)
	encoder.AppendObjectKey("f_bytes1")
	encoder.AppendLiteralBytes(x.FBytes1)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeOptional2 in file tests/proto/cases/references/type_optional.proto
func (x *TypeOptional2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbref.(*TypeOptional2) is nil")
	}
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
		case "f_string1":
			var vv *string
			if vv, err = decoder.ReadPointerString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "f_bytes1":
			var vv []byte
			if vv, err = decoder.ReadLiteralBytes(jsonKey); err != nil {
				return err
			}
			x.FBytes1 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}