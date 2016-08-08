package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(height(n))
	}
}
func height(n int) uint64 {
	h := uint64(1)
	for {
		if n == 0 {
			break
		}
		h *= 2
		n--
		if n == 0 {
			break
		}
		h++
		n--
	}
	return h
}
