package basic

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type EventsListener struct {
	GameStart    func() error
	ChatMsg      func(c *PlayerMessage) error
	SystemMsg    func(c chat.Message, overlay bool) error
	Disconnect   func(reason chat.Message) error
	HealthChange func(health float32) error
	Death        func() error
}

// Attach your event listener to the client.
// The functions are copied when attaching, and modify on [EventListener] doesn't affect after that.
func (e EventsListener) Attach(c *bot.Client) {
	if e.GameStart != nil {
		attachJoinGameHandler(c, e.GameStart)
	}
	if e.ChatMsg != nil {
		attachPlayerMsg(c, e.ChatMsg)
	}
	if e.SystemMsg != nil {
		attachSystemMsg(c, e.SystemMsg)
	}
	if e.Disconnect != nil {
		attachDisconnect(c, e.Disconnect)
	}
	if e.HealthChange != nil || e.Death != nil {
		attachUpdateHealth(c, e.HealthChange, e.Death)
	}
}

func attachJoinGameHandler(c *bot.Client, handler func() error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundLogin,
		F: func(_ pk.Packet) error {
			return handler()
		},
	})
}

type PlayerMessage struct {
}

func (p *PlayerMessage) ReadFrom(r io.Reader) (n int64, err error) {
	// SignedMessageHeader
	// MessageSignature
	// SignedMessageBody
	// Optional<Component>
	// FilterMask
	panic("implement me")
}

type ChatType struct {
	ID         int32
	Name       chat.Message
	TargetName *chat.Message
}

func (c *ChatType) ReadFrom(r io.Reader) (n int64, err error) {
	var hasTargetName pk.Boolean
	n1, err := (*pk.VarInt)(&c.ID).ReadFrom(r)
	if err != nil {
		return n1, err
	}
	n2, err := c.Name.ReadFrom(r)
	if err != nil {
		return n1 + n2, err
	}
	n3, err := hasTargetName.ReadFrom(r)
	if err != nil {
		return n1 + n2 + n3, err
	}
	if hasTargetName {
		c.TargetName = new(chat.Message)
		n4, err := c.TargetName.ReadFrom(r)
		return n1 + n2 + n3 + n4, err
	}
	return n1 + n2 + n3, nil
}

func attachPlayerMsg(c *bot.Client, handler func(c *PlayerMessage) error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundPlayerChat,
		F: func(p pk.Packet) error {
			var message PlayerMessage
			var chatType ChatType
			if err := p.Scan(&message, &chatType); err != nil {
				return Error{err}
			}
			return handler(&message)
		},
	})
}

func attachSystemMsg(c *bot.Client, handler func(c chat.Message, overlay bool) error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundSystemChat,
		F: func(p pk.Packet) error {
			var msg chat.Message
			var pos pk.Boolean
			if err := p.Scan(&msg, &pos); err != nil {
				return Error{err}
			}
			return handler(msg, bool(pos))
		},
	})
}

func attachDisconnect(c *bot.Client, handler func(reason chat.Message) error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundDisconnect,
		F: func(p pk.Packet) error {
			var reason chat.Message
			if err := p.Scan(&reason); err != nil {
				return Error{err}
			}
			return handler(reason)
		},
	})
}

func attachUpdateHealth(c *bot.Client, healthChangeHandler func(health float32) error, deathHandler func() error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundSetHealth,
		F: func(p pk.Packet) error {
			var health pk.Float
			var food pk.VarInt
			var foodSaturation pk.Float

			if err := p.Scan(&health, &food, &foodSaturation); err != nil {
				return Error{err}
			}
			var healthChangeErr, deathErr error
			if healthChangeHandler != nil {
				healthChangeErr = healthChangeHandler(float32(health))
			}
			if deathHandler != nil && health <= 0 {
				healthChangeErr = deathHandler()
			}
			return updateHealthError{healthChangeErr, deathErr}
		},
	})
}

type updateHealthError struct {
	healthChangeErr, deathErr error
}

func (u updateHealthError) Unwrap() error {
	if u.healthChangeErr != nil {
		return u.healthChangeErr
	}
	if u.deathErr != nil {
		return u.deathErr
	}
	return nil
}

func (u updateHealthError) Error() string {
	switch {
	case u.healthChangeErr != nil && u.deathErr != nil:
		return "[" + u.healthChangeErr.Error() + ", " + u.deathErr.Error() + "]"
	case u.healthChangeErr != nil:
		return u.healthChangeErr.Error()
	case u.deathErr != nil:
		return u.deathErr.Error()
	default:
		return "nil"
	}
}
