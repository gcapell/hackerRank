package main

import "fmt"

func main() {
	var heights [26]int
	for j := 0; j < len(heights); j++ {
		fmt.Scan(&heights[j])
	}
	var s string
	fmt.Scan(&s)
	var max int
	for _, c := range s {
		h := heights[c-'a']
		if h > max {
			max = h
		}
	}
	fmt.Println(max * len(s))
}
