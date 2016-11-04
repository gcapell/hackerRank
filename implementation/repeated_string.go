package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n int
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &n)

	whole := n / len(s)
	remainder := n % len(s)
	wholeA := asIn(s)
	remainderA := asIn(s[:remainder])
	fmt.Println(whole*wholeA + remainderA)
}

func asIn(s string) int {
	return strings.Count(s, "a")
}
