package main

import (
	"fmt"

	"github.com/leozz37/hare"
)

func main() {
	r, err := hare.Listen("3000")
	if err != nil {
		panic(err)
	}

	for {
		if r.HasNewMessages() {
			fmt.Println(r.GetMessage())
		}
	}
}
