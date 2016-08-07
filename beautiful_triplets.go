package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)
	a := make([]int, n)
	for j := range a {
		fmt.Scanf("%d", &a[j])
	}
	j,k,l := 0,1,2
	for ; j<n; j++ {
		kTarget := a[j] +d
		for ; k<n && a[k] <= kTarget; k++ {
			if a[k] == kTarget {
				lTarget := kTarget + d
				for l := k+
				
			}
		}
		
	}
}

func minWidth(widths []int) int {
	var m int
	for j, w := range widths {
		if j == 0 || w < m {
			m = w
		}
	}
	return m
}
