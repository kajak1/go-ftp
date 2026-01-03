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
		conn, err := ln.Accept()
		if err != nil {
			logger.Fatalf("Error in accepting %+v\n", err)
		}
		logger.Printf("Accepted connection %+v\n", conn)
		
		buffer := make([]byte, 64)
		n, err := conn.Read(buffer) 
		if err != nil {
			logger.Fatalf("Failed to read byes")
			logger.Fatal(err)
		}

		logger.Println(n)
		logger.Println(string(buffer))
	}
}
