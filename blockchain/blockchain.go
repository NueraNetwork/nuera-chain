package blockchain

import (
	"errors"
	"time"

	"github.com/NueraNetwork/nuera-chain/consensus" // Updated import path
	"github.com/NueraNetwork/nuera-chain/cryptography" // Updated import path
	"github.com/NueraNetwork/nuera-chain/types" // Import the types package
)

// Blockchain represents the blockchain
type Blockchain struct {
	Blocks        []*types.Block // Use types.Block
	Stakeholders []types.Stakeholder // Use types.Stakeholder
}

// CreateBlock creates a new block using PoS consensus
func (bc *Blockchain) CreateBlock(transactions []types.Transaction) (*types.Block, error) { // Use types.Transaction and types.Block
	// Select a validator using PoS
	validator, err := consensus.SelectValidator(bc.Stakeholders)
	if err != nil {
		return nil, err
	}

	// Create the block
	newBlock := &types.Block{
		Index:        len(bc.Blocks),
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     bc.Blocks[len(bc.Blocks)-1].Hash,
		Validator:    validator.Address,
	}

	// Sign the block
	signature, err := cryptography.Sign(newBlock.Hash, validator.PrivateKey)
	if err != nil {
		return nil, err
	}
	newBlock.Signature = signature

	// Validate the block
	if !consensus.ValidateBlock(*newBlock, *validator) {
		return nil, errors.New("block validation failed")
	}

	// Add the block to the chain
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock, nil
}
