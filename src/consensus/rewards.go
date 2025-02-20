package consensus

import (
    "fmt"
)

// CalculateRewards calculates rewards for a validator
func CalculateRewards(validator Validator) int64 {
    // Example reward calculation logic
    reward := validator.Stake / 100 // 1% of the validator's stake
    fmt.Printf("Calculating rewards for validator %s: %d\n", validator.Address, reward)
    return reward
}

// DistributeRewards distributes rewards across all validators
func DistributeRewards(validators []Validator) {
    for _, validator := range validators {
        reward := CalculateRewards(validator)
        fmt.Printf("Distributed %d rewards to validator %s\n", reward, validator.Address)
    }
}
