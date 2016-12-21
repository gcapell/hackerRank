package main

import "fmt"

func main() {
	var n [5]int
	fmt.Scanln(&n[0], &n[1], &n[2], &n[3], &n[4])
	min, max, sum := n[0], n[0], n[0]
	for _, x := range n[1:] {
		sum += x
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	fmt.Println(sum-max, sum-min)
}
