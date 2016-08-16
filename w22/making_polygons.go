package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	side := make([]int, n)
	for j := range side {
		fmt.Scanf("%d", &side[j])
	}
	fmt.Println(cuts(side))
}

func cuts(side []int) int {
	switch len(side) {
	case 1:
		return 2
	case 2:
		return 1
	}
	sort.Ints(side)
	if side[len(side)-1] >= sum(side[:len(side)-1]) {
		return 1
	}
	return 0
}

func sum(side []int) (total int) {
	for _, s := range side {
		total += s
	}
	return total
}
