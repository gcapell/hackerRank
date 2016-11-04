package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	shared := 0
	for j := 0; j < n; j++ {
		var c int
		fmt.Scanf("%d", &c)
		if j != k {
			shared += c
		}
	}
	var bill int
	fmt.Scanf("%d", &bill)
	if bill == shared/2 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(bill - shared/2)
	}
}

func asIn(s string) int {
	return strings.Count(s, "a")
}
