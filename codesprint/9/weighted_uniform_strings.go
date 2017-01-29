package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scanln(&s)
	weights := uniformWeights(s)
	fmt.Scanln(&n)
	for j := 0; j < n; j++ {
		var weight int
		fmt.Scanln(&weight)
		if weights[weight] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func uniformWeights(s string) map[int]bool {
	weights := make(map[int]bool)
	for start := 0; start < len(s); {
		var end int
		for end = start + 1; end < len(s) && s[start] == s[end]; end++ {
		}
		letterWeight := int(s[start] - 'a' + 1)
		w := letterWeight
		for j := 0; j < (end - start); j++ {
			weights[w] = true
			w += letterWeight
		}
		start = end
	}
	return weights
}
