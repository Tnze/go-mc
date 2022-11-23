package world

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/level"
	pk "github.com/Tnze/go-mc/net/packet"
)

type World struct {
	c      *bot.Client
	p      *basic.Player
	events EventsListener

	Columns map[level.ChunkPos]*level.Chunk
}

func NewWorld(c *bot.Client, p *basic.Player, events EventsListener) (w *World) {
	w = &World{
		c: c, p: p,
		events:  events,
		Columns: make(map[level.ChunkPos]*level.Chunk),
	}
	c.Events.AddListener(
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundLogin, F: w.onPlayerSpawn},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundRespawn, F: w.onPlayerSpawn},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundLevelChunkWithLight, F: w.handleLevelChunkWithLightPacket},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundForgetLevelChunk, F: w.handleForgetLevelChunkPacket},
	)
	return
}

func (w *World) onPlayerSpawn(pk.Packet) error {
	// unload all chunks
	w.Columns = make(map[level.ChunkPos]*level.Chunk)
	return nil
}

func (w *World) handleLevelChunkWithLightPacket(packet pk.Packet) error {
	var pos level.ChunkPos
	currentDimType := w.p.WorldInfo.RegistryCodec.DimensionType.Find(w.p.DimensionType)
	chunk := level.EmptyChunk(int(currentDimType.Height) / 16)
	if err := packet.Scan(&pos, chunk); err != nil {
		return err
	}
	w.Columns[pos] = chunk
	if w.events.LoadChunk != nil {
		if err := w.events.LoadChunk(pos); err != nil {
			return err
		}
	}
	return nil
}

func (w *World) handleForgetLevelChunkPacket(packet pk.Packet) error {
	var pos level.ChunkPos
	if err := packet.Scan(&pos); err != nil {
		return err
	}
	var err error
	if w.events.UnloadChunk != nil {
		err = w.events.UnloadChunk(pos)
	}
	delete(w.Columns, pos)
	return err
}
