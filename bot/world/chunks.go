package world

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	. "github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/level/block"
	"math"
)

type World struct {
	Columns        map[ChunkPos]*Chunk
	entities       map[int32]*core.EntityInterface
	entitiesLiving map[int32]*core.EntityLivingInterface
	entitiesPlayer map[int32]*core.EntityPlayerInterface
}

func NewWorld() (w *World) {
	w = &World{
		Columns:        make(map[ChunkPos]*Chunk),
		entities:       make(map[int32]*core.EntityInterface),
		entitiesLiving: make(map[int32]*core.EntityLivingInterface),
		entitiesPlayer: make(map[int32]*core.EntityPlayerInterface),
	}
	return
}

func (w *World) AddEntity(e core.EntityInterface) error {
	if w.isValidEntity(&e) {
		w.entities[e.GetID()] = &e
		return nil
	}
	return fmt.Errorf("invalid entity")
}

func (w *World) RemoveEntity(e *core.EntityInterface) error {
	if w.isValidEntity(e) {
		delete(w.entities, (*e).GetID())
	}
	return fmt.Errorf("invalid entity")
}

func (w *World) GetEntities() map[int32]*core.EntityInterface {
	return w.entities
}

/*func (w *World) GetEntitiesByType(t interface{}) []*interface{} {
	var entities []*interface{}
	for _, e := range w.entities {
		if reflect.TypeOf(*e) == reflect.TypeOf(t) {
			entities = append(entities, e)
		}
	}
	return entities
}*/

func (w *World) GetEntityByID(id int32) (int32, interface{}, error) {
	for i, e := range w.entities {
		if w.isValidEntity(e) && id == (*e).GetID() {
			return i, *e, nil
		}
	}
	return -1, nil, fmt.Errorf("entity not found")
}

func (w *World) isValidEntity(e *core.EntityInterface) bool {
	/*if _, ok := (*e).(core.Entity); ok {
		return true
	} else if _, ok := (*e).(*core.EntityLiving); ok {
		return true
	} else if _, ok := (*e).(*core.EntityPlayer); ok {
		return true
	}
	return false*/
	return true
}

func (w *World) GetBlock(pos maths.Vec3d[float64]) (block.Block, basic.Error) {
	chunkPos := ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		return chunk.GetBlock(pos)
	} else {
		return block.StateList[block.ToStateID[block.Air{}]], basic.Error{Err: basic.InvalidChunk, Info: fmt.Errorf("chunk not found")}
	}
}

func (w *World) SetBlock(d maths.Vec3d[float64], i int) {
	chunkPos := ChunkPos{int32(d.X) >> 4, int32(d.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		chunk.SetBlock(d, i)
	}
}

func (w *World) GetNeighbors(block maths.Vec3d[float64]) []maths.Vec3d[float64] {
	return []maths.Vec3d[float64]{
		{X: block.X + 1, Y: block.Y, Z: block.Z},
		{X: block.X - 1, Y: block.Y, Z: block.Z},
		{X: block.X, Y: block.Y + 1, Z: block.Z},
		{X: block.X, Y: block.Y - 1, Z: block.Z},
		{X: block.X, Y: block.Y, Z: block.Z + 1},
		{X: block.X, Y: block.Y, Z: block.Z - 1},
	}
}

func (w *World) IsBlockLoaded(pos maths.Vec3d[float64]) bool {
	chunkPos := ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		return chunk.IsBlockLoaded(pos)
	}
	return false
}

func (w *World) IsChunkLoaded(pos ChunkPos) bool {
	_, ok := w.Columns[pos]
	return ok
}

func (w *World) RayTrace(start, end maths.Vec3d[float64]) (maths.RayTraceResult, basic.Error) {
	if start == maths.NullVec3d && end == maths.NullVec3d {
		return maths.RayTraceResult{}, basic.Error{Err: basic.NullValue, Info: fmt.Errorf("start and end cannot be null")}
	}

	for _, pos := range maths.RayTraceBlocks(start, end) {
		block, _ := w.GetBlock(pos)
		if block.IsAir() {
			continue
		} else {
			return maths.RayTraceResult{
				Position: pos,
			}, basic.Error{Err: basic.NoError, Info: nil}
		}
	}

	return maths.RayTraceResult{}, basic.Error{Err: basic.NoValue, Info: fmt.Errorf("no block found")}
}

func (w *World) GetBlockDensity(pos maths.Vec3d[float64], bb maths.AxisAlignedBB[float64]) float64 {
	d0 := 1.0 / ((bb.MaxX-bb.MinX)*2.0 + 1.0)
	d1 := 1.0 / ((bb.MaxY-bb.MinY)*2.0 + 1.0)
	d2 := 1.0 / ((bb.MaxZ-bb.MinZ)*2.0 + 1.0)
	d3 := (1.0 - math.Floor(1.0/d0)) * d0 / 2.0
	d4 := (1.0 - math.Floor(1.0/d2)) * d2 / 2.0

	if d0 >= 0.0 && d1 >= 0.0 && d2 >= 0.0 {
		j2 := 0.0
		k2 := 0.0

		for f := 0.0; f <= 1.0; f += d0 {
			for f1 := 0.0; f1 <= 1.0; f1 += d1 {
				for f2 := 0.0; f2 <= 1.0; f2 += d2 {
					d5 := bb.MinX + (bb.MaxX-bb.MinX)*f
					d6 := bb.MinY + (bb.MaxY-bb.MinY)*f1
					d7 := bb.MinZ + (bb.MaxZ-bb.MinZ)*f2

					if _, err := w.RayTrace(maths.Vec3d[float64]{X: d5 + d3, Y: d6, Z: d7 + d4}, pos); err.Is(basic.NoValue) {
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

func (w *World) IsAABBInMaterial(bb maths.AxisAlignedBB[float64]) bool {
	i := int32(math.Floor(bb.MinX))
	j := int32(math.Floor(bb.MaxX))
	k := int32(math.Floor(bb.MinY))
	l := int32(math.Floor(bb.MaxY))
	i1 := int32(math.Floor(bb.MinZ))
	j1 := int32(math.Floor(bb.MaxZ))

	for x := i; x <= j; x++ {
		for y := k; y <= l; y++ {
			for z := i1; z <= j1; z++ {
				if getBlock, err := w.GetBlock(maths.Vec3d[float64]{X: float64(x), Y: float64(y), Z: float64(z)}); !err.Is(basic.InvalidChunk) && !getBlock.IsAir() {
					if getBlock.IsLiquid() {
						return false
					}
				}
			}
		}
	}
	return true
}

func (w *World) GetCollisionBoxes(e core.Entity, aabb maths.AxisAlignedBB[float64]) []maths.AxisAlignedBB[float64] {
	var boxes []maths.AxisAlignedBB[float64]
	/*for _, entity := range w.GetEntitiesInAABB(aabb) {
		if entity != e {
			boxes = append(boxes, entity.GetBoundingBox())
		}
	}*/
	return boxes
}

func (w *World) GetEntitiesInAABB(bb maths.AxisAlignedBB[float64]) []interface{} {
	var entities []interface{}
	/*for _, e := range w.entities {
		if bb.IntersectsWith(e) {
			entities = append(entities, e)
		}
	}*/
	return entities
}

func (w *World) GetEntitiesInAABBExcludingEntity(e core.Entity, bb maths.AxisAlignedBB[float64]) []interface{} {
	var entities []interface{}
	/*for _, entity := range w.entities {
		if entity != e && bb.IntersectsWith(entity) {
			entities = append(entities, entity)
		}
	}*/
	return entities
}
