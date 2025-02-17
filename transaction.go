package main

import (
    "encoding/hex"
    "hash"
)

// Transaction represents a blockchain transaction
type Transaction struct {
    From      string
    To        string
    Amount    float64
    Signature []byte
}

// Hash computes the hash of a transaction
func (tx Transaction) Hash() []byte {
    // Dummy hash function, replace with actual hash computation
    h := hash.New()
    h.Write([]byte(tx.From + tx.To + string(tx.Amount)))
    return h.Sum(nil)
}

// ValidateTransaction checks if a transaction is valid
func ValidateTransaction(tx Transaction) bool {
    if !VerifySignature(hexToPoint(tx.From), tx.Signature, tx.Hash()) {
        return false
    }
    
    // Here you would check for sufficient balance, but for simplicity, we'll skip this.
    return true
}

// Helper function to convert hex string to kyber.Point
func hexToPoint(hexStr string) kyber.Point {
    b, _ := hex.DecodeString(hexStr)
    point := suite.Point()
    point.UnmarshalBinary(b)
    return point
}
