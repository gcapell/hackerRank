package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	fmt.Println(minDelete(a, b))
}

func minDelete(a, b string) int {
	ca, cb := count(a), count(b)
	var del int
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		del += diff(ca[r], cb[r])
	}
	return del
}

func count(s string) map[rune]int {
	reply := make(map[rune]int)
	for _, r := range s {
		reply[r]++
	}
	return reply
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
