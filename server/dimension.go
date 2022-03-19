package server

import (
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/level"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Level interface {
	Init(g *Game)
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
	columns      map[level.ChunkPos]*level.Chunk
}

func (s *SimpleDim) Init(*Game) {}

func NewSimpleDim(secs int) *SimpleDim {
	return &SimpleDim{
		numOfSection: secs,
		columns:      make(map[level.ChunkPos]*level.Chunk),
	}
}

func (s *SimpleDim) LoadChunk(pos level.ChunkPos, c *level.Chunk) {
	s.columns[pos] = c
}

func (s *SimpleDim) Info() LevelInfo {
	return LevelInfo{
		Name:       "minecraft:overworld",
		HashedSeed: 1234567,
	}
}

func (s *SimpleDim) PlayerJoin(p *Player) {
	for pos, column := range s.columns {
		packet := pk.Marshal(
			packetid.ClientboundLevelChunkWithLight,
			pos, column,
		)
		p.WritePacket(Packet758(packet))
	}
}

func (s *SimpleDim) PlayerQuit(*Player) {}
