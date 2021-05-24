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
	scanError
)

type scanner struct {
	step       func(c byte) int
	parseState []int
}

func (s *scanner) stateBeginValue(c byte) int {
	switch c {
	case '{': // beginning of TAG_Compound
		s.step = s.stateCompound
	case '[': // beginning of TAG_List
		s.step = s.stateList
	case '"', '\'': // beginning of TAG_String

	}
	return scanError
}

func (p *scanner) stateCompound(c byte) int {
	return scanError
}

func (p *scanner) stateList(c byte) int {
	return scanError
}
