package server

import (
	"github.com/Tnze/go-mc/chat"
	"image"
	"os"
)

func ExamplePingInfo_standardUsage() {
	// Read server icon
	f, err := os.Open("./server-icon.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	icon, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	// Set server info
	playerList := NewPlayerList(20)
	pingInfo, err := NewPingInfo(playerList, "1.18", 757, chat.Text("A Minecraft Server"), icon)
	if err != nil {
		panic(err)
	}
	// Start listening
	s := Server{
		ListPingHandler: pingInfo,
		LoginHandler:    nil,
		GamePlay:        nil,
	}
	err = s.Listen("0.0.0.0:25565")
	if err != nil {
		return
	}
}
