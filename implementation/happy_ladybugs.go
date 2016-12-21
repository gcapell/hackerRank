package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	response := map[bool]string{
		true:  "YES",
		false: "NO",
	}
	for j := 0; j < n; j++ {
		fmt.Println(response[game()])
	}
}

func game() bool {
	var n int
	var s string
	fmt.Scanln(&n)
	fmt.Scanln(&s)
	count := make(map[rune]int)
	hasGaps := false
	needsGaps := false
	var prev rune
	lastHappy := true
	for _, r := range s {
		if r == '_' {
			hasGaps = true
			continue
		}
		count[r]++

		if r == prev {
			lastHappy = true
		} else {
			if !lastHappy {
				needsGaps = true
			}
			lastHappy = false
		}
		prev = r
	}
	if !lastHappy {
		needsGaps = true
	}
	//fmt.Println(needsGaps, hasGaps, count)
	if !needsGaps {
		return true
	}
	if !hasGaps {
		return false
	}
	for _, v := range count {
		if v == 1 {
			return false
		}
	}
	return true
}
