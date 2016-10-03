package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanf("%d", &n)
		q := make([]int, n)
		for j := range q {
			fmt.Scanf("%d", &q[j])
		}
		fmt.Println(bribes(q))
	}
}

func bribes(q []int) string {
	var reply int
	for pos, e := range q {
		delta := e - (pos + 1)
		switch {
		case delta <= 0:
			continue
		case delta == 1, delta == 2:
			reply += delta
		default:
			return "Too chaotic"
		}
	}
	return strconv.Itoa(reply)
}
