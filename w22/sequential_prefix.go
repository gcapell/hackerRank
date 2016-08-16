package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	seq := sequence{}

	for j := 0; j < n; j++ {
		var cmd string
		fmt.Scanf("%s", &cmd)
		if cmd == "+" {
			var d int
			fmt.Scanf("%d", &d)
			fmt.Println(seq.push(d))
		} else {
			fmt.Println(seq.pop())
		}
	}
}

type sequence struct {
	start, end int // half-open range.  -1 -> open
}

type transition struct {
	from, to int
}

type stateMachine struct {
	sequence []int
	runs     []sequence // longest non-overlapping runs found
	// start of a second-longest run we are tracking
	// -1 if we're not tracking
	secondaryStart int //

	// Inflection points.  Sequence of length 'to' is a suffix of
	// sequence of length 'from', e.g.
	// for "abacabad"
	// "a" is a suffix of "aba"
	// "aba" is a suffix of "abacaba"
	// so we would have entries 3->1, 7->3
	transitions []transition // inflection points
}

func (s *stateMachine) push(added int) int {
	s.seq = append(s.seq, added)
	if len(s.seq) == 1 {
		return
	}
	if openRun(s.runs) {
		s.extend(added)
	} else if next == s.seq[0] {
		// start a run
		s.runs = append(s.runs, sequence{len(s.seq) - 1, -1})
	}
}

func (s *stateMachine) extend(added int) {
	openRun = s.runs[len(s.runs)-1]
	expected := s.seq[len(s.seq)-current.start]
	if added != expected {
		s.closeRun()
		return
	}
	if s.secondaryStart != -1 {
		expected := s.seq[len(s.seq)-secondaryStart]
		if added != expected {
			s.closeSecondary()
		}
	} else if added == s.seq[0] {
		s.secondaryStart = len(s.seq) - 1
	}
}

func openRun(runs []sequence) bool {
	return len(runs) > 0 && runs[len(run)-1].end == -1
}

func (s *sequence) pop() int {
	s.seq = s.seq[:len(s.seq)-1]
	s.sub = s.sub[:len(s.sub)-1]
	if len(s.sub) > 0 {
		return s.sub[len(s.sub)-1]
	}
	return 0
}

// If we've seen p elements from s, and we add e,
// what's the longest prefix of s that we've now got?
func nextLongest(s []int, p int, e int) int {
	// extend longest - easy
	if s[p] == e {
		return p + 1
	}

	// does some suffix of s[:p], when we append e,
	// form a prefix of s?

	// previous occurrence of e?
finde:
	for j := p - 1; j >= 0; j-- {
		if s[j] == e {
			for k := j - 1; k >= 0; k-- {
				if s[k] != s[p-(j-k)] {
					continue finde
				}
			}
			return j + 1
		}
	}
	return 0
}
