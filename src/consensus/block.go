package consensus

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Signature []byte
	Validator string
}

// NewBlock creates a new block
func NewBlock(index int, data string, prevHash string, privateKey ed25519.PrivateKey, validator string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
		Validator: validator,
	}
	block.Hash = block.CalculateHash()

	// Sign the block's hash
	block.Signature = ed25519.Sign(privateKey, []byte(block.Hash))

	return block
}

// CalculateHash computes the hash of the block using SHA-256
func (b *Block) CalculateHash() string {
	// Concatenate block data to create the hash input
	hashInput := fmt.Sprintf("%d%s%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash, b.Validator)

	// Compute the SHA-256 hash
	hash := sha256.Sum256([]byte(hashInput))

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash[:])
}

// VerifyBlock verifies the block's signature
func (b *Block) VerifyBlock(publicKey ed25519.PublicKey) bool {
	// Verify the block's hash using the public key
	return ed25519.Verify(publicKey, []byte(b.Hash), b.Signature)
}
