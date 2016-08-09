package main

import "fmt"

var yn = map[bool]string{true: "YES", false: "NO"}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanf("%d", &n)
		a := make([]int, n)
		for k := range a {
			fmt.Scanf("%d", &a[k])
		}

		fmt.Println(yn[equalSum(a)])
	}
}

func equalSum(a []int) bool {
	cumSum := make([]int, len(a))
	total := 0
	for j, n := range a {
		total += n
		cumSum[j] = total
	}
	for j, n := range cumSum {
		var left int
		if j != 0 {
			left = cumSum[j-1]
		}
		right := total - n
		if left == right {
			return true
		}
	}
	return false
}
