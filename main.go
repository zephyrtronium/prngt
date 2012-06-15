package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, `messages look like
<test> testing <rng>	σ²/μ=<value>	took <time>
σ²/μ is index of dispersion. Lower values indicate higher-quality randomness,
but if a generator never produces high values, then it isn't random.

Available tests:`)
		for i := range Tests {
			fmt.Fprint(os.Stderr, i, " ")
		}
		fmt.Fprintln(os.Stderr, "\n\nAvailable RNGs:")
		for i := range RNGs {
			fmt.Fprint(os.Stderr, i, " ")
		}
		fmt.Fprintln(os.Stderr)
	}
}

var waitgroup sync.WaitGroup

func main() {
	var rng, test string
	var seed uint64
	flag.StringVar(&rng, "rng", "", "select an RNG to test (default all)")
	flag.StringVar(&test, "test", "", "select a test to run (default all)")
	flag.Uint64Var(&seed, "seed", ^uint64(0), "initial seed (default random)")
	flag.Parse()

	if seed == ^uint64(0) {
		binary.Read(crand.Reader, binary.LittleEndian, &seed)
	}
	var rngs map[string]RNGGetter
	var tests map[string]Test
	if rng == "" {
		rngs = RNGs
	} else {
		f, ok := RNGs[rng]
		if !ok {
			fmt.Fprintln(os.Stderr, "unknown rng", rng)
			os.Exit(1)
		}
		rngs = map[string]RNGGetter{rng: f}
	}
	if test == "" {
		tests = Tests
	} else {
		t, ok := Tests[test]
		if !ok {
			fmt.Fprintln(os.Stderr, "unknown test", test)
			os.Exit(1)
		}
		tests = map[string]Test{test: t}
	}

	fmt.Printf("Using seed %#016x\n", seed)
	start := time.Now()
	for rname, r := range rngs {
		for tname, t := range tests {
			waitgroup.Add(1)
			go doTest(r(int64(seed)), t, rname, tname)
		}
	}
	waitgroup.Wait()
	fmt.Println("Total elapsed time:", time.Since(start))
}

func doTest(rng rand.Source, test Test, rname, tname string) {
	start := time.Now()
	variance := test(rng)
	duration := time.Since(start)
	fmt.Printf("%s testing %s:	σ²/μ=%f	took %v\n", tname, rname, variance, duration)
	waitgroup.Done()
}
