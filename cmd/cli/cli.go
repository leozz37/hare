package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/leozz37/hare/websocket"
)

func main() {
	sendFlag := flag.Bool("s", false, "Send message to a given address")
	listenFlag := flag.Bool("l", false, "Listen to a given address")
	portAddress := flag.String("p", "", "Port address to bo operated")
	hostAddress := flag.String("h", "localhost", "Host address to bo operated")
	payload := flag.String("d", "", "Data to be sended")

	flag.Parse()
	if *sendFlag && *listenFlag {
		flag.PrintDefaults()
		log.Fatal("You can only listen OR send to a port")
	}
	if *portAddress == "" {
		flag.PrintDefaults()
		log.Fatal("Flag PORT is required.")
	}
	if *listenFlag {
		err := listenPort(*portAddress, *hostAddress)
		if err != nil {
			panic("Error listening: " + err.Error())
		}
	}
	if *sendFlag {
		err := sendMessage(*portAddress, *hostAddress, *payload)
		if err != nil {
			address := *hostAddress + ":" + *portAddress
			panic("Error sending payload: %s" + address + "\nerror: " + err.Error())
		}
	}
}

func listenPort(port, host string) error {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("Listening to port: %s\n", red(port))

	r, err := websocket.Listen(port)
	if err != nil {
		return err
	}

	for {
		if r.HasNewMessages() {
			fmt.Println(r.GetMessage())
		}
	}
}

func sendMessage(port, host, data string) error {
	err := websocket.Send(port, data)
	if err != nil {
		return err
	}
	return nil
}
