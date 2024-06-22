// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/boundary/cycle1.proto

package pbboundary

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message InlineMessageCycle1 in file tests/proto/cases/boundary/cycle1.proto
func (x *InlineMessageCycle1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(80)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	if a1_3 := x.FMessage1; a1_3 != nil {
		jsonencoder.AppendValStr(enc, "f_string3", a1_3.FString3, false)
		jsonencoder.AppendValStr(enc, "f_string4", a1_3.FString4, false)
		if err := jsonencoder.AppendValMessage(enc, "f_message2", a1_3.FMessage2, false); err != nil {
			return nil, err
		}
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineMessageCycle1 in file tests/proto/cases/boundary/cycle1.proto
func (x *InlineMessageCycle1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*InlineMessageCycle1) is nil")
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
	// declares variables to simple to reference parent field
	var (
		a1_3 *InlineMessageCycle1_EmbedMessage1
	)

	// declares anonymous to init the parent field.
	init_a1_3 := func() error {
		if a1_3 == nil {
			if x.FMessage1 == nil {
				x.FMessage1 = new(InlineMessageCycle1_EmbedMessage1)
			}
			a1_3 = x.FMessage1
		}
		return nil
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
			if err = init_a1_3(); err != nil {
				return err
			}
			if a1_3.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string4":
			if err = init_a1_3(); err != nil {
				return err
			}
			if a1_3.FString4, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_message2":
			if err = init_a1_3(); err != nil {
				return err
			}
			if a1_3.FMessage2, err = jsondecoder.ReadValMessage(dec, a1_3.FMessage2); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message EmbedMessage1 in file tests/proto/cases/boundary/cycle1.proto
func (x *InlineMessageCycle1_EmbedMessage1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(80)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	jsonencoder.AppendValStr(enc, "f_string4", x.FString4, false)
	if err := jsonencoder.AppendValMessage(enc, "f_message2", x.FMessage2, false); err != nil {
		return nil, err
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message EmbedMessage1 in file tests/proto/cases/boundary/cycle1.proto
func (x *InlineMessageCycle1_EmbedMessage1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*InlineMessageCycle1_EmbedMessage1) is nil")
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
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string4":
			if x.FString4, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_message2":
			if x.FMessage2, err = jsondecoder.ReadValMessage(dec, x.FMessage2); err != nil {
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