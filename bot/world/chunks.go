package world

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	. "github.com/Tnze/go-mc/level"
	block2 "github.com/Tnze/go-mc/level/block"
	"math"
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
	if int32(pos.Y) < 0 || int32(pos.Y) > 256 {
		return -1, fmt.Errorf("out of range")
	}
	chunkPos := ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		return chunk.GetBlock(pos)
	} else {
		return -1, fmt.Errorf("chunk not loaded")
	}
}

func (w *World) RayTrace(start, end maths.Vec3d) (core.RayTraceResult, error) {
	if start == maths.NullVec3d {
		return core.RayTraceResult{}, fmt.Errorf("start is null")
	}
	if end == maths.NullVec3d {
		return core.RayTraceResult{}, fmt.Errorf("end is null")
	}

	for _, pos := range maths.RayTraceBlocks(start, end) {
		block, err := w.GetBlock(pos)
		if err != nil {
			return core.RayTraceResult{}, err
		}
		if block2.IsAir(block) {
			continue
		}
		return core.RayTraceResult{
			Position: pos,
			Side:     core.GetClosestFacing(start, pos),
			Block:    block2.StateList[block],
		}, nil
	}

	return core.RayTraceResult{}, fmt.Errorf("no block found")
}

func (w *World) GetBlockDensity(pos maths.Vec3d, bb core.AxisAlignedBB) float32 {
	d0 := 1.0 / ((bb.MaxX-bb.MinX)*2.0 + 1.0)
	d1 := 1.0 / ((bb.MaxY-bb.MinY)*2.0 + 1.0)
	d2 := 1.0 / ((bb.MaxZ-bb.MinZ)*2.0 + 1.0)
	d3 := (1.0 - math.Floor(1.0/d0)*d0) / 2.0
	d4 := (1.0 - math.Floor(1.0/d2)*d2) / 2.0

	if d0 >= 0.0 && d1 >= 0.0 && d2 >= 0.0 {
		j2 := float32(0)
		k2 := float32(0)

		for f := 0.0; f <= 1.0; f += d0 {
			for f1 := 0.0; f1 <= 1.0; f1 += d1 {
				for f2 := 0.0; f2 <= 1.0; f2 += d2 {
					d5 := bb.MinX + (bb.MaxX-bb.MinX)*f
					d6 := bb.MinY + (bb.MaxY-bb.MinY)*f1
					d7 := bb.MinZ + (bb.MaxZ-bb.MinZ)*f2

					if result, err := w.RayTrace(maths.Vec3d{X: float32(d5 + d3), Y: float32(d6), Z: float32(d7 + d4)}, pos); result.Block == nil && err == nil {
						j2++
					}
					k2++
				}
			}
		}

		return j2 / k2
	}
	return 0
}

func (w *World) IsAABBInMaterial(bb core.AxisAlignedBB) bool {
	i := int32(math.Floor(float64(bb.MinX)))
	j := int32(math.Floor(float64(bb.MaxX)))
	k := int32(math.Floor(float64(bb.MinY)))
	l := int32(math.Floor(float64(bb.MaxY)))
	i1 := int32(math.Floor(float64(bb.MinZ)))
	j1 := int32(math.Floor(float64(bb.MaxZ)))

	for x := i; x <= j; x++ {
		for y := k; y <= l; y++ {
			for z := i1; z <= j1; z++ {
				if block, err := w.GetBlock(maths.Vec3d{float32(x), float32(y), float32(z)}); err == nil {
					if block2.StateList[block].ID() == "minecraft:water" { //TODO: fix this
						return false
					}
				}
			}
		}
	}
	return true
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
