package provider

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/effects"
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

	c.Events.AddTicker(
		TickHandler{Priority: 0, F: ApplyPhysics},
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

func (e *EventsListener) SpawnEntity(c *Client, p pk.Packet) basic.Error {
	var (
		EntityID            pk.VarInt
		EntityUUID          pk.UUID
		TypeID              pk.Byte
		X, Y, Z             pk.Double
		Pitch, Yaw, HeadYaw pk.Angle
		Data                pk.VarInt
		vX, vY, vZ          pk.Short
	)

	if err := p.Scan(&EntityID, &EntityUUID, &TypeID, &X, &Y, &Z, &Pitch, &Yaw, &HeadYaw, &Data, &vX, &vY, &vZ); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SpawnEntity packet: %w", err)}
	}

	if err := c.World.AddEntity(core.NewEntity(
		int32(EntityID),
		uuid.UUID(EntityUUID),
		int32(TypeID),
		float64(X), float64(Y), float64(Z),
		float64(Pitch), float64(Yaw),
	)); err != nil {
		return basic.Error{Err: basic.InvalidEntity, Info: err}
	}

	fmt.Println("SpawnEntity", EntityID, EntityUUID, TypeID, X, Y, Z, Pitch, Yaw, HeadYaw, Data, vX, vY, vZ)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SpawnExperienceOrb(c *Client, p pk.Packet) basic.Error {
	var (
		entityID pk.VarInt
		x, y, z  pk.Double
		count    pk.Short
	)

	if err := p.Scan(&entityID, &x, &y, &z, &count); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SpawnExperienceOrb packet: %w", err)}
	}

	fmt.Println("SpawnExperienceOrb", entityID, x, y, z, count)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SpawnPlayer(c *Client, p pk.Packet) basic.Error {
	var (
		EntityID   pk.VarInt
		PlayerUUID pk.UUID
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Angle
	)

	if err := p.Scan(&EntityID, &PlayerUUID, &X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SpawnPlayer packet: %w", err)}
	}

	if err := c.World.AddEntity(core.NewEntity(
		int32(EntityID),
		uuid.UUID(PlayerUUID),
		116, // Player type
		float64(X), float64(Y), float64(Z),
		float64(Pitch), float64(Yaw),
	)); err != nil {
		return basic.Error{Err: basic.InvalidEntity, Info: err}
	}

	fmt.Println("SpawnPlayer", EntityID, PlayerUUID.String(), X, Y, Z, Yaw, Pitch)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityAnimation(c *Client, p pk.Packet) basic.Error {
	var (
		entityID  pk.VarInt
		animation pk.Byte
	)

	if err := p.Scan(&entityID, &animation); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Animation packet: %w", err)}
	}

	fmt.Println("Animation", entityID, animation)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) AwardStatistics(c *Client, p pk.Packet) basic.Error {
	/*var count pk.VarInt
	var statistics []struct {
		Name  pk.String
		Value pk.VarInt
	}*/

	/*if err := p.Scan(&count, &statistics); err != nil {
		return err
	}

	fmt.Println("Statistics", count, statistics)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SetBlockDestroyStage(c *Client, p pk.Packet) basic.Error {
	var (
		entityID pk.VarInt
		location pk.Position
		stage    pk.Byte
	)

	if err := p.Scan(&entityID, &location, &stage); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BlockBreakAnimation packet: %w", err)}
	}

	fmt.Println("BlockBreakAnimation", entityID, location, stage)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) BlockEntityData(c *Client, p pk.Packet) basic.Error {
	/*var location pk.Position
	var action pk.Byte
	var nbtData pk

	if err := p.Scan(&location, &action, &nbtData); err != nil {
		return err
	}

	fmt.Println("UpdateBlockEntity", location, action, nbtData)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) BlockAction(c *Client, p pk.Packet) basic.Error {
	var (
		location    pk.Position
		actionID    pk.Byte
		actionParam pk.Byte
		blockType   pk.VarInt
	)

	if err := p.Scan(&location, &actionID, &actionParam, &blockType); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BlockAction packet: %w", err)}
	}

	fmt.Println("BlockAction", location, actionID, actionParam, blockType)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) BlockChange(c *Client, p pk.Packet) basic.Error {
	var (
		location  pk.Position
		blockType pk.VarInt
	)

	if err := p.Scan(&location, &blockType); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BlockChange packet: %w", err)}
	}

	fmt.Println("BlockChange", location, blockType)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) BossBar(c *Client, p pk.Packet) basic.Error {
	var uuid pk.UUID
	var action pk.Byte

	if err := p.Scan(&uuid, &action); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
	}

	switch action {
	case 0:
		var (
			title     pk.String
			health    pk.Float
			color     pk.Byte
			divisions pk.Byte
			flags     pk.Byte
		)

		if err := p.Scan(&title, &health, &color, &divisions, &flags); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
		}

		fmt.Println("BossBar", uuid, action, title, health, color, divisions, flags)
	case 1:
		var health pk.Float

		if err := p.Scan(&health); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
		}

		fmt.Println("BossBar", uuid, action, health)
	case 2:
		var title pk.String

		if err := p.Scan(&title); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
		}

		fmt.Println("BossBar", uuid, action, title)
	case 3:
		var color pk.Byte

		if err := p.Scan(&color); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
		}

		fmt.Println("BossBar", uuid, action, color)
	case 4:
		var division pk.Byte

		if err := p.Scan(&division); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
		}

		fmt.Println("BossBar", uuid, action, division)
	case 5:
		var flags pk.Byte

		if err := p.Scan(&flags); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read BossBar packet: %w", err)}
		}

		fmt.Println("BossBar", uuid, action, flags)
	case 6:
		fmt.Println("BossBar", uuid, action)
	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ServerDifficulty(c *Client, p pk.Packet) basic.Error {
	var difficulty pk.Byte

	if err := p.Scan(&difficulty); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ServerDifficulty packet: %w", err)}
	}

	fmt.Println("ServerDifficulty", difficulty)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) TabComplete(c *Client, p pk.Packet) basic.Error {
	/*var count pk.VarInt
	var matches []pk.String

	if err := p.Scan(&count, &matches); err != nil {
		return err
	}

	fmt.Println("TabComplete", count, matches)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ChatMessage(c *Client, p pk.Packet) basic.Error {
	var (
		json     pk.String
		position pk.Byte
		/*msg      chat.Message
		pos      pk.VarInt*/
	)

	if err := p.Scan(&json, &position); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ChatMessage packet: %w", err)}
	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) MultiBlockChange(c *Client, p pk.Packet) basic.Error {
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
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ConfirmTransaction(c *Client, p pk.Packet) basic.Error {
	var windowID pk.Byte
	var actionNumber pk.Short
	var accepted pk.Boolean

	if err := p.Scan(&windowID, &actionNumber, &accepted); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ConfirmTransaction packet: %w", err)}
	}

	fmt.Println("ConfirmTransaction", windowID, actionNumber, accepted)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SetWindowContent(c *Client, p pk.Packet) basic.Error {
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
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("failed to scan SetWindowContent")}
	}
	c.Player.Manager.StateID = int32(StateID)
	// copy the slot data to container
	container, ok := c.Player.Manager.Screens[int(ContainerID)]
	if !ok {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("failed to find container with id %d", ContainerID)}
	}
	for i, v := range SlotData {
		err := container.OnSetSlot(i, v)
		if !err.Is(basic.NoError) {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("failed to set slot %d: %w", i, err)}
		}
		/*if m.events.SetSlot != nil {
			if err := m.events.SetSlot(int(ContainerID), i); err != nil {
				return basic.Error{err}
			}
		}*/
	}

	fmt.Println("SetWindowContent", ContainerID, StateID, SlotData, CarriedItem)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) CloseWindow(c *Client, p pk.Packet) basic.Error {
	var windowID pk.Byte

	if err := p.Scan(&windowID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read CloseWindow packet: %w", err)}
	}

	fmt.Println("CloseWindow", windowID)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) OpenWindow(c *Client, p pk.Packet) basic.Error {
	var (
		windowID    pk.Byte
		windowType  pk.String
		windowTitle pk.String
		slotCount   pk.Byte
		entityID    pk.Int
	)

	if err := p.Scan(&windowID, &windowType, &windowTitle, &slotCount, &entityID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read OpenWindow packet: %w", err)}
	}

	fmt.Println("OpenWindow", windowID, windowType, windowTitle, slotCount, entityID)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) WindowItems(c *Client, p pk.Packet) basic.Error {
	/*var windowID pk.Byte
	var count pk.Short
	var items []pk.Slot

	if err := p.Scan(&windowID, &count, &items); err != nil {
		return err
	}

	fmt.Println("WindowItems", windowID, count, items)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) WindowProperty(c *Client, p pk.Packet) basic.Error {
	var windowID pk.Byte
	var property pk.Short
	var value pk.Short

	if err := p.Scan(&windowID, &property, &value); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read WindowProperty packet: %w", err)}
	}

	fmt.Println("WindowProperty", windowID, property, value)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SetSlot(c *Client, p pk.Packet) (err basic.Error) {
	var (
		ContainerID pk.Byte
		StateID     pk.VarInt
		SlotID      pk.Short
		SlotData    Slot
	)
	if err := p.Scan(&ContainerID, &StateID, &SlotID, &SlotData); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("failed to scan SetSlot")}
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
			return basic.Error{err}
		}
	}*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SetCooldown(c *Client, p pk.Packet) basic.Error {
	var (
		itemID pk.VarInt
		ticks  pk.VarInt
	)

	if err := p.Scan(&itemID, &ticks); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SetCooldown packet: %w", err)}
	}

	fmt.Println("SetCooldown", itemID, ticks)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) PluginMessage(c *Client, p pk.Packet) basic.Error {
	var (
		channel pk.String
		data    pk.ByteArray
	)

	if err := p.Scan(&channel, &data); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read PluginMessage packet: %w", err)}
	}

	fmt.Println("PluginMessage", channel, data)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) NamedSoundEffect(c *Client, p pk.Packet) basic.Error {
	var (
		soundName      pk.String
		soundCategory  pk.Byte
		effectPosition pk.Position
		volume         pk.Float
		pitch          pk.Float
	)

	if err := p.Scan(&soundName, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read NamedSoundEffect packet: %w", err)}
	}

	fmt.Println("NamedSoundEffect", soundName, soundCategory, effectPosition, volume, pitch)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Disconnect(c *Client, p pk.Packet) basic.Error {
	var reason chat.Message
	if err := p.Scan(&reason); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("failed to scan Disconnect: %w", err)}
	}

	fmt.Println("Disconnect", reason)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityStatus(c *Client, p pk.Packet) basic.Error {
	var entityID pk.Int
	var entityStatus pk.Byte

	if err := p.Scan(&entityID, &entityStatus); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityStatus packet: %w", err)}
	}

	fmt.Println("EntityStatus", entityID, entityStatus)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Explosion(c *Client, p pk.Packet) basic.Error {
	var (
		x, y, z    pk.Float
		radius     pk.Float
		records    pk.VarInt
		data       = make([][3]pk.VarInt, 0)
		mX, mY, mZ pk.Float
	)

	if err := p.Scan(&x, &y, &z, &radius, &records); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Explosion packet: %w", err)}
	}

	data = make([][3]pk.VarInt, records)
	for i := pk.VarInt(0); i < records; i++ {
		if err := p.Scan(&data[i][0], &data[i][1], &data[i][2]); err != nil {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Explosion packet: %w", err)}
		}
	}

	if err := p.Scan(&mX, &mY, &mZ); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Explosion packet: %w", err)}
	}

	fmt.Println("Explosion", x, y, z, radius, records, data, mX, mY, mZ)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) UnloadChunk(c *Client, p pk.Packet) basic.Error {
	var chunk level.ChunkPos

	if err := p.Scan(&chunk); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read UnloadChunk packet: %w", err)}
	}

	delete(c.World.Columns, chunk)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ChangeGameState(c *Client, p pk.Packet) basic.Error {
	var reason pk.UnsignedByte
	var value pk.Float

	if err := p.Scan(&reason, &value); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ChangeGameState packet: %w", err)}
	}

	fmt.Println("ChangeGameState", reason, value)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) KeepAlive(c *Client, p pk.Packet) basic.Error {
	var keepAliveID pk.Long

	if err := p.Scan(&keepAliveID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read KeepAlive packet: %w", err)}
	}

	if err := c.Conn.WritePacket(
		pk.Marshal(
			packetid.SPacketKeepAlive,
			keepAliveID,
		),
	); !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("unable to write KeepAlive packet: %w", err)}
	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ChunkData(c *Client, p pk.Packet) basic.Error {
	var (
		ChunkPos level.ChunkPos
		Chunk    level.Chunk
	)

	if err := p.Scan(
		&ChunkPos, &Chunk,
	); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ChunkData packet: %w", err)}
	}

	fmt.Println("ChunkData", ChunkPos, len(Chunk.Sections), len(c.World.Columns))
	c.World.Columns[ChunkPos] = &Chunk

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Effect(c *Client, p pk.Packet) basic.Error {
	var effectID pk.Int
	var location pk.Position
	var data pk.Int
	var disableRelativeVolume pk.Boolean

	if err := p.Scan(&effectID, &location, &data, &disableRelativeVolume); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Effect packet: %w", err)}
	}

	fmt.Println("Effect", effectID, location, data, disableRelativeVolume)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Particle(c *Client, p pk.Packet) basic.Error {
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
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) JoinGame(c *Client, p pk.Packet) basic.Error {
	if err := p.Scan(
		(*pk.Int)(&c.Player.ID),
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
	); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read JoinGame packet: %w", err)}
	}
	if err := c.Conn.WritePacket(pk.Marshal( //PluginMessage packet
		packetid.SPacketPluginMessage,
		pk.Identifier("minecraft:brand"),
		pk.String(c.Player.Settings.Brand),
	)); !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("unable to write PluginMessage packet: %w", err)}
	}

	if err := c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientSettings,
		pk.String(c.Player.Settings.Locale),
		pk.Byte(c.Player.Settings.ViewDistance),
		pk.VarInt(c.Player.Settings.ChatMode),
		pk.Boolean(c.Player.Settings.ChatColors),
		pk.UnsignedByte(c.Player.Settings.DisplayedSkinParts),
		pk.VarInt(c.Player.Settings.MainHand),
		pk.Boolean(c.Player.Settings.EnableTextFiltering),
		pk.Boolean(c.Player.Settings.AllowListing),
	)); !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("unable to write ClientSettings packet: %w", err)}
	}

	c.Player.EntityPlayer = core.NewEntityPlayer(c.Player.GetID(), c.Player.GetUUID(), 116, 0, 0, 0, 0, 0)

	// Add the player to the world
	if err := c.World.AddEntity(c.Player.EntityPlayer); err != nil {
		return basic.Error{Err: basic.InvalidEntity, Info: fmt.Errorf("unable to add player to the world: %w", err)}
	}

	fmt.Println("JoinGame", c.Player.EID, c.Player.Hardcore, c.Player.Gamemode, c.Player.PrevGamemode, c.Player.DimensionNames, c.Player.WorldInfo.DimensionCodec, c.Player.WorldInfo.DimensionType, c.Player.DimensionName, c.Player.HashedSeed, c.Player.MaxPlayers, c.Player.ViewDistance, c.Player.SimulationDistance, c.Player.ReducedDebugInfo, c.Player.EnableRespawnScreen, c.Player.IsDebug, c.Player.IsFlat)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Map(c *Client, p pk.Packet) basic.Error {
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
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Entity(c *Client, p pk.Packet) basic.Error {
	var entityID pk.Int

	if err := p.Scan(&entityID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Entity packet: %w", err)}
	}

	fmt.Println("Entity", entityID)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityPosition(c *Client, p pk.Packet) basic.Error {
	var (
		EntityID               pk.VarInt
		DeltaX, DeltaY, DeltaZ pk.Short
		OnGround               pk.Boolean
	)

	if err := p.Scan(&EntityID, &DeltaX, &DeltaY, &DeltaZ, &OnGround); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityPosition packet: %w", err)}
	}

	if _, entity, err := c.World.GetEntityByID(int32(EntityID)); err == nil {
		if t, ok := entity.(*core.Entity); ok {
			t.AddRelativePosition(maths.Vec3d[float64]{X: float64(DeltaX), Y: float64(DeltaY), Z: float64(DeltaZ)})
		}
	}

	fmt.Println("EntityRelativeMove", EntityID, DeltaX, DeltaY, DeltaZ, OnGround)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityPositionRotation(c *Client, p pk.Packet) basic.Error {
	var (
		EntityID               pk.VarInt
		DeltaX, DeltaY, DeltaZ pk.Short
		Yaw, Pitch             pk.Angle
		OnGround               pk.Boolean
	)

	if err := p.Scan(&EntityID, &DeltaX, &DeltaY, &DeltaZ, &Yaw, &Pitch, &OnGround); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityPositionRotation packet: %w", err)}
	}

	fmt.Println("EntityPositionRotation", EntityID, DeltaX, DeltaY, DeltaZ, Yaw, Pitch, OnGround)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityHeadRotation(c *Client, p pk.Packet) basic.Error {
	var (
		EntityID pk.VarInt
		HeadYaw  pk.Angle
	)

	if err := p.Scan(&EntityID, &HeadYaw); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityHeadRotation packet: %w", err)}
	}

	fmt.Println("EntityHeadRotation", EntityID, HeadYaw)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityRotation(c *Client, p pk.Packet) basic.Error {
	var (
		EntityID   pk.VarInt
		Yaw, Pitch pk.Angle
		OnGround   pk.Boolean
	)

	if err := p.Scan(&EntityID, &Yaw, &Pitch, &OnGround); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityRotation packet: %w", err)}
	}

	fmt.Println("EntityRotation", EntityID, Yaw, Pitch, OnGround)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) VehicleMove(c *Client, p pk.Packet) basic.Error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read VehicleMove packet: %w", err)}
	}

	fmt.Println("VehicleMove", X, Y, Z, Yaw, Pitch)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) OpenSignEditor(c *Client, p pk.Packet) basic.Error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read OpenSignEditor packet: %w", err)}
	}

	fmt.Println("OpenSignEditor", location)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) CraftRecipeResponse(c *Client, p pk.Packet) basic.Error {
	var windowID pk.UnsignedByte
	var recipe pk.String

	if err := p.Scan(&windowID, &recipe); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read CraftRecipeResponse packet: %w", err)}
	}

	fmt.Println("CraftRecipeResponse", windowID, recipe)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) PlayerAbilities(c *Client, p pk.Packet) basic.Error {
	var flags pk.UnsignedByte
	var flyingSpeed pk.Float
	var fov pk.Float

	if err := p.Scan(&flags, &flyingSpeed, &fov); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read PlayerAbilities packet: %w", err)}
	}

	fmt.Println("PlayerAbilities", flags, flyingSpeed, fov)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) CombatEvent(c *Client, p pk.Packet) basic.Error {
	var event pk.Byte
	var duration pk.Int
	var entityID pk.Int
	var message pk.String

	if err := p.Scan(&event, &duration, &entityID, &message); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read CombatEvent packet: %w", err)}
	}

	fmt.Println("CombatEvent", event, duration, entityID, message)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) PlayerInfo(c *Client, p pk.Packet) basic.Error {
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
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SyncPlayerPosition(c *Client, p pk.Packet) basic.Error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
		Flags      pk.Byte
		TeleportID pk.VarInt
		Dismount   pk.Boolean
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch, &Flags, &TeleportID, &Dismount); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SyncPlayerPosition packet: %w", err)}
	}

	position := maths.Vec3d[float64]{X: float64(X), Y: float64(Y), Z: float64(Z)}
	rotation := maths.Vec2d[float64]{X: float64(Pitch), Y: float64(Yaw)}

	if Flags&0x01 != 0 {
		c.Player.Position = c.Player.Position.Add(position)
		c.Player.Rotation = c.Player.Rotation.Add(rotation)
	} else {
		c.Player.Position = position
		c.Player.Rotation = rotation
	}

	c.Player.OnGround = Flags&0x02 == 0

	if TeleportID != 0 {
		if err := c.Conn.WritePacket(
			pk.Marshal(
				packetid.SPacketTeleportConfirm,
				TeleportID,
			),
		); !err.Is(basic.NoError) {
			return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("unable to write TeleportConfirm packet: %w", err)}
		}
	}

	fmt.Println("SyncPlayerPosition", position, rotation, TeleportID, Dismount)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) PlayerPositionAndLook(c *Client, p pk.Packet) basic.Error {
	var (
		x          pk.Double
		y          pk.Double
		z          pk.Double
		yaw        pk.Float
		pitch      pk.Float
		flags      pk.Byte
		teleportID pk.VarInt
	)

	if err := p.Scan(&x, &y, &z, &yaw, &pitch, &flags, &teleportID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read PlayerPositionAndLook packet: %w", err)}
	}

	fmt.Println("PlayerPositonAndLook", x, y, z, yaw, pitch, flags, teleportID)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) UseBed(c *Client, p pk.Packet) basic.Error {
	var entityID pk.Int
	var location pk.Position

	if err := p.Scan(&entityID, &location); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read UseBed packet: %w", err)}
	}

	fmt.Println("UseBed", entityID, location)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) UnlockRecipes(c *Client, p pk.Packet) basic.Error {
	/*var action pk.Byte
	var craftingBookOpen pk.Boolean
	var filter pk.Boolean
	var recipes []pk.String

	if err := p.Scan(&action, &craftingBookOpen, &filter, &recipes); err != nil {
		return err
	}

	fmt.Println("UnlockRecipes", action, craftingBookOpen, filter, recipes)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) DestroyEntities(c *Client, p pk.Packet) basic.Error {
	/*var entityIDs []pk.Int

	if err := p.Scan(&entityIDs); err != nil {
		return err
	}

	fmt.Println("DestroyEntities", entityIDs)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) RemoveEntityEffect(c *Client, p pk.Packet) basic.Error {
	var entityID pk.Int
	var effectID pk.Byte

	if err := p.Scan(&entityID, &effectID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read RemoveEntityEffect packet: %w", err)}
	}

	fmt.Println("RemoveEntityEffect", entityID, effectID)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ResourcePackSend(c *Client, p pk.Packet) basic.Error {
	var (
		url  pk.String
		hash pk.String
	)

	if err := p.Scan(&url, &hash); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ResourcePackSend packet: %w", err)}
	}

	fmt.Println("ResourcePackSend", url, hash)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Respawn(c *Client, p pk.Packet) basic.Error {
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
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Respawn packet: %w", err)}
	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SelectAdvancementTab(c *Client, p pk.Packet) basic.Error {
	var hasID pk.Boolean
	var identifier pk.String

	if err := p.Scan(&hasID, &identifier); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SelectAdvancementTab packet: %w", err)}
	}

	fmt.Println("SelectAdvancementTab", hasID, identifier)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) WorldBorder(c *Client, p pk.Packet) basic.Error {
	var (
		action         pk.Byte
		radius         pk.Double
		oldRadius      pk.Double
		speed          pk.VarInt
		x              pk.Int
		z              pk.Int
		portalBoundary pk.Int
		warningTime    pk.VarInt
		warningBlocks  pk.VarInt
	)

	if err := p.Scan(&action, &radius, &oldRadius, &speed, &x, &z, &portalBoundary, &warningTime, &warningBlocks); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read WorldBorder packet: %w", err)}
	}

	fmt.Println("WorldBorder", action, radius, oldRadius, speed, x, z, portalBoundary, warningTime, warningBlocks)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Camera(c *Client, p pk.Packet) basic.Error {
	var cameraID pk.Int

	if err := p.Scan(&cameraID); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Camera packet: %w", err)}
	}

	fmt.Println("Camera", cameraID)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) HeldItemChange(c *Client, p pk.Packet) basic.Error {
	var slot pk.Short

	if err := p.Scan(&slot); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read HeldItemChange packet: %w", err)}
	}

	newSlot, err := c.Player.Inventory.GetHotbarSlotById(uint8(slot))
	if !err.Is(basic.NoError) {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read HeldItemChange packet: %w", err)}
	}

	fmt.Println("HeldItemChange", slot, "New Item:", item.ByID[item.ID(newSlot.ID)].Name)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) DisplayScoreboard(c *Client, p pk.Packet) basic.Error {
	var position pk.Byte
	var name pk.String

	if err := p.Scan(&position, &name); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read DisplayScoreboard packet: %w", err)}
	}

	fmt.Println("DisplayScoreboard", position, name)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityMetadata(c *Client, p pk.Packet) basic.Error {
	var (
		err      error
		EntityID pk.VarInt
		Metadata struct {
			Index pk.UnsignedByte
			Type  pk.VarInt
			Value interface{}
		}
	)

	if err := p.Scan(
		&EntityID,
		&Metadata.Index,
		pk.Opt{
			If: func() bool {
				return Metadata.Index != 0xff
			},
			Value: &Metadata.Type,
		},
	); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityMetadata packet: %w", err)}
	}
	switch Metadata.Type {
	case 0:
		var Value pk.Byte
		err = p.Scan(&Value)
	case 1:
		var Value pk.VarInt
		err = p.Scan(&Value)
	case 2:
		var Value pk.Float
		err = p.Scan(&Value)
	case 3:
		var Value pk.String
		err = p.Scan(&Value)
	case 4:
		var Value chat.Message
		err = p.Scan(&Value)
	case 5:
		var Value struct {
			Present pk.Boolean
			Value   chat.Message
		}
		err = p.Scan(
			&Value.Present,
			pk.Opt{
				If:    Value.Present,
				Value: &Value.Value,
			},
		)
	case 6:
		var Value Slot
		err = p.Scan(&Value)
	case 7:
		var Value pk.Boolean
		err = p.Scan(&Value)
	case 8:
		var Value pk.Position
		err = p.Scan(&Value)
	case 9:
		var Value pk.Position
		err = p.Scan(&Value)
	case 10:
		var Value struct {
			Present pk.Boolean
			Value   pk.Position
		}
		err = p.Scan(
			&Value.Present,
			pk.Opt{
				If:    Value.Present,
				Value: &Value.Value,
			},
		)
	case 11:
		var Value pk.VarInt
		err = p.Scan(&Value)
	case 12:
		var Value struct {
			Present pk.Boolean
			Value   pk.UUID
		}
		err = p.Scan(
			&Value.Present,
			pk.Opt{
				If:    Value.Present,
				Value: &Value.Value,
			},
		)
	case 13:
		var Value pk.VarInt
		err = p.Scan(&Value)
	/*case 14:*/
	case 15:
		var Value struct {
			ID   pk.VarInt
			Data pk.VarInt // TODO: This is a particle data
		}
		err = p.Scan(&Value.ID, &Value.Data)
	case 16:
		var Value pk.ByteArray // TODO: 3 floats
		err = p.Scan(&Value)
	case 17:
		var Value pk.VarInt
		err = p.Scan(&Value)
	case 18:
		var Value pk.VarInt
		err = p.Scan(&Value)
	case 19:
		var Value pk.VarInt
		err = p.Scan(&Value)
	case 20:
		var Value pk.VarInt
		err = p.Scan(&Value)
	/*case 21:*/
	case 22:
		var Value pk.VarInt
		err = p.Scan(&Value)
	}

	if err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityMetadata packet: %w", err)}
	}

	fmt.Println("EntityMetadata", EntityID, Metadata.Index, Metadata.Type)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) AttachEntity(c *Client, p pk.Packet) basic.Error {
	var (
		entityID  pk.Int
		vehicleID pk.Int
		leash     pk.Boolean
	)

	if err := p.Scan(&entityID, &vehicleID, &leash); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read AttachEntity packet: %w", err)}
	}

	fmt.Println("AttachEntity", entityID, vehicleID, leash)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityVelocity(c *Client, p pk.Packet) basic.Error {
	var (
		entityID                        pk.VarInt
		velocityX, velocityY, velocityZ pk.Short
	)

	if err := p.Scan(&entityID, &velocityX, &velocityY, &velocityZ); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityVelocity packet: %w", err)}
	}

	if _, e, err := c.World.GetEntityByID(int32(entityID)); err == nil {
		switch e.(type) {
		case *core.Entity:
			e.(*core.Entity).SetMotion(maths.Vec3d[float64]{X: float64(velocityX) / 8000, Y: float64(velocityY) / 8000, Z: float64(velocityZ) / 8000}.Spread())
		case *core.EntityLiving:
			e.(*core.EntityLiving).SetMotion(maths.Vec3d[float64]{X: float64(velocityX) / 8000, Y: float64(velocityY) / 8000, Z: float64(velocityZ) / 8000}.Spread())
		case *core.EntityPlayer:
			e.(*core.EntityPlayer).SetMotion(maths.Vec3d[float64]{X: float64(velocityX) / 8000, Y: float64(velocityY) / 8000, Z: float64(velocityZ) / 8000}.Spread())
		}
	} else {
		return basic.Error{Err: basic.InvalidEntity, Info: fmt.Errorf("unable to find entity with ID %d", entityID)}
	}

	fmt.Println("EntityVelocity", entityID, velocityX, velocityY, velocityZ)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityEquipment(c *Client, p pk.Packet) basic.Error {
	/*var entityID pk.Int
	var slot pk.Short
	var item pk.Slot

	if err := p.Scan(&entityID, &slot, &item); err != nil {
		return err
	}

	fmt.Println("EntityEquipment", entityID, slot, item)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SetExperience(c *Client, p pk.Packet) basic.Error {
	var (
		experienceBar   pk.Float
		levelInt        pk.VarInt
		totalExperience pk.VarInt
	)

	if err := p.Scan(&experienceBar, &levelInt, &totalExperience); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SetExperience packet: %w", err)}
	}

	fmt.Println("SetExperience", experienceBar, levelInt, totalExperience)
	c.Player.SetExp(float32(experienceBar), int32(levelInt), int32(totalExperience))
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) UpdateHealth(c *Client, p pk.Packet) basic.Error {
	var (
		health         pk.Float
		food           pk.VarInt
		foodSaturation pk.Float
	)

	if err := p.Scan(&health, &food, &foodSaturation); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read UpdateHealth packet: %w", err)}
	}

	fmt.Println("UpdateHealth", health, food, foodSaturation)
	if respawn := c.Player.SetHealth(float32(health)); respawn {
		if err := c.Player.Respawn(c); err != nil {
			return basic.Error{Err: basic.NoError, Info: err}
		}
	}
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) ScoreboardObjective(c *Client, p pk.Packet) basic.Error {
	var (
		name           pk.String
		mode           pk.Byte
		objectiveName  pk.String
		objectiveValue pk.String
		type_          pk.Byte
	)

	if err := p.Scan(&name, &mode, &objectiveName, &objectiveValue, &type_); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read ScoreboardObjective packet: %w", err)}
	}

	fmt.Println("ScoreboardObjective", name, mode, objectiveName, objectiveValue, type_)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SetPassengers(c *Client, p pk.Packet) basic.Error {
	/*var entityID pk.Int
	var passengerCount pk.VarInt
	var passengers pk.Ary[]pk.VarInt

	if err := p.Scan(&entityID, &passengers); err != nil {
		return err
	}

	fmt.Println("SetPassengers", entityID, passengerCount, passengers)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Teams(c *Client, p pk.Packet) basic.Error {
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
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) UpdateScore(c *Client, p pk.Packet) basic.Error {
	var name pk.String
	var action pk.Byte
	var objectiveName pk.String
	var value pk.VarInt

	if err := p.Scan(&name, &action, &objectiveName, &value); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read UpdateScore packet: %w", err)}
	}

	fmt.Println("UpdateScore", name, action, objectiveName, value)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SpawnPosition(c *Client, p pk.Packet) basic.Error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SpawnPosition packet: %w", err)}
	}

	fmt.Println("SpawnPosition", location)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) TimeUpdate(c *Client, p pk.Packet) basic.Error {
	var (
		WorldAge  pk.Long
		TimeOfDay pk.Long
	)

	c.TPS.Update()

	if err := p.Scan(&WorldAge, &TimeOfDay); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read TimeUpdate packet: %w", err)}
	}
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Title(c *Client, p pk.Packet) basic.Error {
	var (
		action    pk.Byte
		fadeIn    pk.Int
		stay      pk.Int
		fadeOut   pk.Int
		title     pk.String
		subtitle  pk.String
		actionBar pk.String
	)

	if err := p.Scan(&action, &fadeIn, &stay, &fadeOut, &title, &subtitle, &actionBar); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Title packet: %w", err)}
	}

	fmt.Println("Title", action, fadeIn, stay, fadeOut, title, subtitle, actionBar)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) SoundEffect(c *Client, p pk.Packet) basic.Error {
	var (
		soundID        pk.VarInt
		soundCategory  pk.VarInt
		effectPosition pk.Position
		volume         pk.Float
		pitch          pk.Float
	)

	if err := p.Scan(&soundID, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read SoundEffect packet: %w", err)}
	}

	fmt.Println("SoundEffect", soundID, soundCategory, effectPosition, volume, pitch)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) PlayerListHeaderAndFooter(c *Client, p pk.Packet) basic.Error {
	var (
		header pk.String
		footer pk.String
	)

	if err := p.Scan(&header, &footer); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read PlayerListHeaderAndFooter packet: %w", err)}
	}

	fmt.Println("PlayerListHeaderAndFooter", header, footer)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) CollectItem(c *Client, p pk.Packet) basic.Error {
	var (
		collectedEntityID pk.Int
		collectorEntityID pk.Int
		pickupCount       pk.Int
	)

	if err := p.Scan(&collectedEntityID, &collectorEntityID, &pickupCount); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read CollectItem packet: %w", err)}
	}

	fmt.Println("CollectItem", collectedEntityID, collectorEntityID, pickupCount)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityTeleport(c *Client, p pk.Packet) basic.Error {
	var entityID pk.VarInt
	var x pk.Double
	var y pk.Double
	var z pk.Double
	var yaw pk.Byte
	var pitch pk.Byte
	var onGround pk.Boolean

	if err := p.Scan(&entityID, &x, &y, &z, &yaw, &pitch, &onGround); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityTeleport packet: %w", err)}
	}

	fmt.Println("EntityTeleport", entityID, x, y, z, yaw, pitch, onGround)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) Advancements(c *Client, p pk.Packet) basic.Error {
	var action pk.Byte
	var data pk.String

	if err := p.Scan(&action, &data); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read Advancements packet: %w", err)}
	}

	fmt.Println("Advancements", action, data)
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityProperties(c *Client, p pk.Packet) basic.Error {
	/*var entityID pk.Int
	var count pk.VarInt
	var properties pk.Ary[]pk.String

	if err := p.Scan(&entityID, &properties); err != nil {
		return err
	}

	fmt.Println("EntityProperties", entityID, count, properties)*/
	return basic.Error{Err: basic.NoError, Info: nil}
}

func (e *EventsListener) EntityEffect(c *Client, p pk.Packet) basic.Error {
	var (
		entityID   pk.VarInt
		effectID   pk.VarInt
		amplifier  pk.Byte
		duration   pk.VarInt
		flags      pk.Byte
		factorData pk.Boolean
		codec      struct {
			PaddingDuration        int     `nbt:"padding_duration"`
			FactorStart            float32 `nbt:"factor_start"`
			FactorTarget           float32 `nbt:"factor_target"`
			FactorCurrent          float32 `nbt:"factor_current"`
			EffectChangedTimeStamp int     `nbt:"effect_changed_timestamp"`
			FactorPreviousFrame    float32 `nbt:"factor_previous_frame"`
			HadEffectLastTick      bool    `nbt:"had_effect_last_tick"`
		}
	)

	if err := p.Scan(
		pk.Tuple{
			&entityID,
			&effectID,
			&amplifier,
			&duration,
			&flags,
			pk.Opt{
				If:    &factorData,
				Value: pk.NBT(&codec),
			},
		},
	); err != nil {
		return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("unable to read EntityEffect packet: %w", err)}
	}

	if effect, ok := effects.ByID[int32(effectID)]; ok {
		effectStatus := &effects.EffectStatus{
			ID:            int32(effectID),
			Amplifier:     byte(amplifier),
			Duration:      int32(duration),
			ShowParticles: flags&0x01 == 0x01,
			ShowIcon:      flags&0x04 == 0x04,
		}
		c.Player.ActivePotionEffects[effectStatus.ID] = effectStatus
		fmt.Println("EntityEffect", entityID, effect, amplifier, duration, flags, codec)
	}
	return basic.Error{Err: basic.NoError, Info: nil}
}
