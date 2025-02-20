package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

// GenerateKeyPair generates a new ed25519 key pair
func GenerateKeyPair() (ed25519.PrivateKey, ed25519.PublicKey) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err) // Handle key generation error
	}
	return privateKey, publicKey
}
