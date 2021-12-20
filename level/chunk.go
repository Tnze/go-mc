package level

import (
	"bytes"
	"io"
	"math/bits"
	"strings"
	"sync"
	"unsafe"

	"github.com/Tnze/go-mc/data/block"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/save"
)

type ChunkPos struct{ X, Z int }
type Chunk struct {
	sync.Mutex
	Sections   []Section
	HeightMaps HeightMaps
}
type HeightMaps struct {
	MotionBlocking *BitStorage
	WorldSurface   *BitStorage
}

func EmptyChunk(secs int) *Chunk {
	sections := make([]Section, secs)
	for i := range sections {
		sections[i] = Section{
			blockCount: 0,
			States:     NewStatesPaletteContainer(16*16*16, 0),
			Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
		}
	}
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking: NewBitStorage(bits.Len(uint(secs)*16), 16*16, nil),
		},
	}
}

var biomesIDs = map[string]int{"ocean": 0,
	"deep_ocean":               24,
	"frozen_ocean":             10,
	"deep_frozen_ocean":        50,
	"cold_ocean":               46,
	"deep_cold_ocean":          49,
	"lukewarm_ocean":           45,
	"deep_lukewarm_ocean":      48,
	"warm_ocean":               44,
	"river":                    7,
	"frozen_river":             11,
	"beach":                    16,
	"stony_shore":              25,
	"snowy_beach":              26,
	"forest":                   4,
	"flower_forest":            132,
	"birch_forest":             27,
	"old_growth_birch_forest":  155,
	"dark_forest":              29,
	"jungle":                   21,
	"sparse_jungle":            23,
	"bamboo_jungle":            168,
	"taiga":                    5,
	"snowy_taiga":              30,
	"old_growth_pine_taiga":    32,
	"old_growth_spruce_taiga":  160,
	"mushroom_fields":          14,
	"swamp":                    6,
	"savanna":                  35,
	"savanna_plateau":          36,
	"windswept_savanna":        163,
	"plains":                   1,
	"sunflower_plains":         129,
	"desert":                   2,
	"snowy_plains":             12,
	"ice_spikes":               140,
	"windswept_hills":          3,
	"windswept_forest":         34,
	"windswept_gravelly_hills": 131,
	"badlands":                 37,
	"wooded_badlands":          38,
	"eroded_badlands":          165,
	"dripstone_caves":          174,
	"lush_caves":               175,
	"nether_wastes":            8,
	"crimson_forest":           171,
	"warped_forest":            172,
	"soul_sand_valley":         170,
	"basalt_deltas":            173,
	"the_end":                  9,
	"small_end_islands":        40,
	"end_midlands":             41,
	"end_highlands":            42,
	"end_barrens":              43,
	"the_void":                 127,
	"meadow":                   177,
	"grove":                    178,
	"snowy_slopes":             179,
	"frozen_peaks":             180,
	"jagged_peaks":             181,
	"stony_peaks":              182,
}

func ChunkFromSave(c *save.Chunk, secs int) *Chunk {
	sections := make([]Section, secs)
	for _, v := range c.Sections {
		var blockCount int16
		stateData := *(*[]uint64)((unsafe.Pointer)(&v.BlockStates.Data))
		statePalette := v.BlockStates.Palette
		stateRawPalette := make([]int, len(statePalette))
		for i, v := range statePalette {
			// TODO: Consider the properties of block, not only index the block name
			stateRawPalette[i] = int(stateIDs[strings.TrimPrefix(v.Name, "minecraft:")])
			if v.Name != "minecraft:air" {
				blockCount++
			}
		}

		biomesData := *(*[]uint64)((unsafe.Pointer)(&v.BlockStates.Data))
		biomesPalette := v.Biomes.Palette
		biomesRawPalette := make([]int, len(biomesPalette))
		for i, v := range biomesPalette {
			biomesRawPalette[i] = biomesIDs[strings.TrimPrefix(v, "minecraft:")]
		}

		i := int32(int8(v.Y)) - c.YPos
		sections[i].blockCount = blockCount
		sections[i].States = NewStatesPaletteContainerWithData(16*16*16, stateData, stateRawPalette)
		sections[i].Biomes = NewBiomesPaletteContainerWithData(16*16*16*2, biomesData, biomesRawPalette)
	}
	for i := range sections {
		if sections[i].States == nil {
			sections[i] = Section{
				blockCount: 0,
				States:     NewStatesPaletteContainer(16*16*16, 0),
				Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
			}
		}
	}

	motionBlocking := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.MotionBlocking))
	worldSurface := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.WorldSurface))
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking: NewBitStorage(bits.Len(uint(secs)), 16*16, motionBlocking),
			WorldSurface:   NewBitStorage(bits.Len(uint(secs)), 16*16, worldSurface),
		},
	}
}

// TODO: This map should be moved to data/block.
var stateIDs = make(map[string]uint32)

func init() {
	for i, v := range block.StateID {
		name := block.ByID[v].Name
		if _, ok := stateIDs[name]; !ok {
			stateIDs[name] = i
		}
	}
}

func (c *Chunk) WriteTo(w io.Writer) (int64, error) {
	data, err := c.Data()
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		// Heightmaps
		pk.NBT(struct {
			MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
			WorldSurface   []uint64 `nbt:"WORLD_SURFACE"`
		}{
			MotionBlocking: c.HeightMaps.MotionBlocking.Raw(),
			WorldSurface:   c.HeightMaps.MotionBlocking.Raw(),
		}),
		pk.ByteArray(data),
		pk.VarInt(0), // TODO: Block Entity
		&lightData{
			SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
			BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
			SkyLight:       []pk.ByteArray{},
			BlockLight:     []pk.ByteArray{},
		},
	}.WriteTo(w)
}

func (c *Chunk) Data() ([]byte, error) {
	var buff bytes.Buffer
	for _, section := range c.Sections {
		_, err := section.WriteTo(&buff)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

type Section struct {
	blockCount int16
	States     *PaletteContainer
	Biomes     *PaletteContainer
}

func (s *Section) GetBlock(i int) int {
	return s.States.Get(i)
}
func (s *Section) SetBlock(i int, v int) {
	// TODO: Handle cave air and void air
	if s.States.Get(i) != 0 {
		s.blockCount--
	}
	if v != 0 {
		s.blockCount++
	}
	s.States.Set(i, v)
}

func (s *Section) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Short(s.blockCount),
		s.States,
		s.Biomes,
	}.WriteTo(w)
}

func (s *Section) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		pk.Short(s.blockCount),
		s.States,
		s.Biomes,
	}.ReadFrom(r)
}

type lightData struct {
	SkyLightMask   pk.BitSet
	BlockLightMask pk.BitSet
	SkyLight       []pk.ByteArray
	BlockLight     []pk.ByteArray
}

func bitSetRev(set pk.BitSet) pk.BitSet {
	rev := make(pk.BitSet, len(set))
	for i := range rev {
		rev[i] = ^set[i]
	}
	return rev
}

func (l *lightData) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Boolean(true), // Trust Edges
		l.SkyLightMask,
		l.BlockLightMask,
		bitSetRev(l.SkyLightMask),
		bitSetRev(l.BlockLightMask),
		pk.Array(l.SkyLight),
		pk.Array(l.BlockLight),
	}.WriteTo(w)
}
