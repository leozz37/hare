package websocket

import (
	"testing"
	"time"
)

// TestSendMessages checks if the message is sended to the right port, the payload
// is correct and HasNewMessages is true when the message is received
func TestSendMessages(t *testing.T) {
	testTable := []struct {
		name     string
		expected string
		port     string
		payload  string
	}{
		{"Send to port 3001", "Hello", "3001", "Hello"},
		{"Send to port 3003", "Hello", "3003", "Hello"},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			r, _ := Listen(test.port)
			Send(test.port, test.payload)

			time.Sleep(time.Second)
			if !r.HasNewMessages() {
				t.Error("Listen didn't receive any message")
			}

			if r.GetMessage() != test.expected {
				t.Error("Output different from input")
			}
		})
	}
}

// TestListen checks if Listen receives a message and HasNewMessages is working
func TestListen(t *testing.T) {
	testTable := []struct {
		name    string
		payload string
		port    string
	}{
		{"Send to port 3002", "Hello", "3002"},
		{"Send to port 4000", "World", "4000"},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			r, err := Listen(test.port)
			if err != nil {
				t.Errorf("Error listening: %s", err.Error())
			}

			Send(test.port, test.payload)
			time.Sleep(time.Second)
			if !r.HasNewMessages() {
				t.Error("Listen didn't receive any message")
			}
		})
	}
}

// TestListenStop checks if the listen stops listening when the port is closed
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
