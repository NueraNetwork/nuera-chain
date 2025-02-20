package consensus

import (
	"crypto/rand"
	"math/big"
	"github.com/NueraNetwork/nuera-chain/cryptography" // Updated import path
	"github.com/NueraNetwork/nuera-chain/types" // Import the types package
)

// SelectValidator selects the next block validator based on stake
func SelectValidator(stakeholders []types.Stakeholder) (*types.Stakeholder, error) { // Use types.Stakeholder
	totalStake := new(big.Int)
	for _, sh := range stakeholders {
		totalStake.Add(totalStake, sh.Stake)
	}

	// Select a random stake-weighted validator
	randVal, err := rand.Int(rand.Reader, totalStake)
	if err != nil {
		return nil, err
	}

	runningTotal := new(big.Int)
	for _, sh := range stakeholders {
		runningTotal.Add(runningTotal, sh.Stake)
		if runningTotal.Cmp(randVal) > 0 {
			return &sh, nil
		}
	}

	return nil, nil
}

// ValidateBlock validates a block using the validator's signature
func ValidateBlock(block types.Block, validator types.Stakeholder) bool { // Use types.Block and types.Stakeholder
	// Convert block.Hash (string) to []byte
	hashBytes := []byte(block.Hash)

	// Verify the block signature
	return cryptography.VerifySignature(hashBytes, block.Signature, validator.PublicKey)
}
