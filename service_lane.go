package main

import (
	"fmt"
)

func main() {
	var n, t int
	fmt.Scanf("%d %d", &n, &t)
	widths := make([]int, n)
	for j := range widths {
		fmt.Scanf("%d", &widths[j])
	}
	for j := 0; j < t; j++ {
		var in, out int
		fmt.Scanf("%d %d", &in, &out)
		fmt.Println(minWidth(widths[in : out+1]))
	}
}

func minWidth(widths []int) int {
	var m int
	for j, w := range widths {
		if j == 0 || w < m {
			m = w
		}
	}
	return m
}
