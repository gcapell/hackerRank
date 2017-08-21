package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanln(&t)
	for j := 0; j < t; j++ {
		var s string
		fmt.Scanln(&s)
		fmt.Println(anagramPairs(s))
	}
}

func anagramPairs(s string) int {
	counts := make(map[string]int)
	for j := 0; j < len(s); j++ {
		for k := j + 1; k <= len(s); k++ {
			counts[sorted(s[j:k])]++
		}
	}
	pairs := 0
	for _, v := range counts {
		if v > 1 {
			pairs += v * (v - 1) / 2
		}
	}
	return pairs
}

type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

func sorted(s string) string {
	b := byteSlice(s)
	sort.Sort(b)
	return string(b)
}
