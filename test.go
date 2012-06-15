package main

import "math/rand"

// Statistical test of a random number generator. The returned value is the
// variance from the expected distribution in [0, ∞) with smaller values
// considered "better."
type Test func(rand.Source) float64

var Tests = map[string]Test{
	"survival":        Survival,
	"survival-parity": SurvivalParity,
	"popcount":        Popcount,
}
