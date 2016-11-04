package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	receive := 5
	total := 0
	for j := 0; j < n; j++ {
		like := receive / 2
		total += like
		receive = like * 3
	}

	fmt.Println(total)
}
