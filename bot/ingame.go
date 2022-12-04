package bot

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	pk "github.com/Tnze/go-mc/net/packet"
)

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (c *Client) HandleGame() basic.Error {
	var p pk.Packet
	for {
		//Read packets
		if err := c.Conn.ReadPacket(&p); !err.Is(basic.NoError) {
			fmt.Printf("read packet error %x: %v", p.ID, err)
			return err
		}

		//handle packets
		if err := c.handlePacket(p); !err.Is(basic.NoError) {
			fmt.Println("handle packet error:", err.Err)
			return err
		}
	}
}

func (c *Client) handlePacket(p pk.Packet) (err basic.Error) {
	if c.Events.generic != nil {
		for _, handler := range *c.Events.generic {
			if err := handler.F(c, p); !err.Is(basic.NoError) {
				return err
			}
		}
	}
	if listeners := c.Events.handlers[p.ID]; listeners != nil {
		for _, handler := range *listeners {
			if err := handler.F(c, p); !err.Is(basic.NoError) {
				return err
			}
		}
	}
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (c *Client) handleTickers() basic.Error {
	if c.Events.tickers != nil {
		for _, handler := range *c.Events.tickers {
			if err := handler.F(c); !err.Is(basic.NoError) {
				return err
			}
		}
	}
	return basic.Error{Err: basic.NoError, Info: nil}
}
