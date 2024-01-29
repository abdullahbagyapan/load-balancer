package main

import (
	"log"
	"net"
)

var (
	listenAddr = "localhost:8080"

	servers = []string{
		"localhost:5000",
		"localhost:5001",
		"localhost:5002",
		"localhost:5003	",
	}
)

func main() {

	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("error listening %s, error: %s", listenAddr, err)
	}

	defer listener.Close()

	conn, err := listener.Accept()

	if err != nil {
		log.Printf("error accepting connection, error: %s", err)
	}

}

func chooseBackend() string {
	//TODO: Choose randomly
	return servers[0]
}
