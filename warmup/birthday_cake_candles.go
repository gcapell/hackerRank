package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var max, count int

	for j := 0; j < n; j++ {
		var t int
		fmt.Scan(&t)
		switch {
		case t > max:
			max = t
			count = 1
		case t == max:
			count++
		}
	}
	fmt.Println(count)
}
