package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendListStr append an JSON key with list of string to JSON contents.
func AppendListStr(enc *Encoder, jsonKey string, val []string, omitempty bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValStr(item)
	}
	enc.AppendListEnd()
}

// AppendListBytes append an JSON key with list of bytes to JSON contents.
func AppendListBytes(enc *Encoder, jsonKey string, val [][]byte, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendListBegin()
	for _, item := range val {
		if err := enc.appendValBytes(item); err != nil {
			return err
		}
	}
	enc.AppendListEnd()
	return nil
}

// AppendListBool append an JSON key with list of bool to JSON contents.
func AppendListBool(enc *Encoder, jsonKey string, val []bool, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValBool(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListI32 append an JSON key with list of int32 to JSON contents.
func AppendListI32(enc *Encoder, jsonKey string, val []int32, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValI32(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListI64 append an JSON key with list of int64 to JSON contents.
func AppendListI64(enc *Encoder, jsonKey string, val []int64, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValI64(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListU32 append an JSON key with list of uint32 to JSON contents.
func AppendListU32(enc *Encoder, jsonKey string, val []uint32, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValU32(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListU64 append an JSON key with list of uint64 to JSON contents.
func AppendListU64(enc *Encoder, jsonKey string, val []uint64, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValU64(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListF32 append an JSON key with list of float32 to JSON contents.
func AppendListF32(enc *Encoder, jsonKey string, val []float32, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValF32(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListF64 append an JSON key with list of float64 to JSON contents.
func AppendListF64(enc *Encoder, jsonKey string, val []float64, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValF64(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListEnumNum append an JSON key with list of proto ENUM to JSON contents.
func AppendListEnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val []T, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValEnumNum(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListEnumStr append an JSON key with list of proto ENUM to JSON contents.
func AppendListEnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val []T, omitempty bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValEnumStr(item)
	}
	enc.AppendListEnd()
}

// AppendListMessage append an JSON key with list of proto MESSAGE to JSON contents.
func AppendListMessage[T protoreflect.ProtoMessage](enc *Encoder, jsonKey string, val []T, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendListBegin()
	for _, item := range val {
		if err := enc.appendValMessage(item); err != nil {
			return err
		}
	}
	enc.AppendListEnd()
	return nil
}

// AppendListWKTAnyObject append an JSON key with list of WKT any to JSON contents.
func AppendListWKTAnyObject(enc *Encoder, jsonKey string, val []*anypb.Any, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendListBegin()
	for _, item := range val {
		if err := enc.appendValWKTAnyObject(item); err != nil {
			return err
		}
	}
	enc.AppendListEnd()
	return nil
}

// AppendListWKTAnyProto append an JSON key with list of WKT any to JSON contents.
func AppendListWKTAnyProto(enc *Encoder, jsonKey string, val []*anypb.Any, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendListBegin()
	for _, item := range val {
		if err := enc.appendValWKTAnyProto(item); err != nil {
			return err
		}
	}
	enc.AppendListEnd()
	return nil
}

// AppendListWKTDurObject append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurObject(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendListBegin()
	for _, item := range val {
		if err := enc.appendValWKTDurObject(item); err != nil {
			return err
		}
	}
	enc.AppendListEnd()
	return nil
}

// AppendListWKTDurTimeStr append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurTimeStr(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurTimeStr(item)
	}
	enc.AppendListEnd()
}

// AppendListWKTDurNano append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurNano(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurNano(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTDurMicro append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurMicro(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurMicro(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTDurMilli append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurMilli(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurMilli(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTDurSecond append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurSecond(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurSecond(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTDurMinute append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurMinute(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurMinute(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTDurHour append an JSON key with list of WKT duration to JSON contents.
func AppendListWKTDurHour(enc *Encoder, jsonKey string, val []*durationpb.Duration, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTDurHour(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTTsObject append an JSON key with list of WKT timestamp to JSON contents.
func AppendListWKTTsObject(enc *Encoder, jsonKey string, val []*timestamppb.Timestamp, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return nil
	}
	enc.AppendListBegin()
	for _, item := range val {
		if err := enc.appendValWKTTsObject(item); err != nil {
			return err
		}
	}
	enc.AppendListEnd()
	return nil
}

// AppendListWKTTsLayout append an JSON key with list of WKT timestamp to JSON contents.
func AppendListWKTTsLayout(enc *Encoder, jsonKey string, val []*timestamppb.Timestamp, omitempty bool, layout string) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTTsLayout(item, layout)
	}
	enc.AppendListEnd()
}

// AppendListWKTTsUnixNano append an JSON key with list of WKT timestamp to JSON contents.
func AppendListWKTTsUnixNano(enc *Encoder, jsonKey string, val []*timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTTsUnixNano(item, quote)
	}
	enc.AppendListEnd()

}

// AppendListWKTTsUnixMicro append an JSON key with list of WKT timestamp to JSON contents.
func AppendListWKTTsUnixMicro(enc *Encoder, jsonKey string, val []*timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTTsUnixMicro(item, quote)
	}
	enc.AppendListEnd()
}

// AppendListWKTTsUnixMilli append an JSON key with list of WKT timestamp to JSON contents.
func AppendListWKTTsUnixMilli(enc *Encoder, jsonKey string, val []*timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTTsUnixMilli(item, quote)
	}
	enc.AppendListEnd()

}

// AppendListWKTTsUnixSec append an JSON key with list of WKT timestamp to JSON contents.
func AppendListWKTTsUnixSec(enc *Encoder, jsonKey string, val []*timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.AppendListBegin()
	for _, item := range val {
		enc.appendValWKTTsUnixSec(item, quote)
	}
	enc.AppendListEnd()
}
