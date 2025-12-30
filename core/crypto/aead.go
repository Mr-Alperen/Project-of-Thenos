package crypto

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
)

// NewAEAD returns an AEAD instance (XChaCha20-Poly1305) for the given key.
// Key must be AEADKeySize bytes long.
func NewAEAD(key []byte) (chacha20poly1305.AEAD, error) {
	if len(key) != AEADKeySize {
		return nil, errors.New("invalid key length for AEAD")
	}
	return chacha20poly1305.NewX(key)
}

// Encrypt encrypts plaintext with AEAD (XChaCha20-Poly1305), generates a random nonce
// and returns (nonce, ciphertext). AAD may be nil.
func Encrypt(key, plaintext, aad []byte) (nonce, ciphertext []byte, err error) {
	aead, err := NewAEAD(key)
	if err != nil {
		return nil, nil, err
	}

	nonce = make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		return nil, nil, err
	}
	ct := aead.Seal(nil, nonce, plaintext, aad)
	return nonce, ct, nil
}

// Decrypt decrypts ciphertext produced by Encrypt using the provided key, nonce and aad.
func Decrypt(key, nonce, ciphertext, aad []byte) ([]byte, error) {
	aead, err := NewAEAD(key)
	if err != nil {
		return nil, err
	}
	if len(nonce) != chacha20poly1305.NonceSizeX {
		return nil, errors.New("invalid nonce size")
	}
	pt, err := aead.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		return nil, err
	}
	return pt, nil
}
