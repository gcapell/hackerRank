package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	counts := make(map[int]int)
	for j := 0; j < n; j++ {
		var d int
		fmt.Scanf("%d", &d)
		counts[d]++
	}
	highest := 0
	for _, v := range counts {
		if v > highest {
			highest = v
		}
	}
	fmt.Println(n - highest)

}
