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
	GetMessage     func() string
	HasNewMessages func() bool
}

// MessageManager manages message storage
type MessageManager struct {
	HasNewMessages bool
	Message        string
}

var listener Listener
var messageManager MessageManager

func listening() error {
	for {
		c, _ := listener.SocketListener.Accept()
		message, _ := bufio.NewReader(c).ReadString('\n')
		messageManager.Message = message
		messageManager.HasNewMessages = true
	}
}

// Listen to socket port
func Listen(port string) (*Listener, error) {
	var err error
	listener.SocketListener, err = net.Listen(connType, connHost+":"+port)
	if err != nil {
		return nil, err
	}

	// GetMessage return the last message
	listener.GetMessage = func() string {
		messageManager.HasNewMessages = false
		return messageManager.Message
	}

	// HasNewMessages return if there's new messages on socket
	listener.HasNewMessages = func() bool {
		return messageManager.HasNewMessages
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
