package server

import (
	"container/list"
	"github.com/google/uuid"
	"strconv"
	"sync"

	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Player struct {
	*net.Conn

	Name string
	uuid.UUID
	EntityID int32
	Gamemode byte

	packetQueue *PacketQueue
	errChan     chan error
}

func NewPlayer(conn *net.Conn, name string, id uuid.UUID, eid int32, gamemode byte) (p *Player) {
	p = &Player{
		Conn:        conn,
		Name:        name,
		UUID:        id,
		EntityID:    eid,
		Gamemode:    gamemode,
		packetQueue: NewPacketQueue(),
		errChan:     make(chan error, 1),
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
	return
}

func (p *Player) Close() {
	p.packetQueue.Close()
}

// Packet758 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type Packet758 pk.Packet
type Packet757 pk.Packet

// WritePacket to player client. The type of parameter will update per version.
func (p *Player) WritePacket(packet Packet758) {
	p.packetQueue.Push(pk.Packet(packet))
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

type PacketQueue struct {
	queue  *list.List
	closed bool
	cond   sync.Cond
}

func NewPacketQueue() (p *PacketQueue) {
	p = &PacketQueue{
		queue: list.New(),
		cond:  sync.Cond{L: new(sync.Mutex)},
	}
	return p
}

func (p *PacketQueue) Push(packet pk.Packet) {
	p.cond.L.Lock()
	if !p.closed {
		p.queue.PushBack(packet)
	}
	p.cond.Signal()
	p.cond.L.Unlock()
}

func (p *PacketQueue) Pull() (packet pk.Packet, ok bool) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for p.queue.Front() == nil && !p.closed {
		p.cond.Wait()
	}
	if p.closed {
		return pk.Packet{}, false
	}
	packet = p.queue.Remove(p.queue.Front()).(pk.Packet)
	ok = true
	return
}

func (p *PacketQueue) Close() {
	p.cond.L.Lock()
	p.closed = true
	p.cond.Broadcast()
	p.cond.L.Unlock()
}
