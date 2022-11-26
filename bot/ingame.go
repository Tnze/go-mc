package bot

import (
	"fmt"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (c *Client) HandleGame() error {
	var p pk.Packet
	for {
		//Read packets
		if err := c.Conn.ReadPacket(&p); err != nil {
			return err
		}

		//handle packets
		err := c.handlePacket(p)
		if err != nil {
			return err
		}
	}
}

type PacketHandlerError struct {
	ID  packetid.ClientboundPacketID
	Err error
}

func (d PacketHandlerError) Error() string {
	return fmt.Sprintf("handle packet %v error: %v", d.ID, d.Err)
}

func (d PacketHandlerError) Unwrap() error {
	return d.Err
}

func (c *Client) handlePacket(p pk.Packet) (err error) {
	packetID := packetid.ClientboundPacketID(p.ID)
	if c.Events.generic != nil {
		for _, handler := range *c.Events.generic {
			if err = handler.F(p); err != nil {
				return PacketHandlerError{ID: packetID, Err: err}
			}
		}
	}
	if listeners := c.Events.handlers[packetID]; listeners != nil {
		for _, handler := range *listeners {
			err = handler.F(p)
			if err != nil {
				return PacketHandlerError{ID: packetID, Err: err}
			}
		}
	}
	return
}
