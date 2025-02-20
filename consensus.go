package main

import (
	"crypto/rand"
	"math/big"
	"time"
)

// Stake represents the amount of tokens a node has staked
type Stake struct {
	NodeID string
	Amount int64
}

// Consensus implements the Proof of Stake logic
type Consensus struct {
	Stakers []Stake
}

// NewConsensus creates a new Consensus instance
func NewConsensus() *Consensus {
	return &Consensus{
		Stakers: make([]Stake, 0),
	}
}

// AddStaker adds a new staker to the list
func (c *Consensus) AddStaker(nodeID string, amount int64) {
	c.Stakers = append(c.Stakers, Stake{NodeID: nodeID, Amount: amount})
}

// SelectValidator selects a validator based on their stake
func (c *Consensus) SelectValidator() string {
	totalStake := int64(0)
	for _, staker := range c.Stakers {
		totalStake += staker.Amount
	}

	// Generate a random number between 0 and totalStake
	randNum, _ := rand.Int(rand.Reader, big.NewInt(totalStake))

	// Select the validator based on the random number
	cumulativeStake := int64(0)
	for _, staker := range c.Stakers {
		cumulativeStake += staker.Amount
		if randNum.Int64() < cumulativeStake {
			return staker.NodeID
		}
	}

	return ""
}

// ValidateBlock validates a block based on the consensus rules
func (c *Consensus) ValidateBlock(block *Block) bool {
	// For now, we'll just check if the block has a valid timestamp
	return block.Timestamp <= time.Now().Unix()
}
