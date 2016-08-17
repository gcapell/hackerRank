package main

import (
	"fmt"
	"log"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	permuted := make([]int, n)
	pos := make([]int, n+1)
	for j := range permuted {
		var d int
		fmt.Scanf("%d", &d)
		permuted[j] = d
		pos[d] = j
	}
	for j := 0; j < len(permuted) && k > 0; j++ {
		got := permuted[j]
		if j != pos[got] {
			log.Fatal(j, got, pos)
		}
		want := n - j
		if got != want {
			permuted[pos[want]], permuted[pos[got]] = permuted[pos[got]], permuted[pos[want]]
			pos[got], pos[want] = pos[want], pos[got]
			k--
		}
	}
	s := fmt.Sprint(permuted)
	fmt.Println(s[1 : len(s)-1])
}
