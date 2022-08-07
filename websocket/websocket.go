package websocket

import (
	"bufio"
	"net"
)

const (
	connectionHost = "0.0.0.0"
	connectionType = "tcp"
)

// Listener tools for socket listening
type Listener struct {
	socketListener net.Listener
	messageManager MessageManager
	running        bool
}

// MessageManager manages message storage
type MessageManager struct {
	HasNewMessage bool
	Message       string
}

// Listen to socket port
func Listen(port string) (Listener, error) {
	listener, err := net.Listen(connectionType, connectionHost+":"+port)
	if err != nil {
		return Listener{}, err
	}

	l := Listener{socketListener: listener}
	go l.listening()

	return Listener{socketListener: listener}, nil
}

// GetMessage return the last message
func (l *Listener) GetMessage() string {
	l.messageManager.HasNewMessage = false
	return l.messageManager.Message
}

// HasNewMessages return if there's new messages on socket
func (l *Listener) HasNewMessages() bool {
	return l.messageManager.HasNewMessage
}

// Stop the listener
func (l *Listener) Stop() {
	l.running = false
}

// listening listen thread for new messages on socket port
func (l *Listener) listening() error {
	for l.running {
		c, _ := l.socketListener.Accept()
		message, _ := bufio.NewReader(c).ReadString('\n')
		l.messageManager.Message = message
		l.messageManager.HasNewMessage = true
	}
	l.socketListener.Close()
	return nil
}

// Send message to socket port
func Send(port, message string) error {
	conn, err := net.Dial(connectionType, connectionHost+":"+port)
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Write([]byte(message))
	return nil
}
