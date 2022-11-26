package bot

import (
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/yggdrasil/user"
)

// Client is used to access Minecraft server
type Client struct {
	Conn    *net.Conn
	Auth    Auth
	KeyPair user.KeyPairResp

	World  *world.World
	Player *Player
	TPS    *maths.TpsCalculator

	EventHandlers EventsListener
	Events        Events
	LoginPlugin   map[string]func(data []byte) ([]byte, error)
}

func (c *Client) Close() error {
	return c.Conn.Close()
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() *Client {
	c := &Client{
		Auth:   Auth{Name: "Steve"},
		Events: Events{handlers: make(map[int32]*handlerHeap)},
	}
	c.Player = NewPlayer(c, basic.DefaultSettings)
	c.World = world.NewWorld()
	c.TPS = new(maths.TpsCalculator)
	return c
}
