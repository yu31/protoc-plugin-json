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
		return fmt.Sprintf("json: unmarshal the input in offset %d: %s", e.Offset, e.Reason)
	}
	return fmt.Sprintf("json: unmarshal the value of field %s in offset %d: %s", e.Field, e.Offset, e.Reason)
}

func errorWrap(jsonKey string, offset int, err error) error {
	//if _err, ok := err.(*Error); ok {
	//	return _err
	//}
	return &Error{
		Field:  fmt.Sprintf("%s", jsonKey),
		Offset: offset - 1,
		Reason: err.Error(),
	}
}
