package main

import (
	"fmt"
	"log"
)

func scanln(a ...interface{}) {
	if _, err := fmt.Scanln(a...); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var t int
	scanln(&t)
	for j := 0; j < t; j++ {
		var n, a, b int
		scanln(&n)
		scanln(&a)
		scanln(&b)
		stones(n, a, b)
	}
}

func stones(n, a, b int) {
	if a > b {
		a, b = b, a
	}
	total := (n - 1) * a
	if a == b {
		fmt.Println(total)
		return
	}
	for j := 0; j < n; j++ {
		if j != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", total)
		total += b - a
	}
	fmt.Println()
}
