package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&a[j])
	}
	sort.Ints(a)
	min := a[1] - a[0]
	for j := 2; j < len(a); j++ {
		diff := a[j] - a[j-1]
		if diff < min {
			min = diff
			if min == 0 {
				break
			}
		}
	}
	fmt.Println(min)
}
