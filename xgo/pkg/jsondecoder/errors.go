package jsondecoder

import (
	"fmt"
)

type Error struct {
	Field  string // The key of JSON. Maybe empty if error in top level.
	Offset int    // The offset in JSON input.
	Reason string // The message of error reason.
}

func (e *Error) Error() string {
	if e.Field == "" {
		return fmt.Sprintf("json: unmarshal: the input in offset %d: %s", e.Offset, e.Reason)
	}
	return fmt.Sprintf("json: unmarshal: the value of field %s in offset %d: %s", e.Field, e.Offset, e.Reason)
}

func errorWrap(dec *Decoder, err error) error {
	//if _err, ok := err.(*Error); ok {
	//	return _err
	//}
	return &Error{
		Field:  fmt.Sprintf("%s", dec.jsonKey),
		Offset: dec.offset - 1,
		Reason: err.Error(),
	}
}

// FIXME: fromat with Error
func ErrUnknownField(dec *Decoder) error {
	return fmt.Errorf("json: unmarshal: unknown field %q", dec.jsonKey)
}

// FIXME: fromat with Error
func ErrOneOfConflict(dec *Decoder) error {
	return fmt.Errorf("json: unmarshal: the field %s is type oneof, allow contains only one", dec.jsonKey)
}

// FIXME: fromat with Error
func ErrUnknownOneOfField(dec *Decoder, oneOfKey string) error {
	return fmt.Errorf("json: unmarshal: unknown field %q in oneof parts", oneOfKey)
}
