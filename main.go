package main

import (
	"log"
	"net"
)

func main() {
	logger := log.Default()
	ln, listenErr := net.Listen("tcp", "127.0.0.1:1234")

	if listenErr != nil {
		log.Fatal(listenErr)
	}

	logger.Println("[tcp] 127.0.0.1:1234 listening...")
	for {
		ln.Accept()
	}
	// net.Dial("tcp", "127.0.0.1")
}
