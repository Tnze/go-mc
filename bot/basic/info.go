package basic

import (
	"github.com/Tnze/go-mc/chat"
	"unsafe"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

// WorldInfo content player info in server.
type WorldInfo struct {
	RegistryCodec       RegistryCodec
	DimensionType       string
	DimensionNames      []string // Identifiers for all worlds on the server.
	DimensionName       string   // Name of the world being spawned into.
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

type ChatType struct {
	Chat      chat.Decoration `nbt:"chat"`
	Narration chat.Decoration `nbt:"narration"`
}

type RegistryCodec struct {
	// What is Below? (wiki.vg)
	ChatType      Registry[ChatType]       `nbt:"minecraft:chat_type"`
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

func (r *Registry[E]) FindByID(id int32) *E {
	if id >= 0 && id < int32(len(r.Value)) && r.Value[id].ID == id {
		return &r.Value[id].Element
	}
	for i := range r.Value {
		if r.Value[i].ID == id {
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

func (p *Player) handleLoginPacket(packet pk.Packet) error {
	err := packet.Scan(
		(*pk.Int)(&p.EID),
		(*pk.Boolean)(&p.Hardcore),
		(*pk.UnsignedByte)(&p.Gamemode),
		(*pk.Byte)(&p.PrevGamemode),
		pk.Array((*[]pk.Identifier)(unsafe.Pointer(&p.DimensionNames))),
		pk.NBT(&p.WorldInfo.RegistryCodec),
		(*pk.Identifier)(&p.WorldInfo.DimensionType),
		(*pk.Identifier)(&p.DimensionName),
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
	err = p.c.Conn.WritePacket(pk.Marshal( //PluginMessage packet
		int32(packetid.ServerboundCustomPayload),
		pk.Identifier("minecraft:brand"),
		pk.String(p.Settings.Brand),
	))
	if err != nil {
		return Error{err}
	}

	err = p.c.Conn.WritePacket(pk.Marshal(
		int32(packetid.ServerboundClientInformation), // Client settings
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
	err := packet.Scan(
		(*pk.String)(&p.DimensionType),
		(*pk.Identifier)(&p.DimensionName),
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
	return nil
}
