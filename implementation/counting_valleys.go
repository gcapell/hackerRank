package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var level, valleys int
	for j := 0; j < n; j++ {
		var c string
		fmt.Scanf("%1s", &c)
		switch c {
		case "U":
			level++

		case "D":
			if level == 0 {
				valleys++
			}
			level--
		default:
			log.Fatal(c)

		}
	}
	fmt.Println(valleys)
}
