package bot

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"

	pk "github.com/Tnze/go-mc/net/packet"
)

// These events usually happen when the server sends us a relevant packet

type eventBroker struct {
	OnGameBegin      func() error // Happens when the player connects to a server and enters the world
	OnChatMessage    func(message chat.Message, pos byte) error
	OnDisconnect     func(reason chat.Message) error
	OnHPChange       func() error // When the player health changes
	OnDeath          func() error
	OnRespawn        func() error 
	OnSound          func(name string, category int, x, y, z float64, volume, pitch float32) error
	OnPluginMessage  func(channel string, data []byte) error
	OnHeldItemChange func(slot int) error

	OnWindowsItem       func(id byte, slots []entity.Slot) error
	OnWindowsItemChange func(id byte, slotID int, slot entity.Slot) error

	// Called whenever a new packet reaches us, default handler only runs when pass == false
	OnPacketRecieve func(p pk.Packet) (pass bool, err error)
}
