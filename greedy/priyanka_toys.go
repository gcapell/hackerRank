package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	weights := make([]int, n)
	for j := range weights {
		fmt.Scanf("%d", &weights[j])
	}
	sort.Ints(weights)
	purchases := 0
	for len(weights) > 0 {
		purchases++
		var w int
		w, weights = weights[0], weights[1:]
		for len(weights) > 0 && weights[0] <= w+4 {
			weights = weights[1:]
		}
	}
	fmt.Println(purchases)
}
