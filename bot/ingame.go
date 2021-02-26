package bot

import (
	"fmt"
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
	ID  int32
	Err error
}

func (d PacketHandlerError) Error() string {
	return fmt.Sprintf("handle packet 0x%X error: %v", d.ID, d.Err)
}

func (d PacketHandlerError) Unwrap() error {
	return d.Err
}

func (c *Client) handlePacket(p pk.Packet) (err error) {
	if listeners := c.Events.handlers[p.ID]; listeners != nil {
		for _, handler := range *listeners {
			err = handler.F(p)
			if err != nil {
				return PacketHandlerError{ID: p.ID, Err: err}
			}
		}
	}
	return
}
