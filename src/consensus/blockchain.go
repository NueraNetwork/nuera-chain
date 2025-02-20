package consensus

import (
	"crypto/ed25519"
	"log"
	"github.com/NueraNetwork/nuera-chain/src/crypto"
)

// Blockchain represents the blockchain
type Blockchain struct {
	Blocks     []*Block
	PublicKey  ed25519.PublicKey // Store the public key for validation
	Validators []string          // List of validators
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain(publicKey ed25519.PublicKey) *Blockchain {
	privateKey, _ := crypto.GenerateKeyPair()
	validators := []string{"Node1", "Node2", "Node3"} // Initialize with some validators
	genesisBlock := NewBlock(0, "Genesis Block", "", privateKey, validators[0])
	return &Blockchain{
		Blocks:     []*Block{genesisBlock},
		PublicKey:  publicKey,
		Validators: validators,
	}
}

// SelectValidator selects a validator from the list
func (bc *Blockchain) SelectValidator() string {
	// For now, we'll just rotate through the validators
	nextValidator := bc.Validators[len(bc.Blocks)%len(bc.Validators)]
	return nextValidator
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string, privateKey ed25519.PrivateKey) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	validator := bc.SelectValidator() // Select a validator
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash, privateKey, validator)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// ValidateBlockchain validates the integrity of the blockchain
func (bc *Blockchain) ValidateBlockchain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		// Check if the current block's hash is valid
		if currentBlock.Hash != currentBlock.CalculateHash() {
			log.Printf("Block %d: Invalid hash (expected: %s, got: %s)\n", currentBlock.Index, currentBlock.CalculateHash(), currentBlock.Hash)
			return false
		}

		// Check if the previous block's hash matches
		if currentBlock.PrevHash != prevBlock.Hash {
			log.Printf("Block %d: Previous hash mismatch (expected: %s, got: %s)\n", currentBlock.Index, prevBlock.Hash, currentBlock.PrevHash)
			return false
		}

		// Verify the block's signature using the stored public key
		if !currentBlock.VerifyBlock(bc.PublicKey) {
			log.Printf("Block %d: Invalid signature\n", currentBlock.Index)
			return false
		}
	}
	return true
}
