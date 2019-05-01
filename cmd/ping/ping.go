package main

import (
	bot "github.com/Tnze/gomcbot"
	"log"
)

func main() {
	resp, err := bot.PingAndList("play.miaoscraft.cn", 25565)
	if err != nil {
		log.Fatalf("ping and list server fail: %v", err)
	}
	log.Println("Status:" + resp)
}
