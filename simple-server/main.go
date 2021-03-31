package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not start server: %s", err.Error())
	}
	defer l.Close()
	clients := make([]net.Conn, 0)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("could not accept connection: %s", err.Error())
			continue
		}
		clients := append(clients, conn)
		fmt.Println(clients)
		go func(c net.Conn) {
			for _, client := range clients {
				io.Copy(c, client)
			}
			c.Close()
		}(conn)
	}
}
