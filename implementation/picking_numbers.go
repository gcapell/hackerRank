package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	counts := make(map[int]int)
	for j := 0; j < n; j++ {
		var e int
		fmt.Scan(&e)
		counts[e]++
	}
	var keys []int
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	max := counts[keys[0]]
	for pos := 1; pos < len(keys); pos++ {
		count := counts[keys[pos]]
		if keys[pos] == 1+keys[pos-1] {
			count += counts[keys[pos-1]]
		}
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
}
