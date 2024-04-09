package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendMapBoolI32 append an JSON key map of key int32 and value int32 to JSON contents.
func AppendMapBoolI32(enc *Encoder, jsonKey string, val map[bool]int32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValI32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolI64 append an JSON key map of key int32 and value int64 to JSON contents.
func AppendMapBoolI64(enc *Encoder, jsonKey string, val map[bool]int64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValI64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolU32 append an JSON key map of key int32 and value uint32 to JSON contents.
func AppendMapBoolU32(enc *Encoder, jsonKey string, val map[bool]uint32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValU32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolU64 append an JSON key map of key int32 and value uint64 to JSON contents.
func AppendMapBoolU64(enc *Encoder, jsonKey string, val map[bool]uint64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValU64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolF32 append an JSON key map of key int32 and value float32 to JSON contents.
func AppendMapBoolF32(enc *Encoder, jsonKey string, val map[bool]float32, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValF32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolF64 append an JSON key map of key int32 and value float64 to JSON contents.
func AppendMapBoolF64(enc *Encoder, jsonKey string, val map[bool]float64, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValF64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolBool append an JSON key map of key int32 and value bool to JSON contents.
func AppendMapBoolBool(enc *Encoder, jsonKey string, val map[bool]bool, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValBool(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolStr append an JSON key map of key int32 and value string to JSON contents.
func AppendMapBoolStr(enc *Encoder, jsonKey string, val map[bool]string, omitempty, quoteKey bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolBytes append an JSON key map of key int32 and value bytes to JSON contents.
func AppendMapBoolBytes(enc *Encoder, jsonKey string, val map[bool][]byte, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyBool(mk, quoteKey)
		if err := enc.appendValBytes(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapBoolEnumNum append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapBoolEnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[bool]T, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValEnumNum(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolEnumStr append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapBoolEnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[bool]T, omitempty, quoteKey bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValEnumStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapBoolMessage append an JSON key map of key int32 and value proto MESSAGE to JSON contents.
func AppendMapBoolMessage[T protoreflect.ProtoMessage](enc *Encoder, jsonKey string, val map[bool]T, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyBool(mk, quoteKey)
		if err := enc.appendValMessage(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapBoolWKTAnyObject append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapBoolWKTAnyObject(enc *Encoder, jsonKey string, val map[bool]*anypb.Any, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyBool(mk, quoteKey)
		if err := enc.appendValWKTAnyObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapBoolWKTAnyProto append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapBoolWKTAnyProto(enc *Encoder, jsonKey string, val map[bool]*anypb.Any, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyBool(mk, quoteKey)
		if err := enc.appendValWKTAnyProto(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapBoolWKTDurObject append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurObject(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyBool(mk, quoteKey)
		if err := enc.appendValWKTDurObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapBoolWKTDurTimeStr append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurTimeStr(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurTimeStr(mv)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTDurNano append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurNano(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTDurMicro append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurMicro(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTDurMilli append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurMilli(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTDurSecond append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurSecond(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurSecond(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTDurMinute append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurMinute(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurMinute(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTDurHour append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapBoolWKTDurHour(enc *Encoder, jsonKey string, val map[bool]*durationpb.Duration, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTDurHour(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTTsObject append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapBoolWKTTsObject(enc *Encoder, jsonKey string, val map[bool]*timestamppb.Timestamp, omitempty, quoteKey bool) error {
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
		enc.appendMapKeyBool(mk, quoteKey)
		if err := enc.appendValWKTTsObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapBoolWKTTsLayout append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapBoolWKTTsLayout(enc *Encoder, jsonKey string, val map[bool]*timestamppb.Timestamp, omitempty, quoteKey bool, layout string) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTTsLayout(mv, layout)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTTsUnixNano append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapBoolWKTTsUnixNano(enc *Encoder, jsonKey string, val map[bool]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTTsUnixNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTTsUnixMicro append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapBoolWKTTsUnixMicro(enc *Encoder, jsonKey string, val map[bool]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTTsUnixMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTTsUnixMilli append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapBoolWKTTsUnixMilli(enc *Encoder, jsonKey string, val map[bool]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTTsUnixMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}

// AppendMapBoolWKTTsUnixSec append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapBoolWKTTsUnixSec(enc *Encoder, jsonKey string, val map[bool]*timestamppb.Timestamp, omitempty, quoteKey, quoteVal bool) {
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
		enc.appendMapKeyBool(mk, quoteKey)
		enc.appendValWKTTsUnixSec(mv, quoteVal)
	}
	enc.AppendObjectEnd()
	return
}
