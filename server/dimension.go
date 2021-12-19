package server

import (
	"bytes"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/save"
	"io"
	"sync"
)

type Dimension interface {
	Info() DimInfo
	PlayerJoin(p *Player)
	PlayerQuit(p *Player)
}

type DimInfo struct {
	Name       string
	HashedSeed uint64
}

type SimpleDim struct {
	Columns map[struct{ X, Z int }]*struct {
		sync.Mutex
	}
}

type chunkData struct {
	HeightMaps save.BitStorage
	BlockState save.BitStorage
	Biomes     save.BitStorage
}

func (c *chunkData) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		// Heightmaps
		pk.NBT(struct {
			MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
		}{c.HeightMaps.Raw()}),
		pk.ByteArray(c.Data()), // TODO: Chunk Data
		pk.VarInt(0),           // TODO: Block Entity
	}.WriteTo(w)
}

func (c *chunkData) Data() []byte {
	var buff bytes.Buffer
	_, _ = pk.Short(0).WriteTo(&buff)
	return buff.Bytes()
}

type lightData struct {
	SkyLightMask   pk.BitSet
	BlockLightMask pk.BitSet
	SkyLight       []pk.ByteArray
	BlockLight     []pk.ByteArray
}

func bitSetRev(set pk.BitSet) pk.BitSet {
	rev := make(pk.BitSet, len(set))
	for i := range rev {
		rev[i] = ^set[i]
	}
	return rev
}

func (l *lightData) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Boolean(true), // Trust Edges
		l.SkyLightMask,
		l.BlockLightMask,
		bitSetRev(l.SkyLightMask),
		bitSetRev(l.BlockLightMask),
		pk.Array(l.SkyLight),
		pk.Array(l.BlockLight),
	}.WriteTo(w)
}

func (s *SimpleDim) PlayerJoin(p *Player) {
	for pos, column := range s.Columns {
		column.Lock()
		packet := pk.Marshal(
			packetid.ClientboundLevelChunkWithLight,
			pk.Int(pos.X), pk.Int(pos.Z),
			&chunkData{},
			&lightData{
				SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
				BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
				SkyLight:       []pk.ByteArray{},
				BlockLight:     []pk.ByteArray{},
			},
		)
		column.Unlock()

		err := p.WritePacket(Packet757(packet))
		if err != nil {
			return
		}
	}
}

func (s *SimpleDim) PlayerQuit(p *Player) {

}
