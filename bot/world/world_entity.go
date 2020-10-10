package world

import (
	"fmt"

	"github.com/Tnze/go-mc/bot/world/entity"
	e "github.com/Tnze/go-mc/data/entity"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/net/ptypes"
	"github.com/google/uuid"
)

// PlayerEntities returns a list of players on the server within viewing range.
func (w *World) PlayerEntities() []entity.Entity {
	w.entityLock.RLock()
	defer w.entityLock.RUnlock()
	out := make([]entity.Entity, 0, 12)
	for _, ent := range w.Entities {
		if ent.Base.ID == e.Player.ID {
			out = append(out, *ent)
		}
	}
	return out
}

// OnSpawnEntity should be called when a SpawnEntity packet
// is recieved.
func (w *World) OnSpawnEntity(pkt ptypes.SpawnEntity) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	base, ok := e.ByID[e.ID(pkt.Type)]
	if !ok {
		return fmt.Errorf("unknown entity ID %v", pkt.Type)
	}

	w.Entities[int32(pkt.ID)] = &entity.Entity{
		ID:    int32(pkt.ID),
		Base:  base,
		Data:  int32(pkt.Data),
		UUID:  uuid.UUID(pkt.UUID),
		X:     float64(pkt.X),
		Y:     float64(pkt.Y),
		Z:     float64(pkt.Z),
		Pitch: int8(pkt.Pitch),
		Yaw:   int8(pkt.Yaw),
		VelX:  int16(pkt.VelX),
		VelY:  int16(pkt.VelY),
		VelZ:  int16(pkt.VelZ),
	}

	return nil
}

// OnSpawnLivingEntity should be called when a SpawnLivingEntity packet
// is recieved.
func (w *World) OnSpawnLivingEntity(pkt ptypes.SpawnLivingEntity) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	base, ok := e.ByID[e.ID(pkt.Type)]
	if !ok {
		return fmt.Errorf("unknown entity ID %v", pkt.Type)
	}

	// fmt.Printf("SpawnLivingEntity %q\n", base.Name)
	w.Entities[int32(pkt.ID)] = &entity.Entity{
		ID:        int32(pkt.ID),
		Base:      base,
		UUID:      uuid.UUID(pkt.UUID),
		X:         float64(pkt.X),
		Y:         float64(pkt.Y),
		Z:         float64(pkt.Z),
		Pitch:     int8(pkt.Pitch),
		Yaw:       int8(pkt.Yaw),
		VelX:      int16(pkt.VelX),
		VelY:      int16(pkt.VelY),
		VelZ:      int16(pkt.VelZ),
		HeadPitch: int8(pkt.HeadPitch),
	}
	return nil
}

// OnSpawnPlayer should be called when a SpawnPlayer packet
// is recieved.
func (w *World) OnSpawnPlayer(pkt ptypes.SpawnPlayer) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	// fmt.Printf("SpawnPlayer %v\n", pkt)
	w.Entities[int32(pkt.ID)] = &entity.Entity{
		ID:    int32(pkt.ID),
		Base:  &e.Player,
		UUID:  uuid.UUID(pkt.UUID),
		X:     float64(pkt.X),
		Y:     float64(pkt.Y),
		Z:     float64(pkt.Z),
		Pitch: int8(pkt.Pitch),
		Yaw:   int8(pkt.Yaw),
	}
	return nil
}

// OnEntityPosUpdate should be called when an EntityPosition packet
// is recieved.
func (w *World) OnEntityPosUpdate(pkt ptypes.EntityPosition) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	ent, ok := w.Entities[int32(pkt.ID)]
	if !ok {
		return fmt.Errorf("cannot handle position update for unknown entity %d", pkt.ID)
	}

	ent.X += float64(pkt.X) / (128 * 32)
	ent.Y += float64(pkt.Y) / (128 * 32)
	ent.Z += float64(pkt.Z) / (128 * 32)
	ent.OnGround = bool(pkt.OnGround)
	return nil
}

// OnEntityPosLookUpdate should be called when an EntityPositionLook packet
// is recieved.
func (w *World) OnEntityPosLookUpdate(pkt ptypes.EntityPositionLook) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	ent, ok := w.Entities[int32(pkt.ID)]
	if !ok {
		return fmt.Errorf("cannot handle position look update for unknown entity %d", pkt.ID)
	}

	ent.X += float64(pkt.X) / (128 * 32)
	ent.Y += float64(pkt.Y) / (128 * 32)
	ent.Z += float64(pkt.Z) / (128 * 32)
	ent.OnGround = bool(pkt.OnGround)
	ent.Pitch, ent.Yaw = int8(pkt.Pitch), int8(pkt.Yaw)
	return nil
}

// OnEntityLookUpdate should be called when an EntityRotation packet
// is recieved.
func (w *World) OnEntityLookUpdate(pkt ptypes.EntityRotation) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	ent, ok := w.Entities[int32(pkt.ID)]
	if !ok {
		return fmt.Errorf("cannot handle look update for unknown entity %d", pkt.ID)
	}

	ent.Pitch, ent.Yaw = int8(pkt.Pitch), int8(pkt.Yaw)
	ent.OnGround = bool(pkt.OnGround)
	return nil
}

// OnEntityDestroy should be called when a DestroyEntities packet
// is recieved.
func (w *World) OnEntityDestroy(eIDs []pk.VarInt) error {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	for _, eID := range eIDs {
		delete(w.Entities, int32(eID))
	}
	return nil
}
