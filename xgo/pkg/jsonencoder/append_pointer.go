package jsonencoder

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// AppendPtrI32 append an JSON key with pointer of int32 to JSON contents.
func AppendPtrI32(enc *Encoder, jsonKey string, val *int32, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValI32(*val, quote)
}

// AppendPtrI64 append an JSON key with pointer of int64 to JSON contents.
func AppendPtrI64(enc *Encoder, jsonKey string, val *int64, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValI64(*val, quote)
}

// AppendPtrU32 append an JSON key with pointer of uint32 to JSON contents.
func AppendPtrU32(enc *Encoder, jsonKey string, val *uint32, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValU32(*val, quote)
}

// AppendPtrU64 append an JSON key with pointer of uint64 to JSON contents.
func AppendPtrU64(enc *Encoder, jsonKey string, val *uint64, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValU64(*val, quote)
}

// AppendPtrF32 append an JSON key with pointer of float32 to JSON contents.
func AppendPtrF32(enc *Encoder, jsonKey string, val *float32, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValF32(*val, quote)
}

// AppendPtrF64 append an JSON key with pointer of float64 to JSON contents.
func AppendPtrF64(enc *Encoder, jsonKey string, val *float64, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValF64(*val, quote)
}

// AppendPtrBool append an JSON key with pointer of bool to JSON contents.
func AppendPtrBool(enc *Encoder, jsonKey string, val *bool, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValBool(*val, quote)
}

// AppendPtrStr append an JSON key with pointer of string to JSON contents.
func AppendPtrStr(enc *Encoder, jsonKey string, val *string, omitempty bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValStr(*val)
}

// AppendPtrEnumNum append an JSON key with value of proto enum to JSON contents.
func AppendPtrEnumNum[T protoreflect.Enum](enc *Encoder, jsonKey string, val *T, omitempty, quote bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValEnumNum(*val, quote)
}

// AppendPtrEnumStr append an JSON key with value of proto enum to JSON contents.
func AppendPtrEnumStr[T protoreflect.Enum](enc *Encoder, jsonKey string, val *T, omitempty bool) {
	if omitempty && val == nil {
		return
	}
	enc.AppendObjectKey(jsonKey)
	if val == nil {
		enc.AppendValNULL()
		return
	}
	enc.appendValEnumStr(*val)
}
