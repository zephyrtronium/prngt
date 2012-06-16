package rngs

import "io"

type ReaderSource struct {
	R io.Reader
}

func (r ReaderSource) Int63() int64 {
	var b [8]byte
	r.R.Read(b[:]) // panicking would suck, so ignore errors
	var x int64
	for _, v := range b {
		x = x<<8 | int64(v)
	}
	return x & 0x7fffffffffffffff
}

// This is a no-op.
func (r ReaderSource) Seed(int64) {}
