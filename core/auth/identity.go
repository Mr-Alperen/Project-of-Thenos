package auth

import (
	"crypto/ed25519"
	"crypto/rand"
)

// GenerateIdentityKeypair creates a new Ed25519 identity keypair.
func GenerateIdentityKeypair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	return pub, priv, err
}

// SignIdentity signs message with the given Ed25519 private key.
func SignIdentity(priv ed25519.PrivateKey, message []byte) []byte {
	return ed25519.Sign(priv, message)
}

// VerifyIdentity verifies the Ed25519 signature for message.
func VerifyIdentity(pub ed25519.PublicKey, message, sig []byte) bool {
	return ed25519.Verify(pub, message, sig)
}
