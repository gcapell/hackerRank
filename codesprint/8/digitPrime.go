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

var totals [maxDigits]int

func main() {
	findTotals()

	var q int
	fmt.Scan(&q)
	for j := 0; j < q; j++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(totals[n])
	}
}

func findTotals() {
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

	stateNum := sequencer(make(map[state]int))
	for dst, srcs := range transitions {
		stateNum.add(dst)
		for _, s := range srcs {
			stateNum.add(s)
		}
	}

	var m1 [][2]int
	var m2 [][3]int
	var m3 [][4]int
	var m4 [][5]int
	var m5 [][6]int
	sn := stateNum
	for dst, srcs := range transitions {
		switch len(srcs) {
		case 1:
			m1 = append(m1, [2]int{sn[dst], sn[srcs[0]]})
		case 2:
			m2 = append(m2, [3]int{sn[dst], sn[srcs[0]], sn[srcs[1]]})
		case 3:
			m3 = append(m3, [4]int{sn[dst], sn[srcs[0]], sn[srcs[1]], sn[srcs[2]]})
		case 4:
			m4 = append(m4, [5]int{sn[dst], sn[srcs[0]], sn[srcs[1]], sn[srcs[2]], sn[srcs[3]]})
		case 5:
			m5 = append(m5, [6]int{sn[dst], sn[srcs[0]], sn[srcs[1]], sn[srcs[2]], sn[srcs[3]], sn[srcs[4]]})
		default:
			log.Fatal("too many", len(srcs))
		}
	}

	s1 := make([]int, len(stateNum))
	for s, n := range stateNum {
		if s.a != 0 {
			s1[n] = 1
		}
	}
	s2 := make([]int, len(stateNum))

	totals[1] = 9 // Cheated and copied these.
	totals[2] = 90
	totals[3] = 303
	totals[4] = 280

	start := time.Now()
	for k := 5; k < len(totals); k++ {
		total := 0
		if debug {
			fmt.Println(s1)
		}
		for _, m := range m1 {
			s2[m[0]] = s1[m[1]]
			total += s1[m[1]]
		}
		for _, m := range m2 {
			i := s1[m[1]] + s1[m[2]]
			s2[m[0]] = i % modulus
			total += i
		}
		for _, m := range m3 {
			i := s1[m[1]] + s1[m[2]] + s1[m[3]]
			s2[m[0]] = i % modulus
			total += i
		}
		for _, m := range m4 {
			i := s1[m[1]] + s1[m[2]] + s1[m[3]] + s1[m[4]]
			s2[m[0]] = i % modulus
			total += i
		}
		for _, m := range m5 {
			i := s1[m[1]] + s1[m[2]] + s1[m[3]] + s1[m[4]] + s1[m[5]]
			s2[m[0]] = i % modulus
			total += i
		}
		totals[k] = total % modulus
		s1, s2 = s2, s1
	}
	if debug {
		fmt.Printf("%d in %v\n", len(totals), time.Since(start))
	}
}

type state struct{ a, b, c, d int }

func (s state) String() string {
	return fmt.Sprintf("%d%d%d%d", s.a, s.b, s.c, s.d)
}

type sequencer map[state]int

func (m sequencer) add(s state) {
	if _, ok := m[s]; !ok {
		m[s] = len(m)
	}
}
