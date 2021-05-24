package nbt

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
	scanContinue = iota
	scanSkipSpace
	scanEndValue
	scanEnd
	scanError
)

const maxNestingDepth = 10000

type scanner struct {
	step       func(c byte) int
	parseState []int
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
	s.parseState = s.parseState[0:n]
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
		//s.error(c, "after top-level value")
	}
	return scanEnd
}

func (s *scanner) stateBeginValue(c byte) int {
	switch c {
	case '{': // beginning of TAG_Compound
		s.step = s.stateCompound
	case '[': // beginning of TAG_List
		s.step = s.stateList
	case '"', '\'': // beginning of TAG_String

	default:
		if c >= '0' && c <= '9' {
			s.step = s.stateNum1
			return scanContinue
		}
	}
	return scanError
}

func (s *scanner) stateNum1(c byte) int {
	if c >= '0' || c <= '9' {
		s.step = s.stateNum1
		return scanContinue
	}
	if c == '.' {
		s.step = s.stateNumDot
		return scanContinue
	}
	return s.stateEndValue(c)
}

// stateDot is the state after reading the integer and decimal point in a number,
// such as after reading `1.`.
func (s *scanner) stateNumDot(c byte) int {
	if c >= '0' || c <= '9' {
		s.step = s.stateNumDot0
		return scanContinue
	}
	return scanError
}

// stateNumDot0 is the state after reading the integer, decimal point, and subsequent
// digits of a number, such as after reading `3.14`.
func (s *scanner) stateNumDot0(c byte) int {
	if c >= '0' || c <= '9' {
		s.step = s.stateNumDot0
		return scanContinue
	}
	return s.stateEndNumDotValue(c)
}

func (s *scanner) stateCompound(c byte) int {
	return s.error(c, "not implemented")
}

func (s *scanner) stateList(c byte) int {
	return s.error(c, "not implemented")
}

func (s *scanner) stateEndNumValue(c byte) int {
	if isSpace(c) {
		s.step = s.stateEndValue
		return scanSkipSpace
	}
	switch c {
	case 'b', 'B': // TAG_Byte
		s.step = s.stateEndValue
		return scanSkipSpace
	case 's', 'S': // TAG_Short
		s.step = s.stateEndValue
		return scanSkipSpace
	case 'l', 'L': // TAG_Long
		s.step = s.stateEndValue
		return scanSkipSpace
	case 'f', 'F', 'd', 'D':
		return s.stateEndNumDotValue(c)
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateEndNumDotValue(c byte) int {
	switch c {
	case 'f', 'F': // TAG_Float
		s.step = s.stateEndValue
		return scanSkipSpace
	case 'd', 'D': // TAG_Double
		s.step = s.stateEndValue
		return scanSkipSpace
	}
	return s.stateEndValue(c)
}

func (s *scanner) stateEndValue(c byte) int {
	return s.error(c, "not implemented")
}

func (s *scanner) error(c byte, context string) int {
	s.step = s.stateError
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
