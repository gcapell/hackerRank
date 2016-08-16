package main

import "fmt"
import "strings"

func main() {
	n, m := 16, 100000
	fmt.Println(n, m)
	for j := 0; j < m; j++ {
		fmt.Println(2, 99, strings.Repeat("1", n))
	}
}
