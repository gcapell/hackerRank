package main

import (
	"fmt"
	"sort"
)

type delta struct {
	start, delta int
}

type ByStart []delta

func (a ByStart) Len() int      { return len(a) }
func (a ByStart) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool {
	if a[i].start == a[j].start {
		return a[i].delta < a[j].delta
	}
	return a[i].start < a[j].start
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	deltas := make([]delta, 0, m*2)
	for j := 0; j < m; j++ {
		var a, b, d int
		fmt.Scanln(&a, &b, &d)
		deltas = append(deltas, delta{a, d})
		deltas = append(deltas, delta{b + 1, -d})
	}
	sort.Sort(ByStart(deltas))

	var max, total int
	for _, d := range deltas {
		total += d.delta
		if total > max {
			max = total
		}
	}
	fmt.Println(max)
}
