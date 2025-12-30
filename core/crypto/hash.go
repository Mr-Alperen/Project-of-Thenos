package crypto

import "crypto/sha256"

// SHA256Sum returns the 32-byte SHA-256 sum of data.
func SHA256Sum(data []byte) [32]byte {
	return sha256.Sum256(data)
}

// SHA256 returns the SHA-256 digest as a byte slice.
func SHA256(data []byte) []byte {
	s := sha256.Sum256(data)
	b := make([]byte, len(s))
	copy(b, s[:])
	return b
}
