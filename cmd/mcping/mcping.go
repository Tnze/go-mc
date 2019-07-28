// Usage: go run cmd/ping/ping.go localhost
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
)

type status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	//favicon ignored
}

func main() {
	addr, port := getAddr()

	fmt.Printf("MCPING (%s:%d):\n", addr, port)

	resp, delay, err := bot.PingAndList(addr, port)
	if err != nil {
		fmt.Printf("ping and list server fail: %v", err)
		os.Exit(1)
	}

	var s status
	err = json.Unmarshal(resp, &s)
	if err != nil {
		fmt.Print("unmarshal resp fail:", err)
		os.Exit(1)
	}

	fmt.Print(s)
	fmt.Println("Delay:", delay)
}

func getAddr() (string, int) {
	const usage = "Usage: mcping <hostname>[:port]"
	if len(os.Args) < 2 {
		fmt.Println("no host name.", usage)
		os.Exit(1)
	}

	addr := strings.Split(os.Args[1], ":")
	var port int
	switch len(addr) {
	case 1:
		port = 25565
	case 2:
		var err error
		port, err = strconv.Atoi(addr[1])
		if err != nil {
			fmt.Println(err, usage)
			os.Exit(1)
		}
	}

	return addr[0], port
}

func (s status) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "Server:", s.Version.Name)
	fmt.Fprintln(&sb, "Protocol:", s.Version.Protocol)
	fmt.Fprintln(&sb, "Description:", s.Description)
	fmt.Fprintf(&sb, "Players: %d/%d\n", s.Players.Online, s.Players.Max)
	for _, v := range s.Players.Sample {
		fmt.Fprintf(&sb, "- [%s] %v\n", v.Name, v.ID)
	}
	return sb.String()
}
