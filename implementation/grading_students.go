package main

import "fmt"

func main() {
	var students int
	fmt.Scanln(&students)
	for j := 0; j < students; j++ {
		var grade int
		fmt.Scanln(&grade)
		fmt.Println(round(grade))
	}
}

func round(g int) int {
	if g < 38 {
		return g
	}
	switch g % 5 {
	case 3, 4:
		return ((g / 5) + 1) * 5
	}
	return g
}
