package world

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/maths"
	. "github.com/Tnze/go-mc/level"
)

type World struct {
	Columns map[ChunkPos]*Chunk
}

func NewWorld() (w *World) {
	w = &World{
		Columns: make(map[ChunkPos]*Chunk),
	}
	return
}

func (w *World) GetBlock(pos maths.Vec3d) (BlocksState, error) {
	if int32(pos.Y) < 0 || int32(pos.Y) > 320 {
		return -1, fmt.Errorf("out of range")
	}
	chunkPos := ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		return chunk.GetBlock(pos)
	} else {
		return -1, fmt.Errorf("chunk not loaded")
	}
}

/*func (w *World) onPlayerSpawn(pk.Packet) error {
	// unload all chunks
	w.Columns = make(map[level.ChunkPos]*level.Chunk)
	return nil
}

func (w *World) handleLevelChunkWithLightPacket(packet pk.Packet) error {
	var pos level.ChunkPos
	currentDimType := w.p.WorldInfo.DimensionCodec.DimensionType.Find(w.p.DimensionType)
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
}*/
