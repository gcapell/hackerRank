package main

import "fmt"

func main() {
	var s, t string
	var k int
	fmt.Scan(&s, &t, &k)
	c := commonPrefixLength(s, t)
	delta := len(s) + len(t) - 2*c
	r := "No"
	switch {
	case delta > k:
		r = "No"
	case delta%2 == k%2:
		r = "Yes"
	case len(s)+len(t) < k:
		r = "Yes"
	}
	fmt.Println(r)

}

func commonPrefixLength(s, t string) int {
	j := 0
	for j < len(s) && j < len(t) {
		if s[j] != t[j] {
			break
		}
		j++
	}
	return j
}
