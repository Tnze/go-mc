package dimension

import (
	"container/list"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/level"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server"
	"github.com/Tnze/go-mc/server/clientinfo"
	"github.com/Tnze/go-mc/server/internal/bvh"
)

type vec2d = bvh.Vec2[float64]
type sphere2d = bvh.Sphere[float64, vec2d]

type manager struct {
	storage
	activeChunks *list.List
	chunkElement map[level.ChunkPos]*list.Element
	players      bvh.Tree[float64, sphere2d, *server.Player]
	clients      *clientinfo.ClientInformation
}

type chunkHandler struct {
	level.ChunkPos
	*level.Chunk
	// records all players loaded this chunk
	players map[*server.Player]struct{}
}

func (m *manager) refresh() error {
	all := func(b sphere2d) bool { return true }
	m.players.Find(all, func(n *bvh.Node[float64, sphere2d, *server.Player]) bool {
		p := n.Value
		v := m.clients.Players[p.UUID].ViewDistance
		for i := 1 - v; i < v; i++ {
			for j := 1 - v; j < v; j++ {
				pos := level.ChunkPos{
					X: int(n.Box.Center[0]) >> 4,
					Z: int(n.Box.Center[1]) >> 4,
				}
				point := vec2d{float64(pos.X + i), float64(pos.Z + j)}
				if _, exist := m.chunkElement[pos]; !exist && n.Box.WithIn(point) {
					c, err := m.GetChunk(pos)
					if err != nil {
						return false
					}
					m.chunkElement[pos] = m.activeChunks.PushBack(c)
				}
			}
		}
		return true
	})
	for e := m.activeChunks.Front(); e != nil; {
		ch := e.Value.(*chunkHandler)
		point := vec2d{float64(ch.X), float64(ch.Z)}
		filter := bvh.TouchPoint[vec2d, sphere2d](point)
		newPlayers := make(map[*server.Player]struct{})
		m.players.Find(filter, func(n *bvh.Node[float64, sphere2d, *server.Player]) bool {
			p := n.Value
			if _, ok := ch.players[p]; ok {
				delete(ch.players, p)
			} else {
				playerLoadChunk(p, ch)
			}
			newPlayers[p] = struct{}{}
			return true
		})
		for p := range ch.players {
			playerUnloadChunk(p, ch)
		}
		if len(newPlayers) > 0 {
			ch.players = newPlayers
			e = e.Next()
		} else {
			// no player around this chunk, unload it
			if err := m.PutChunk(ch.ChunkPos, ch.Chunk); err != nil {
				return err
			}
			next := e.Next()
			m.activeChunks.Remove(e)
			delete(m.chunkElement, ch.ChunkPos)
			e = next
		}
	}
	return nil
}

func playerLoadChunk(p *server.Player, c *chunkHandler) {
	p.WritePacket(server.Packet758(pk.Marshal(
		packetid.ClientboundLevelChunkWithLight,
		c.ChunkPos, c.Chunk,
	)))
}

func playerUnloadChunk(p *server.Player, c *chunkHandler) {
	p.WritePacket(server.Packet758(pk.Marshal(
		packetid.ClientboundForgetLevelChunk,
		c.ChunkPos,
	)))
}
