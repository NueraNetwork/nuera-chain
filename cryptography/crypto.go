package cryptography

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// Sign signs a message using a private key
func Sign(message []byte, privateKey []byte) ([]byte, error) {
	privKey := new(ecdsa.PrivateKey)
	privKey.Curve = elliptic.P256()
	privKey.D = new(big.Int).SetBytes(privateKey)
	privKey.PublicKey.X, privKey.PublicKey.Y = privKey.Curve.ScalarBaseMult(privateKey)

	hashed := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hashed[:])
	if err != nil {
		return nil, err
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

// VerifySignature verifies a signature using a public key
func VerifySignature(message []byte, signature []byte, publicKey []byte) bool {
	// Parse the public key
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, publicKey)
	if x == nil || y == nil {
		return false // Invalid public key
	}

	pubKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	// Hash the message
	hashed := sha256.Sum256(message)

	// Verify the signature
	return ecdsa.VerifyASN1(pubKey, hashed[:], signature)
}
