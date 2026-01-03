package main

import (
	"fmt"
	"log"
	"net"

	// "net"
	"os"
)

func main() {
	logger := log.Default()

	hostname, port := parseArgs()

	address := createAddress(hostname, port)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Connection to %s created", address)
	
	conn.Write([]byte("hello"))
}

func parseArgs() (hostname string, port string) {
	args := os.Args[1:]

	hostname = args[0]
	port = args[1]

	return hostname, port
}

func createAddress(hostname string, port string) string {
	return fmt.Sprintf("%s:%s", hostname, port)
}
