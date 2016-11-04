package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sock := make([]int, n)
	for j := range sock {
		fmt.Scanf("%d", &sock[j])
	}
	sort.Ints(sock)
	pairs := 0
	matched := true
	last := 0
	for _, s := range sock {
		if !matched && last == s {
			matched = true
			pairs++
		} else {
			matched = false
		}
		last = s
	}
	fmt.Println(pairs)
}
