package server

import (
	"container/list"
	"strconv"
	"sync"
	"sync/atomic"

	pk "github.com/Tnze/go-mc/net/packet"
)

// Packet758 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type (
	Packet758 pk.Packet
	Packet757 pk.Packet
)

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

type PacketQueue interface {
	Push(packet pk.Packet)
	Pull() (packet pk.Packet, ok bool)
	Close()
}

func NewPacketQueue() (p PacketQueue) {
	p = &LinkedListPacketQueue{
		queue: list.New(),
		cond:  sync.Cond{L: new(sync.Mutex)},
	}
	return p
}

type LinkedListPacketQueue struct {
	queue  *list.List
	closed bool
	cond   sync.Cond
}

func (p *LinkedListPacketQueue) Push(packet pk.Packet) {
	p.cond.L.Lock()
	if !p.closed {
		p.queue.PushBack(packet)
	}
	p.cond.Signal()
	p.cond.L.Unlock()
}

func (p *LinkedListPacketQueue) Pull() (packet pk.Packet, ok bool) {
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

func (p *LinkedListPacketQueue) Close() {
	p.cond.L.Lock()
	p.closed = true
	p.cond.Broadcast()
	p.cond.L.Unlock()
}

type ChannelPacketQueue struct {
	c      chan pk.Packet
	closed atomic.Bool
}

func (c ChannelPacketQueue) Push(packet pk.Packet) {
	if c.closed.Load() {
		return
	}
	select {
	case c.c <- packet:
	default:
		c.closed.Store(true)
	}
}

func (c ChannelPacketQueue) Pull() (packet pk.Packet, ok bool) {
	if !c.closed.Load() {
		packet, ok = <-c.c
	}
	return
}

func (c ChannelPacketQueue) Close() {
	c.closed.Store(true)
}
