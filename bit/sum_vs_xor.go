package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	reply := 1
	for ; n > 0; n /= 2 {
		if n%2 == 0 {
			reply *= 2
		}
	}
	fmt.Println(reply)
}
