package server

import (
	"log"
	"net"
)

func StartListener(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Println("Thenos listening on", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
