package consensus

// Validator represents a node that participates in consensus
type Validator struct {
    Address string
    Stake   int64
}

// TotalSupply represents the total supply of tokens in the network
var TotalSupply int64 = 10000000 // Example value
