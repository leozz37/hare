package example

import (
	"github.com/leozz37/hare"
)

func SendExample() {
	err := hare.Send("3000", "Hey")
	if err != nil {
		panic(err)
	}
}
