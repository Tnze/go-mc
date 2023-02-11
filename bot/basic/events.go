package basic

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// EventsListener is a collection of event handlers.
// Fill the fields with your handler functions and pass it to [NewPlayer] to create the Player manager.
// For the event you don't want to handle, just leave it nil.
type EventsListener struct {
	// GameStart event is called when the login process is completed and the player is ready to play.
	//
	// If you want to do some action when the bot joined the server like sending a chat message,
	// this event is the right place to do it.
	GameStart func() error

	// Disconnect event is called before the server disconnects your client.
	// When the server willfully disconnects the client, it will send a ClientboundDisconnect packet and tell you why.
	// On vanilla client, the reason is displayed in the disconnect screen.
	//
	// This information may be very useful for debugging, and generally you should record it into the log.
	//
	// If the connection is disconnected due to network reasons or the client's initiative,
	// this event will not be triggered.
	Disconnect func(reason chat.Message) error

	// HealthChange event is called when the player's health or food changed.
	HealthChange func(health float32, foodLevel int32, foodSaturation float32) error

	// Death event is a special case of HealthChange.
	// It will be called after HealthChange handler called (if it isn't nil)
	// when the player's health is less than or equal to 0.
	//
	// Typically, you should call [Player.Respawn] in this handler.
	Death func() error

	// Teleported event is called when the server think the player position in the client side is wrong,
	// and send a ClientboundPlayerPosition packet to correct the client.
	//
	// Typically, you need to do two things in this handler:
	// - Update the player's position and rotation you tracked to the correct position.
	// - Call [Player.AcceptTeleportation] to send a teleport confirmation packet to the server.
	//
	// Before you confirm the teleportation, the server will not accept any player motion packets.
	//
	// The position coordinates and rotation are absolute or relative depends on the flags.
	// The flag byte is a bitfield, specifies whether each coordinate value is absolute or relative.
	// For more information, see https://wiki.vg/Protocol#Synchronize_Player_Position
	Teleported func(x, y, z float64, yaw, pitch float32, flags byte, teleportID int32, dismountVehicle bool) error
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
	if e.Teleported != nil {
		attachPlayerPosition(p.c, e.Teleported)
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

func attachUpdateHealth(c *bot.Client, healthChangeHandler func(health float32, food int32, saturation float32) error, deathHandler func() error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundSetHealth,
		F: func(p pk.Packet) error {
			var health pk.Float
			var food pk.VarInt
			var saturation pk.Float

			if err := p.Scan(&health, &food, &saturation); err != nil {
				return Error{err}
			}
			var healthChangeErr, deathErr error
			if healthChangeHandler != nil {
				healthChangeErr = healthChangeHandler(float32(health), int32(food), float32(saturation))
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

func attachPlayerPosition(c *bot.Client, handler func(x, y, z float64, yaw, pitch float32, flag byte, teleportID int32, dismountVehicle bool) error) {
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundPlayerPosition,
		F: func(p pk.Packet) error {
			var (
				X, Y, Z         pk.Double
				Yaw, Pitch      pk.Float
				Flags           pk.Byte
				TeleportID      pk.VarInt
				DismountVehicle pk.Boolean
			)
			if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch, &Flags, &TeleportID, &DismountVehicle); err != nil {
				return Error{err}
			}
			return handler(float64(X), float64(Y), float64(Z), float32(Yaw), float32(Pitch), byte(Flags), int32(TeleportID), bool(DismountVehicle))
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
