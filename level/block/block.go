package block

type Block interface {
	ID() string
}

var toStateID = make(map[Block]int)
var fromStateID []Block
var fromID = make(map[string]Block)

func init() {
	//regState := func(s Block) {
	//	if _, ok := toStateID[s]; ok {
	//		panic(fmt.Errorf("state %#v already exist", s))
	//	}
	//	toStateID[s] = len(fromStateID)
	//	fromStateID = append(fromStateID, s)
	//}
	//regBlock := func(b Block) {
	//	fromID[b.ID()] = b
	//	b.forEachState(regState)
	//}
	//regBlock(Air{})
}

func NewFromStateID(stateID int) Block {
	return fromStateID[stateID]
}
