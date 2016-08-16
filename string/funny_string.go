package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanf("%s", &s)
		if !funny(s) {
			fmt.Printf("Not ")
		}
		fmt.Println("Funny")
	}
}

func funny(s string) bool {
	for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
		dj := s[j] - s[j+1]
		dk := s[k] - s[k-1]
		if !(dj == dk || dj == -dk) {
			return false
		}
	}
	return true
}
