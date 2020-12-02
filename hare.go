package main

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

// Listen to a socket port
func Listen(port string) {
	l, err := net.Listen(connType, connHost+":"+port)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		message, _  := bufio.NewReader(c).ReadString('\n')
		log.Print(message)
	}
}


func main() {
	Listen("3000")
}