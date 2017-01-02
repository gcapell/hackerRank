package main

import (
	"fmt"
	"sort"
)

func main() {
	var cases int
	fmt.Scanln(&cases)
	yn := map[bool]string{
		false: "NO",
		true:  "YES",
	}
	for j := 0; j < cases; j++ {
		fmt.Println(yn[solve()])
	}
}

func solve() bool {
	var size int
	fmt.Scanln(&size)
	g := make([]string, size)
	for j := 0; j < size; j++ {
		fmt.Scanln(&g[j])
		g[j] = sortChars(g[j])
	}
	return sortedCols(g)
}

type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

func sortChars(s string) string {
	b := []byte(s)
	sort.Sort(byteSlice(b))
	return string(b)
}

func sortedCols(g []string) bool {
	for j := 0; j < len(g); j++ {
		for k := 1; k < len(g); k++ {
			if g[k][j] < g[k-1][j] {
				return false
			}
		}
	}
	return true
}
