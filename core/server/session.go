package server

import (
	"log"
	"net"

	"github.com/Mr-Alperen/Project-of-Thenos/core/protocol"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("New connection:", conn.RemoteAddr())

	for {
		frame, err := protocol.ReadFrame(conn)
		if err != nil {
			log.Println("Connection closed:", err)
			return
		}

		log.Printf("Frame received | Type: %d | Size: %d bytes\n",
			frame.Type, frame.Length)
	}
}
