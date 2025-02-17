package main

import (
    "flag"
    "log"
    "github.com/NueraNetwork/nuera-chain/src/consensus"
    "github.com/NueraNetwork/nuera-chain/src/crypto"
    "github.com/NueraNetwork/nuera-chain/src/p2p"
)

func main() {
    // Parse command-line flags
    port := flag.String("port", "8080", "Port to listen on")
    peer := flag.String("peer", "", "Address of peer to connect to")
    flag.Parse()

    // Create a new node
    node := p2p.NewNode("127.0.0.1:" + *port)

    // Start the server
    go node.StartServer()

    // Connect to a peer if specified
    if *peer != "" {
        node.ConnectToPeer(*peer)
    }

    // Start peer discovery
    go node.PeerDiscovery()

    // Generate a key pair for the node
    privateKey, publicKey := crypto.GenerateKeyPair()

    // Create a new blockchain
    blockchain := consensus.NewBlockchain()

    // Add some blocks to the blockchain
    blockchain.AddBlock("First Block", privateKey)
    blockchain.AddBlock("Second Block", privateKey)
    blockchain.AddBlock("Third Block", privateKey)

    // Validate the blockchain
    if blockchain.ValidateBlockchain(publicKey) {
        log.Println("Blockchain is valid!")
    } else {
        log.Println("Blockchain is invalid!")
    }

    // Log the blockchain
    for _, block := range blockchain.Blocks {
        log.Printf("Block %d: %s\n", block.Index, block.Hash)
    }

    // Create a new consensus instance
    consensus := consensus.NewConsensus(node)

    // Broadcast the latest block
    latestBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    consensus.BroadcastBlock(latestBlock)

    // Keep the program running
    log.Println("Node started successfully. Waiting for connections...")
    select {}
}
