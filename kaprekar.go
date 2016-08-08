package main

import (
	"fmt"
	"log"
)

func main() {
	var p, q int
	fmt.Scanf("%d", &p)
	fmt.Scanf("%d", &q)
	first := true
	for j := p; j <= q; j++ {
		if kaprekar(j) {
			if !first {
				fmt.Print(" ")
			}
			first = false
			fmt.Print(j)
		}
	}
	if first {
		fmt.Print("INVALID RANGE")
	}
	fmt.Println()
}

func kaprekar(n int) bool {
	n2 := n * n
	m := minMod(n2)
	return (n2/m)+(n2%m) == n
}

var modulus = []struct{ n, m int }{
	{10, 10},
	{100, 10},
	{1000, 100},
	{10000, 100},
	{100000, 1000},
	{1000000, 1000},
	{10000000, 10000},
	{100000000, 10000},
	{1000000000, 100000},
	{10000000000, 100000},
	{100000000000, 1000000},
}

// We could be tricky and avoid looping through
// early part of modulus array each time, but meh.
func minMod(n int) (m int) {
	for _, mod := range modulus {
		if n < mod.n {
			return mod.m
		}
	}
	log.Fatal("minMod", n)
	return -1
}
