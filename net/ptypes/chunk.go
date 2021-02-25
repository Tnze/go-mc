// Package ptypes implements encoding and decoding for high-level packets.
package ptypes

import (
	"io"
	"math"

	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ChunkData is a client-bound packet which describes a chunk.
type ChunkData struct {
	X, Z           pk.Int
	FullChunk      pk.Boolean
	PrimaryBitMask pk.VarInt
	Heightmaps     struct{}
	Biomes         biomesData
	Data           pk.ByteArray
	BlockEntities  blockEntities
}

func (p *ChunkData) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		&p.X,
		&p.Z,
		&p.FullChunk,
		&p.PrimaryBitMask,
		&pk.NBT{V: &p.Heightmaps},
		pk.Opt{Has: &p.FullChunk, Field: &p.Biomes},
		&p.Data,
		&p.BlockEntities,
	}.ReadFrom(r)
}

type biomesData struct {
	data []pk.VarInt
}

func (b *biomesData) ReadFrom(r io.Reader) (int64, error) {
	var n pk.VarInt // Number of Biomes Data
	return pk.Tuple{
		&n, pk.Ary{Len: &n, Ary: []pk.VarInt{}},
	}.ReadFrom(r)
}

type blockEntities []entity.BlockEntity

// Decode implement net.packet.FieldDecoder
func (b *blockEntities) ReadFrom(r io.Reader) (n int64, err error) {
	var sz pk.VarInt // Number of BlockEntities
	if nn, err := sz.ReadFrom(r); err != nil {
		return nn, err
	} else {
		n += nn
	}
	*b = make(blockEntities, sz)
	lr := &io.LimitedReader{R: r, N: math.MaxInt64}
	d := nbt.NewDecoder(lr)
	for i := 0; i < int(sz); i++ {
		if err := d.Decode(&(*b)[i]); err != nil {
			return math.MaxInt64 - lr.N, err
		}
	}
	return math.MaxInt64 - lr.N, nil
}

// TileEntityData describes a change to a tile entity.
type TileEntityData struct {
	Pos    pk.Position
	Action pk.UnsignedByte
	Data   entity.BlockEntity
}

func (p *TileEntityData) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		&p.Pos,
		&p.Action,
		&pk.NBT{V: &p.Data},
	}.ReadFrom(r)
}
