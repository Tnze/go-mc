package bot

import (
	"sync"
	"time"

	"github.com/Tnze/go-mc/bot/phy"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// Client is used to access Minecraft server
type Client struct {
	conn *net.Conn
	Auth

	player.Player
	PlayInfo
	ServInfo
	abilities PlayerAbilities
	settings  Settings

	Wd        world.World //the map data
	Inputs    phy.Inputs
	Physics   phy.State
	lastPosTx time.Time

	// Delegate allows you push a function to let HandleGame run.
	// Do not send at the same goroutine!
	Delegate chan func() error
	Events   eventBroker

	closing chan struct{}
	inbound chan pk.Packet
	wg      sync.WaitGroup
}

func (c *Client) Close() error {
	close(c.closing)
	err := c.disconnect()
	c.wg.Wait()
	return err
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() *Client {
	return &Client{
		settings: DefaultSettings,
		Auth:     Auth{Name: "Steve"},
		Delegate: make(chan func() error),
		Wd: world.World{
			Entities: make(map[int32]entity.Entity),
			Chunks:   make(map[world.ChunkLoc]*world.Chunk),
		},
		closing: make(chan struct{}),
		inbound: make(chan pk.Packet, 5),
	}
}

//PlayInfo content player info in server.
type PlayInfo struct {
	Gamemode         int      //游戏模式
	Hardcore         bool     //是否是极限模式
	Dimension        int      //维度
	Difficulty       int      //难度
	ViewDistance     int      //视距
	ReducedDebugInfo bool     //减少调试信息
	WorldName        string   //当前世界的名字
	IsDebug          bool     //调试
	IsFlat           bool     //超平坦世界
	SpawnPosition    Position //主世界出生点
}

// ServInfo contains information about the server implementation.
type ServInfo struct {
	Brand string
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
