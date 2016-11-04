package main

import (
	"fmt"
)

func hits(tree, h1, h2, fruits int) int {
	reply := 0
	for j := 0; j < fruits; j++ {
		var f int
		fmt.Scanf("%d", &f)
		pos := tree + f
		if pos >= h1 && pos <= h2 {
			reply++
		}
	}
	return reply
}

func main() {
	var s, t, a, b, m, n int
	fmt.Scanf("%d %d", &s, &t)
	fmt.Scanf("%d %d", &a, &b)
	fmt.Scanf("%d %d", &m, &n)
	fmt.Println(hits(a, s, t, m))
	fmt.Println(hits(b, s, t, n))
}
