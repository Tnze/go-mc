package world

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/data/block"
	pk "github.com/Tnze/go-mc/net/packet"
)

// World record all of the things in the world where player at
type World struct {
	Entities map[int32]entity.Entity
	Chunks   map[ChunkLoc]*Chunk
}

// Chunk store a 256*16*16 column blocks
type Chunk struct {
	Sections [16]Section
}

// Section store a 16*16*16 cube blocks
type Section interface {
	// GetBlock return block status, offset can be calculate by SectionOffset.
	GetBlock(offset uint) BlockStatus
	// SetBlock is the reverse operation of GetBlock.
	SetBlock(offset uint, s BlockStatus)
}

func SectionIdx(x, y, z int) uint {
	// According to wiki.vg: Data Array is given for each block with increasing x coordinates,
	// within rows of increasing z coordinates, within layers of increasing y coordinates.
	// So offset equals to ( x*16^0 + z*16^1 + y*16^2 )*(bits per block).
	return uint(((y & 15) << 8) | (z << 4) | x)
}

type BlockStatus uint32

type ChunkLoc struct {
	X, Z int
}

// //Entity 表示一个实体
// type Entity interface {
// 	EntityID() int32
// }

// //Face is a face of a block
// type Face byte

// // All six faces in a block
// const (
// 	Bottom Face = iota
// 	Top
// 	North
// 	South
// 	West
// 	East
// )

// getBlock return the block in the position (x, y, z)
func (w *World) GetBlockStatus(x, y, z int) BlockStatus {
	// Use n>>4 rather then n/16. It acts wrong if n<0.
	c := w.Chunks[ChunkLoc{x >> 4, z >> 4}]
	if c != nil && y >= 0 {
		if sec := c.Sections[y>>4]; sec != nil {
			return sec.GetBlock(SectionIdx(x&15, y&15, z&15))
		}
	}
	return 0
}

func (w *World) UnaryBlockUpdate(pos pk.Position, bStateID BlockStatus) bool {
	c := w.Chunks[ChunkLoc{X: pos.X >> 4, Z: pos.Z >> 4}]
	if c == nil {
		return false
	}
	sIdx, bIdx := pos.Y>>4, SectionIdx(pos.X&15, pos.Y&15, pos.Z&15)

	if sec := c.Sections[sIdx]; sec == nil {
		sec = newSectionWithSize(uint(block.BitsPerBlock))
		sec.SetBlock(bIdx, bStateID)
		c.Sections[sIdx] = sec
	} else {
		sec.SetBlock(bIdx, bStateID)
	}
	return true
}

func (w *World) MultiBlockUpdate(loc ChunkLoc, sectionY int, blocks []pk.VarLong) bool {
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
		bStateID := b >> 12
		x, z, y := (b>>8)&0xf, (b>>4)&0xf, b&0xf
		sec.SetBlock(SectionIdx(int(x&15), int(y&15), int(z&15)), BlockStatus(bStateID))
	}

	return true
}

// func (b Block) String() string {
// 	return blockNameByID[b.id]
// }

//LoadChunk load chunk at (x, z)
func (w *World) LoadChunk(x, z int, c *Chunk) {
	w.Chunks[ChunkLoc{X: x, Z: z}] = c
}
