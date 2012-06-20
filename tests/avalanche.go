package tests

import "math/rand"

const avalancheN = 100000 // this test can take VERY long with slow seeding
const avalancheFuse = 0

// The Hamming distances between the nth values of one RNG seeded with two
// values with Hamming distance 1 should be binomially distributed. This test
// derives its name from the bit avalanche effect.
//
// NOTE: ReaderSource RNGs (e.g. crypto-reader) will be incorrect.
func Avalanche(r rand.Source) float64 {
	// hax teh copypasta from popcount
	counts := [64]int64{}
	v := r.Int63()
	var a, b int64
	for i := 0; i < avalancheN; i++ {
		for j := uint(0); j < 63; j++ {
			r.Seed(v)
			for k := 0; k < avalancheFuse; k++ {
				r.Int63()
			}
			a = r.Int63()
			r.Seed(v ^ (1 << j))
			for k := 0; k < avalancheFuse; k++ {
				r.Int63()
			}
			b = r.Int63()
			// Hamming distance from a to b is equal to popcount(a^b)
			x := a ^ b
			// popcount_3() from http://en.wikipedia.org/wiki/Hamming_weight
			x -= x >> 1 & 0x5555555555555555
			x = x&0x3333333333333333 + x>>2&0x3333333333333333
			x = (x + x>>4) & 0x0f0f0f0f0f0f0f0f
			x = x * 0x0101010101010101 >> 56 // popcount
			counts[x]++
		}
		v = b // silly to use a^b because that depends upon what we're testing
	}
	var chi2 float64
	for i, v := range counts {
		d := float64(v)*binomialBitScale - binomialBitPMF[i]
		chi2 += d*d / binomialBitPMF[i]
	}
	return 1 - LowerGamma(30, chi2/2)*1.13099628864477e-31
}
