// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/boundary/cycle5.proto

package pbboundary

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message InlineMessageCycle5 in file tests/proto/cases/boundary/cycle5.proto
func (x *InlineMessageCycle5) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(392)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	if a1_3 := x.FMessage1; a1_3 != nil {
		jsonencoder.AppendValStr(enc, "f_string3", a1_3.FString3, false)
		jsonencoder.AppendValStr(enc, "f_string4", a1_3.FString4, false)
		switch o2_6 := a1_3.OneType2.(type) {
		case *InlineMessageCycle5_EmbedMessage1_FMessage2:
			if err := jsonencoder.AppendValMessage(enc, "f_message2", o2_6.FMessage2, false); err != nil {
				return nil, err
			}
		case *InlineMessageCycle5_EmbedMessage1_FString5:
			jsonencoder.AppendValStr(enc, "f_string5", o2_6.FString5, false)
		case *InlineMessageCycle5_EmbedMessage1_FString6:
			jsonencoder.AppendValStr(enc, "f_string6", o2_6.FString6, false)
		}
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineMessageCycle5 in file tests/proto/cases/boundary/cycle5.proto
func (x *InlineMessageCycle5) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbboundary", "InlineMessageCycle5")
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
		isLoad_o2_6 bool
	)

	// declares variables for simple to reference parent field
	var (
		a1_3 *InlineMessageCycle5_EmbedMessage1
	)

	// declares anonymous func to init the parent field.
	init_a1_3 := func() error {
		if a1_3 == nil {
			if x.FMessage1 == nil {
				x.FMessage1 = new(InlineMessageCycle5_EmbedMessage1)
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
			if isLoad_o2_6 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o2_6 = true
			o2_6, ok := a1_3.OneType2.(*InlineMessageCycle5_EmbedMessage1_FMessage2)
			if !ok {
				o2_6 = new(InlineMessageCycle5_EmbedMessage1_FMessage2)
				a1_3.OneType2 = o2_6
			}
			if o2_6.FMessage2, err = jsondecoder.ReadValMessage(dec, o2_6.FMessage2); err != nil {
				return err
			}
		case "f_string5":
			if err = init_a1_3(); err != nil {
				return err
			}
			if isLoad_o2_6 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o2_6 = true
			o2_6, ok := a1_3.OneType2.(*InlineMessageCycle5_EmbedMessage1_FString5)
			if !ok {
				o2_6 = new(InlineMessageCycle5_EmbedMessage1_FString5)
				a1_3.OneType2 = o2_6
			}
			if o2_6.FString5, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string6":
			if err = init_a1_3(); err != nil {
				return err
			}
			if isLoad_o2_6 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o2_6 = true
			o2_6, ok := a1_3.OneType2.(*InlineMessageCycle5_EmbedMessage1_FString6)
			if !ok {
				o2_6 = new(InlineMessageCycle5_EmbedMessage1_FString6)
				a1_3.OneType2 = o2_6
			}
			if o2_6.FString6, err = jsondecoder.ReadValStr(dec); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message EmbedMessage1 in file tests/proto/cases/boundary/cycle5.proto
func (x *InlineMessageCycle5_EmbedMessage1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(296)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	jsonencoder.AppendValStr(enc, "f_string4", x.FString4, false)
	switch o1_3 := x.OneType2.(type) {
	case *InlineMessageCycle5_EmbedMessage1_FMessage2:
		if err := jsonencoder.AppendValMessage(enc, "f_message2", o1_3.FMessage2, false); err != nil {
			return nil, err
		}
	case *InlineMessageCycle5_EmbedMessage1_FString5:
		jsonencoder.AppendValStr(enc, "f_string5", o1_3.FString5, false)
	case *InlineMessageCycle5_EmbedMessage1_FString6:
		jsonencoder.AppendValStr(enc, "f_string6", o1_3.FString6, false)
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message EmbedMessage1 in file tests/proto/cases/boundary/cycle5.proto
func (x *InlineMessageCycle5_EmbedMessage1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbboundary", "InlineMessageCycle5_EmbedMessage1")
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
		isLoad_o1_3 bool
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
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string4":
			if x.FString4, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_message2":
			if isLoad_o1_3 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_3 = true
			o1_3, ok := x.OneType2.(*InlineMessageCycle5_EmbedMessage1_FMessage2)
			if !ok {
				o1_3 = new(InlineMessageCycle5_EmbedMessage1_FMessage2)
				x.OneType2 = o1_3
			}
			if o1_3.FMessage2, err = jsondecoder.ReadValMessage(dec, o1_3.FMessage2); err != nil {
				return err
			}
		case "f_string5":
			if isLoad_o1_3 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_3 = true
			o1_3, ok := x.OneType2.(*InlineMessageCycle5_EmbedMessage1_FString5)
			if !ok {
				o1_3 = new(InlineMessageCycle5_EmbedMessage1_FString5)
				x.OneType2 = o1_3
			}
			if o1_3.FString5, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string6":
			if isLoad_o1_3 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_3 = true
			o1_3, ok := x.OneType2.(*InlineMessageCycle5_EmbedMessage1_FString6)
			if !ok {
				o1_3 = new(InlineMessageCycle5_EmbedMessage1_FString6)
				x.OneType2 = o1_3
			}
			if o1_3.FString6, err = jsondecoder.ReadValStr(dec); err != nil {
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
