package consensus

import (
    "fmt"
)

// SlashValidator slashes a validator for misbehavior
func SlashValidator(validator Validator, slashingPercentage int64) {
    slashedAmount := validator.Stake * slashingPercentage / 100
    validator.Stake -= slashedAmount
    fmt.Printf("Slashed %d tokens from validator %s\n", slashedAmount, validator.Address)
}
