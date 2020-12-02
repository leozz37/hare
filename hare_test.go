package hare

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	expected := "Hello World"

	r := Listen("3000")
	Send("3000", expected)

	if r.HasNewMessages {
		if GetMessage() != expected {
			t.Errorf("Output different from input")
		}
	}
}

func TestListen(t *testing.T) {
	r := Listen("3001")
	Send("3001", "Hey Hey")

	// waiting for the message to be sended
	time.Sleep(time.Second)

	if r.HasNewMessages == false {
		t.Errorf("Listen didn't receive any messsage")
	}
}
