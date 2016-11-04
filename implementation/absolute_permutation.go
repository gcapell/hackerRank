package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n, k int
		fmt.Scanf("%d %d", &n, &k)
		switch {
		case k == 0:
			enumerate(n)
		case n%(2*k) == 0:
			permute(n, k)
		default:
			fmt.Println(-1)
		}
	}
}

func enumerate(n int) {
	for j := 1; j <= n; j++ {
		fmt.Print(j, " ")
	}
	fmt.Println()
}

func permute(n, k int) {
	for base := 1; base < n; base += 2 * k {
		for j := 0; j < k; j++ {
			fmt.Print(base+k+j, " ")
		}
		for j := 0; j < k; j++ {
			fmt.Print(base+j, " ")
		}
	}
	fmt.Println()
}
