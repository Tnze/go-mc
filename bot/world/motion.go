package gomcbot

// import (
// 	"math"
// 	"time"
// )

// // SetPosition method move your character around.
// // Server will ignore this if changes too much.
// func (g *Client) SetPosition(x, y, z float64, onGround bool) {
// 	g.motion <- func() {
// 		g.player.X, g.player.Y, g.player.Z = x, y, z
// 		g.player.OnGround = onGround
// 		sendPlayerPositionPacket(g) //向服务器更新位置
// 	}
// }

// // LookAt method turn player's hand and make it look at a point.
// func (g *Client) LookAt(x, y, z float64) {
// 	x0, y0, z0 := g.player.X, g.player.Y, g.player.Z
// 	x, y, z = x-x0, y-y0, z-z0

// 	r := math.Sqrt(x*x + y*y + z*z)
// 	yaw := -math.Atan2(x, z) / math.Pi * 180
// 	for yaw < 0 {
// 		yaw = 360 + yaw
// 	}
// 	pitch := -math.Asin(y/r) / math.Pi * 180

// 	g.LookYawPitch(float32(yaw), float32(pitch))
// }

// // LookYawPitch set player's hand to the direct by yaw and pitch.
// // yaw can be [0, 360) and pitch can be (-180, 180).
// // if |pitch|>90 the player's hand will be very strange.
// func (g *Client) LookYawPitch(yaw, pitch float32) {
// 	g.motion <- func() {
// 		g.player.Yaw, g.player.Pitch = yaw, pitch
// 		sendPlayerLookPacket(g) //向服务器更新朝向
// 	}
// }

// // SwingHand sent when the player's arm swings.
// // if hand is true, swing the main hand
// func (g *Client) SwingHand(hand bool) {
// 	if hand {
// 		sendAnimationPacket(g, 0)
// 	} else {
// 		sendAnimationPacket(g, 1)
// 	}
// }

// // Dig a block in the position and wait for it's breaked
// func (g *Client) Dig(x, y, z int) error {
// 	b := g.GetBlock(x, y, z).id
// 	sendPlayerDiggingPacket(g, 0, x, y, z, Top) //start
// 	sendPlayerDiggingPacket(g, 2, x, y, z, Top) //end

// 	for {
// 		time.Sleep(time.Millisecond * 50)
// 		if g.GetBlock(x, y, z).id != b {
// 			break
// 		}
// 		g.SwingHand(true)
// 	}

// 	return nil
// }

// // UseItem use the item in hand.
// // if hand is true, swing the main hand
// func (g *Client) UseItem(hand bool) {
// 	if hand {
// 		sendUseItemPacket(g, 0)
// 	} else {
// 		sendUseItemPacket(g, 1)
// 	}
// }
