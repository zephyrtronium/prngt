package tests

import "math"

func LowerGamma(a, x float64) (z float64) {
	// x**a Γ(a) e**-x Σ{k=0..∞}x**k/Γ(a+k+1)
	const ε = 1e-20
	if x == 0 {
		return 0 // γ(a, x) is an integral from 0 to x
	}
	d := math.Gamma(a)
	m := math.Pow(x, a) * d * math.Exp(-x)
	if m == 0 {
		// overflow
		return d // lim{x→∞}γ(a, x) = Γ(a)
	}
	s := 1 / (d * a) // x**0 / Γ(a+0+1)
	z = s
	for k := a + 1.0; s > ε; k += 1.0 {
		s *= x / k
		z += s
	}
	return m * z
}
