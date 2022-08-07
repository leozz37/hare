package websocket

import (
	"testing"
)

// TestSendFailConnection tries to send a message to a non-existent port and should return error
// ps: should use a different port than the one used by another test
func TestSendFailConnection(t *testing.T) {
	err := Send("7561", "Hello")
	if err == nil {
		t.Errorf("Sended message to unknow port")
	}
}

// TestListenBusyPort tries to connect to a busy port and should return error
func TestListenBusyPort(t *testing.T) {
	Listen("3214")
	_, err := Listen("3214")
	if err == nil {
		t.Errorf("Listen opened a busy port")
	}
}
