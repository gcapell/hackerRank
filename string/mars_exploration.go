package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	var delta int
	for j := range s {
		if s[j] != "SOS"[j%3] {
			delta++
		}
	}
	fmt.Println(delta)
}
