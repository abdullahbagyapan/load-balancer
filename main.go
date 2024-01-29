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
		"localhost:5003",
	}

	counter = 0
)

func main() {

	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("error listening %s, error: %s", listenAddr, err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("error accepting connection, error: %s", err)
		}

		go func() {
			bserver := chooseBackend()

			err := proxy(bserver, conn)

			if err != nil {
				log.Printf("WARNING: proxying failed %v", err)
			}
		}()

	}

}

func proxy(addr string, c net.Conn) error {
	bc, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalf("error connecting to backend server : %s, error: %s", addr, err)
	}

	go io.Copy(bc, c)

	go io.Copy(c, bc)
	return nil
}

func chooseBackend() string {

	server := servers[counter%len(servers)]
	counter++

	return server
}
