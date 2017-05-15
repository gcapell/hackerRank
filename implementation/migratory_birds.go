package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var counts [6]int

	for j := 0; j < n; j++ {
		var t int
		fmt.Scan(&t)
		counts[t]++
	}
	var max, maxT int
	for j := 1; j <= 5; j++ {
		if counts[j] > max {
			max = counts[j]
			maxT = j
		}
	}
	fmt.Println(maxT)
}
