package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	max := 0
	for j := 0; j < n; j++ {
		var x int
		fmt.Scan(&x)
		if x > max {
			max = x
		}
	}
	if max > k {
		fmt.Println(max - k)
	} else {
		fmt.Println(0)
	}
}
