package bot

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"

	pk "github.com/Tnze/go-mc/net/packet"
)

type eventBroker struct {
	GameStart      func() error
	ChatMsg        func(msg chat.Message, pos byte, sender string) error
	Disconnect     func(reason chat.Message) error
	HealthChange   func() error
	Die            func() error
	SoundPlay      func(name string, category int, x, y, z float64, volume, pitch float32) error
	PluginMessage  func(channel string, data []byte) error
	HeldItemChange func(slot int) error

	WindowsItem       func(id byte, slots []entity.Slot) error
	WindowsItemChange func(id byte, slotID int, slot entity.Slot) error

	// ReceivePacket will be called when new packet arrive.
	// Default handler will run only if pass == false.
	ReceivePacket func(p pk.Packet) (pass bool, err error)
}
