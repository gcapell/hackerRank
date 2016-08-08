package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	var total int
	var important []int
	for j := 0; j < n; j++ {
		var l, i int
		fmt.Scanf("%d %d", &l, &i)
		total += l
		if i == 1 {
			important = append(important, l)
		}
	}
	if len(important) > k {
		sort.Sort(sort.Reverse(sort.IntSlice(important)))
		total -= 2 * sum(important[k:])
	}
	fmt.Println(total)
}

func sum(s []int) int {
	var t int
	for _, n := range s {
		t += n
	}
	return t
}
