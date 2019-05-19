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
	// Do not send at the same goroutin!
	Delegate chan func() error
	Events   eventBroker

	Inventory [46]entity.Slot
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is useable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() (c *Client) {
	c = new(Client)

	//init Client
	c.settings = DefaultSettings
	c.Name = "Steve"
	c.Delegate = make(chan func() error)

	c.Wd.Entities = make(map[int32]entity.Entity)
	c.Wd.Chunks = make(map[world.ChunkLoc]*world.Chunk)

	return
}

//PlayInfo content player info in server.
type PlayInfo struct {
	Gamemode         int    //游戏模式
	Hardcore         bool   //是否是极限模式
	Dimension        int    //维度
	Difficulty       int    //难度
	LevelType        string //地图类型
	ViewDistance     int    //视距
	ReducedDebugInfo bool   //减少调试信息
	// SpawnPosition    Position //主世界出生点
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

//HotBar return the hotbar of inventory
func (c *Client) HotBar() []entity.Slot {
	return c.Inventory[36:45]
}

// MainInventory return the main inventory slots
func (c *Client) MainInventory() []entity.Slot {
	return c.Inventory[9:36]
}
