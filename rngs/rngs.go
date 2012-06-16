package rngs

import (
	"crypto/des"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"github.com/zephyrtronium/xer"
	"hash/crc64"
	"math/rand"
)

type RNGGetter func(seed int64) rand.Source

var RNGs = map[string]RNGGetter{
	"plan9":     rand.NewSource,
	"xer256":    func(seed int64) rand.Source { return xer.New(seed, 256) },
	"xer312":    func(seed int64) rand.Source { return xer.New(seed, 312) },
	"xer65536":  func(seed int64) rand.Source { return xer.New(seed, 65536) },
	"knuth-lcg": NewKnuthLCG,
	"des-ctr": func(seed int64) rand.Source {
		ctr, _ := des.NewCipher(Uint64ToBytes(uint64(seed), 8))
		return NewCtrRNG(ctr, seed)
	},
	"crypto-reader": func(int64) rand.Source { return ReaderSource{crand.Reader} },
	//"repeat":        func(seed int64) rand.Source { return &Repeater{seed & 0x7fffffffffffffff} },
	"crc64-ecma": func(seed int64) rand.Source { return &HashOFB{crc64.New(crc64.MakeTable(crc64.ECMA))} },
	"md5":        func(seed int64) rand.Source { return &HashOFB{md5.New()} },
	"sha1":       func(seed int64) rand.Source { return &HashOFB{sha1.New()} },
}

var CryptoRNGs = map[string]bool{
	"des-ctr":       true,
	"crypto-reader": true,
	"md5":           true,
	"sha1":          true,
}
