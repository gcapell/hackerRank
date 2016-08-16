package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	rem := m % n
	if rem == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(n - rem)
	}
}
