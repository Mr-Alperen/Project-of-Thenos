package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"

	"golang.org/x/crypto/curve25519"

	"github.com/Mr-Alperen/Project-of-Thenos/core/crypto"
)

// GenerateX25519Keypair generates a new X25519 ephemeral key pair (pub, priv).
func GenerateX25519Keypair() (pub, priv []byte, err error) {
	priv = make([]byte, 32)
	if _, err = rand.Read(priv); err != nil {
		return nil, nil, err
	}
	pub, err = curve25519.X25519(priv, curve25519.Basepoint)
	if err != nil {
		return nil, nil, err
	}
	return pub, priv, nil
}

// ComputeSharedSecret computes the X25519 shared secret.
func ComputeSharedSecret(priv, peerPub []byte) ([]byte, error) {
	return curve25519.X25519(priv, peerPub)
}

// DeriveSessionKey derives AEAD key material from the shared secret and nonces.
func DeriveSessionKey(sharedSecret, clientNonce, serverNonce []byte) ([]byte, error) {
	if len(sharedSecret) == 0 {
		return nil, errors.New("empty shared secret")
	}
	salt := append(clientNonce, serverNonce...)
	return crypto.DeriveKey(sharedSecret, salt, []byte("thenos session"), crypto.AEADKeySize)
}

// ClientHello is a simple wire-serializable structure used during handshake.
type ClientHello struct {
	Version     uint16
	ClientPub   []byte // 32
	ClientNonce []byte // recommended 24
	IdentityPub []byte // variable (Ed25519)
}

// Serialize serializes ClientHello into bytes: version|len(pub)|pub|len(nonce)|nonce|len(id)|id
func (c *ClientHello) Serialize() []byte {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, c.Version)
	binary.Write(buf, binary.BigEndian, uint16(len(c.ClientPub)))
	buf.Write(c.ClientPub)
	binary.Write(buf, binary.BigEndian, uint16(len(c.ClientNonce)))
	buf.Write(c.ClientNonce)
	binary.Write(buf, binary.BigEndian, uint16(len(c.IdentityPub)))
	buf.Write(c.IdentityPub)
	return buf.Bytes()
}

// ParseClientHello parses a ClientHello from bytes.
func ParseClientHello(b []byte) (*ClientHello, error) {
	r := bytes.NewReader(b)
	var ver uint16
	if err := binary.Read(r, binary.BigEndian, &ver); err != nil {
		return nil, err
	}
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	clientPub := make([]byte, l)
	if _, err := r.Read(clientPub); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	clientNonce := make([]byte, l)
	if _, err := r.Read(clientNonce); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	id := make([]byte, l)
	if _, err := r.Read(id); err != nil {
		return nil, err
	}
	return &ClientHello{Version: ver, ClientPub: clientPub, ClientNonce: clientNonce, IdentityPub: id}, nil
}
