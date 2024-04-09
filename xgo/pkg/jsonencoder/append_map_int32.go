package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendMapI32I32 append an JSON key map of key int32 and value int32 to JSON contents.
func AppendMapI32I32(enc *Encoder, jsonKey string, val map[int32]int32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValI32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32I64 append an JSON key map of key int32 and value int64 to JSON contents.
func AppendMapI32I64(enc *Encoder, jsonKey string, val map[int32]int64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValI64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32U32 append an JSON key map of key int32 and value uint32 to JSON contents.
func AppendMapI32U32(enc *Encoder, jsonKey string, val map[int32]uint32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValU32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32U64 append an JSON key map of key int32 and value uint64 to JSON contents.
func AppendMapI32U64(enc *Encoder, jsonKey string, val map[int32]uint64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValU64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32F32 append an JSON key map of key int32 and value float32 to JSON contents.
func AppendMapI32F32(enc *Encoder, jsonKey string, val map[int32]float32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValF32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32F64 append an JSON key map of key int32 and value float64 to JSON contents.
func AppendMapI32F64(enc *Encoder, jsonKey string, val map[int32]float64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValF64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32Bool append an JSON key map of key int32 and value bool to JSON contents.
func AppendMapI32Bool(enc *Encoder, jsonKey string, val map[int32]bool, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValBool(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32Str append an JSON key map of key int32 and value string to JSON contents.
func AppendMapI32Str(enc *Encoder, jsonKey string, val map[int32]string, omitempty, quoteKey bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32Bytes append an JSON key map of key int32 and value bytes to JSON contents.
func AppendMapI32Bytes(enc *Encoder, jsonKey string, val map[int32][]byte, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI32(mk, quoteKey)
		if err := enc.appendValBytes(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI32EnumNum append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapI32EnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[int32]T, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValEnumNum(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32EnumStr append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapI32EnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[int32]T, omitempty, quoteKey bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValEnumStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32Message append an JSON key map of key int32 and value proto MESSAGE to JSON contents.
func AppendMapI32Message[T protoreflect.ProtoMessage](enc *Encoder, jsonKey string, val map[int32]T, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI32(mk, quoteKey)
		if err := enc.appendValMessage(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI32WKTAnyObject append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapI32WKTAnyObject(enc *Encoder, jsonKey string, val map[int32]*anypb.Any, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI32(mk, quoteKey)
		if err := enc.appendValWKTAnyObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI32WKTAnyProto append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapI32WKTAnyProto(enc *Encoder, jsonKey string, val map[int32]*anypb.Any, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI32(mk, quoteKey)
		if err := enc.appendValWKTAnyProto(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI32WKTDurObject append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurObject(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI32(mk, quoteKey)
		if err := enc.appendValWKTDurObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI32WKTDurTimeStr append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurTimeStr(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurTimeStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTDurNano append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurNano(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTDurMicro append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurMicro(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTDurMilli append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurMilli(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTDurSecond append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurSecond(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurSecond(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTDurMinute append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurMinute(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurMinute(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTDurHour append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapI32WKTDurHour(enc *Encoder, jsonKey string, val map[int32]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTDurHour(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTTsObject append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI32WKTTsObject(enc *Encoder, jsonKey string, val map[int32]*timestamppb.Timestamp, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyI32(mk, quoteKey)
		if err := enc.appendValWKTTsObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapI32WKTTsLayout append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI32WKTTsLayout(enc *Encoder, jsonKey string, val map[int32]*timestamppb.Timestamp, omitempty, quoteKey bool, layout string) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTTsLayout(mv, layout)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTTsUnixNano append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI32WKTTsUnixNano(enc *Encoder, jsonKey string, val map[int32]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTTsUnixNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapI32WKTTsUnixMicro append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI32WKTTsUnixMicro(enc *Encoder, jsonKey string, val map[int32]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTTsUnixMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapI32WKTTsUnixMilli append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI32WKTTsUnixMilli(enc *Encoder, jsonKey string, val map[int32]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTTsUnixMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapI32WKTTsUnixSec append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapI32WKTTsUnixSec(enc *Encoder, jsonKey string, val map[int32]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyI32(mk, quoteKey)
		enc.appendValWKTTsUnixSec(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}
