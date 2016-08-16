package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a, b := readN(n), readN(n)
	var pairs int

	for k, v := range a {
		v2 := b[k]
		switch {
		case v2 == v:
			pairs += v
		case v2 < v:
			pairs += v2
		case v2 > v:
			pairs += v
		}
	}
	if pairs == n {
		pairs--
	} else {
		pairs++
	}
	fmt.Println(pairs)
}

func readN(n int) map[int]int {
	a := make(map[int]int)
	for j := 0; j < n; j++ {
		var d int
		fmt.Scanf("%d", &d)
		a[d]++
	}
	return a
}
