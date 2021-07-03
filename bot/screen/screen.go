package screen

import (
	"io"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Manager struct {
	Screens map[byte]Container
}

func NewManager(c *bot.Client) *Manager {
	m := &Manager{Screens: make(map[byte]Container)}
	m.Screens[0] = &Inventory{}
	c.Events.AddListener(
		bot.PacketHandler{Priority: 64, ID: packetid.OpenWindow, F: m.onOpenScreen},
		bot.PacketHandler{Priority: 64, ID: packetid.WindowItems, F: m.onSetContentPacket},
	)
	return m
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
	// TODO: Create the specified container
	return nil
}

func (m *Manager) onSetContentPacket(p pk.Packet) error {
	var (
		ContainerID pk.UnsignedByte
		Count       pk.Short
		SlotData    []slot
	)
	if err := p.Scan(
		&ContainerID,
		&Count, pk.Ary{
			Len: &Count,
			Ary: &SlotData,
		}); err != nil {
		return Error{err}
	}
	container := m.Screens[byte(ContainerID)]
	for i, v := range SlotData {
		container.SetSlot(i, int32(v.id), byte(v.count), v.nbt)
	}
	return nil
}

type slot struct {
	present pk.Boolean
	id      pk.VarInt
	count   pk.Byte
	nbt     nbt.RawMessage
}

func (s *slot) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		&s.present, pk.Opt{Has: &s.present,
			Field: pk.Tuple{
				&s.id, &s.count, pk.NBT(&s.nbt),
			},
		},
	}.ReadFrom(r)
}

type Container interface {
	SetSlot(i int, id int32, count byte, NBT nbt.RawMessage)
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
