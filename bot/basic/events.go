package basic

import (
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type EventsListener struct {
	GameStart    func() error
	ChatMsg      func(c *PlayerMessage) error
	SystemMsg    func(c chat.Message, pos byte) error
	Disconnect   func(reason chat.Message) error
	HealthChange func(health float32) error
	Death        func() error
}

func (e EventsListener) Attach(c *bot.Client) {
	c.Events.AddListener(
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundLogin, F: e.onJoinGame},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundSystemChat, F: e.onSystemMsg},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundPlayerChat, F: e.onPlayerMsg},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundDisconnect, F: e.onDisconnect},
		bot.PacketHandler{Priority: 64, ID: packetid.ClientboundSetHealth, F: e.onUpdateHealth},
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

type PlayerMessage struct {
	SignedMessage     chat.Message
	Unsigned          bool
	UnsignedMessage   chat.Message
	Position          int32
	Sender            uuid.UUID
	SenderDisplayName chat.Message
	HasSenderTeam     bool
	SenderTeamName    chat.Message
	TimeStamp         time.Time
}

func (e *EventsListener) onPlayerMsg(p pk.Packet) error {
	if e.ChatMsg != nil {
		var message PlayerMessage
		var senderDisplayName pk.String
		var senderTeamName pk.String
		var timeStamp pk.Long
		var salt pk.Long
		var signature pk.ByteArray
		if err := p.Scan(&message.SignedMessage,
			(*pk.Boolean)(&message.Unsigned),
			pk.Opt{
				Has:   &message.Unsigned,
				Field: &message.UnsignedMessage,
			},
			(*pk.VarInt)(&message.Position),
			(*pk.UUID)(&message.Sender),
			&senderDisplayName,
			(*pk.Boolean)(&message.HasSenderTeam),
			pk.Opt{
				Has:   &message.HasSenderTeam,
				Field: &senderTeamName,
			},
			&timeStamp,
			&salt,
			&signature); err != nil {
			return Error{err}
		}
		if err := message.SenderDisplayName.UnmarshalJSON([]byte(senderDisplayName)); err != nil {
			return Error{err}
		}
		if err := message.SenderTeamName.UnmarshalJSON([]byte(senderDisplayName)); err != nil {
			return Error{err}
		}
		return e.ChatMsg(&message)
	}
	return nil
}
func (e *EventsListener) onSystemMsg(p pk.Packet) error {
	if e.SystemMsg != nil {
		var msg chat.Message
		var pos pk.VarInt

		if err := p.Scan(&msg, &pos); err != nil {
			return Error{err}
		}
		return e.SystemMsg(msg, byte(pos))
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
