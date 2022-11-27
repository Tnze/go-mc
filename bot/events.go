package bot

import (
	"errors"
	"fmt"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/item"
	"github.com/Tnze/go-mc/level"
	"time"
	"unsafe"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	. "github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

type EventsListener struct{}

func (e EventsListener) Attach(c *Client) {
	c.Events.AddListener(
		PacketHandler{Priority: 64, ID: packetid.CPacketLogin, F: e.JoinGame},
		PacketHandler{Priority: 64, ID: packetid.CPacketKeepAlive, F: e.KeepAlive},
		PacketHandler{Priority: 64, ID: packetid.CPacketChatMessage, F: e.ChatMessage},
		PacketHandler{Priority: 64, ID: packetid.CPacketDisconnect, F: e.Disconnect},
		PacketHandler{Priority: 64, ID: packetid.CPacketUpdateHealth, F: e.UpdateHealth},
		PacketHandler{Priority: int(^uint(0) >> 1), ID: packetid.CPacketSetTime, F: e.TimeUpdate},
	)
}

type PlayerMessage struct {
	SignedMessage     chat.Message
	Unsigned          bool
	UnsignedMessage   chat.Message
	Position          int32
	Sender            uuid.UUID
	SenderDisplayName chat.Message
	HasSenderTeam     bool
	SenderTeamName    chat.Message
	TimeStamp         time.Time
}

func (e *EventsListener) SpawnEntity(c *Client, p pk.Packet) error {
	var (
		EntityID                        pk.VarInt
		EntityUUID                      pk.UUID
		TypeID                          pk.Byte
		X, Y, Z                         pk.Double
		Pitch, Yaw, HeadYaw             pk.Angle
		Data                            pk.VarInt
		VelocityX, VelocityY, VelocityZ pk.Short
	)

	if err := p.Scan(&EntityID, &EntityUUID, &TypeID, &X, &Y, &Z, &Pitch, &Yaw, &HeadYaw, &Data, &VelocityX, &VelocityY, &VelocityZ); err != nil {
		return err
	}

	fmt.Println("SpawnEntity", EntityID, EntityUUID, TypeID, X, Y, Z, Pitch, Yaw, HeadYaw, Data, VelocityX, VelocityY, VelocityZ)
	return nil
}

func (e *EventsListener) SpawnExperienceOrb(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var x, y, z pk.Double
	var count pk.Short

	if err := p.Scan(&entityID, &x, &y, &z, &count); err != nil {
		return err
	}

	fmt.Println("SpawnExperienceOrb", entityID, x, y, z, count)
	return nil
}

func (e *EventsListener) SpawnGlobalEntity(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var typeID pk.Byte
	var x, y, z pk.Double

	if err := p.Scan(&entityID, &typeID, &x, &y, &z); err != nil {
		return err
	}

	fmt.Println("SpawnGlobalEntity", entityID, typeID, x, y, z)
	return nil
}

func (e *EventsListener) SpawnMob(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var u pk.UUID
	var typeID pk.Byte
	var x, y, z pk.Double
	var yaw, pitch, headYaw pk.Angle
	var velocityX, velocityY, velocityZ pk.Short
	//var metadata pk.Metadata // TODO

	if err := p.Scan(&entityID, &u, &typeID, &x, &y, &z, &yaw, &pitch, &headYaw, &velocityX, &velocityY, &velocityZ /*&metadata*/); err != nil {
		return err
	}

	fmt.Println("SpawnMob", entityID, u, typeID, x, y, z, yaw, pitch, headYaw, velocityX, velocityY, velocityZ /*metadata*/)
	return nil
}

func (e *EventsListener) SpawnPainting(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var uuid pk.UUID
	var title pk.String
	var location pk.Position
	var direction pk.Byte

	if err := p.Scan(&entityID, &uuid, &title, &location, &direction); err != nil {
		return err
	}

	fmt.Println("SpawnPainting", entityID, uuid, title, location, direction)
	return nil
}

func (e *EventsListener) SpawnPlayer(c *Client, p pk.Packet) error {
	var (
		EntityID   pk.VarInt
		PlayerUUID pk.UUID
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Angle
	)

	if err := p.Scan(&EntityID, &PlayerUUID, &X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return err
	}
	fmt.Println("SpawnPlayer", EntityID, PlayerUUID.String(), X, Y, Z, Yaw, Pitch)
	return nil
}

func (e *EventsListener) Animation(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var animation pk.Byte

	if err := p.Scan(&entityID, &animation); err != nil {
		return err
	}

	fmt.Println("Animation", entityID, animation)
	return nil
}

func (e *EventsListener) Statistics(c *Client, p pk.Packet) error {
	/*var count pk.VarInt
	var statistics []struct {
		Name  pk.String
		Value pk.VarInt
	}*/

	/*if err := p.Scan(&count, &statistics); err != nil {
		return err
	}

	fmt.Println("Statistics", count, statistics)*/
	return nil
}

func (e *EventsListener) BlockBreakAnimation(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var location pk.Position
	var destroyStage pk.Byte

	if err := p.Scan(&entityID, &location, &destroyStage); err != nil {
		return err
	}

	fmt.Println("BlockBreakAnimation", entityID, location, destroyStage)
	return nil
}

func (e *EventsListener) UpdateBlockEntity(c *Client, p pk.Packet) error {
	/*var location pk.Position
	var action pk.Byte
	var nbtData pk

	if err := p.Scan(&location, &action, &nbtData); err != nil {
		return err
	}

	fmt.Println("UpdateBlockEntity", location, action, nbtData)*/
	return nil
}

func (e *EventsListener) BlockAction(c *Client, p pk.Packet) error {
	var location pk.Position
	var actionID pk.Byte
	var actionParam pk.Byte
	var blockType pk.VarInt

	if err := p.Scan(&location, &actionID, &actionParam, &blockType); err != nil {
		return err
	}

	fmt.Println("BlockAction", location, actionID, actionParam, blockType)
	return nil
}

func (e *EventsListener) BlockChange(c *Client, p pk.Packet) error {
	var location pk.Position
	var blockType pk.VarInt

	if err := p.Scan(&location, &blockType); err != nil {
		return err
	}

	fmt.Println("BlockChange", location, blockType)
	return nil
}

func (e *EventsListener) BossBar(c *Client, p pk.Packet) error {
	var uuid pk.UUID
	var action pk.Byte

	if err := p.Scan(&uuid, &action); err != nil {
		return err
	}

	switch action {
	case 0:
		var title pk.String
		var health pk.Float
		var color pk.Byte
		var division pk.Byte
		var flags pk.Byte

		if err := p.Scan(&title, &health, &color, &division, &flags); err != nil {
			return err
		}

		fmt.Println("BossBar", uuid, action, title, health, color, division, flags)
	case 1:
		var health pk.Float

		if err := p.Scan(&health); err != nil {
			return err
		}

		fmt.Println("BossBar", uuid, action, health)
	case 2:
		var title pk.String

		if err := p.Scan(&title); err != nil {
			return err
		}

		fmt.Println("BossBar", uuid, action, title)
	case 3:
		var color pk.Byte

		if err := p.Scan(&color); err != nil {
			return err
		}

		fmt.Println("BossBar", uuid, action, color)
	case 4:
		var division pk.Byte

		if err := p.Scan(&division); err != nil {
			return err
		}

		fmt.Println("BossBar", uuid, action, division)
	case 5:
		var flags pk.Byte

		if err := p.Scan(&flags); err != nil {
			return err
		}

		fmt.Println("BossBar", uuid, action, flags)
	case 6:
		fmt.Println("BossBar", uuid, action)
	}

	return nil
}

func (e *EventsListener) ServerDifficulty(c *Client, p pk.Packet) error {
	var difficulty pk.Byte

	if err := p.Scan(&difficulty); err != nil {
		return err
	}

	fmt.Println("ServerDifficulty", difficulty)
	return nil
}

func (e *EventsListener) TabComplete(c *Client, p pk.Packet) error {
	/*var count pk.VarInt
	var matches []pk.String

	if err := p.Scan(&count, &matches); err != nil {
		return err
	}

	fmt.Println("TabComplete", count, matches)*/
	return nil
}

func (e *EventsListener) ChatMessage(c *Client, p pk.Packet) error {
	var json pk.String
	var position pk.Byte

	if err := p.Scan(&json, &position); err != nil {
		return err
	}

	switch position {
	case 0:
		/*var (
			Message           PlayerMessage
			SenderDisplayName pk.String
			SenderTeamName    pk.String
			Timestamp         pk.Long
		)

		if err := p.Scan(&SenderDisplayName, &SenderTeamName, &Timestamp); err != nil {
			return err
		}
		if err := json.Unmarshal(&Message); err != nil {
			return err
		}*/

		var msg chat.Message
		var pos pk.VarInt

		if err := p.Scan(&msg, &pos); err != nil {
			return err
		}
		fmt.Println("ChatMessage", msg, pos)
	case 1:
		var msg chat.Message
		var pos pk.VarInt

		if err := p.Scan(&msg, &pos); err != nil {
			return err
		}
		fmt.Println("SystemMessage", msg, pos)
	}

	return nil
}

func (e *EventsListener) MultiBlockChange(c *Client, p pk.Packet) error {
	/*var chunkX pk.Int
	var chunkZ pk.Int
	var recordCount pk.VarInt
	var records []struct {
		Position pk.Byte
		BlockID  pk.VarInt
	}

	if err := p.Scan(&chunkX, &chunkZ, &recordCount, &records); err != nil {
		return err
	}

	fmt.Println("MultiBlockChange", chunkX, chunkZ, recordCount, records)*/
	return nil
}

func (e *EventsListener) ConfirmTransaction(c *Client, p pk.Packet) error {
	var windowID pk.Byte
	var actionNumber pk.Short
	var accepted pk.Boolean

	if err := p.Scan(&windowID, &actionNumber, &accepted); err != nil {
		return err
	}

	fmt.Println("ConfirmTransaction", windowID, actionNumber, accepted)
	return nil
}

func (e *EventsListener) SetWindowContent(c *Client, p pk.Packet) error {
	var (
		ContainerID pk.UnsignedByte
		StateID     pk.VarInt
		SlotData    []Slot
		CarriedItem Slot
	)
	if err := p.Scan(
		&ContainerID,
		&StateID,
		pk.Array(&SlotData),
		&CarriedItem,
	); err != nil {
		return Error{err}
	}
	c.Player.Manager.StateID = int32(StateID)
	// copy the slot data to container
	container, ok := c.Player.Manager.Screens[int(ContainerID)]
	if !ok {
		return Error{errors.New("setting content of non-exist container")}
	}
	for i, v := range SlotData {
		err := container.OnSetSlot(i, v)
		if err != nil {
			return Error{err}
		}
		/*if m.events.SetSlot != nil {
			if err := m.events.SetSlot(int(ContainerID), i); err != nil {
				return Error{err}
			}
		}*/
	}

	fmt.Println("SetWindowContent", ContainerID, StateID, SlotData, CarriedItem)
	return nil
}

func (e *EventsListener) CloseWindow(c *Client, p pk.Packet) error {
	var windowID pk.Byte

	if err := p.Scan(&windowID); err != nil {
		return err
	}

	fmt.Println("CloseWindow", windowID)
	return nil
}

func (e *EventsListener) OpenWindow(c *Client, p pk.Packet) error {
	var windowID pk.Byte
	var windowType pk.String
	var windowTitle pk.String
	var slotCount pk.Byte
	var entityID pk.Int

	if err := p.Scan(&windowID, &windowType, &windowTitle, &slotCount, &entityID); err != nil {
		return err
	}

	fmt.Println("OpenWindow", windowID, windowType, windowTitle, slotCount, entityID)
	return nil
}

func (e *EventsListener) WindowItems(c *Client, p pk.Packet) error {
	/*var windowID pk.Byte
	var count pk.Short
	var items []pk.Slot

	if err := p.Scan(&windowID, &count, &items); err != nil {
		return err
	}

	fmt.Println("WindowItems", windowID, count, items)*/
	return nil
}

func (e *EventsListener) WindowProperty(c *Client, p pk.Packet) error {
	var windowID pk.Byte
	var property pk.Short
	var value pk.Short

	if err := p.Scan(&windowID, &property, &value); err != nil {
		return err
	}

	fmt.Println("WindowProperty", windowID, property, value)
	return nil
}

func (e *EventsListener) SetSlot(c *Client, p pk.Packet) (err error) {
	var (
		ContainerID pk.Byte
		StateID     pk.VarInt
		SlotID      pk.Short
		SlotData    Slot
	)
	if err := p.Scan(&ContainerID, &StateID, &SlotID, &SlotData); err != nil {
		return Error{err}
	}

	c.Player.Manager.StateID = int32(StateID)
	if ContainerID == -1 && SlotID == -1 {
		c.Player.Manager.Cursor = SlotData
	} else if ContainerID == -2 {
		err = c.Player.Manager.Inventory.OnSetSlot(int(SlotID), SlotData)
	} else if c, ok := c.Player.Manager.Screens[int(ContainerID)]; ok {
		err = c.OnSetSlot(int(SlotID), SlotData)
	}

	/*if m.events.SetSlot != nil {
		if err := m.events.SetSlot(int(ContainerID), int(SlotID)); err != nil {
			return Error{err}
		}
	}*/
	if err != nil {
		return Error{err}
	}
	return nil
}

func (e *EventsListener) SetCooldown(c *Client, p pk.Packet) error {
	var itemID pk.VarInt
	var cooldown pk.VarInt

	if err := p.Scan(&itemID, &cooldown); err != nil {
		return err
	}

	fmt.Println("SetCooldown", itemID, cooldown)
	return nil
}

func (e *EventsListener) PluginMessage(c *Client, p pk.Packet) error {
	var channel pk.String
	var data pk.ByteArray

	if err := p.Scan(&channel, &data); err != nil {
		return err
	}

	fmt.Println("PluginMessage", channel, data)
	return nil
}

func (e *EventsListener) NamedSoundEffect(c *Client, p pk.Packet) error {
	var soundName pk.String
	var soundCategory pk.Byte
	var effectPosition pk.Position
	var volume pk.Float
	var pitch pk.Float

	if err := p.Scan(&soundName, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return err
	}

	fmt.Println("NamedSoundEffect", soundName, soundCategory, effectPosition, volume, pitch)
	return nil
}

func (e *EventsListener) Disconnect(c *Client, p pk.Packet) error {
	var reason chat.Message
	if err := p.Scan(&reason); err != nil {
		return Error{err}
	}

	return nil
}

func (e *EventsListener) EntityStatus(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var entityStatus pk.Byte

	if err := p.Scan(&entityID, &entityStatus); err != nil {
		return err
	}

	fmt.Println("EntityStatus", entityID, entityStatus)
	return nil
}

func (e *EventsListener) Explosion(c *Client, p pk.Packet) error {
	/*var x pk.Float
	var y pk.Float
	var z pk.Float
	var radius pk.Float
	var recordCount pk.Int
	var records []struct {
		X pk.Byte
		Y pk.Byte
		Z pk.Byte
	}
	var playerMotionX pk.Float
	var playerMotionY pk.Float
	var playerMotionZ pk.Float

	if err := p.Scan(&x, &y, &z, &radius, &recordCount, &records, &playerMotionX, &playerMotionY, &playerMotionZ); err != nil {
		return err
	}

	fmt.Println("Explosion", x, y, z, radius, recordCount, records, playerMotionX, playerMotionY, playerMotionZ)*/
	return nil
}

func (e *EventsListener) UnloadChunk(c *Client, p pk.Packet) error {
	var chunk level.ChunkPos

	if err := p.Scan(&chunk); err != nil {
		return err
	}

	delete(c.World.Columns, chunk)
	return nil
}

func (e *EventsListener) ChangeGameState(c *Client, p pk.Packet) error {
	var reason pk.UnsignedByte
	var value pk.Float

	if err := p.Scan(&reason, &value); err != nil {
		return err
	}

	fmt.Println("ChangeGameState", reason, value)
	return nil
}

func (e *EventsListener) KeepAlive(c *Client, p pk.Packet) error {
	var keepAliveID pk.Long

	if err := p.Scan(&keepAliveID); err != nil {
		return err
	}

	if err := c.Conn.WritePacket(
		pk.Marshal(
			packetid.SPacketKeepAlive,
			keepAliveID,
		),
	); err != nil {
		return err
	}

	return nil
}

func (e *EventsListener) ChunkData(c *Client, p pk.Packet) error {
	var (
		ChunkPos level.ChunkPos
		Chunk    level.Chunk
	)

	if err := p.Scan(
		&ChunkPos, &Chunk,
	); err != nil {
		return err
	}

	fmt.Println("ChunkData", ChunkPos, len(Chunk.Sections), len(c.World.Columns))
	c.World.Columns[ChunkPos] = &Chunk

	return nil
}

func (e *EventsListener) Effect(c *Client, p pk.Packet) error {
	var effectID pk.Int
	var location pk.Position
	var data pk.Int
	var disableRelativeVolume pk.Boolean

	if err := p.Scan(&effectID, &location, &data, &disableRelativeVolume); err != nil {
		return err
	}

	fmt.Println("Effect", effectID, location, data, disableRelativeVolume)
	return nil
}

func (e *EventsListener) Particle(c *Client, p pk.Packet) error {
	/*var particleID pk.String
	var longDistance pk.Boolean
	var x pk.Float
	var y pk.Float
	var z pk.Float
	var offsetX pk.Float
	var offsetY pk.Float
	var offsetZ pk.Float
	var particleData pk.Float
	var particleCount pk.Int
	var data []pk.Int

	if err := p.Scan(&particleID, &longDistance, &x, &y, &z, &offsetX, &offsetY, &offsetZ, &particleData, &particleCount, &data); err != nil {
		return err
	}

	fmt.Println("Particle", particleID, longDistance, x, y, z, offsetX, offsetY, offsetZ, particleData, particleCount, data)*/
	return nil
}

func (e *EventsListener) JoinGame(c *Client, p pk.Packet) error {
	err := p.Scan(
		(*pk.Int)(&c.Player.EID),
		(*pk.Boolean)(&c.Player.Hardcore),
		(*pk.UnsignedByte)(&c.Player.Gamemode),
		(*pk.Byte)(&c.Player.PrevGamemode),
		pk.Array((*[]pk.Identifier)(unsafe.Pointer(&c.Player.DimensionNames))),
		pk.NBT(&c.Player.WorldInfo.DimensionCodec),
		(*pk.Identifier)(&c.Player.WorldInfo.DimensionType),
		(*pk.Identifier)(&c.Player.DimensionName),
		(*pk.Long)(&c.Player.HashedSeed),
		(*pk.VarInt)(&c.Player.MaxPlayers),
		(*pk.VarInt)(&c.Player.ViewDistance),
		(*pk.VarInt)(&c.Player.SimulationDistance),
		(*pk.Boolean)(&c.Player.ReducedDebugInfo),
		(*pk.Boolean)(&c.Player.EnableRespawnScreen),
		(*pk.Boolean)(&c.Player.IsDebug),
		(*pk.Boolean)(&c.Player.IsFlat),
		pk.Opt{
			If:    (*pk.Boolean)(&c.Player.HasDeathLocation),
			Value: (*pk.Position)(&c.Player.WorldInfo.DeathPosition),
		},
	)
	if err != nil {
		return Error{err}
	}
	err = c.Conn.WritePacket(pk.Marshal( //PluginMessage packet
		packetid.SPacketPluginMessage,
		pk.Identifier("minecraft:brand"),
		pk.String(c.Player.Settings.Brand),
	))
	if err != nil {
		return Error{err}
	}

	err = c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientSettings,
		pk.String(c.Player.Settings.Locale),
		pk.Byte(c.Player.Settings.ViewDistance),
		pk.VarInt(c.Player.Settings.ChatMode),
		pk.Boolean(c.Player.Settings.ChatColors),
		pk.UnsignedByte(c.Player.Settings.DisplayedSkinParts),
		pk.VarInt(c.Player.Settings.MainHand),
		pk.Boolean(c.Player.Settings.EnableTextFiltering),
		pk.Boolean(c.Player.Settings.AllowListing),
	))
	if err != nil {
		return Error{err}
	}

	fmt.Println("JoinGame", c.Player.EID, c.Player.Hardcore, c.Player.Gamemode, c.Player.PrevGamemode, c.Player.DimensionNames, c.Player.WorldInfo.DimensionCodec, c.Player.WorldInfo.DimensionType, c.Player.DimensionName, c.Player.HashedSeed, c.Player.MaxPlayers, c.Player.ViewDistance, c.Player.SimulationDistance, c.Player.ReducedDebugInfo, c.Player.EnableRespawnScreen, c.Player.IsDebug, c.Player.IsFlat)
	return nil
}

func (e *EventsListener) Map(c *Client, p pk.Packet) error {
	/*var itemDamage pk.VarInt
	var scale pk.Byte
	var trackingPosition pk.Boolean
	var icons []struct {
		Direction pk.Byte
		X         pk.Byte
		Y         pk.Byte
	}
	var columns pk.VarInt
	var rows pk.VarInt
	var x pk.VarInt
	var z pk.VarInt
	var data pk.ByteArray

	if err := p.Scan(&itemDamage, &scale, &trackingPosition, &icons, &columns, &rows, &x, &z, &data); err != nil {
		return err
	}

	fmt.Println("Map", itemDamage, scale, trackingPosition, icons, columns, rows, x, z, data)*/
	return nil
}

func (e *EventsListener) Entity(c *Client, p pk.Packet) error {
	var entityID pk.Int

	if err := p.Scan(&entityID); err != nil {
		return err
	}

	fmt.Println("Entity", entityID)
	return nil
}

func (e *EventsListener) EntityPosition(c *Client, p pk.Packet) error {
	var (
		EntityID               pk.VarInt
		DeltaX, DeltaY, DeltaZ pk.Short
		OnGround               pk.Boolean
	)

	if err := p.Scan(&EntityID, &DeltaX, &DeltaY, &DeltaZ, &OnGround); err != nil {
		return err
	}

	fmt.Println("EntityRelativeMove", EntityID, DeltaX, DeltaY, DeltaZ, OnGround)
	return nil
}

func (e *EventsListener) EntityPositionRotation(c *Client, p pk.Packet) error {
	var (
		EntityID               pk.VarInt
		DeltaX, DeltaY, DeltaZ pk.Short
		Yaw, Pitch             pk.Angle
		OnGround               pk.Boolean
	)

	if err := p.Scan(&EntityID, &DeltaX, &DeltaY, &DeltaZ, &Yaw, &Pitch, &OnGround); err != nil {
		return err
	}

	fmt.Println("EntityPositionRotation", EntityID, DeltaX, DeltaY, DeltaZ, Yaw, Pitch, OnGround)
	return nil
}

func (e *EventsListener) EntityHeadRotation(c *Client, p pk.Packet) error {
	var (
		EntityID pk.VarInt
		HeadYaw  pk.Angle
	)

	if err := p.Scan(&EntityID, &HeadYaw); err != nil {
		return err
	}

	fmt.Println("EntityHeadRotation", EntityID, HeadYaw)
	return nil
}

func (e *EventsListener) EntityRotation(c *Client, p pk.Packet) error {
	var (
		EntityID   pk.VarInt
		Yaw, Pitch pk.Angle
		OnGround   pk.Boolean
	)

	if err := p.Scan(&EntityID, &Yaw, &Pitch, &OnGround); err != nil {
		return err
	}

	fmt.Println("EntityRotation", EntityID, Yaw, Pitch, OnGround)
	return nil
}

func (e *EventsListener) VehicleMove(c *Client, p pk.Packet) error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return err
	}

	fmt.Println("VehicleMove", X, Y, Z, Yaw, Pitch)
	return nil
}

func (e *EventsListener) OpenSignEditor(c *Client, p pk.Packet) error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return err
	}

	fmt.Println("OpenSignEditor", location)
	return nil
}

func (e *EventsListener) CraftRecipeResponse(c *Client, p pk.Packet) error {
	var windowID pk.UnsignedByte
	var recipe pk.String

	if err := p.Scan(&windowID, &recipe); err != nil {
		return err
	}

	fmt.Println("CraftRecipeResponse", windowID, recipe)
	return nil
}

func (e *EventsListener) PlayerAbilities(c *Client, p pk.Packet) error {
	var flags pk.UnsignedByte
	var flyingSpeed pk.Float
	var fov pk.Float

	if err := p.Scan(&flags, &flyingSpeed, &fov); err != nil {
		return err
	}

	fmt.Println("PlayerAbilities", flags, flyingSpeed, fov)
	return nil
}

func (e *EventsListener) CombatEvent(c *Client, p pk.Packet) error {
	var event pk.Byte
	var duration pk.Int
	var entityID pk.Int
	var message pk.String

	if err := p.Scan(&event, &duration, &entityID, &message); err != nil {
		return err
	}

	fmt.Println("CombatEvent", event, duration, entityID, message)
	return nil
}

func (e *EventsListener) PlayerInfo(c *Client, p pk.Packet) error {
	/*var action pk.Byte
	var players []struct {
		UUID       pk.UUID
		Name       pk.String
		Properties []struct {
			Name      pk.String
			Value     pk.String
			Signature pk.String
		}
		GameMode    pk.UnsignedByte
		Latency     pk.VarInt
		DisplayName pk.String
	}

	if err := p.Scan(&action, &players); err != nil {
		return err
	}

	fmt.Println("PlayerInfo", action, players)*/
	return nil
}

func (e *EventsListener) SyncPlayerPosition(c *Client, p pk.Packet) error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
		Flags      pk.Byte
		TeleportID pk.VarInt
		Dismount   pk.Boolean
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch, &Flags, &TeleportID, &Dismount); err != nil {
		return err
	}

	position := maths.Vec3d{X: float32(X), Y: float32(Y), Z: float32(Z)}
	rotation := maths.Vec2d{X: float32(Pitch), Y: float32(Yaw)}

	if Flags&0x01 != 0 {
		c.Player.Position = c.Player.Position.Add(position)
		c.Player.Rotation = c.Player.Rotation.Add(rotation)
	} else {
		c.Player.Position = position
		c.Player.Rotation = rotation
	}
	result, err := c.World.RayTrace(c.Player.Rotation, c.Player.GetEyePos(), 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.Block)
	}
	fmt.Println("SyncPlayerPosition", position, rotation, TeleportID, Dismount)
	return nil
}

func (e *EventsListener) PlayerPositionAndLook(c *Client, p pk.Packet) error {
	var x pk.Double
	var y pk.Double
	var z pk.Double
	var yaw pk.Angle
	var pitch pk.Angle
	var flags pk.Byte
	var teleportID pk.VarInt

	if err := p.Scan(&x, &y, &z, &yaw, &pitch, &flags, &teleportID); err != nil {
		return err
	}

	fmt.Println("PlayerPositonAndLook", x, y, z, yaw, pitch, flags, teleportID)
	return nil
}

func (e *EventsListener) UseBed(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var location pk.Position

	if err := p.Scan(&entityID, &location); err != nil {
		return err
	}

	fmt.Println("UseBed", entityID, location)
	return nil
}

func (e *EventsListener) UnlockRecipes(c *Client, p pk.Packet) error {
	/*var action pk.Byte
	var craftingBookOpen pk.Boolean
	var filter pk.Boolean
	var recipes []pk.String

	if err := p.Scan(&action, &craftingBookOpen, &filter, &recipes); err != nil {
		return err
	}

	fmt.Println("UnlockRecipes", action, craftingBookOpen, filter, recipes)*/
	return nil
}

func (e *EventsListener) DestroyEntities(c *Client, p pk.Packet) error {
	/*var entityIDs []pk.Int

	if err := p.Scan(&entityIDs); err != nil {
		return err
	}

	fmt.Println("DestroyEntities", entityIDs)*/
	return nil
}

func (e *EventsListener) RemoveEntityEffect(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var effectID pk.Byte

	if err := p.Scan(&entityID, &effectID); err != nil {
		return err
	}

	fmt.Println("RemoveEntityEffect", entityID, effectID)
	return nil
}

func (e *EventsListener) ResourcePackSend(c *Client, p pk.Packet) error {
	var url pk.String
	var hash pk.String

	if err := p.Scan(&url, &hash); err != nil {
		return err
	}

	fmt.Println("ResourcePackSend", url, hash)
	return nil
}

func (e *EventsListener) Respawn(c *Client, p pk.Packet) error {
	var copyMeta bool
	if err := p.Scan(
		(*pk.String)(&c.Player.DimensionType),
		(*pk.Identifier)(&c.Player.DimensionName),
		(*pk.Long)(&c.Player.HashedSeed),
		(*pk.UnsignedByte)(&c.Player.Gamemode),
		(*pk.Byte)(&c.Player.PrevGamemode),
		(*pk.Boolean)(&c.Player.IsDebug),
		(*pk.Boolean)(&c.Player.IsFlat),
		(*pk.Boolean)(&copyMeta),
	); err != nil {
		return err
	}

	return nil
}

func (e *EventsListener) SelectAdvancementTab(c *Client, p pk.Packet) error {
	var hasID pk.Boolean
	var identifier pk.String

	if err := p.Scan(&hasID, &identifier); err != nil {
		return err
	}

	fmt.Println("SelectAdvancementTab", hasID, identifier)
	return nil
}

func (e *EventsListener) WorldBorder(c *Client, p pk.Packet) error {
	var action pk.Byte
	var radius pk.Double
	var oldRadius pk.Double
	var speed pk.VarInt
	var x pk.Int
	var z pk.Int
	var portalBoundary pk.VarInt
	var warningTime pk.VarInt
	var warningBlocks pk.VarInt

	if err := p.Scan(&action, &radius, &oldRadius, &speed, &x, &z, &portalBoundary, &warningTime, &warningBlocks); err != nil {
		return err
	}

	fmt.Println("WorldBorder", action, radius, oldRadius, speed, x, z, portalBoundary, warningTime, warningBlocks)
	return nil
}

func (e *EventsListener) Camera(c *Client, p pk.Packet) error {
	var cameraID pk.Int

	if err := p.Scan(&cameraID); err != nil {
		return err
	}

	fmt.Println("Camera", cameraID)
	return nil
}

func (e *EventsListener) HeldItemChange(c *Client, p pk.Packet) error {
	var slot pk.Short

	if err := p.Scan(&slot); err != nil {
		return err
	}

	newSlot, err := c.Player.Inventory.GetHotbarSlotById(uint8(slot))
	if err != nil {
		return err
	}
	fmt.Println("HeldItemChange", slot, "New Item:", item.ByID[item.ID(newSlot.ID)].Name)
	return nil
}

func (e *EventsListener) DisplayScoreboard(c *Client, p pk.Packet) error {
	var position pk.Byte
	var name pk.String

	if err := p.Scan(&position, &name); err != nil {
		return err
	}

	fmt.Println("DisplayScoreboard", position, name)
	return nil
}

func (e *EventsListener) EntityMetadata(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var metadata pk.Metadata

	if err := p.Scan(&entityID, &metadata); err != nil {
		return err
	}

	fmt.Println("EntityMetadata", entityID, metadata)*/
	return nil
}

func (e *EventsListener) AttachEntity(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var vehicleID pk.Int
	var leash pk.Boolean

	if err := p.Scan(&entityID, &vehicleID, &leash); err != nil {
		return err
	}

	fmt.Println("AttachEntity", entityID, vehicleID, leash)
	return nil
}

func (e *EventsListener) EntityVelocity(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var velocityX, velocityY, velocityZ pk.Short

	if err := p.Scan(&entityID, &velocityX, &velocityY, &velocityZ); err != nil {
		return err
	}

	fmt.Println("EntityVelocity", entityID, velocityX, velocityY, velocityZ)
	return nil
}

func (e *EventsListener) EntityEquipment(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var slot pk.Short
	var item pk.Slot

	if err := p.Scan(&entityID, &slot, &item); err != nil {
		return err
	}

	fmt.Println("EntityEquipment", entityID, slot, item)*/
	return nil
}

func (e *EventsListener) SetExperience(c *Client, p pk.Packet) error {
	var experienceBar pk.Float
	var level pk.VarInt
	var totalExperience pk.VarInt

	if err := p.Scan(&experienceBar, &level, &totalExperience); err != nil {
		return err
	}

	fmt.Println("SetExperience", experienceBar, level, totalExperience)
	c.Player.SetExp(float32(experienceBar), int32(level), int32(totalExperience))
	return nil
}

func (e *EventsListener) UpdateHealth(c *Client, p pk.Packet) error {
	var (
		health         pk.Float
		food           pk.VarInt
		foodSaturation pk.Float
	)

	if err := p.Scan(&health, &food, &foodSaturation); err != nil {
		return err
	}

	fmt.Println("UpdateHealth", health, food, foodSaturation)
	if respawn := c.Player.SetHealth(float32(health)); respawn {
		if err := c.Player.Respawn(); err != nil {
			return err
		}
	}
	return nil
}

func (e *EventsListener) ScoreboardObjective(c *Client, p pk.Packet) error {
	var (
		name           pk.String
		mode           pk.Byte
		objectiveName  pk.String
		objectiveValue pk.String
		type_          pk.Byte
	)

	if err := p.Scan(&name, &mode, &objectiveName, &objectiveValue, &type_); err != nil {
		return err
	}

	fmt.Println("ScoreboardObjective", name, mode, objectiveName, objectiveValue, type_)
	return nil
}

func (e *EventsListener) SetPassengers(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var passengerCount pk.VarInt
	var passengers pk.Ary[]pk.VarInt

	if err := p.Scan(&entityID, &passengers); err != nil {
		return err
	}

	fmt.Println("SetPassengers", entityID, passengerCount, passengers)*/
	return nil
}

func (e *EventsListener) Teams(c *Client, p pk.Packet) error {
	/*var name pk.String
	var mode pk.Byte
	var teamName pk.String
	var displayName pk.String
	var prefix pk.String
	var suffix pk.String
	var friendlyFire pk.Byte
	var nameTagVisibility pk.String
	var collisionRule pk.String
	var color pk.Byte
	var playerCount pk.VarInt
	var players pk.Ary[]pk.String

	if err := p.Scan(&name, &mode, &teamName, &displayName, &prefix, &suffix, &friendlyFire, &nameTagVisibility, &collisionRule, &color, &players); err != nil {
		return err
	}

	fmt.Println("Teams", name, mode, teamName, displayName, prefix, suffix, friendlyFire, nameTagVisibility, collisionRule, color, playerCount, players)*/
	return nil
}

func (e *EventsListener) UpdateScore(c *Client, p pk.Packet) error {
	var name pk.String
	var action pk.Byte
	var objectiveName pk.String
	var value pk.VarInt

	if err := p.Scan(&name, &action, &objectiveName, &value); err != nil {
		return err
	}

	fmt.Println("UpdateScore", name, action, objectiveName, value)
	return nil
}

func (e *EventsListener) SpawnPosition(c *Client, p pk.Packet) error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return err
	}

	fmt.Println("SpawnPosition", location)
	return nil
}

func (e *EventsListener) TimeUpdate(c *Client, p pk.Packet) error {
	var (
		WorldAge  pk.Long
		TimeOfDay pk.Long
	)

	c.TPS.Update()

	if err := p.Scan(&WorldAge, &TimeOfDay); err != nil {
		return err
	}
	return nil
}

func (e *EventsListener) Title(c *Client, p pk.Packet) error {
	var action pk.Byte
	var fadeIn pk.Int
	var stay pk.Int
	var fadeOut pk.Int
	var title pk.String
	var subtitle pk.String
	var actionBar pk.String

	if err := p.Scan(&action, &fadeIn, &stay, &fadeOut, &title, &subtitle, &actionBar); err != nil {
		return err
	}

	fmt.Println("Title", action, fadeIn, stay, fadeOut, title, subtitle, actionBar)
	return nil
}

func (e *EventsListener) SoundEffect(c *Client, p pk.Packet) error {
	var soundID pk.VarInt
	var soundCategory pk.VarInt
	var effectPosition pk.Position
	var volume pk.Float
	var pitch pk.Float

	if err := p.Scan(&soundID, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return err
	}

	fmt.Println("SoundEffect", soundID, soundCategory, effectPosition, volume, pitch)
	return nil
}

func (e *EventsListener) PlayerListHeaderAndFooter(c *Client, p pk.Packet) error {
	var header pk.String
	var footer pk.String

	if err := p.Scan(&header, &footer); err != nil {
		return err
	}

	fmt.Println("PlayerListHeaderAndFooter", header, footer)
	return nil
}

func (e *EventsListener) CollectItem(c *Client, p pk.Packet) error {
	var collectedEntityID pk.Int
	var collectorEntityID pk.Int
	var pickupCount pk.VarInt

	if err := p.Scan(&collectedEntityID, &collectorEntityID, &pickupCount); err != nil {
		return err
	}

	fmt.Println("CollectItem", collectedEntityID, collectorEntityID, pickupCount)
	return nil
}

func (e *EventsListener) EntityTeleport(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var x pk.Double
	var y pk.Double
	var z pk.Double
	var yaw pk.Byte
	var pitch pk.Byte
	var onGround pk.Boolean

	if err := p.Scan(&entityID, &x, &y, &z, &yaw, &pitch, &onGround); err != nil {
		return err
	}

	fmt.Println("EntityTeleport", entityID, x, y, z, yaw, pitch, onGround)
	return nil
}

func (e *EventsListener) Advancements(c *Client, p pk.Packet) error {
	var action pk.Byte
	var data pk.String

	if err := p.Scan(&action, &data); err != nil {
		return err
	}

	fmt.Println("Advancements", action, data)
	return nil
}

func (e *EventsListener) EntityProperties(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var count pk.VarInt
	var properties pk.Ary[]pk.String

	if err := p.Scan(&entityID, &properties); err != nil {
		return err
	}

	fmt.Println("EntityProperties", entityID, count, properties)*/
	return nil
}

func (e *EventsListener) EntityEffect(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var effectID pk.Byte
	var amplifier pk.Byte
	var duration pk.VarInt
	var hideParticles pk.Boolean

	if err := p.Scan(&entityID, &effectID, &amplifier, &duration, &hideParticles); err != nil {
		return err
	}

	fmt.Println("EntityEffect", entityID, effectID, amplifier, duration, hideParticles)
	return nil
}
