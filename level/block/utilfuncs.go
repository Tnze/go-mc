package block

func IsAir(s int) bool {
	switch StateList[s].(type) {
	case Air, CaveAir, VoidAir:
		return true
	default:
		return false
	}
}
