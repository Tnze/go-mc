package main

import (
	"github.com/Tnze/go-mc/bot"
	"log"
)

func main() {
	resp, delay, err := bot.PingAndList("play.miaoscraft.cn", 25565)
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}
	log.Println("Status:"+string(resp), "\nDealy:", delay)
}
