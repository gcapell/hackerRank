package main

import (
	"fmt"
	"regexp"
	"strings"
)

var yn = map[bool]string{
	false: "NO",
	true:  "YES",
}

func main() {
	var t int
	fmt.Scanln(&t)
	for j := 0; j < t; j++ {
		fmt.Println(yn[gridSearch()])
	}
}

func gridSearch() bool {
	var r, c int
	fmt.Scanln(&r, &c)
	rows := readRows(r)
	grid := strings.Join(rows, "")

	var pr, pc int
	fmt.Scanln(&pr, &pc)
	rows = readRows(pr)
	pattern := strings.Join(rows, strings.Repeat(".", c-pc))
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringIndex(grid, -1)
	for _, m := range matches {
		if m[0]%c <= c-pc {
			return true
		}
	}
	return false

}

func readRows(n int) []string {
	var rows []string
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanln(&s)
		rows = append(rows, s)
	}
	return rows
}
