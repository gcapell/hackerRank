package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var count, last int
	open := false
	for j := 0; j < n; j++ {
		var b int
		fmt.Scanf("%d", &b)
		if b%2 == 0 {
			continue
		}
		if open {
			count += j - last
		} else {
			last = j
		}
		open = !open
	}
	if open {
		fmt.Println("NO")
	} else {
		fmt.Println(count * 2)
	}
}
