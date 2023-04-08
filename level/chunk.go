package level

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/maths"
	"io"
	"log"
	"math/bits"
	"strings"

	"github.com/Tnze/go-mc/level/block"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/save"
)

type ChunkPos [2]int32

func (c ChunkPos) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.Int(c[0]).WriteTo(w)
	if err != nil {
		return
	}
	n1, err := pk.Int(c[1]).WriteTo(w)
	return n + n1, err
}

func (c *ChunkPos) ReadFrom(r io.Reader) (n int64, err error) {
	var x, z pk.Int
	if n, err = x.ReadFrom(r); err != nil {
		return n, err
	}
	var n1 int64
	if n1, err = z.ReadFrom(r); err != nil {
		return n + n1, err
	}
	*c = ChunkPos{int32(x), int32(z)}
	return n + n1, nil
}

type Chunk struct {
	Sections    []Section
	HeightMaps  HeightMaps
	BlockEntity []BlockEntity
	Status      ChunkStatus
}

func (c *Chunk) GetBlock(vec3d maths.Vec3d[float64]) (block.Block, basic.Error) {
	X, Y, Z := int(vec3d.X), int(vec3d.Y), int(vec3d.Z)
	Y += 64 // Offset so that Y=-64 is the index 0 of the array
	if Y < 0 || Y >= len(c.Sections)*16 {
		return block.StateList[block.ToStateID[block.Air{}]], basic.Error{Err: basic.NoError, Info: fmt.Errorf("y=%d out of bound", Y)} // Safe check
	}
	if t := c.Sections[Y>>4]; t.States != nil {
		return block.StateList[t.States.Get(Y&15<<8|Z&15<<4|X&15)], basic.Error{Err: basic.NoError, Info: nil}
	} else {
		return block.StateList[block.ToStateID[block.Air{}]], basic.Error{Err: basic.NoError, Info: fmt.Errorf("y=%d out of bound", Y)}
	}
}

func (c *Chunk) SetBlock(d maths.Vec3d[float64], i int) {
	X, Y, Z := int(d.X), int(d.Y), int(d.Z)
	Y += 64 // Offset so that Y=-64 is the index 0 of the array
	if Y < 0 || Y >= len(c.Sections)*16 {
		return // Safe check
	}
	if t := c.Sections[Y>>4]; t.States != nil {
		t.States.Set(Y&15<<8|Z&15<<4|X&15, BlocksState(i))
	}
}

var biomesIDs map[string]BiomesState
var BitsPerBiome int

var biomesNames = []string{
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

func init() {
	biomesIDs = make(map[string]BiomesState, len(biomesNames))
	for i, v := range biomesNames {
		biomesIDs[v] = BiomesState(i)
	}
	BitsPerBiome = bits.Len(uint(len(biomesNames)))
}

func EmptyChunk(secs int) *Chunk {
	sections := make([]Section, secs)
	for i := range sections {
		sections[i] = Section{
			BlockCount: 0,
			States:     NewStatesPaletteContainer(16*16*16, 0),
			Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
		}
	}
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking: NewBitStorage(bits.Len(uint(secs)*16), 16*16, nil),
		},
		Status: StatusEmpty,
	}
}

// ChunkFromSave convert save.Chunk to level.Chunk.
func ChunkFromSave(c *save.Chunk) (*Chunk, error) {
	secs := len(c.Sections)
	sections := make([]Section, secs)
	for _, v := range c.Sections {
		i := int32(v.Y) - c.YPos
		if i < 0 || i >= int32(secs) {
			return nil, fmt.Errorf("section Y value %d out of bounds", v.Y)
		}
		var err error
		BlockCount, States, err := readStatesPalette(v.BlockStates.Palette, v.BlockStates.Data)
		if err != nil {
			return nil, err
		}
		Biomes, err := readBiomesPalette(v.Biomes.Palette, v.Biomes.Data)
		if err != nil {
			return nil, err
		}
		sections[i] = Section{
			BlockCount: BlockCount,
			States:     States,
			Biomes:     Biomes,
			SkyLight:   v.SkyLight,
			BlockLight: v.BlockLight,
		}
	}

	motionBlocking := c.Heightmaps.MotionBlocking
	motionBlockingNoLeaves := c.Heightmaps.MotionBlockingNoLeaves
	oceanFloor := c.Heightmaps.OceanFloor
	worldSurface := c.Heightmaps.WorldSurface

	bitsForHeight := bits.Len( /* chunk height in blocks */ uint(secs) * 16)
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking:         NewBitStorage(bitsForHeight, 16*16, motionBlocking),
			MotionBlockingNoLeaves: NewBitStorage(bitsForHeight, 16*16, motionBlockingNoLeaves),
			OceanFloor:             NewBitStorage(bitsForHeight, 16*16, oceanFloor),
			WorldSurface:           NewBitStorage(bitsForHeight, 16*16, worldSurface),
		},
		Status: ChunkStatus(c.Status),
	}, nil
}

func readStatesPalette(palette []save.BlockState, data []uint64) (blockCount int16, paletteData *PaletteContainer[BlocksState], err error) {
	statePalette := make([]BlocksState, len(palette))
	for i, v := range palette {
		b, ok := block.FromID[v.Name]
		if !ok {
			return 0, nil, fmt.Errorf("unknown block id: %v", v.Name)
		}
		if v.Properties.Data != nil {
			if err := v.Properties.Unmarshal(&b); err != nil {
				return 0, nil, fmt.Errorf("unmarshal block properties fail: %v", err)
			}
		}
		s, ok := block.ToStateID[b]
		if !ok {
			return 0, nil, fmt.Errorf("unknown block: %v", b)
		}
		if !block.IsAir(s) {
			blockCount++
		}
		statePalette[i] = s
	}
	paletteData = NewStatesPaletteContainerWithData(16*16*16, data, statePalette)
	return
}

func readBiomesPalette(palette []string, data []uint64) (*PaletteContainer[BiomesState], error) {
	biomesRawPalette := make([]BiomesState, len(palette))
	var ok bool
	for i, v := range palette {
		biomesRawPalette[i], ok = biomesIDs[strings.TrimPrefix(v, "minecraft:")]
		if !ok {
			return nil, fmt.Errorf("unknown biomes: %s", v)
		}
	}
	return NewBiomesPaletteContainerWithData(4*4*4, data, biomesRawPalette), nil
}

// ChunkToSave convert level.Chunk to save.Chunk
func ChunkToSave(c *Chunk, dst *save.Chunk) (err error) {
	secs := len(c.Sections)
	sections := make([]save.Section, secs)
	for i, v := range c.Sections {
		s := &sections[i]
		states := &s.BlockStates
		biomes := &s.Biomes
		s.Y = int8(int32(i) + dst.YPos)
		states.Palette, states.Data, err = writeStatesPalette(v.States)
		if err != nil {
			return
		}
		biomes.Palette, biomes.Data = writeBiomesPalette(v.Biomes)
		s.SkyLight = v.SkyLight
		s.BlockLight = v.BlockLight
	}
	dst.Sections = sections
	dst.Heightmaps.MotionBlocking = c.HeightMaps.MotionBlocking.Raw()
	dst.Status = string(c.Status)
	return
}

func writeStatesPalette(paletteData *PaletteContainer[BlocksState]) (palette []save.BlockState, data []uint64, err error) {
	rawPalette := paletteData.palette.export()
	palette = make([]save.BlockState, len(rawPalette))
	var buffer bytes.Buffer
	for i, v := range rawPalette {
		b := block.StateList[v]
		palette[i].Name = b.ID()

		buffer.Reset()
		err = nbt.NewEncoder(&buffer).Encode(b, "")
		if err != nil {
			return
		}
		_, err = nbt.NewDecoder(&buffer).Decode(&palette[i].Properties)
		if err != nil {
			return
		}
	}
	data = append(data, paletteData.data.Raw()...)

	return
}

func writeBiomesPalette(paletteData *PaletteContainer[BiomesState]) (palette []string, data []uint64) {
	rawPalette := paletteData.palette.export()
	palette = make([]string, len(rawPalette))
	for i, v := range rawPalette {
		palette[i] = biomesNames[v]
	}
	data = append(data, paletteData.data.Raw()...)

	return
}

func (c *Chunk) WriteTo(w io.Writer) (int64, error) {
	data, err := c.Data()
	if err != nil {
		return 0, err
	}
	light := LightData{
		SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
		BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
		SkyLight:       []pk.ByteArray{},
		BlockLight:     []pk.ByteArray{},
	}
	for i, v := range c.Sections {
		if v.SkyLight != nil {
			light.SkyLightMask.Set(int(i), true)
			light.SkyLight = append(light.SkyLight, v.SkyLight)
		}
		if v.BlockLight != nil {
			light.BlockLightMask.Set(int(i), true)
			light.BlockLight = append(light.BlockLight, v.BlockLight)
		}
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
		pk.Array(c.BlockEntity),
		&light,
	}.WriteTo(w)
}

func (c *Chunk) ReadFrom(r io.Reader) (int64, error) {
	/*
		From https://github.com/maxsupermanhd/WebChunk/blob/7ba5b2394ddc7a8d3ab90c31fb511c920ca2c62c/proxy/chunkProcessor.go#L437
	*/
	var (
		HeightMaps struct {
			MotionBlocking         []uint64 `nbt:"MOTION_BLOCKING"`
			MotionBlockingNoLeaves []uint64 `nbt:"MOTION_BLOCKING_NO_LEAVES"`
			OceanFloor             []uint64 `nbt:"OCEAN_FLOOR"`
			WorldSurface           []uint64 `nbt:"WORLD_SURFACE"`
		}
		Sections pk.ByteArray
	)

	n, err := pk.Tuple{
		pk.NBT(&HeightMaps),
		&Sections,
		pk.Array(&c.BlockEntity),
		&LightData{
			SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
			BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
			SkyLight:       []pk.ByteArray{},
			BlockLight:     []pk.ByteArray{},
		},
	}.ReadFrom(r)

	d := bytes.NewReader(Sections)
	dl := len(Sections)
	c.Sections = make([]Section, 0)
	for {
		if dl == 0 {
			break
		}
		if dl < 200 { // whole chunk structure is 207 if completely empty?
			log.Printf("Leaving %d bytes behind while parsing chunk data!", dl)
			break
		}
		ss := &Section{
			BlockCount: 0,
			States:     NewStatesPaletteContainer(16*16*16, 0),
			Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
		}
		n, err := ss.ReadFrom(d)
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			log.Printf("EOF while decoding chunk data while reading section %d", len(Sections))
			break
		}
		if n == 0 {
			continue
		}
		dl -= int(n)
		c.Sections = append(c.Sections, *ss)
	}

	return n, err
}

func (c *Chunk) Data() ([]byte, error) {
	var buff bytes.Buffer
	for i := range c.Sections {
		section := c.Sections[i]
		_, err := section.WriteTo(&buff)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func (c *Chunk) PutData(data []byte) error {
	r := bytes.NewReader(data)
	for _, section := range c.Sections {
		_, err := section.ReadFrom(r)
		if err != nil {
			return err
		}
	}
	return nil
}

type HeightMaps struct {
	MotionBlocking         *BitStorage
	MotionBlockingNoLeaves *BitStorage
	OceanFloor             *BitStorage
	WorldSurface           *BitStorage
}

type BlockEntity struct {
	XZ   int8
	Y    int16
	Type int32
	Data nbt.RawMessage
}

func (b BlockEntity) UnpackXZ() (X, Z int) {
	return int((uint8(b.XZ) >> 4) & 0xF), int(uint8(b.XZ) & 0xF)
}

func (b BlockEntity) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.Byte(b.XZ),
		pk.Short(b.Y),
		pk.VarInt(b.Type),
		pk.NBT(b.Data),
	}.WriteTo(w)
}

func (b *BlockEntity) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		(*pk.Byte)(&b.XZ),
		(*pk.Short)(&b.Y),
		(*pk.VarInt)(&b.Type),
		pk.NBT(&b.Data),
	}.ReadFrom(r)
}

type Section struct {
	BlockCount int16
	States     *PaletteContainer[BlocksState]
	Biomes     *PaletteContainer[BiomesState]
	// Half a byte per light value.
	// Could be nil if not exist
	SkyLight   []byte // len() == 2048
	BlockLight []byte // len() == 2048
}

func (s *Section) GetBlock(i int) BlocksState {
	return s.States.Get(i)
}

func (s *Section) SetBlock(i int, v BlocksState) {
	if block.IsAir(s.States.Get(i)) {
		s.BlockCount--
	}
	if v != 0 {
		s.BlockCount++
	}
	s.States.Set(i, v)
}

func (s *Section) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Short(s.BlockCount),
		s.States,
		s.Biomes,
	}.WriteTo(w)
}

func (s *Section) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		(*pk.Short)(&s.BlockCount),
		s.States,
		s.Biomes,
	}.ReadFrom(r)
}

type LightData struct {
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

func (l *LightData) WriteTo(w io.Writer) (int64, error) {
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

func (l *LightData) ReadFrom(r io.Reader) (int64, error) {
	var TrustEdges pk.Boolean
	var RevSkyLightMask, RevBlockLightMask pk.BitSet
	return pk.Tuple{
		&TrustEdges, // Trust Edges
		&l.SkyLightMask,
		&l.BlockLightMask,
		&RevSkyLightMask,
		&RevBlockLightMask,
		pk.Array(&l.SkyLight),
		pk.Array(&l.BlockLight),
	}.ReadFrom(r)
}
