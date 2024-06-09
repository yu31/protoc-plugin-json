// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/errors/invalid_syntax_map.proto

package pberrors

import (
	errors "errors"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message InvalidSyntaxMap in file tests/proto/cases/errors/invalid_syntax_map.proto
func (x *InvalidSyntaxMap) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(136)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapI32I32(enc, "f_int32_key_string_val_numeric", x.FInt32KeyStringValNumeric, false, true, false)
	jsonencoder.AppendMapI32I32(enc, "f_int32_key_numeric_val_string", x.FInt32KeyNumericValString, false, true, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message InvalidSyntaxMap in file tests/proto/cases/errors/invalid_syntax_map.proto
func (x *InvalidSyntaxMap) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pberrors.(*InvalidSyntaxMap) is nil")
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
		case "f_int32_key_string_val_numeric":
			if x.FInt32KeyStringValNumeric, err = jsondecoder.ReadMapI32I32(dec, x.FInt32KeyStringValNumeric, true, false); err != nil {
				return err
			}
		case "f_int32_key_numeric_val_string":
			if x.FInt32KeyNumericValString, err = jsondecoder.ReadMapI32I32(dec, x.FInt32KeyNumericValString, true, false); err != nil {
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