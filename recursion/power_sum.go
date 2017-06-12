package main

import "fmt"

func main() {
	var x, n int
	fmt.Scanln(&x)
	fmt.Scanln(&n)
	fmt.Println(sumPow(x, n, 1))
}

// How many ways can x be expressed as sums of nth powers
// of numbers >= j ?
func sumPow(x, n, j int) int {
	p := pow(j, n)
	switch {
	case p > x:
		return 0
	case p == x:
		return 1
	default:
		return sumPow(x-p, n, j+1) + sumPow(x, n, j+1)
	}
}

var cache = make(map[int]int)

func pow(j, n int) (reply int) {
	if reply, ok := cache[j]; ok {
		return reply
	}
	defer func() {
		cache[j] = reply
	}()
	reply = 1
	for k := 0; k < n; k++ {
		reply *= j
	}
	return reply
}
