package rngs

import "hash"

// HashOFB feeds the output from a hash back into itself to produce values.
type HashOFB struct {
	hash.Hash
}

func (h *HashOFB) Seed(seed int64) {
	b := Uint64ToBytes(uint64(seed), 8)
	h.Reset()
	h.Write(b)
}

func (h *HashOFB) Int63() (x int64) {
	n := 0
	b := make([]byte, 0, h.Size())
	for n < 8 {
		b = h.Sum(b[:0])
		h.Write(b)
		n += len(b)
		for _, v := range b {
			x = x<<8 | int64(v)
		}
	}
	return
}
