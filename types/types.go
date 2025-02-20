package types

import (
	"math/big"
)

// Block represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Signature    []byte
	Validator    string
}

// Transaction represents a transaction in the blockchain
type Transaction struct {
	Sender    string
	Receiver  string
	Amount    int
	Signature []byte
}

// Stakeholder represents a node participating in the PoS consensus
type Stakeholder struct {
	Address    string
	Stake      *big.Int // Amount of tokens staked
	PrivateKey []byte   // Private key for signing blocks
	PublicKey  []byte   // Public key for verification
}
