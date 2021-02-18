package ptypes

import pk "github.com/Tnze/go-mc/net/packet"

// SpawnEntity is a client-bound packet used to spawn a non-mob entity.
type SpawnEntity struct {
	ID               pk.VarInt
	UUID             pk.UUID
	Type             pk.VarInt
	X, Y, Z          pk.Double
	Pitch, Yaw       pk.Angle
	Data             pk.Int
	VelX, VelY, VelZ pk.Short
}

func (p *SpawnEntity) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.UUID, &p.Type,
		&p.X, &p.Y, &p.Z, &p.Pitch, &p.Yaw,
		&p.Data, &p.VelX, &p.VelY, &p.VelZ)
}

// SpawnPlayer is a client-bound packet used to describe a player entering
// visible range.
type SpawnPlayer struct {
	ID         pk.VarInt
	UUID       pk.UUID
	X, Y, Z    pk.Double
	Yaw, Pitch pk.Angle
}

func (p *SpawnPlayer) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.UUID, &p.X, &p.Y, &p.Z, &p.Yaw, &p.Pitch)
}

// SpawnLivingEntity is a client-bound packet used to spawn a mob.
type SpawnLivingEntity struct {
	ID               pk.VarInt
	UUID             pk.UUID
	Type             pk.VarInt
	X, Y, Z          pk.Double
	Yaw, Pitch       pk.Angle
	HeadPitch        pk.Angle
	VelX, VelY, VelZ pk.Short
}

func (p *SpawnLivingEntity) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.UUID, &p.Type,
		&p.X, &p.Y, &p.Z, &p.Yaw, &p.Pitch,
		&p.HeadPitch, &p.VelX, &p.VelY, &p.VelZ)
}

// EntityAnimationClientbound updates the animation state of an entity.
type EntityAnimationClientbound struct {
	ID        pk.VarInt
	Animation pk.UnsignedByte
}

func (p *EntityAnimationClientbound) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.Animation)
}

// EntityPosition is a clientbound packet used to update an entities position.
type EntityPosition struct {
	ID       pk.VarInt
	X, Y, Z  pk.Short // Deltas
	OnGround pk.Boolean
}

func (p *EntityPosition) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.X, &p.Y, &p.Z, &p.OnGround)
}

// EntityPosition is a clientbound packet used to update an entities position
// and its rotation.
type EntityPositionLook struct {
	ID         pk.VarInt
	X, Y, Z    pk.Short // Deltas
	Yaw, Pitch pk.Angle
	OnGround   pk.Boolean
}

func (p *EntityPositionLook) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.X, &p.Y, &p.Z, &p.Yaw, &p.Pitch, &p.OnGround)
}

// EntityRotation is a clientbound packet used to update an entities rotation.
type EntityRotation struct {
	ID         pk.VarInt
	Yaw, Pitch pk.Angle
	OnGround   pk.Boolean
}

func (p *EntityRotation) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.ID, &p.Yaw, &p.Pitch, &p.OnGround)
}
