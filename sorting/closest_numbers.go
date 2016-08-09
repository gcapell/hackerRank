package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for j := range a {
		fmt.Scanf("%d", &a[j])
	}
	sort.Ints(a)
	minDiff := a[1] - a[0]
	minPairs := []int{a[0], a[1]}

	for j := 2; j < len(a); j++ {
		d := a[j] - a[j-1]

		switch {
		case d == minDiff:
			minPairs = append(minPairs, a[j-1], a[j])
		case d < minDiff:
			minDiff = d
			minPairs = []int{a[j-1], a[j]}
		}
	}

	for j, n := range minPairs {
		if j != 0 {
			fmt.Print(" ")
		}
		fmt.Print(n)
	}
	fmt.Println()
}
