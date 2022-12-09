package queue

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type Queue[T any] interface {
	Push(v T)
	Pull() (v T, ok bool)
	Close()
}

func NewLinkedQueue[T any]() (p Queue[T]) {
	p = &LinkedListPacketQueue[T]{
		queue: list.New(),
		cond:  sync.Cond{L: new(sync.Mutex)},
	}
	return p
}

type LinkedListPacketQueue[T any] struct {
	queue  *list.List
	closed bool
	cond   sync.Cond
}

func (p *LinkedListPacketQueue[T]) Push(v T) {
	p.cond.L.Lock()
	if !p.closed {
		p.queue.PushBack(v)
	}
	p.cond.Signal()
	p.cond.L.Unlock()
}

func (p *LinkedListPacketQueue[T]) Pull() (v T, ok bool) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for p.queue.Front() == nil && !p.closed {
		p.cond.Wait()
	}
	if p.closed {
		return
	}
	v = p.queue.Remove(p.queue.Front()).(T)
	ok = true
	return
}

func (p *LinkedListPacketQueue[T]) Close() {
	p.cond.L.Lock()
	p.closed = true
	p.cond.Broadcast()
	p.cond.L.Unlock()
}

type ChannelPacketQueue[T any] struct {
	c      chan T
	closed atomic.Bool
}

func (c ChannelPacketQueue[T]) Push(v T) {
	if c.closed.Load() {
		return
	}
	select {
	case c.c <- v:
	default:
		c.closed.Store(true)
	}
}

func (c ChannelPacketQueue[T]) Pull() (v T, ok bool) {
	if !c.closed.Load() {
		v, ok = <-c.c
	}
	return
}

func (c ChannelPacketQueue[T]) Close() {
	c.closed.Store(true)
}
