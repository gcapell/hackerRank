package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(dups(s))
	}
}

func dups(s string) int {
	var prev rune
	var count int
	for _, r := range s {
		if r == prev {
			count++
		}
		prev = r
	}
	return count
}
