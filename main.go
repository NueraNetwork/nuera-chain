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

	// Create a new P2P node
	node := p2p.NewNode("/ip4/127.0.0.1/tcp/" + *port)

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

	// Create a new blockchain with the public key
	blockchain := consensus.NewBlockchain(publicKey)

	// Add some blocks to the blockchain
	blockchain.AddBlock("First Block", privateKey)
	blockchain.AddBlock("Second Block", privateKey)
	blockchain.AddBlock("Third Block", privateKey)

	// TEMPORARY: Add a fourth block for testing block propagation
	blockchain.AddBlock("Fourth Block", privateKey)

	// Validate the blockchain
	if blockchain.ValidateBlockchain() {
		log.Println("Blockchain is valid!")
	} else {
		log.Println("Blockchain is invalid!")
	}

	// Log the blockchain
	for _, block := range blockchain.Blocks {
		log.Printf("Block %d: %s (Validator: %s)\n", block.Index, block.Hash, block.Validator)
	}

	// Keep the program running
	log.Println("Node started successfully. Waiting for connections...")
	select {}
}
