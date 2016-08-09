package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(anagram(s))
	}
}

func anagram(s string) (reply int) {
	n := len(s)
	if n%2 == 1 {
		return -1
	}
	a, b := count(s[:n/2]), count(s[n/2:])
	for letter, n := range a {
		if n > b[letter] {
			reply += n - b[letter]
		}
	}
	return reply
}

func count(s string) map[rune]int {
	reply := make(map[rune]int)
	for _, r := range s {
		reply[r]++
	}
	return reply
}