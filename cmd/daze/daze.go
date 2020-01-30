package main

import (
	"bytes"
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	pk "github.com/Tnze/go-mc/net/packet"
)

func main() {
	c := bot.NewClient("Steve")

	//Login
	err := c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//Register event handlers
	c.Events.OnGameBegin = onGameBegin
	c.Events.OnChatMessage = onChatMessage
	c.Events.OnDisconnect = onDisconnect
	c.Events.OnPluginMessage = onPluginMessage

	//JoinGame
	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onGameBegin() error {
	log.Println("Game start")
	return nil //if err isn't nil, HandleGame() will return it.
}

func onChatMessage(c chat.Message, pos byte) error {
	log.Println("Chat:", c.ClearString()) // output chat message without any format code (like color or bold)
	return nil
}

func onDisconnect(c chat.Message) error {
	log.Println("Disconnect:", c)
	return nil
}

func onPluginMessage(channel string, data []byte) error {
	switch channel {
	case "minecraft:brand":
		var brand pk.String
		if err := brand.Decode(bytes.NewReader(data)); err != nil {
			return err
		}
		log.Println("Server brand is:", brand)

	default:
		log.Println("PluginMessage", channel, data)
	}
	return nil
}
