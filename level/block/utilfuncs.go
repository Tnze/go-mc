package block

func IsAir(s StateID) bool {
	return IsAirBlock(StateList[s])
}

func IsAirBlock(b Block) bool {
	switch b.(type) {
	case Air, CaveAir, VoidAir:
		return true
	default:
		return false
	}
}
