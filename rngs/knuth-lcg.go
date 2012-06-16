package rngs

import "math/rand"

type KnuthLCG struct {
	x int64
}

func NewKnuthLCG(seed int64) rand.Source {
	return &KnuthLCG{seed}
}

func (r *KnuthLCG) Seed(seed int64) {
	r.x = seed
}

func (r *KnuthLCG) Int63() (s int64) {
	const a int64 = 6364136223846793005
	const c int64 = 1442695040888963407
	r.x = a*r.x + c
	s = r.x >> 32
	r.x = a*r.x + c
	s |= int64(r.x & 0x7fffffff00000000)
	return
}
