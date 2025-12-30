package crypto

// Package-level constants for cryptographic parameters used across the project.

const (
	// AEAD key size in bytes (XChaCha20-Poly1305 uses 32-byte keys)
	AEADKeySize = 32

	// XChaCha20-Poly1305 nonce size (24 bytes)
	XChaCha20NonceSize = 24

	// Maximum allowed payload for a single frame (16 MiB default)
	MaxPayloadSize = 16 * 1024 * 1024
)
