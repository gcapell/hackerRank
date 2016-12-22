package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var s string
	fmt.Scanln(&s)
	fmt.Println(twoCharLen(s))
}

type pair struct {
	last rune
	len  int
}

func twoCharLen(s string) int {
	byRune := make([][]*pair, 26)
	var all []*pair
	for j := 'a'; j <= 'z'; j++ {
		for k := j + 1; k <= 'z'; k++ {
			p := &pair{}
			byRune[j-'a'] = append(byRune[j-'a'], p)
			byRune[k-'a'] = append(byRune[k-'a'], p)
			all = append(all, p)
		}
	}
	for _, r := range s {
		for _, p := range byRune[r-'a'] {
			p.add(r)
		}
	}
	max := 0
	for _, p := range all {
		if p.last != '!' && p.len > max {
			max = p.len
		}
	}
	if max == 1 {
		max = 0
	}
	return max
}

func (p *pair) add(r rune) {
	switch p.last {
	case r:
		p.last = '!'
	case '!':
	default:
		p.last = r
	}
	p.len++
}
