package bot

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/net/ptypes"
)

// //GetPosition return the player's position
// func (p *Player) GetPosition() (x, y, z float64) {
// 	return p.X, p.Y, p.Z
// }

// //GetBlockPos return the position of the Block at player's feet
// func (p *Player) GetBlockPos() (x, y, z int) {
// 	return int(math.Floor(p.X)), int(math.Floor(p.Y)), int(math.Floor(p.Z))
// }

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (c *Client) HandleGame() error {
	for {
		select {
		case task := <-c.Delegate:
			if err := task(); err != nil {
				return err
			}
		default:
			//Read packets
			p, err := c.conn.ReadPacket()
			if err != nil {
				return fmt.Errorf("bot: read packet fail: %w", err)
			}
			//handle packets
			disconnect, err := c.handlePacket(p)
			if err != nil {
				return fmt.Errorf("handle packet 0x%X error: %w", p.ID, err)
			}
			if disconnect {
				return nil
			}
		}
	}
}

func (c *Client) handlePacket(p pk.Packet) (disconnect bool, err error) {
	if c.Events.ReceivePacket != nil {
		pass, err := c.Events.ReceivePacket(p)
		if err != nil {
			return false, err
		}
		if pass {
			return false, nil
		}
	}

	switch data.PktID(p.ID) {
	case data.Login:
		err = handleJoinGamePacket(c, p)

		if err == nil && c.Events.GameStart != nil {
			err = c.Events.GameStart()
		}

		_ = c.conn.WritePacket(
			//PluginMessage packet (serverbound) - sending minecraft brand.
			pk.Marshal(
				data.CustomPayloadServerbound,
				pk.Identifier("minecraft:brand"),
				pk.String(c.settings.Brand),
			),
		)
		if err2 := c.Events.updateSeenPackets(seenJoinGame); err == nil {
			err = err2
		}
	case data.CustomPayloadClientbound:
		err = handlePluginPacket(c, p)
	case data.Difficulty:
		err = handleServerDifficultyPacket(c, p)
		if err == nil && c.Events.ServerDifficultyChange != nil {
			err = c.Events.ServerDifficultyChange(c.Difficulty)
		}
		if err2 := c.Events.updateSeenPackets(seenServerDifficulty); err == nil {
			err = err2
		}
	case data.SpawnPosition:
		err = handleSpawnPositionPacket(c, p)
		if err2 := c.Events.updateSeenPackets(seenSpawnPos); err == nil {
			err = err2
		}
	case data.AbilitiesClientbound:
		err = handlePlayerAbilitiesPacket(c, p)
		_ = c.conn.WritePacket(
			//ClientSettings packet (serverbound)
			pk.Marshal(
				data.Settings,
				pk.String(c.settings.Locale),
				pk.Byte(c.settings.ViewDistance),
				pk.VarInt(c.settings.ChatMode),
				pk.Boolean(c.settings.ChatColors),
				pk.UnsignedByte(c.settings.DisplayedSkinParts),
				pk.VarInt(c.settings.MainHand),
			),
		)
		if err2 := c.Events.updateSeenPackets(seenPlayerAbilities); err == nil {
			err = err2
		}
	case data.HeldItemSlotClientbound:
		err = handleHeldItemPacket(c, p)
	case data.UpdateLight:
		err = c.Events.updateSeenPackets(seenUpdateLight)
	case data.MapChunk:
		err = handleChunkDataPacket(c, p)
		if err2 := c.Events.updateSeenPackets(seenChunkData); err == nil {
			err = err2
		}
	case data.PositionClientbound:
		err = handlePlayerPositionAndLookPacket(c, p)
		sendPlayerPositionAndLookPacket(c) // to confirm the position
		if err2 := c.Events.updateSeenPackets(seenPlayerPositionAndLook); err == nil {
			err = err2
		}
	case data.DeclareRecipes:
		// handleDeclareRecipesPacket(g, reader)
	case data.KeepAliveClientbound:
		err = handleKeepAlivePacket(c, p)
	case data.Entity:
		//handleEntityPacket(g, reader)
	case data.NamedEntitySpawn:
		// err = handleSpawnPlayerPacket(g, reader)
	case data.WindowItems:
		err = handleWindowItemsPacket(c, p)
	case data.UpdateHealth:
		err = handleUpdateHealthPacket(c, p)
	case data.ChatClientbound:
		err = handleChatMessagePacket(c, p)
	case data.BlockChange:
		////err = handleBlockChangePacket(c, p)
	case data.MultiBlockChange:
		////err = handleMultiBlockChangePacket(c, p)
	case data.KickDisconnect:
		err = handleDisconnectPacket(c, p)
		disconnect = true
	case data.SetSlot:
		err = handleSetSlotPacket(c, p)
	case data.SoundEffect:
		err = handleSoundEffect(c, p)
	case data.NamedSoundEffect:
		err = handleNamedSoundEffect(c, p)
	case data.Experience:
		err = handleSetExperience(c, p)
	default:
		// fmt.Printf("ignore pack id %X\n", p.ID)
	}
	return
}

func handleSoundEffect(c *Client, p pk.Packet) error {
	var s ptypes.SoundEffect
	if err := s.Decode(p); err != nil {
		return err
	}

	if c.Events.SoundPlay != nil {
		return c.Events.SoundPlay(
			data.SoundNames[s.Sound], int(s.Category),
			float64(s.X)/8, float64(s.Y)/8, float64(s.Z)/8,
			float32(s.Volume), float32(s.Pitch))
	}

	return nil
}

func handleNamedSoundEffect(c *Client, p pk.Packet) error {
	var s ptypes.NamedSoundEffect
	if err := s.Decode(p); err != nil {
		return err
	}

	if c.Events.SoundPlay != nil {
		return c.Events.SoundPlay(
			string(s.Sound), int(s.Category),
			float64(s.X)/8, float64(s.Y)/8, float64(s.Z)/8,
			float32(s.Volume), float32(s.Pitch))
	}

	return nil
}

func handleDisconnectPacket(c *Client, p pk.Packet) error {
	var reason chat.Message

	err := p.Scan(&reason)
	if err != nil {
		return err
	}

	if c.Events.Disconnect != nil {
		return c.Events.Disconnect(reason)
	}
	return nil
}

func handleSetSlotPacket(c *Client, p pk.Packet) error {
	if c.Events.WindowsItemChange == nil {
		return nil
	}
	var pkt ptypes.SetSlot
	if err := pkt.Decode(p); err != nil {
		return err
	}

	return c.Events.WindowsItemChange(byte(pkt.WindowID), int(pkt.Slot), pkt.SlotData)
}

// func handleMultiBlockChangePacket(c *Client, p pk.Packet) error {
// 	if !c.settings.ReceiveMap {
// 		return nil
// 	}

// 	var cX, cY pk.Int

// 	err := p.Scan(&cX, &cY)
// 	if err != nil {
// 		return err
// 	}

// 	c := g.wd.chunks[chunkLoc{int(cX), int(cY)}]
// 	if c != nil {
// 		RecordCount, err := pk.UnpackVarInt(r)
// 		if err != nil {
// 			return err
// 		}

// 		for i := int32(0); i < RecordCount; i++ {
// 			xz, err := r.ReadByte()
// 			if err != nil {
// 				return err
// 			}
// 			y, err := r.ReadByte()
// 			if err != nil {
// 				return err
// 			}
// 			BlockID, err := pk.UnpackVarInt(r)
// 			if err != nil {
// 				return err
// 			}
// 			x, z := xz>>4, xz&0x0F

// 			c.sections[y/16].blocks[x][y%16][z] = Block{id: uint(BlockID)}
// 		}
// 	}

// 	return nil
// }

// func handleBlockChangePacket(c *Client, p pk.Packet) error {
// 	if !c.settings.ReceiveMap {
// 		return nil
// 	}
// 	var pos pk.Position
// 	err := p.Scan(&pos)
// 	if err != nil {
// 		return err
// 	}

// 	c := g.wd.chunks[chunkLoc{x >> 4, z >> 4}]
// 	if c != nil {
// 		id, err := pk.UnpackVarInt(r)
// 		if err != nil {
// 			return err
// 		}
// 		c.sections[y/16].blocks[x&15][y&15][z&15] = Block{id: uint(id)}
// 	}

// 	return nil
// }

func handleChatMessagePacket(c *Client, p pk.Packet) (err error) {
	var msg ptypes.ChatMessageClientbound
	if err := msg.Decode(p); err != nil {
		return err
	}

	if c.Events.ChatMsg != nil {
		return c.Events.ChatMsg(msg.S, byte(msg.Pos), uuid.UUID(msg.Sender))
	}
	return nil
}

func handleUpdateHealthPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.UpdateHealth
	if err := pkt.Decode(p); err != nil {
		return err
	}

	c.Health = float32(pkt.Health)
	c.Food = int32(pkt.Food)
	c.FoodSaturation = float32(pkt.FoodSaturation)

	if c.Events.HealthChange != nil {
		if err := c.Events.HealthChange(); err != nil {
			return err
		}
	}
	if c.Health < 1 { //player is dead
		sendPlayerPositionAndLookPacket(c)
		if c.Events.Die != nil {

			if err := c.Events.Die(); err != nil {
				return err
			}
		}
	}
	return nil
}

func handleJoinGamePacket(c *Client, p pk.Packet) error {
	var pkt ptypes.JoinGame
	if err := pkt.Decode(p); err != nil {
		return err
	}

	c.EntityID = int(pkt.PlayerEntity)
	c.Gamemode = int(pkt.Gamemode & 0x7)
	c.Hardcore = pkt.Gamemode&0x8 != 0
	c.Dimension = int(pkt.Dimension)
	c.WorldName = string(pkt.WorldName)
	c.ViewDistance = int(pkt.ViewDistance)
	c.ReducedDebugInfo = bool(pkt.RDI)
	c.IsDebug = bool(pkt.IsDebug)
	c.IsFlat = bool(pkt.IsFlat)

	return nil
}

func handlePluginPacket(c *Client, p pk.Packet) error {
	var msg ptypes.PluginMessage
	if err := msg.Decode(p); err != nil {
		return err
	}

	switch msg.Channel {
	case "minecraft:brand":
		var brandRaw pk.String
		if err := brandRaw.Decode(bytes.NewReader(msg.Data)); err != nil {
			return err
		}
		c.ServInfo.Brand = string(brandRaw)
	}

	if c.Events.PluginMessage != nil {
		return c.Events.PluginMessage(string(msg.Channel), []byte(msg.Data))
	}
	return nil
}

func handleServerDifficultyPacket(c *Client, p pk.Packet) error {
	var difficulty pk.Byte
	err := p.Scan(&difficulty)
	if err != nil {
		return err
	}
	c.Difficulty = int(difficulty)
	return nil
}

func handleSpawnPositionPacket(c *Client, p pk.Packet) error {
	var pos pk.Position
	err := p.Scan(&pos)
	if err != nil {
		return err
	}
	c.SpawnPosition.X, c.SpawnPosition.Y, c.SpawnPosition.Z =
		pos.X, pos.Y, pos.Z
	return nil
}

func handlePlayerAbilitiesPacket(g *Client, p pk.Packet) error {
	var (
		flags    pk.Byte
		flySpeed pk.Float
		viewMod  pk.Float
	)
	err := p.Scan(&flags, &flySpeed, &viewMod)
	if err != nil {
		return err
	}
	g.abilities.Flags = int8(flags)
	g.abilities.FlyingSpeed = float32(flySpeed)
	g.abilities.FieldofViewModifier = float32(viewMod)
	return nil
}

func handleHeldItemPacket(c *Client, p pk.Packet) error {
	var hi pk.Byte
	if err := p.Scan(&hi); err != nil {
		return err
	}
	c.HeldItem = int(hi)

	if c.Events.HeldItemChange != nil {
		return c.Events.HeldItemChange(c.HeldItem)
	}
	return nil
}

func handleChunkDataPacket(c *Client, p pk.Packet) error {
	if !c.settings.ReceiveMap {
		return nil
	}

	var pkt ptypes.ChunkData
	if err := pkt.Decode(p); err != nil {
		return err
	}

	chunk, err := world.DecodeChunkColumn(int32(pkt.PrimaryBitMask), pkt.Data)
	if err != nil {
		return fmt.Errorf("decode chunk column: %w", err)
	}

	c.Wd.LoadChunk(int(pkt.X), int(pkt.Z), chunk)
	return nil
}

func handlePlayerPositionAndLookPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.PositionAndLookClientbound
	if err := pkt.Decode(p); err != nil {
		return err
	}

	if pkt.RelativeX() {
		c.X = float64(pkt.X)
	} else {
		c.X += float64(pkt.X)
	}
	if pkt.RelativeY() {
		c.Y = float64(pkt.Y)
	} else {
		c.Y += float64(pkt.Y)
	}
	if pkt.RelativeZ() {
		c.Z = float64(pkt.Z)
	} else {
		c.Z += float64(pkt.Z)
	}
	if pkt.RelativeYaw() {
		c.Yaw = float32(pkt.Yaw)
	} else {
		c.Yaw += float32(pkt.Yaw)
	}
	if pkt.RelativePitch() {
		c.Pitch = float32(pkt.Pitch)
	} else {
		c.Pitch += float32(pkt.Pitch)
	}

	//Confirm
	return c.conn.WritePacket(pk.Marshal(
		data.TeleportConfirm,
		pk.VarInt(pkt.TeleportID),
	))
}

func handleKeepAlivePacket(c *Client, p pk.Packet) error {
	var KeepAliveID pk.Long
	if err := p.Scan(&KeepAliveID); err != nil {
		return err
	}
	//Response
	return c.conn.WritePacket(pk.Marshal(
		data.KeepAliveServerbound,
		KeepAliveID,
	))
}

func handleWindowItemsPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.WindowItems
	if err := pkt.Decode(p); err != nil {
		return err
	}

	if pkt.WindowID == 0 { // Window ID 0 is the players' inventory.
		if err := c.Events.updateSeenPackets(seenPlayerInventory); err != nil {
			return err
		}
	}
	if c.Events.WindowsItem != nil {
		return c.Events.WindowsItem(byte(pkt.WindowID), pkt.Slots)
	}
	return nil
}

func handleSetExperience(c *Client, p pk.Packet) (err error) {
	var (
		bar   pk.Float
		level pk.VarInt
		total pk.VarInt
	)

	if err := p.Scan(&bar, &level, &total); err != nil {
		return err
	}

	c.Level = int32(level)

	if c.Events.ExperienceChange != nil {
		return c.Events.ExperienceChange(float32(bar), int32(level), int32(total))
	}

	return nil
}

func sendPlayerPositionAndLookPacket(c *Client) {
	c.conn.WritePacket(pk.Marshal(
		data.PositionLook,
		pk.Double(c.X),
		pk.Double(c.Y),
		pk.Double(c.Z),
		pk.Float(c.Yaw),
		pk.Float(c.Pitch),
		pk.Boolean(c.OnGround),
	))
}
