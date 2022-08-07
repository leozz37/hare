package example

import (
	"fmt"

	"github.com/leozz37/hare/websocket"
)

// ListenExample is an example of how to listen for messages
func ListenExample() {
	r, err := websocket.Listen("3000")
	if err != nil {
		panic(err)
	}

	for {
		if r.HasNewMessages() {
			fmt.Println(r.GetMessage())
		}
	}
}
