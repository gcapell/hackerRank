package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var n int
		fmt.Scanf("%d", &n)
		var xor int
		for j := 0; j < n; j++ {
			var d int
			fmt.Scanf("%d", &d)
			xor ^= d
		}
		reply := "First"

		if xor == 0 {
			reply = "Second"
		}
		fmt.Println(reply)
	}
}
