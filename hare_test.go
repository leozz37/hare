package hare

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	expected := "Hello World"

	r, _ := Listen("3000")
	Send("3000", expected)

	if r.HasNewMessages {
		if GetMessage() != expected {
			t.Errorf("Output different from input")
		}
	}
}

func TestSendFailConnect(t *testing.T) {
	err := Send("8005", "Hello")
	if err == nil {
		t.Errorf("Sended message to unknow port")
	}
}

func TestListen(t *testing.T) {
	r, _ := Listen("3001")
	Send("3001", "Hey Hey")

	// waiting for the message to be sended
	time.Sleep(time.Second)

	if r.HasNewMessages == false {
		t.Errorf("Listen didn't receive any messsage")
	}
}

func TestListenBusyPort(t *testing.T) {
	Listen("3000")
	_, err := Listen("3000")
	if err == nil {
		t.Errorf("Listen opened a busy port")
	}
}
