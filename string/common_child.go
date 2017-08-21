package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(common(a, b))
}

// common returns length of longest common
// child of a,b.mate
func common(a, b string) int {
	prev := make([]int, len(b)+1)
	var next []int
	for j := 0; j < len(a); j++ {
		next = make([]int, len(b)+1)
		for k := 0; k < len(b); k++ {
			if a[j] == b[k] {
				next[k+1] = max(next[k], prev[k]+1)
			} else {
				next[k+1] = max(next[k], prev[k+1])
			}
		}
		//fmt.Println(string(a[j]), next)
		prev = next
	}
	return prev[len(b)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
