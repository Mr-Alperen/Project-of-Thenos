package protocol

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	TypeTextMessage byte = 0x01
	TypeFileMeta    byte = 0x02
	TypeFileChunk   byte = 0x03
	TypeFileEnd     byte = 0x04
	TypeHeartbeat   byte = 0x05
)

type Frame struct {
	Type   byte
	Length uint32
	Payload []byte
}

func ReadFrame(r io.Reader) (*Frame, error) {
	header := make([]byte, 5)

	_, err := io.ReadFull(r, header)
	if err != nil {
		return nil, err
	}

	frameType := header[0]
	length := binary.BigEndian.Uint32(header[1:5])

	if length > 10*1024*1024 {
		return nil, errors.New("frame too large")
	}

	payload := make([]byte, length)
	_, err = io.ReadFull(r, payload)
	if err != nil {
		return nil, err
	}

	return &Frame{
		Type: frameType,
		Length: length,
		Payload: payload,
	}, nil
}

func WriteFrame(w io.Writer, frameType byte, payload []byte) error {
	header := make([]byte, 5)
	header[0] = frameType
	binary.BigEndian.PutUint32(header[1:], uint32(len(payload)))

	if _, err := w.Write(header); err != nil {
		return err
	}
	if _, err := w.Write(payload); err != nil {
		return err
	}
	return nil
}
