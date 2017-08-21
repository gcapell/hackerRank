package main

import "fmt"

var yesno = map[bool]string{
	true:  "YES",
	false: "NO",
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(yesno[valid(counts(s))])
}

func valid(counts map[int]int) bool {
	switch len(counts) {
	case 0, 1:
		return true
	case 2: // well, let's see
	default:
		return false
	}

	var ks, vs []int
	for k, v := range counts {
		if k == 1 && v == 1 {
			return true // totally remove a character
		}
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return vs[0] == 1 && ks[0] == ks[1]+1 ||
		vs[1] == 1 && ks[1] == ks[0]+1
}

func counts(s string) map[int]int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	reply := make(map[int]int)
	for _, v := range count {
		reply[v]++
	}
	return reply
}
