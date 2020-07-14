package net

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	c := make(chan int, 1)
	go server(t, c)
	<-c
	client(t)
	<-c
}

func server(t *testing.T, c chan<- int) {
	l, err := ListenRCON("localhost:25575")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	c <- 1 // prepared

	conn, err := l.Accept()
	if err != nil {
		t.Fatal(err)
	}

	err = conn.AcceptLogin("RightPassword")
	if err != nil {
		t.Fatal("password wrong")
	}

	cmd, err := conn.AcceptCmd()
	if err != nil {
		t.Log(err)
		return
	}

	resp := handleCommand(cmd)
	err = conn.RespCmd(resp)
	if err != nil {
		t.Fatal(err)
	}

	c <- 2 // finished
}

func handleCommand(cmd string) (resp string) {
	return fmt.Sprintf("your command is %q", cmd)
}

func client(t *testing.T) {
	conn, err := DialRCON("localhost:25575", "RightPassword")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = conn.Cmd("TEST COMMAND")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := conn.Resp()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Server response: %q", resp)
}
