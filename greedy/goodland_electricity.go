package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	tower := make([]bool, n)
	for j := 0; j < n; j++ {
		var t int
		fmt.Scan(&t)
		tower[j] = t == 1
	}
	fmt.Println(minOn(tower, k))
}

func minOn(tower []bool, k int) int {
	unpowered := 0
	on := 0
	for unpowered < len(tower) {
		m, ok := furthest(tower, unpowered, k)
		if !ok {
			return -1
		}
		on++
		unpowered = m + k
	}
	return on
}

func furthest(tower []bool, start, k int) (int, bool) {
	far := start + k - 1
	if far > len(tower)-1 {
		far = len(tower) - 1
	}
	near := start - k + 1
	if near < 0 {
		near = 0
	}
	for j := far; j >= near; j-- {
		if tower[j] {
			return j, true
		}
	}
	fmt.Println(start, k, far, near)
	return -1, false
}
