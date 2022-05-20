package block

func IsAir(s StateID) bool {
	switch StateList[s].(type) {
	case Air, CaveAir, VoidAir:
		return true
	default:
		return false
	}
}
