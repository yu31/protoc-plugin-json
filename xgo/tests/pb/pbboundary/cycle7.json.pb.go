// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/boundary/cycle7.proto

package pbboundary

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message InlineMessageCycle7 in file tests/proto/cases/boundary/cycle7.proto
func (x *InlineMessageCycle7) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(592)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	if a1_3 := x.FMessage1; a1_3 != nil {
		jsonencoder.AppendValStr(enc, "f_string3", a1_3.FString3, false)
		jsonencoder.AppendValStr(enc, "f_string4", a1_3.FString4, false)
		switch o2_6 := a1_3.OneType2.(type) {
		case *InlineMessageCycle7_EmbedMessage1_FMessage2:
			if p2_7 := o2_6.FMessage2; p2_7 != nil {
				jsonencoder.AppendValStr(enc, "f_string7", p2_7.FString7, false)
				jsonencoder.AppendValStr(enc, "f_string8", p2_7.FString8, false)
				enc.AppendObjectKey("OneType3")
				if p2_7.OneType3 != nil {
					enc.AppendObjectBegin()
					switch o3_10 := p2_7.OneType3.(type) {
					case *InlineMessageCycle7_EmbedMessage2_FMessage3:
						if err := jsonencoder.AppendValMessage(enc, "f_message3", o3_10.FMessage3, false); err != nil {
							return nil, err
						}
					case *InlineMessageCycle7_EmbedMessage2_FString9:
						jsonencoder.AppendValStr(enc, "f_string9", o3_10.FString9, false)
					case *InlineMessageCycle7_EmbedMessage2_FString10:
						jsonencoder.AppendValStr(enc, "f_string10", o3_10.FString10, false)
					}
					enc.AppendObjectEnd()
				} else {
					enc.AppendValNULL()
				}
			}
		case *InlineMessageCycle7_EmbedMessage1_FString5:
			jsonencoder.AppendValStr(enc, "f_string5", o2_6.FString5, false)
		case *InlineMessageCycle7_EmbedMessage1_FString6:
			jsonencoder.AppendValStr(enc, "f_string6", o2_6.FString6, false)
		}
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message InlineMessageCycle7 in file tests/proto/cases/boundary/cycle7.proto
func (x *InlineMessageCycle7) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbboundary", "InlineMessageCycle7")
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
		isLoad_o3_10 bool
		isLoad_o2_6  bool
	)

	// declares variables for simple to reference parent field
	var (
		a1_3 *InlineMessageCycle7_EmbedMessage1
		p2_7 *InlineMessageCycle7_EmbedMessage2
	)

	// declares anonymous func to init the parent field.
	init_a1_3 := func() error {
		if a1_3 == nil {
			if x.FMessage1 == nil {
				x.FMessage1 = new(InlineMessageCycle7_EmbedMessage1)
			}
			a1_3 = x.FMessage1
		}
		return nil
	}
	init_p2_7 := func() error {
		if p2_7 == nil {
			if _err := init_a1_3(); _err != nil {
				return _err
			}
			if isLoad_o2_6 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o2_6 = true
			o2_6, ok := a1_3.OneType2.(*InlineMessageCycle7_EmbedMessage1_FMessage2)
			if !ok {
				o2_6 = new(InlineMessageCycle7_EmbedMessage1_FMessage2)
				a1_3.OneType2 = o2_6
			}
			if o2_6.FMessage2 == nil {
				o2_6.FMessage2 = new(InlineMessageCycle7_EmbedMessage2)
			}
			p2_7 = o2_6.FMessage2
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
		case "f_string7":
			if err = init_p2_7(); err != nil {
				return err
			}
			if p2_7.FString7, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string8":
			if err = init_p2_7(); err != nil {
				return err
			}
			if p2_7.FString8, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "OneType3":
			if err = init_p2_7(); err != nil {
				return err
			}
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				p2_7.OneType3 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey3_10 string
					if oneOfKey3_10, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey3_10 { // match the oneof key
					case "f_message3":
						if isLoad_o3_10 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o3_10 = true
						o3_10, ok := p2_7.OneType3.(*InlineMessageCycle7_EmbedMessage2_FMessage3)
						if !ok {
							o3_10 = new(InlineMessageCycle7_EmbedMessage2_FMessage3)
							p2_7.OneType3 = o3_10
						}
						if o3_10.FMessage3, err = jsondecoder.ReadValMessage(dec, o3_10.FMessage3); err != nil {
							return err
						}
					case "f_string9":
						if isLoad_o3_10 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o3_10 = true
						o3_10, ok := p2_7.OneType3.(*InlineMessageCycle7_EmbedMessage2_FString9)
						if !ok {
							o3_10 = new(InlineMessageCycle7_EmbedMessage2_FString9)
							p2_7.OneType3 = o3_10
						}
						if o3_10.FString9, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					case "f_string10":
						if isLoad_o3_10 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o3_10 = true
						o3_10, ok := p2_7.OneType3.(*InlineMessageCycle7_EmbedMessage2_FString10)
						if !ok {
							o3_10 = new(InlineMessageCycle7_EmbedMessage2_FString10)
							p2_7.OneType3 = o3_10
						}
						if o3_10.FString10, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "f_string5":
			if err = init_a1_3(); err != nil {
				return err
			}
			if isLoad_o2_6 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o2_6 = true
			o2_6, ok := a1_3.OneType2.(*InlineMessageCycle7_EmbedMessage1_FString5)
			if !ok {
				o2_6 = new(InlineMessageCycle7_EmbedMessage1_FString5)
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
			o2_6, ok := a1_3.OneType2.(*InlineMessageCycle7_EmbedMessage1_FString6)
			if !ok {
				o2_6 = new(InlineMessageCycle7_EmbedMessage1_FString6)
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

// MarshalJSON implements interface json.Marshaler for proto message EmbedMessage1 in file tests/proto/cases/boundary/cycle7.proto
func (x *InlineMessageCycle7_EmbedMessage1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(496)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	jsonencoder.AppendValStr(enc, "f_string4", x.FString4, false)
	switch o1_3 := x.OneType2.(type) {
	case *InlineMessageCycle7_EmbedMessage1_FMessage2:
		if p1_4 := o1_3.FMessage2; p1_4 != nil {
			jsonencoder.AppendValStr(enc, "f_string7", p1_4.FString7, false)
			jsonencoder.AppendValStr(enc, "f_string8", p1_4.FString8, false)
			enc.AppendObjectKey("OneType3")
			if p1_4.OneType3 != nil {
				enc.AppendObjectBegin()
				switch o2_7 := p1_4.OneType3.(type) {
				case *InlineMessageCycle7_EmbedMessage2_FMessage3:
					if err := jsonencoder.AppendValMessage(enc, "f_message3", o2_7.FMessage3, false); err != nil {
						return nil, err
					}
				case *InlineMessageCycle7_EmbedMessage2_FString9:
					jsonencoder.AppendValStr(enc, "f_string9", o2_7.FString9, false)
				case *InlineMessageCycle7_EmbedMessage2_FString10:
					jsonencoder.AppendValStr(enc, "f_string10", o2_7.FString10, false)
				}
				enc.AppendObjectEnd()
			} else {
				enc.AppendValNULL()
			}
		}
	case *InlineMessageCycle7_EmbedMessage1_FString5:
		jsonencoder.AppendValStr(enc, "f_string5", o1_3.FString5, false)
	case *InlineMessageCycle7_EmbedMessage1_FString6:
		jsonencoder.AppendValStr(enc, "f_string6", o1_3.FString6, false)
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message EmbedMessage1 in file tests/proto/cases/boundary/cycle7.proto
func (x *InlineMessageCycle7_EmbedMessage1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbboundary", "InlineMessageCycle7_EmbedMessage1")
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
		isLoad_o2_7 bool
		isLoad_o1_3 bool
	)

	// declares variables for simple to reference parent field
	var (
		p1_4 *InlineMessageCycle7_EmbedMessage2
	)

	// declares anonymous func to init the parent field.
	init_p1_4 := func() error {
		if p1_4 == nil {
			if isLoad_o1_3 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_3 = true
			o1_3, ok := x.OneType2.(*InlineMessageCycle7_EmbedMessage1_FMessage2)
			if !ok {
				o1_3 = new(InlineMessageCycle7_EmbedMessage1_FMessage2)
				x.OneType2 = o1_3
			}
			if o1_3.FMessage2 == nil {
				o1_3.FMessage2 = new(InlineMessageCycle7_EmbedMessage2)
			}
			p1_4 = o1_3.FMessage2
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
		case "f_string3":
			if x.FString3, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string4":
			if x.FString4, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string7":
			if err = init_p1_4(); err != nil {
				return err
			}
			if p1_4.FString7, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string8":
			if err = init_p1_4(); err != nil {
				return err
			}
			if p1_4.FString8, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "OneType3":
			if err = init_p1_4(); err != nil {
				return err
			}
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				p1_4.OneType3 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey2_7 string
					if oneOfKey2_7, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey2_7 { // match the oneof key
					case "f_message3":
						if isLoad_o2_7 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o2_7 = true
						o2_7, ok := p1_4.OneType3.(*InlineMessageCycle7_EmbedMessage2_FMessage3)
						if !ok {
							o2_7 = new(InlineMessageCycle7_EmbedMessage2_FMessage3)
							p1_4.OneType3 = o2_7
						}
						if o2_7.FMessage3, err = jsondecoder.ReadValMessage(dec, o2_7.FMessage3); err != nil {
							return err
						}
					case "f_string9":
						if isLoad_o2_7 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o2_7 = true
						o2_7, ok := p1_4.OneType3.(*InlineMessageCycle7_EmbedMessage2_FString9)
						if !ok {
							o2_7 = new(InlineMessageCycle7_EmbedMessage2_FString9)
							p1_4.OneType3 = o2_7
						}
						if o2_7.FString9, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					case "f_string10":
						if isLoad_o2_7 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o2_7 = true
						o2_7, ok := p1_4.OneType3.(*InlineMessageCycle7_EmbedMessage2_FString10)
						if !ok {
							o2_7 = new(InlineMessageCycle7_EmbedMessage2_FString10)
							p1_4.OneType3 = o2_7
						}
						if o2_7.FString10, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "f_string5":
			if isLoad_o1_3 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_3 = true
			o1_3, ok := x.OneType2.(*InlineMessageCycle7_EmbedMessage1_FString5)
			if !ok {
				o1_3 = new(InlineMessageCycle7_EmbedMessage1_FString5)
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
			o1_3, ok := x.OneType2.(*InlineMessageCycle7_EmbedMessage1_FString6)
			if !ok {
				o1_3 = new(InlineMessageCycle7_EmbedMessage1_FString6)
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

// MarshalJSON implements interface json.Marshaler for proto message EmbedMessage2 in file tests/proto/cases/boundary/cycle7.proto
func (x *InlineMessageCycle7_EmbedMessage2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(304)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string7", x.FString7, false)
	jsonencoder.AppendValStr(enc, "f_string8", x.FString8, false)
	enc.AppendObjectKey("OneType3")
	if x.OneType3 != nil {
		enc.AppendObjectBegin()
		switch o1_3 := x.OneType3.(type) {
		case *InlineMessageCycle7_EmbedMessage2_FMessage3:
			if err := jsonencoder.AppendValMessage(enc, "f_message3", o1_3.FMessage3, false); err != nil {
				return nil, err
			}
		case *InlineMessageCycle7_EmbedMessage2_FString9:
			jsonencoder.AppendValStr(enc, "f_string9", o1_3.FString9, false)
		case *InlineMessageCycle7_EmbedMessage2_FString10:
			jsonencoder.AppendValStr(enc, "f_string10", o1_3.FString10, false)
		}
		enc.AppendObjectEnd()
	} else {
		enc.AppendValNULL()
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message EmbedMessage2 in file tests/proto/cases/boundary/cycle7.proto
func (x *InlineMessageCycle7_EmbedMessage2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbboundary", "InlineMessageCycle7_EmbedMessage2")
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
		case "f_string7":
			if x.FString7, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "f_string8":
			if x.FString8, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "OneType3":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.OneType3 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_3 string
					if oneOfKey1_3, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_3 { // match the oneof key
					case "f_message3":
						if isLoad_o1_3 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_3 = true
						o1_3, ok := x.OneType3.(*InlineMessageCycle7_EmbedMessage2_FMessage3)
						if !ok {
							o1_3 = new(InlineMessageCycle7_EmbedMessage2_FMessage3)
							x.OneType3 = o1_3
						}
						if o1_3.FMessage3, err = jsondecoder.ReadValMessage(dec, o1_3.FMessage3); err != nil {
							return err
						}
					case "f_string9":
						if isLoad_o1_3 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_3 = true
						o1_3, ok := x.OneType3.(*InlineMessageCycle7_EmbedMessage2_FString9)
						if !ok {
							o1_3 = new(InlineMessageCycle7_EmbedMessage2_FString9)
							x.OneType3 = o1_3
						}
						if o1_3.FString9, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					case "f_string10":
						if isLoad_o1_3 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_3 = true
						o1_3, ok := x.OneType3.(*InlineMessageCycle7_EmbedMessage2_FString10)
						if !ok {
							o1_3 = new(InlineMessageCycle7_EmbedMessage2_FString10)
							x.OneType3 = o1_3
						}
						if o1_3.FString10, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		default:
			if err = dec.DiscardValue(); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
