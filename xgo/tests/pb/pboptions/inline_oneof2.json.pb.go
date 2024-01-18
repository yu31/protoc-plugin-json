// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/options/inline_oneof2.proto

package pboptions

import (
	errors "errors"
	fmt "fmt"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message InlineOneOf2Message3 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2Message3) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("t_string3")
	encoder.AppendLiteralString(x.FString1)
	switch ov := x.OneType01.(type) {
	case *InlineOneOf2Message3_FDouble:
		encoder.AppendObjectKey("t_double")
		encoder.AppendLiteralFloat64(ov.FDouble, false)
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineOneOf2Message3 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2Message3) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*InlineOneOf2Message3) is nil")
	}
	var (
		oneOfIsFill_OneType01 bool
	)
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
		case "t_string3":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "t_double":
			if oneOfIsFill_OneType01 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType01 = true

			var ok bool
			var ot *InlineOneOf2Message3_FDouble
			if ot, ok = x.OneType01.(*InlineOneOf2Message3_FDouble); !ok {
				ot = new(InlineOneOf2Message3_FDouble)
			}
			var vv float64
			if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
				return err
			}
			ot.FDouble = vv
			x.OneType01 = ot
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message InlineOneOf2Message2 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2Message2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("t_string2")
	encoder.AppendLiteralString(x.FString1)
	switch ov := x.OneType01.(type) {
	case *InlineOneOf2Message2_FMessage1:
		encoder.AppendObjectKey("t_message2")
		if err = encoder.AppendLiteralInterface(ov.FMessage1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineOneOf2Message2 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2Message2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*InlineOneOf2Message2) is nil")
	}
	var (
		oneOfIsFill_OneType01 bool
	)
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
		case "t_string2":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "t_message2":
			if oneOfIsFill_OneType01 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType01 = true

			var ok bool
			var ot *InlineOneOf2Message2_FMessage1
			if ot, ok = x.OneType01.(*InlineOneOf2Message2_FMessage1); !ok {
				ot = new(InlineOneOf2Message2_FMessage1)
			}
			var vv *InlineOneOf2Message3
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if ot.FMessage1 != nil {
					vv = ot.FMessage1
				} else {
					vv = new(InlineOneOf2Message3)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			ot.FMessage1 = vv
			x.OneType01 = ot
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message InlineOneOf2Message1 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2Message1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("t_string1")
	encoder.AppendLiteralString(x.FString1)
	switch ov := x.OneType01.(type) {
	case *InlineOneOf2Message1_FMessage1:
		encoder.AppendObjectKey("t_message1")
		if err = encoder.AppendLiteralInterface(ov.FMessage1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineOneOf2Message1 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2Message1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*InlineOneOf2Message1) is nil")
	}
	var (
		oneOfIsFill_OneType01 bool
	)
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
		case "t_string1":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "t_message1":
			if oneOfIsFill_OneType01 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType01 = true

			var ok bool
			var ot *InlineOneOf2Message1_FMessage1
			if ot, ok = x.OneType01.(*InlineOneOf2Message1_FMessage1); !ok {
				ot = new(InlineOneOf2Message1_FMessage1)
			}
			var vv *InlineOneOf2Message2
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if ot.FMessage1 != nil {
					vv = ot.FMessage1
				} else {
					vv = new(InlineOneOf2Message2)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			ot.FMessage1 = vv
			x.OneType01 = ot
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message InlineOneOf2 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(72)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendObjectKey("t_string0")
	encoder.AppendLiteralString(x.FString1)
	switch ov := x.OneType01.(type) {
	case *InlineOneOf2_FMessage1:
		encoder.AppendObjectKey("t_message0")
		if err = encoder.AppendLiteralInterface(ov.FMessage1); err != nil {
			return nil, err
		}
	default:
		_ = ov // to avoids unused panics
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineOneOf2 in file tests/proto/cases/options/inline_oneof2.proto
func (x *InlineOneOf2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*InlineOneOf2) is nil")
	}
	var (
		oneOfIsFill_OneType01 bool
	)
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
		case "t_string0":
			var vv string
			if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
				return err
			}
			x.FString1 = vv
		case "t_message0":
			if oneOfIsFill_OneType01 {
				return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", jsonKey)
			}
			oneOfIsFill_OneType01 = true

			var ok bool
			var ot *InlineOneOf2_FMessage1
			if ot, ok = x.OneType01.(*InlineOneOf2_FMessage1); !ok {
				ot = new(InlineOneOf2_FMessage1)
			}
			var vv *InlineOneOf2Message1
			if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
				return err
			}
			if !isNULL {
				if ot.FMessage1 != nil {
					vv = ot.FMessage1
				} else {
					vv = new(InlineOneOf2Message1)
				}
				if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
					return err
				}
			}
			ot.FMessage1 = vv
			x.OneType01 = ot
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}