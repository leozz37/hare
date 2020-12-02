package hare

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connType = "tcp"
)

// Listener tools for socket listening
type Listener struct {
	SocketListener net.Listener
	HasNewMessages bool
	Message        string
}

var listener Listener

func listening() {
	for {
		c, err := listener.SocketListener.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		message, _ := bufio.NewReader(c).ReadString('\n')
		listener.Message = message
		listener.HasNewMessages = true
	}
}

// GetMessage from port
func GetMessage() string {
	listener.HasNewMessages = false
	return listener.Message
}

// Listen to socket port
func Listen(port string) *Listener {
	var err error
	listener.SocketListener, err = net.Listen(connType, connHost+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	go listening()

	return &listener
}

// Send message to socket port
func Send(port, message string) {
	conn, err := net.Dial(connType, connHost+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
