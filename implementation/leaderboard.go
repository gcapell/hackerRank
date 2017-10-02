package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := dense(readn())
	alice := readn()
	for _, a := range alice {
		r := sort.Search(len(scores), func(p int) bool {
			return scores[p] <= a
		})
		fmt.Println(r + 1)
	}
}

func dense(scores []int) []int {
	dst := 0
	for src := 1; src < len(scores); src++ {
		if scores[src] != scores[dst] {
			dst++
			if dst != src {
				scores[dst] = scores[src]
			}
		}
	}
	return scores[:dst+1]
}

func readn() []int {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&a[j])
	}
	return a
}
