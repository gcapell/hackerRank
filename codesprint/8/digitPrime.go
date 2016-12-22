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

	for ns[0] < 5 {
		ns = ns[1:]
	}
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
	if n < 1 {
		log.Fatal("n", n)
	}
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

var matrixSize int

const modulus = 1e9 + 7

type matrix struct {
	rows, cols sparseVector
}

type sparseVector []map[int]int

func newMatrix() matrix {
	return matrix{
		rows: newSparseVector(),
		cols: newSparseVector(),
	}
}

func newSparseVector() sparseVector {
	return make([]map[int]int, matrixSize)
}

func (m matrix) set(r, c, val int) {
	m.rows.set(r, c, val)
	m.cols.set(c, r, val)
}

func (v sparseVector) set(r, c, val int) {
	if v[r] == nil {
		v[r] = make(map[int]int)
	}
	v[r][c] = val
}

var multiplies int

func mul(a, b matrix) matrix {
	multiplies++
	reply := newMatrix()

	for r := 0; r < matrixSize; r++ {
		for c := 0; c < matrixSize; c++ {
			v := sparseDotProduct(a.rows[r], b.cols[c])
			if v != 0 {
				reply.set(r, c, v%modulus)
			}
		}
	}
	return reply
}

func sparseDotProduct(a, b map[int]int) int {
	if a == nil || b == nil {
		return 0
	}
	var reply int
	for k, v := range a {
		if v2, ok := b[k]; ok {
			reply += v * v2
		}
	}
	return reply
}

func (m matrix) sumCols(cols []int) int {
	reply := 0
	for _, c := range cols {
		for _, v := range m.cols[c] {
			reply += v
		}
	}
	return reply % modulus
}
