package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendMapStrI32 append an JSON key map of key string and value int32 to JSON contents.
func AppendMapStrI32(enc *Encoder, jsonKey string, val map[string]int32, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValI32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrI64 append an JSON key map of key string and value int64 to JSON contents.
func AppendMapStrI64(enc *Encoder, jsonKey string, val map[string]int64, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValI64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrU32 append an JSON key map of key string and value uint32 to JSON contents.
func AppendMapStrU32(enc *Encoder, jsonKey string, val map[string]uint32, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValU32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrU64 append an JSON key map of key string and value uint64 to JSON contents.
func AppendMapStrU64(enc *Encoder, jsonKey string, val map[string]uint64, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValU64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrF32 append an JSON key map of key string and value float32 to JSON contents.
func AppendMapStrF32(enc *Encoder, jsonKey string, val map[string]float32, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValF32(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrF64 append an JSON key map of key string and value float64 to JSON contents.
func AppendMapStrF64(enc *Encoder, jsonKey string, val map[string]float64, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValF64(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrBool append an JSON key map of key string and value bool to JSON contents.
func AppendMapStrBool(enc *Encoder, jsonKey string, val map[string]bool, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValBool(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrStr append an JSON key map of key string and value string to JSON contents.
func AppendMapStrStr(enc *Encoder, jsonKey string, val map[string]string, omitempty bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrBytes append an JSON key map of key string and value bytes to JSON contents.
func AppendMapStrBytes(enc *Encoder, jsonKey string, val map[string][]byte, omitempty bool) error {
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
		enc.appendMapKeyStr(mk)
		if err := enc.appendValBytes(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapStrEnumNum append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapStrEnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[string]T, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValEnumNum(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrEnumStr append an JSON key map of key int32 and value proto ENUM to JSON contents.
func AppendMapStrEnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val map[string]T, omitempty bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValEnumStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrMessage append an JSON key map of key string and value proto MESSAGE to JSON contents.
func AppendMapStrMessage[T protoreflect.ProtoMessage](enc *Encoder, jsonKey string, val map[string]T, omitempty bool) error {
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
		enc.appendMapKeyStr(mk)
		if err := enc.appendValMessage(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapStrWKTAnyObject append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapStrWKTAnyObject(enc *Encoder, jsonKey string, val map[string]*anypb.Any, omitempty bool) error {
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
		enc.appendMapKeyStr(mk)
		if err := enc.appendValWKTAnyObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapStrWKTAnyProto append an JSON key map of key int32 and value WKT any to JSON contents.
func AppendMapStrWKTAnyProto(enc *Encoder, jsonKey string, val map[string]*anypb.Any, omitempty bool) error {
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
		enc.appendMapKeyStr(mk)
		if err := enc.appendValWKTAnyProto(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapStrWKTDurObject append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurObject(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty bool) error {
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
		enc.appendMapKeyStr(mk)
		if err := enc.appendValWKTDurObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapStrWKTDurTimeStr append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurTimeStr(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurTimeStr(mv)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTDurNano append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurNano(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTDurMicro append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurMicro(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTDurMilli append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurMilli(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTDurSecond append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurSecond(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurSecond(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTDurMinute append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurMinute(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurMinute(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTDurHour append an JSON key map of key int32 and value WKT duration to JSON contents.
func AppendMapStrWKTDurHour(enc *Encoder, jsonKey string, val map[string]*durationpb.Duration, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTDurHour(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTTsObject append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapStrWKTTsObject(enc *Encoder, jsonKey string, val map[string]*timestamppb.Timestamp, omitempty bool) error {
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
		enc.appendMapKeyStr(mk)
		if err := enc.appendValWKTTsObject(mv); err != nil {
			return err
		}
	}
	enc.AppendObjectEnd()
	return nil
}

// AppendMapStrWKTTsLayout append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapStrWKTTsLayout(enc *Encoder, jsonKey string, val map[string]*timestamppb.Timestamp, omitempty bool, layout string) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTTsLayout(mv, layout)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTTsUnixNano append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapStrWKTTsUnixNano(enc *Encoder, jsonKey string, val map[string]*timestamppb.Timestamp, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTTsUnixNano(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTTsUnixMicro append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapStrWKTTsUnixMicro(enc *Encoder, jsonKey string, val map[string]*timestamppb.Timestamp, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTTsUnixMicro(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTTsUnixMilli append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapStrWKTTsUnixMilli(enc *Encoder, jsonKey string, val map[string]*timestamppb.Timestamp, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTTsUnixMilli(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}

// AppendMapStrWKTTsUnixSec append an JSON key map of key int32 and value WKT timestamp to JSON contents.
func AppendMapStrWKTTsUnixSec(enc *Encoder, jsonKey string, val map[string]*timestamppb.Timestamp, omitempty, quoteVal bool) {
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
		enc.appendMapKeyStr(mk)
		enc.appendValWKTTsUnixSec(mv, quoteVal)
	}
	enc.AppendObjectEnd()
}
