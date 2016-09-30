package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for j := 0; j < t; j++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		s := "First"
		if ((x-1)/2)%2 ==0 && ((y-1)/2)%2 ==0 {
			s = "Second"
		}
		fmt.Println(s)
	}
}
