package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for j := 0; j < n; j++ {
		var grade int
		fmt.Scanln(&grade)
		fmt.Println(round(grade))
	}
}

func round(g int) int {
	if g >= 38 {
		switch g % 5 {
		case 3:
			return g + 2
		case 4:
			return g + 1
		}
	}
	return g
}
