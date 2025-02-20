package consensus

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	block := &Block{
		Index:     0,
		Timestamp: "2025-02-19 07:49:42",
		Data:      "Genesis Block",
		PrevHash:  "",
		Validator: "Node1",
	}

	expectedHash := "42b909559a4449ce1f0771abe0d0f5d0cf5ee960f5301bb7a9f1aeb960e88e7a"
	actualHash := block.CalculateHash()

	if actualHash != expectedHash {
		t.Errorf("CalculateHash() failed: expected %s, got %s", expectedHash, actualHash)
	}
}
