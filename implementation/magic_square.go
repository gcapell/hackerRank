package main

import "fmt"

var squares = [][]int{
	{8, 1, 6, 3, 5, 7, 4, 9, 2},
	{6, 1, 8, 7, 5, 3, 2, 9, 4},
	{4, 9, 2, 3, 5, 7, 8, 1, 6},
	{2, 9, 4, 7, 5, 3, 6, 1, 8},
	{8, 3, 4, 1, 5, 9, 6, 7, 2},
	{4, 3, 8, 9, 5, 1, 2, 7, 6},
	{6, 7, 2, 1, 5, 9, 8, 3, 4},
	{2, 7, 6, 9, 5, 1, 4, 3, 8},
}

func main() {
	var s [9]int
	for j := 0; j < 9; j++ {
		fmt.Scan(&s[j])
	}
	fmt.Println(min(diffs(s[:], squares)))
}

func diffs(s []int, magic [][]int) []int {
	var reply []int
	for _, m := range magic {
		reply = append(reply, diff(s, m))
	}
	return reply
}

func min(n []int) int {
	m := n[0]
	for _, o := range n[1:] {
		if o < m {
			m = o
		}
	}
	return m
}

func diff(a, b []int) int {
	total := 0
	for pos := range a {
		delta := a[pos] - b[pos]
		if delta < 0 {
			delta = -delta
		}
		total += delta

	}
	return total
}
