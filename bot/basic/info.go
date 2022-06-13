package basic

import (
	"golang.org/x/exp/slices"
	"unsafe"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// WorldInfo content player info in server.
type WorldInfo struct {
	DimensionCodec      DimensionCodec
	Dimension           Dimension
	WorldNames          []string // Identifiers for all worlds on the server.
	WorldName           string   // Name of the world being spawned into.
	HashedSeed          int64    // First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise
	MaxPlayers          int32    // Was once used by the client to draw the player list, but now is ignored.
	ViewDistance        int32    // Render distance (2-32).
	SimulationDistance  int32    // The distance that the client will process specific things, such as entities.
	ReducedDebugInfo    bool     // If true, a Notchian client shows reduced information on the debug screen. For servers in development, this should almost always be false.
	EnableRespawnScreen bool     // Set to false when the doImmediateRespawn gamerule is true.
	IsDebug             bool     // True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks.
	IsFlat              bool     // True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63.
}
type Dimension struct {
	Name    string `nbt:"name"`
	Id      int    `nbt:"id"`
	Element struct {
		PiglinSafe                  byte    `nbt:"piglin_safe"`
		Natural                     byte    `nbt:"natural"`
		AmbientLight                float64 `nbt:"ambient_light"`
		MonsterSpawnBlockLightLimit int     `nbt:"monster_spawn_block_light_limit"`
		Infiniburn                  string  `nbt:"infiniburn"`
		RespawnAnchorWorks          byte    `nbt:"respawn_anchor_works"`
		HasSkylight                 byte    `nbt:"has_skylight"`
		BedWorks                    byte    `nbt:"bed_works"`
		Effects                     string  `nbt:"effects"`
		HasRaids                    byte    `nbt:"has_raids"`
		Shrunk                      byte    `nbt:"shrunk"`
		LogicalHeight               int     `nbt:"logical_height"`
		CoordinateScale             float64 `nbt:"coordinate_scale"`
		MinY                        int     `nbt:"min_y"`
		MonsterSpawnLightLevel      int     `nbt:"monster_spawn_light_level"`
		Ultrawarm                   byte    `nbt:"ultrawarm"`
		HasCeiling                  byte    `nbt:"has_ceiling"`
		Height                      int     `nbt:"height"`
		FixedTime                   int64   `nbt:"fixed_time,omitempty"`
	} `nbt:"element"`
}
type DimensionCodec struct {
	// What is Below (wik.vg)
	ChatType struct {
		Type  string `nbt:"type"`
		Value []struct {
			Name    string `nbt:"name"`
			Id      int    `nbt:"id"`
			Element struct {
				Chat struct {
					Decoration struct {
						TranslationKey interface{} `nbt:"translation_key"`
						Style          struct {
							Color  interface{} `nbt:"color,omitempty"`
							Italic interface{} `nbt:"italic,omitempty"`
						} `nbt:"style"`
						Parameters []interface{} `nbt:"parameters"`
					} `nbt:"decoration,omitempty"`
				} `nbt:"chat,omitempty"`
				Narration struct {
					Priority   interface{} `nbt:"priority"`
					Decoration struct {
						TranslationKey interface{} `nbt:"translation_key"`
						Style          struct {
						} `nbt:"style"`
						Parameters []interface{} `nbt:"parameters"`
					} `nbt:"decoration,omitempty"`
				} `nbt:"narration,omitempty"`
				Overlay struct {
				} `nbt:"overlay,omitempty"`
			} `nbt:"element"`
		} `nbt:"value"`
	} `nbt:"minecraft:chat_type"`
	DimensionType struct {
		Type  string      `nbt:"type"`
		Value []Dimension `nbt:"value"`
	} `nbt:"minecraft:dimension_type"`
	WorldGenBiome struct {
		Type  string `nbt:"type"`
		Value []struct {
			Name    string `nbt:"name"`
			Id      int    `nbt:"id"`
			Element struct {
				Precipitation interface{} `nbt:"precipitation"`
				Effects       struct {
					SkyColor      int `nbt:"sky_color"`
					WaterFogColor int `nbt:"water_fog_color"`
					FogColor      int `nbt:"fog_color"`
					WaterColor    int `nbt:"water_color"`
					MoodSound     struct {
						TickDelay         int         `nbt:"tick_delay"`
						Offset            interface{} `nbt:"offset"`
						Sound             string      `nbt:"sound"`
						BlockSearchExtent int         `nbt:"block_search_extent"`
					} `nbt:"mood_sound"`
					GrassColorModifier interface{} `nbt:"grass_color_modifier,omitempty"`
					Music              struct {
						ReplaceCurrentMusic interface{} `nbt:"replace_current_music"`
						MaxDelay            int         `nbt:"max_delay"`
						Sound               string      `nbt:"sound"`
						MinDelay            int         `nbt:"min_delay"`
					} `nbt:"music,omitempty"`
					FoliageColor   int    `nbt:"foliage_color,omitempty"`
					GrassColor     int    `nbt:"grass_color,omitempty"`
					AmbientSound   string `nbt:"ambient_sound,omitempty"`
					AdditionsSound struct {
						Sound      string      `nbt:"sound"`
						TickChance interface{} `nbt:"tick_chance"`
					} `nbt:"additions_sound,omitempty"`
					Particle struct {
						Probability interface{} `nbt:"probability"`
						Options     struct {
							Type string `nbt:"type"`
						} `nbt:"options"`
					} `nbt:"particle,omitempty"`
				} `nbt:"effects"`
				Temperature         interface{} `nbt:"temperature"`
				Downfall            interface{} `nbt:"downfall"`
				TemperatureModifier interface{} `nbt:"temperature_modifier,omitempty"`
			} `nbt:"element"`
		} `nbt:"value"`
	} `nbt:"minecraft:worldgen/biome"`
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

func (p *Player) handleLoginPacket(packet pk.Packet) error {
	var WorldNames = make([]pk.Identifier, 0)
	var currentDimension pk.Identifier
	err := packet.Scan(
		(*pk.Int)(&p.EID),
		(*pk.Boolean)(&p.Hardcore),
		(*pk.UnsignedByte)(&p.Gamemode),
		(*pk.Byte)(&p.PrevGamemode),
		pk.Array(&WorldNames),
		pk.NBT(&p.WorldInfo.DimensionCodec),
		&currentDimension,
		(*pk.Identifier)(&p.WorldName),
		(*pk.Long)(&p.HashedSeed),
		(*pk.VarInt)(&p.MaxPlayers),
		(*pk.VarInt)(&p.ViewDistance),
		(*pk.VarInt)(&p.SimulationDistance),
		(*pk.Boolean)(&p.ReducedDebugInfo),
		(*pk.Boolean)(&p.EnableRespawnScreen),
		(*pk.Boolean)(&p.IsDebug),
		(*pk.Boolean)(&p.IsFlat),
	)
	if err != nil {
		return Error{err}
	}
	// This line should work "like" the following code but without copy things
	//	p.WorldNames = make([]string, len(WorldNames))
	//	for i, v := range WorldNames {
	//		p.WorldNames[i] = string(v)
	//	}
	p.WorldNames = *(*[]string)(unsafe.Pointer(&WorldNames))
	index := slices.IndexFunc(p.DimensionCodec.DimensionType.Value, func(d Dimension) bool {
		return d.Name == string(currentDimension)
	})
	if index != -1 {
		p.Dimension = p.DimensionCodec.DimensionType.Value[index]
	}

	err = p.c.Conn.WritePacket(pk.Marshal( //PluginMessage packet
		packetid.ServerboundCustomPayload,
		pk.Identifier("minecraft:brand"),
		pk.String(p.Settings.Brand),
	))
	if err != nil {
		return Error{err}
	}

	err = p.c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundClientInformation, // Client settings
		pk.String(p.Settings.Locale),
		pk.Byte(p.Settings.ViewDistance),
		pk.VarInt(p.Settings.ChatMode),
		pk.Boolean(p.Settings.ChatColors),
		pk.UnsignedByte(p.Settings.DisplayedSkinParts),
		pk.VarInt(p.Settings.MainHand),
		pk.Boolean(p.Settings.EnableTextFiltering),
		pk.Boolean(p.Settings.AllowListing),
	))
	if err != nil {
		return Error{err}
	}
	return nil
}
func (p *Player) handleRespawnPacket(packet pk.Packet) error {
	var copyMeta bool
	var currentDimension pk.Identifier
	err := packet.Scan(
		//pk.NBT(&p.WorldInfo.Dimension),
		&currentDimension,
		(*pk.Identifier)(&p.WorldName),
		(*pk.Long)(&p.HashedSeed),
		(*pk.UnsignedByte)(&p.Gamemode),
		(*pk.Byte)(&p.PrevGamemode),
		(*pk.Boolean)(&p.IsDebug),
		(*pk.Boolean)(&p.IsFlat),
		(*pk.Boolean)(&copyMeta),
	)
	if err != nil {
		return Error{err}
	}
	index := slices.IndexFunc(p.DimensionCodec.DimensionType.Value, func(d Dimension) bool {
		return d.Name == string(currentDimension)
	})
	if index != -1 {
		p.Dimension = p.DimensionCodec.DimensionType.Value[index]
	}
	return nil
}
