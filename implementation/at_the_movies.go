package main

import "fmt"

func main() {
	var i, j, k int
	fmt.Scanln(&i, &j, &k)
	count := 0
	for ; i <= j; i++ {
		if beautiful(i, k) {
			count++
		}
	}
	fmt.Println(count)
}

func beautiful(i, k int) bool {
	return absDiff(i, reverse(i))%k == 0
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func reverse(i int) int {
	reply := 0
	for i > 0 {
		n := i % 10
		i /= 10
		reply = reply*10 + n
	}
	return reply
}
