package bot

import (
	pk "github.com/Tnze/go-mc/net/packet"
)

type Events struct {
	generic  *handlerHeap           // for every packet
	handlers map[int32]*handlerHeap // for specific packet id only
	tickers  *tickerHeap
}

// AddListener adds a listener to the event.
// The listener will be called when the packet with the same ID is received.
// The listener will be called in the order of priority.
// The listeners cannot have multiple same ID.
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

// AddGeneric adds listeners like AddListener, but the packet ID is ignored.
// Generic listener is always called before specific packet listener.
func (e *Events) AddGeneric(listeners ...PacketHandler) {
	for _, l := range listeners {
		if e.generic == nil {
			e.generic = &handlerHeap{l}
		} else {
			e.generic.Push(l)
		}
	}
}

type TickHandler struct {
	Priority int
	F        func(*Client) error
}

func (e *Events) AddTicker(tickers ...TickHandler) {
	for _, t := range tickers {
		if e.tickers == nil {
			e.tickers = &tickerHeap{t}
		} else {
			e.tickers.Push(t)
		}
	}
}

type PacketHandlerFunc func(*Client, pk.Packet) error
type PacketHandler struct {
	ID       int32
	Priority int
	F        PacketHandlerFunc
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

// tickerHeap is PriorityQueue<TickHandlerFunc>
type tickerHeap []TickHandler

func (h tickerHeap) Len() int { return len(h) }
func (h tickerHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}
func (h tickerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *tickerHeap) Push(x interface{}) {
	*h = append(*h, x.(TickHandler))
}
func (h *tickerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[0 : n-1]
	return old[n-1]
}
