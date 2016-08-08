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
	fmt.Println(a[len(a)/2])
}
