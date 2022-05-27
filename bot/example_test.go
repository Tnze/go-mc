package bot

import (
	"log"

	"encoding/hex"
	"github.com/Tnze/go-mc/offline"
	"github.com/Tnze/go-mc/yggdrasil"
)

func ExamplePingAndList() {
	resp, delay, err := PingAndList("localhost:25565")
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}

	log.Println("Status:", string(resp))
	log.Println("Delay:", delay)
}

func ExampleClient_JoinServer_offline() {
	c := NewClient()
	c.Auth.Name = "Tnze" // set it's name before login.

	id := offline.NameToUUID(c.Auth.Name) // optional, get uuid of offline mode game
	c.Auth.UUID = hex.EncodeToString(id[:])

	//Login
	err := c.JoinServer("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// Register event handlers
	// 	c.Events.GameStart = onGameStartFunc
	// 	c.Events.ChatMsg = onChatMsgFunc
	// 	c.Events.Disconnect = onDisconnectFunc
	//	...

	//JoinGame
	err = c.HandleGame()
	if err != nil {
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
