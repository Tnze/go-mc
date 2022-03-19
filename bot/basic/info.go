package basic

import (
	"unsafe"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

// WorldInfo content player info in server.
type WorldInfo struct {
	DimensionCodec      nbt.StringifiedMessage
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
	PiglinSafe         int8    `nbt:"piglin_safe"`
	Natural            int8    `nbt:"natural"`
	AmbientLight       float32 `nbt:"ambient_light"`
	FixedTime          *int64  `nbt:"fixed_time"`
	Infiniburn         string  `nbt:"infiniburn"`
	RespawnAnchorWorks int8    `nbt:"respawn_anchor_works"`
	HasSkylight        int8    `nbt:"has_skylight"`
	BedWorks           int8    `nbt:"bed_works"`
	Effects            string  `nbt:"effects"`
	HasRaids           int8    `nbt:"has_raids"`
	LogicalHeight      int32   `nbt:"logical_height"`
	CoordinateScale    float64 `nbt:"coordinate_scale"`
	MinY               int32   `nbt:"min_y"`
	HasCeiling         int8    `nbt:"has_ceiling"`
	Ultrawarm          int8    `nbt:"ultrawarm"`
	Height             int32   `nbt:"height"`
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
	err := packet.Scan(
		(*pk.Int)(&p.EID),
		(*pk.Boolean)(&p.Hardcore),
		(*pk.UnsignedByte)(&p.Gamemode),
		(*pk.Byte)(&p.PrevGamemode),
		pk.Ary[pk.VarInt, *pk.VarInt]{Ary: &WorldNames},
		pk.NBT(&p.WorldInfo.DimensionCodec),
		pk.NBT(&p.WorldInfo.Dimension),
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
