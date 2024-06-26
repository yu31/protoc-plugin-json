package jsondecoder

import (
	"errors"
	"fmt"
)

var (
	errInvalidFormatJSONKey = errors.New("the format of json key is invalid")
)

// Error used for invalid format of JSON contents.
type Error struct {
	JSONKey string // The key of JSON. Maybe empty if error in top level.
	Offset  int    // The offset in JSON input.
	Reason  string // The message of error reason.
}

func (e *Error) Error() string {
	if e.JSONKey == "" {
		return fmt.Sprintf("json: unmarshal: the input in offset %d: %s", e.Offset, e.Reason)
	}
	return fmt.Sprintf("json: unmarshal: the field %s in offset %d: %s", e.JSONKey, e.Offset, e.Reason)
}

func errorWrap(dec *Decoder, err error) error {
	if _err, ok := err.(*InternalError); ok {
		return _err
	}
	return &Error{
		JSONKey: fmt.Sprintf("%s", dec.jsonKey),
		Offset:  dec.offset - 1,
		Reason:  err.Error(),
	}
}

func ErrUnknownField(dec *Decoder) error {
	return &Error{
		JSONKey: fmt.Sprintf("%s", dec.jsonKey),
		Offset:  dec.offset - 1,
		Reason:  "unknown field",
	}
}

func ErrOneOfConflict(dec *Decoder) error {
	return &Error{
		JSONKey: fmt.Sprintf("%s", dec.jsonKey),
		Offset:  dec.offset - 1,
		Reason:  "type oneof allow contains only one",
	}
}

func ErrUnknownOneOfField(dec *Decoder, oneOfKey string) error {
	return &Error{
		JSONKey: fmt.Sprintf("%s", dec.jsonKey),
		Offset:  dec.offset - 1,
		Reason:  fmt.Sprintf("unknown field %q in oneof", oneOfKey),
	}
}

func ErrStructIsNIL(importPath string, msgName string) error {
	return fmt.Errorf("json: Unmarshal: nil pointer %s.(*%s)", importPath, msgName)
}
