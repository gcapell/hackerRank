package main

import (
	"fmt"
	"strings"
)

const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	words := 1
	for {
		n := strings.IndexAny(s, upper)
		if n == -1 {
			break
		}
		s = s[n+1:]
		words++
	}
	fmt.Println(words)
}
