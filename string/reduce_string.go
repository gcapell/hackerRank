package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	r := reduce(s)
	if len(r) == 0 {
		r = "Empty String"
	}

	fmt.Println(r)
}

func reduce(s string) string {
	if len(s) < 2 {
		return s
	}
	a := []byte(s)
	p := 1 // next position to add
	for j := 1; j < len(a); j++ {
		if p == 0 || a[p-1] != a[j] {
			a[p] = a[j]
			p++
		} else {
			p--
		}
	}
	return string(a[:p])
}
