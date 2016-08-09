package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	var plain string
	fmt.Scanf("%d\n%s\n%d", &n, &plain, &k)
	k %= 26
	fmt.Println(strings.Map(func(r rune) rune {
		r = modRange(r, rune('a'), rune('z'), k)
		r = modRange(r, rune('A'), rune('Z'), k)
		return r
	}, plain))
}

func modRange(r, lower, upper rune, k int) rune {
	if r < lower || r > upper {
		return r
	}
	r += rune(k)
	if r > upper {
		r -= (upper + 1 - lower)
	}
	return r
}
