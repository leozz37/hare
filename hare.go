package hare

import (
	"bufio"
	"net"
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

func listening() error {
	for {
		c, _ := listener.SocketListener.Accept()
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
func Listen(port string) (*Listener, error) {
	var err error
	listener.SocketListener, err = net.Listen(connType, connHost+":"+port)
	if err != nil {
		return nil, err
	}

	go listening()

	return &listener, nil
}

// Send message to socket port
func Send(port, message string) error {
	conn, err := net.Dial(connType, connHost+":"+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Write([]byte(message))
	return nil
}
