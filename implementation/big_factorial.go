package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	b, c, d := big.NewInt(1), big.NewInt(1), big.NewInt(1)
	for j := 0; j < n-1; j++ {
		c.Add(c, d)
		b.Mul(b, c)
	}
	fmt.Println(b)
}
