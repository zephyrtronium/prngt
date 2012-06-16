package rngs

type Repeater struct {
	V int64
}

func (r *Repeater) Seed(seed int64) { r.V = seed & 0x7fffffffffffffff }
func (r *Repeater) Int63() int64    { return r.V }
