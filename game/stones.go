package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanf("%d", &n)
		if n%7 < 2 {
			fmt.Println("Second")
		} else {
			fmt.Println("First")
		}
	}
}
