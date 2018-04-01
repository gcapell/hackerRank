package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var q int
	fmt.Scanln(&q)
	for j := 0; j < q; j++ {
		var x uint64
		fmt.Scanln(&x)
		fmt.Println(1<<uint(bits.Len64(x)) - 1 - x)
	}
}
