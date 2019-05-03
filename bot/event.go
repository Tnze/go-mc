package bot

import (
	"github.com/Tnze/go-mc/chat"
)

type eventBroker struct {
	GameStart    func() error
	ChatMsg      func(msg chat.Message) error
	Disconnect   func(reason chat.Message) error
	HealhtChange func() error
	Die          func() error
}
