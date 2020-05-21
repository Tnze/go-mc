package world

import (
	"github.com/Tnze/go-mc/bot/world/entity"
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
	GetBlock(offset int) BlockStatus
	// SetBlock is the reverse operation of GetBlock.
	SetBlock(offset int, s BlockStatus)
}

func SectionOffset(x, y, z int) (offset int) {
	// According to wiki.vg: Data Array is given for each block with increasing x coordinates,
	// within rows of increasing z coordinates, within layers of increasing y coordinates.
	// So offset equals to ( x*16^0 + z*16^1 + y*16^2 )*(bits per block).
	return x + z*16 + y*16*16
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
	if c != nil {
		// (n&(16-1)) == (n<0 ? n%16+16 : n%16)
		if sec := c.Sections[y>>4]; sec != nil {
			return sec.GetBlock(SectionOffset(x&15, y&15, z&15))
		}
	}
	return 0
}

// func (b Block) String() string {
// 	return blockNameByID[b.id]
// }

//LoadChunk load chunk at (x, z)
func (w *World) LoadChunk(x, z int, c *Chunk) {
	w.Chunks[ChunkLoc{X: x, Z: z}] = c
}
