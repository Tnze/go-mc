package phy

// Inputs describes the desired movements of the player.
type Inputs struct {
	Yaw, Pitch float32

	ThrottleX, ThrottleZ float64
}
