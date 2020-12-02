package main

import (
	"log"
	"net"
)

const (
	connHost = "localhost"
	connType = "tcp"
)

func Sender(port, message string) {
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
