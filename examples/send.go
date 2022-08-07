package main

import (
	"github.com/leozz37/hare/websocket"
)

// SendExample is an example of how to send messages
func SendExample() {
	err := websocket.Send("3000", "Hey")
	if err != nil {
		panic(err)
	}
}
