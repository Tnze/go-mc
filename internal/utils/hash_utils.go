package utils

func HashString(s string) uint64 {
	h := uint64(5381) // choose a random initial value
	for i := 0; i < len(s); i++ {
		h = h*33 + uint64(s[i])
	}
	return h
}
