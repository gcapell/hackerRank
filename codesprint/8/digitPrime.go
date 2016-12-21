package main

import (
	"fmt"
	"sort"
)

type state struct {
	a, b, c, d int
}

func (s state) String() string {
	return fmt.Sprintf("%d%d%d%d", s.a, s.b, s.c, s.d)
}

func main() {
	prime := make([]bool, 46)
	for _, p := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43} {
		prime[p] = true
	}
	transitions := make(map[state][]state)
	initialState := make(map[state]int)
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
							initialState[s] = 1
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
	fmt.Println("size", matrixSize)
	m := newMatrix()
	for src, dsts := range transitions {
		for _, dst := range dsts {
			fmt.Println(src, stateNum[src], dst, stateNum[dst])
			m.set(stateNum[src], stateNum[dst], 1)
		}
	}
	fmt.Println(m)

	thisState := initialState
	for j := 0; j < 20; j++ {
		internal := 0
		nextState := make(map[state]int)
		for s, count := range thisState {
			if len(transitions[s]) == 0 {
				continue
			}
			internal++
			for _, v := range transitions[s] {
				nextState[v] += count
			}
		}
		var throughPaths, terminalPaths int
		for s, v := range nextState {
			if len(transitions[s]) > 0 {
				throughPaths += v
			} else {
				terminalPaths += v
			}
		}
		showState(nextState)
		fmt.Println()
		fmt.Printf("through:%d, terminal:%d\n", throughPaths, terminalPaths)
		fmt.Printf("\n################# internal nodes:%d, total nodes: %d\n", internal, len(thisState))
		thisState = nextState
	}
	fmt.Println("internals")
	for k := range thisState {
		if len(transitions[k])>0 {
		fmt.Println(k, transitions[k])
		}
	}
}

var matrixSize int
type matrix []int

func newMatrix() matrix {
	return matrix(make([]int,matrixSize*matrixSize))
}

func (m matrix)set(r,c,val int) {
	fmt.Println(len(m), "set", r, c, val, r*matrixSize+c)
	m[r*matrixSize+c] = val
}

func showState(s map[state]int) {
	var keys []state
	for k := range s {
		keys = append(keys, k)
	}
	sort.Sort(ByState(keys))
	for _, k := range keys {
		fmt.Printf("%s:%d, ", k, s[k])
	}
}

type ByState []state

func (a ByState) Len() int      { return len(a) }
func (a ByState) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByState) Less(i, j int) bool {
	return a[i].String() < a[j].String()
}
