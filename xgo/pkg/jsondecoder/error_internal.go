package jsondecoder

import (
	"fmt"
	"runtime"
)

//const PhasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

// InternalError is used as error when we end up with something that
// shouldn't happen. It can indicate a bug in the JSON decoder, or that
// something is editing the data slice while the decoder executes.
type InternalError struct {
	Reason string // The message of error reason.
	Stack  string // The runtime stack of error occurs.
}

func (e *InternalError) Error() string {
	return e.Reason
}

func phaseDecodeError(funcName string, opCode OpCode) *InternalError {
	buf := make([]byte, 8192)
	_ = runtime.Stack(buf, true)

	return &InternalError{
		Reason: fmt.Sprintf("JSON decoder out of sync - func: %s opCode: %d", funcName, opCode),
		Stack:  string(buf),
	}
}
