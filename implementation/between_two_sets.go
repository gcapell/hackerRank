package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for j := 0; j < n; j++ {
		fmt.Scanf("%d", &a[j])
	}
	for j := 0; j < m; j++ {
		fmt.Scanf("%d", &b[j])
	}

	start := max(a)
	end := min(b)
	inc := min(a)
	count := 0
	for j := start; j <= end; j += inc {
		if allFactors(j, a) && allMultiples(j, b) {
			count++
		}
	}
	fmt.Println(count)
}

func max(a []int) int {
	m := a[0]
	for j := 1; j < len(a); j++ {
		if a[j] > m {
			m = a[j]
		}
	}
	return m
}
func min(a []int) int {
	m := a[0]
	for j := 1; j < len(a); j++ {
		if a[j] < m {
			m = a[j]
		}
	}
	return m
}

func allFactors(n int, as []int) bool {
	for _, a := range as {
		if n%a != 0 {
			return false
		}
	}
	return true
}

func allMultiples(n int, as []int) bool {
	for _, a := range as {
		if a%n != 0 {
			return false
		}
	}
	return true
}
