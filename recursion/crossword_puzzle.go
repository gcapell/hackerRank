package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func main() {
	const boardSize = 10
	board := make([]string, boardSize)
	for j := 0; j < boardSize; j++ {
		fmt.Scanln(&board[j])
	}

	ss := slots(board)

	var wordIn string
	fmt.Scanln(&wordIn)
	words := strings.Split(wordIn, ";")

	available := newAvailable(words)

	if !solve(ss, available) {
		log.Fatal("no solution")
	}
	out := make([][]byte, boardSize)
	for j := 0; j < boardSize; j++ {
		out[j] = bytes.Repeat([]byte("+"), boardSize)
	}
	for _, s := range ss {
		p := s.head
		for j := 0; j < s.length; j++ {
			out[p.r][p.c] = s.word[j]
			if s.down {
				p.r++
			} else {
				p.c++
			}
		}
	}
	for _, line := range out {
		fmt.Println(string(line))
	}
}

func solve(ss []*slot, words availableStrings) bool {
	s, ws := mostConstrained(ss, words)
	if len(ws) == 0 {
		return s == nil
	}
	for _, w := range ws {
		s.word = w
		words.remove(w)
		if solve(ss, words) {
			return true
		}
		s.word = ""
		words.add(w)
	}
	return false
}

type availableStrings map[int]map[string]struct{}

func newAvailable(words []string) availableStrings {
	a := availableStrings(make(map[int]map[string]struct{}))
	for _, w := range words {
		a.add(w)
	}
	return a
}

func (a availableStrings) remove(w string) {
	delete(a[len(w)], w)
}

func (a availableStrings) add(w string) {
	m, ok := a[len(w)]
	if !ok {
		m = make(map[string]struct{})
		a[len(w)] = m
	}
	m[w] = struct{}{}
}

func mostConstrained(slots []*slot, words availableStrings) (*slot, []string) {
	var smallest []string
	var mostConstrainedSlot *slot
	first := true

	for _, s := range slots {
		if s.word != "" {
			continue
		}
		var possible []string
		for w, _ := range words[s.length] {
			if !s.conflict(w) {
				possible = append(possible, w)
			}
		}
		if first || len(possible) < len(smallest) {
			first = false
			smallest = possible
			mostConstrainedSlot = s
		}
	}
	return mostConstrainedSlot, smallest
}

func (s slot) conflict(w string) bool {
	for _, c := range s.conflicts {
		if c.peer.word != "" && w[c.me] != c.peer.word[c.them] {
			if false {
				fmt.Printf("%s[%d]=%s, %s[%d]=%s\n",
					w, c.me, string(w[c.me]),
					c.peer.word, c.them, string(c.peer.word[c.them]))
			}
			return true
		}
	}
	return false
}

func (s *slot) String() string {
	return fmt.Sprintf("%d@%v:%s", s.length, s.head, s.word)
}

func (d *slot) addConflict(r *slot) {
	if !(d.head.c >= r.head.c && d.head.c <= r.head.c+r.length &&
		r.head.r >= d.head.r && r.head.r <= d.head.r+d.length) {
		return
	}
	dOff := r.head.r - d.head.r
	rOff := d.head.c - r.head.c
	d.conflicts = append(d.conflicts, conflict{r, dOff, rOff})
	r.conflicts = append(r.conflicts, conflict{d, rOff, dOff})
}

type pos struct {
	r, c int
}

type conflict struct {
	peer     *slot
	me, them int
}
type slot struct {
	head      pos
	length    int
	conflicts []conflict
	word      string
	down      bool
}

func slots(board []string) []*slot {
	var downSlots, rightSlots []*slot
	// Map from members of a slot to head of slot
	downMembers := make(map[pos]pos)
	rightMembers := make(map[pos]pos)

	for r, line := range board {
		for c, ch := range line {
			if ch == '+' {
				continue
			}
			if ch != '-' {
				log.Fatal(r, c, ch, line, board)
			}
			if s, ok := down(r, c, board, downMembers); ok {
				downSlots = append(downSlots, s)
			}
			if s, ok := right(r, c, board, rightMembers); ok {
				rightSlots = append(rightSlots, s)
			}
		}
	}
	for _, d := range downSlots {
		for _, r := range rightSlots {
			d.addConflict(r)
		}
	}
	return append(downSlots, rightSlots...)
}

func down(r, c int, board []string, members map[pos]pos) (*slot, bool) {
	s := &slot{head: pos{r, c}, down: true}
	if _, ok := members[s.head]; ok {
		return s, false
	}
	j := r + 1
	for ; j < len(board) && board[j][c] == '-'; j++ {
		members[pos{j, c}] = s.head
	}
	s.length = j - r
	return s, s.length > 1
}

func right(r, c int, board []string, members map[pos]pos) (*slot, bool) {
	s := &slot{head: pos{r, c}, down: false}
	if _, ok := members[s.head]; ok {
		return s, false
	}
	j := c + 1
	for ; j < len(board[r]) && board[r][j] == '-'; j++ {
		members[pos{r, j}] = s.head
	}
	s.length = j - c
	return s, s.length > 1
}
