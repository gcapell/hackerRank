package main

import (
	"fmt"
	"log"
	"time"
)

const (
	modulus   = 1e9 + 7
	maxDigits = 400e3
	debug     = false
)

var totals [maxDigits + 1]int

type state struct{ a, b, c, d int }

func (s state) String() string {
	return fmt.Sprintf("%d%d%d%d", s.a, s.b, s.c, s.d)
}

func main() {
	findTotals()
	var q int
	fmt.Scanln(&q)
	for j := 0; j < q; j++ {
		var d int
		fmt.Scanln(&d)
		fmt.Println(totals[d])
	}
}

func findTotals() {
	start := time.Now()
	prime := make([]bool, 46)
	for _, p := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43} {
		prime[p] = true
	}
	var prime4 []state // 4-tuples that are 3-sum-prime and 4-sum-prime
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				if !prime[a+b+c] {
					continue
				}
				for d := 0; d <= 9; d++ {
					if prime[a+b+c+d] && prime[b+c+d] {
						prime4 = append(prime4, state{a, b, c, d})
					}
				}
			}
		}
	}

	transitions := make(map[state][]state) // dst -> []src
	for _, s := range prime4 {
		for e := 0; e <= 9; e++ {
			p := prime
			if p[s.c+s.d+e] && p[s.b+s.c+s.d+e] && p[s.a+s.b+s.c+s.d+e] {
				dst := state{s.b, s.c, s.d, e}
				transitions[dst] = append(transitions[dst], s)
			}
		}
	}

	counts := make(map[state]int)
	var nextCount map[state]int
	for dst, srcs := range transitions {
		if dst.a != 0 {
			counts[dst] = 1
		}
		for _, s := range srcs {
			if s.a != 0 {
				counts[s] = 1
			}
		}
	}
	totals[1] = 9
	totals[2] = 90
	totals[3] = 303
	totals[4] = 280

	digits := 4
	for {
		nextCount = make(map[state]int)
		rd := removeEmptyDsts(transitions)
		if debug && rd > 0 {
			fmt.Printf("removed %d empty dst\n", rd)
		}
		for dst, srcs := range transitions {
			total := 0
			for _, s := range srcs {
				total += counts[s]
			}
			nextCount[dst] = total
		}
		rs := removeUselessSrcs(transitions)
		if debug && rs > 0 {
			fmt.Printf("removed %d useless src\n", rs)
		}
		counts = nextCount
		digits++
		totals[digits] = sum(counts)
		if rs == 0 && rd == 0 {
			break
		}
	}

	// Map every state to an int ID
	seq := sequencer(make(map[state]int))
	for dst, srcs := range transitions {
		seq.add(dst)
		for _, s := range srcs {
			seq.add(s)
		}
	}

	// Store counts of states indexing into an array with state's ID
	nCount := make([]int, len(seq))
	for s, val := range counts {
		nCount[seq[s]] = val
	}
	nextNCount := make([]int, len(seq))

	// replace transitions map with three transition slices (using IDs)
	var t1 [][2]int
	var t2 [][3]int
	var t3 [][4]int
	for dst, srcs := range transitions {
		switch len(srcs) {
		case 1:
			t1 = append(t1, [2]int{seq[dst], seq[srcs[0]]})
		case 2:
			t2 = append(t2, [3]int{seq[dst], seq[srcs[0]], seq[srcs[1]]})
		case 3:
			t3 = append(t3, [4]int{seq[dst], seq[srcs[0]], seq[srcs[1]], seq[srcs[2]]})
		default:
			log.Fatal(dst, srcs)
		}
	}

	for digits++; digits < len(totals); digits++ {
		total := 0
		for _, t := range t1 {
			sub := nCount[t[1]]
			nextNCount[t[0]] = sub
			total += sub
		}
		for _, t := range t2 {
			sub := (nCount[t[1]] + nCount[t[2]]) % modulus
			nextNCount[t[0]] = sub
			total += sub
		}
		for _, t := range t3 {
			sub := (nCount[t[1]] + nCount[t[2]] + nCount[t[3]]) % modulus
			nextNCount[t[0]] = sub
			total += sub
		}
		totals[digits] = total % modulus
		nCount, nextNCount = nextNCount, nCount
	}
	if debug {
		fmt.Printf("%d in %s\n", len(totals), time.Since(start))
		fmt.Println(totals[:20])
	}
}

func removeEmptyDsts(transitions map[state][]state) int {
	del := 0
	for s, v := range transitions {
		if len(v) == 0 {
			delete(transitions, s)
			del++
		}
	}
	return del
}

func removeUselessSrcs(transitions map[state][]state) int {
	del := 0
	for dst, srcs := range transitions {
		var good []state
		for _, s := range srcs {
			if _, ok := transitions[s]; ok {
				good = append(good, s)
			}
		}
		if len(good) < len(srcs) {
			transitions[dst] = good
			del += len(srcs) - len(good)
		}
	}
	return del
}

func sum(counts map[state]int) int {
	total := 0
	for _, v := range counts {
		total += v
	}

	return total
}

type sequencer map[state]int

func (m sequencer) add(s state) {
	if _, ok := m[s]; !ok {
		m[s] = len(m)
	}
}
