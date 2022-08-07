package main

import (
	"github.com/leozz37/hare"
)

func main() {
	err := hare.Send("3000", "Hey")
	if err != nil {
		panic(err)
	}
}
