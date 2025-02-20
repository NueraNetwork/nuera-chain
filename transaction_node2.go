package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"time" // Add this line
)

// Transaction represents a basic transaction
type Transaction struct {
	Sender    string  `json:"sender"`    // Sender's address
	Receiver  string  `json:"receiver"`  // Receiver's address
	Amount    float64 `json:"amount"`    // Amount to send
	Signature []byte  `json:"signature"` // Digital signature
}

// Sign signs the transaction with the sender's private key
func (t *Transaction) Sign(privKey *ecdsa.PrivateKey) error {
	// Create a copy of the transaction without the signature
	txCopy := &Transaction{
		Sender:   t.Sender,
		Receiver: t.Receiver,
		Amount:   t.Amount,
	}

	// Convert the transaction copy to a JSON string
	txData, err := json.Marshal(txCopy)
	if err != nil {
		return err
	}

	// Hash the transaction data
	hash := sha256.Sum256(txData)
	fmt.Printf("Transaction Hash (for signing): %x\n", hash)

	// Sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return err
	}

	// Store the signature
	t.Signature = append(r.Bytes(), s.Bytes()...)
	fmt.Printf("Signature: %x\n", t.Signature)
	return nil
}

// Verify verifies the transaction's signature
func (t *Transaction) Verify(pubKey *ecdsa.PublicKey) bool {
	// Create a copy of the transaction without the signature
	txCopy := &Transaction{
		Sender:   t.Sender,
		Receiver: t.Receiver,
		Amount:   t.Amount,
	}

	// Convert the transaction copy to a JSON string
	txData, err := json.Marshal(txCopy)
	if err != nil {
		fmt.Println("Error marshaling transaction for verification:", err)
		return false
	}

	// Hash the transaction data
	hash := sha256.Sum256(txData)
	fmt.Printf("Transaction Hash (for verification): %x\n", hash)

	// Split the signature into r and s
	r := new(big.Int).SetBytes(t.Signature[:len(t.Signature)/2])
	s := new(big.Int).SetBytes(t.Signature[len(t.Signature)/2:])
	fmt.Printf("Signature R: %x\n", r.Bytes())
	fmt.Printf("Signature S: %x\n", s.Bytes())

	// Verify the signature
	valid := ecdsa.Verify(pubKey, hash[:], r, s)
	fmt.Printf("Signature Valid: %v\n", valid)
	return valid
}

// GeneratePrivateKey generates a sample ECDSA private key
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// GetPublicKey returns the public key from a private key
func GetPublicKey(privKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return &privKey.PublicKey
}

// BroadcastTransaction sends the transaction to all connected peers
func BroadcastTransaction(tx *Transaction, node *Node) {
	// Convert the transaction to JSON
	txData, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("Error marshaling transaction:", err)
		return
	}

	// Broadcast the transaction to all peers
	node.BroadcastMessage(string(txData))
}

// ValidateTransaction checks if the transaction is valid
func ValidateTransaction(tx *Transaction, pubKey *ecdsa.PublicKey) bool {
	// Check if the sender and receiver are valid addresses
	if tx.Sender == "" || tx.Receiver == "" {
		return false
	}

	// Check if the amount is positive
	if tx.Amount <= 0 {
		return false
	}

	// Verify the signature
	return tx.Verify(pubKey)
}

func main() {
	// Example usage
	fmt.Println("Creating a transaction...")

	// Initialize a P2P node
	node := NewNode("127.0.0.1:8081") // Second node listens on 127.0.0.1:8081
	go node.StartServer()

	// Add a delay to ensure the first node is ready
	fmt.Println("Waiting for the first node to initialize...")
	time.Sleep(5 * time.Second) // Wait for 5 seconds

	// Connect to the first node
	fmt.Println("Connecting to the first node...")
	node.ConnectToPeer("127.0.0.1:8080")

	// Generate a sample private key
	privKey, err := GeneratePrivateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Create a new transaction
	tx := &Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10.5,
	}

	// Sign the transaction
	err = tx.Sign(privKey)
	if err != nil {
		fmt.Println("Error signing transaction:", err)
		return
	}

	// Validate the transaction
	pubKey := GetPublicKey(privKey)
	if ValidateTransaction(tx, pubKey) {
		fmt.Println("Transaction is valid!")
	} else {
		fmt.Println("Transaction is invalid!")
		return
	}

	// Broadcast the transaction
	BroadcastTransaction(tx, node)
}
