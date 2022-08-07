package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/leozz37/hare/websocket"
)

func main() {
	sendPtr := flag.Bool("s", false, "Send message to a given address")
	listenPtr := flag.Bool("l", false, "Listen to a given address")
	portPtr := flag.String("p", "", "Port address to bo operated")
	hostPtr := flag.String("h", "localhost", "Host address to bo operated")
	dataPtr := flag.String("d", "", "Data to be sended")

	flag.Parse()
	if *sendPtr && *listenPtr {
		flag.PrintDefaults()
		fmt.Errorf("You can only listen OR send to a port")
	}
	if *portPtr == "" {
		flag.PrintDefaults()
		fmt.Errorf("Flag PORT is required.")
	}

	if *listenPtr {
		err := listenPort(*portPtr, *hostPtr)
		if err != nil {
			fmt.Errorf("Could listen to address")
		}
	}
	if *sendPtr {
		err := sendMessage(*portPtr, *hostPtr, *dataPtr)
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
		panic(err)
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
		panic(err)
	}
	return nil
}
