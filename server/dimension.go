package server

import (
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/level"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Level interface {
	Info() LevelInfo
	PlayerJoin(p *Player)
	PlayerQuit(p *Player)
}

type LevelInfo struct {
	Name       string
	HashedSeed uint64
}

type SimpleDim struct {
	numOfSection int
	Columns      map[level.ChunkPos]*level.Chunk
}

func NewSimpleDim(secs int) *SimpleDim {
	return &SimpleDim{
		numOfSection: secs,
		Columns:      make(map[level.ChunkPos]*level.Chunk),
	}
}

func (s *SimpleDim) LoadChunk(pos level.ChunkPos, c *level.Chunk) {
	s.Columns[pos] = c
}

func (s *SimpleDim) Info() LevelInfo {
	return LevelInfo{
		Name:       "minecraft:overworld",
		HashedSeed: 1234567,
	}
}

func (s *SimpleDim) PlayerJoin(p *Player) {
	for pos, column := range s.Columns {
		column.Lock()
		packet := pk.Marshal(
			packetid.ClientboundLevelChunkWithLight,
			pk.Int(pos.X), pk.Int(pos.Z),
			column,
		)
		column.Unlock()

		p.WritePacket(Packet758(packet))
	}

	p.WritePacket(Packet758(pk.Marshal(
		packetid.ClientboundPlayerPosition,
		pk.Double(0), pk.Double(143), pk.Double(0),
		pk.Float(0), pk.Float(0),
		pk.Byte(0),
		pk.VarInt(0),
		pk.Boolean(true),
	)))
}

func (s *SimpleDim) PlayerQuit(p *Player) {

}
