package bot

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
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
	var (
		SoundID       pk.VarInt
		SoundCategory pk.VarInt
		x, y, z       pk.Int
		Volume, Pitch pk.Float
	)
	err := p.Scan(&SoundID, &SoundCategory, &x, &y, &z, &Volume, &Pitch)
	if err != nil {
		return err
	}

	if c.Events.SoundPlay != nil {
		err = c.Events.SoundPlay(
			data.SoundNames[SoundID], int(SoundCategory),
			float64(x)/8, float64(y)/8, float64(z)/8,
			float32(Volume), float32(Pitch))
	}

	return nil
}

func handleNamedSoundEffect(c *Client, p pk.Packet) error {
	var (
		SoundName     pk.String
		SoundCategory pk.VarInt
		x, y, z       pk.Int
		Volume, Pitch pk.Float
	)
	err := p.Scan(&SoundName, &SoundCategory, &x, &y, &z, &Volume, &Pitch)
	if err != nil {
		return err
	}

	if c.Events.SoundPlay != nil {
		err = c.Events.SoundPlay(
			string(SoundName), int(SoundCategory),
			float64(x)/8, float64(y)/8, float64(z)/8,
			float32(Volume), float32(Pitch))
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
	var (
		windowID pk.Byte
		slotI    pk.Short
		slot     entity.Slot
	)
	if err := p.Scan(&windowID, &slotI, &slot); err != nil && !errors.Is(err, nbt.ErrEND) {
		return err
	}

	return c.Events.WindowsItemChange(byte(windowID), int(slotI), slot)
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
	var (
		s      chat.Message
		pos    pk.Byte
		sender pk.UUID
	)

	err = p.Scan(&s, &pos, &sender)
	if err != nil {
		return err
	}

	if c.Events.ChatMsg != nil {
		err = c.Events.ChatMsg(s, byte(pos), uuid.UUID(sender))
	}

	return err
}

func handleUpdateHealthPacket(c *Client, p pk.Packet) (err error) {
	var (
		Health         pk.Float
		Food           pk.VarInt
		FoodSaturation pk.Float
	)

	err = p.Scan(&Health, &Food, &FoodSaturation)
	if err != nil {
		return
	}

	c.Health = float32(Health)
	c.Food = int32(Food)
	c.FoodSaturation = float32(FoodSaturation)

	if c.Events.HealthChange != nil {
		err = c.Events.HealthChange()
		if err != nil {
			return
		}
	}
	if c.Health < 1 { //player is dead
		sendPlayerPositionAndLookPacket(c)
		if c.Events.Die != nil {
			err = c.Events.Die()
			if err != nil {
				return
			}
		}
	}
	return
}

func handleJoinGamePacket(c *Client, p pk.Packet) error {
	var (
		eid        pk.Int
		hardcore   pk.Boolean
		gamemode   pk.UnsignedByte
		previousGm pk.UnsignedByte
		worldCount pk.VarInt
		worldNames pk.Identifier
		_          pk.NBT
		//dimensionCodec pk.NBT
		dimension    pk.Int
		worldName    pk.Identifier
		hashedSeed   pk.Long
		maxPlayers   pk.VarInt
		viewDistance pk.VarInt
		rdi          pk.Boolean // Reduced Debug Info
		ers          pk.Boolean // Enable respawn screen
		isDebug      pk.Boolean
		isFlat       pk.Boolean
	)
	err := p.Scan(&eid, &hardcore, &gamemode, &previousGm, &worldCount, &worldNames, &dimension, &worldName,
		&hashedSeed, &maxPlayers, &rdi, &ers, &isDebug, &isFlat)
	if err != nil {
		return err
	}

	c.EntityID = int(eid)
	c.Gamemode = int(gamemode & 0x7)
	c.Hardcore = gamemode&0x8 != 0
	c.Dimension = int(dimension)
	c.WorldName = string(worldName)
	c.ViewDistance = int(viewDistance)
	c.ReducedDebugInfo = bool(rdi)
	c.IsDebug = bool(isDebug)
	c.IsFlat = bool(isFlat)

	return nil
}

// The PluginMessageData only used in recive PluginMessage packet.
// When decode it, read to end.
type pluginMessageData []byte

//Encode a PluginMessageData
func (p pluginMessageData) Encode() []byte {
	return []byte(p)
}

//Decode a PluginMessageData
func (p *pluginMessageData) Decode(r pk.DecodeReader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	*p = data
	return nil
}

func handlePluginPacket(c *Client, p pk.Packet) error {
	var (
		Channel pk.Identifier
		Data    pluginMessageData
	)
	if err := p.Scan(&Channel, &Data); err != nil {
		return err
	}

	switch Channel {
	case "minecraft:brand":
		var brandRaw pk.String
		if err := brandRaw.Decode(bytes.NewReader(Data)); err != nil {
			return err
		}
		c.ServInfo.Brand = string(brandRaw)
	}

	if c.Events.PluginMessage != nil {
		return c.Events.PluginMessage(string(Channel), []byte(Data))
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

	var (
		X, Z           pk.Int
		FullChunk      pk.Boolean
		PrimaryBitMask pk.VarInt
		Heightmaps     struct{}
		Biomes         = biomesData{fullChunk: (*bool)(&FullChunk)}
		Data           chunkData
		BlockEntities  blockEntities
	)
	if err := p.Scan(&X, &Z, &FullChunk, &PrimaryBitMask, pk.NBT{V: &Heightmaps}, &Biomes, &Data, &BlockEntities); err != nil {
		return err
	}
	chunk, err := world.DecodeChunkColumn(int32(PrimaryBitMask), Data)
	if err != nil {
		return fmt.Errorf("decode chunk column fail: %w", err)
	}

	c.Wd.LoadChunk(int(X), int(Z), chunk)

	return err
}

type biomesData struct {
	fullChunk *bool
	data      []pk.VarInt
}

func (b *biomesData) Decode(r pk.DecodeReader) error {
	if b.fullChunk == nil || !*b.fullChunk {
		return nil
	}

	var nobd pk.VarInt // Number of Biome Datums
	if err := nobd.Decode(r); err != nil {
		return err
	}

	b.data = make([]pk.VarInt, nobd)

	for i := 0; i < int(nobd); i++ {
		var d pk.VarInt
		if err := d.Decode(r); err != nil {
			return err
		}
		b.data[i] = d
	}

	return nil
}

type chunkData []byte
type blockEntities []blockEntitie
type blockEntitie struct {
}

// Decode implement net.packet.FieldDecoder
func (c *chunkData) Decode(r pk.DecodeReader) error {
	var Size pk.VarInt
	if err := Size.Decode(r); err != nil {
		return err
	}
	*c = make([]byte, Size)
	if _, err := r.Read(*c); err != nil {
		return err
	}
	return nil
}

// Decode implement net.packet.FieldDecoder
func (b *blockEntities) Decode(r pk.DecodeReader) error {
	var nobe pk.VarInt // Number of BlockEntities
	if err := nobe.Decode(r); err != nil {
		return err
	}
	*b = make(blockEntities, nobe)
	decoder := nbt.NewDecoder(r)
	for i := 0; i < int(nobe); i++ {
		if err := decoder.Decode(&(*b)[i]); err != nil {
			return err
		}
	}
	return nil
}

func handlePlayerPositionAndLookPacket(c *Client, p pk.Packet) error {
	var (
		x, y, z    pk.Double
		yaw, pitch pk.Float
		flags      pk.Byte
		TeleportID pk.VarInt
	)

	err := p.Scan(&x, &y, &z, &yaw, &pitch, &flags, &TeleportID)
	if err != nil {
		return err
	}

	if flags&0x01 == 0 {
		c.X = float64(x)
	} else {
		c.X += float64(x)
	}
	if flags&0x02 == 0 {
		c.Y = float64(y)
	} else {
		c.Y += float64(y)
	}
	if flags&0x04 == 0 {
		c.Z = float64(z)
	} else {
		c.Z += float64(z)
	}
	if flags&0x08 == 0 {
		c.Yaw = float32(yaw)
	} else {
		c.Yaw += float32(yaw)
	}
	if flags&0x10 == 0 {
		c.Pitch = float32(pitch)
	} else {
		c.Pitch += float32(pitch)
	}

	//Confirm
	return c.conn.WritePacket(pk.Marshal(
		data.TeleportConfirm,
		pk.VarInt(TeleportID),
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
	r := bytes.NewReader(p.Data)
	var (
		windowID pk.Byte
		count    pk.Short
		slots    []entity.Slot
	)
	if err := windowID.Decode(r); err != nil {
		return err
	}
	if err := count.Decode(r); err != nil {
		return err
	}
	for i := 0; i < int(count); i++ {
		var slot entity.Slot
		if err := slot.Decode(r); err != nil && !errors.Is(err, nbt.ErrEND) {
			return err
		}
		slots = append(slots, slot)
	}

	if windowID == 0 { // Window ID 0 is the players' inventory.
		if err := c.Events.updateSeenPackets(seenPlayerInventory); err != nil {
			return err
		}
	}
	if c.Events.WindowsItem == nil {
		return nil
	}
	return c.Events.WindowsItem(byte(windowID), slots)
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
