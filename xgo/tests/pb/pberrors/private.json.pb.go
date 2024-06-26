// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/errors/private.proto

package pberrors

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message MessageCommon1 in file tests/proto/cases/errors/private.proto
func (x *MessageCommon1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(136)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageCommon1 in file tests/proto/cases/errors/private.proto
func (x *MessageCommon1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pberrors", "MessageCommon1")
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
			if x.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Embed1 in file tests/proto/cases/errors/private.proto
func (x *MessageCommon1_Embed1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(136)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed1 in file tests/proto/cases/errors/private.proto
func (x *MessageCommon1_Embed1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pberrors", "MessageCommon1_Embed1")
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
			if x.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Embed2 in file tests/proto/cases/errors/private.proto
func (x *MessageCommon1_Embed1_Embed2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(136)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Embed2 in file tests/proto/cases/errors/private.proto
func (x *MessageCommon1_Embed1_Embed2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pberrors", "MessageCommon1_Embed1_Embed2")
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
			if x.FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string2":
			if x.FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
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
