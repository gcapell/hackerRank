package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	sticks := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&sticks[j])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sticks)))
	for biggest := 0; biggest <= len(sticks)-3; biggest++ {
		if sticks[biggest] < sticks[biggest+1]+sticks[biggest+2] {
			fmt.Println(sticks[biggest+2], sticks[biggest+1], sticks[biggest])
			return
		}
	}
	fmt.Println(-1)
}
