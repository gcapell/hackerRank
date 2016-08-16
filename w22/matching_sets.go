package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	x, y := readArray(n), readArray(n)
	fmt.Println(diff(x, y))
}

func readArray(n int) []int {
	a := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scanf("%d", &a[j])
	}
	sort.Ints(a)
	return a
}

func diff(a, b []int) int {
	var posDelta, delta int
	for j := range a {
		d := a[j] - b[j]
		if d > 0 {
			posDelta += d
		}
		delta += d
	}
	if delta != 0 {
		return -1
	}
	return posDelta
}
