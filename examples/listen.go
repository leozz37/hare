package main

import (
	"fmt"

	"github.com/leozz37/hare"
)

func main() {
	r := hare.Listen("3000")

	for {
		if r.HasNewMessages {
			fmt.Println(hare.GetMessage())
		}
	}
}
