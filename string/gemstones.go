package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var intersection uint32
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanf("%s", &s)
		n := bitmap(s)
		if j == 0 {
			intersection = n
		} else {
			intersection &= n
		}
	}
	fmt.Println(bits(intersection))
}

func bitmap(s string) uint32 {
	var b uint32
	for _, r := range s {
		b |= 1 << uint(r-'a')
	}
	return b
}

func bits(n uint32) int {
	count := 0
	for j := 0; j < 26; j++ {
		if n&1 != 0 {
			count++
		}
		n >>= 1
	}
	return count
}