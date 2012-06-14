package main

import "math/rand"

const survivalN = 1000000

// The survival test measures through how many iterations each bit in the
// sequence "survives", or fails to change. If the sequence is uniformly
// distributed, then survival times should be exponentially distributed.
func Survival(r rand.Source) float64 {
	consec := [63]int{}
	bits := [63]bool{}
	counts := make(map[int]int)
	for i := 0; i < survivalN; i++ {
		x := r.Int63()
		for b := 0; b < 63; b++ {
			if ((x & (1 << uint(b))) != 0) == bits[b] { // bit survived
				consec[b]++
			} else { // bit changed
				bits[b] = !bits[b]
				counts[consec[b]] = counts[consec[b]] + 1
				consec[b] = 0
			}
		}
	}
	var maximum int
	var mean float64
	for i, v := range counts {
		if i > maximum {
			maximum = i
		}
		mean += float64(v)
	}
	mean /= float64(maximum)
	// If the source is truly random, then there should be half as many hits
	// for counts[n] as there were for counts[n-1], with counts[0] being the
	// maximum.
	//TODO: E should be calculated from the median
	E := float64(counts[0])
	var s2n float64
	for i := 1; i < maximum; i++ {
		E /= 2
		d := float64(counts[i]) - E
		s2n += d*d
	}
	return s2n / float64(maximum) / mean // index of dispersion
}
