package server

import (
	"context"
	_ "embed"
	"sync"
	"sync/atomic"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

type GamePlay interface {
	// AcceptPlayer handle everything after "LoginSuccess" is sent.
	//
	// Note: the connection will be closed after this function returned.
	// You don't need to close the connection, but to keep not returning while the player is playing.
	AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn)
}

type Game struct {
	Dim        Level
	components []Component
	handlers   map[int32][]*PacketHandler

	eid int32
}

type Component interface {
	Init(g *Game)
	Run(ctx context.Context)
	AddPlayer(p *Player)
	RemovePlayer(p *Player)
}

type PacketHandler struct {
	ID int32
	F  packetHandlerFunc
}

type packetHandlerFunc func(player *Player, packet Packet757) error

//go:embed DimensionCodec.snbt
var dimensionCodecSNBT string

//go:embed Dimension.snbt
var dimensionSNBT string

func NewGame(dim Level, components ...Component) *Game {
	g := &Game{
		Dim:        dim,
		components: components,
		handlers:   make(map[int32][]*PacketHandler),
	}
	for _, v := range components {
		v.Init(g)
	}
	return g
}

func (g *Game) AddHandler(ph *PacketHandler) {
	g.handlers[ph.ID] = append(g.handlers[ph.ID], ph)
}

func (g *Game) Run(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(len(g.components))
	for _, c := range g.components {
		go func(c Component) {
			defer wg.Done()
			c.Run(ctx)
		}(c)
	}
	wg.Wait()
}

func (g *Game) AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn) {
	p := &Player{
		Conn:        conn,
		Name:        name,
		UUID:        id,
		EntityID:    g.newEID(),
		Gamemode:    1,
		packetQueue: NewPacketQueue(),
		errChan:     make(chan error, 1),
	}
	dimInfo := g.Dim.Info()
	err := p.Conn.WritePacket(pk.Marshal(
		packetid.ClientboundLogin,
		pk.Int(p.EntityID),  // Entity ID
		pk.Boolean(false),   // Is hardcore
		pk.Byte(p.Gamemode), // Gamemode
		pk.Byte(-1),         // Prev Gamemode
		pk.Array([]pk.Identifier{pk.Identifier(dimInfo.Name)}),
		pk.NBT(nbt.StringifiedMessage(dimensionCodecSNBT)),
		pk.NBT(nbt.StringifiedMessage(dimensionSNBT)),
		pk.Identifier(dimInfo.Name), // World Name
		pk.Long(dimInfo.HashedSeed), // Hashed seed
		pk.VarInt(0),                // Max Players (Ignored by client)
		pk.VarInt(15),               // View Distance
		pk.VarInt(15),               // Simulation Distance
		pk.Boolean(false),           // Reduced Debug Info
		pk.Boolean(true),            // Enable respawn screen
		pk.Boolean(false),           // Is Debug
		pk.Boolean(true),            // Is Flat
	))
	if err != nil {
		return
	}

	go func() {
		for {
			packet, ok := p.packetQueue.Pull()
			if !ok {
				break
			}
			err := p.Conn.WritePacket(packet)
			if err != nil {
				p.PutErr(err)
				break
			}
		}
	}()
	defer p.packetQueue.Close()

	g.Dim.PlayerJoin(p)
	defer g.Dim.PlayerQuit(p)

	for _, c := range g.components {
		c.AddPlayer(p)
		if err := p.GetErr(); err != nil {
			return
		}
		//goland:noinspection GoDeferInLoop
		defer c.RemovePlayer(p)
	}

	var packet pk.Packet
	for {
		if err := p.ReadPacket(&packet); err != nil {
			return
		}
		for _, ph := range g.handlers[packet.ID] {
			if err := ph.F(p, Packet757(packet)); err != nil {
				return
			}
			if err := p.GetErr(); err != nil {
				return
			}
		}
	}
}

func (g *Game) newEID() int32 {
	return atomic.AddInt32(&g.eid, 1)
}
