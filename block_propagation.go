package main

import (
    "context"
    "fmt"
    "sync"

    "github.com/libp2p/go-libp2p"
    "github.com/libp2p/go-libp2p/core/host"
    "github.com/libp2p/go-libp2p/core/peer"
    "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// Block represents a block in the blockchain
type Block struct {
    Data         string
    Transactions []*Transaction
}

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

// Blockchain represents the blockchain structure
type Blockchain struct {
    Blocks          []*Block
    TransactionPool *TransactionPool
    mutex           sync.Mutex
}

// AddBlock adds a new block to the blockchain and broadcasts it
func (bc *Blockchain) AddBlock(data string) {
    bc.mutex.Lock()
    defer bc.mutex.Unlock()

    // Get transactions from the pool
    transactions := bc.TransactionPool.GetTransactions()

    // Create a new block with the transactions
    newBlock := &Block{
        Data:         data,
        Transactions: transactions,
    }

    // Add the block to the blockchain
    bc.Blocks = append(bc.Blocks, newBlock)

    // Broadcast the block to peers
    bc.BroadcastBlock(newBlock)

    // Clear the transaction pool
    bc.TransactionPool.ClearTransactions()
}

// BroadcastBlock simulates broadcasting a block to all peers
func (bc *Blockchain) BroadcastBlock(block *Block) {
    fmt.Printf("Broadcasting block: %s\n", block.Data)
    fmt.Println("Transactions in block:")
    for _, tx := range block.Transactions {
        fmt.Printf("  Transaction ID: %s, Data: %s\n", tx.ID, tx.Data)
    }
}

// DiscoveryNotifee handles peer discovery events
type DiscoveryNotifee struct {
    h host.Host
}

// HandlePeerFound is called when a new peer is discovered
func (n *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
    fmt.Printf("Discovered new peer: %s\n", pi.ID)
    err := n.h.Connect(context.Background(), pi)
    if err != nil {
        fmt.Printf("Failed to connect to peer %s: %v\n", pi.ID, err)
    }
}

func main() {
    // Create a new libp2p host
    h, err := libp2p.New()
    if err != nil {
        panic(err)
    }

    // Print the host's peer ID and addresses
    fmt.Printf("Host created with ID %s\n", h.ID())

    // Set up a simple blockchain with a transaction pool
    bc := &Blockchain{
        TransactionPool: &TransactionPool{},
    }

    // Add some transactions to the pool
    bc.TransactionPool.AddTransaction(&Transaction{ID: "tx1", Data: "Hello, World!"})
    bc.TransactionPool.AddTransaction(&Transaction{ID: "tx2", Data: "Another transaction"})

    // Add a block to the blockchain (this will include the transactions)
    bc.AddBlock("Genesis Block")

    // Set up peer discovery using mDNS
    notifee := &DiscoveryNotifee{h: h}
    mdnsService := mdns.NewMdnsService(h, "nuera-network", notifee)

    // Start the mDNS service
    defer mdnsService.Close()

    // Keep the program running
    fmt.Println("Listening for peers...")
    select {}
}
