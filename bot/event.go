package bot

import (
	"github.com/Tnze/go-mc/chat"
)

type eventBroker struct {
	GameStart     func() error
	ChatMsg       func(msg chat.Message, pos byte) error
	Disconnect    func(reason chat.Message) error
	HealthChange  func() error
	Die           func() error
	SoundPlay     func(name string, category int, x, y, z float64, volume, pitch float32) error
	PluginMessage func(channel string, data []byte) error
}
