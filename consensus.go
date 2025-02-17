package main

import (
    "math/big"
    "hash"
)

// ProofOfWork computes the nonce that results in a hash below the target difficulty
func ProofOfWork(block Block) ([]byte, int) {
    var nonce int
    var hashInt big.Int
    target := big.NewInt(1)
    target.Lsh(target, 256-16) // 16 is the difficulty level

    for nonce < 1000000 { // Arbitrary limit to prevent infinite loop
        data := block.PrepareData(nonce)
        hash := hash.Sum(data)
        
        hashInt.SetBytes(hash)
        if hashInt.Cmp(target) == -1 {
            return hash, nonce
        }
        nonce++
    }
    return nil, -1 // Indicates failure if no solution found within iteration limit
}

// Block is a simplified blockchain block
type Block struct {
    Index        int
    Transactions []Transaction
    Nonce        int
}

// PrepareData prepares the data for hashing, including the nonce
func (b Block) PrepareData(nonce int) []byte {
    // Dummy preparation, replace with actual data preparation
    return []byte(string(b.Index) + string(nonce))
}
