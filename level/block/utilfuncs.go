package block

func IsAir(s StateID) bool {
	switch StateList[s].IBlock {
	case Air{}, CaveAir{}, VoidAir{}:
		return true
	default:
		return false
	}
}
