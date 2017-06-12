package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	a := make(map[int]int)
	for j := 0; j < n; j++ {
		var k int
		fmt.Scan(&k)
		a[k]++
	}
	var m int
	fmt.Scanln(&m)
	missing := make(map[int]bool)
	for j := 0; j < m; j++ {
		var k int
		fmt.Scan(&k)
		if a[k] == 0 {
			missing[k] = true
		} else {
			a[k]--
		}
	}
	var missingL []int
	for k := range missing {
		missingL = append(missingL, k)
	}
	sort.Ints(missingL)

	for p, k := range missingL {
		if p != 0 {
			fmt.Print(" ")
		}
		fmt.Print(k)
	}
	fmt.Println()
}
