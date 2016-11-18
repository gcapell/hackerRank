package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	fmt.Scanln(&s)
	f := math.Sqrt(float64(len(s)))
	r := int(math.Floor(f))
	c := int(math.Ceil(f))
	if r*c < len(s) {
		r++
	}
	//fmt.Println(r, c)
	for j := 0; j < c; j++ {
		for k := 0; k < r; k++ {
			pos := k*c + j
			if pos < len(s) {
				fmt.Print(string(s[pos]))
			}
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
