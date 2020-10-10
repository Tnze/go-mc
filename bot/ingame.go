package bot

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/net/ptypes"
)

func (c *Client) updateServerPos(pos player.Pos) error {
	prev := c.Player.Pos
	c.Player.Pos = pos

	switch {
	case !prev.LookEqual(pos) && !prev.PosEqual(pos):
		sendPlayerPositionAndLookPacket(c)
	case !prev.PosEqual(pos):
		sendPlayerPositionPacket(c)
	case !prev.LookEqual(pos):
		sendPlayerLookPacket(c)
	case prev.OnGround != pos.OnGround:
		c.conn.WritePacket(
			pk.Marshal(
				data.Flying,
				pk.Boolean(pos.OnGround),
			),
		)
	}

	if c.justTeleported || time.Now().Add(-time.Second).After(c.lastPosTx) {
		c.justTeleported = false
		c.lastPosTx = time.Now()
		sendPlayerPositionPacket(c)
	}

	if c.Events.PositionChange != nil && !prev.Equal(pos) {
		if err := c.Events.PositionChange(pos); err != nil {
			return err
		}
	}
	return nil
}

// HandleGame receive server packet and response them correctly.
// Note that HandleGame will block if you don't receive from Events.
func (c *Client) HandleGame() error {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.closing:
				return

			default:
				//Read packets
				p, err := c.conn.ReadPacket()
				if err != nil {
					if e, ok := err.(*net.OpError); ok && e.Err.Error() != "use of closed network connection" {
						fmt.Fprintf(os.Stderr, "ReadPacket error: %v\n", err)
					}
					return
				}
				c.inbound <- p
			}
		}
	}()

	cTick := time.NewTicker(time.Second / 10 / 2)
	defer cTick.Stop()
	for {
		select {
		case <-c.closing:
			return http.ErrServerClosed
		case <-cTick.C:
			if c.Events.PrePhysics != nil {
				if err := c.Events.PrePhysics(); err != nil {
					return err
				}
			}
			if err := c.Physics.Tick(c.Inputs, &c.Wd); err != nil {
				c.disconnect()
				return err
			}
			c.updateServerPos(c.Physics.Position())

		case task := <-c.Delegate:
			if err := task(); err != nil {
				c.disconnect()
				return err
			}
		case p := <-c.inbound:
			//handle packets
			disconnect, err := c.handlePacket(p)
			if err != nil {
				c.disconnect()
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
	case data.CustomPayloadClientbound:
		err = handlePluginPacket(c, p)
	case data.Difficulty:
		err = handleServerDifficultyPacket(c, p)
	case data.SpawnPosition:
		err = handleSpawnPositionPacket(c, p)
		if err2 := c.Events.updateSeenPackets(seenSpawnPos); err == nil {
			err = err2
		}
	case data.AbilitiesClientbound:
		err = handlePlayerAbilitiesPacket(c, p)
	case data.UpdateHealth:
		err = handleUpdateHealthPacket(c, p)
	case data.ChatClientbound:
		err = handleChatMessagePacket(c, p)

	case data.HeldItemSlotClientbound:
		err = handleHeldItemPacket(c, p)
	case data.WindowItems:
		err = handleWindowItemsPacket(c, p)
	case data.OpenWindow:
		err = handleOpenWindowPacket(c, p)
	case data.TransactionClientbound:
		err = handleWindowConfirmationPacket(c, p)

	case data.DeclareRecipes:
		// handleDeclareRecipesPacket(g, reader)
	case data.KeepAliveClientbound:
		err = handleKeepAlivePacket(c, p)

	case data.SpawnEntity:
		err = handleSpawnEntityPacket(c, p)
	case data.NamedEntitySpawn:
		err = handleSpawnPlayerPacket(c, p)
	case data.SpawnEntityLiving:
		err = handleSpawnLivingEntityPacket(c, p)
	case data.Animation:
		err = handleEntityAnimationPacket(c, p)
	case data.EntityStatus:
		err = handleEntityStatusPacket(c, p)
	case data.EntityDestroy:
		err = handleDestroyEntitiesPacket(c, p)
	case data.RelEntityMove:
		err = handleEntityPositionPacket(c, p)
	case data.EntityMoveLook:
		err = handleEntityPositionLookPacket(c, p)
	case data.EntityLook:
		err = handleEntityLookPacket(c, p)
	case data.Entity:
		err = handleEntityMovePacket(c, p)

	case data.UpdateLight:
		err = c.Events.updateSeenPackets(seenUpdateLight)
	case data.MapChunk:
		err = handleChunkDataPacket(c, p)
	case data.BlockChange:
		err = handleBlockChangePacket(c, p)
	case data.MultiBlockChange:
		err = handleMultiBlockChangePacket(c, p)
	case data.UnloadChunk:
		err = handleUnloadChunkPacket(c, p)
	case data.TileEntityData:
		err = handleTileEntityDataPacket(c, p)

	case data.PositionClientbound:
		err = handlePlayerPositionAndLookPacket(c, p)
		sendPlayerPositionAndLookPacket(c) // to confirm the position
		if err2 := c.Events.updateSeenPackets(seenPlayerPositionAndLook); err == nil {
			err = err2
		}

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

func handleSpawnEntityPacket(c *Client, p pk.Packet) error {
	var se ptypes.SpawnEntity
	if err := se.Decode(p); err != nil {
		return err
	}
	return c.Wd.OnSpawnEntity(se)
}

func handleSpawnLivingEntityPacket(c *Client, p pk.Packet) error {
	var se ptypes.SpawnLivingEntity
	if err := se.Decode(p); err != nil {
		return err
	}
	return c.Wd.OnSpawnLivingEntity(se)
}

func handleSpawnPlayerPacket(c *Client, p pk.Packet) error {
	var se ptypes.SpawnPlayer
	if err := se.Decode(p); err != nil {
		return err
	}
	fmt.Println(se)
	return c.Wd.OnSpawnPlayer(se)
}

func handleEntityPositionPacket(c *Client, p pk.Packet) error {
	var pu ptypes.EntityPosition
	if err := pu.Decode(p); err != nil {
		return err
	}
	return c.Wd.OnEntityPosUpdate(pu)
}

func handleEntityPositionLookPacket(c *Client, p pk.Packet) error {
	var epr ptypes.EntityPositionLook
	if err := epr.Decode(p); err != nil {
		return err
	}
	return c.Wd.OnEntityPosLookUpdate(epr)
}

func handleEntityLookPacket(c *Client, p pk.Packet) error {
	var er ptypes.EntityRotation
	if err := er.Decode(p); err != nil {
		return err
	}
	return c.Wd.OnEntityLookUpdate(er)
}

func handleEntityMovePacket(c *Client, p pk.Packet) error {
	var id pk.VarInt
	if err := p.Scan(&id); err != nil {
		return err
	}
	fmt.Printf("EntityMove (probs didnt for players): %+v\n", id)
	return nil
}

func handleEntityAnimationPacket(c *Client, p pk.Packet) error {
	var se ptypes.EntityAnimationClientbound
	if err := se.Decode(p); err != nil {
		return err
	}
	// fmt.Printf("EntityAnimationClientbound: %+v\n", se)
	return nil
}

func handleEntityStatusPacket(c *Client, p pk.Packet) error {
	var (
		id     pk.Int
		status pk.Byte
	)
	if err := p.Scan(&id, &status); err != nil {
		return err
	}
	// fmt.Printf("EntityStatus: %v, %v\n", id, status)
	return nil
}

func handleDestroyEntitiesPacket(c *Client, p pk.Packet) error {
	var (
		count pk.VarInt
		r     = bytes.NewReader(p.Data)
	)
	if err := count.Decode(r); err != nil {
		return err
	}

	entities := make([]pk.VarInt, int(count))
	for i := 0; i < int(count); i++ {
		if err := entities[i].Decode(r); err != nil {
			return err
		}
	}

	return c.Wd.OnEntityDestroy(entities)
}

func handleSoundEffect(c *Client, p pk.Packet) error {
	var s ptypes.SoundEffect
	if err := s.Decode(p); err != nil {
		return err
	}

	if c.Events.SoundPlay != nil {
		if int(s.Sound) < len(data.SoundNames) {
			return c.Events.SoundPlay(
				data.SoundNames[s.Sound], int(s.Category),
				float64(s.X)/8, float64(s.Y)/8, float64(s.Z)/8,
				float32(s.Volume), float32(s.Pitch))
		} else {
			fmt.Fprintf(os.Stderr, "WARN: Unknown sound name. is data.SoundNames out of date?\n")
		}
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

func handleMultiBlockChangePacket(c *Client, p pk.Packet) error {
	if !c.settings.ReceiveMap {
		return nil
	}
	r := bytes.NewReader(p.Data)

	var (
		loc            pk.Long
		dontTrustEdges pk.Boolean
		sz             pk.VarInt
	)

	if err := loc.Decode(r); err != nil {
		return fmt.Errorf("packed location: %v", err)
	}
	if err := dontTrustEdges.Decode(r); err != nil {
		return fmt.Errorf("unknown 1: %v", err)
	}
	if err := sz.Decode(r); err != nil {
		return fmt.Errorf("array size: %v", err)
	}

	packedBlocks := make([]pk.VarLong, int(sz))
	for i := 0; i < int(sz); i++ {
		if err := packedBlocks[i].Decode(r); err != nil {
			return fmt.Errorf("block[%d]: %v", i, err)
		}
	}

	x, z, y := int((loc>>42)&((1<<22)-1)),
		int((loc>>20)&((1<<22)-1)),
		int(loc&((1<<20)-1))

	// Apply transform into negative (these numbers are signed)
	if x >= 1<<21 {
		x -= 1 << 22
	}
	if z >= 1<<21 {
		z -= 1 << 22
	}

	c.Wd.MultiBlockUpdate(world.ChunkLoc{X: x, Z: z}, y, packedBlocks)
	return nil
}

func handleBlockChangePacket(c *Client, p pk.Packet) error {
	if !c.settings.ReceiveMap {
		return nil
	}
	var (
		pos pk.Position
		bID pk.VarInt
	)

	if err := p.Scan(&pos, &bID); err != nil {
		return err
	}

	c.Wd.UnaryBlockUpdate(pos, world.BlockStatus(bID))
	return nil
}

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
		c.Physics.Run = false
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

	c.Player.ID = int32(pkt.PlayerEntity)
	c.Gamemode = int(pkt.Gamemode & 0x7)
	c.Hardcore = pkt.Gamemode&0x8 != 0
	c.Dimension = int(pkt.Dimension)
	c.WorldName = string(pkt.WorldName)
	c.ViewDistance = int(pkt.ViewDistance)
	c.ReducedDebugInfo = bool(pkt.RDI)
	c.IsDebug = bool(pkt.IsDebug)
	c.IsFlat = bool(pkt.IsFlat)

	if c.Events.GameStart != nil {
		if err := c.Events.GameStart(); err != nil {
			return err
		}
	}

	c.conn.WritePacket(
		//PluginMessage packet (serverbound) - sending minecraft brand.
		pk.Marshal(
			data.CustomPayloadServerbound,
			pk.Identifier("minecraft:brand"),
			pk.String(c.settings.Brand),
		),
	)
	if err := c.Events.updateSeenPackets(seenJoinGame); err != nil {
		return err
	}
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
	if err := p.Scan(&difficulty); err != nil {
		return err
	}
	c.Difficulty = int(difficulty)

	if c.Events.ServerDifficultyChange != nil {
		if err := c.Events.ServerDifficultyChange(c.Difficulty); err != nil {
			return err
		}
	}
	return c.Events.updateSeenPackets(seenServerDifficulty)
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

func handlePlayerAbilitiesPacket(c *Client, p pk.Packet) error {
	var (
		flags    pk.Byte
		flySpeed pk.Float
		viewMod  pk.Float
	)
	err := p.Scan(&flags, &flySpeed, &viewMod)
	if err != nil {
		return err
	}
	c.abilities.Flags = int8(flags)
	c.abilities.FlyingSpeed = float32(flySpeed)
	c.abilities.FieldofViewModifier = float32(viewMod)

	c.conn.WritePacket(
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
	return c.Events.updateSeenPackets(seenPlayerAbilities)
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

func handleUnloadChunkPacket(c *Client, p pk.Packet) error {
	if !c.settings.ReceiveMap {
		return nil
	}

	var x, z pk.Int
	if err := p.Scan(&x, &z); err != nil {
		return err
	}
	c.Wd.UnloadChunk(world.ChunkLoc{X: int(x) >> 4, Z: int(z) >> 4})
	return nil
}

func handleChunkDataPacket(c *Client, p pk.Packet) error {
	if err := c.Events.updateSeenPackets(seenChunkData); err != nil {
		return err
	}
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
	chunk.TileEntities = make(map[world.TilePosition]entity.BlockEntity, 64)
	for _, e := range pkt.BlockEntities {
		chunk.TileEntities[world.ToTilePos(e.X, e.Y, e.Z)] = e
	}

	c.Wd.LoadChunk(int(pkt.X), int(pkt.Z), chunk)
	return nil
}

func handleTileEntityDataPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.TileEntityData
	if err := pkt.Decode(p); err != nil {
		return err
	}
	return c.Wd.TileEntityUpdate(pkt)
}

func handlePlayerPositionAndLookPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.PositionAndLookClientbound
	if err := pkt.Decode(p); err != nil {
		return err
	}

	pp := c.Player.Pos
	if pkt.RelativeX() {
		pp.X += float64(pkt.X)
	} else {
		pp.X = float64(pkt.X)
	}
	if pkt.RelativeY() {
		pp.Y += float64(pkt.Y)
	} else {
		pp.Y = float64(pkt.Y)
	}
	if pkt.RelativeZ() {
		pp.Z += float64(pkt.Z)
	} else {
		pp.Z = float64(pkt.Z)
	}
	if pkt.RelativeYaw() {
		pp.Yaw += float32(pkt.Yaw)
	} else {
		pp.Yaw = float32(pkt.Yaw)
	}
	if pkt.RelativePitch() {
		pp.Pitch += float32(pkt.Pitch)
	} else {
		pp.Pitch = float32(pkt.Pitch)
	}
	if err := c.Physics.ServerPositionUpdate(pp, &c.Wd); err != nil {
		return err
	}
	c.Player.Pos = pp
	c.justTeleported = true

	if c.Events.PositionChange != nil {
		if err := c.Events.PositionChange(pp); err != nil {
			return err
		}
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

func handleOpenWindowPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.OpenWindow
	if err := pkt.Decode(p); err != nil {
		return err
	}

	if c.Events.OpenWindow != nil {
		return c.Events.OpenWindow(pkt)
	}
	return nil
}

func handleWindowConfirmationPacket(c *Client, p pk.Packet) error {
	var pkt ptypes.ConfirmTransaction
	if err := pkt.Decode(p); err != nil {
		return err
	}

	if c.Events.WindowConfirmation != nil {
		return c.Events.WindowConfirmation(pkt)
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

func sendPlayerPositionAndLookPacket(c *Client) error {
	return c.conn.WritePacket(ptypes.PositionAndLookServerbound{
		X:        pk.Double(c.Pos.X),
		Y:        pk.Double(c.Pos.Y),
		Z:        pk.Double(c.Pos.Z),
		Yaw:      pk.Float(c.Pos.Yaw),
		Pitch:    pk.Float(c.Pos.Pitch),
		OnGround: pk.Boolean(c.Pos.OnGround),
	}.Encode())
}

func sendPlayerPositionPacket(c *Client) error {
	return c.conn.WritePacket(ptypes.Position{
		X:        pk.Double(c.Pos.X),
		Y:        pk.Double(c.Pos.Y),
		Z:        pk.Double(c.Pos.Z),
		OnGround: pk.Boolean(c.Pos.OnGround),
	}.Encode())
}

func sendPlayerLookPacket(c *Client) error {
	return c.conn.WritePacket(ptypes.Look{
		Yaw:      pk.Float(c.Pos.Yaw),
		Pitch:    pk.Float(c.Pos.Pitch),
		OnGround: pk.Boolean(c.Pos.OnGround),
	}.Encode())
}
