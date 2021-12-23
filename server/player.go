package server

import (
	"sync"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Player struct {
	*net.Conn
	writeLock sync.Mutex

	Name string
	uuid.UUID
	EntityID int32
	Gamemode byte
	handlers map[int32][]packetHandlerFunc
}

// Packet757 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type Packet757 pk.Packet

// WritePacket to player client. The type of parameter will update per version.
func (p *Player) WritePacket(packet Packet757) error {
	p.writeLock.Lock()
	defer p.writeLock.Unlock()
	return p.Conn.WritePacket(pk.Packet(packet))
}

type PacketHandler struct {
	ID int32
	F  packetHandlerFunc
}

type packetHandlerFunc func(packet Packet757) error

func (p *Player) Add(ph PacketHandler) {

	if p.handlers == nil {
		p.handlers = make(map[int32][]packetHandlerFunc)
	}
	p.handlers[ph.ID] = append(p.handlers[ph.ID], ph.F)
}

func (p *Player) PutErr(err error) {
	// TODO: handle errors
}
