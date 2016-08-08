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
		var b, w, x, y, z int
		scanln(&b, &w)
		scanln(&x, &y, &z)
		fmt.Println(cost(b, w, x, y, z))
	}
}

func cost(b, w, x, y, z int) int {
	return b*min(x, y+z) + w*min(y, x+z)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}