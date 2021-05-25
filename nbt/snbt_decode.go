package nbt

import (
	"errors"
)

type token int

const (
	ILLEGAL token = iota

	IDENT // name

	INT // 12345
	FLT // 12345.67

	BYTE   // b or B
	SHORT  // s or S
	LONG   // l or L
	FLOAT  // f or F
	DOUBLE // d or D

	STRING // "abc" 'def'

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
)

const (
	scanContinue        = iota // uninteresting byte
	scanBeginCompound          // begin TAG_Compound (after left-brace )
	scanBeginList              // begin TAG_List (after left-brack)
	scanListValue              // just finished read list value
	scanListType               // just finished read list type (after "B;" or "L;")
	scanCompoundTagName        // just finished read tag name (before colon)
	scanCompoundValue          // just finished read value (before comma or right-brace )
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
		//s.endTop = true
	} else {
		s.step = s.stateEndValue
	}
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
		return scanContinue
	default:
		if isNumber(c) {
			return s.stateNum1(c)
		}
		if isNumOrLetter(c) {
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
		s.step = s.stateInSqString
		return scanContinue
	case '"':
		s.step = s.stateInDqString
		return scanContinue
	default:
		if isNumOrLetter(c) {
			s.step = s.stateInPureString
			return scanContinue
		}
	}
	return s.error(c, "looking for beginning of string")
}

func (s *scanner) stateInSqString(c byte) int {
	if c == '\\' {
		s.step = s.stateInSqStringEsc
		return scanContinue
	}
	if c == '\'' {
		s.step = s.stateEndValue
		return scanContinue
	}
	if isNumOrLetter(c) {
		return scanContinue
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateInSqStringEsc(c byte) int {
	switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '\'':
		s.step = s.stateInSqString
		return scanContinue
	}
	return s.error(c, "in string escape code")
}

func (s *scanner) stateInDqString(c byte) int {
	if c == '\\' {
		s.step = s.stateInDqStringEsc
		return scanContinue
	}
	if c == '"' {
		s.step = s.stateEndValue
		return scanContinue
	}
	if isNumOrLetter(c) {
		return scanContinue
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateInDqStringEsc(c byte) int {
	switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
		s.step = s.stateInDqString
		return scanContinue
	}
	return s.error(c, "in string escape code")
}

func (s *scanner) stateInPureString(c byte) int {
	if isNumOrLetter(c) {
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
		return scanContinue
	case ']':
		return s.stateEndValue(c)
	default:
		return s.stateBeginValue(c)
	}
}

func (s *scanner) stateListOrArrayT(c byte) int {
	if c == ';' {
		s.step = s.stateBeginValue
		return scanListType
	}
	return s.stateInPureString(c)
}

func (s *scanner) stateNeg(c byte) int {
	if !isNumber(c) {
		s.error(c, "not a number after '-'")
	}
	s.step = s.stateNum1
	return scanContinue
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
	if isSpace(c) {
		s.step = s.stateEndValue
		return scanSkipSpace
	}
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
		//s.endTop = true
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
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isNumOrLetter(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || isNumber(c) {
		return true
	}
	return false
}
