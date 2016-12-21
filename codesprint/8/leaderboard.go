package main

import (
	"fmt"
	"sort"
)

func main() {
	board := removeDups(scanInts())
	reverse(board)
	scores := scanInts()
	for _, s := range scores {
		pos := sort.SearchInts(board, s)
		if pos < len(board) && board[pos] == s {
			pos++
		}
		fmt.Println(len(board) - pos + 1)
	}
}

func reverse(a []int) {
	for j, k := 0, len(a)-1; j < k; j, k = j+1, k-1 {
		a[j], a[k] = a[k], a[j]
	}
}

func removeDups(a []int) []int {
	var reply []int
	prev := -1
	for _, n := range a {
		if n == prev {
			continue
		}
		reply = append(reply, n)
		prev = n
	}
	return reply
}

func scanInts() []int {
	var n int
	fmt.Scanln(&n)
	reply := make([]int, n)
	for p := range reply {
		fmt.Scan(&reply[p])
	}
	return reply
}
