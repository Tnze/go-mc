package bot

import (
	"github.com/Tnze/go-mc/yggdrasil"
	"log"
)

func ExamplePingAndList() {
	resp, delay, err := PingAndList("localhost", 25565)
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}

	log.Println("Status:", string(resp))
	log.Println("Delay:", delay)
}

func ExampleClient_JoinServer_offline() {
	c := NewClient()
	c.Name = "Tnze" // set it's name before login.

	//Login
	err := c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// Regist event handlers
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

	//Login Mojang account to get AccessToken
	auth, err := yggdrasil.Authenticate("Your E-mail", "Your Password")
	if err != nil {
		panic(err)
	}

	c.Auth.UUID, c.Name = auth.SelectedProfile()
	c.AsTk = auth.AccessToken()

	//Connect server
	err = c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// Regist event handlers
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
