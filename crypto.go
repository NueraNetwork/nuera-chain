package main

import (
    "crypto/cipher"
    "crypto/rand"
    "fmt"            // For printing output
    "io"
    "golang.org/x/crypto/chacha20" // For creating a cipher.Stream
    "go.dedis.ch/kyber/v3"         // Kyber v3.0.13
    "go.dedis.ch/kyber/v3/group/edwards25519" // Group for Kyber
    "go.dedis.ch/kyber/v3/sign/schnorr" // For signing and verification
)

var suite = edwards25519.NewBlakeSHA256Ed25519()

// newCipherStream creates a cipher.Stream from crypto/rand
func newCipherStream() (cipher.Stream, error) {
    key := make([]byte, 32) // 256-bit key for ChaCha20
    if _, err := io.ReadFull(rand.Reader, key); err != nil {
        return nil, err
    }

    nonce := make([]byte, chacha20.NonceSize) // Nonce for ChaCha20
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return chacha20.NewUnauthenticatedCipher(key, nonce)
}

// GenerateKeyPair generates a new private/public key pair
func GenerateKeyPair() (kyber.Scalar, kyber.Point, error) {
    stream, err := newCipherStream()
    if err != nil {
        return nil, nil, err
    }

    privateKey := suite.Scalar().Pick(stream) // Generate private key
    publicKey := suite.Point().Mul(privateKey, nil) // Derive public key

    return privateKey, publicKey, nil
}

// Sign signs a message with the private key
func Sign(privateKey kyber.Scalar, msg []byte) []byte {
    // Use Kyber's Schnorr implementation for signing
    signature, err := schnorr.Sign(suite, privateKey, msg)
    if err != nil {
        panic(err)
    }
    return signature
}

// VerifySignature checks if the signature is valid for the message with the public key
func VerifySignature(publicKey kyber.Point, signature, msg []byte) bool {
    // Use Kyber's Schnorr implementation for verification
    err := schnorr.Verify(suite, publicKey, msg, signature)
    return err == nil
}

func main() {
    // Example usage of the GenerateKeyPair function
    privateKey, publicKey, err := GenerateKeyPair()
    if err != nil {
        panic(err)
    }

    fmt.Println("Private Key:", privateKey)
    fmt.Println("Public Key:", publicKey)

    // Example usage of the Sign and VerifySignature functions
    msg := []byte("Hello, Nuera Network!")
    signature := Sign(privateKey, msg)
    fmt.Println("Signature:", signature)

    valid := VerifySignature(publicKey, signature, msg)
    fmt.Println("Signature Valid:", valid)
}
