// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/references/type_map.proto

package pbref

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeMap1 in file tests/proto/cases/references/type_map.proto
func (x *TypeMap1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(552)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapI32I32(enc, "f_int32a", x.FInt32A, false, true, false)
	jsonencoder.AppendMapI32I32(enc, "f_int32b", x.FInt32B, false, true, true)
	jsonencoder.AppendMapI64I64(enc, "f_int64a", x.FInt64A, false, true, false)
	jsonencoder.AppendMapI64I64(enc, "f_int64b", x.FInt64B, false, true, true)
	jsonencoder.AppendMapU32U32(enc, "f_uint32a", x.FUint32A, false, true, false)
	jsonencoder.AppendMapU32U32(enc, "f_uint32b", x.FUint32B, false, true, true)
	jsonencoder.AppendMapU64U64(enc, "f_uint64a", x.FUint64A, false, true, false)
	jsonencoder.AppendMapU64U64(enc, "f_uint64b", x.FUint64B, false, true, true)
	jsonencoder.AppendMapI32I32(enc, "f_sint32a", x.FSint32A, false, true, false)
	jsonencoder.AppendMapI32I32(enc, "f_sint32b", x.FSint32B, false, true, true)
	jsonencoder.AppendMapI64I64(enc, "f_sint64a", x.FSint64A, false, true, false)
	jsonencoder.AppendMapI64I64(enc, "f_sint64b", x.FSint64B, false, true, true)
	jsonencoder.AppendMapI32I32(enc, "f_sfixed32a", x.FSfixed32A, false, true, false)
	jsonencoder.AppendMapI32I32(enc, "f_sfixed32b", x.FSfixed32B, false, true, true)
	jsonencoder.AppendMapI64I64(enc, "f_sfixed64a", x.FSfixed64A, false, true, false)
	jsonencoder.AppendMapI64I64(enc, "f_sfixed64b", x.FSfixed64B, false, true, true)
	jsonencoder.AppendMapU32U32(enc, "f_fixed32a", x.FFixed32A, false, true, false)
	jsonencoder.AppendMapU32U32(enc, "f_fixed32b", x.FFixed32B, false, true, true)
	jsonencoder.AppendMapU64U64(enc, "f_fixed64a", x.FFixed64A, false, true, false)
	jsonencoder.AppendMapU64U64(enc, "f_fixed64b", x.FFixed64B, false, true, true)
	jsonencoder.AppendMapStrF32(enc, "f_float1", x.FFloat1, false, false)
	jsonencoder.AppendMapStrF32(enc, "f_float2", x.FFloat2, false, true)
	jsonencoder.AppendMapStrF64(enc, "f_double1", x.FDouble1, false, false)
	jsonencoder.AppendMapStrF64(enc, "f_double2", x.FDouble2, false, true)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeMap1 in file tests/proto/cases/references/type_map.proto
func (x *TypeMap1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbref.(*TypeMap1) is nil")
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
		case "f_int32a":
			if x.FInt32A, err = jsondecoder.ReadMapI32I32(dec, x.FInt32A, true, false); err != nil {
				return err
			}
		case "f_int32b":
			if x.FInt32B, err = jsondecoder.ReadMapI32I32(dec, x.FInt32B, true, true); err != nil {
				return err
			}
		case "f_int64a":
			if x.FInt64A, err = jsondecoder.ReadMapI64I64(dec, x.FInt64A, true, false); err != nil {
				return err
			}
		case "f_int64b":
			if x.FInt64B, err = jsondecoder.ReadMapI64I64(dec, x.FInt64B, true, true); err != nil {
				return err
			}
		case "f_uint32a":
			if x.FUint32A, err = jsondecoder.ReadMapU32U32(dec, x.FUint32A, true, false); err != nil {
				return err
			}
		case "f_uint32b":
			if x.FUint32B, err = jsondecoder.ReadMapU32U32(dec, x.FUint32B, true, true); err != nil {
				return err
			}
		case "f_uint64a":
			if x.FUint64A, err = jsondecoder.ReadMapU64U64(dec, x.FUint64A, true, false); err != nil {
				return err
			}
		case "f_uint64b":
			if x.FUint64B, err = jsondecoder.ReadMapU64U64(dec, x.FUint64B, true, true); err != nil {
				return err
			}
		case "f_sint32a":
			if x.FSint32A, err = jsondecoder.ReadMapI32I32(dec, x.FSint32A, true, false); err != nil {
				return err
			}
		case "f_sint32b":
			if x.FSint32B, err = jsondecoder.ReadMapI32I32(dec, x.FSint32B, true, true); err != nil {
				return err
			}
		case "f_sint64a":
			if x.FSint64A, err = jsondecoder.ReadMapI64I64(dec, x.FSint64A, true, false); err != nil {
				return err
			}
		case "f_sint64b":
			if x.FSint64B, err = jsondecoder.ReadMapI64I64(dec, x.FSint64B, true, true); err != nil {
				return err
			}
		case "f_sfixed32a":
			if x.FSfixed32A, err = jsondecoder.ReadMapI32I32(dec, x.FSfixed32A, true, false); err != nil {
				return err
			}
		case "f_sfixed32b":
			if x.FSfixed32B, err = jsondecoder.ReadMapI32I32(dec, x.FSfixed32B, true, true); err != nil {
				return err
			}
		case "f_sfixed64a":
			if x.FSfixed64A, err = jsondecoder.ReadMapI64I64(dec, x.FSfixed64A, true, false); err != nil {
				return err
			}
		case "f_sfixed64b":
			if x.FSfixed64B, err = jsondecoder.ReadMapI64I64(dec, x.FSfixed64B, true, true); err != nil {
				return err
			}
		case "f_fixed32a":
			if x.FFixed32A, err = jsondecoder.ReadMapU32U32(dec, x.FFixed32A, true, false); err != nil {
				return err
			}
		case "f_fixed32b":
			if x.FFixed32B, err = jsondecoder.ReadMapU32U32(dec, x.FFixed32B, true, true); err != nil {
				return err
			}
		case "f_fixed64a":
			if x.FFixed64A, err = jsondecoder.ReadMapU64U64(dec, x.FFixed64A, true, false); err != nil {
				return err
			}
		case "f_fixed64b":
			if x.FFixed64B, err = jsondecoder.ReadMapU64U64(dec, x.FFixed64B, true, true); err != nil {
				return err
			}
		case "f_float1":
			if x.FFloat1, err = jsondecoder.ReadMapStrF32(dec, x.FFloat1, false); err != nil {
				return err
			}
		case "f_float2":
			if x.FFloat2, err = jsondecoder.ReadMapStrF32(dec, x.FFloat2, true); err != nil {
				return err
			}
		case "f_double1":
			if x.FDouble1, err = jsondecoder.ReadMapStrF64(dec, x.FDouble1, false); err != nil {
				return err
			}
		case "f_double2":
			if x.FDouble2, err = jsondecoder.ReadMapStrF64(dec, x.FDouble2, true); err != nil {
				return err
			}
		default:
			if err = dec.DiscardValue(); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message TypeMap2 in file tests/proto/cases/references/type_map.proto
func (x *TypeMap2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(584)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapI32I32(enc, "f_int32a", x.FInt32A, false, false, false)
	jsonencoder.AppendMapI32I32(enc, "f_int32b", x.FInt32B, false, true, true)
	jsonencoder.AppendMapI64I64(enc, "f_int64a", x.FInt64A, false, false, false)
	jsonencoder.AppendMapI64I64(enc, "f_int64b", x.FInt64B, false, true, true)
	jsonencoder.AppendMapU32U32(enc, "f_uint32a", x.FUint32A, false, false, false)
	jsonencoder.AppendMapU32U32(enc, "f_uint32b", x.FUint32B, false, true, true)
	jsonencoder.AppendMapU64U64(enc, "f_uint64a", x.FUint64A, false, false, false)
	jsonencoder.AppendMapU64U64(enc, "f_uint64b", x.FUint64B, false, true, true)
	jsonencoder.AppendMapI32I32(enc, "f_sint32a", x.FSint32A, false, false, false)
	jsonencoder.AppendMapI32I32(enc, "f_sint32b", x.FSint32B, false, true, true)
	jsonencoder.AppendMapI64I64(enc, "f_sint64a", x.FSint64A, false, false, false)
	jsonencoder.AppendMapI64I64(enc, "f_sint64b", x.FSint64B, false, true, true)
	jsonencoder.AppendMapI32I32(enc, "f_sfixed32a", x.FSfixed32A, false, false, false)
	jsonencoder.AppendMapI32I32(enc, "f_sfixed32b", x.FSfixed32B, false, true, true)
	jsonencoder.AppendMapI64I64(enc, "f_sfixed64a", x.FSfixed64A, false, false, false)
	jsonencoder.AppendMapI64I64(enc, "f_sfixed64b", x.FSfixed64B, false, true, true)
	jsonencoder.AppendMapU32U32(enc, "f_fixed32a", x.FFixed32A, false, false, false)
	jsonencoder.AppendMapU32U32(enc, "f_fixed32b", x.FFixed32B, false, true, true)
	jsonencoder.AppendMapU64U64(enc, "f_fixed64a", x.FFixed64A, false, false, false)
	jsonencoder.AppendMapU64U64(enc, "f_fixed64b", x.FFixed64B, false, true, true)
	jsonencoder.AppendMapStrF32(enc, "f_float1", x.FFloat1, false, false)
	jsonencoder.AppendMapStrF32(enc, "f_float2", x.FFloat2, false, true)
	jsonencoder.AppendMapStrF64(enc, "f_double1", x.FDouble1, false, false)
	jsonencoder.AppendMapStrF64(enc, "f_double2", x.FDouble2, false, true)
	jsonencoder.AppendMapBoolBool(enc, "f_bool1", x.FBool1, false, false, false)
	jsonencoder.AppendMapBoolBool(enc, "f_bool2", x.FBool2, false, true, true)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeMap2 in file tests/proto/cases/references/type_map.proto
func (x *TypeMap2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbref.(*TypeMap2) is nil")
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
		case "f_int32a":
			if x.FInt32A, err = jsondecoder.ReadMapI32I32(dec, x.FInt32A, false, false); err != nil {
				return err
			}
		case "f_int32b":
			if x.FInt32B, err = jsondecoder.ReadMapI32I32(dec, x.FInt32B, true, true); err != nil {
				return err
			}
		case "f_int64a":
			if x.FInt64A, err = jsondecoder.ReadMapI64I64(dec, x.FInt64A, false, false); err != nil {
				return err
			}
		case "f_int64b":
			if x.FInt64B, err = jsondecoder.ReadMapI64I64(dec, x.FInt64B, true, true); err != nil {
				return err
			}
		case "f_uint32a":
			if x.FUint32A, err = jsondecoder.ReadMapU32U32(dec, x.FUint32A, false, false); err != nil {
				return err
			}
		case "f_uint32b":
			if x.FUint32B, err = jsondecoder.ReadMapU32U32(dec, x.FUint32B, true, true); err != nil {
				return err
			}
		case "f_uint64a":
			if x.FUint64A, err = jsondecoder.ReadMapU64U64(dec, x.FUint64A, false, false); err != nil {
				return err
			}
		case "f_uint64b":
			if x.FUint64B, err = jsondecoder.ReadMapU64U64(dec, x.FUint64B, true, true); err != nil {
				return err
			}
		case "f_sint32a":
			if x.FSint32A, err = jsondecoder.ReadMapI32I32(dec, x.FSint32A, false, false); err != nil {
				return err
			}
		case "f_sint32b":
			if x.FSint32B, err = jsondecoder.ReadMapI32I32(dec, x.FSint32B, true, true); err != nil {
				return err
			}
		case "f_sint64a":
			if x.FSint64A, err = jsondecoder.ReadMapI64I64(dec, x.FSint64A, false, false); err != nil {
				return err
			}
		case "f_sint64b":
			if x.FSint64B, err = jsondecoder.ReadMapI64I64(dec, x.FSint64B, true, true); err != nil {
				return err
			}
		case "f_sfixed32a":
			if x.FSfixed32A, err = jsondecoder.ReadMapI32I32(dec, x.FSfixed32A, false, false); err != nil {
				return err
			}
		case "f_sfixed32b":
			if x.FSfixed32B, err = jsondecoder.ReadMapI32I32(dec, x.FSfixed32B, true, true); err != nil {
				return err
			}
		case "f_sfixed64a":
			if x.FSfixed64A, err = jsondecoder.ReadMapI64I64(dec, x.FSfixed64A, false, false); err != nil {
				return err
			}
		case "f_sfixed64b":
			if x.FSfixed64B, err = jsondecoder.ReadMapI64I64(dec, x.FSfixed64B, true, true); err != nil {
				return err
			}
		case "f_fixed32a":
			if x.FFixed32A, err = jsondecoder.ReadMapU32U32(dec, x.FFixed32A, false, false); err != nil {
				return err
			}
		case "f_fixed32b":
			if x.FFixed32B, err = jsondecoder.ReadMapU32U32(dec, x.FFixed32B, true, true); err != nil {
				return err
			}
		case "f_fixed64a":
			if x.FFixed64A, err = jsondecoder.ReadMapU64U64(dec, x.FFixed64A, false, false); err != nil {
				return err
			}
		case "f_fixed64b":
			if x.FFixed64B, err = jsondecoder.ReadMapU64U64(dec, x.FFixed64B, true, true); err != nil {
				return err
			}
		case "f_float1":
			if x.FFloat1, err = jsondecoder.ReadMapStrF32(dec, x.FFloat1, false); err != nil {
				return err
			}
		case "f_float2":
			if x.FFloat2, err = jsondecoder.ReadMapStrF32(dec, x.FFloat2, true); err != nil {
				return err
			}
		case "f_double1":
			if x.FDouble1, err = jsondecoder.ReadMapStrF64(dec, x.FDouble1, false); err != nil {
				return err
			}
		case "f_double2":
			if x.FDouble2, err = jsondecoder.ReadMapStrF64(dec, x.FDouble2, true); err != nil {
				return err
			}
		case "f_bool1":
			if x.FBool1, err = jsondecoder.ReadMapBoolBool(dec, x.FBool1, false, false); err != nil {
				return err
			}
		case "f_bool2":
			if x.FBool2, err = jsondecoder.ReadMapBoolBool(dec, x.FBool2, true, true); err != nil {
				return err
			}
		default:
			if err = dec.DiscardValue(); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message TypeMap3 in file tests/proto/cases/references/type_map.proto
func (x *TypeMap3) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(48)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapStrStr(enc, "f_string1", x.FString1, false)
	if err := jsonencoder.AppendMapStrBytes(enc, "f_bytes1", x.FBytes1, false); err != nil {
		return nil, err
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeMap3 in file tests/proto/cases/references/type_map.proto
func (x *TypeMap3) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbref.(*TypeMap3) is nil")
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
		case "f_string1":
			if x.FString1, err = jsondecoder.ReadMapStrStr(dec, x.FString1); err != nil {
				return err
			}
		case "f_bytes1":
			if x.FBytes1, err = jsondecoder.ReadMapStrBytes(dec, x.FBytes1); err != nil {
				return err
			}
		default:
			if err = dec.DiscardValue(); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
