package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	perm := make([]int, n+1)
	for j := 1; j <= n; j++ {
		var pos int
		fmt.Scan(&pos)
		perm[pos] = j
	}
	for j := 1; j <= n; j++ {
		fmt.Println(perm[perm[j]])
	}
}
