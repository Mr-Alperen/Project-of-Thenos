package router

import (
	"encoding/binary"
	"errors"
)

// Message represents a simple application-level message between clients.
type Message struct {
	From string
	To   string
	Body []byte
}

// ParseRecipientPrefixed assumes payload format: recipient_len(uint16)|recipient(bytes)|message...
// returns recipient id and message bytes.
func ParseRecipientPrefixed(payload []byte) (string, []byte, error) {
	if len(payload) < 2 {
		return "", nil, errors.New("payload too short for recipient length")
	}
	l := int(binary.BigEndian.Uint16(payload[:2]))
	if len(payload) < 2+l {
		return "", nil, errors.New("payload too short for recipient")
	}
	recipient := string(payload[2 : 2+l])
	msg := payload[2+l:]
	return recipient, msg, nil
}

// BuildRecipientPrefixed builds a payload with recipient prefix used by Dispatch.
func BuildRecipientPrefixed(recipient string, msg []byte) []byte {
	b := make([]byte, 2+len(recipient)+len(msg))
	binary.BigEndian.PutUint16(b[0:2], uint16(len(recipient)))
	copy(b[2:2+len(recipient)], []byte(recipient))
	copy(b[2+len(recipient):], msg)
	return b
}
