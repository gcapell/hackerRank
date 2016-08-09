package main

import "fmt"

var yes = map[bool]string{
	false: "NO",
	true:  "YES",
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var a, b string
		fmt.Scanf("%s\n%s", &a, &b)

		fmt.Println(yes[common(a, b)])
	}
}

func common(a, b string) bool {
	al := make(map[byte]bool)
	bl := make(map[byte]bool)
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for j := 0; j < minLen; j++ {
		al[a[j]] = true
		if bl[a[j]] {
			return true
		}
		bl[b[j]] = true
		if al[b[j]] {
			return true
		}
	}
	for _, r := range a[minLen:] {
		if bl[byte(r)] {
			return true
		}
	}
	for _, r := range b[minLen:] {
		if al[byte(r)] {
			return true
		}
	}
	return false
}
