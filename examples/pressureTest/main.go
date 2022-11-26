package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	//"github.com/mattn/go-colorable"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
)

var address = flag.String("address", "127.0.0.1", "The server address")
var number = flag.Int("number", 1023, "The number of clients")

func main() {
	flag.Parse()
	//log.SetOutput(colorable.NewColorableStdout())

	for i := 0; i < *number; i++ {
		go func(i int) {
			for {
				ind := newIndividual(i, "Player"+strconv.Itoa(i))
				ind.run(*address)
				time.Sleep(time.Second * 3)
			}
		}(i)
		time.Sleep(time.Millisecond)
	}
	select {}
}

type individual struct {
	id     int
	client *bot.Client
	player *bot.Player
}

func newIndividual(id int, name string) (i *individual) {
	i = new(individual)
	i.id = id
	i.client = bot.NewClient()
	i.client.Auth.Name = name
	i.player = bot.NewPlayer(i.client, bot.DefaultSettings)
	bot.EventsListener{
		GameStart:  i.onGameStart,
		Disconnect: onDisconnect,
	}.Attach(i.client)
	return
}

func (i *individual) run(address string) {
	//Login
	err := i.client.JoinServer(address)
	if err != nil {
		log.Printf("[%d]Login fail: %v", i.id, err)
		return
	}
	log.Printf("[%d]Login success", i.id)

	//JoinGame
	if err = i.client.HandleGame(); err == nil {
		panic("HandleGame never return nil")
	}
	log.Printf("[%d] Handle game error: %v", i.id, err)
}

func (i *individual) onGameStart() error {
	log.Printf("[%d]Game start", i.id)
	return nil
}

type DisconnectErr struct {
	Reason chat.Message
}

func (d DisconnectErr) Error() string {
	return "disconnect: " + d.Reason.String()
}

func onDisconnect(reason chat.Message) error {
	return DisconnectErr{Reason: reason}
}
