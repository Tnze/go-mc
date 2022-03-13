package server

type EntitySet struct {
}

type entityPosition struct {
	player     *Player
	x, y, z    float64
	yaw, pitch float32
}

func NewEntitySet() *EntitySet {
	return &EntitySet{}
}
