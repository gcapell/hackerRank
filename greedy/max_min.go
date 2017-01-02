package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanln(&n)
	fmt.Scanln(&k)
	ns := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scanln(&ns[j])
	}
	sort.Ints(ns)
	var best int
	for start, end := 0, k-1; end < n; start, end = start+1, end+1 {
		unfair := ns[end] - ns[start]
		if start == 0 || unfair < best {
			best = unfair
		}
	}
	fmt.Println(best)
}
