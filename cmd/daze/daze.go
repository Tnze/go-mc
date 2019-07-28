package main

import (
	"bytes"
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	// "github.com/Tnze/go-mc/authenticate"
)

func main() {
	c := bot.NewClient()
	// For online-mode, you need login your Mojang account
	// and load your Name and UUID to client:
	//
	// 	auth, err := authenticate.Authenticate("Your E-mail", "Your Password")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	c.Name, c.Auth.UUID, c.AsTk =  auth.SelectedProfile.Name, auth.SelectedProfile.ID, auth.AccessToken

	//Login
	err := c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//Register event handlers
	c.Events.GameStart = onGameStart
	c.Events.ChatMsg = onChatMsg
	c.Events.Disconnect = onDisconnect
	c.Events.PluginMessage = onPluginMessage

	//JoinGame
	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onGameStart() error {
	log.Println("Game start")
	return nil //if err isn't nil, HandleGame() will return it.
}

func onChatMsg(c chat.Message, pos byte) error {
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
