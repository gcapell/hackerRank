package main

import (
	"fmt"
	"math/big"
)

func main() {
	var t1, t2, n int

	fmt.Scanf("%d %d %d", &t1, &t2, &n)

	b1 := big.NewInt(int64(t1))
	b2 := big.NewInt(int64(t2))
	c := big.NewInt(0)
	for j := 0; j < n-2; j++ {
		c.Mul(b2, b2)
		c.Add(c, b1)
		b1, b2, c = b2, c, b1
	}
	fmt.Printf("%s\n", b2)
}
