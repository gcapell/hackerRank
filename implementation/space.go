package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scanf("%d %d", &n, &m)
	station := make([]int, m)
	for j := range station {
		fmt.Scanf("%d", &station[j])
	}
	sort.Ints(station)
	var d max
	var last int
	for j, s := range station {
		//fmt.Printf("s:%d ", s)
		switch {
		case m == 1:
			d.update(n - 1 - s)
			d.update(s)
		case j == 0:
			d.update(s)
		case j == m-1:
			// This case first for correctness when m==1
			d.update(n - 1 - s)
			d.update((s - last) / 2)
		default:
			d.update((s - last) / 2)
		}
		last = s
	}
	fmt.Println(d)
}

type max int

func (d *max) update(n int) {
	//fmt.Println("n:", n)
	if n > int(*d) {
		*d = max(n)
	}
}
