package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var m, n int
		fmt.Scanf("%d\n%d", &m, &n)
		flavours := make([]int, n)
		for k := range flavours {
			fmt.Scanf("%d", &flavours[k])
		}
		a, b := indexSum(flavours, m)
		if a > b {
			a, b = b, a
		}
		fmt.Println(a, b)
	}
}

type indexedCost struct {
	cost, index int
}

type byCost []indexedCost

func (a byCost) Len() int           { return len(a) }
func (a byCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCost) Less(i, j int) bool { return a[i].cost < a[j].cost }

// indexSum returns the indices into cost whose values add to total.
func indexSum(cost []int, total int) (int, int) {
	a := make([]indexedCost, len(cost))
	for j, c := range cost {
		a[j] = indexedCost{c, j + 1}
	}
	sort.Sort(byCost(a))
	//fmt.Println(a, total)
	left := 0
	right := len(a) - 1
	for a[right].cost+a[left].cost > total {
		right--
	}
	// fmt.Println(right, a[right])
	for a[right].cost+a[left].cost != total {
		left++
		for a[right].cost+a[left].cost > total {
			right--
		}
	}
	return a[left].index, a[right].index
}
