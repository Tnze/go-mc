package server

import (
	"strconv"
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

	errChan chan error
}

// Packet757 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type Packet757 pk.Packet

// WritePacket to player client. The type of parameter will update per version.
func (p *Player) WritePacket(packet Packet757) error {
	p.writeLock.Lock()
	defer p.writeLock.Unlock()
	err := p.Conn.WritePacket(pk.Packet(packet))
	if err != nil {
		return WritePacketError{Err: err, ID: packet.ID}
	}
	return nil
}

type WritePacketError struct {
	Err error
	ID  int32
}

func (s WritePacketError) Error() string {
	return "server: send packet " + strconv.FormatInt(int64(s.ID), 16) + " error: " + s.Err.Error()
}

func (s WritePacketError) Unwrap() error {
	return s.Err
}

type PacketHandler struct {
	ID int32
	F  packetHandlerFunc
}

type packetHandlerFunc func(packet Packet757) error

func (p *Player) AddHandler(ph PacketHandler) {
	if p.handlers == nil {
		p.handlers = make(map[int32][]packetHandlerFunc)
	}
	p.handlers[ph.ID] = append(p.handlers[ph.ID], ph.F)
}

func (p *Player) PutErr(err error) {
	select {
	case p.errChan <- err:
	default:
		// previous error exist, ignore this.
	}
}

func (p *Player) GetErr() error {
	select {
	case err := <-p.errChan:
		return err
	default:
		return nil
	}
}
