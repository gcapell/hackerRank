package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(cost(s))
	}
}

func cost(s string) (reply int) {
	for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
		reply += delta(s[j], s[k])
	}
	return reply
}

func delta(a, b byte) int {
	if a < b {
		a, b = b, a
	}
	return int(a - b)
}