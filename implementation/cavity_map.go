package main

import (
	"fmt"
	"log"
)

func scanln(a ...interface{}) {
	if _, err := fmt.Scanln(a...); err != nil {
		log.Fatal(err)
	}
}

func mapInts(s string) []int {
	var reply []int
	for _, r := range s {
		reply = append(reply, int(r)-'0')
	}
	return reply
}

func main() {
	var s int
	scanln(&s)
	a := make([][]int, s)

	for j := range a {
		var line string
		scanln(&line)
		if len(line) != s {
			log.Fatal(line)
		}
		a[j] = mapInts(line)
	}
	hole := cavities(a)
	for r, row := range a {
		for c, n := range row {
			s := '0' + n
			if hole[coord{r, c}] {
				s = 'X'
			}
			fmt.Printf("%c", s)
		}
		fmt.Println()
	}
}

type coord struct {
	r, c int
}

func cavities(a [][]int) map[coord]bool {
	reply := make(map[coord]bool)

	for r := 1; r < len(a)-1; r++ {
		for c := 1; c < len(a[0])-1; c++ {
			n := a[r][c]
			if n > a[r-1][c] &&
				n > a[r+1][c] &&
				n > a[r][c-1] &&
				n > a[r][c+1] {
				reply[coord{r, c}] = true
			}
		}
	}
	return reply
}
