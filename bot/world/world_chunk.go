package world

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/data/block"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/net/ptypes"
)

// Chunk store a 256*16*16 area of blocks, sharded on the Y axis into 16
// sections.
type Chunk struct {
	Sections     [16]Section
	TileEntities map[TilePosition]entity.BlockEntity
}

// Section implements storage of blocks within a fixed 16*16*16 area.
type Section interface {
	// GetBlock return block status, offset can be calculate by SectionOffset.
	GetBlock(offset uint) BlockStatus
	// SetBlock is the reverse operation of GetBlock.
	SetBlock(offset uint, s BlockStatus)
}

func sectionIdx(x, y, z int) uint {
	return uint(((y & 15) << 8) | (z << 4) | x)
}

type BlockStatus uint32

type ChunkLoc struct {
	X, Z int
}

// Signs returns a list of signs on the server within viewing range.
func (w *World) Signs() []entity.BlockEntity {
	w.chunkLock.RLock()
	defer w.chunkLock.RUnlock()

	out := make([]entity.BlockEntity, 0, 4)
	for _, c := range w.Chunks {
		for _, e := range c.TileEntities {
			if e.ID == "minecraft:sign" {
				out = append(out, e)
			}
		}
	}

	return out
}

// GetBlockStatus return the state ID of the block at the given position.
func (w *World) GetBlockStatus(x, y, z int) BlockStatus {
	w.chunkLock.RLock()
	defer w.chunkLock.RUnlock()

	c := w.Chunks[ChunkLoc{x >> 4, z >> 4}]
	if c != nil && y >= 0 {
		if sec := c.Sections[y>>4]; sec != nil {
			return sec.GetBlock(sectionIdx(x&15, y&15, z&15))
		}
	}
	return 0
}

// UnloadChunk unloads a chunk from the world.
func (w *World) UnloadChunk(loc ChunkLoc) {
	w.chunkLock.Lock()
	delete(w.Chunks, loc)
	w.chunkLock.Unlock()
}

// UnaryBlockUpdate updates the block at the specified position with a new
// state ID.
func (w *World) UnaryBlockUpdate(pos pk.Position, bStateID BlockStatus) bool {
	w.chunkLock.Lock()
	defer w.chunkLock.Unlock()

	c := w.Chunks[ChunkLoc{X: pos.X >> 4, Z: pos.Z >> 4}]
	if c == nil {
		return false
	}
	sIdx, bIdx := pos.Y>>4, sectionIdx(pos.X&15, pos.Y&15, pos.Z&15)

	if sec := c.Sections[sIdx]; sec == nil {
		sec = newSectionWithSize(uint(block.BitsPerBlock))
		sec.SetBlock(bIdx, bStateID)
		c.Sections[sIdx] = sec
	} else {
		tp := ToTilePos(pos.X, pos.Y, pos.Z)
		if _, ok := c.TileEntities[tp]; ok && sec.GetBlock(bIdx) != bStateID {
			delete(c.TileEntities, tp)
		}
		sec.SetBlock(bIdx, bStateID)
	}
	return true
}

// MultiBlockUpdate updates a packed specification of blocks within a single
// section of a chunk.
func (w *World) MultiBlockUpdate(loc ChunkLoc, sectionY int, blocks []pk.VarLong) bool {
	w.chunkLock.Lock()
	defer w.chunkLock.Unlock()

	c := w.Chunks[loc]
	if c == nil {
		return false // not loaded
	}

	sec := c.Sections[sectionY]
	if sec == nil {
		sec = newSectionWithSize(uint(block.BitsPerBlock))
		c.Sections[sectionY] = sec
	}

	for _, b := range blocks {
		var (
			bStateID = BlockStatus(b >> 12)
			x, z, y  = (b >> 8) & 0xf, (b >> 4) & 0xf, b & 0xf
			bIdx     = sectionIdx(int(x&15), int(y&15), int(z&15))
			tp       = ToTilePos(int(x), int(y), int(z))
		)
		if _, ok := c.TileEntities[tp]; ok && sec.GetBlock(bIdx) != bStateID {
			delete(c.TileEntities, tp)
		}
		sec.SetBlock(bIdx, bStateID)
	}

	return true
}

//LoadChunk adds the given chunk to the world.
func (w *World) LoadChunk(x, z int, c *Chunk) {
	w.chunkLock.Lock()
	w.Chunks[ChunkLoc{X: x, Z: z}] = c
	w.chunkLock.Unlock()
}

func (w *World) TileEntityUpdate(pkt ptypes.TileEntityData) error {
	w.chunkLock.Lock()
	defer w.chunkLock.Unlock()
	c := w.Chunks[ChunkLoc{X: pkt.Pos.X >> 4, Z: pkt.Pos.Z >> 4}]
	if c == nil {
		return nil
	}

	switch pkt.Action {
	case 9: // Sign update
		c.TileEntities[ToTilePos(pkt.Pos.X, pkt.Pos.Y, pkt.Pos.Z)] = pkt.Data
	}
	return nil
}
