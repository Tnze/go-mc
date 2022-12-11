package queue

import (
	"container/list"
	"sync"
)

type Queue[T any] interface {
	Push(v T) (ok bool)
	Pull() (v T, ok bool)
	Close()
}

func NewLinkedQueue[T any]() (q Queue[T]) {
	return &LinkedListQueue[T]{
		queue: list.New(),
		cond:  sync.Cond{L: new(sync.Mutex)},
	}
}

type LinkedListQueue[T any] struct {
	queue  *list.List
	closed bool
	cond   sync.Cond
}

func (p *LinkedListQueue[T]) Push(v T) bool {
	p.cond.L.Lock()
	if p.closed {
		panic("push on closed queue")
	}
	p.queue.PushBack(v)
	p.cond.Signal()
	p.cond.L.Unlock()
	return true
}

func (p *LinkedListQueue[T]) Pull() (v T, ok bool) {
	p.cond.L.Lock()
	for {
		if elem := p.queue.Front(); elem != nil {
			v = p.queue.Remove(elem).(T)
			ok = true
			break
		} else if p.closed {
			break
		}
		p.cond.Wait()
	}
	p.cond.L.Unlock()
	return
}

func (p *LinkedListQueue[T]) Close() {
	p.cond.L.Lock()
	p.closed = true
	p.cond.Broadcast()
	p.cond.L.Unlock()
}

func NewChannelQueue[T any](n int) (q Queue[T]) {
	return make(ChannelQueue[T], n)
}

type ChannelQueue[T any] chan T

func (c ChannelQueue[T]) Push(v T) bool {
	select {
	case c <- v:
		return true
	default:
		return false
	}
}

func (c ChannelQueue[T]) Pull() (v T, ok bool) {
	v, ok = <-c
	return
}

func (c ChannelQueue[T]) Close() {
	close(c)
}
