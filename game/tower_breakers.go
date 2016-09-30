package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		reply := 1
		if m == 1 || n%2 == 0 {
			reply = 2
		}
		fmt.Println(reply)
	}
}
