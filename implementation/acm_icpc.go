package main

import (
	"fmt"
	"log"
)

func scanln(a ...interface{}) {
	if _, err := fmt.Scanln(a...); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var n, m int
	scanln(&n, &m)
	lines := make([]string, n)
	for j := range lines {
		scanln(&lines[j])
	}
	topics, teams := max(lines)
	fmt.Println(topics)
	fmt.Println(teams)
}

func max(lines []string) (topics, teams int) {
	for j := 0; j < len(lines); j++ {
		for k := j + 1; k < len(lines); k++ {
			o := overlap(lines[j], lines[k])
			if o > topics {
				topics, teams = o, 1
			} else if o == topics {
				teams++
			}
		}
	}
	return
}

func overlap(a, b string) int {
	var reply int
	for j := 0; j < len(a); j++ {
		if a[j] == '1' || b[j] == '1' {
			reply++
		}
	}
	return reply
}
