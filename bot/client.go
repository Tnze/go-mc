package bot

import (
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/Tnze/go-mc/net"
)

// Client is the Object used to access Minecraft server
type Client struct {
	conn *net.Conn
	Auth

	player.Player
	PlayInfo
	abilities PlayerAbilities
	settings  Settings
	// wd        world //the map data

}

//NewClient init and return a new Client
func NewClient() (c *Client) {
	c = new(Client)

	//init Client
	c.settings = DefaultSettings
	c.Name = "Steve"

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
