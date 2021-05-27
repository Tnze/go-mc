package nbt

import (
	"errors"
)

const (
	scanContinue        = iota // uninteresting byte
	scanBeginLiteral           // end implied by next result != scanContinue
	scanBeginCompound          // begin TAG_Compound (after left-brace )
	scanBeginList              // begin TAG_List (after left-brack)
	scanListValue              // just finished read list value (after comma)
	scanListType               // just finished read list type (after "B;" or "L;")
	scanCompoundTagName        // just finished read tag name (before colon)
	scanCompoundValue          // just finished read value (after comma)
	scanSkipSpace              // space byte; can skip; known to be last "continue" result
	scanEndValue

	scanEnd
	scanError
)

// These values are stored in the parseState stack.
// They give the current state of a composite value
// being scanned. If the parser is inside a nested value
// the parseState describes the nested state, outermost at entry 0.
const (
	parseCompoundName  = iota // parsing tag name (before colon)
	parseCompoundValue        // parsing value (after colon)
	parseListValue            // parsing list
)

const maxNestingDepth = 10000

type scanner struct {
	step       func(c byte) int
	parseState []int
	err        error
	endTop     bool
}

// reset prepares the scanner for use.
// It must be called before calling s.step.
func (s *scanner) reset() {
	s.step = s.stateBeginValue
	s.parseState = s.parseState[0:0]
}

// pushParseState pushes a new parse state p onto the parse stack.
// an error state is returned if maxNestingDepth was exceeded, otherwise successState is returned.
func (s *scanner) pushParseState(c byte, newParseState int, successState int) int {
	s.parseState = append(s.parseState, newParseState)
	if len(s.parseState) <= maxNestingDepth {
		return successState
	}
	return scanError
}

// popParseState pops a parse state (already obtained) off the stack
// and updates s.step accordingly.
func (s *scanner) popParseState() {
	n := len(s.parseState) - 1
	s.parseState = s.parseState[:n]
	if n == 0 {
		s.step = s.stateEndTop
		s.endTop = true
	} else {
		s.step = s.stateEndValue
	}
}

// eof tells the scanner that the end of input has been reached.
// It returns a scan status just as s.step does.
func (s *scanner) eof() int {
	if s.err != nil {
		return scanError
	}
	if s.endTop {
		return scanEnd
	}
	s.step(' ')
	if s.endTop {
		return scanEnd
	}
	if s.err == nil {
		s.err = errors.New("unexpected end of JSON input")
	}
	return scanError
}

// stateEndTop is the state after finishing the top-level value,
// such as after reading `{}` or `[1,2,3]`.
// Only space characters should be seen now.
func (s *scanner) stateEndTop(c byte) int {
	if !isSpace(c) {
		// Complain about non-space byte on next call.
		s.error(c, "after top-level value")
	}
	return scanEnd
}

func (s *scanner) stateBeginValue(c byte) int {
	if isSpace(c) {
		s.step = s.stateBeginValue
		return scanSkipSpace
	}
	switch c {
	case '{': // beginning of TAG_Compound
		s.step = s.stateCompoundOrEmpty
		return s.pushParseState(c, parseCompoundName, scanBeginCompound)
	case '[': // beginning of TAG_List
		s.step = s.stateListOrArray
		return s.pushParseState(c, parseListValue, scanBeginList)
	case '"', '\'': // beginning of TAG_String
		return s.stateBeginString(c)
	case '-': // beginning of negative number
		s.step = s.stateNeg
		return scanBeginLiteral
	default:
		if isNumber(c) {
			s.stateNum0(c)
			return scanBeginLiteral
		}
		if isAllowedInUnquotedString(c) {
			return s.stateBeginString(c)
		}
	}
	return s.error(c, "looking for beginning of value")
}

func (s *scanner) stateCompoundOrEmpty(c byte) int {
	if isSpace(c) {
		return scanSkipSpace
	}
	if c == '}' {
		n := len(s.parseState)
		s.parseState[n-1] = parseCompoundValue
		return s.stateEndValue(c)
	}
	return s.stateBeginString(c)
}

func (s *scanner) stateBeginString(c byte) int {
	if isSpace(c) {
		return scanSkipSpace
	}
	switch c {
	case '\'':
		s.step = s.stateInSingleQuotedString
		return scanBeginLiteral
	case '"':
		s.step = s.stateInDoubleQuotedString
		return scanBeginLiteral
	default:
		if isAllowedInUnquotedString(c) {
			s.step = s.stateInUnquotedString
			return scanBeginLiteral
		}
	}
	return s.error(c, "looking for beginning of string")
}

func (s *scanner) stateInSingleQuotedString(c byte) int {
	if c == '\\' {
		s.step = s.stateInSingleQuotedStringEsc
		return scanContinue
	}
	if c == '\'' {
		s.step = s.stateEndValue
		return scanContinue
	}
	return scanContinue
}

func (s *scanner) stateInSingleQuotedStringEsc(c byte) int {
	switch c {
	case '\\', '\'':
		s.step = s.stateInSingleQuotedString
		return scanContinue
	}
	return s.error(c, "in string escape code")
}

func (s *scanner) stateInDoubleQuotedString(c byte) int {
	if c == '\\' {
		s.step = s.stateInDqStringEsc
		return scanContinue
	}
	if c == '"' {
		s.step = s.stateEndValue
		return scanContinue
	}
	return scanContinue
}

func (s *scanner) stateInDqStringEsc(c byte) int {
	switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
		s.step = s.stateInDoubleQuotedString
		return scanContinue
	}
	return s.error(c, "in string escape code")
}

func (s *scanner) stateInUnquotedString(c byte) int {
	if isAllowedInUnquotedString(c) {
		return scanContinue
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateListOrArray(c byte) int {
	if isSpace(c) {
		return scanSkipSpace
	}
	switch c {
	case 'B', 'I', 'L':
		s.step = s.stateListOrArrayT
		return scanBeginLiteral
	case ']':
		return s.stateEndValue(c)
	default:
		return s.stateBeginValue(c)
	}
}

func (s *scanner) stateListOrArrayT(c byte) int {
	if c == ';' {
		s.step = s.stateArrayT
		return scanListType
	}
	return s.stateInUnquotedString(c)
}

func (s *scanner) stateArrayT(c byte) int {
	if isSpace(c) {
		return scanSkipSpace
	}
	if c == ']' { // empty array
		return scanEndValue
	}
	return s.stateBeginValue(c)
}

func (s *scanner) stateNeg(c byte) int {
	if isNumber(c) {
		s.step = s.stateNum0
		return scanBeginLiteral
	}
	if isAllowedInUnquotedString(c) {
		s.step = s.stateInUnquotedString
		return scanBeginLiteral
	}
	return s.error(c, "not a number after '-'")
}

func (s *scanner) stateNum0(c byte) int {
	if isNumber(c) {
		s.step = s.stateNum1
		return scanContinue
	}
	return s.stateEndNumValue(c)
}

func (s *scanner) stateNum1(c byte) int {
	if isNumber(c) {
		s.step = s.stateNum1
		return scanContinue
	}
	if c == '.' {
		s.step = s.stateNumDot
		return scanContinue
	}
	return s.stateEndNumValue(c)
}

// stateDot is the state after reading the integer and decimal point in a number,
// such as after reading `1.`.
func (s *scanner) stateNumDot(c byte) int {
	if isNumber(c) {
		s.step = s.stateNumDot0
		return scanContinue
	}
	if isAllowedInUnquotedString(c) {
		s.step = s.stateInUnquotedString
		return scanContinue
	}
	return s.error(c, "after decimal point in numeric literal")
}

// stateNumDot0 is the state after reading the integer, decimal point, and subsequent
// digits of a number, such as after reading `3.14`.
func (s *scanner) stateNumDot0(c byte) int {
	if isNumber(c) {
		s.step = s.stateNumDot0
		return scanContinue
	}
	return s.stateEndNumDotValue(c)
}

func (s *scanner) stateEndNumValue(c byte) int {
	switch c {
	case 'b', 'B': // TAG_Byte
		s.step = s.stateEndValue
		return scanContinue
	case 's', 'S': // TAG_Short
		s.step = s.stateEndValue
		return scanContinue
	case 'l', 'L': // TAG_Long
		s.step = s.stateEndValue
		return scanContinue
	case 'f', 'F', 'd', 'D':
		return s.stateEndNumDotValue(c)
	}
	if isAllowedInUnquotedString(c) {
		s.step = s.stateInUnquotedString
		return scanContinue
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateEndNumDotValue(c byte) int {
	switch c {
	case 'f', 'F': // TAG_Float
		s.step = s.stateEndValue
		return scanContinue
	case 'd', 'D': // TAG_Double
		s.step = s.stateEndValue
		return scanContinue
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateEndValue(c byte) int {
	n := len(s.parseState)
	if n == 0 {
		// Completed top-level before the current byte.
		s.step = s.stateEndTop
		s.endTop = true
		return s.stateEndTop(c)
	}
	if isSpace(c) {
		return scanSkipSpace
	}

	ps := s.parseState[n-1]
	switch ps {
	case parseCompoundName:
		if c == ':' {
			s.parseState[n-1] = parseCompoundValue
			s.step = s.stateBeginValue
			return scanCompoundTagName
		}
		return s.error(c, "after compound tag name")
	case parseCompoundValue:
		switch c {
		case ',':
			s.parseState[n-1] = parseCompoundName
			s.step = s.stateBeginString
			return scanCompoundValue
		case '}':
			s.popParseState()
			return scanEndValue
		}
		return s.error(c, "after compound value")
	case parseListValue:
		switch c {
		case ',':
			s.step = s.stateBeginValue
			return scanListValue
		case ']':
			s.popParseState()
			return scanEndValue
		}
		return s.error(c, "after list element")
	}
	return s.error(c, "")
}

func (s *scanner) error(c byte, context string) int {
	s.step = s.stateError
	s.err = errors.New(context)
	return scanError
}

// stateError is the state after reaching a syntax error,
// such as after reading `[1}` or `5.1.2`.
func (s *scanner) stateError(c byte) int {
	return scanError
}

func isSpace(c byte) bool {
	return c <= ' ' && (c == ' ' || c == '\t' || c == '\r' || c == '\n')
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAllowedInUnquotedString(c byte) bool {
	return c == '_' || c == '-' ||
		c == '.' || c == '+' ||
		c >= '0' && c <= '9' ||
		c >= 'A' && c <= 'Z' ||
		c >= 'a' && c <= 'z'
}
