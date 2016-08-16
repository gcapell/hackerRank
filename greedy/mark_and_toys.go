package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	toys := make([]int, n)
	for j := range toys {
		fmt.Scanf("%d", &toys[j])
	}
	sort.Ints(toys)
	var j int
	for ; j < len(toys) && k > toys[j]; j++ {
		k -= toys[j]
	}
	fmt.Println(j)
}
