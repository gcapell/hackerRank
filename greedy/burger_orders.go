package main

import (
	"fmt"
	"sort"
)

type order struct {
	id, t, d int
}

type byDelivery []order

func (a byDelivery) Len() int      { return len(a) }
func (a byDelivery) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDelivery) Less(i, j int) bool {
	iVal := a[i].t + a[i].d
	jVal := a[j].t + a[j].d
	switch {
	case iVal < jVal:
		return true
	case iVal > jVal:
		return false
	}
	return a[i].id < a[j].id
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var orders []order
	for j := 0; j < n; j++ {
		var t, d int
		fmt.Scanf("%d %d", &t, &d)
		orders = append(orders, order{j + 1, t, d})
	}
	sort.Sort(byDelivery(orders))

	for j, o := range orders {
		if j != 0 {
			fmt.Print(" ")
		}
		fmt.Print(o.id)
	}
	fmt.Println()
}
