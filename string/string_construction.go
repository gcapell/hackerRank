package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(uniqueChars(s))
	}
}

func uniqueChars(s string) int {
	d := make(map[rune]bool)
	for _, r := range s {
		d[r] = true
	}
	return len(d)
}
