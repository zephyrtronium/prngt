package main

import (
	"github.com/zephyrtronium/xer"
	"math/rand"
)

type RNGGetter func(seed int64) rand.Source

var RNGs = map[string]RNGGetter{
	"plan9":     rand.NewSource,
	"xer256":    func(seed int64) rand.Source { return xer.New(seed, 256) },
	"xer312":    func(seed int64) rand.Source { return xer.New(seed, 312) },
	"xer65536":  func(seed int64) rand.Source { return xer.New(seed, 65536) },
	"knuth-lcg": NewKnuthLCG,
}
