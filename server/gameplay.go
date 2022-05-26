package server

import (
	"context"
	_ "embed"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server/ecs"
)

type GamePlay interface {
	// AcceptPlayer handle everything after "LoginSuccess" is sent.
	//
	// Note: the connection will be closed after this function returned.
	// You don't need to close the connection, but to keep not returning while the player is playing.
	AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn)
}

type Game struct {
	*ecs.World
	*ecs.Dispatcher
	handlers   map[int32][]*PacketHandler
	components []Component
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
	ClientLeft(c *Client, p *Player)
}

func NewGame(components ...Component) *Game {
	g := &Game{
		World:      ecs.NewWorld(),
		Dispatcher: ecs.NewDispatcher(),
		handlers:   make(map[int32][]*PacketHandler),
		components: components,
	}
	for _, c := range components {
		c.Init(g)
	}
	ecs.Register[Client, *ecs.HashMapStorage[Client]](g.World)
	ecs.Register[Player, *ecs.HashMapStorage[Player]](g.World)
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
			g.Dispatcher.Run(g.World)
		case <-ctx.Done():
			return
		}
	}
}

func (g *Game) AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn) {
	eid := g.CreateEntity(
		Client{
			Conn:        conn,
			Protocol:    protocol,
			packetQueue: NewPacketQueue(),
			errChan:     make(chan error, 1),
		},
		Player{
			UUID: id,
			Name: name,
		},
	)
	c := ecs.GetComponent[Client](g.World).GetValue(eid)
	p := ecs.GetComponent[Player](g.World).GetValue(eid)
	defer c.packetQueue.Close()

	go func() {
		for {
			packet, ok := c.packetQueue.Pull()
			if !ok {
				break
			}
			err := c.Conn.WritePacket(packet)
			if err != nil {
				c.PutErr(err)
				break
			}
		}
	}()

	for _, component := range g.components {
		component.ClientJoin(c, p)
		defer component.ClientLeft(c, p)
	}

	var packet pk.Packet
	for {
		if err := c.ReadPacket(&packet); err != nil {
			return
		}
		for _, ph := range g.handlers[packet.ID] {
			if err := ph.F(c, p, Packet758(packet)); err != nil {
				return
			}
			if err := c.GetErr(); err != nil {
				return
			}
		}
	}
}
