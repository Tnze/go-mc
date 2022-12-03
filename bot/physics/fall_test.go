package physics

import (
	"github.com/Tnze/go-mc/bot/maths"
	"testing"
)

func Test_Falling(t *testing.T) {
	ticks := float32(0.0)
	position := maths.Vec3d{X: 0, Y: 256, Z: 0}
	velocity := maths.NullVec3d
	for position.Y > 0 {
		velocity = maths.CalculateFallVelocity(ticks)
		position = position.Add(velocity)
		ticks++
		nextVelocity := maths.CalculateFallVelocity(ticks + 1)
		if position.Add(nextVelocity).Y < 0 {
			position = maths.NullVec3d
		}
		t.Logf("Ticks: %v, Position: %v, Velocity: %v", ticks, position, velocity)
	}
}
