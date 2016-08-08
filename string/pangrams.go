package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	if !pangram(line) {
		fmt.Print("not ")
	}
	fmt.Println("pangram")
}

func pangram(line string) bool {
	letters := make(map[rune]bool)
	for _, r := range line {
		letters[unicode.ToLower(r)] = true
	}
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if !letters[r] {
			return false
		}
	}
	return true
}
