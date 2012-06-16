package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/zephyrtronium/prngt/rngs"
	"github.com/zephyrtronium/prngt/tests"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Usage = func() {
		x := make([]string, 0, len(tests.Tests))
		for i := range tests.Tests {
			x = append(x, i)
		}
		y := make([]string, 0, len(rngs.RNGs))
		for i := range rngs.RNGs {
			y = append(y, i)
		}
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nAvailable tests:")
		fmt.Fprintln(os.Stderr, strings.Join(x, ", "))
		fmt.Fprintln(os.Stderr, "\nAvailable RNGs:")
		fmt.Fprintln(os.Stderr, strings.Join(y, ", "))
	}
}

var waitgroup sync.WaitGroup

func main() {
	var rng, test string
	var seed uint64
	var nocrypto bool
	flag.StringVar(&rng, "rng", "", "comma-separated list of RNGs to test (default all)")
	flag.StringVar(&test, "test", "", "comma-separated list of tests to run (default all)")
	flag.Uint64Var(&seed, "seed", ^uint64(0), "initial seed (default random)")
	flag.BoolVar(&nocrypto, "no-crypto", false, "skip cryptographic RNGs (overridden by -rng)")
	flag.Parse()

	if seed == ^uint64(0) {
		binary.Read(crand.Reader, binary.LittleEndian, &seed)
	}
	rngm := make(map[string]rngs.RNGGetter)
	testm := make(map[string]tests.Test)
	if rng == "" {
		if nocrypto {
			for nm, r := range rngs.RNGs {
				if !rngs.CryptoRNGs[nm] {
					rngm[nm] = r
				}
			}
		} else {
			rngm = rngs.RNGs
		}
	} else {
		for _, nm := range strings.Split(rng, ",") {
			f, ok := rngs.RNGs[nm]
			if !ok {
				fmt.Fprintln(os.Stderr, "unknown rng", rng)
			}
			rngm[nm] = f
		}
	}
	if test == "" {
		testm = tests.Tests
	} else {
		for _, nm := range strings.Split(test, ",") {
			t, ok := tests.Tests[test]
			if !ok {
				fmt.Fprintln(os.Stderr, "unknown test", test)
			}
			testm[nm] = t
		}
	}
	if len(rngm) == 0 {
		fmt.Fprintln(os.Stderr, "no rngs to test")
		os.Exit(1)
	}
	if len(testm) == 0 {
		fmt.Fprintln(os.Stderr, "no tests to run")
		os.Exit(1)
	}

	fmt.Printf("Using seed %#016x\n", seed)
	fmt.Println("RNG  \t\tTest \t\tMetric\t\tTime ")
	start := time.Now()
	for rname, r := range rngm {
		for tname, t := range testm {
			waitgroup.Add(1)
			go doTest(r(int64(seed)), t, rname, tname)
		}
	}
	waitgroup.Wait()
	fmt.Println("Total elapsed time:", time.Since(start))
}

func doTest(rng rand.Source, test tests.Test, rname, tname string) {
	start := time.Now()
	variance := test(rng)
	duration := time.Since(start)
	fmt.Printf("%s\t%s\t%f\t%v\n", rname, tname, variance, duration)
	waitgroup.Done()
}
