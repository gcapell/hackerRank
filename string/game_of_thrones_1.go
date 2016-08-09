package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(map[bool]string{
		true: "YES", false: "NO",
	}[canPal(s)])
}

func canPal(s string) bool {
	odd := make(map[rune]bool)
	for _, r := range s {
		odd[r] = !odd[r]
	}
	odds := 0
	for _, v := range odd {
		if !v {
			continue
		}
		odds++
		if odds > 1 {
			return false
		}
	}

	return true
}
