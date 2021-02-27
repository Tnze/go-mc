package main

import (
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	"github.com/mattn/go-colorable"
)

const timeout = 45

var (
	c *bot.Client
	p *basic.Player

	watch chan time.Time
)

func main() {
	log.SetOutput(colorable.NewColorableStdout())
	c = bot.NewClient()
	p = basic.NewPlayer(c, basic.DefaultSettings)

	//Register event handlers
	basic.EventsListener{
		GameStart:  onGameStart,
		ChatMsg:    onChatMsg,
		Disconnect: onDisconnect,
		Death:      onDeath,
	}.Attach(c)
	c.Events.AddListener(bot.PacketHandler{
		ID:       packetid.NamedSoundEffect,
		Priority: 0,
		F: func(p pk.Packet) error {
			return onSound()
		},
	})

	//Login
	err := c.JoinServer("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//JoinGame
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

	return c.UseItem(0)
}

//goland:noinspection SpellCheckingInspection
func onSound(name string, category int, x, y, z float64, volume, pitch float32) error {
	if name == "entity.fishing_bobber.splash" {
		if err := c.UseItem(0); err != nil { //retrieve
			return err
		}
		log.Println("gra~")
		time.Sleep(time.Millisecond * 300)
		if err := c.UseItem(0); err != nil { //throw
			return err
		}
		watch <- time.Now()
	}
	return nil
}

func onChatMsg(c chat.Message, pos byte, uuid uuid.UUID) error {
	log.Println("Chat:", c)
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
			if err := c.UseItem(0); err != nil {
				panic(err)
			}
		}
		to.Reset(time.Second * timeout)
	}
}
