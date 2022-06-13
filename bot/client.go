package bot

import (
	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/yggdrasil/userApi"
	"github.com/google/uuid"
)

// Client is used to access Minecraft server
type Client struct {
	Conn    *net.Conn
	Auth    Auth
	KeyPair *userApi.KeyPairResp

	Name string
	UUID uuid.UUID

	Events      Events
	LoginPlugin map[string]func(data []byte) ([]byte, error)
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
	return &Client{
		Auth:   Auth{Name: "Steve"},
		Events: Events{handlers: make(map[int32]*handlerHeap)},
	}
}

//Position is a 3D vector.
type Position struct {
	X, Y, Z int
}
