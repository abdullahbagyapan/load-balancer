package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
)

var (
	listenAddr = "localhost:8080"

	server = []string{
		"localhost:3000",
		"localhost:3001",
		"localhost:3002",
	}
)

func main() {

	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("Error listening %s , err = %v", listenAddr, err)
	}
	defer listener.Close()

	for {
		// client connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection %v", err)
		}
		backend := chooseBackend()

		go func() {
			err = proxy(backend, conn)
			if err != nil {
				log.Printf("Error proxying : %v", err)
			}
		}()

	}

}

func proxy(backend string, inConn net.Conn) error {

	// connection to backend
	outConn, err := net.Dial("tcp", backend) //  reverse proxy

	if err != nil {
		return fmt.Errorf("error connecting server %s , err %v", backend, err)
	}

	go io.Copy(outConn, inConn)

	go io.Copy(inConn, outConn)

	return nil
}

func chooseBackend() string {
	randInt := rand.Intn(3)

	return server[randInt]
}
