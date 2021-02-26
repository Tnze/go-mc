package bot

import (
	pk "github.com/Tnze/go-mc/net/packet"
)

type Events struct {
	handlers map[int32]*handlerHeap
}

func (e *Events) AddListener(listeners ...PacketHandler) {
	for _, l := range listeners {
		var s *handlerHeap
		var ok bool
		if s, ok = e.handlers[l.ID]; !ok {
			s = &handlerHeap{l}
			e.handlers[l.ID] = s
		} else {
			s.Push(l)
		}
	}
}

type PacketHandlerFunc func(p pk.Packet) error
type PacketHandler struct {
	ID       int32
	Priority int
	F        func(p pk.Packet) error
}

// handlerHeap is PriorityQueue<PacketHandlerFunc>
type handlerHeap []PacketHandler

func (h handlerHeap) Len() int            { return len(h) }
func (h handlerHeap) Less(i, j int) bool  { return h[i].Priority < h[j].Priority }
func (h handlerHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *handlerHeap) Push(x interface{}) { *h = append(*h, x.(PacketHandler)) }
func (h *handlerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[0 : n-1]
	return old[n-1]
}
