package basic

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (p *Player) handleKeepAlivePacket(packet pk.Packet) error {
	var KeepAliveID pk.Long
	if err := packet.Scan(&KeepAliveID); err != nil {
		return Error{err}
	}
	// Response
	err := p.c.Conn.WritePacket(pk.Packet{
		ID:   int32(packetid.ServerboundKeepAlive),
		Data: packet.Data,
	})
	if err != nil {
		return Error{err}
	}
	return nil
}

func (p *Player) handlePlayerPosition(packet pk.Packet) error {
	var (
		X, Y, Z         pk.Double
		Yaw, Pitch      pk.Float
		Flags           pk.Byte
		TeleportID      pk.VarInt
		DismountVehicle pk.Boolean
	)
	if err := packet.Scan(&X, &Y, &Z, &Yaw, &Pitch, &Flags, &TeleportID, &DismountVehicle); err != nil {
		return Error{err}
	}

	// Teleport Confirm
	err := p.c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundAcceptTeleportation,
		TeleportID,
	))
	if err != nil {
		return Error{err}
	}

	if !p.isSpawn {
		// PlayerPositionAndRotation to confirm the spawn position
		err = p.c.Conn.WritePacket(pk.Marshal(
			packetid.ServerboundMovePlayerPosRot,
			X, Y-1.62, Z,
			Yaw, Pitch,
			pk.Boolean(true),
		))
		if err != nil {
			return Error{err}
		}
		p.isSpawn = true
	}

	return nil
}
