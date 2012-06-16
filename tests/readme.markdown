Implemented tests:

 - avalanche: The Hamming distances between the nth values of one RNG seeded with two values whose Hamming distance is 1 should follow a binomial distribution.
   This tests the [bit avalanche](http://en.wikipedia.org/wiki/Avalanche_effect).
   **NOTE**: This test tends to be _very_ slow, as it has to seed each RNG twice in each of its hundred thousand iterations, and seeding algorithms can be very slow. xer65536 took just under three hours on my system.
   **NOTE**: The crypto-reader RNG does not operate properly under this test, as seeding is a no-op.
 - survival: The numbers of iterations of an RNG through which the bit in each position remains the same should follow an exponential distribution.
   This is a form of the [next-bit test](http://en.wikipedia.org/wiki/Next-bit_test).
 - survival-lag*n*: Same as survival, but with *n* iterations of the RNG ignored between samples.
   This test will later be replaced with lagging on RNGs.
 - survival-parity: The numbers of iterations through which the parity of values from an RNG remains the same should follow an exponential distribution.
   This is another form of the next-bit test, applied norizontally rather than vertically.

The metric currently returned by tests is one such that lower values usually mean higher-quality randomness. Later this metric will be replaced with p-values (or possibly another goodness of fit metric).
