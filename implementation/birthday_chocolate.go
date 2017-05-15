package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&nums[j])
	}
	var d, m int
	fmt.Scan(&d, &m)

	var count, total int
	for j := 0; j < m; j++ {
		total += nums[j]
	}
	if total == d {
		count++
	}
	start, end := 0, m-1
	for {
		end++
		if end == n {
			break
		}
		total += nums[end]
		total -= nums[start]
		start++
		if total == d {
			count++
		}
	}

	fmt.Println(count)
}
