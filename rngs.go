package main

import (
	"github.com/zephyrtronium/xer"
	"math/rand"
)

var RNGs = map[string]func(seed int64) rand.Source{
	"plan9":    rand.NewSource,
	"xer256":   func(seed int64) rand.Source { return xer.New(seed, 256) },
	"xer65536": func(seed int64) rand.Source { return xer.New(seed, 65536) },
}
