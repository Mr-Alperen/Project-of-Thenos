package router

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/Mr-Alperen/Project-of-Thenos/core/protocol"
)

// Dispatcher routes protocol frames between connected clients.
// It keeps a map of client IDs to net.Conn and provides thread-safe
// Register/Unregister and Dispatch helpers.
type Dispatcher struct {
	mu    sync.RWMutex
	conns map[string]net.Conn
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{conns: make(map[string]net.Conn)}
}

// Register associates a client ID with a connection.
func (d *Dispatcher) Register(id string, conn net.Conn) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.conns[id] = conn
	log.Println("router: registered", id)
}

// Unregister removes a client from the dispatcher.
func (d *Dispatcher) Unregister(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if conn, ok := d.conns[id]; ok {
		conn.Close()
		delete(d.conns, id)
	}
	log.Println("router: unregistered", id)
}

// SendFrame sends a protocol.Frame to the given client ID.
func (d *Dispatcher) SendFrame(id string, fr *protocol.Frame) error {
	d.mu.RLock()
	conn, ok := d.conns[id]
	d.mu.RUnlock()
	if !ok {
		return fmt.Errorf("no connection for id %s", id)
	}
	return protocol.WriteFrame(conn, fr.Type, fr.Payload)
}

// Dispatch is a generic entrypoint for frames coming from a sender.
// For application frames it extracts recipient(s) via router conventions
// and forwards the payload. This function is intentionally simple —
// higher-level routing policies (ACLs, rooms, groups) should be implemented
// by the caller.
func (d *Dispatcher) Dispatch(senderID string, fr *protocol.Frame) error {
	switch fr.Type {
	case protocol.TypeTextMessage:
		// payload format: recipient_len(uint16)|recipient(bytes)|message...
		recipient, msg, err := ParseRecipientPrefixed(fr.Payload)
		if err != nil {
			return err
		}
		out := make([]byte, 0, len(msg)+2)
		// forward as-is as an application frame
		return d.SendFrame(recipient, &protocol.Frame{Type: protocol.TypeTextMessage, Payload: msg, Length: uint32(len(msg))})
	default:
		// by default try to forward to sender (echo) or ignore
		log.Printf("router: unhandled frame type %02x from %s\n", fr.Type, senderID)
		return nil
	}
}
