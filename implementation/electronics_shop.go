package main

import (
	"fmt"
	"sort"
)

func main() {
	var money, keyboards, drives int
	fmt.Scanln(&money, &keyboards, &drives)
	keyPrices := scanAndSort(keyboards)
	drivePrices := scanAndSort(drives)
	fmt.Println(sumLE(money, keyPrices, drivePrices))
}

// sumLE returns the largest sum of one number from a and one from b
// which is <= total.  (or -1)
func sumLE(total int, as, bs []int) int {
	a := 0
	b := sort.SearchInts(bs, total-as[a])
	if b == len(bs) {
		b--
	}
	for b > 0 && as[a]+bs[b] > total {
		b--
	}
	max := as[a] + bs[b]
	if max > total {
		return -1
	}
	if max == total {
		return max
	}
outer:
	for {
		a++
		if a == len(as) {
			break
		}
		for as[a]+bs[b] > total {
			b--
			if b < 0 {
				break outer
			}
		}
		//fmt.Println(a, b)
		if as[a]+bs[b] > max {
			max = as[a] + bs[b]
			if max == total {
				break
			}
		}
	}
	return max
}

func scanAndSort(n int) []int {
	reply := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&reply[j])
	}
	sort.Ints(reply)
	return reply
}
