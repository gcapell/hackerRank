package main

import "fmt"

var yn = map[bool]string{
	false: "NO",
	true:  "YES",
}

func main() {
	var t int
	fmt.Scanln(&t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanln(&n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
		}
		fmt.Println(yn[evenParity(a)])
	}
}

func evenParity(n []int) bool {
	even := true
	for j := 0; j < len(n); j++ {
		for k := j + 1; k < len(n); k++ {
			if n[k] < n[j] {
				even = !even
			}
		}
	}
	return even
}
