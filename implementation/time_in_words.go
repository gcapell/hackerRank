package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, m int
	fmt.Scanln(&h)
	fmt.Scanln(&m)
	fmt.Println(words(h, m))
}

var word = strings.Fields(`
	zero one two three four five six
	seven eight nine ten eleven twelve thirteen
	fourteen fifteen sixteen seventeen eighteen
	nineteen twenty twenty-one twenty-two twenty-three
	twenty-four twenty-five twenty-six twenty-seven
	twenty-eight twenty-nine
`)

func words(h, m int) string {
	switch {
	case m == 0:
		return fmt.Sprintf("%s o' clock", word[h])
	case m == 15:
		return fmt.Sprintf("quarter past %s", word[h])
	case m == 30:
		return fmt.Sprintf("half past %s", word[h])
	case m == 45:
		return fmt.Sprintf("quarter to %s", word[succ(h)])
	case m < 30:
		return fmt.Sprintf("%s minute%s past %s", removeDash(word[m]), plural(m), word[h])
	default:
		return fmt.Sprintf("%s minute%s to %s", removeDash(word[60-m]), plural(60-m), word[succ(h)])
	}
}

func removeDash(s string) string {
	return strings.Replace(s, "-", " ", 1)
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func succ(n int) int {
	n++
	if n == 13 {
		n = 1
	}
	return n
}
