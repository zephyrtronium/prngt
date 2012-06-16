package rngs

import "crypto/cipher"

// CtrRNG turns a block cipher into a random number generator in counter mode.
type CtrRNG struct {
	b cipher.Block
	s cipher.Stream
}

func NewCtrRNG(c cipher.Block, seed int64) *CtrRNG {
	ctr := &CtrRNG{c, nil}
	ctr.Seed(seed)
	return ctr
}

func (ctr *CtrRNG) Seed(seed int64) {
	ctr.s = cipher.NewCTR(ctr.b, Uint64ToBytes(uint64(seed), ctr.b.BlockSize()))
}

func (ctr *CtrRNG) Int63() int64 {
	b := make([]byte, 8)
	ctr.s.XORKeyStream(b, b)
	var x int64
	for _, v := range b {
		x = x<<8 | int64(v)
	}
	return x & 0x7fffffffffffffff
}

func Uint64ToBytes(s uint64, n int) []byte {
	iv := make([]byte, n)
	for i := range iv {
		iv[i] = byte(s)
		s >>= 8
	}
	return iv
}
