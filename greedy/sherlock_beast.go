package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(decent(n))
	}
}

func decent(n int) string {
	for j := 0; j < 3 && n >= 0; j++ {
		if n%3 == 0 {
			return strings.Repeat("5", n) + strings.Repeat("3", j*5)
		}
		n -= 5
	}
	return "-1"
}
