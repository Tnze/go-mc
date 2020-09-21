// Package ptypes implements encoding and decoding for high-level packets.
package ptypes

import (
	"bytes"
	"fmt"

	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ChunkData is a clientbound packet which describes a chunk.
type ChunkData struct {
	X, Z           pk.Int
	FullChunk      pk.Boolean
	PrimaryBitMask pk.VarInt
	Heightmaps     struct{}
	Biomes         biomesData
	Data           chunkData
	BlockEntities  blockEntities
}

func (p *ChunkData) Decode(pkt pk.Packet) error {
	r := bytes.NewReader(pkt.Data)
	if err := p.X.Decode(r); err != nil {
		return fmt.Errorf("X: %v", err)
	}
	if err := p.Z.Decode(r); err != nil {
		return fmt.Errorf("Z: %v", err)
	}
	if err := p.FullChunk.Decode(r); err != nil {
		return fmt.Errorf("full chunk: %v", err)
	}
	if err := p.PrimaryBitMask.Decode(r); err != nil {
		return fmt.Errorf("bit mask: %v", err)
	}
	if err := (pk.NBT{V: &p.Heightmaps}).Decode(r); err != nil {
		return fmt.Errorf("heightmaps: %v", err)
	}

	// Biome data is only present for full chunks.
	if p.FullChunk {
		if err := p.Biomes.Decode(r); err != nil {
			return fmt.Errorf("heightmaps: %v", err)
		}
	}

	if err := p.Data.Decode(r); err != nil {
		return fmt.Errorf("data: %v", err)
	}
	if err := p.BlockEntities.Decode(r); err != nil {
		return fmt.Errorf("block entities: %v", err)
	}
	return nil
}

type biomesData struct {
	data []pk.VarInt
}

func (b *biomesData) Decode(r pk.DecodeReader) error {
	var nobd pk.VarInt // Number of Biome Datums
	if err := nobd.Decode(r); err != nil {
		return err
	}
	b.data = make([]pk.VarInt, nobd)

	for i := 0; i < int(nobd); i++ {
		var d pk.VarInt
		if err := d.Decode(r); err != nil {
			return err
		}
		b.data[i] = d
	}

	return nil
}

type chunkData []byte
type blockEntities []entity.BlockEntity

// Decode implement net.packet.FieldDecoder
func (c *chunkData) Decode(r pk.DecodeReader) error {
	var sz pk.VarInt
	if err := sz.Decode(r); err != nil {
		return err
	}
	*c = make([]byte, sz)
	if _, err := r.Read(*c); err != nil {
		return err
	}
	return nil
}

// Decode implement net.packet.FieldDecoder
func (b *blockEntities) Decode(r pk.DecodeReader) error {
	var sz pk.VarInt // Number of BlockEntities
	if err := sz.Decode(r); err != nil {
		return err
	}
	*b = make(blockEntities, sz)
	decoder := nbt.NewDecoder(r)
	for i := 0; i < int(sz); i++ {
		if err := decoder.Decode(&(*b)[i]); err != nil {
			return err
		}
	}
	return nil
}

// TileEntityData describes a change to a tile entity.
type TileEntityData struct {
	Pos    pk.Position
	Action pk.UnsignedByte
	Data   entity.BlockEntity
}

func (p *TileEntityData) Decode(pkt pk.Packet) error {
	r := bytes.NewReader(pkt.Data)
	if err := p.Pos.Decode(r); err != nil {
		return fmt.Errorf("position: %v", err)
	}
	if err := p.Action.Decode(r); err != nil {
		return fmt.Errorf("action: %v", err)
	}
	return nbt.NewDecoder(r).Decode(&p.Data)
}
