package main

import (
	"log"
	"time"

	//"github.com/mattn/go-colorable"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

const timeout = 45

var (
	c *bot.Client
	p *bot.Player

	watch chan time.Time
)

func main() {
	//log.SetOutput(colorable.NewColorableStdout()) // optional for colorable output
	c = bot.NewClient()
	p = bot.NewPlayer(c, bot.DefaultSettings)

	//Register event handlers
	bot.EventsListener{
		GameStart:  onGameStart,
		ChatMsg:    onChatMsg,
		SystemMsg:  onSystemMsg,
		Disconnect: onDisconnect,
		Death:      onDeath,
	}.Attach(c)
	c.Events.AddListener(soundListener)

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

	return UseItem(0)
}

var soundListener = bot.PacketHandler{
	ID:       packetid.ClientboundSound,
	Priority: 0,
	F: func(p pk.Packet) error {
		var (
			SoundID       pk.VarInt
			SoundCategory pk.VarInt
			X, Y, Z       pk.Int
			Volume, Pitch pk.Float
		)
		if err := p.Scan(&SoundID, &SoundCategory, &X, &Y, &Z, &Volume, &Pitch); err != nil {
			return err
		}
		return onSound(int(SoundID), int(SoundCategory), float64(X)/8, float64(Y)/8, float64(Z)/8, float32(Volume), float32(Pitch))
	},
}

func UseItem(hand int32) error {
	return c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundUseItem,
		pk.VarInt(hand),
	))
}

//goland:noinspection SpellCheckingInspection
func onSound(id int, category int, x, y, z float64, volume, pitch float32) error {
	if id == 369 {
		if err := UseItem(0); err != nil { //retrieve
			return err
		}
		log.Println("gra~")
		time.Sleep(time.Millisecond * 300)
		if err := UseItem(0); err != nil { //throw
			return err
		}
		watch <- time.Now()
	}
	return nil
}

func onChatMsg(c *bot.PlayerMessage) error {
	log.Println("Chat:", c.SignedMessage)
	return nil
}

func onSystemMsg(c chat.Message, pos byte) error {
	log.Printf("System: %v, Location: %v", c, pos)
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
			if err := UseItem(0); err != nil {
				panic(err)
			}
		}
		to.Reset(time.Second * timeout)
	}
}
