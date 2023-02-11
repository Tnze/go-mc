// Daze could join an offline-mode server as client.
// Just standing there and do nothing. Automatically reborn after five seconds of death.
package main

import (
	"errors"
	"flag"
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	//_ "github.com/Tnze/go-mc/data/lang/zh-cn"
)

var (
	address     = flag.String("address", "127.0.0.1:25565", "The server address")
	name        = flag.String("name", "Daze", "The player's name")
	playerID    = flag.String("uuid", "", "The player's UUID")
	accessToken = flag.String("token", "", "AccessToken")
)

var (
	client *bot.Client
	player *basic.Player
)

func main() {
	flag.Parse()
	client = bot.NewClient()
	client.Auth = bot.Auth{
		Name: *name,
		UUID: *playerID,
		AsTk: *accessToken,
	}
	player = basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{
		Disconnect: onDisconnect,
	})

	// Login
	err := client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// JoinGame
	for {
		if err = client.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}

		if err2 := new(bot.PacketHandlerError); errors.As(err, err2) {
			if err := new(DisconnectErr); errors.As(err2, err) {
				log.Print("Disconnect, reason: ", err.Reason)
				return
			} else {
				// normal packet handler error, ignore and continue.
				log.Print(err2)
			}
		} else {
			// if the error is not a PacketHandlerError, the connection is broken.
			// stop the program
			log.Fatal(err)
		}
	}
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.ClearString()
}

func onDisconnect(reason chat.Message) error {
	// return an error value so that we can stop main loop
	return DisconnectErr{Reason: reason}
}
