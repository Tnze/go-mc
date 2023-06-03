package provider

import (
	"encoding/hex"
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/internal/utils"
	"github.com/Tnze/go-mc/offline"
	"github.com/Tnze/go-mc/yggdrasil"
	auth "github.com/maxsupermanhd/go-mc-ms-auth"
	"log"
	"testing"
)

func TestExamplePingAndList(t *testing.T) {
	resp, delay, err := PingAndList("")
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}

	log.Println("Status:", string(resp))
	log.Println("Delay:", delay)
}

func TestExampleClient_JoinServer_online(t *testing.T) {
	c := NewClient()

	if mcAuth, err := auth.GetMCcredentials(utils.GetCacheDirectory(), "88650e7e-efee-4857-b9a9-cf580a00ef43"); err != nil {
		fmt.Printf("GetMCcredentials failed: %v\nWe can safely ignore this error if you are connecting on a offline server\n", err)
		c.Auth.Name = "Steve"
		c.Auth.UUID = offline.NameToUUID(c.Auth.Name).String()
	} else {
		c.Auth = Auth(mcAuth)
	}

	//Login
	if err := c.JoinServer("localhost:25565"); !err.Is(basic.NoError) {
		log.Fatal(err)
	}
	log.Println("Login success")

	Attach(c)
	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler{ID: packetid.CPacketSetContainerContent, Priority: 0, F: SetContainerContent},
		PacketHandler{ID: packetid.CPacketSetContainerSlot, Priority: 0, F: SetContainerSlot},
		PacketHandler{ID: packetid.CPacketSetContainerProperty, Priority: 0, F: SetContainerProperty},

		/* Physic */
		PacketHandler{ID: packetid.CPacketChunkData, Priority: 0, F: ChunkData},
		PacketHandler{ID: packetid.CPacketExplosion, Priority: 0, F: Explosion},

		/* Entities */
		PacketHandler{ID: packetid.CPacketSpawnEntity, Priority: 0, F: SpawnEntity},
		PacketHandler{ID: packetid.CPacketSpawnExperienceOrb, Priority: 0, F: SpawnExperienceOrb},
		PacketHandler{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: SpawnPlayer},
		PacketHandler{ID: packetid.CPacketEntityAnimation, Priority: 0, F: EntityAnimation},
		PacketHandler{ID: packetid.CPacketBlockEntityData, Priority: 0, F: BlockEntityData},
		PacketHandler{ID: packetid.CPacketBlockAction, Priority: 0, F: BlockAction},
		PacketHandler{ID: packetid.CPacketBlockUpdate, Priority: 0, F: BlockChange},
		PacketHandler{ID: packetid.CPacketEntityPosition, Priority: 0, F: EntityPosition},
		PacketHandler{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: EntityPositionRotation},
		PacketHandler{ID: packetid.CPacketEntityRotation, Priority: 0, F: EntityRotation},
		PacketHandler{ID: packetid.CPacketVehicleMove, Priority: 0, F: VehicleMove},
		PacketHandler{ID: packetid.CPacketLookAt, Priority: 0, F: LookAt},
		PacketHandler{ID: packetid.CPacketSyncPosition, Priority: 0, F: SyncPlayerPosition},
		PacketHandler{ID: packetid.CPacketEntityEffect, Priority: 0, F: EntityEffect},
		PacketHandler{ID: packetid.CPacketEntityVelocity, Priority: 0, F: EntityVelocity},

		PacketHandler{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); !err.Is(basic.NoError) {
		log.Fatal(err)
	}
}

func TestExampleClient_JoinServer_offline(t *testing.T) {
	c := NewClient()
	c.Auth.Name = "Tnze" // set its name before login.

	id := offline.NameToUUID(c.Auth.Name) // optional, get uuid of offline mode game
	c.Auth.UUID = hex.EncodeToString(id[:])

	//Login
	if err := c.JoinServer("127.0.0.1"); !err.Is(basic.NoError) {
		log.Fatal(err)
	}
	log.Println("Login success")

	Attach(c)
	//Register event handlers
	c.Events.AddListener(
		/* Inventory transactions */
		PacketHandler{ID: packetid.CPacketSetContainerContent, Priority: 0, F: SetContainerContent},
		PacketHandler{ID: packetid.CPacketSetContainerSlot, Priority: 0, F: SetContainerSlot},
		PacketHandler{ID: packetid.CPacketSetContainerProperty, Priority: 0, F: SetContainerProperty},

		/* Physic */
		PacketHandler{ID: packetid.CPacketChunkData, Priority: 0, F: ChunkData},
		PacketHandler{ID: packetid.CPacketExplosion, Priority: 0, F: Explosion},

		/* Entities */
		PacketHandler{ID: packetid.CPacketSpawnEntity, Priority: 0, F: SpawnEntity},
		PacketHandler{ID: packetid.CPacketSpawnExperienceOrb, Priority: 0, F: SpawnExperienceOrb},
		PacketHandler{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: SpawnPlayer},
		PacketHandler{ID: packetid.CPacketEntityAnimation, Priority: 0, F: EntityAnimation},
		PacketHandler{ID: packetid.CPacketBlockEntityData, Priority: 0, F: BlockEntityData},
		PacketHandler{ID: packetid.CPacketBlockAction, Priority: 0, F: BlockAction},
		PacketHandler{ID: packetid.CPacketBlockUpdate, Priority: 0, F: BlockChange},
		PacketHandler{ID: packetid.CPacketEntityPosition, Priority: 0, F: EntityPosition},
		PacketHandler{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: EntityPositionRotation},
		PacketHandler{ID: packetid.CPacketEntityRotation, Priority: 0, F: EntityRotation},
		PacketHandler{ID: packetid.CPacketVehicleMove, Priority: 0, F: VehicleMove},
		PacketHandler{ID: packetid.CPacketLookAt, Priority: 0, F: LookAt},
		PacketHandler{ID: packetid.CPacketSyncPosition, Priority: 0, F: SyncPlayerPosition},
		PacketHandler{ID: packetid.CPacketEntityEffect, Priority: 0, F: EntityEffect},
		PacketHandler{ID: packetid.CPacketEntityVelocity, Priority: 0, F: EntityVelocity},

		PacketHandler{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: PlayerAbilities},
	)

	//JoinGame
	if err := c.HandleGame(); !err.Is(basic.NoError) {
		log.Fatal(err)
	}
}

func ExampleClient_JoinServer_online() {
	c := NewClient()

	// Login Mojang account to get AccessToken
	// To use Microsoft Account, see issue #106
	// https://github.com/Tnze/go-mc/issues/106
	auth, err := yggdrasil.Authenticate("Your E-mail", "Your Password")
	if err != nil {
		panic(err)
	}

	// As long as you set these three fields correctly,
	// the client can connect to the online-mode server
	c.Auth.UUID, c.Auth.Name = auth.SelectedProfile()
	c.Auth.AsTk = auth.AccessToken()

	//Connect server
	err = c.JoinServer("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// Register event handlers
	// 	c.Events.GameStart = onGameStartFunc
	// 	c.Events.ChatMsg = onChatMsgFunc
	// 	c.Events.Disconnect = onDisconnectFunc
	//	...

	//Join the game
	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}
