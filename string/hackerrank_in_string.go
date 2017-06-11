package main

import "fmt"

var yes = map[bool]string{
	false: "NO",
	true:  "YES",
}

func main() {
	var n int
	fmt.Scanln(&n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanln(&s)
		fmt.Println(yes[contains(s, "hackerrank")])
	}
}

func contains(s, target string) bool {
	for _, c := range s {
		if c == rune(target[0]) {
			target = target[1:]
			if len(target) == 0 {
				return true
			}
		}
	}
	return false
}
