package world

import (
	"sync"

	"github.com/Tnze/go-mc/bot/world/entity"
)

// World record all of the things in the world where player at
type World struct {
	entityLock sync.RWMutex
	Entities   map[int32]*entity.Entity
	chunkLock  sync.RWMutex
	Chunks     map[ChunkLoc]*Chunk
}
