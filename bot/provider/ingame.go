package provider

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
			fmt.Printf("read packet error %x: %v", p.ID, err)
			return err
		}

		//handle packets
		if err := c.handlePacket(p); err != nil {
			fmt.Println("handle packet error:", err)
			return err
		}
	}
}

func (c *Client) handlePacket(p pk.Packet) (err error) {
	if c.Events.generic != nil {
		for _, handler := range *c.Events.generic {
			err = handler.F(c, p)
			if err != nil {
				return
			}
		}
	}
	if listeners := c.Events.handlers[p.ID]; listeners != nil {
		for _, handler := range *listeners {
			err = handler.F(c, p)
			if err != nil {
				return
			}
		}
	}
	return
}

func (c *Client) handleTickers() (err error) {
	if c.Events.tickers != nil {
		for _, handler := range *c.Events.tickers {
			err = handler.F(c)
			if err != nil {
				return
			}
		}
	}
	return
}
