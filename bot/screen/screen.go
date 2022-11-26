package screen

import (
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type Manager struct {
	Screens       map[int]Container
	CurrentScreen int
	Inventory     Inventory
	Cursor        slots.Slot
	// The last received State ID from server
	StateID int32
}

func NewManager() *Manager {
	m := &Manager{
		Screens: make(map[int]Container),
	}
	m.Screens[0] = &m.Inventory
	return m
}

type ChangedSlots map[int]*slots.Slot

func (c ChangedSlots) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.VarInt(len(c)).WriteTo(w)
	if err != nil {
		return
	}
	for i, v := range c {
		n1, err := pk.Short(i).WriteTo(w)
		if err != nil {
			return n + n1, err
		}
		n2, err := v.WriteTo(w)
		if err != nil {
			return n + n1 + n2, err
		}
		n += n1 + n2
	}
	return
}

type Container interface {
	OnSetSlot(i int, s slots.Slot) error
	onClose() error
}

type Error struct {
	Err error
}

func (e Error) Error() string {
	return "bot/screen: " + e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
