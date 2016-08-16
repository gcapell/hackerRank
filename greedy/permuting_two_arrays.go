package main

import (
	"fmt"
	"sort"
)

func main() {
	var q int

	fmt.Scanf("%d", &q)
	for j := 0; j < q; j++ {
		if query() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func query() bool {
	var n, total int
	fmt.Scanf("%d %d", &n, &total)
	as, bs := readAndSort(n), readAndSort(n)
	for j, k := 0, len(as)-1; j <= k; j, k = j+1, k-1 {
		if as[j]+bs[k] < total {
			return false
		}
	}
	return true
}

func readAndSort(n int) []int {
	a := make([]int, n)
	for j := range a {
		fmt.Scanf("%d", &a[j])
	}
	sort.Ints(a)
	return a
}
