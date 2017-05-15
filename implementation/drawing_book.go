package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Scan(&n, &p)
	reply := p / 2

	if n%2 == 0 {
		n++
	}
	backward := (n - p) / 2

	if backward < reply {
		reply = backward
	}
	fmt.Println(reply)

}
