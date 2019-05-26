package bot

import (
	"errors"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

//SwingArm swing player's arm.
//hand could be one of 0: main hand, 1: off hand.
//It's just animation.
func (c *Client) SwingArm(hand int) error {
	return c.conn.WritePacket(pk.Marshal(
		data.AnimationServerbound,
		pk.VarInt(hand),
	))
}

//Respawn the player when it was dead.
func (c *Client) Respawn() error {
	return c.conn.WritePacket(pk.Marshal(
		data.ClientStatus,
		pk.VarInt(0),
	))
}

//UseItem use the item player handing.
//hand could be one of 0: main hand, 1: off hand
func (c *Client) UseItem(hand int) error {
	return c.conn.WritePacket(pk.Marshal(
		data.UseItem,
		pk.VarInt(hand),
	))
}

//Chat send chat as chat message or command at textbox.
func (c *Client) Chat(msg string) error {
	if len(msg) > 256 {
		return errors.New("message too long")
	}

	return c.conn.WritePacket(pk.Marshal(
		data.ChatMessageServerbound,
		pk.String(msg),
	))
}
