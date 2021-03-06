package tests

import (
	"math"
	"math/rand"
)

const survivalParityN = laggedSurvN * 31

// Survival Parity applies Survival to the parity of the random values.
func SurvivalParity(r rand.Source) float64 {
	consec := 0
	var parity int64 = -1 // prevent the first iteration from matching
	counts := map[int]int{0: -1}
	for i := 0; i < survivalParityN; i++ {
		// software parity because I do not enjoy setting up asm to be used
		// http://www-graphics.stanford.edu/~seander/bithacks.html#ParityMultiply
		x := r.Int63()
		x ^= x >> 1
		x ^= x >> 2
		x = (x & 0x1111111111111111) * 0x1111111111111111
		x = x >> 60 & 1 // parity complete
		if x == parity {
			consec++
		} else {
			parity = x
			counts[consec] = counts[consec] + 1
			consec = 0
		}
	}
	// copypasta
	var maximum int
	for i := range counts {
		if i > maximum {
			maximum = i
		}
	}
	// If the source is truly random, then there should be half as many hits
	// for counts[n] as there were for counts[n-1], with counts[0] being the
	// maximum.
	//TODO: E should be calculated from the median. This is causing crc64-ecma to NaN.
	E := float64(counts[0])
	var chi2 float64
	for i := 1; i < maximum; i++ {
		E /= 2
		d := float64(counts[i]) - E
		chi2 += d * d / E
	}
	k_2 := float64(maximum-2) / 2
	return 1 - LowerGamma(k_2, chi2/2)/math.Gamma(k_2)
}
