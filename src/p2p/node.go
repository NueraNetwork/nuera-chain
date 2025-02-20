package p2p

import (
	"context"
	"log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// Node represents a P2P node
type Node struct {
	Host host.Host
}

// NewNode creates a new P2P node
func NewNode(listenAddr string) *Node {
	// Create a new libp2p host
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(listenAddr),
	)
	if err != nil {
		log.Fatalf("Failed to create P2P node: %v", err)
	}

	log.Printf("Node started with ID: %s\n", h.ID())

	return &Node{
		Host: h,
	}
}

// StartServer starts the P2P server
func (n *Node) StartServer() {
	log.Printf("Node listening on %s\n", n.Host.Addrs())
}

// ConnectToPeer connects to a peer
func (n *Node) ConnectToPeer(peerAddr string) {
	// Parse the peer address
	peerInfo, err := peer.AddrInfoFromString(peerAddr)
	if err != nil {
		log.Fatalf("Failed to parse peer address: %v", err)
	}

	// Connect to the peer
	err = n.Host.Connect(context.Background(), *peerInfo)
	if err != nil {
		log.Fatalf("Failed to connect to peer: %v", err)
	}

	log.Printf("Connected to peer: %s\n", peerInfo.ID)
}

// PeerDiscovery sets up peer discovery using mDNS
func (n *Node) PeerDiscovery() {
	// Set up mDNS service discovery
	service := mdns.NewMdnsService(n.Host, "nuera-network", n)
	err := service.Start()
	if err != nil {
		log.Fatalf("Failed to start mDNS service: %v", err)
	}
}

// HandlePeerFound is called when a new peer is discovered
func (n *Node) HandlePeerFound(peerInfo peer.AddrInfo) {
	log.Printf("Discovered peer: %s\n", peerInfo.ID)
}

// BroadcastMessage sends a message to all connected peers
func (n *Node) BroadcastMessage(message []byte) {
	// Get the list of connected peers
	peers := n.Host.Network().Peers()

	// Send the message to each peer
	for _, peerID := range peers {
		stream, err := n.Host.NewStream(context.Background(), peerID, "/nuera/1.0.0")
		if err != nil {
			log.Printf("Failed to open stream to peer %s: %v\n", peerID, err)
			continue
		}

		_, err = stream.Write(message)
		if err != nil {
			log.Printf("Failed to send message to peer %s: %v\n", peerID, err)
		}

		stream.Close()
	}
}
