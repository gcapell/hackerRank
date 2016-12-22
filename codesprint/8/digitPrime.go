package main

import "fmt"

func main() {
	powers, initialStateNums := prep()
	var q int
	fmt.Scan(&q)
	for j := 0; j < q; j++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(exp(n-4, powers).sumCols(initialStateNums))
	}
}

type state struct{ a, b, c, d int }

func prep() ([]matrix, []int) {
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
	powers := []matrix{nil, m}
	for j := 2; j < 40e3; j *= 2 {
		m = mul(m, m)
		powers = append(powers, m)
	}
	var initialStateNums []int
	for _, s := range initialStates {
		initialStateNums = append(initialStateNums, stateNum[s])
	}
	return powers, initialStateNums
}

func exp(n int, powers []matrix) matrix {
	var m matrix
	p := 1
	for n > 0 {
		if n%2 == 1 {
			if m == nil {
				m = powers[p]
			} else {
				m = mul(m, powers[p])
			}
		}
		n /= 2
		p++
	}
	return m
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

func mul(a, b matrix) matrix {
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
