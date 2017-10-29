package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	buildings := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&buildings[n-j-1])
	}
	required := 0
	for _, h := range buildings {
		//fmt.Println("r:", required, "h:", h)
		if h >= required {
			required += (h - required + 1) / 2
		} else {
			required -= (required - h) / 2
		}
	}
	fmt.Println(required)
}
