package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	costs := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&costs[j])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(costs)))
	total := 0
	for j, c := range costs {
		total += c * (j/k + 1)
	}
	fmt.Println(total)
}
