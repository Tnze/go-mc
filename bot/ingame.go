package bot

import (
	"fmt"
	pk "github.com/Tnze/go-mc/net/packet"
	"time"
)

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (c *Client) HandleGame() error {
	var ticker = time.Now()
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

		//handle tickers
		if time.Since(ticker) >= time.Duration(50*c.TPS.TickAverage())*time.Millisecond { // Server synchronization
			if err := c.handleTickers(); err != nil {
				fmt.Println("handle tickers error:", err)
			}
			ticker = time.Now()
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
	if c.Events.generic != nil {
		for _, handler := range *c.Events.generic {
			if err = handler.F(c, p); err != nil {
				return PacketHandlerError{ID: p.ID, Err: err}
			}
		}
	}
	if listeners := c.Events.handlers[p.ID]; listeners != nil {
		for _, handler := range *listeners {
			if err = handler.F(c, p); err != nil {
				return PacketHandlerError{ID: p.ID, Err: err}
			}
		}
	}
	return
}

func (c *Client) handleTickers() error {
	if c.Events.tickers != nil {
		for _, handler := range *c.Events.tickers {
			if err := handler.F(c); err != nil {
				return err
			}
		}
	}
	return nil
}
