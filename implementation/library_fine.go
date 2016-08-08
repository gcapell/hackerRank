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

type date struct {
	d, m, y int
}

func main() {
	var a, b date
	scanln(&a.d, &a.m, &a.y)
	scanln(&b.d, &b.m, &b.y)
	fmt.Println(fine(a, b))
}

func fine(returned, expected date) int {
	switch {
	case returned.y > expected.y:
		return 10000
	case returned.y < expected.y || returned.m < expected.m:
		return 0
	case returned.m > expected.m:
		return 500 * (returned.m - expected.m)
	case returned.d > expected.d:
		return 15 * (returned.d - expected.d)
	}
	return 0
}
