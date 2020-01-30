package main

import (
	"log"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	//"github.com/mattn/go-colorable" // this package is nice but cannot get in china mainland because it import golang.org/x/sys
)

const timeout = 45

var (
	c     *bot.Client
	watch chan time.Time
)

func main() {
	//log.SetOutput(colorable.NewColorableStdout())
	c = bot.NewClient("Steve")

	//Login
	err := c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//Register event handlers
	c.Events.OnGameBegin = onGameStart
	c.Events.OnChatMessage = onChatMsg
	c.Events.OnDisconnect = onDisconnect
	c.Events.OnSound = onSound
	c.Events.OnDeath = onDeath
	c.Events.OnRespawn = onRespawn

	//JoinGame
	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onRespawn() error {
	log.Println("Respawned")
	return nil
}

func onDeath() error {
	log.Println("Died")
	c.Respawn() // If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	return nil
}

func onGameStart() error {
	log.Println("Game start")

	watch = make(chan time.Time)
	go watchDog()

	return c.UseItem(0)
}

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

func onChatMsg(m chat.Message, pos byte) error {
	log.Println("Chat:", m)
	return nil
}

func onDisconnect(m chat.Message) error {
	log.Println("Disconnect:", m)
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
