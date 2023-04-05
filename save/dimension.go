package save

// https://minecraft.fandom.com/wiki/Custom_dimension
type DimensionType struct {
	Ultrawarm bool `nbt:"ultrawarm"`

	// When false, compasses spin randomly, and using a bed to set the respawn point or sleep, is disabled. When true, nether portals can spawn zombified piglins.
	Natural bool `nbt:"natural"`

	// The multiplier applied to coordinates when leaving the dimension. Value between 0.00001 and 30000000.0 (both inclusive)
	CoordinatesScale float64 `nbt:"coordinate_scale"`

	HasSkylight  bool    `nbt:"has_skylight"`
	HasCeiling   bool    `nbt:"has_ceiling"`
	AmbientLight float32 `nbt:"ambient_light"`
	FixedTime    int64   `nbt:"fixed_time,omitempty"`

	// idk what this really is, looks like an integer
	MonsterSpawnLightLevel      int32 `nbt:"monster_spawn_light_level"`
	MonsterSpawnBlockLightLimit int32 `nbt:"monster_spawn_block_light_limit"`

	PiglinSafe         bool   `nbt:"piglin_safe"`
	BedWorks           bool   `nbt:"bed_works"`
	RespawnAnchorWorks bool   `nbt:"respawn_anchor_works"`
	HasRaids           bool   `nbt:"has_raids"`
	LogicalHeight      int32  `nbt:"logical_height"`
	MinY               int32  `nbt:"min_y"`
	Height             int32  `nbt:"height"`
	Infiniburn         string `nbt:"infiniburn"`
	Effects            string `nbt:"effects"`
}

type WorldGenSettings struct {
	BonusChest       bool                          `nbt:"bonus_chest"`
	GenerateFeatures bool                          `nbt:"generate_features"`
	Seed             int64                         `nbt:"seed"`
	Dimensions       map[string]DimensionGenerator `nbt:"dimensions"`
}

type DimensionGenerator struct {
	Type      string         `nbt:"type"`
	Generator map[string]any `nbt:"generator"`
}

var (
	// as of 1.19.2
	DefaultDimensionsGenerators = map[string]DimensionGenerator{
		"minecraft:overworld": {
			Type: "minecraft:overworld",
			Generator: map[string]any{
				"type":     "minecraft:noise",
				"settings": "minecraft:overworld",
				"biome_source": map[string]string{
					"preset": "minecraft:overworld",
					"type":   "minecraft:multi_noise",
				},
			},
		},
		"minecraft:the_end": {
			Type: "minecraft:the_end",
			Generator: map[string]any{
				"type":     "minecraft:noise",
				"settings": "minecraft:end",
				"biome_source": map[string]string{
					"type": "minecraft:the_end",
				},
			},
		},
		"minecraft:the_nether": {
			Type: "minecraft:the_nether",
			Generator: map[string]any{
				"type":     "minecraft:noise",
				"settings": "minecraft:nether",
				"biome_source": map[string]string{
					"preset": "minecraft:nether",
					"type":   "minecraft:multi_noise",
				},
			},
		},
	}
	DefaultDimensionsTypes = map[string]DimensionType{
		"minecraft:overworld": {
			FixedTime:                   0,
			HasSkylight:                 true,
			HasCeiling:                  false,
			Ultrawarm:                   false,
			Natural:                     true,
			CoordinatesScale:            1.0,
			BedWorks:                    true,
			RespawnAnchorWorks:          false,
			MinY:                        -64,
			Height:                      384,
			LogicalHeight:               384,
			Infiniburn:                  "minecraft:infiniburn_overworld",
			Effects:                     "minecraft:overworld",
			AmbientLight:                0.0,
			PiglinSafe:                  false,
			HasRaids:                    true,
			MonsterSpawnLightLevel:      7,
			MonsterSpawnBlockLightLimit: 0,
		},
		"minecraft:the_nether": {
			FixedTime:                   18000,
			HasSkylight:                 false,
			HasCeiling:                  true,
			Ultrawarm:                   true,
			Natural:                     false,
			CoordinatesScale:            8.0,
			BedWorks:                    false,
			RespawnAnchorWorks:          true,
			MinY:                        0,
			Height:                      256,
			LogicalHeight:               128,
			Infiniburn:                  "minecraft:infiniburn_nether",
			Effects:                     "minecraft:the_nether",
			AmbientLight:                0.1,
			PiglinSafe:                  true,
			HasRaids:                    false,
			MonsterSpawnLightLevel:      11,
			MonsterSpawnBlockLightLimit: 15,
		},
		"minecraft:the_end": {
			FixedTime:                   6000,
			HasSkylight:                 false,
			HasCeiling:                  false,
			Ultrawarm:                   false,
			Natural:                     false,
			CoordinatesScale:            1.0,
			BedWorks:                    false,
			RespawnAnchorWorks:          false,
			MinY:                        0,
			Height:                      256,
			LogicalHeight:               256,
			Infiniburn:                  "minecraft:infiniburn_end",
			Effects:                     "minecraft:the_end",
			AmbientLight:                0.0,
			PiglinSafe:                  false,
			HasRaids:                    true,
			MonsterSpawnLightLevel:      7,
			MonsterSpawnBlockLightLimit: 0,
		},
		"minecraft:overworld_caves": {
			FixedTime:                   0,
			HasSkylight:                 true,
			HasCeiling:                  true,
			Ultrawarm:                   false,
			Natural:                     false,
			CoordinatesScale:            1.0,
			BedWorks:                    true,
			RespawnAnchorWorks:          false,
			MinY:                        -64,
			Height:                      384,
			LogicalHeight:               384,
			Infiniburn:                  "minecraft:infiniburn_overworld",
			Effects:                     "minecraft:overworld",
			AmbientLight:                0.0,
			PiglinSafe:                  false,
			HasRaids:                    true,
			MonsterSpawnLightLevel:      7,
			MonsterSpawnBlockLightLimit: 0,
		},
	}
)
