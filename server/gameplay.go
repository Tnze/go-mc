package server

import (
	_ "embed"
	"github.com/Tnze/go-mc/nbt"
	"sync/atomic"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/data/packetid"
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
	eid int32
	Dim Dimension
}

//go:embed DimensionCodec.snbt
var dimensionCodecSNBT string

//go:embed Dimension.snbt
var dimensionSNBT string

func (g *Game) AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn) {
	p := &Player{
		Conn:     conn,
		EntityID: g.newEID(),
		Gamemode: 1,
	}
	dimInfo := g.Dim.Info()
	err := p.WritePacket(Packet757(pk.Marshal(
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
	)))
	if err != nil {
		return
	}
	g.Dim.PlayerJoin(p)
}

func (g *Game) newEID() int32 {
	return atomic.AddInt32(&g.eid, 1)
}
