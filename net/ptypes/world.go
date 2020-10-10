package ptypes

import (
	pk "github.com/Tnze/go-mc/net/packet"
)

// JoinGame encodes global/world information from the server.
type JoinGame struct {
	PlayerEntity pk.Int
	Hardcore     pk.Boolean
	Gamemode     pk.UnsignedByte
	PrevGamemode pk.UnsignedByte
	WorldCount   pk.VarInt
	WorldNames   pk.Identifier
	//DimensionCodec pk.NBT
	Dimension    pk.Int
	WorldName    pk.Identifier
	HashedSeed   pk.Long
	maxPlayers   pk.VarInt // Now ignored
	ViewDistance pk.VarInt
	RDI          pk.Boolean // Reduced Debug Info
	ERS          pk.Boolean // Enable respawn screen
	IsDebug      pk.Boolean
	IsFlat       pk.Boolean
}

func (p *JoinGame) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.PlayerEntity, &p.Hardcore, &p.Gamemode, &p.PrevGamemode,
		&p.WorldCount, &p.WorldNames, &p.Dimension,
		&p.WorldName, &p.HashedSeed, &p.maxPlayers, &p.ViewDistance,
		&p.RDI, &p.ERS, &p.IsDebug, &p.IsFlat)
}
