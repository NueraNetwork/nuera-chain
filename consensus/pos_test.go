
package consensus

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"testing"
	"github.com/NueraNetwork/nuera-chain/types"
)

func TestSelectValidator(t *testing.T) {
	// Create a list of stakeholders
	stakeholders := []types.Stakeholder{
		{Address: "addr1", Stake: big.NewInt(100)},
		{Address: "addr2", Stake: big.NewInt(200)},
		{Address: "addr3", Stake: big.NewInt(300)},
	}

	// Test selecting a validator
	validator, err := SelectValidator(stakeholders)
	if err != nil {
		t.Fatalf("Failed to select validator: %v", err)
	}

	if validator == nil {
		t.Fatal("No validator selected")
	}

	t.Logf("Selected validator: %s", validator.Address)
}

func TestValidateBlock(t *testing.T) {
	// Generate a private key for the validator
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Create a mock block
	block := types.Block{
		Index:        0,
		Timestamp:    1234567890,
		Transactions: []types.Transaction{},
		PrevHash:     "prevhash",
		Hash:         "blockhash",
		Signature:    nil, // Will be set after signing
		Validator:    "validator",
	}

	// Hash the block data
	hash := sha256.Sum256([]byte(block.Hash))

	// Sign the block hash
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		t.Fatalf("Failed to sign block: %v", err)
	}
	block.Signature = signature

	// Create a mock stakeholder (validator)
	validator := types.Stakeholder{
		Address:    "validator",
		Stake:      big.NewInt(100),
		PrivateKey: nil, // Not needed for validation
		PublicKey:  elliptic.Marshal(privateKey.PublicKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y),
	}

	// Test validating the block
	valid := ValidateBlock(block, validator)
	if !valid {
		t.Fatal("Block validation failed")
	}

	t.Log("Block validation succeeded")
}
