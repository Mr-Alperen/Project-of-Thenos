package crypto

import (
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/hkdf"
)

// DeriveKey uses HKDF-SHA256 to derive a key of length "length" bytes
// from the provided secret, salt and info values.
func DeriveKey(secret, salt, info []byte, length int) ([]byte, error) {
	hk := hkdf.New(sha256.New, secret, salt, info)
	out := make([]byte, length)
	if _, err := io.ReadFull(hk, out); err != nil {
		return nil, err
	}
	return out, nil
}
