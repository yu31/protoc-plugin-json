// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/inline/inline_message24.proto

package pbinline

import (
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	_ "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
)

// MarshalJSON implements interface json.Marshaler for proto message MessageLevel24 in file tests/proto/cases/inline/inline_message24.proto
func (x *MessageLevel24) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(2864)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "level24_c_f_string1", x.Level24FString1, false)
	jsonencoder.AppendValStr(enc, "level24_f_string2", x.Level24FString2, false)
	jsonencoder.AppendPtrStr(enc, "level24_c_p_string1", x.Level24PString1, false)
	jsonencoder.AppendPtrStr(enc, "level24_p_string2", x.Level24PString2, false)
	jsonencoder.AppendListStr(enc, "level24_c_r_string1", x.Level24RString1, false)
	jsonencoder.AppendListStr(enc, "level24_r_string2", x.Level24RString2, false)
	jsonencoder.AppendMapStrStr(enc, "level24_c_m_string1", x.Level24MString1, false)
	jsonencoder.AppendMapStrStr(enc, "level24_m_string2", x.Level24MString2, false)
	if err := jsonencoder.AppendValMessage(enc, "level24_f_message_extern1", x.Level24FMessageExtern1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "level24_f_message_extern2", x.Level24FMessageExtern2, false); err != nil {
		return nil, err
	}
	enc.AppendObjectKey("level24_one1_extern1")
	if x.Level24OneOfExtern1 != nil {
		enc.AppendObjectBegin()
		switch o1_11 := x.Level24OneOfExtern1.(type) {
		case *MessageLevel24_Level24One1String1:
			jsonencoder.AppendValStr(enc, "level24_c_one1_string1", o1_11.Level24One1String1, false)
		case *MessageLevel24_Level24One1MessageExtern1:
			if err := jsonencoder.AppendValMessage(enc, "level24_one1_message_extern1", o1_11.Level24One1MessageExtern1, false); err != nil {
				return nil, err
			}
		case *MessageLevel24_Level24One1MessageExtern2:
			if err := jsonencoder.AppendValMessage(enc, "level24_one1_message_extern2", o1_11.Level24One1MessageExtern2, false); err != nil {
				return nil, err
			}
		case *MessageLevel24_Level24One1MessageInlineEmtpy:
			if err := jsonencoder.AppendValMessage(enc, "level24_one1_message_inline_emtpy", o1_11.Level24One1MessageInlineEmtpy, false); err != nil {
				return nil, err
			}
		case *MessageLevel24_Level24One1MessageInlineIgnoreFields:
			if err := jsonencoder.AppendValMessage(enc, "level24_one1_message_inline_ignore_fields", o1_11.Level24One1MessageInlineIgnoreFields, false); err != nil {
				return nil, err
			}
		}
		enc.AppendObjectEnd()
	} else {
		enc.AppendValNULL()
	}
	switch o1_17 := x.Level24OneOfInline1.(type) {
	case *MessageLevel24_Level24One2String1:
		jsonencoder.AppendValStr(enc, "level24_c_one2_string1", o1_17.Level24One2String1, false)
	case *MessageLevel24_Level24One2MessageExtern1:
		if err := jsonencoder.AppendValMessage(enc, "level24_one2_message_extern1", o1_17.Level24One2MessageExtern1, false); err != nil {
			return nil, err
		}
	case *MessageLevel24_Level24One2MessageExtern2:
		if err := jsonencoder.AppendValMessage(enc, "level24_one2_message_extern2", o1_17.Level24One2MessageExtern2, false); err != nil {
			return nil, err
		}
	case *MessageLevel24_Level24One2MessageInlineEmtpy:
		if err := jsonencoder.AppendValMessage(enc, "level24_one2_message_inline_emtpy", o1_17.Level24One2MessageInlineEmtpy, false); err != nil {
			return nil, err
		}
	case *MessageLevel24_Level24One2MessageInlineIgnoreFields:
		if err := jsonencoder.AppendValMessage(enc, "level24_one2_message_inline_ignore_fields", o1_17.Level24One2MessageInlineIgnoreFields, false); err != nil {
			return nil, err
		}
	}
	switch o1_23 := x.Level24OneOfOmitempty1.(type) {
	case *MessageLevel24_Level24One3String1:
		jsonencoder.AppendValStr(enc, "level24_one3_string1", o1_23.Level24One3String1, true)
	case *MessageLevel24_Level24One3Int32A:
		jsonencoder.AppendValI32(enc, "level24_one3_int32a", o1_23.Level24One3Int32A, true, false)
	}
	if err := jsonencoder.AppendValMessage(enc, "level24_f_message_inline_emtpy", x.Level24FMessageInlineEmtpy, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendValMessage(enc, "level24_f_message_inline_ignore_fields", x.Level24FMessageInlineIgnoreFields, false); err != nil {
		return nil, err
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message MessageLevel24 in file tests/proto/cases/inline/inline_message24.proto
func (x *MessageLevel24) UnmarshalJSON(b []byte) error {
	if x == nil {
		return jsondecoder.ErrStructIsNIL("xgo/tests/pb/pbinline", "MessageLevel24")
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
		isLoad_o1_11 bool
		isLoad_o1_17 bool
		isLoad_o1_23 bool
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
		case "level24_c_f_string1":
			if x.Level24FString1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "level24_f_string2":
			if x.Level24FString2, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "level24_c_p_string1":
			if x.Level24PString1, err = jsondecoder.ReadPtrStr(dec, x.Level24PString1); err != nil {
				return err
			}
		case "level24_p_string2":
			if x.Level24PString2, err = jsondecoder.ReadPtrStr(dec, x.Level24PString2); err != nil {
				return err
			}
		case "level24_c_r_string1":
			if x.Level24RString1, err = jsondecoder.ReadListStr(dec, x.Level24RString1); err != nil {
				return err
			}
		case "level24_r_string2":
			if x.Level24RString2, err = jsondecoder.ReadListStr(dec, x.Level24RString2); err != nil {
				return err
			}
		case "level24_c_m_string1":
			if x.Level24MString1, err = jsondecoder.ReadMapStrStr(dec, x.Level24MString1); err != nil {
				return err
			}
		case "level24_m_string2":
			if x.Level24MString2, err = jsondecoder.ReadMapStrStr(dec, x.Level24MString2); err != nil {
				return err
			}
		case "level24_f_message_extern1":
			if x.Level24FMessageExtern1, err = jsondecoder.ReadValMessage(dec, x.Level24FMessageExtern1); err != nil {
				return err
			}
		case "level24_f_message_extern2":
			if x.Level24FMessageExtern2, err = jsondecoder.ReadValMessage(dec, x.Level24FMessageExtern2); err != nil {
				return err
			}
		case "level24_one1_extern1":
			if isNULL, err = dec.BeforeReadObject(); err != nil {
				return err
			}
			if isNULL {
				x.Level24OneOfExtern1 = nil
			} else {
				for { // Loop to read oneof fields
					if isEnd, err = dec.BeforeReadNext(); err != nil {
						return err
					}
					if isEnd {
						break
					}
					var oneOfKey1_11 string
					if oneOfKey1_11, err = dec.ReadObjectKey(); err != nil {
						return err
					}
					switch oneOfKey1_11 { // match the oneof key
					case "level24_c_one1_string1":
						if isLoad_o1_11 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_11 = true
						o1_11, ok := x.Level24OneOfExtern1.(*MessageLevel24_Level24One1String1)
						if !ok {
							o1_11 = new(MessageLevel24_Level24One1String1)
							x.Level24OneOfExtern1 = o1_11
						}
						if o1_11.Level24One1String1, err = jsondecoder.ReadValStr(dec); err != nil {
							return err
						}
					case "level24_one1_message_extern1":
						if isLoad_o1_11 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_11 = true
						o1_11, ok := x.Level24OneOfExtern1.(*MessageLevel24_Level24One1MessageExtern1)
						if !ok {
							o1_11 = new(MessageLevel24_Level24One1MessageExtern1)
							x.Level24OneOfExtern1 = o1_11
						}
						if o1_11.Level24One1MessageExtern1, err = jsondecoder.ReadValMessage(dec, o1_11.Level24One1MessageExtern1); err != nil {
							return err
						}
					case "level24_one1_message_extern2":
						if isLoad_o1_11 {
							return jsondecoder.ErrOneOfConflict(dec)
						}
						isLoad_o1_11 = true
						o1_11, ok := x.Level24OneOfExtern1.(*MessageLevel24_Level24One1MessageExtern2)
						if !ok {
							o1_11 = new(MessageLevel24_Level24One1MessageExtern2)
							x.Level24OneOfExtern1 = o1_11
						}
						if o1_11.Level24One1MessageExtern2, err = jsondecoder.ReadValMessage(dec, o1_11.Level24One1MessageExtern2); err != nil {
							return err
						}
					default:
						if err = dec.DiscardValue(); err != nil {
							return err
						}
					}
				}
			}
		case "level24_c_one2_string1":
			if isLoad_o1_17 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_17 = true
			o1_17, ok := x.Level24OneOfInline1.(*MessageLevel24_Level24One2String1)
			if !ok {
				o1_17 = new(MessageLevel24_Level24One2String1)
				x.Level24OneOfInline1 = o1_17
			}
			if o1_17.Level24One2String1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "level24_one2_message_extern1":
			if isLoad_o1_17 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_17 = true
			o1_17, ok := x.Level24OneOfInline1.(*MessageLevel24_Level24One2MessageExtern1)
			if !ok {
				o1_17 = new(MessageLevel24_Level24One2MessageExtern1)
				x.Level24OneOfInline1 = o1_17
			}
			if o1_17.Level24One2MessageExtern1, err = jsondecoder.ReadValMessage(dec, o1_17.Level24One2MessageExtern1); err != nil {
				return err
			}
		case "level24_one2_message_extern2":
			if isLoad_o1_17 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_17 = true
			o1_17, ok := x.Level24OneOfInline1.(*MessageLevel24_Level24One2MessageExtern2)
			if !ok {
				o1_17 = new(MessageLevel24_Level24One2MessageExtern2)
				x.Level24OneOfInline1 = o1_17
			}
			if o1_17.Level24One2MessageExtern2, err = jsondecoder.ReadValMessage(dec, o1_17.Level24One2MessageExtern2); err != nil {
				return err
			}
		case "level24_one3_string1":
			if isLoad_o1_23 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_23 = true
			o1_23, ok := x.Level24OneOfOmitempty1.(*MessageLevel24_Level24One3String1)
			if !ok {
				o1_23 = new(MessageLevel24_Level24One3String1)
				x.Level24OneOfOmitempty1 = o1_23
			}
			if o1_23.Level24One3String1, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "level24_one3_int32a":
			if isLoad_o1_23 {
				return jsondecoder.ErrOneOfConflict(dec)
			}
			isLoad_o1_23 = true
			o1_23, ok := x.Level24OneOfOmitempty1.(*MessageLevel24_Level24One3Int32A)
			if !ok {
				o1_23 = new(MessageLevel24_Level24One3Int32A)
				x.Level24OneOfOmitempty1 = o1_23
			}
			if o1_23.Level24One3Int32A, err = jsondecoder.ReadValI32(dec, false); err != nil {
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
