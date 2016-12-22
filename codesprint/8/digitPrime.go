package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	m, initialStateNums := prep()
	expCache[1] = m

	var q int
	fmt.Scan(&q)
	ns := make([]int, q)
	for j := 0; j < q; j++ {
		fmt.Scan(&ns[j])
	}
	sort.Ints(ns)

	prev := ns[0]
	m = exp(prev - 4)
	fmt.Println(m.sumCols(initialStateNums))

	for _, n := range ns {
		diff := n - prev
		if diff > 0 {
			diffM := exp(diff)
			m = mul(m, diffM)
		}
		fmt.Println(m.sumCols(initialStateNums))
		prev = n
	}
	fmt.Println("multiplies: ", multiplies, " cacheHits: ", cacheHits)
}

type state struct{ a, b, c, d int }

func prep() (matrix, []int) {
	prime := make([]bool, 46)
	for _, p := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43} {
		prime[p] = true
	}
	transitions := make(map[state][]state)
	var initialStates []state
	stateNum := make(map[state]int)
	nextStateNum := 0
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					if prime[a+b+c+d] && prime[a+b+c] && prime[b+c+d] {
						s := state{a, b, c, d}
						transitions[s] = nil
						stateNum[s] = nextStateNum
						nextStateNum++
						if a > 0 {
							initialStates = append(initialStates, s)
						}
					}
				}
			}
		}
	}

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
	matrixSize = len(transitions)
	m := newMatrix()
	for src, dsts := range transitions {
		for _, dst := range dsts {
			m.set(stateNum[dst], stateNum[src], 1)
		}
	}
	var initialStateNums []int
	for _, s := range initialStates {
		initialStateNums = append(initialStateNums, stateNum[s])
	}
	return m, initialStateNums
}

var expCache = make(map[int]matrix)
var cacheHits int

func exp(n int) matrix {
	if m, ok := expCache[n]; ok {
		cacheHits++
		return m
	}
	p := largestPowerOf2LTE(n)

	if p == n {
		m2 := exp(n / 2)
		m := mul(m2, m2)
		expCache[n] = m
		return m
	}
	m := mul(exp(p), exp(n-p))
	if n < 1024 {
		expCache[n] = m
	}
	return m
}

func largestPowerOf2LTE(n int) int {
	j := 1
	for j <= n {
		j *= 2
	}
	return j / 2
}

func power(powers []matrix, n int) matrix {
	if n >= len(powers) {
		log.Fatal("n", n)
	}
	if powers[n] != nil {
		return powers[n]
	}
	smaller := power(powers, n-1)
	pow := mul(smaller, smaller)
	powers[n] = pow
	return pow
}

var matrixSize int

const modulus = 1e9 + 7

type matrix []int

func (m matrix) sumCols(cols []int) int {
	reply := 0
	for r := 0; r < matrixSize; r++ {
		for _, c := range cols {
			reply += m[r*matrixSize+c]
		}
	}
	return reply % modulus
}

func newMatrix() matrix {
	return matrix(make([]int, matrixSize*matrixSize))
}

func (m matrix) set(r, c, val int) {
	m[r*matrixSize+c] = val
}

var multiplies int

func mul(a, b matrix) matrix {
	multiplies++
	reply := newMatrix()
	for r := 0; r < matrixSize; r++ {
		for c := 0; c < matrixSize; c++ {
			n := 0
			for k := 0; k < matrixSize; k++ {
				n += a[r*matrixSize+k] * b[k*matrixSize+c]
			}
			reply[r*matrixSize+c] = n % modulus
		}
	}
	return reply
}
