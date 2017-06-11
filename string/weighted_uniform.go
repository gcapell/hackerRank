package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	w := weights(s)
	var queries int
	fmt.Scanln(&queries)
	for j := 0; j < queries; j++ {
		var score int
		fmt.Scanln(&score)
		reply := "No"
		if _, ok := w[score]; ok {
			reply = "Yes"
		}
		fmt.Println(reply)
	}
}

func weights(s string) map[int]bool {
	reply := make(map[int]bool)
	total := 0
	prev := 'Z'
	for _, c := range s {
		if c != prev {
			total = 0
		}
		total += int(c - 'a' + 1)
		reply[total] = true
		prev = c
	}
	return reply
}
