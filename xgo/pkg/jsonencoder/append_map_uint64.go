package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendMapU64I32 append an JSON key map of key uint64 and value int32 to JSON contents.
func AppendMapU64I32(enc *Encoder, jsonKey string, val map[uint64]int32, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValI32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64I64 append an JSON key map of key uint64 and value int64 to JSON contents.
func AppendMapU64I64(enc *Encoder, jsonKey string, val map[uint64]int64, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValI64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64U32 append an JSON key map of key uint64 and value uint32 to JSON contents.
func AppendMapU64U32(enc *Encoder, jsonKey string, val map[uint64]uint32, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValU32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64U64 append an JSON key map of key uint64 and value uint64 to JSON contents.
func AppendMapU64U64(enc *Encoder, jsonKey string, val map[uint64]uint64, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValU64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64F32 append an JSON key map of key uint64 and value float32 to JSON contents.
func AppendMapU64F32(enc *Encoder, jsonKey string, val map[uint64]float32, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValF32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64F64 append an JSON key map of key uint64 and value float64 to JSON contents.
func AppendMapU64F64(enc *Encoder, jsonKey string, val map[uint64]float64, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValF64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64Bool append an JSON key map of key uint64 and value bool to JSON contents.
func AppendMapU64Bool(enc *Encoder, jsonKey string, val map[uint64]bool, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValBool(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64Str append an JSON key map of key uint64 and value string to JSON contents.
func AppendMapU64Str(enc *Encoder, jsonKey string, val map[uint64]string, omitempty, quoteKey bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64Bytes append an JSON key map of key uint64 and value bytes to JSON contents.
func AppendMapU64Bytes(enc *Encoder, jsonKey string, val map[uint64][]byte, omitempty, quoteKey bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		if err := enc.appendValBytes(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapU64EnumNum append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapU64EnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[uint64]T, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValEnumNum(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64EnumStr append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapU64EnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[uint64]T, omitempty, quoteKey bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValEnumStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64Message append an JSON key map of key uint64 and value proto MESSAGE to JSON contents.
func AppendMapU64Message[T protoreflect.ProtoMessage](enc *Encoder, jsonKey string, val map[uint64]T, omitempty, quoteKey bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		if err := enc.appendValMessage(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapU64WKTAnyObject append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapU64WKTAnyObject(enc *Encoder, jsonKey string, val map[uint64]*anypb.Any, omitempty, quoteKey bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		if err := enc.appendValWKTAnyObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapU64WKTAnyProto append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapU64WKTAnyProto(enc *Encoder, jsonKey string, val map[uint64]*anypb.Any, omitempty, quoteKey bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		if err := enc.appendValWKTAnyProto(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapU64WKTDurObject append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurObject(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		if err := enc.appendValWKTDurObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapU64WKTDurTimeStr append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurTimeStr(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurTimeStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTDurNano append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurNano(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTDurMicro append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurMicro(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTDurMilli append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurMilli(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTDurSecond append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurSecond(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurSecond(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTDurMinute append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurMinute(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurMinute(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTDurHour append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapU64WKTDurHour(enc *Encoder, jsonKey string, val map[uint64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTDurHour(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTTsObject append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapU64WKTTsObject(enc *Encoder, jsonKey string, val map[uint64]*timestamppb.Timestamp, omitempty, quoteKey bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		if err := enc.appendValWKTTsObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapU64WKTTsLayout append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapU64WKTTsLayout(enc *Encoder, jsonKey string, val map[uint64]*timestamppb.Timestamp, omitempty, quoteKey bool, layout string) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTTsLayout(mv, layout)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTTsUnixNano append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapU64WKTTsUnixNano(enc *Encoder, jsonKey string, val map[uint64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTTsUnixNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTTsUnixMicro append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapU64WKTTsUnixMicro(enc *Encoder, jsonKey string, val map[uint64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTTsUnixMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTTsUnixMilli append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapU64WKTTsUnixMilli(enc *Encoder, jsonKey string, val map[uint64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTTsUnixMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapU64WKTTsUnixSec append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapU64WKTTsUnixSec(enc *Encoder, jsonKey string, val map[uint64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendObjectBegin()
	for mk, mv := range val {
		enc.appendMapKeyU64(mk, quoteKey)
		enc.appendValWKTTsUnixSec(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}
