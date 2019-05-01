package main

import (
	bot "github.com/Tnze/gomcbot"
	"log"
)

func main() {
	c := bot.NewClient()

	err := c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}
