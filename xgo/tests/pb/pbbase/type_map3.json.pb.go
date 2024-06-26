// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/base/type_map3.proto

package pbbase

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeMap3 in file tests/proto/cases/base/type_map3.proto
func (x *TypeMap3) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(1280)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapBoolStr(enc, "f_string1", x.FString1, false, true)
	jsonencoder.AppendMapBoolI32(enc, "f_int32", x.FInt32, false, true, false)
	jsonencoder.AppendMapBoolI64(enc, "f_int64", x.FInt64, false, true, false)
	jsonencoder.AppendMapBoolU32(enc, "f_uint32", x.FUint32, false, true, false)
	jsonencoder.AppendMapBoolU64(enc, "f_uint64", x.FUint64, false, true, false)
	jsonencoder.AppendMapBoolI32(enc, "f_sint32", x.FSint32, false, true, false)
	jsonencoder.AppendMapBoolI64(enc, "f_sint64", x.FSint64, false, true, false)
	jsonencoder.AppendMapBoolI32(enc, "f_sfixed32", x.FSfixed32, false, true, false)
	jsonencoder.AppendMapBoolI64(enc, "f_sfixed64", x.FSfixed64, false, true, false)
	jsonencoder.AppendMapBoolU32(enc, "f_fixed32", x.FFixed32, false, true, false)
	jsonencoder.AppendMapBoolU64(enc, "f_fixed64", x.FFixed64, false, true, false)
	jsonencoder.AppendMapBoolBool(enc, "f_bool", x.FBool, false, true, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeMap3 in file tests/proto/cases/base/type_map3.proto
func (x *TypeMap3) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbbase", "TypeMap3")
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
			if x.FString1, err = jsondecoder.ReadMapBoolStr(dec, x.FString1, true); err != nil {
				return err
			}
		case "f_int32":
			if x.FInt32, err = jsondecoder.ReadMapBoolI32(dec, x.FInt32, true, false); err != nil {
				return err
			}
		case "f_int64":
			if x.FInt64, err = jsondecoder.ReadMapBoolI64(dec, x.FInt64, true, false); err != nil {
				return err
			}
		case "f_uint32":
			if x.FUint32, err = jsondecoder.ReadMapBoolU32(dec, x.FUint32, true, false); err != nil {
				return err
			}
		case "f_uint64":
			if x.FUint64, err = jsondecoder.ReadMapBoolU64(dec, x.FUint64, true, false); err != nil {
				return err
			}
		case "f_sint32":
			if x.FSint32, err = jsondecoder.ReadMapBoolI32(dec, x.FSint32, true, false); err != nil {
				return err
			}
		case "f_sint64":
			if x.FSint64, err = jsondecoder.ReadMapBoolI64(dec, x.FSint64, true, false); err != nil {
				return err
			}
		case "f_sfixed32":
			if x.FSfixed32, err = jsondecoder.ReadMapBoolI32(dec, x.FSfixed32, true, false); err != nil {
				return err
			}
		case "f_sfixed64":
			if x.FSfixed64, err = jsondecoder.ReadMapBoolI64(dec, x.FSfixed64, true, false); err != nil {
				return err
			}
		case "f_fixed32":
			if x.FFixed32, err = jsondecoder.ReadMapBoolU32(dec, x.FFixed32, true, false); err != nil {
				return err
			}
		case "f_fixed64":
			if x.FFixed64, err = jsondecoder.ReadMapBoolU64(dec, x.FFixed64, true, false); err != nil {
				return err
			}
		case "f_bool":
			if x.FBool, err = jsondecoder.ReadMapBoolBool(dec, x.FBool, true, false); err != nil {
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
