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
	total := 0
	for j := 0; j < len(a); j++ {
		total += a[len(a)-j-1] << uint32(j)
	}
	fmt.Println(total)
}
