package net

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	p := make(chan int, 1)
	go server(t, p)
	<-p
	client(t)
}

func server(t *testing.T, prepare chan<- int) {
	l, err := ListenRCON("localhost:25575")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	prepare <- 1

	for {
		conn, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func(conn RCONServerConn) {
			err := conn.AcceptLogin("RightPassword")
			if err != nil {
				t.Fatal("password wrong")
			}

			for {
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
			}
		}(conn)
	}
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
