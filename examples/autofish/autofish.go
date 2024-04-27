package main

import (
	"errors"
	"flag"
	"log"
	"time"

	//"github.com/mattn/go-colorable"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/msg"
	"github.com/Tnze/go-mc/bot/playerlist"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/item"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/data/registry/entitytype"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/level/block"
	pk "github.com/Tnze/go-mc/net/packet"
)

const timeout = 120

var (
	address     = flag.String("address", "127.0.0.1:25565", "The server address")
	name        = flag.String("name", "Daze", "The player's name")
	playerID    = flag.String("uuid", "", "The player's UUID")
	accessToken = flag.String("token", "", "AccessToken")
)

var (
	c *bot.Client
	p *basic.Player

	playerList    *playerlist.PlayerList
	chatHandler   *msg.Manager
	worldManager  *world.World
	screenManager *screen.Manager

	watch chan time.Time
)

func main() {
	flag.Parse()
	// log.SetOutput(colorable.NewColorableStdout()) // optional for colorable output
	c = bot.NewClient()
	c.Auth = bot.Auth{
		Name: *name,
		UUID: *playerID,
		AsTk: *accessToken,
	}
	p = basic.NewPlayer(c, basic.DefaultSettings, basic.EventsListener{
		GameStart:  onGameStart,
		Disconnect: onDisconnect,
		Death:      onDeath,
	})
	playerList = playerlist.New(c)
	chatHandler = msg.New(c, p, playerList, msg.EventsHandler{
		SystemChat:        onSystemChat,
		PlayerChatMessage: onPlayerChat,
		DisguisedChat:     onDisguisedChat,
	})
	worldManager = world.NewWorld(c, p, world.EventsListener{
		LoadChunk:   onChunkLoad,
		UnloadChunk: onChunkUnload,
	})
	screenManager = screen.NewManager(c, screen.EventsListener{
		Open:    nil,
		SetSlot: onScreenSlotChange,
		Close:   nil,
	})
	// Register event handlers

	c.Events.AddListener(bobberListeners...)

	// Login
	err := c.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// JoinGame
	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onDeath() error {
	log.Println("Died and Respawned")
	// If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	return p.Respawn()
}

func onGameStart() error {
	log.Println("Game start")

	watch = make(chan time.Time)
	go watchDog()

	return UseItem(0)
}

func onChunkLoad(pos level.ChunkPos) error {
	log.Println("Load chunk:", pos)
	return nil
}

func onChunkUnload(pos level.ChunkPos) error {
	log.Println("Unload chunk:", pos)
	return nil
}

var (
	playerBobberEID int32
	playerBobberPos [3]float64
)

var bobberListeners = []bot.PacketHandler{
	{
		ID: packetid.ClientboundAddEntity,
		F: func(packet pk.Packet) error {
			var (
				eID, eType          pk.VarInt
				eUUID               pk.UUID
				x, y, z             pk.Double
				pitch, yaw, headYaw pk.Angle
				Data                pk.VarInt
				vX, vY, vZ          pk.Short
			)
			if err := packet.Scan(&eID, &eUUID, &eType, &x, &y, &z, &pitch, &yaw, &headYaw, &Data, &vX, &vY, &vZ); err != nil {
				return err
			}
			entityType := entitytype.EntityType(eType)
			if entityType == entitytype.FishingBobber {
				if ownerID := int32(Data); ownerID == p.EID {
					log.Print("My bobber id: ", eID)
					playerBobberEID = int32(eID)
					playerBobberPos = [3]float64{float64(x), float64(y), float64(z)}
					observe()
				}
			}
			return nil
		},
	},
	{
		ID: packetid.ClientboundMoveEntityPos,
		F: func(packet pk.Packet) error {
			var (
				eID     pk.VarInt
				x, y, z pk.Short
			)
			if err := packet.Scan(&eID, &x, &y, &z); err != nil {
				return err
			}
			if int32(eID) == playerBobberEID {
				playerBobberPos[0] += float64(x) / (128 * 32)
				playerBobberPos[1] += float64(y) / (128 * 32)
				playerBobberPos[2] += float64(z) / (128 * 32)
				if err := observe(); err != nil {
					return err
				}
			}
			return nil
		},
	},
	{
		ID: packetid.ClientboundMoveEntityPosRot,
		F: func(packet pk.Packet) error {
			var (
				eID        pk.VarInt
				x, y, z    pk.Short
				yaw, pitch pk.Angle
			)
			if err := packet.Scan(&eID, &x, &y, &z, &yaw, &pitch); err != nil {
				return err
			}
			if int32(eID) == playerBobberEID {
				playerBobberPos[0] += float64(x) / (128 * 32)
				playerBobberPos[1] += float64(y) / (128 * 32)
				playerBobberPos[2] += float64(z) / (128 * 32)
				if err := observe(); err != nil {
					return err
				}
			}
			return nil
		},
	},
	{
		ID: packetid.ClientboundTeleportEntity,
		F: func(packet pk.Packet) error {
			var (
				eID        pk.VarInt
				x, y, z    pk.Double
				yaw, pitch pk.Angle
			)
			if err := packet.Scan(&eID, &x, &y, &z, &yaw, &pitch); err != nil {
				return err
			}
			if int32(eID) == playerBobberEID {
				playerBobberPos[0] += float64(x)
				playerBobberPos[1] += float64(y)
				playerBobberPos[2] += float64(z)
				if err := observe(); err != nil {
					return err
				}
			}
			return nil
		},
	},
}

func onScreenSlotChange(id, index int) error {
	if id == -2 {
		log.Printf("Slot: inventory: %v", screenManager.Inventory.Slots[index])
	} else if id == -1 && index == -1 {
		log.Printf("Slot: cursor: %v", screenManager.Cursor)
	} else {
		container, ok := screenManager.Screens[id]
		if ok {
			// Currently, only inventory container is supported
			switch container := container.(type) {
			case *screen.Inventory:
				slot := container.Slots[index]
				itemInfo := item.ByID[item.ID(slot.ID)]
				log.Printf("Slot: Screen[%d].Slot[%d]: [%v] * %d | NBT: %v", id, index, itemInfo, slot.Count, slot.NBT)
			}
		}
	}
	return nil
}

func observe() error {
	// log.Printf("fishing bobber at %.2f", playerBobberPos)
	// How to judge if the bobber is under water?
	var (
		x = int32(playerBobberPos[0])
		y = int32(playerBobberPos[1] + 1.1)
		z = int32(playerBobberPos[2])
	)
	chunkPos := level.ChunkPos{x >> 4, z >> 4}
	chunk := worldManager.Columns[chunkPos]
	log.Print("chunk pos ", chunkPos)
	if chunk == nil {
		return errors.New("error chunk index")
	}
	_, currentDimType := c.ConfigHandler.(*bot.ConfigData).Registries.DimensionType.Find(p.DimensionType)
	if currentDimType == nil {
		return errors.New("dimension type " + p.DimensionType + " not found")
	}
	sec := (y - currentDimType.MinY) >> 4
	blkIndex := int((y%16)*16*16 + (z%16)*16 + (x % 16))
	if blkIndex >= 16*16*16 {
		log.Panic("error block index", blkIndex)
	}
	if sec < 0 || sec >= int32(len(chunk.Sections)) {
		log.Panic("error section index", sec)
	}
	blkID := chunk.Sections[sec].GetBlock(blkIndex)
	blk := block.StateList[blkID]
	if _, ok := blk.(block.Water); ok {
		return onFish()
	}
	return nil
}

var sequence pk.VarInt

func UseItem(hand int32) error {
	sequence++
	return c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundUseItem,
		pk.VarInt(hand),
		sequence,
	))
}

func onFish() error {
	if err := UseItem(0); err != nil { // retrieve
		return err
	}
	log.Println("gra~")
	time.Sleep(time.Millisecond * 600)
	if err := UseItem(0); err != nil { // throw
		return err
	}
	watch <- time.Now()
	return nil
}

func onSystemChat(c chat.Message, overlay bool) error {
	log.Printf("System Chat: %v, Overlay: %v", c, overlay)
	return nil
}

func onPlayerChat(c chat.Message, _ bool) error {
	log.Println("Player Chat:", c)
	return nil
}

func onDisguisedChat(c chat.Message) error {
	log.Println("Disguised Chat:", c)
	return nil
}

func onDisconnect(c chat.Message) error {
	log.Println("Disconnect:", c)
	return nil
}

func watchDog() {
	to := time.NewTimer(time.Second * timeout)
	for {
		select {
		case <-watch:
		case <-to.C:
			log.Println("rethrow")
			if err := UseItem(0); err != nil {
				panic(err)
			}
		}
		to.Reset(time.Second * timeout)
	}
}
