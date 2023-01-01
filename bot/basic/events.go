package basic

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type EventsListener struct {
	GameStart    func() error
	Disconnect   func(reason chat.Message) error
	HealthChange func(health float32) error
	Death        func() error
}

// attach your event listener to the client.
// The functions are copied when attaching, and modify on [EventListener] doesn't affect after that.
func (e EventsListener) attach(p *Player) {
	if e.GameStart != nil {
		attachJoinGameHandler(p.c, e.GameStart)
	}
	if e.Disconnect != nil {
		attachDisconnect(p.c, e.Disconnect)
	}
	if e.HealthChange != nil || e.Death != nil {
		attachUpdateHealth(p.c, e.HealthChange, e.Death)
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
				deathErr = deathHandler()
			}
			if healthChangeErr != nil || deathErr != nil {
				return updateHealthError{healthChangeErr, deathErr}
			}
			return nil
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
