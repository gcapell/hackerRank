package main

import "fmt"

func main() {
	var q int
	fmt.Scanln(&q)
	for j := 0; j < q; j++ {
		var x, y, z int
		fmt.Scanln(&x, &y, &z)
		a := delta(x, z)
		b := delta(y, z)
		switch {
		case a == b:
			fmt.Println("Mouse C")
		case a < b:
			fmt.Println("Cat A")
		default:
			fmt.Println("Cat B")
		}
	}
}

func delta(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
