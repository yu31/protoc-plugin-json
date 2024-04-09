// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/boundary/boundary1.proto

package pbboundary

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
)

// MarshalJSON implements interface json.Marshaler for proto message Message1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Message1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(72)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string1", x.FString1, false)
	jsonencoder.AppendValStr(enc, "f_string2", x.FString2, false)
	jsonencoder.AppendValStr(enc, "f_string3", x.FString3, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Message1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Message1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Message1) is nil")
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

// MarshalJSON implements interface json.Marshaler for proto message Repeated1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Repeated1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(32)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendListStr(enc, "r_string1", x.RString1, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Repeated1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Repeated1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Repeated1) is nil")
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
		case "r_string1":
			if x.RString1, err = jsondecoder.ReadListStr(dec, x.RString1); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Map1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Map1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(32)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendMapStrStr(enc, "m_string1", x.MString1, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Map1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Map1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Map1) is nil")
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
		case "m_string1":
			if x.MString1, err = jsondecoder.ReadMapStrStr(dec, x.MString1); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Complex1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(104)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValI32(enc, "f_int32", x.FInt32, false, false)
	jsonencoder.AppendListI64(enc, "r_int64", x.RInt64, false, false)
	if err := jsonencoder.AppendValMessage(enc, "f_message1", x.FMessage1, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendMapStrI32(enc, "m_int32", x.MInt32, false, false)
	jsonencoder.AppendPtrI64(enc, "f_int64", x.FInt64, false, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Complex1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex1) is nil")
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
		case "f_int32":
			if x.FInt32, err = jsondecoder.ReadValI32(dec, false); err != nil {
				return err
			}
		case "r_int64":
			if x.RInt64, err = jsondecoder.ReadListI64(dec, x.RInt64, false); err != nil {
				return err
			}
		case "f_message1":
			if x.FMessage1, err = jsondecoder.ReadValMessage(dec, x.FMessage1); err != nil {
				return err
			}
		case "m_int32":
			if x.MInt32, err = jsondecoder.ReadMapStrI32(dec, x.MInt32, false); err != nil {
				return err
			}
		case "f_int64":
			if x.FInt64, err = jsondecoder.ReadPtrI64(dec, false); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Complex2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(104)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string", x.FString, false)
	jsonencoder.AppendListI64(enc, "r_int64", x.RInt64, false, false)
	if err := jsonencoder.AppendValMessage(enc, "level1", x.Level1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendListMessage(enc, "r_level1", x.RLevel1, false); err != nil {
		return nil, err
	}
	if err := jsonencoder.AppendListMessage(enc, "r_level2", x.RLevel2, false); err != nil {
		return nil, err
	}
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Complex2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2) is nil")
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
		case "f_string":
			if x.FString, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "r_int64":
			if x.RInt64, err = jsondecoder.ReadListI64(dec, x.RInt64, false); err != nil {
				return err
			}
		case "level1":
			if x.Level1, err = jsondecoder.ReadValMessage(dec, x.Level1); err != nil {
				return err
			}
		case "r_level1":
			if x.RLevel1, err = jsondecoder.ReadListMessage(dec, x.RLevel1); err != nil {
				return err
			}
		case "r_level2":
			if x.RLevel2, err = jsondecoder.ReadListMessage(dec, x.RLevel2); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Level1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(64)
	enc.AppendObjectBegin() // Add begin JSON identifier

	if err := jsonencoder.AppendValMessage(enc, "level2", x.Level2, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendValStr(enc, "f_string", x.FString, false)
	jsonencoder.AppendListStr(enc, "r_string", x.RString, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Level1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2_Level1) is nil")
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
		case "level2":
			if x.Level2, err = jsondecoder.ReadValMessage(dec, x.Level2); err != nil {
				return err
			}
		case "f_string":
			if x.FString, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "r_string":
			if x.RString, err = jsondecoder.ReadListStr(dec, x.RString); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Level2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(64)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValStr(enc, "f_string", x.FString, false)
	if err := jsonencoder.AppendValMessage(enc, "level3", x.Level3, false); err != nil {
		return nil, err
	}
	jsonencoder.AppendListI64(enc, "r_int64", x.RInt64, false, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Level2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2_Level2) is nil")
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
		case "f_string":
			if x.FString, err = jsondecoder.ReadValStr(dec); err != nil {
				return err
			}
		case "level3":
			if x.Level3, err = jsondecoder.ReadValMessage(dec, x.Level3); err != nil {
				return err
			}
		case "r_int64":
			if x.RInt64, err = jsondecoder.ReadListI64(dec, x.RInt64, false); err != nil {
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

// MarshalJSON implements interface json.Marshaler for proto message Level3 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level3) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	enc := jsonencoder.New(80)
	enc.AppendObjectBegin() // Add begin JSON identifier

	jsonencoder.AppendValI32(enc, "f_int32", x.FInt32, false, false)
	jsonencoder.AppendListI64(enc, "r_int64", x.RInt64, false, false)
	jsonencoder.AppendMapStrI32(enc, "m_int32", x.MInt32, false, false)
	jsonencoder.AppendPtrI64(enc, "p_int64", x.PInt64, false, false)
	enc.AppendObjectEnd() // Add end JSON identifier
	return enc.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler for proto message Level3 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level3) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2_Level3) is nil")
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
		case "f_int32":
			if x.FInt32, err = jsondecoder.ReadValI32(dec, false); err != nil {
				return err
			}
		case "r_int64":
			if x.RInt64, err = jsondecoder.ReadListI64(dec, x.RInt64, false); err != nil {
				return err
			}
		case "m_int32":
			if x.MInt32, err = jsondecoder.ReadMapStrI32(dec, x.MInt32, false); err != nil {
				return err
			}
		case "p_int64":
			if x.PInt64, err = jsondecoder.ReadPtrI64(dec, false); err != nil {
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
