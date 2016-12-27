package main

import (
	"flag"
	"fmt"
	"sort"
)

const (
	modulus = 1e9 + 7
)

var verbose = flag.Bool("v", false, "verbose")

type state struct{ a, b, c, d int }

func (s state) String() string {
	return fmt.Sprintf("%d%d%d%d", s.a, s.b, s.c, s.d)
}

func main() {
	prime := make([]bool, 46)
	for _, p := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43} {
		prime[p] = true
	}
	transitions := make(map[state][]state)
	// Find all the sequences which are 3-digit and 4-digit prime
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				if !prime[a+b+c] {
					continue
				}
				for d := 0; d <= 9; d++ {
					if !(prime[a+b+c+d] && prime[b+c+d]) {
						continue
					}
					s := state{a, b, c, d}
					transitions[s] = nil
				}
			}
		}
	}

	// For each state, find all the possible next states
	// (i.e. find the possible five-digit sequences from this
	// four-digit sequence)
	for s := range transitions {
		for e := 0; e <= 9; e++ {
			if prime[s.c+s.d+e] &&
				prime[s.b+s.c+s.d+e] &&
				prime[s.a+s.b+s.c+s.d+e] {
				dest := state{s.b, s.c, s.d, e}
				transitions[s] = append(transitions[s], dest)
			}
		}
	}

	// How many ways can we get to this state (trailing 4 digits?)
	counts := make(map[state]int)

	for s := range transitions {
		if s.a != 0 {
			counts[s] = 1
		}
	}

	for digits := 4; digits < 20; digits++ {
		fmt.Println(digits, sum(counts))
		nextCount := make(map[state]int)
		for s, count := range counts {
			for _, dst := range transitions[s] {
				nextCount[dst] += count
			}
		}
		counts = nextCount
	}
}

func sum(counts map[state]int) int {
	total := 0
	var states []state
	for s, v := range counts {
		if *verbose && v > 0 {
			states = append(states, s)
		}
		total += v
	}
	if *verbose {
		sort.Sort(ByS(states))
		fmt.Println("###########################################\n")
		for _, s := range states {
			fmt.Println(s, counts[s])
		}
	}
	return total
}

type ByS []state

func (a ByS) Len() int           { return len(a) }
func (a ByS) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByS) Less(i, j int) bool { return a[i].String() < a[j].String() }
