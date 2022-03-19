package screen

import (
	"errors"
	"io"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Manager struct {
	c *bot.Client

	Screens   map[int]Container
	Inventory Inventory
	Cursor    Slot
	events    EventsListener
	// The last received State ID from server
	stateID int32
}

func NewManager(c *bot.Client, e EventsListener) *Manager {
	m := &Manager{
		c:       c,
		Screens: make(map[int]Container),
		events:  e,
	}
	m.Screens[0] = &m.Inventory
	c.Events.AddListener(
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundOpenScreen, F: m.onOpenScreen},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundContainerSetContent, F: m.onSetContentPacket},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundContainerClose, F: m.onCloseScreen},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundContainerSetSlot, F: m.onSetSlot},
	)
	return m
}

type ChangedSlots map[int]*Slot

func (m *Manager) ContainerClick(id int, slot int16, button byte, mode int32, slots ChangedSlots, carried *Slot) error {
	return m.c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundContainerClick,
		pk.UnsignedByte(id),
		pk.VarInt(m.stateID),
		pk.Short(slot),
		pk.Byte(button),
		pk.VarInt(mode),
		slots,
		carried,
	))
}

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

func (m *Manager) onOpenScreen(p pk.Packet) error {
	var (
		ContainerID pk.VarInt
		Type        pk.VarInt
		Title       chat.Message
	)
	if err := p.Scan(&ContainerID, &Type, &Title); err != nil {
		return Error{err}
	}
	//if c, ok := m.Screens[byte(ContainerID)]; ok {
	// TODO: Create the specified container
	//}
	if m.events.Open != nil {
		if err := m.events.Open(int(ContainerID)); err != nil {
			return Error{err}
		}
	}
	return nil
}

func (m *Manager) onSetContentPacket(p pk.Packet) error {
	var (
		ContainerID pk.UnsignedByte
		StateID     pk.VarInt
		SlotData    []Slot
		CarriedItem Slot
	)
	if err := p.Scan(
		&ContainerID,
		&StateID,
		pk.Ary[pk.VarInt, *pk.VarInt]{
			Ary: &SlotData,
		},
		&CarriedItem,
	); err != nil {
		return Error{err}
	}
	m.stateID = int32(StateID)
	// copy the slot data to container
	container, ok := m.Screens[int(ContainerID)]
	if !ok {
		return Error{errors.New("setting content of non-exist container")}
	}
	for i, v := range SlotData {
		err := container.onSetSlot(i, v)
		if err != nil {
			return Error{err}
		}
		if m.events.SetSlot != nil {
			if err := m.events.SetSlot(int(ContainerID), i); err != nil {
				return Error{err}
			}
		}
	}
	return nil
}

func (m *Manager) onCloseScreen(p pk.Packet) error {
	var ContainerID pk.UnsignedByte
	if err := p.Scan(&ContainerID); err != nil {
		return Error{err}
	}
	if c, ok := m.Screens[int(ContainerID)]; ok {
		delete(m.Screens, int(ContainerID))
		if err := c.onClose(); err != nil {
			return Error{err}
		}
		if m.events.Close != nil {
			if err := m.events.Close(int(ContainerID)); err != nil {
				return Error{err}
			}
		}
	}
	return nil
}

func (m *Manager) onSetSlot(p pk.Packet) (err error) {
	var (
		ContainerID pk.Byte
		StateID     pk.VarInt
		SlotID      pk.Short
		SlotData    Slot
	)
	if err := p.Scan(&ContainerID, &StateID, &SlotID, &SlotData); err != nil {
		return Error{err}
	}

	m.stateID = int32(StateID)
	if ContainerID == -1 && SlotID == -1 {
		m.Cursor = SlotData
	} else if ContainerID == -2 {
		err = m.Inventory.onSetSlot(int(SlotID), SlotData)
	} else if c, ok := m.Screens[int(ContainerID)]; ok {
		err = c.onSetSlot(int(SlotID), SlotData)
	}

	if m.events.SetSlot != nil {
		if err := m.events.SetSlot(int(ContainerID), int(SlotID)); err != nil {
			return Error{err}
		}
	}
	if err != nil {
		return Error{err}
	}
	return nil
}

type Slot struct {
	ID    pk.VarInt
	Count pk.Byte
	NBT   nbt.RawMessage
}

func (s *Slot) WriteTo(w io.Writer) (n int64, err error) {
	var present pk.Boolean = s != nil
	return pk.Tuple{
		present, pk.Opt{
			Has: present,
			Field: pk.Tuple{
				&s.ID, &s.Count, pk.NBT(&s.NBT),
			},
		},
	}.WriteTo(w)
}

func (s *Slot) ReadFrom(r io.Reader) (n int64, err error) {
	var present pk.Boolean
	return pk.Tuple{
		&present, pk.Opt{
			Has: &present,
			Field: pk.Tuple{
				&s.ID, &s.Count, pk.NBT(&s.NBT),
			},
		},
	}.ReadFrom(r)
}

type Container interface {
	onSetSlot(i int, s Slot) error
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
