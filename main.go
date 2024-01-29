package main

import (
	"io"
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

	bserver := chooseBackend()
	proxy(bserver, conn)

}

func proxy(addr string, c net.Conn) error {
	bc, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalf("error connecting to backend server : %s, error: %s", addr, err)
	}

	go io.Copy(bc, c)

	go io.Copy(c, bc)

}

func chooseBackend() string {
	//TODO: Choose randomly
	return servers[0]
}
