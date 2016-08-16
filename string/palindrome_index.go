package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(palindromeIndex(s))
	}
}

func palindromeIndex(s string) (index int) {
	for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
		if s[j] == s[k] {
			continue
		}
		switch {
		case s[j+1] == s[k] && s[j] == s[k-1]:
			switch {
			case palindrome(s[j+1 : k+1]):
				return j
			case palindrome(s[j:k]):
				return k
			}
		case s[j+1] == s[k]:
			if palindrome(s[j+1 : k+1]) {
				return j
			}
		case s[j] == s[k-1]:
			if palindrome(s[j:k]) {
				return k
			}
		}
		return -1
	}
	return -1
}

func palindrome(s string) bool {
	for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
		if s[j] != s[k] {
			return false
		}
	}
	return true
}
