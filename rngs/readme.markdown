Implemented RNGs:

 - des-ctr: DES operating in counter mode with key and iv set to the initial seed.
   This is a cryptographic PRNG.
 - crc64-ecma: CRC-64 with the ECMA polynomial 0xC96C5795D7870F42, operating in output-feedback mode (each generated value is written back into the hasher).
   This tends to be the slowest available PRNG that is not cryptographically secure.
 - md5: MD5 operating in output-feedback mode. This is considered a cryptographic PRNG.
 - sha1: SHA-1 operating in output-feedback mode. This is a cryptographic PRNG.
 - knuth-lcg: Knuth's 64-bit [linear congruential generator](http://en.wikipedia.org/wiki/Linear_congruential_generator), with a=6364136223846793005 and c=1442695040888963407.
   Each yielded value of this PRNG is the concatenation of the 32 most significant bits of two successive iterations of the LCG.
 - crypto-reader: This uses http://golang.org/pkg/crypto/rand/#variables and is therefore dependent upon the host operating system.
   This is considered a cryptographic RNG.
   **NOTE**: This source cannot be seeded, and so the avalanche test does not apply.
 - plan9: This uses the algorithm in Go's standard library. Fast and high quality.
   The source for this PRNG is at <http://golang.org/src/pkg/math/rand/rng.go>.
 - xer*n*: This is just [an experiment of mine](https://github.com/zephyrtronium/xer). *n* gives the state size used.

All PRNGs yield 63 (assumedly) pseudo-random bits with each iteration.

There is also code for a "reference PRNG" that yields its seed infinitely, but using it caused extreme memory usage in the avalanche and popcount tests, which use arbitrary-precision numbers to calculate the test metric. It has been disabled by default; to enable it, uncomment its line in rngs.go and rebuild.
