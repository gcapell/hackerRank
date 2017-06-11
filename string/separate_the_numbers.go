package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanln(&s)
		if first, ok := pretty(s); ok {
			fmt.Println("YES", first)
		} else {
			fmt.Println("NO")
		}
	}
}

func pretty(s string) (int, bool) {
	if len(s) < 2 || s[0] == '0' {
		return 0, false
	}
	for size := 1; size <= len(s)/2; size++ {
		n, _ := strconv.Atoi(s[:size])
		s2 := gen(n, len(s))
		if s == s2 {
			return n, true
		}
	}
	return 0, false
}

func gen(n, size int) string {
	s := strconv.Itoa(n)
	for len(s) < size {
		n++
		s += strconv.Itoa(n)
	}
	return s
}
