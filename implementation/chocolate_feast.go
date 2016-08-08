package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n, c, m int
		fmt.Scanf("%d %d %d", &n, &c, &m)
		fmt.Println(wrap(n/c, m))
	}
}

func wrap(ch, m int) int {
	w := ch
	for w >= m {
		ch += w / m
		w = w%m + w/m
	}
	return ch
}