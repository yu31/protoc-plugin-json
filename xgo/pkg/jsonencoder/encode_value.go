package jsonencoder

import (
	"encoding/base64"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// appendValI32 append the object value of type int32.
func (enc *Encoder) appendValI32(v int32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(int64(v), quote)
}

// appendValI64 append the object value of type int64.
func (enc *Encoder) appendValI64(v int64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(v, quote)
}

// appendValU32 append the object value of type uint32.
func (enc *Encoder) appendValU32(v uint32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(uint64(v), quote)
}

// appendValU64 append the object value of type uint64.
func (enc *Encoder) appendValU64(v uint64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendUint(v, quote)
}

// appendValF32 append the object value of type float32.
func (enc *Encoder) appendValF32(v float32, quote bool) {
	enc.appendCommaSeparator()
	enc.appendFloat(float64(v), 32, quote)
}

// appendValF64 append the object value of type float64.
func (enc *Encoder) appendValF64(v float64, quote bool) {
	enc.appendCommaSeparator()
	enc.appendFloat(v, 64, quote)
}

// appendValBool append the object value of type bool.
func (enc *Encoder) appendValBool(v bool, quote bool) {
	enc.appendCommaSeparator()
	enc.appendBool(v, quote)
}

// appendValStr append the object value of type string.
func (enc *Encoder) appendValStr(v string) {
	enc.appendCommaSeparator()
	enc.appendString(v)
}

// appendValBytes append the object value of type bytes.
func (enc *Encoder) appendValBytes(v []byte) error {
	enc.appendCommaSeparator()
	if v == nil {
		enc.writeString("null")
		return nil
	}

	enc.writeByte('"')
	if len(v) != 0 {
		encodedLen := base64.StdEncoding.EncodedLen(len(v))
		// TODO: Improved the alloc logic.
		if encodedLen <= 1024 {
			// The encoded bytes are short enough to allocate for, and
			// Encoding.Encode is still cheaper.
			dst := make([]byte, encodedLen)
			base64.StdEncoding.Encode(dst, v)
			enc.writeBytes(dst)
		} else {
			// The encoded bytes are too long to cheaply allocate, and
			// Encoding.Encode is no longer noticeably cheaper.
			be := base64.NewEncoder(base64.StdEncoding, enc)
			if _, err := be.Write(v); err != nil {
				return err
			}
			if err := be.Close(); err != nil {
				return err
			}
		}
	}
	enc.writeByte('"')
	return nil
}

// appendValInterface append the object value of type interface.
func (enc *Encoder) appendValInterface(v interface{}) error {
	enc.appendCommaSeparator()
	return enc.appendInterface(v)
}

// appendValEnumNum append the object value of type proto enum.
func (enc *Encoder) appendValEnumNum(v protoreflect.Enum, quote bool) {
	enc.appendCommaSeparator()
	enc.appendInt(int64(v.Number()), quote)
}

// appendValEnumStr append the object value of type proto enum.
func (enc *Encoder) appendValEnumStr(v protoreflect.Enum) {
	enc.appendCommaSeparator()
	enc.appendString(protoimpl.X.EnumStringOf(v.Descriptor(), v.Number()))
}

// appendValMessage append the object value of type proto message.
func (enc *Encoder) appendValMessage(v protoreflect.ProtoMessage) error {
	enc.appendCommaSeparator()
	return enc.appendInterface(v)
}

// appendValWKTAnyObject append the object value of type WKT any.
func (enc *Encoder) appendValWKTAnyObject(v *anypb.Any) error {
	//enc.appendCommaSeparator()
	if err := enc.appendValMessage(v); err != nil {
		return err
	}
	return nil
}

// appendValWKTAnyProto append the object value of type WKT any.
func (enc *Encoder) appendValWKTAnyProto(v *anypb.Any) error {
	enc.appendCommaSeparator()
	b, err := pMarshal.Marshal(v)
	if err != nil {
		return err
	}
	enc.writeBytes(b)
	return nil
}

// appendValWKTDurObject append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurObject(v *durationpb.Duration) error {
	//enc.appendCommaSeparator()
	if err := enc.appendValMessage(v); err != nil {
		return err
	}
	return nil
}

// appendValWKTDurTimeStr append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurTimeStr(v *durationpb.Duration) {
	//enc.appendCommaSeparator()
	enc.appendValStr(v.AsDuration().String())
}

// appendValWKTDurNano append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurNano(v *durationpb.Duration, quote bool) {
	enc.appendValI64(v.AsDuration().Nanoseconds(), quote)
}

// appendValWKTDurMicro append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurMicro(v *durationpb.Duration, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValI64(v.AsDuration().Microseconds(), quote)
}

// appendValWKTDurMilli append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurMilli(v *durationpb.Duration, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValI64(v.AsDuration().Milliseconds(), quote)
}

// appendValWKTDurSecond append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurSecond(v *durationpb.Duration, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValF64(v.AsDuration().Seconds(), quote)
}

// appendValWKTDurMinute append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurMinute(v *durationpb.Duration, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValF64(v.AsDuration().Minutes(), quote)
}

// appendValWKTDurHour append the object value of type WKT duration.
func (enc *Encoder) appendValWKTDurHour(v *durationpb.Duration, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValF64(v.AsDuration().Hours(), quote)
}

// appendValWKTTsObject append the object value of type WKT timestamp.
func (enc *Encoder) appendValWKTTsObject(v *timestamppb.Timestamp) error {
	//enc.appendCommaSeparator()
	if err := enc.appendValMessage(v); err != nil {
		return err
	}
	return nil
}

// appendValWKTTsLayout append the object value of type WKT timestamp.
func (enc *Encoder) appendValWKTTsLayout(v *timestamppb.Timestamp, layout string) {
	//enc.appendCommaSeparator()
	enc.appendValStr(v.AsTime().Format(layout))
}

// appendValWKTTsUnixNano append the object value of type WKT timestamp.
func (enc *Encoder) appendValWKTTsUnixNano(v *timestamppb.Timestamp, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValI64(v.AsTime().UnixNano(), quote)
}

// appendValWKTTsUnixMicro append the object value of type WKT timestamp.
func (enc *Encoder) appendValWKTTsUnixMicro(v *timestamppb.Timestamp, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValI64(v.AsTime().UnixMicro(), quote)
}

// appendValWKTTsUnixMilli append the object value of type WKT timestamp.
func (enc *Encoder) appendValWKTTsUnixMilli(v *timestamppb.Timestamp, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValI64(v.AsTime().UnixMilli(), quote)
}

// appendValWKTTsUnixSec append the object value of type WKT timestamp.
func (enc *Encoder) appendValWKTTsUnixSec(v *timestamppb.Timestamp, quote bool) {
	//enc.appendCommaSeparator()
	enc.appendValI64(v.AsTime().Unix(), quote)
}
