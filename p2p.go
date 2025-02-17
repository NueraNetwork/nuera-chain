package main

import (
    "net"
    "log"
)

// Node represents a node in the network
type Node struct {
    address string
}

// DiscoverPeers attempts to connect to known bootstrap nodes
func DiscoverPeers() {
    bootstrapNodes := []string{"127.0.0.1:3000", "127.0.0.1:3001"}
    for _, addr := range bootstrapNodes {
        err := ConnectToPeer(addr)
        if err != nil {
            log.Printf("Failed to connect to peer %s: %v", addr, err)
        }
    }
}

// ConnectToPeer connects to a peer 
func ConnectToPeer(address string) error {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return err
    }
    defer conn.Close()
    
    log.Printf("Connected to peer: %s", address)
    return nil
}

func main() {
    // Example usage
    node := Node{address: "127.0.0.1:3000"}
    DiscoverPeers()
}
