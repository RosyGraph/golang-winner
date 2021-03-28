package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	port := ":1100"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not initialize server: %s", err.Error())
	}

	defer l.Close()
	log.Printf("started server on %s", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("could not accept connnection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
