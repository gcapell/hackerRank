package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(squaresBetween(a, b))
	}
}

func squaresBetween(a, b int) int {
	sa := math.Ceil(math.Sqrt(float64(a)))
	sb := math.Floor(math.Sqrt(float64(b)))
	return int(sb-sa) + 1
}
