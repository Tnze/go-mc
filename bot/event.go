package bot

import (
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"

	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/net/ptypes"
)

type seenPacketFlags uint8

// Valid seenPacketFlags values.
const (
	seenJoinGame seenPacketFlags = 1 << iota
	seenServerDifficulty
	seenPlayerAbilities
	seenPlayerInventory
	seenUpdateLight
	seenChunkData
	seenPlayerPositionAndLook
	seenSpawnPos

	// gameReadyMinPackets are the minimum set of packets that must be seen, for
	// the GameReady callback to be invoked.
	gameReadyMinPackets = seenJoinGame | seenChunkData | seenUpdateLight |
		seenPlayerAbilities | seenPlayerInventory | seenServerDifficulty |
		seenPlayerPositionAndLook | seenSpawnPos
)

type eventBroker struct {
	seenPackets seenPacketFlags
	isReady     bool

	GameStart      func() error
	ChatMsg        func(msg chat.Message, pos byte, sender uuid.UUID) error
	Disconnect     func(reason chat.Message) error
	HealthChange   func() error
	Die            func() error
	SoundPlay      func(name string, category int, x, y, z float64, volume, pitch float32) error
	PluginMessage  func(channel string, data []byte) error
	HeldItemChange func(slot int) error
	OpenWindow     func(pkt ptypes.OpenWindow) error

	// ExperienceChange will be called every time player's experience level updates.
	// Parameters:
	//   bar - state of the experience bar from 0.0 to 1.0;
	//   level - current level;
	//   total - total amount of experience received from level 0.
	ExperienceChange func(bar float32, level int32, total int32) error

	WindowsItem        func(id byte, slots []entity.Slot) error
	WindowsItemChange  func(id byte, slotID int, slot entity.Slot) error
	WindowConfirmation func(pkt ptypes.ConfirmTransaction) error

	// ServerDifficultyChange is called whenever the gamemode of the server changes.
	// At time of writing (1.16.3), difficulty values of 0, 1, 2, and 3 correspond
	// to peaceful, easy, normal, and hard respectively.
	ServerDifficultyChange func(difficulty int) error

	// GameReady is called after the client has joined the server and successfully
	// received player state. Additionally, the server has begun sending world
	// state (such as lighting and chunk information).
	//
	// Use this callback as a signal as to when your bot should start 'doing'
	// things.
	GameReady func() error

	// PositionChange is called whenever the player position is updated.
	PositionChange func(pos player.Pos) error

	// ReceivePacket will be called when new packets arrive.
	// The default handler will run only if pass == false.
	ReceivePacket func(p pk.Packet) (pass bool, err error)

	// PrePhysics will be called before a phyiscs tick.
	PrePhysics func() error
}

func (b *eventBroker) updateSeenPackets(f seenPacketFlags) error {
	b.seenPackets |= f
	if (^b.seenPackets)&gameReadyMinPackets == 0 && b.GameReady != nil && !b.isReady {
		b.isReady = true
		return b.GameReady()
	}
	return nil
}
