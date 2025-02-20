package main

import (
    "sync"
)

// Transaction represents a transaction in the blockchain
type Transaction struct {
    ID   string
    Data string
}

// TransactionPool manages pending transactions
type TransactionPool struct {
    Transactions []*Transaction
    mutex        sync.Mutex
}

// AddTransaction adds a transaction to the pool
func (tp *TransactionPool) AddTransaction(tx *Transaction) {
    tp.mutex.Lock()
    defer tp.mutex.Unlock()
    tp.Transactions = append(tp.Transactions, tx)
}

// GetTransactions retrieves all transactions from the pool
func (tp *TransactionPool) GetTransactions() []*Transaction {
    tp.mutex.Lock()
    defer tp.mutex.Unlock()
    return tp.Transactions
}

// ClearTransactions removes all transactions from the pool
func (tp *TransactionPool) ClearTransactions() {
    tp.mutex.Lock()
    defer tp.mutex.Unlock()
    tp.Transactions = []*Transaction{}
}

func main() {
    // Create a new transaction pool
    tp := &TransactionPool{}

    // Add a transaction to the pool
    tp.AddTransaction(&Transaction{ID: "tx1", Data: "Hello, World!"})

    // Retrieve transactions from the pool
    transactions := tp.GetTransactions()
    for _, tx := range transactions {
        println("Transaction ID:", tx.ID, "Data:", tx.Data)
    }

    // Clear the transaction pool
    tp.ClearTransactions()
    println("Transaction pool cleared. Number of transactions:", len(tp.GetTransactions()))
}
