package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	c := make([]int, m)
	for j := 0; j < m; j++ {
		fmt.Scan(&c[j])
	}
	fmt.Println(change(n, c))
}

func change(n int, coins []int) int {
	prev := make([]int, n+1)
	next := make([]int, n+1)
	for _, c := range coins {
		for j := 0; j < n; j++ {
			if prev[j] == 0 {
				continue
			}
			for k := j + c; k <= n; k += c {
				next[k] += prev[j]
			}
		}
		prev, next = next, prev
	}
	return prev[n]
}
