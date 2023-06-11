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
	var err error
	encoder := jsonencoder.New(68)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string1")
	encoder.AppendValueString(x.FString1)
	encoder.AppendJSONKey("f_string2")
	encoder.AppendValueString(x.FString2)
	encoder.AppendJSONKey("f_string3")
	encoder.AppendValueString(x.FString3)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Message1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Message1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Message1) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string1":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString1 = vv
		case "f_string2":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString2 = vv
		case "f_string3":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString3 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Repeated1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Repeated1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(24)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("r_string1")
	if x.RString1 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RString1 {
			encoder.AppendValueString(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Repeated1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Repeated1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Repeated1) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "r_string1":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RString1 = nil
				continue LOOP_SCAN
			}
			if x.RString1 == nil {
				x.RString1 = make([]string, 0)
			}
			i := 0
			length := len(x.RString1)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv string
				vv, err = decoder.ReadArrayElemString(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RString1[i] = vv
				} else {
					x.RString1 = append(x.RString1, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RString1 = x.RString1[:i]
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Map1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Map1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(24)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("m_string1")
	if x.MString1 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MString1 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendValueString(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Map1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Map1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Map1) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "m_string1":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MString1 = nil
				continue LOOP_SCAN
			}
			if x.MString1 == nil {
				x.MString1 = make(map[string]string)
			}
			for {
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mapKey string
				mapKey, err = decoder.ReadMapKeyString(jsonKey)
				if err != nil {
					return err
				}
				var mapVal string
				mapVal, err = decoder.ReadMapValueString(jsonKey)
				if err != nil {
					return err
				}
				x.MString1[mapKey] = mapVal
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Complex1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(100)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_int32")
	encoder.AppendValueInt32(x.FInt32)
	encoder.AppendJSONKey("r_int64")
	if x.RInt64 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt64 {
			encoder.AppendValueInt64(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_message1")
	err = encoder.AppendValueInterface(x.FMessage1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("m_int32")
	if x.MInt32 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MInt32 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendValueInt32(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("f_int64")
	encoder.AppendPointerInt64(x.FInt64)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Complex1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex1) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_int32":
			var vv int32
			vv, err = decoder.ReadValueInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FInt32 = vv
		case "r_int64":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RInt64 = nil
				continue LOOP_SCAN
			}
			if x.RInt64 == nil {
				x.RInt64 = make([]int64, 0)
			}
			i := 0
			length := len(x.RInt64)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv int64
				vv, err = decoder.ReadArrayElemInt64(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RInt64[i] = vv
				} else {
					x.RInt64 = append(x.RInt64, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RInt64 = x.RInt64[:i]
			}
		case "f_message1":
			var vv *Message1
			initFN := func() interface{} {
				if x.FMessage1 != nil {
					vv = x.FMessage1
				} else {
					vv = new(Message1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.FMessage1 = vv
		case "m_int32":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MInt32 = nil
				continue LOOP_SCAN
			}
			if x.MInt32 == nil {
				x.MInt32 = make(map[string]int32)
			}
			for {
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mapKey string
				mapKey, err = decoder.ReadMapKeyString(jsonKey)
				if err != nil {
					return err
				}
				var mapVal int32
				mapVal, err = decoder.ReadMapValueInt32(jsonKey)
				if err != nil {
					return err
				}
				x.MInt32[mapKey] = mapVal
			}
		case "f_int64":
			var vv *int64
			vv, err = decoder.ReadPointerInt64(jsonKey)
			if err != nil {
				return err
			}
			x.FInt64 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Complex2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(96)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string")
	encoder.AppendValueString(x.FString)
	encoder.AppendJSONKey("r_int64")
	if x.RInt64 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt64 {
			encoder.AppendValueInt64(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("level1")
	err = encoder.AppendValueInterface(x.Level1)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("r_level1")
	if x.RLevel1 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RLevel1 {
			err = encoder.AppendValueInterface(ri)
			if err != nil {
				return nil, err
			}
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("r_level2")
	if x.RLevel2 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RLevel2 {
			err = encoder.AppendValueInterface(ri)
			if err != nil {
				return nil, err
			}
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Complex2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString = vv
		case "r_int64":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RInt64 = nil
				continue LOOP_SCAN
			}
			if x.RInt64 == nil {
				x.RInt64 = make([]int64, 0)
			}
			i := 0
			length := len(x.RInt64)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv int64
				vv, err = decoder.ReadArrayElemInt64(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RInt64[i] = vv
				} else {
					x.RInt64 = append(x.RInt64, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RInt64 = x.RInt64[:i]
			}
		case "level1":
			var vv *Complex2_Level1
			initFN := func() interface{} {
				if x.Level1 != nil {
					vv = x.Level1
				} else {
					vv = new(Complex2_Level1)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.Level1 = vv
		case "r_level1":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RLevel1 = nil
				continue LOOP_SCAN
			}
			if x.RLevel1 == nil {
				x.RLevel1 = make([]*Complex2_Level1, 0)
			}
			i := 0
			length := len(x.RLevel1)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv *Complex2_Level1
				initFN := func() interface{} {
					if i < length {
						vv = x.RLevel1[i]
					}
					if vv == nil {
						vv = new(Complex2_Level1)
					}
					return vv
				}
				err = decoder.ReadArrayElemInterface(jsonKey, initFN)
				if err != nil {
					return err
				}
				if i < length {
					x.RLevel1[i] = vv
				} else {
					x.RLevel1 = append(x.RLevel1, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RLevel1 = x.RLevel1[:i]
			}
		case "r_level2":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RLevel2 = nil
				continue LOOP_SCAN
			}
			if x.RLevel2 == nil {
				x.RLevel2 = make([]*Complex2_Level1, 0)
			}
			i := 0
			length := len(x.RLevel2)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv *Complex2_Level1
				initFN := func() interface{} {
					if i < length {
						vv = x.RLevel2[i]
					}
					if vv == nil {
						vv = new(Complex2_Level1)
					}
					return vv
				}
				err = decoder.ReadArrayElemInterface(jsonKey, initFN)
				if err != nil {
					return err
				}
				if i < length {
					x.RLevel2[i] = vv
				} else {
					x.RLevel2 = append(x.RLevel2, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RLevel2 = x.RLevel2[:i]
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Level1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(58)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("level2")
	err = encoder.AppendValueInterface(x.Level2)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("f_string")
	encoder.AppendValueString(x.FString)
	encoder.AppendJSONKey("r_string")
	if x.RString != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RString {
			encoder.AppendValueString(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Level1 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2_Level1) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "level2":
			var vv *Complex2_Level2
			initFN := func() interface{} {
				if x.Level2 != nil {
					vv = x.Level2
				} else {
					vv = new(Complex2_Level2)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.Level2 = vv
		case "f_string":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString = vv
		case "r_string":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RString = nil
				continue LOOP_SCAN
			}
			if x.RString == nil {
				x.RString = make([]string, 0)
			}
			i := 0
			length := len(x.RString)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv string
				vv, err = decoder.ReadArrayElemString(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RString[i] = vv
				} else {
					x.RString = append(x.RString, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RString = x.RString[:i]
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Level2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level2) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(56)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_string")
	encoder.AppendValueString(x.FString)
	encoder.AppendJSONKey("level3")
	err = encoder.AppendValueInterface(x.Level3)
	if err != nil {
		return nil, err
	}
	encoder.AppendJSONKey("r_int64")
	if x.RInt64 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt64 {
			encoder.AppendValueInt64(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Level2 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level2) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2_Level2) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_string":
			var vv string
			vv, err = decoder.ReadValueString(jsonKey)
			if err != nil {
				return err
			}
			x.FString = vv
		case "level3":
			var vv *Complex2_Level3
			initFN := func() interface{} {
				if x.Level3 != nil {
					vv = x.Level3
				} else {
					vv = new(Complex2_Level3)
				}
				return vv
			}
			err = decoder.ReadValueInterface(jsonKey, initFN)
			if err != nil {
				return err
			}
			x.Level3 = vv
		case "r_int64":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RInt64 = nil
				continue LOOP_SCAN
			}
			if x.RInt64 == nil {
				x.RInt64 = make([]int64, 0)
			}
			i := 0
			length := len(x.RInt64)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv int64
				vv, err = decoder.ReadArrayElemInt64(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RInt64[i] = vv
				} else {
					x.RInt64 = append(x.RInt64, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RInt64 = x.RInt64[:i]
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON implements interface json.Marshaler for proto message Level3 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level3) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(76)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	encoder.AppendJSONKey("f_int32")
	encoder.AppendValueInt32(x.FInt32)
	encoder.AppendJSONKey("r_int64")
	if x.RInt64 != nil {
		encoder.AppendArrayBegin()
		for _, ri := range x.RInt64 {
			encoder.AppendValueInt64(ri)
		}
		encoder.AppendArrayEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("m_int32")
	if x.MInt32 != nil {
		encoder.AppendObjectBegin()
		for mk, mv := range x.MInt32 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendValueInt32(mv)
		}
		encoder.AppendObjectEnd()
	} else {
		encoder.AppendValueNULL()
	}
	encoder.AppendJSONKey("p_int64")
	encoder.AppendPointerInt64(x.PInt64)

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message Level3 in file tests/proto/cases/boundary/boundary1.proto
func (x *Complex2_Level3) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pbboundary.(*Complex2_Level3) is nil")
	}

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
		if isEnd, err = decoder.BeforeReadJSONKey(); err != nil {
			return err
		}
		if isEnd {
			break LOOP_SCAN
		}
		if jsonKey, err = decoder.ReadJSONKey(); err != nil {
			return err
		}
		switch jsonKey { // match the JSON key
		case "f_int32":
			var vv int32
			vv, err = decoder.ReadValueInt32(jsonKey)
			if err != nil {
				return err
			}
			x.FInt32 = vv
		case "r_int64":
			if isNULL, err = decoder.BeforeReadArray(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.RInt64 = nil
				continue LOOP_SCAN
			}
			if x.RInt64 == nil {
				x.RInt64 = make([]int64, 0)
			}
			i := 0
			length := len(x.RInt64)
			for {
				if isEnd, err = decoder.BeforeReadArrayElem(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var vv int64
				vv, err = decoder.ReadArrayElemInt64(jsonKey)
				if err != nil {
					return err
				}
				if i < length {
					x.RInt64[i] = vv
				} else {
					x.RInt64 = append(x.RInt64, vv)
				}
				i++
			}
			if i < length {
				// truncate the slice to Consistent with standard library json.
				x.RInt64 = x.RInt64[:i]
			}
		case "m_int32":
			if isNULL, err = decoder.BeforeReadObject(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.MInt32 = nil
				continue LOOP_SCAN
			}
			if x.MInt32 == nil {
				x.MInt32 = make(map[string]int32)
			}
			for {
				if isEnd, err = decoder.BeforeReadObjectKey(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mapKey string
				mapKey, err = decoder.ReadMapKeyString(jsonKey)
				if err != nil {
					return err
				}
				var mapVal int32
				mapVal, err = decoder.ReadMapValueInt32(jsonKey)
				if err != nil {
					return err
				}
				x.MInt32[mapKey] = mapVal
			}
		case "p_int64":
			var vv *int64
			vv, err = decoder.ReadPointerInt64(jsonKey)
			if err != nil {
				return err
			}
			x.PInt64 = vv
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		}
	}
	return nil
}
