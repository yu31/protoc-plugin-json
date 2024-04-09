package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppendValI32 append an JSON key with value of int32 to JSON contents.
func AppendValI32(enc *Encoder, jsonKey string, val int32, omitempty, quote bool) {
	if omitempty && val == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValI32(val, quote)
}

// AppendValI64 append an JSON key with value of int64 to JSON contents.
func AppendValI64(enc *Encoder, jsonKey string, val int64, omitempty, quote bool) {
	if omitempty && val == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValI64(val, quote)
}

// AppendValU32 append an JSON key with value of uint32 to JSON contents.
func AppendValU32(enc *Encoder, jsonKey string, val uint32, omitempty, quote bool) {
	if omitempty && val == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValU32(val, quote)
}

// AppendValU64 append an JSON key with value of uint64 to JSON contents.
func AppendValU64(enc *Encoder, jsonKey string, val uint64, omitempty, quote bool) {
	if omitempty && val == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValU64(val, quote)
}

// AppendValF32 append an JSON key with value of float32 to JSON contents.
func AppendValF32(enc *Encoder, jsonKey string, val float32, omitempty, quote bool) {
	if omitempty && val == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValF32(val, quote)
}

// AppendValF64 append an JSON key with value of float64 to JSON contents.
func AppendValF64(enc *Encoder, jsonKey string, val float64, omitempty, quote bool) {
	if omitempty && val == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValF64(val, quote)
}

// AppendValBool append an JSON key with value of bool to JSON contents.
func AppendValBool(enc *Encoder, jsonKey string, val bool, omitempty, quote bool) {
	if omitempty && !val {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValBool(val, quote)
}

// AppendValStr append an JSON key with value of string to JSON contents.
func AppendValStr(enc *Encoder, jsonKey string, val string, omitempty bool) {
	if omitempty && len(val) == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValStr(val)
}

// AppendValBytes append an JSON key with value of bytes to JSON contents.
func AppendValBytes(enc *Encoder, jsonKey string, val []byte, omitempty bool) error {
	if omitempty && len(val) == 0 {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if err := enc.appendValBytes(val); err != nil {
		return err
	}
	return nil
}

// AppendValEnumNum append an JSON key with value of proto enum to JSON contents.
func AppendValEnumNum(enc *Encoder, jsonKey string, val protoreflect.Enum, omitempty, quote bool) {
	if omitempty && val.Number() == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValEnumNum(val, quote)
}

// AppendValEnumStr append an JSON key with value of proto enum to JSON contents.
func AppendValEnumStr(enc *Encoder, jsonKey string, val protoreflect.Enum, omitempty bool) {
	if omitempty && val.Number() == 0 {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValEnumStr(val)
}

// AppendValMessage append an JSON key with value of proto MESSAGE to JSON contents.
//func AppendValProtoMessage[T PtrGoAny[U], U any](enc *Encoder,  jsonKey string,
//	val T) error {
func AppendValMessage[U any, T *U](enc *Encoder, jsonKey string, val T, omitempty bool) error {
	if omitempty && val == nil {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if err := enc.appendValInterface(val); err != nil {
		return err
	}
	//if omitempty && reflect.ValueOf(val).IsNil() {
	//	// In Golang. The val will not be considered as nil if it has a specific type(stupid design !).
	//	return nil
	//}
	//if err := enc.appendValMessage(val); err != nil {
	//	return err
	//}
	return nil
}

// AppendValWKTAnyObject append an JSON key with value of WKT any to JSON contents.
func AppendValWKTAnyObject(enc *Encoder, jsonKey string, val *anypb.Any, omitempty bool) error {
	if omitempty && val == nil {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if err := enc.appendValWKTAnyObject(val); err != nil {
		return err
	}
	return nil
}

// AppendValWKTAnyProto append an JSON key with value of WKT any to JSON contents.
func AppendValWKTAnyProto(enc *Encoder, jsonKey string, val *anypb.Any, omitempty bool) error {
	if omitempty && val == nil {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if err := enc.appendValWKTAnyProto(val); err != nil {
		return err
	}
	return nil
}

// AppendValWKTDurObject append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurObject(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty bool) error {
	if omitempty && val == nil {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if err := enc.appendValWKTDurObject(val); err != nil {
		return err
	}
	return nil
}

// AppendValWKTDurTimeStr append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurTimeStr(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurTimeStr(val)
}

// AppendValWKTDurNano append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurNano(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurNano(val, quote)
}

// AppendValWKTDurMicro append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurMicro(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurMicro(val, quote)
}

// AppendValWKTDurMilli append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurMilli(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurMilli(val, quote)
}

// AppendValWKTDurSecond append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurSecond(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurSecond(val, quote)
}

// AppendValWKTDurMinute append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurMinute(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurMinute(val, quote)
}

// AppendValWKTDurHour append an JSON key with value of WKT duration to JSON contents.
func AppendValWKTDurHour(enc *Encoder, jsonKey string, val *durationpb.Duration, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTDurHour(val, quote)
}

// AppendValWKTTsObject append an JSON key with value of WKT timestamp to JSON contents.
func AppendValWKTTsObject(enc *Encoder, jsonKey string, val *timestamppb.Timestamp, omitempty bool) error {
	if omitempty && val == nil {
		return nil
	}
	enc.AppendObjectKey(jsonKey)
	if err := enc.appendValWKTTsObject(val); err != nil {
		return err
	}
	return nil
}

// AppendValWKTTsLayout append an JSON key with value of WKT timestamp to JSON contents.
func AppendValWKTTsLayout(enc *Encoder, jsonKey string, val *timestamppb.Timestamp, omitempty bool, layout string) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTTsLayout(val, layout)
}

// AppendValWKTTsUnixNano append an JSON key with value of WKT timestamp to JSON contents.
func AppendValWKTTsUnixNano(enc *Encoder, jsonKey string, val *timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTTsUnixNano(val, quote)
}

// AppendValWKTTsUnixMicro append an JSON key with value of WKT timestamp to JSON contents.
func AppendValWKTTsUnixMicro(enc *Encoder, jsonKey string, val *timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTTsUnixMicro(val, quote)
}

// AppendValWKTTsUnixMilli append an JSON key with value of WKT timestamp to JSON contents.
func AppendValWKTTsUnixMilli(enc *Encoder, jsonKey string, val *timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTTsUnixMilli(val, quote)
}

// AppendValWKTTsUnixSec append an JSON key with value of WKT timestamp to JSON contents.
func AppendValWKTTsUnixSec(enc *Encoder, jsonKey string, val *timestamppb.Timestamp, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	enc.appendValWKTTsUnixSec(val, quote)
}
