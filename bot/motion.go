package bot

import (
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

//SwingArm swing player's arm.
//hand could be 0: main hand, 1: off hand
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

func (c *Client) UseItem(hand int) error {
	return c.conn.WritePacket(pk.Marshal(
		data.UseItem,
		pk.VarInt(hand),
	))
}
