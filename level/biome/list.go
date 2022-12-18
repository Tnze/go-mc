package biome

import "math/bits"

type Type int

var (
	BitsPerBiome int
	BiomesIDs    map[string]Type
	BiomesNames  = []string{
		"the_void",
		"plains",
		"sunflower_plains",
		"snowy_plains",
		"ice_spikes",
		"desert",
		"swamp",
		"mangrove_swamp",
		"forest",
		"flower_forest",
		"birch_forest",
		"dark_forest",
		"old_growth_birch_forest",
		"old_growth_pine_taiga",
		"old_growth_spruce_taiga",
		"taiga",
		"snowy_taiga",
		"savanna",
		"savanna_plateau",
		"windswept_hills",
		"windswept_gravelly_hills",
		"windswept_forest",
		"windswept_savanna",
		"jungle",
		"sparse_jungle",
		"bamboo_jungle",
		"badlands",
		"eroded_badlands",
		"wooded_badlands",
		"meadow",
		"grove",
		"snowy_slopes",
		"frozen_peaks",
		"jagged_peaks",
		"stony_peaks",
		"river",
		"frozen_river",
		"beach",
		"snowy_beach",
		"stony_shore",
		"warm_ocean",
		"lukewarm_ocean",
		"deep_lukewarm_ocean",
		"ocean",
		"deep_ocean",
		"cold_ocean",
		"deep_cold_ocean",
		"frozen_ocean",
		"deep_frozen_ocean",
		"mushroom_fields",
		"dripstone_caves",
		"lush_caves",
		"deep_dark",
		"nether_wastes",
		"warped_forest",
		"crimson_forest",
		"soul_sand_valley",
		"basalt_deltas",
		"the_end",
		"end_highlands",
		"end_midlands",
		"small_end_islands",
		"end_barrens",
	}
)

func init() {
	BitsPerBiome = bits.Len(uint(len(BiomesNames)))
	BiomesIDs = make(map[string]Type, len(BiomesNames))
	for i, v := range BiomesNames {
		BiomesIDs[v] = Type(i)
	}
}
