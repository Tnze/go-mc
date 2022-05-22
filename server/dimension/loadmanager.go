package dimension

import (
	"container/list"

	"github.com/Tnze/go-mc/level"
)

type manager struct {
	storage
	elements     map[level.ChunkPos]*list.Element
	activeChunks *list.List

	chunkLoadQueue   []level.ChunkPos
	chunkUnloadQueue []level.ChunkPos
}

type chunkHandler struct {
	level.ChunkPos
	*level.Chunk
}

func (m *manager) refresh(players [][2]int) {
	m.chunkLoadQueue = m.chunkLoadQueue[:0]
	m.chunkUnloadQueue = m.chunkUnloadQueue[:0]

	newActives := list.New()
	const N = 16
	for _, p := range players {
		for i := 1 - N; i < N; i++ {
			for j := 1 - N; j < N; j++ {
				pos := level.ChunkPos{
					X: p[0] + i,
					Z: p[1] + j,
				}
				if e := m.elements[pos]; e != nil {
					// chunk exist, move into newActives
					v := m.activeChunks.Remove(e)
					m.elements[pos] = newActives.PushBack(v)
				} else {
					// not exist, load from storage
					m.chunkLoadQueue = append(m.chunkLoadQueue, pos)
				}
			}
		}
	}
	for e := m.activeChunks.Front(); e != nil; e = e.Next() {
		pos := e.Value.(chunkHandler).ChunkPos
		m.chunkUnloadQueue = append(m.chunkUnloadQueue, pos)
		m.elements[pos] = newActives.PushBack(e.Value)
	}
	m.activeChunks = newActives
}
