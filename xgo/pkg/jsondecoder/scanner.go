package jsondecoder

import "strconv"

type OpCode int8

// These values are returned by the state transition functions
// assigned to scanner.state and the method scanner.eof.
// They give details about the current state of the scan that
// callers might be interested to know about.
// It is okay to ignore the return value of any particular
// call to scanner.state: if one call returns scanError,
// every subsequent call will return scanError too.
const (
	// label for Continue

	scanContinue     OpCode = iota // 0 uninteresting byte
	scanLiteralBegin               // 1 end implied by next result != scanContinue
	scanObjectBegin                // 2 begin object
	scanObjectKey                  // 3 just finished object key (string)
	scanObjectValue                // 4 just finished non-last object value
	scanObjectEnd                  // 5 end object (implies scanObjectValue if possible)
	scanArrayBegin                 // 6 begin array
	scanArrayElem                  // 7 just finished array elements
	scanArrayEnd                   // 8 end array (implies scanArrayElem if possible)
	scanSkipSpace                  // 9 space byte; can skip; known to be last "continue" result

	// label for Stop

	scanEnd   // 10 top-level value ended *before* this byte; known to be first "stop" result
	scanError // 11 hit an error, scanner.err.
)

type ParsePhase int8

// These values are stored in the parseState stack.
// They give the current state of a composite value
// being scanned. If the parser is inside a nested value
// the parseState describes the nested state, outermost at entry 0.
const (
	parseObjectKey   ParsePhase = iota // parsing object key (before colon)
	parseObjectValue                   // parsing object value (after colon)
	parseArrayElem                     // parsing array elements
)

// This limits the max nesting depth to prevent stack overflow.
// This is permitted by https://tools.ietf.org/html/rfc7159#section-9
const maxNestingDepth = 10000

// A scanner is a JSON scanning state machine.
// Callers call scan.reset and then pass bytes in one at a time
// by calling scan.step(&scan, c) for each byte.
// The return value, referred to as an opcode, tells the
// caller about significant parsing events like beginning
// and ending literals, objects, and arrays, so that the
// caller can follow along if it wishes.
// The return value scanEnd indicates that a single top-level
// JSON value has been completed, *before* the byte that
// just got passed in.  (The indication must be delayed in order
// to recognize the end of numbers: is 123 a whole value or
// the beginning of 12345e+6?).
type scanner struct {
	// The step is a func to be called to execute the next transition.
	// Also tried using an integer constant and a single func
	// with a switch, but using the func directly was 10% faster
	// on a 64-bit Mac Mini, and it's nicer to read.
	step func(*scanner, byte) OpCode

	// Reached end of top-level value.
	endTop bool

	// Stack of what we're in the middle of - array values, object keys, object values.
	parseState []ParsePhase

	// Error that happened, if any.
	err error

	// total bytes consumed, updated by decoder.Decode (and deliberately
	// not set to zero by scan.reset)
	bytes int
}

// eof tells the scanner that the end of input has been reached.
// It returns a scan status just as s.step does.
func (s *scanner) eof() OpCode {
	if s.err != nil {
		return scanError
	}
	if s.endTop {
		return scanEnd
	}
	s.step(s, ' ')
	if s.endTop {
		return scanEnd
	}
	if s.err == nil {
		s.err = &SyntaxError{reason: "unexpected end of JSON input", Offset: s.bytes}
	}
	return scanError
}

// pushParseState pushes a new parse state p onto the parse stack.
// an error state is returned if maxNestingDepth was exceeded, otherwise successState is returned.
func (s *scanner) pushParseState(c byte, parseState ParsePhase, successState OpCode) OpCode {
	s.parseState = append(s.parseState, parseState)
	if len(s.parseState) <= maxNestingDepth {
		return successState
	}
	return s.error(c, "exceeded max depth")
}

// popParseState pops a parse state (already obtained) offset the stack
// and updates s.step accordingly.
func (s *scanner) popParseState() {
	n := len(s.parseState) - 1
	s.parseState = s.parseState[0:n]
	if n == 0 {
		s.step = stateEndTop
		s.endTop = true
	} else {
		s.step = stateEndValue
	}
}

// stateBeginElementOrEmpty is the state after reading `[`.
func stateBeginElementOrEmpty(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return scanSkipSpace
	}
	if c == ']' {
		return stateEndValue(s, c)
	}
	return stateBeginValue(s, c)
}

// stateBeginValue is the state at the beginning of the input.
func stateBeginValue(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return scanSkipSpace
	}
	switch c {
	case '{':
		s.step = stateBeginStringOrEmpty
		return s.pushParseState(c, parseObjectKey, scanObjectBegin)
	case '[':
		s.step = stateBeginElementOrEmpty
		return s.pushParseState(c, parseArrayElem, scanArrayBegin)
	case '"':
		s.step = stateInString
		return scanLiteralBegin
	case '-':
		s.step = stateNeg
		return scanLiteralBegin
	case '0': // beginning of 0.123
		s.step = state0
		return scanLiteralBegin
	case 't': // beginning of true
		s.step = stateT
		return scanLiteralBegin
	case 'f': // beginning of false
		s.step = stateF
		return scanLiteralBegin
	case 'n': // beginning of null
		s.step = stateN
		return scanLiteralBegin
	}
	if '1' <= c && c <= '9' { // beginning of 1234 or 1234.5
		s.step = state1
		return scanLiteralBegin
	}
	return s.error(c, "looking for beginning of value")
}

// stateBeginStringOrEmpty is the state after reading `{`.
func stateBeginStringOrEmpty(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return scanSkipSpace
	}
	if c == '}' {
		n := len(s.parseState)
		s.parseState[n-1] = parseObjectValue
		return stateEndValue(s, c)
	}
	return stateBeginLiteral(s, c)
}

// stateBeginLiteral is the state after reading `{"key": value,` or `{key: value,`
func stateBeginLiteral(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return scanSkipSpace
	}
	switch c {
	case '"':
		s.step = stateInString
		return scanLiteralBegin
	case '-':
		s.step = stateNeg
		return scanLiteralBegin
	case 't': // beginning of true
		s.step = stateT
		return scanLiteralBegin
	case 'f': // beginning of false
		s.step = stateF
		return scanLiteralBegin
	}
	if '1' <= c && c <= '9' { // beginning of 1234 or 1234.5
		s.step = state1
		return scanLiteralBegin
	}
	return s.error(c, "looking for beginning of object key")
}

//// stateBeginLiteral is the state after reading `{"key": value,`.
//func stateBeginLiteral(s *scanner, c byte) OpCode {
//	if isSpace(c) {
//		return scanSkipSpace
//	}
//	if c == '"' {
//		s.step = stateInString
//		return scanLiteralBegin
//	}
//	return s.error(c, "looking for beginning of object key string")
//}

// stateEndValue is the state after completing a value,
// such as after reading `{}` or `true` or `["x"`.
func stateEndValue(s *scanner, c byte) OpCode {
	n := len(s.parseState)
	if n == 0 {
		// Completed top-level before the current byte.
		s.step = stateEndTop
		s.endTop = true
		return stateEndTop(s, c)
	}
	if isSpace(c) {
		s.step = stateEndValue
		return scanSkipSpace
	}
	ps := s.parseState[n-1]
	switch ps {
	case parseObjectKey:
		if c == ':' {
			s.parseState[n-1] = parseObjectValue
			s.step = stateBeginValue
			return scanObjectKey
		}
		return s.error(c, "after object key")
	case parseObjectValue:
		if c == ',' {
			s.parseState[n-1] = parseObjectKey
			s.step = stateBeginLiteral
			return scanObjectValue
		}
		if c == '}' {
			s.popParseState()
			return scanObjectEnd
		}
		return s.error(c, "after object key:value pair")
	case parseArrayElem:
		if c == ',' {
			s.step = stateBeginValue
			return scanArrayElem
		}
		if c == ']' {
			s.popParseState()
			return scanArrayEnd
		}
		return s.error(c, "after array element")
	}
	return s.error(c, "")
}

// stateEndTop is the state after finishing the top-level value,
// such as after reading `{}` or `[1,2,3]`.
// Only space characters should be seen now.
func stateEndTop(s *scanner, c byte) OpCode {
	if !isSpace(c) {
		// Complain about non-space byte on next call.
		s.error(c, "after top-level value")
	}
	return scanEnd
}

// stateInString is the state after reading `"`.
func stateInString(s *scanner, c byte) OpCode {
	if c == '"' {
		s.step = stateEndValue
		return scanContinue
	}
	if c == '\\' {
		s.step = stateInStringEsc
		return scanContinue
	}
	if c < 0x20 {
		return s.error(c, "in string literal")
	}
	return scanContinue
}

// stateInStringEsc is the state after reading `"\` during a quoted string.
func stateInStringEsc(s *scanner, c byte) OpCode {
	switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
		s.step = stateInString
		return scanContinue
	case 'u':
		s.step = stateInStringEscU
		return scanContinue
	}
	return s.error(c, "in string escape code")
}

// stateInStringEscU is the state after reading `"\u` during a quoted string.
func stateInStringEscU(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInStringEscU1
		return scanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateInStringEscU1 is the state after reading `"\u1` during a quoted string.
func stateInStringEscU1(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInStringEscU12
		return scanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateInStringEscU12 is the state after reading `"\u12` during a quoted string.
func stateInStringEscU12(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInStringEscU123
		return scanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateInStringEscU123 is the state after reading `"\u123` during a quoted string.
func stateInStringEscU123(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInString
		return scanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateNeg is the state after reading `-` during a number.
func stateNeg(s *scanner, c byte) OpCode {
	if c == '0' {
		s.step = state0
		return scanContinue
	}
	if '1' <= c && c <= '9' {
		s.step = state1
		return scanContinue
	}
	return s.error(c, "in numeric literal")
}

// state1 is the state after reading a non-zero integer during a number,
// such as after reading `1` or `100` but not `0`.
func state1(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		s.step = state1
		return scanContinue
	}
	return state0(s, c)
}

// state0 is the state after reading `0` during a number.
func state0(s *scanner, c byte) OpCode {
	if c == '.' {
		s.step = stateDot
		return scanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		return scanContinue
	}
	return stateEndValue(s, c)
}

// stateDot is the state after reading the integer and decimal point in a number,
// such as after reading `1.`.
func stateDot(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		s.step = stateDot0
		return scanContinue
	}
	return s.error(c, "after decimal point in numeric literal")
}

// stateDot0 is the state after reading the integer, decimal point, and subsequent
// digits of a number, such as after reading `3.14`.
func stateDot0(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		return scanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		return scanContinue
	}
	return stateEndValue(s, c)
}

// stateE is the state after reading the mantissa and e in a number,
// such as after reading `314e` or `0.314e`.
func stateE(s *scanner, c byte) OpCode {
	if c == '+' || c == '-' {
		s.step = stateESign
		return scanContinue
	}
	return stateESign(s, c)
}

// stateESign is the state after reading the mantissa, e, and sign in a number,
// such as after reading `314e-` or `0.314e+`.
func stateESign(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		s.step = stateE0
		return scanContinue
	}
	return s.error(c, "in exponent of numeric literal")
}

// stateE0 is the state after reading the mantissa, e, optional sign,
// and at least one digit of the exponent in a number,
// such as after reading `314e-2` or `0.314e+1` or `3.14e0`.
func stateE0(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		return scanContinue
	}
	return stateEndValue(s, c)
}

// stateT is the state after reading `t`.
func stateT(s *scanner, c byte) OpCode {
	if c == 'r' {
		s.step = stateTr
		return scanContinue
	}
	return s.error(c, "in literal true (expecting 'r')")
}

// stateTr is the state after reading `tr`.
func stateTr(s *scanner, c byte) OpCode {
	if c == 'u' {
		s.step = stateTru
		return scanContinue
	}
	return s.error(c, "in literal true (expecting 'u')")
}

// stateTru is the state after reading `tru`.
func stateTru(s *scanner, c byte) OpCode {
	if c == 'e' {
		s.step = stateEndValue
		return scanContinue
	}
	return s.error(c, "in literal true (expecting 'e')")
}

// stateF is the state after reading `f`.
func stateF(s *scanner, c byte) OpCode {
	if c == 'a' {
		s.step = stateFa
		return scanContinue
	}
	return s.error(c, "in literal false (expecting 'a')")
}

// stateFa is the state after reading `fa`.
func stateFa(s *scanner, c byte) OpCode {
	if c == 'l' {
		s.step = stateFal
		return scanContinue
	}
	return s.error(c, "in literal false (expecting 'l')")
}

// stateFal is the state after reading `fal`.
func stateFal(s *scanner, c byte) OpCode {
	if c == 's' {
		s.step = stateFals
		return scanContinue
	}
	return s.error(c, "in literal false (expecting 's')")
}

// stateFals is the state after reading `fals`.
func stateFals(s *scanner, c byte) OpCode {
	if c == 'e' {
		s.step = stateEndValue
		return scanContinue
	}
	return s.error(c, "in literal false (expecting 'e')")
}

// stateN is the state after reading `n`.
func stateN(s *scanner, c byte) OpCode {
	if c == 'u' {
		s.step = stateNu
		return scanContinue
	}
	return s.error(c, "in literal null (expecting 'u')")
}

// stateNu is the state after reading `nu`.
func stateNu(s *scanner, c byte) OpCode {
	if c == 'l' {
		s.step = stateNul
		return scanContinue
	}
	return s.error(c, "in literal null (expecting 'l')")
}

// stateNul is the state after reading `nul`.
func stateNul(s *scanner, c byte) OpCode {
	if c == 'l' {
		s.step = stateEndValue
		return scanContinue
	}
	return s.error(c, "in literal null (expecting 'l')")
}

// stateError is the state after reaching a syntax error,
// such as after reading `[1}` or `5.1.2`.
func stateError(s *scanner, c byte) OpCode {
	return scanError
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) OpCode {
	s.step = stateError
	s.err = &SyntaxError{
		reason: "invalid character " + quoteChar(c) + " " + context,
		Offset: s.bytes,
	}
	return scanError
}

func isSpace(c byte) bool {
	return c <= ' ' && (c == ' ' || c == '\t' || c == '\r' || c == '\n')
}

// quoteChar formats c as a quoted character literal
func quoteChar(c byte) string {
	// special cases - different from quoted strings
	if c == '\'' {
		return `'\''`
	}
	if c == '"' {
		return `'"'`
	}

	// use quoted string with different quotation marks
	s := strconv.Quote(string(c))
	return "'" + s[1:len(s)-1] + "'"
}
