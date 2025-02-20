package consensus

import (
	"github.com/NueraNetwork/nuera-chain/src/p2p"
)

// Consensus implements the Proof of Stake logic
type Consensus struct {
	Node *p2p.Node
}

// NewConsensus creates a new Consensus instance
func NewConsensus(node *p2p.Node) *Consensus {
	return &Consensus{
		Node: node,
	}
}

// BroadcastBlock broadcasts a block to all peers
func (c *Consensus) BroadcastBlock(block *Block) {
	// Serialize the block (you can use JSON, Protobuf, etc.)
	message := []byte("Block data") // Replace with actual serialization logic

	// Broadcast the message
	c.Node.BroadcastMessage(message)
}
