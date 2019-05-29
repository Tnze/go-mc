package main

import (
	"log"
	"os"

	"github.com/Tnze/go-mc/bot"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no host name. Useage: ping [hostname]")
	}

	resp, delay, err := bot.PingAndList(os.Args[1], 25565)
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}
	log.Println("Status:"+string(resp), "\nDealy:", delay)
}
