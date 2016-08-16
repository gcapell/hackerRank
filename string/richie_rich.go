package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d %d\n%s", &n, &k, &s)
	fmt.Println(maxPal(s, k))
}

func maxPal(s string, ch int) string {
	b := []byte(s)
	changed := make(map[int]bool)
	// make palindrome
	for j, k := 0, len(b)-1; j < k; j, k = j+1, k-1 {
		if b[j] == b[k] {
			continue
		}
		if ch == 0 {
			return "-1"
		}
		if b[j] > b[k] {
			b[k] = b[j]
		} else {
			b[j] = b[k]
		}
		changed[j] = true
		ch--
	}

	// make bigger
	for j, k := 0, len(b)-1; ch > 0 && j <= k; j, k = j+1, k-1 {
		if b[j] == '9' {
			continue
		}
		switch {
		case changed[j], j == k:
			ch--
		case ch > 1:
			ch -= 2
		default:
			continue
		}
		b[j], b[k] = '9', '9'
	}

	return string(b)
}
