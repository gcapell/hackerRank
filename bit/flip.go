package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n uint32
		fmt.Scanf("%d", &n)
		fmt.Println(^n)
	}
}
