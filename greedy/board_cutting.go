package main

import (
	"fmt"
	"sort"
)

func main() {
	var queries int
	fmt.Scanln(&queries)
	for j := 0; j < queries; j++ {
		query()
	}
}

func query() {
	var m, n int
	fmt.Scanln(&m, &n)
	var cuts []cut
	addCuts(&cuts, m-1, false)
	addCuts(&cuts, n-1, true)
	sort.Sort(byCost(cuts))
	counts := map[bool]int{false: 1, true: 1}
	total := 0
	for _, c := range cuts {
		total = (total + counts[c.vertical]*c.cost) % (1e9 + 7)
		counts[!c.vertical]++
	}
	fmt.Println(total)
}

type cut struct {
	vertical bool
	cost     int
}

type byCost []cut

func (a byCost) Len() int           { return len(a) }
func (a byCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCost) Less(i, j int) bool { return a[i].cost > a[j].cost }

func addCuts(cuts *[]cut, n int, vertical bool) {
	for j := 0; j < n; j++ {
		var cost int
		fmt.Scan(&cost)
		*cuts = append(*cuts, cut{vertical, cost})
	}
}
