package level

type ChunkStatus string

const (
	StatusEmpty               ChunkStatus = "empty"
	StatusStructureStarts     ChunkStatus = "structure_starts"
	StatusStructureReferences ChunkStatus = "structure_references"
	StatusBiomes              ChunkStatus = "biomes"
	StatusNoise               ChunkStatus = "noise"
	StatusSurface             ChunkStatus = "surface"
	StatusCarvers             ChunkStatus = "carvers"
	StatusLiquidCarvers       ChunkStatus = "liquid_carvers"
	StatusFeatures            ChunkStatus = "features"
	StatusLight               ChunkStatus = "light"
	StatusSpawn               ChunkStatus = "spawn"
	StatusHeightmaps          ChunkStatus = "heightmaps"
	StatusFull                ChunkStatus = "full"
)
