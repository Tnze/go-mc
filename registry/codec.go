package registry

import (
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/nbt"
)

type NetworkCodec struct {
	ChatType      Registry[ChatType]       `nbt:"minecraft:chat_type"`
	DamageType    Registry[DamageType]     `nbt:"minecraft:damage_type"`
	DimensionType Registry[Dimension]      `nbt:"minecraft:dimension_type"`
	TrimMaterial  Registry[nbt.RawMessage] `nbt:"minecraft:trim_material"`
	TrimPattern   Registry[nbt.RawMessage] `nbt:"minecraft:trim_pattern"`
	WorldGenBiome Registry[nbt.RawMessage] `nbt:"minecraft:worldgen/biome"`
}

type ChatType struct {
	Chat      chat.Decoration `nbt:"chat"`
	Narration chat.Decoration `nbt:"narration"`
}

type DamageType struct {
	MessageID        string  `nbt:"message_id"`
	Scaling          string  `nbt:"scaling"`
	Exhaustion       float32 `nbt:"exhaustion"`
	Effects          string  `nbt:"effects,omitempty"`
	DeathMessageType string  `nbt:"death_message_type,omitempty"`
}

type Dimension struct {
	FixedTime          int64   `nbt:"fixed_time,omitempty"`
	HasSkylight        bool    `nbt:"has_skylight"`
	HasCeiling         bool    `nbt:"has_ceiling"`
	Ultrawarm          bool    `nbt:"ultrawarm"`
	Natural            bool    `nbt:"natural"`
	CoordinateScale    float64 `nbt:"coordinate_scale"`
	BedWorks           bool    `nbt:"bed_works"`
	RespawnAnchorWorks byte    `nbt:"respawn_anchor_works"`
	MinY               int32   `nbt:"min_y"`
	Height             int32   `nbt:"height"`
	LogicalHeight      int32   `nbt:"logical_height"`
	InfiniteBurn       string  `nbt:"infiniburn"`
	Effects            string  `nbt:"effects"`
	AmbientLight       float64 `nbt:"ambient_light"`

	PiglinSafe                  byte           `nbt:"piglin_safe"`
	HasRaids                    byte           `nbt:"has_raids"`
	MonsterSpawnLightLevel      nbt.RawMessage `nbt:"monster_spawn_light_level"` // Tag_Int or {type:"minecraft:uniform", value:{min_inclusive: Tag_Int, max_inclusive: Tag_Int}}
	MonsterSpawnBlockLightLimit int32          `nbt:"monster_spawn_block_light_limit"`
}
