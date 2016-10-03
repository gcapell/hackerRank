package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	var m int
	fmt.Scanln(&m)
	fmt.Println((m / period(s)) % (1e9 + 7))
}

func period(s string) int {
	p := 1
	for j := 1; j < len(s); j++ {
		if s[j] != s[j%p] {
			p = j + 1
		}
	}
	return p
}
