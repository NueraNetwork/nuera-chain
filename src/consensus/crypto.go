package consensus

import (
	"crypto/ed25519"
	"crypto/rand"
)

// GenerateKeyPair generates a new ed25519 key pair
func GenerateKeyPair() (ed25519.PrivateKey, ed25519.PublicKey) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	return privateKey, publicKey
}
