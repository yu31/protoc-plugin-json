package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendMapI64I32 append an JSON key map of key int64 and value int32 to JSON contents.
func AppendMapI64I32(enc *Encoder, jsonKey string, val map[int64]int32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValI32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64I64 append an JSON key map of key int64 and value int64 to JSON contents.
func AppendMapI64I64(enc *Encoder, jsonKey string, val map[int64]int64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValI64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64U32 append an JSON key map of key int64 and value uint32 to JSON contents.
func AppendMapI64U32(enc *Encoder, jsonKey string, val map[int64]uint32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValU32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64U64 append an JSON key map of key int64 and value uint64 to JSON contents.
func AppendMapI64U64(enc *Encoder, jsonKey string, val map[int64]uint64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValU64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64F32 append an JSON key map of key int64 and value float32 to JSON contents.
func AppendMapI64F32(enc *Encoder, jsonKey string, val map[int64]float32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValF32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64F64 append an JSON key map of key int64 and value float64 to JSON contents.
func AppendMapI64F64(enc *Encoder, jsonKey string, val map[int64]float64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValF64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64Bool append an JSON key map of key int64 and value bool to JSON contents.
func AppendMapI64Bool(enc *Encoder, jsonKey string, val map[int64]bool, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValBool(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64Str append an JSON key map of key int64 and value string to JSON contents.
func AppendMapI64Str(enc *Encoder, jsonKey string, val map[int64]string, omitempty, quoteKey bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64Bytes append an JSON key map of key int64 and value bytes to JSON contents.
func AppendMapI64Bytes(enc *Encoder, jsonKey string, val map[int64][]byte, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI64(mk, quoteKey)
		if err := enc.appendValBytes(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI64EnumNum append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapI64EnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[int64]T, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValEnumNum(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64EnumStr append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapI64EnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[int64]T, omitempty, quoteKey bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValEnumStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64Message append an JSON key map of key int64 and value proto MESSAGE to JSON contents.
func AppendMapI64Message[T protoreflect.ProtoMessage](enc *Encoder, jsonKey string, val map[int64]T, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI64(mk, quoteKey)
		if err := enc.appendValMessage(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI64WKTAnyObject append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapI64WKTAnyObject(enc *Encoder, jsonKey string, val map[int64]*anypb.Any, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI64(mk, quoteKey)
		if err := enc.appendValWKTAnyObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI64WKTAnyProto append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapI64WKTAnyProto(enc *Encoder, jsonKey string, val map[int64]*anypb.Any, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI64(mk, quoteKey)
		if err := enc.appendValWKTAnyProto(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI64WKTDurObject append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurObject(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI64(mk, quoteKey)
		if err := enc.appendValWKTDurObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI64WKTDurTimeStr append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurTimeStr(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurTimeStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTDurNano append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurNano(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTDurMicro append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurMicro(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTDurMilli append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurMilli(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTDurSecond append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurSecond(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurSecond(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTDurMinute append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurMinute(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurMinute(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTDurHour append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI64WKTDurHour(enc *Encoder, jsonKey string, val map[int64]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTDurHour(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTTsObject append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI64WKTTsObject(enc *Encoder, jsonKey string, val map[int64]*timestamppb.Timestamp, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI64(mk, quoteKey)
		if err := enc.appendValWKTTsObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI64WKTTsLayout append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI64WKTTsLayout(enc *Encoder, jsonKey string, val map[int64]*timestamppb.Timestamp, omitempty, quoteKey bool, layout string) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTTsLayout(mv, layout)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTTsUnixNano append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI64WKTTsUnixNano(enc *Encoder, jsonKey string, val map[int64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTTsUnixNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTTsUnixMicro append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI64WKTTsUnixMicro(enc *Encoder, jsonKey string, val map[int64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTTsUnixMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTTsUnixMilli append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI64WKTTsUnixMilli(enc *Encoder, jsonKey string, val map[int64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTTsUnixMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI64WKTTsUnixSec append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI64WKTTsUnixSec(enc *Encoder, jsonKey string, val map[int64]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI64(mk, quoteKey)
		enc.appendValWKTTsUnixSec(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}
