package consensus_test

import (
	"testing"
	"github.com/NueraNetwork/nuera-chain/src/consensus"
)

func TestDistributeRewards(t *testing.T) {
	// Create a list of validators
	validators := []*consensus.Validator{
		{Address: "validator1", Stake: 1000},
		{Address: "validator2", Stake: 2000},
	}

	// Test distributing rewards from a reward pool of 300
	err := consensus.DistributeRewards(validators, 300)
	if err != nil {
		t.Fatalf("DistributeRewards failed: %v", err)
	}

	// Check if the rewards were distributed correctly
	if validators[0].Stake != 1100 {
		t.Errorf("Expected validator1 stake to be 1100, got %d", validators[0].Stake)
	}

	if validators[1].Stake != 2200 {
		t.Errorf("Expected validator2 stake to be 2200, got %d", validators[1].Stake)
	}

	// Test distributing rewards with a negative reward pool
	err = consensus.DistributeRewards(validators, -100)
	if err == nil {
		t.Error("Expected error for negative reward pool, got nil")
	}

	// Test distributing rewards exceeding the total supply
	err = consensus.DistributeRewards(validators, consensus.TotalSupply+1)
	if err == nil {
		t.Error("Expected error for reward pool exceeding total supply, got nil")
	}
}
