package tests

import "math/rand"

// Statistical test of a random number generator. The returned value is the
// variance from the expected distribution in [0, ∞) with smaller values
// usually considered "better."
type Test func(rand.Source) float64

var Tests = map[string]Test{
	"survival":        NewLaggedSurvival(0),
	"survival-parity": SurvivalParity,
	"popcount":        Popcount,
	"survival-lag1":   NewLaggedSurvival(1),
	"survival-lag2":   NewLaggedSurvival(2),
	"survival-lag5":   NewLaggedSurvival(5),
	"survival-lag11":  NewLaggedSurvival(11),
	"survival-lag16":  NewLaggedSurvival(16),
	"avalanche":       Avalanche,
}
