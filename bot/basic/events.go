package basic

import (
	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type EventsListener struct {
	GameStart    func() error
	ChatMsg      func(c chat.Message, pos byte, uuid uuid.UUID) error
	Disconnect   func(reason chat.Message) error
	HealthChange func(health float32) error
	Death        func() error
}

func (e EventsListener) Attach(c *bot.Client) {
	c.Events.AddListener(
		bot.PacketHandler{Priority: 64, ID: packetid.Login, F: e.onJoinGame},
		bot.PacketHandler{Priority: 64, ID: packetid.ChatClientbound, F: e.onChatMsg},
		bot.PacketHandler{Priority: 64, ID: packetid.KickDisconnect, F: e.onDisconnect},
		bot.PacketHandler{Priority: 64, ID: packetid.UpdateHealth, F: e.onUpdateHealth},
	)
}

func (e *EventsListener) onJoinGame(_ pk.Packet) error {
	if e.GameStart != nil {
		return e.GameStart()
	}
	return nil
}

func (e *EventsListener) onDisconnect(p pk.Packet) error {
	if e.Disconnect != nil {
		var reason chat.Message
		if err := p.Scan(&reason); err != nil {
			return Error{err}
		}
		return e.Disconnect(reason)
	}
	return nil
}

func (e *EventsListener) onChatMsg(p pk.Packet) error {
	if e.ChatMsg != nil {
		var msg chat.Message
		var pos pk.Byte
		var sender pk.UUID

		if err := p.Scan(&msg, &pos, &sender); err != nil {
			return Error{err}
		}

		return e.ChatMsg(msg, byte(pos), uuid.UUID(sender))
	}
	return nil
}

func (e *EventsListener) onUpdateHealth(p pk.Packet) error {
	if e.ChatMsg != nil {
		var health pk.Float
		var food pk.VarInt
		var foodSaturation pk.Float

		if err := p.Scan(&health, &food, &foodSaturation); err != nil {
			return Error{err}
		}
		if e.HealthChange != nil {
			if err := e.HealthChange(float32(health)); err != nil {
				return err
			}
		}
		if e.Death != nil && health <= 0 {
			if err := e.Death(); err != nil {
				return err
			}
		}
	}
	return nil
}
