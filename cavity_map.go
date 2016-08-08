package main

import "fmt"

func main() {
	var s int
	fmt.Scanf("%d", &s)
	map := make([][]int, s)
	for r := range map {
	
		row := make([]int, s)
		for c := range row {
			fmt.Scanf("%d
