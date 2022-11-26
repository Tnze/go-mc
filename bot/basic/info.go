package basic

import (
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/nbt"
)

// WorldInfo content player info in server.
type WorldInfo struct {
	DimensionCodec      DimensionCodec
	DimensionType       string
	DimensionNames      []string    // Identifiers for all worlds on the server.
	DimensionName       string      // Name of the world being spawned into.
	HashedSeed          int64       // First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise
	MaxPlayers          int32       // Was once used by the client to draw the player list, but now is ignored.
	ViewDistance        int32       // Render distance (2-32).
	SimulationDistance  int32       // The distance that the client will process specific things, such as entities.
	ReducedDebugInfo    bool        // If true, a Notchian client shows reduced information on the debug screen. For servers in development, this should almost always be false.
	EnableRespawnScreen bool        // Set to false when the doImmediateRespawn gamerule is true.
	IsDebug             bool        // True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks.
	IsFlat              bool        // True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63.
	HasDeathLocation    bool        // If true, then the next two fields are present.
	DeathDimensionName  string      // The name of the dimension the player died in.
	DeathPosition       maths.Vec3d // The position the player died at.
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

type DimensionCodec struct {
	// What is Below? (wiki.vg)
	ChatType      Registry[nbt.RawMessage] `nbt:"minecraft:chat_type"`
	DimensionType Registry[Dimension]      `nbt:"minecraft:dimension_type"`
	WorldGenBiome Registry[nbt.RawMessage] `nbt:"minecraft:worldgen/biome"`
}

type Registry[E any] struct {
	Type  string `nbt:"type"`
	Value []struct {
		Name    string `nbt:"name"`
		ID      int32  `nbt:"id"`
		Element E      `nbt:"element"`
	} `nbt:"value"`
}

func (r *Registry[E]) Find(name string) *E {
	for i := range r.Value {
		if r.Value[i].Name == name {
			return &r.Value[i].Element
		}
	}
	return nil
}

type PlayerInfo struct {
	EID          int32 // The player's Entity ID (EID).
	Hardcore     bool  // Is hardcore
	Gamemode     byte  // Gamemode. 0: Survival, 1: Creative, 2: Adventure, 3: Spectator.
	PrevGamemode int8  // Previous Gamemode
}

// ServInfo contains information about the server implementation.
type ServInfo struct {
	Brand string
}
