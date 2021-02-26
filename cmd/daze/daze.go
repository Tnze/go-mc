package main

import (
	"errors"
	"flag"
	"log"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/zh-cn"
)

var address = flag.String("address", "127.0.0.1", "The server address")
var client *bot.Client
var player *basic.Player

func main() {
	flag.Parse()
	client = bot.NewClient()
	player = basic.NewPlayer(client, basic.DefaultSettings)
	basic.EventsListener{
		GameStart:  onGameStart,
		ChatMsg:    onChatMsg,
		Disconnect: onDisconnect,
		Death:      onDeath,
	}.Attach(client)

	//Login
	err := client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//Register event handlers

	//JoinGame
	for {
		if err = client.HandleGame(); err != nil {
			var interErr *bot.PacketHandlerError
			if errors.As(err, &interErr) {
				log.Print("Internal bugs: ", interErr)
			} else {
				log.Fatal(err)
			}
		} else {
			break
		}
	}
}

func onDeath() error {
	log.Println("Died and Respawned")
	// If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	return player.Respawn()
}

func onGameStart() error {
	log.Println("Game start")
	return nil //if err isn't nil, HandleGame() will return it.
}

func onChatMsg(c chat.Message, pos byte, uuid uuid.UUID) error {
	log.Println("Chat:", c.ClearString()) // output chat message without any format code (like color or bold)
	return nil
}

func onDisconnect(reason chat.Message) error {
	log.Println("Disconnect:", reason)
	return nil
}
