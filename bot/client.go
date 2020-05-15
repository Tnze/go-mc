package bot

import (
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/Tnze/go-mc/net"
)

// Client is used to access Minecraft server
type Client struct {
	conn *net.Conn
	Auth

	player.Player
	PlayInfo
	abilities PlayerAbilities
	settings  Settings
	Wd        world.World //the map data

	// Delegate allows you push a function to let HandleGame run.
	// Do not send at the same goroutine!
	Delegate chan func() error
	Events   eventBroker
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() (c *Client) {
	c = new(Client)

	//init Client
	c.settings = DefaultSettings
	c.Name = "Steve"
	c.Delegate = make(chan func() error)

	c.Wd = world.World{
		Entities: make(map[int32]entity.Entity),
		Chunks:   make(map[world.ChunkLoc]*world.Chunk),
	}

	return
}

//PlayInfo content player info in server.
type PlayInfo struct {
	Gamemode         int    // Game mode (self explanatory)
	Hardcore         bool   // Hardcode (self explanatory)
	Dimension        int    // The currently loaded Dimension (DIM in gamefiles)
	Difficulty       int    // Difficulty, 0 peaceful -> 3 hard
	LevelType        string // type of map that is generated (NORMAL for vanilla, BIOMESOP for biomes o plenty and etc)
	ViewDistance     int    // radius in which chunks are loaded
	ReducedDebugInfo bool   // Reduce debug info. This settigns changes how much debug info is displayed on vanilla clients
	// SpawnPosition    Position // Worldspawn
}


// PlayerAbilities defines what player can do.
type PlayerAbilities struct {
	Flags               int8
	FlyingSpeed         float32
	FieldofViewModifier float32
}

//Position is a 3D vector.
type Position struct {
	X, Y, Z int
}
