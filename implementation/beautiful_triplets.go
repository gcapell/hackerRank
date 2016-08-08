package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)
	a := make([]int, n)
	for j := range a {
		fmt.Scanf("%d", &a[j])
	}
	var triples int
	j, k, l := 0, 1, 2
	for ; ; j++ {
		kTarget := a[j] + d
		for ; k < n && a[k] < kTarget; k++ {
		}
		if k == n {
			break
		}
		if a[k] != kTarget {
			continue
		}
		if k+1 > l {
			l = k + 1
		}
		lTarget := kTarget + d
		for ; l < n && a[l] < lTarget; l++ {
		}
		if l == n {
			break
		}
		if a[l] == lTarget {
			triples++
		}
	}
	fmt.Println(triples)
}
