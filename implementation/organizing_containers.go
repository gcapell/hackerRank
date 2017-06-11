package main

import (
	"fmt"
	"sort"
)

func main() {
	var q int
	fmt.Scanln(&q)
	for j := 0; j < q; j++ {
		query()
	}
}

func query() {
	var n int
	fmt.Scanln(&n)
	m := scanSquare(n)
	r, c := sumRows(m), sumCols(m)
	sort.Ints(r)
	sort.Ints(c)
	if same(r, c) {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}

func sumRows(n [][]int) []int {
	reply := make([]int, len(n))
	for j := 0; j < len(n); j++ {
		for k := 0; k < len(n); k++ {
			reply[j] += n[j][k]
		}
	}
	return reply
}

func sumCols(n [][]int) []int {
	reply := make([]int, len(n))
	for j := 0; j < len(n); j++ {
		for k := 0; k < len(n); k++ {
			reply[j] += n[k][j]
		}
	}
	return reply
}

func same(a, b []int) bool {
	for p := range a {
		if a[p] != b[p] {
			return false
		}
	}
	return true
}

func scanSquare(n int) [][]int {
	var reply [][]int
	for j := 0; j < n; j++ {
		row := make([]int, n)
		for k := 0; k < n; k++ {
			fmt.Scan(&row[k])
		}
		reply = append(reply, row)
	}
	return reply
}
