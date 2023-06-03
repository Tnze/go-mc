package provider

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/effects"
	"github.com/Tnze/go-mc/level"
	"runtime"
	"time"
	"unsafe"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	. "github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

type EventsListener struct{}

func Attach(c *Client) {
	c.Events.AddListener(
		PacketHandler{Priority: 64, ID: packetid.CPacketLogin, F: JoinGame},
		PacketHandler{Priority: 64, ID: packetid.CPacketKeepAlive, F: KeepAlive},
		PacketHandler{Priority: 64, ID: packetid.CPacketChatMessage, F: ChatMessage},
		PacketHandler{Priority: 64, ID: packetid.CPacketSystemMessage, F: ChatMessage},
		PacketHandler{Priority: 64, ID: packetid.CPacketDisconnect, F: Disconnect},
		PacketHandler{Priority: 64, ID: packetid.CPacketUpdateHealth, F: UpdateHealth},
		PacketHandler{Priority: int(^uint(0) >> 1), ID: packetid.CPacketSetTime, F: TimeUpdate},
	)

	c.Events.AddTicker(
		TickHandler{Priority: int(^uint(0) >> 1), F: ApplyPhysics},
		TickHandler{Priority: int(^uint(0) >> 1), F: RunTransactions},
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

func SpawnEntity(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read SpawnEntity packet: %w", err)
	}

	if err := c.World.AddEntity(core.NewEntity(
		int32(EntityID),
		uuid.UUID(EntityUUID),
		int32(TypeID),
		float64(X), float64(Y), float64(Z),
		float64(Pitch), float64(Yaw),
	)); err != nil {
		return err
	}

	fmt.Println("SpawnEntity", EntityID, EntityUUID, TypeID, X, Y, Z, Pitch, Yaw, HeadYaw, Data, vX, vY, vZ)
	return nil
}

func SpawnExperienceOrb(c *Client, p pk.Packet) error {
	var (
		entityID pk.VarInt
		x, y, z  pk.Double
		count    pk.Short
	)

	if err := p.Scan(&entityID, &x, &y, &z, &count); err != nil {
		return fmt.Errorf("unable to read SpawnExperienceOrb packet: %w", err)
	}

	return nil
}

func SpawnPlayer(c *Client, p pk.Packet) error {
	var (
		EntityID   pk.VarInt
		PlayerUUID pk.UUID
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Angle
	)

	if err := p.Scan(&EntityID, &PlayerUUID, &X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return fmt.Errorf("unable to read SpawnPlayer packet: %w", err)
	}

	if err := c.World.AddEntity(core.NewEntity(
		int32(EntityID),
		uuid.UUID(PlayerUUID),
		116, // Player type
		float64(X), float64(Y), float64(Z),
		float64(Pitch), float64(Yaw),
	)); err != nil {
		return err
	}

	fmt.Println("SpawnPlayer", EntityID, PlayerUUID.String(), X, Y, Z, Yaw, Pitch)
	return nil
}

func EntityAnimation(c *Client, p pk.Packet) error {
	var (
		entityID  pk.VarInt
		animation pk.Byte
	)

	if err := p.Scan(&entityID, &animation); err != nil {
		return fmt.Errorf("unable to read Animation packet: %w", err)
	}

	return nil
}

func AwardStatistics(c *Client, p pk.Packet) error {
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

func SetBlockDestroyStage(c *Client, p pk.Packet) error {
	var (
		entityID pk.VarInt
		location pk.Position
		stage    pk.Byte
	)

	if err := p.Scan(&entityID, &location, &stage); err != nil {
		return fmt.Errorf("unable to read BlockBreakAnimation packet: %w", err)
	}

	return nil
}

func BlockEntityData(c *Client, p pk.Packet) error {
	/*var location pk.Position
	var action pk.Byte
	var nbtData pk

	if err := p.Scan(&location, &action, &nbtData); err != nil {
		return err
	}

	fmt.Println("UpdateBlockEntity", location, action, nbtData)*/
	return nil
}

func BlockAction(c *Client, p pk.Packet) error {
	var (
		location    pk.Position
		actionID    pk.Byte
		actionParam pk.Byte
		blockType   pk.VarInt
	)

	if err := p.Scan(&location, &actionID, &actionParam, &blockType); err != nil {
		return fmt.Errorf("unable to read BlockAction packet: %w", err)
	}

	return nil
}

func BlockChange(c *Client, p pk.Packet) error {
	var (
		location  pk.Position
		blockType pk.VarInt
	)

	if err := p.Scan(&location, &blockType); err != nil {
		return fmt.Errorf("unable to read BlockChange packet: %w", err)
	}

	return nil
}

func BossBar(c *Client, p pk.Packet) error {
	var uuid pk.UUID
	var action pk.Byte

	var err error

	if err = p.Scan(&uuid, &action); err != nil {
		return fmt.Errorf("unable to read BossBar packet: %w", err)
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

		if err = p.Scan(&title, &health, &color, &divisions, &flags); err != nil {
			return fmt.Errorf("unable to read BossBar packet: %w", err)
		}

	case 1:
		var health pk.Float

		err = p.Scan(&health)
	case 2:
		var title pk.String

		err = p.Scan(&title)
	case 3:
		var color pk.Byte

		err = p.Scan(&color)
	case 4:
		var division pk.Byte

		err = p.Scan(&division)
	case 5:
		var flags pk.Byte

		err = p.Scan(&flags)
	case 6:
		//fmt.Println("BossBar", uuid, action)
	}

	return nil
}

func ServerDifficulty(c *Client, p pk.Packet) (err error) {
	var difficulty pk.Byte

	return p.Scan(&difficulty)
}

func TabComplete(c *Client, p pk.Packet) error {
	/*var count pk.VarInt
	var matches []pk.String

	if err := p.Scan(&count, &matches); err != nil {
		return err
	}

	fmt.Println("TabComplete", count, matches)*/
	return nil
}

func ChatMessage(c *Client, p pk.Packet) error {
	var message chat.Message

	if err := p.Scan(&message); err != nil {
		return fmt.Errorf("unable to read ChatMessage packet: %w", err)
	}
	/*
		// Get 2 random items from the inventory
		var (
			item1      *Slot
			item2      *Slot
			item1Index int32
			item2Index int32
		)
		for i, v := range c.Player.Inventory.Slots {
			if v.ID != 0 {
				if item1 == nil {
					item1 = &v
					item1Index = int32(i)
				} else if item2 == nil {
					item2 = &v
					item2Index = int32(i)
					break
				}
			}
		}

		if item1 == nil || item2 == nil {
			return nil
		}
		// Create a new inventory transaction
		builder := transactions.TransactionBuilder{
			WindowID: 0,
			StateID:  pk.VarInt(c.Player.StateID),
			Actions:  transactions.SwitchSlot(item1Index, item1, item2Index, item2),
		}
		c.Player.Transactions.Post(builder.Build())*/

	fmt.Println("ChatMessage", message)

	return nil
}

func MultiBlockChange(c *Client, p pk.Packet) error {
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

func SetContainerContent(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("failed to scan SetContainerContent: %w", err)
	}

	c.Player.Manager.StateID = int32(StateID)
	var container screen.Container
	if ContainerID == 0 {
		container = c.Player.Inventory
	} else {
		container, _ = c.Player.Manager.Screens[int(ContainerID)] // Let's assume that the container exists
	}

	// copy the slot data to container
	container.ApplyData(SlotData)

	fmt.Println("SetContainerContent", ContainerID, StateID, SlotData, CarriedItem)
	return nil
}

func CloseContainer(c *Client, p pk.Packet) error {
	var windowID pk.Byte

	if err := p.Scan(&windowID); err != nil {
		return fmt.Errorf("unable to read CloseContainer packet: %w", err)
	}

	return nil
}

func CloseWindow(c *Client, p pk.Packet) error {
	var windowID pk.Byte

	if err := p.Scan(&windowID); err != nil {
		return fmt.Errorf("unable to read CloseWindow packet: %w", err)
	}

	return nil
}

func SetContainerProperty(c *Client, p pk.Packet) error {
	var (
		windowID pk.Byte
		property pk.Short
		value    pk.Short
	)

	if err := p.Scan(&windowID, &property, &value); err != nil {
		return fmt.Errorf("unable to read WindowProperty packet: %w", err)
	}

	return nil
}

func SetContainerSlot(c *Client, p pk.Packet) (err error) {
	var (
		ContainerID pk.Byte
		StateID     pk.VarInt
		SlotID      pk.Short
		SlotData    Slot
	)
	if err = p.Scan(&ContainerID, &StateID, &SlotID, &SlotData); err != nil {
		return fmt.Errorf("failed to scan SetSlot: %w", err)
	}

	c.Player.Manager.StateID = int32(StateID)
	switch ContainerID {
	case -1:
		c.Player.Manager.Cursor = &SlotData
	case -2:
		err = c.Player.Manager.Inventory.SetSlot(int(SlotID), SlotData)
	default:
		if container, ok := c.Player.Manager.Screens[int(ContainerID)]; !ok {
			return fmt.Errorf("failed to find container with id %d", ContainerID)
		} else {
			err = container.SetSlot(int(SlotID), SlotData)
		}
	}

	return
}

func SetCooldown(c *Client, p pk.Packet) error {
	var (
		itemID pk.VarInt
		ticks  pk.VarInt
	)

	if err := p.Scan(&itemID, &ticks); err != nil {
		return fmt.Errorf("unable to read SetCooldown packet: %w", err)
	}

	return nil
}

func PluginMessage(c *Client, p pk.Packet) error {
	var (
		channel pk.String
		data    pk.ByteArray
	)

	if err := p.Scan(&channel, &data); err != nil {
		return fmt.Errorf("unable to read PluginMessage packet: %w", err)
	}

	return nil
}

func NamedSoundEffect(c *Client, p pk.Packet) error {
	var (
		soundName      pk.String
		soundCategory  pk.Byte
		effectPosition pk.Position
		volume         pk.Float
		pitch          pk.Float
	)

	if err := p.Scan(&soundName, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return fmt.Errorf("unable to read NamedSoundEffect packet: %w", err)
	}

	return nil
}

func Disconnect(c *Client, p pk.Packet) error {
	var reason chat.Message
	if err := p.Scan(&reason); err != nil {
		return fmt.Errorf("failed to scan Disconnect: %w", err)
	}

	fmt.Println("Disconnect:", reason)
	return nil
}

func EntityStatus(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var entityStatus pk.Byte

	if err := p.Scan(&entityID, &entityStatus); err != nil {
		return fmt.Errorf("unable to read EntityStatus packet: %w", err)
	}

	return nil
}

func Explosion(c *Client, p pk.Packet) error {
	var (
		x, y, z    pk.Float
		radius     pk.Float
		records    pk.VarInt
		data       = make([][3]pk.VarInt, 0)
		mX, mY, mZ pk.Float
	)

	if err := p.Scan(&x, &y, &z, &radius, &records); err != nil {
		return fmt.Errorf("unable to read Explosion packet: %w", err)
	}

	data = make([][3]pk.VarInt, records)
	for i := pk.VarInt(0); i < records; i++ {
		if err := p.Scan(&data[i][0], &data[i][1], &data[i][2]); err != nil {
			return fmt.Errorf("unable to read Explosion packet: %w", err)
		}
	}

	if err := p.Scan(&mX, &mY, &mZ); err != nil {
		return fmt.Errorf("unable to read Explosion packet: %w", err)
	}

	fmt.Println("Explosion", x, y, z, radius, records, data, mX, mY, mZ)
	return nil
}

func UnloadChunk(c *Client, p pk.Packet) error {
	var chunk level.ChunkPos

	if err := p.Scan(&chunk); err != nil {
		return fmt.Errorf("unable to read UnloadChunk packet: %w", err)
	}

	delete(c.World.Columns, chunk)
	return nil
}

func ChangeGameState(c *Client, p pk.Packet) error {
	var reason pk.UnsignedByte
	var value pk.Float

	if err := p.Scan(&reason, &value); err != nil {
		return fmt.Errorf("unable to read ChangeGameState packet: %w", err)
	}

	return nil
}

func KeepAlive(c *Client, p pk.Packet) error {
	var keepAliveID pk.Long

	if err := p.Scan(&keepAliveID); err != nil {
		return fmt.Errorf("unable to read KeepAlive packet: %w", err)
	}

	if err := c.Conn.WritePacket(
		pk.Marshal(
			packetid.SPacketKeepAlive,
			keepAliveID,
		),
	); err != nil {
		return fmt.Errorf("unable to write KeepAlive packet: %w", err)
	}

	// DEV
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Alloc = ", m.Alloc/1024/1024, "MiB", "\tTotalAlloc = ", m.TotalAlloc/1024/1024, "MiB", "\tSys = ", m.Sys/1024/1024, "MiB", "\tNumGC = ", m.NumGC)

	return nil
}

func ChunkData(c *Client, p pk.Packet) error {
	var (
		ChunkPos level.ChunkPos
		Chunk    level.Chunk
	)

	if err := p.Scan(
		&ChunkPos, &Chunk,
	); err != nil {
		return fmt.Errorf("unable to read ChunkData packet: %w", err)
	}

	c.World.Columns[ChunkPos] = &Chunk

	return nil
}

func Effect(c *Client, p pk.Packet) error {
	var effectID pk.Int
	var location pk.Position
	var data pk.Int
	var disableRelativeVolume pk.Boolean

	if err := p.Scan(&effectID, &location, &data, &disableRelativeVolume); err != nil {
		return fmt.Errorf("unable to read Effect packet: %w", err)
	}

	return nil
}

func Particle(c *Client, p pk.Packet) error {
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

func JoinGame(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read JoinGame packet: %w", err)
	}
	if err := c.Conn.WritePacket(pk.Marshal( //PluginMessage packet
		packetid.SPacketPluginMessage,
		pk.Identifier("minecraft:brand"),
		pk.String(c.Player.Settings.Brand),
	)); err != nil {
		return fmt.Errorf("unable to write PluginMessage packet: %w", err)
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
	)); err != nil {
		return fmt.Errorf("unable to write ClientSettings packet: %w", err)
	}

	c.Player.EntityPlayer = core.NewEntityPlayer(c.Player.GetID(), c.Player.GetUUID(), 116, 0, 0, 0, 0, 0)

	// Add the player to the world
	if err := c.World.AddEntity(c.Player.EntityPlayer); err != nil {
		return fmt.Errorf("unable to add player to the world: %w", err)
	}

	return nil
}

func Map(c *Client, p pk.Packet) error {
	var Map world.Map

	if err := p.Scan(&Map); err != nil {
		return fmt.Errorf("unable to read Map packet: %w", err)
	}

	return nil
}

func Entity(c *Client, p pk.Packet) error {
	var entityID pk.Int

	if err := p.Scan(&entityID); err != nil {
		return fmt.Errorf("unable to read Entity packet: %w", err)
	}

	return nil
}

func EntityPosition(c *Client, p pk.Packet) error {
	var (
		EntityID               pk.VarInt
		DeltaX, DeltaY, DeltaZ pk.Short
		OnGround               pk.Boolean
	)

	if err := p.Scan(&EntityID, &DeltaX, &DeltaY, &DeltaZ, &OnGround); err != nil {
		return fmt.Errorf("unable to read EntityPosition packet: %w", err)
	}

	if _, entity, err := c.World.GetEntityByID(int32(EntityID)); err == nil {
		if t, ok := entity.(*core.Entity); ok {
			t.AddRelativePosition(maths.Vec3d[float64]{X: float64(DeltaX), Y: float64(DeltaY), Z: float64(DeltaZ)})
		}
	}

	return nil
}

func EntityPositionRotation(c *Client, p pk.Packet) error {
	var (
		EntityID               pk.VarInt
		DeltaX, DeltaY, DeltaZ pk.Short
		Yaw, Pitch             pk.Angle
		OnGround               pk.Boolean
	)

	if err := p.Scan(&EntityID, &DeltaX, &DeltaY, &DeltaZ, &Yaw, &Pitch, &OnGround); err != nil {
		return fmt.Errorf("unable to read EntityPositionRotation packet: %w", err)
	}

	return nil
}

func EntityHeadRotation(c *Client, p pk.Packet) error {
	var (
		EntityID pk.VarInt
		HeadYaw  pk.Angle
	)

	if err := p.Scan(&EntityID, &HeadYaw); err != nil {
		return fmt.Errorf("unable to read EntityHeadRotation packet: %w", err)
	}

	return nil
}

func EntityRotation(c *Client, p pk.Packet) error {
	var (
		EntityID   pk.VarInt
		Yaw, Pitch pk.Angle
		OnGround   pk.Boolean
	)

	if err := p.Scan(&EntityID, &Yaw, &Pitch, &OnGround); err != nil {
		return fmt.Errorf("unable to read EntityRotation packet: %w", err)
	}

	return nil
}

func VehicleMove(c *Client, p pk.Packet) error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch); err != nil {
		return fmt.Errorf("unable to read VehicleMove packet: %w", err)
	}

	return nil
}

func OpenSignEditor(c *Client, p pk.Packet) error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return fmt.Errorf("unable to read OpenSignEditor packet: %w", err)
	}

	return nil
}

func CraftRecipeResponse(c *Client, p pk.Packet) error {
	var windowID pk.UnsignedByte
	var recipe pk.String

	if err := p.Scan(&windowID, &recipe); err != nil {
		return fmt.Errorf("unable to read CraftRecipeResponse packet: %w", err)
	}

	return nil
}

func PlayerAbilities(c *Client, p pk.Packet) error {
	var flags pk.UnsignedByte
	var flyingSpeed pk.Float
	var fov pk.Float

	if err := p.Scan(&flags, &flyingSpeed, &fov); err != nil {
		return fmt.Errorf("unable to read PlayerAbilities packet: %w", err)
	}

	return nil
}

func CombatEvent(c *Client, p pk.Packet) error {
	var event pk.Byte
	var duration pk.Int
	var entityID pk.Int
	var message pk.String

	if err := p.Scan(&event, &duration, &entityID, &message); err != nil {
		return fmt.Errorf("unable to read CombatEvent packet: %w", err)
	}

	fmt.Println("CombatEvent", event, duration, entityID, message)
	return nil
}

func PlayerInfo(c *Client, p pk.Packet) error {
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

func SyncPlayerPosition(c *Client, p pk.Packet) error {
	var (
		X, Y, Z    pk.Double
		Yaw, Pitch pk.Float
		Flags      pk.Byte
		TeleportID pk.VarInt
		Dismount   pk.Boolean
	)

	if err := p.Scan(&X, &Y, &Z, &Yaw, &Pitch, &Flags, &TeleportID, &Dismount); err != nil {
		return fmt.Errorf("unable to read SyncPlayerPosition packet: %w", err)
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
		); err != nil {
			return fmt.Errorf("unable to write TeleportConfirm packet: %w", err)
		}
	}

	return nil
}

func PlayerPositionAndLook(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read PlayerPositionAndLook packet: %w", err)
	}

	return nil
}

func UseBed(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var location pk.Position

	if err := p.Scan(&entityID, &location); err != nil {
		return fmt.Errorf("unable to read UseBed packet: %w", err)
	}

	return nil
}

func UnlockRecipes(c *Client, p pk.Packet) error {
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

func DestroyEntities(c *Client, p pk.Packet) error {
	/*var entityIDs []pk.Int

	if err := p.Scan(&entityIDs); err != nil {
		return err
	}

	fmt.Println("DestroyEntities", entityIDs)*/
	return nil
}

func RemoveEntityEffect(c *Client, p pk.Packet) error {
	var entityID pk.Int
	var effectID pk.Byte

	if err := p.Scan(&entityID, &effectID); err != nil {
		return fmt.Errorf("unable to read RemoveEntityEffect packet: %w", err)
	}

	return nil
}

func ResourcePackSend(c *Client, p pk.Packet) error {
	var (
		url  pk.String
		hash pk.String
	)

	if err := p.Scan(&url, &hash); err != nil {
		return fmt.Errorf("unable to read ResourcePackSend packet: %w", err)
	}

	return nil
}

func Respawn(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read Respawn packet: %w", err)
	}

	return nil
}

func SelectAdvancementTab(c *Client, p pk.Packet) error {
	var hasID pk.Boolean
	var identifier pk.String

	if err := p.Scan(&hasID, &identifier); err != nil {
		return fmt.Errorf("unable to read SelectAdvancementTab packet: %w", err)
	}

	return nil
}

func WorldBorder(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read WorldBorder packet: %w", err)
	}

	return nil
}

func Camera(c *Client, p pk.Packet) error {
	var cameraID pk.Int

	if err := p.Scan(&cameraID); err != nil {
		return fmt.Errorf("unable to read Camera packet: %w", err)
	}

	return nil
}

func SetHeldItem(c *Client, p pk.Packet) error {
	var slot pk.Short

	if err := p.Scan(&slot); err != nil {
		return fmt.Errorf("unable to read SetHeldItem packet: %w", err)
	}

	c.Player.Manager.HeldItem = c.Player.Manager.Inventory.GetHotbarSlot(int(slot))

	return nil
}

func DisplayScoreboard(c *Client, p pk.Packet) error {
	var position pk.Byte
	var name pk.String

	if err := p.Scan(&position, &name); err != nil {
		return fmt.Errorf("unable to read DisplayScoreboard packet: %w", err)
	}

	return nil
}

func EntityMetadata(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read EntityMetadata packet: %w", err)
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
		return fmt.Errorf("unable to read EntityMetadata packet: %w", err)
	}

	return nil
}

func AttachEntity(c *Client, p pk.Packet) error {
	var (
		entityID  pk.Int
		vehicleID pk.Int
		leash     pk.Boolean
	)

	if err := p.Scan(&entityID, &vehicleID, &leash); err != nil {
		return fmt.Errorf("unable to read AttachEntity packet: %w", err)
	}

	return nil
}

func EntityVelocity(c *Client, p pk.Packet) error {
	var (
		entityID                        pk.VarInt
		velocityX, velocityY, velocityZ pk.Short
	)

	if err := p.Scan(&entityID, &velocityX, &velocityY, &velocityZ); err != nil {
		return fmt.Errorf("unable to read EntityVelocity packet: %w", err)
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
		return nil
	}

	return nil
}

func EntityEquipment(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var slot pk.Short
	var item pk.Slot

	if err := p.Scan(&entityID, &slot, &item); err != nil {
		return err
	}

	fmt.Println("EntityEquipment", entityID, slot, item)*/
	return nil
}

func SetExperience(c *Client, p pk.Packet) error {
	var (
		experienceBar   pk.Float
		levelInt        pk.VarInt
		totalExperience pk.VarInt
	)

	if err := p.Scan(&experienceBar, &levelInt, &totalExperience); err != nil {
		return fmt.Errorf("unable to read SetExperience packet: %w", err)
	}

	c.Player.SetExp(float32(experienceBar), int32(levelInt), int32(totalExperience))
	return nil
}

func UpdateHealth(c *Client, p pk.Packet) error {
	var (
		health         pk.Float
		food           pk.VarInt
		foodSaturation pk.Float
	)

	if err := p.Scan(&health, &food, &foodSaturation); err != nil {
		return fmt.Errorf("unable to read UpdateHealth packet: %w", err)
	}

	if respawn := c.Player.SetHealth(float32(health)); respawn {
		if err := c.Player.Respawn(c); err != nil {
			return nil
		}
	}
	return nil
}

func ScoreboardObjective(c *Client, p pk.Packet) error {
	var (
		name           pk.String
		mode           pk.Byte
		objectiveName  pk.String
		objectiveValue pk.String
		type_          pk.Byte
	)

	if err := p.Scan(&name, &mode, &objectiveName, &objectiveValue, &type_); err != nil {
		return fmt.Errorf("unable to read ScoreboardObjective packet: %w", err)
	}

	return nil
}

func SetPassengers(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var passengerCount pk.VarInt
	var passengers pk.Ary[]pk.VarInt

	if err := p.Scan(&entityID, &passengers); err != nil {
		return err
	}*/

	return nil
}

func Teams(c *Client, p pk.Packet) error {
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
	}*/

	return nil
}

func UpdateScore(c *Client, p pk.Packet) error {
	var name pk.String
	var action pk.Byte
	var objectiveName pk.String
	var value pk.VarInt

	if err := p.Scan(&name, &action, &objectiveName, &value); err != nil {
		return fmt.Errorf("unable to read UpdateScore packet: %w", err)
	}

	return nil
}

func SpawnPosition(c *Client, p pk.Packet) error {
	var location pk.Position

	if err := p.Scan(&location); err != nil {
		return fmt.Errorf("unable to read SpawnPosition packet: %w", err)
	}

	return nil
}

func TimeUpdate(c *Client, p pk.Packet) error {
	var (
		WorldAge  pk.Long
		TimeOfDay pk.Long
	)

	c.TPS.Update()

	if err := p.Scan(&WorldAge, &TimeOfDay); err != nil {
		return fmt.Errorf("unable to read TimeUpdate packet: %w", err)
	}
	return nil
}

func Title(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read Title packet: %w", err)
	}

	return nil
}

func SoundEffect(c *Client, p pk.Packet) error {
	var (
		soundID        pk.VarInt
		soundCategory  pk.VarInt
		effectPosition pk.Position
		volume         pk.Float
		pitch          pk.Float
	)

	if err := p.Scan(&soundID, &soundCategory, &effectPosition, &volume, &pitch); err != nil {
		return fmt.Errorf("unable to read SoundEffect packet: %w", err)
	}

	return nil
}

func PlayerListHeaderAndFooter(c *Client, p pk.Packet) error {
	var (
		header pk.String
		footer pk.String
	)

	if err := p.Scan(&header, &footer); err != nil {
		return fmt.Errorf("unable to read PlayerListHeaderAndFooter packet: %w", err)
	}

	return nil
}

func CollectItem(c *Client, p pk.Packet) error {
	var (
		collectedEntityID pk.Int
		collectorEntityID pk.Int
		pickupCount       pk.Int
	)

	if err := p.Scan(&collectedEntityID, &collectorEntityID, &pickupCount); err != nil {
		return fmt.Errorf("unable to read CollectItem packet: %w", err)
	}

	return nil
}

func EntityTeleport(c *Client, p pk.Packet) error {
	var entityID pk.VarInt
	var x pk.Double
	var y pk.Double
	var z pk.Double
	var yaw pk.Byte
	var pitch pk.Byte
	var onGround pk.Boolean

	if err := p.Scan(&entityID, &x, &y, &z, &yaw, &pitch, &onGround); err != nil {
		return fmt.Errorf("unable to read EntityTeleport packet: %w", err)
	}

	return nil
}

func Advancements(c *Client, p pk.Packet) error {
	var action pk.Byte
	var data pk.String

	if err := p.Scan(&action, &data); err != nil {
		return fmt.Errorf("unable to read Advancements packet: %w", err)
	}

	return nil
}

func EntityProperties(c *Client, p pk.Packet) error {
	/*var entityID pk.Int
	var count pk.VarInt
	var properties pk.Ary[]pk.String

	if err := p.Scan(&entityID, &properties); err != nil {
		return err
	}

	fmt.Println("EntityProperties", entityID, count, properties)*/
	return nil
}

func EntityEffect(c *Client, p pk.Packet) error {
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
		return fmt.Errorf("unable to read EntityEffect packet: %w", err)
	}

	if _, ok := effects.ByID[int32(effectID)]; ok {
		effectStatus := &effects.EffectStatus{
			ID:            int32(effectID),
			Amplifier:     byte(amplifier),
			Duration:      int32(duration),
			ShowParticles: flags&0x01 == 0x01,
			ShowIcon:      flags&0x04 == 0x04,
		}
		c.Player.ActivePotionEffects[effectStatus.ID] = effectStatus
	}
	return nil
}

func LookAt(client *Client, packet pk.Packet) error {
	var (
		targetEnum   pk.VarInt
		X, Y, Z      pk.Double
		isEntity     pk.Boolean
		entityID     pk.VarInt
		entityTarget pk.VarInt
	)

	if err := packet.Scan(
		pk.Tuple{
			&targetEnum,
			&X, &Y, &Z,
			pk.Opt{
				If: &isEntity,
				Value: pk.Tuple{
					&entityID,
					&entityTarget,
				},
			},
		},
	); err != nil {
		return fmt.Errorf("unable to read LookAt packet: %w", err)
	}

	return nil
}
