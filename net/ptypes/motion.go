// Package ptypes implements encoding and decoding for high-level packets.
package ptypes

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// PositionAndLookClientbound describes the location and orientation of
// the player.
type PositionAndLookClientbound struct {
	X, Y, Z    pk.Double
	Yaw, Pitch pk.Float
	Flags      pk.Byte
	TeleportID pk.VarInt
}

func (p *PositionAndLookClientbound) RelativeX() bool {
	return p.Flags&0x01 != 0
}
func (p *PositionAndLookClientbound) RelativeY() bool {
	return p.Flags&0x02 != 0
}
func (p *PositionAndLookClientbound) RelativeZ() bool {
	return p.Flags&0x04 != 0
}
func (p *PositionAndLookClientbound) RelativeYaw() bool {
	return p.Flags&0x08 != 0
}
func (p *PositionAndLookClientbound) RelativePitch() bool {
	return p.Flags&0x10 != 0
}

func (p *PositionAndLookClientbound) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.X, &p.Y, &p.Z, &p.Yaw, &p.Pitch, &p.Flags, &p.TeleportID)
}

// PositionAndLookServerbound describes the location and orientation of
// the player.
type PositionAndLookServerbound struct {
	X, Y, Z    pk.Double
	Yaw, Pitch pk.Float
	OnGround   pk.Boolean
}

func (p PositionAndLookServerbound) Encode() pk.Packet {
	return pk.Marshal(
		packetid.PositionLook,
		p.X, p.Y, p.Z,
		p.Yaw, p.Pitch,
		p.OnGround,
	)
}

// Position describes the position of the player.
type Position struct {
	X, Y, Z  pk.Double
	OnGround pk.Boolean
}

func (p Position) Encode() pk.Packet {
	return pk.Marshal(
		packetid.PositionServerbound,
		p.X, p.Y, p.Z,
		p.OnGround,
	)
}

// Look describes the rotation of the player.
type Look struct {
	Yaw, Pitch pk.Float
	OnGround   pk.Boolean
}

func (p Look) Encode() pk.Packet {
	return pk.Marshal(
		packetid.Look,
		p.Yaw, p.Pitch,
		p.OnGround,
	)
}
