package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var s string
	fmt.Scanln(&s)

	freq := make(map[byte]int)
	for _, r := range s {
		freq[byte(r)]++
	}
	upper, lower, best, target := 0, 0, n, n/4
bigloop:
	for {
		for anyGreater(freq, target) {
			freq[s[upper]]--
			upper++
			if upper == n {
				break bigloop
			}
		}
		for !anyGreater(freq, target) {
			freq[s[lower]]++
			lower++
		}
		delta := upper - lower + 1
		if delta < best {
			best = delta
		}
	}
	fmt.Println(best)
}

func anyGreater(f map[byte]int, target int) bool {
	for _, v := range f {
		if v > target {
			return true
		}
	}
	return false
}
