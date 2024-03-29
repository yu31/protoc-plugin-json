// Code generated by protoc-gen-gojson. DO NOT EDIT.
// versions:
// 		protoc-gen-gojson 0.0.1
// source: tests/proto/cases/options/type_map.proto

package pboptions

import (
	errors "errors"
	_ "github.com/yu31/protoc-plugin-json/xgo/pb/pbjson"
	jsondecoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsondecoder"
	jsonencoder "github.com/yu31/protoc-plugin-json/xgo/pkg/jsonencoder"
	pbexternal "github.com/yu31/protoc-plugin-json/xgo/tests/pb/pbexternal"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// MarshalJSON implements interface json.Marshaler for proto message TypeMap1 in file tests/proto/cases/options/type_map.proto
func (x *TypeMap1) MarshalJSON() ([]byte, error) {
	if x == nil {
		return []byte("null"), nil
	}
	var err error
	encoder := jsonencoder.New(424)

	// Add begin JSON identifier
	encoder.AppendObjectBegin()

	if len(x.FInt32) != 0 {
		encoder.AppendObjectKey("t_int32")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FInt32 {
			encoder.AppendMapKeyInt32(mk, true)
			encoder.AppendLiteralInt32(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FInt64) != 0 {
		encoder.AppendObjectKey("t_int64")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FInt64 {
			encoder.AppendMapKeyInt64(mk, true)
			encoder.AppendLiteralInt64(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FUint32) != 0 {
		encoder.AppendObjectKey("t_uint32")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FUint32 {
			encoder.AppendMapKeyUInt32(mk, true)
			encoder.AppendLiteralUint32(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FUint64) != 0 {
		encoder.AppendObjectKey("t_uint64")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FUint64 {
			encoder.AppendMapKeyUInt64(mk, true)
			encoder.AppendLiteralUint64(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FSint32) != 0 {
		encoder.AppendObjectKey("t_sint32")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FSint32 {
			encoder.AppendMapKeyInt32(mk, true)
			encoder.AppendLiteralInt32(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FSint64) != 0 {
		encoder.AppendObjectKey("t_sint64")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FSint64 {
			encoder.AppendMapKeyInt64(mk, true)
			encoder.AppendLiteralInt64(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FSfixed32) != 0 {
		encoder.AppendObjectKey("t_sfixed32")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FSfixed32 {
			encoder.AppendMapKeyInt32(mk, true)
			encoder.AppendLiteralInt32(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FSfixed64) != 0 {
		encoder.AppendObjectKey("t_sfixed64")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FSfixed64 {
			encoder.AppendMapKeyInt64(mk, true)
			encoder.AppendLiteralInt64(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FFixed32) != 0 {
		encoder.AppendObjectKey("t_fixed32")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FFixed32 {
			encoder.AppendMapKeyUInt32(mk, true)
			encoder.AppendLiteralUint32(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FFixed64) != 0 {
		encoder.AppendObjectKey("t_fixed64")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FFixed64 {
			encoder.AppendMapKeyUInt64(mk, true)
			encoder.AppendLiteralUint64(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FFloat) != 0 {
		encoder.AppendObjectKey("t_float")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FFloat {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralFloat32(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FDouble) != 0 {
		encoder.AppendObjectKey("t_double")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FDouble {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralFloat64(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FBool1) != 0 {
		encoder.AppendObjectKey("t_bool1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FBool1 {
			encoder.AppendMapKeyBool(mk, true)
			encoder.AppendLiteralBool(mv, false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FString1) != 0 {
		encoder.AppendObjectKey("t_string1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FString1 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralString(mv)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FBytes1) != 0 {
		encoder.AppendObjectKey("t_bytes1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FBytes1 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralBytes(mv)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FEnum1) != 0 {
		encoder.AppendObjectKey("t_enum1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FEnum1 {
			encoder.AppendMapKeyString(mk)
			encoder.AppendLiteralInt32(int32(mv.Number()), false)
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FMessage1) != 0 {
		encoder.AppendObjectKey("t_message1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FMessage1 {
			encoder.AppendMapKeyString(mk)
			if err = encoder.AppendLiteralInterface(mv); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FAny1) != 0 {
		encoder.AppendObjectKey("t_any1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FAny1 {
			encoder.AppendMapKeyString(mk)
			if err = encoder.AppendLiteralInterface(mv); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FDuration1) != 0 {
		encoder.AppendObjectKey("t_duration1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FDuration1 {
			encoder.AppendMapKeyString(mk)
			if err = encoder.AppendLiteralInterface(mv); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	}
	if len(x.FTimestamp1) != 0 {
		encoder.AppendObjectKey("t_timestamp1")
		encoder.AppendObjectBegin()
		for mk, mv := range x.FTimestamp1 {
			encoder.AppendMapKeyString(mk)
			if err = encoder.AppendLiteralInterface(mv); err != nil {
				return nil, err
			}
		}
		encoder.AppendObjectEnd()
	}

	// Add end JSON identifier
	encoder.AppendObjectEnd()
	return encoder.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler for proto message TypeMap1 in file tests/proto/cases/options/type_map.proto
func (x *TypeMap1) UnmarshalJSON(b []byte) error {
	if x == nil {
		return errors.New("json: Unmarshal: xgo/tests/pb/pboptions.(*TypeMap1) is nil")
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
		case "t_int32":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FInt32 = nil
				continue LOOP_SCAN
			}
			if x.FInt32 == nil {
				x.FInt32 = make(map[int32]int32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int32
				if mk, err = decoder.ReadMapKeyInt32(jsonKey, true); err != nil {
					return err
				}
				var vv int32
				if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
					return err
				}
				x.FInt32[mk] = vv
			}
		case "t_int64":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FInt64 = nil
				continue LOOP_SCAN
			}
			if x.FInt64 == nil {
				x.FInt64 = make(map[int64]int64)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int64
				if mk, err = decoder.ReadMapKeyInt64(jsonKey, true); err != nil {
					return err
				}
				var vv int64
				if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
					return err
				}
				x.FInt64[mk] = vv
			}
		case "t_uint32":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FUint32 = nil
				continue LOOP_SCAN
			}
			if x.FUint32 == nil {
				x.FUint32 = make(map[uint32]uint32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk uint32
				if mk, err = decoder.ReadMapKeyUint32(jsonKey, true); err != nil {
					return err
				}
				var vv uint32
				if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
					return err
				}
				x.FUint32[mk] = vv
			}
		case "t_uint64":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FUint64 = nil
				continue LOOP_SCAN
			}
			if x.FUint64 == nil {
				x.FUint64 = make(map[uint64]uint64)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk uint64
				if mk, err = decoder.ReadMapKeyUint64(jsonKey, true); err != nil {
					return err
				}
				var vv uint64
				if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
					return err
				}
				x.FUint64[mk] = vv
			}
		case "t_sint32":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FSint32 = nil
				continue LOOP_SCAN
			}
			if x.FSint32 == nil {
				x.FSint32 = make(map[int32]int32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int32
				if mk, err = decoder.ReadMapKeyInt32(jsonKey, true); err != nil {
					return err
				}
				var vv int32
				if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
					return err
				}
				x.FSint32[mk] = vv
			}
		case "t_sint64":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FSint64 = nil
				continue LOOP_SCAN
			}
			if x.FSint64 == nil {
				x.FSint64 = make(map[int64]int64)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int64
				if mk, err = decoder.ReadMapKeyInt64(jsonKey, true); err != nil {
					return err
				}
				var vv int64
				if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
					return err
				}
				x.FSint64[mk] = vv
			}
		case "t_sfixed32":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FSfixed32 = nil
				continue LOOP_SCAN
			}
			if x.FSfixed32 == nil {
				x.FSfixed32 = make(map[int32]int32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int32
				if mk, err = decoder.ReadMapKeyInt32(jsonKey, true); err != nil {
					return err
				}
				var vv int32
				if vv, err = decoder.ReadLiteralInt32(jsonKey, false); err != nil {
					return err
				}
				x.FSfixed32[mk] = vv
			}
		case "t_sfixed64":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FSfixed64 = nil
				continue LOOP_SCAN
			}
			if x.FSfixed64 == nil {
				x.FSfixed64 = make(map[int64]int64)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk int64
				if mk, err = decoder.ReadMapKeyInt64(jsonKey, true); err != nil {
					return err
				}
				var vv int64
				if vv, err = decoder.ReadLiteralInt64(jsonKey, false); err != nil {
					return err
				}
				x.FSfixed64[mk] = vv
			}
		case "t_fixed32":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FFixed32 = nil
				continue LOOP_SCAN
			}
			if x.FFixed32 == nil {
				x.FFixed32 = make(map[uint32]uint32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk uint32
				if mk, err = decoder.ReadMapKeyUint32(jsonKey, true); err != nil {
					return err
				}
				var vv uint32
				if vv, err = decoder.ReadLiteralUint32(jsonKey, false); err != nil {
					return err
				}
				x.FFixed32[mk] = vv
			}
		case "t_fixed64":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FFixed64 = nil
				continue LOOP_SCAN
			}
			if x.FFixed64 == nil {
				x.FFixed64 = make(map[uint64]uint64)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk uint64
				if mk, err = decoder.ReadMapKeyUint64(jsonKey, true); err != nil {
					return err
				}
				var vv uint64
				if vv, err = decoder.ReadLiteralUint64(jsonKey, false); err != nil {
					return err
				}
				x.FFixed64[mk] = vv
			}
		case "t_float":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FFloat = nil
				continue LOOP_SCAN
			}
			if x.FFloat == nil {
				x.FFloat = make(map[string]float32)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv float32
				if vv, err = decoder.ReadLiteralFloat32(jsonKey, false); err != nil {
					return err
				}
				x.FFloat[mk] = vv
			}
		case "t_double":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FDouble = nil
				continue LOOP_SCAN
			}
			if x.FDouble == nil {
				x.FDouble = make(map[string]float64)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv float64
				if vv, err = decoder.ReadLiteralFloat64(jsonKey, false); err != nil {
					return err
				}
				x.FDouble[mk] = vv
			}
		case "t_bool1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FBool1 = nil
				continue LOOP_SCAN
			}
			if x.FBool1 == nil {
				x.FBool1 = make(map[bool]bool)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk bool
				if mk, err = decoder.ReadMapKeyBool(jsonKey, true); err != nil {
					return err
				}
				var vv bool
				if vv, err = decoder.ReadLiteralBool(jsonKey, false); err != nil {
					return err
				}
				x.FBool1[mk] = vv
			}
		case "t_string1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FString1 = nil
				continue LOOP_SCAN
			}
			if x.FString1 == nil {
				x.FString1 = make(map[string]string)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv string
				if vv, err = decoder.ReadLiteralString(jsonKey); err != nil {
					return err
				}
				x.FString1[mk] = vv
			}
		case "t_bytes1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FBytes1 = nil
				continue LOOP_SCAN
			}
			if x.FBytes1 == nil {
				x.FBytes1 = make(map[string][]byte)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv []byte
				if vv, err = decoder.ReadLiteralBytes(jsonKey); err != nil {
					return err
				}
				x.FBytes1[mk] = vv
			}
		case "t_enum1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FEnum1 = nil
				continue LOOP_SCAN
			}
			if x.FEnum1 == nil {
				x.FEnum1 = make(map[string]pbexternal.Enum1)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv pbexternal.Enum1
				var v1 int32
				if v1, err = decoder.ReadLiteralEnumNumber(jsonKey, false); err != nil {
					return err
				}
				vv = pbexternal.Enum1(v1)
				x.FEnum1[mk] = vv
			}
		case "t_message1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FMessage1 = nil
				continue LOOP_SCAN
			}
			if x.FMessage1 == nil {
				x.FMessage1 = make(map[string]*pbexternal.Message1)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv *pbexternal.Message1
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.FMessage1[mk] != nil {
						vv = x.FMessage1[mk]
					} else {
						vv = new(pbexternal.Message1)
					}
					if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
						return err
					}
				}
				x.FMessage1[mk] = vv
			}
		case "t_any1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FAny1 = nil
				continue LOOP_SCAN
			}
			if x.FAny1 == nil {
				x.FAny1 = make(map[string]*anypb.Any)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv *anypb.Any
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.FAny1[mk] != nil {
						vv = x.FAny1[mk]
					} else {
						vv = new(anypb.Any)
					}
					if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
						return err
					}
				}
				x.FAny1[mk] = vv
			}
		case "t_duration1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FDuration1 = nil
				continue LOOP_SCAN
			}
			if x.FDuration1 == nil {
				x.FDuration1 = make(map[string]*durationpb.Duration)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv *durationpb.Duration
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.FDuration1[mk] != nil {
						vv = x.FDuration1[mk]
					} else {
						vv = new(durationpb.Duration)
					}
					if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
						return err
					}
				}
				x.FDuration1[mk] = vv
			}
		case "t_timestamp1":
			if isNULL, err = decoder.BeforeReadMap(jsonKey); err != nil {
				return err
			}
			if isNULL {
				x.FTimestamp1 = nil
				continue LOOP_SCAN
			}
			if x.FTimestamp1 == nil {
				x.FTimestamp1 = make(map[string]*timestamppb.Timestamp)
			}
			for {
				if isEnd, err = decoder.BeforeReadNext(jsonKey); err != nil {
					return err
				}
				if isEnd {
					break
				}
				var mk string
				if mk, err = decoder.ReadMapKeyString(jsonKey); err != nil {
					return err
				}
				var vv *timestamppb.Timestamp
				if isNULL, err = decoder.NextLiteralIsNULL(jsonKey); err != nil {
					return err
				}
				if !isNULL {
					if x.FTimestamp1[mk] != nil {
						vv = x.FTimestamp1[mk]
					} else {
						vv = new(timestamppb.Timestamp)
					}
					if err = decoder.ReadLiteralInterface(jsonKey, vv); err != nil {
						return err
					}
				}
				x.FTimestamp1[mk] = vv
			}
		default:
			if err = decoder.DiscardValue(jsonKey); err != nil {
				return err
			}
		} // end switch
	}
	return nil
}
