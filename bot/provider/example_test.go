package provider

import (
	"encoding/hex"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/offline"
	"github.com/Tnze/go-mc/yggdrasil"
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

	c.EventHandlers.Attach(c)
	//Register event handlers
	c.Events.AddListener(
		PacketHandler{ID: packetid.CPacketSetContainerContent, Priority: 0, F: c.EventHandlers.SetWindowContent},
		PacketHandler{ID: packetid.CPacketSpawnEntity, Priority: 0, F: c.EventHandlers.SpawnEntity},
		PacketHandler{ID: packetid.CPacketSpawnPlayer, Priority: 0, F: c.EventHandlers.SpawnPlayer},
		PacketHandler{ID: packetid.CPacketEntityEffect, Priority: 0, F: c.EventHandlers.EntityEffect},
		PacketHandler{ID: packetid.CPacketEntityVelocity, Priority: 0, F: c.EventHandlers.EntityVelocity},
		PacketHandler{ID: packetid.CPacketEntityPositionRotation, Priority: 0, F: c.EventHandlers.EntityPositionRotation},
		PacketHandler{ID: packetid.CPacketPlayerAbilities, Priority: 0, F: c.EventHandlers.PlayerAbilities},
		PacketHandler{ID: packetid.CPacketSyncPosition, Priority: 0, F: c.EventHandlers.SyncPlayerPosition},
		PacketHandler{ID: packetid.CPacketChunkData, Priority: 0, F: c.EventHandlers.ChunkData},
		PacketHandler{ID: packetid.CPacketExplosion, Priority: 0, F: c.EventHandlers.Explosion},
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
