package example

import (
	"github.com/leozz37/hare"
)

// SendExample is an example of how to send messages
func SendExample() {
	err := hare.Send("3000", "Hey")
	if err != nil {
		panic(err)
	}
}
