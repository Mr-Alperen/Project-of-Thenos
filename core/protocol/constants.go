package protocol

const (
	// Protocol identity
	ProtocolName    = "THENOS"
	ProtocolVersion = 1

	// Frame header
	HeaderSize = 5 // TYPE (1) + LENGTH (4)

	// Limits
	MaxFrameSize   = 10 * 1024 * 1024 // 10 MB
	MaxMessageSize = 64 * 1024        // 64 KB text message
	MaxFileChunk   = 64 * 1024        // 64 KB file chunk
)

const (
	// Frame types
	TypeTextMessage byte = 0x01
	TypeFileMeta    byte = 0x02
	TypeFileChunk   byte = 0x03
	TypeFileEnd     byte = 0x04
	TypeHeartbeat   byte = 0x05

	// Control / future
	TypeAuthInit   byte = 0x10
	TypeAuthProof  byte = 0x11
	TypeAuthResult byte = 0x12
)
