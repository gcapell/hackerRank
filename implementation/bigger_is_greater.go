package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(inc(s))
	}
}

func inc(s string) string {
	if len(s) < 2 {
		return "no answer"
	}
	b := []byte(s)
	for j := len(b) - 2; j >= 0; j-- {
		if b[j] < b[j+1] {
			pivot(b[j:])
			return string(b)
		}
	}
	return "no answer"
}

func pivot(b []byte) {
	for j := len(b) - 1; ; j-- {
		if b[j] > b[0] {
			b[j], b[0] = b[0], b[j]
			break
		}
	}
	reverse(b[1:])
}

func reverse(b []byte) {
	for j, k := 0, len(b)-1; j < k; j, k = j+1, k-1 {
		b[j], b[k] = b[k], b[j]
	}
}
