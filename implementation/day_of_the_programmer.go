package main

import (
	"fmt"
)

func main() {
	var y int
	fmt.Scanln(&y)
	d := 13
	switch {
	case y == 1918:
		d = 26
	case (y < 1918 && y%4 == 0) || (y%400 == 0 || (y%4 == 0 && y%100 != 0)):
		d--
	}
	fmt.Printf("%d.09.%d\n", d, y)
}
