package server

import (
	"context"
	"crypto/rsa"
	_ "embed"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/net"
)

type GamePlay interface {
	// AcceptPlayer handle everything after "LoginSuccess" is sent.
	//
	// Note: the connection will be closed after this function returned.
	// You don't need to close the connection, but to keep not returning while the player is playing.
	AcceptPlayer(name string, id uuid.UUID, profilePubKey *rsa.PublicKey, protocol int32, conn *net.Conn)
}

type Game struct {
	WorldLocker sync.Mutex
	handlers    map[int32][]*PacketHandler
	components  []Component
}

func (g *Game) AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn) {
	conn.Close()
}

type PacketHandler struct {
	ID int32
	F  packetHandlerFunc
}

type packetHandlerFunc func(client *Client, player *Player, packet Packet758) error

type Component interface {
	Init(g *Game)
	Run(ctx context.Context)
	ClientJoin(c *Client, p *Player)
	ClientLeft(c *Client, p *Player, reason error)
}

func NewGame(components ...Component) *Game {
	g := &Game{
		handlers:   make(map[int32][]*PacketHandler),
		components: components,
	}
	for _, c := range components {
		c.Init(g)
	}
	return g
}

func (g *Game) AddHandler(ph *PacketHandler) {
	g.handlers[ph.ID] = append(g.handlers[ph.ID], ph)
}

func (g *Game) Run(ctx context.Context) {
	for _, c := range g.components {
		go c.Run(ctx)
	}
	ticker := time.NewTicker(time.Second / 20)
	for {
		select {
		case <-ticker.C:
			g.WorldLocker.Lock()
			//g.Dispatcher.Run(g.World)
			g.WorldLocker.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
