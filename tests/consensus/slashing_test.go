package consensus_test

import (
	"testing"
	"github.com/NueraNetwork/nuera-chain/src/consensus"
)

func TestSlashValidator(t *testing.T) {
	// Create a validator with 1000 tokens staked
	validator := &consensus.Validator{
		Address: "validator1",
		Stake:   1000,
	}

	// Test slashing 5% of the validator's stake
	err := consensus.SlashValidator(validator, 5)
	if err != nil {
		t.Fatalf("SlashValidator failed: %v", err)
	}

	// Check if the validator's stake is now 950 (1000 - 5%)
	if validator.Stake != 950 {
		t.Errorf("Expected stake to be 950, got %d", validator.Stake)
	}

	// Test slashing with an invalid penalty percentage (110%)
	err = consensus.SlashValidator(validator, 110)
	if err == nil {
		t.Error("Expected error for invalid penalty percentage, got nil")
	}
}
