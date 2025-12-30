package protocol

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidFrameType = errors.New("invalid frame type")
	ErrPayloadTooLarge  = errors.New("payload exceeds allowed size")
	ErrInvalidPayload   = errors.New("invalid payload format")
)

func ValidateFrame(frame *Frame) error {
	switch frame.Type {

	case TypeTextMessage:
		if frame.Length > MaxMessageSize {
			return ErrPayloadTooLarge
		}
		return validateJSON(frame.Payload)

	case TypeFileMeta:
		return validateJSON(frame.Payload)

	case TypeFileChunk:
		if frame.Length > MaxFileChunk {
			return ErrPayloadTooLarge
		}
		return nil

	case TypeFileEnd:
		if frame.Length != 0 {
			return ErrInvalidPayload
		}
		return nil

	case TypeHeartbeat:
		return nil

	case TypeAuthInit, TypeAuthProof, TypeAuthResult:
		return validateJSON(frame.Payload)

	default:
		return ErrInvalidFrameType
	}
}

func validateJSON(data []byte) error {
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return ErrInvalidPayload
	}
	return nil
}
