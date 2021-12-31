package hare

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	expected := "Hello World"

	r, _ := Listen("3000")
	Send("3000", expected)

	if r.HasNewMessages() {
		if r.GetMessage() != expected {
			t.Errorf("Output different from input")
		}
	}
}

func TestSendFailConnect(t *testing.T) {
	err := Send("3001", "Hello")
	if err == nil {
		t.Errorf("Sended message to unknow port")
	}
}

func TestListen(t *testing.T) {
	r, err := Listen("3002")
	if err != nil {
		panic(err)
	}
	Send("3002", "Hey Hey")

	// waiting for the message to be sended
	time.Sleep(time.Second)

	if r.HasNewMessages() == false {
		t.Errorf("Listen didn't receive any messsage")
	}
}

func TestListenBusyPort(t *testing.T) {
	Listen("3003")
	_, err := Listen("3003")
	if err == nil {
		t.Errorf("Listen opened a busy port")
	}
}

func TestGetMessage(t *testing.T) {
	expected := "Hey gopher!"

	r, err := Listen("3004")
	if err != nil {
		t.Errorf("Couldn't open port")
	}

	Send("3004", expected)
	time.Sleep(time.Second)

	if r.GetMessage() != expected {
		t.Errorf("Input message different from output")
	}

}

func TestListenStop(t *testing.T) {
	expected := "Hello World"

	r, _ := Listen("3005")
	Send("3005", expected)

	if r.HasNewMessages() {
		if r.GetMessage() != expected {
			t.Errorf("Output different from input")
		}
	}

	r.Stop()
	err := Send("3005", "This should fails")
	if err == nil {
		t.Errorf("Couldn't close connection")
	}
}
