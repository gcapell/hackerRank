package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	g := NewSuffixGenerator(a, 1, false)
	l := g.next()
	for !l.empty() {
		c, pos, id := l.next()
		fmt.Println(string(c), pos, id)
	}

	fmt.Println("*****************")
	l = g.next()
	for !l.empty() {
		c, pos, id := l.next()
		fmt.Println(string(c), pos, id)
	}
	pals(a)

	//pals(b)
}

// pals prints palindromes found in s
func pals(s string) {
	root := New()
	root.addSuffixes(NewSuffixGenerator(s, 0, true))
	root.addSuffixes(NewSuffixGenerator(s, 1, false))
	root.pprint("")
}

type node struct {
	letter map[byte]*node
	paths  map[int][]int
}

func New() *node {
	return &node{
		letter: make(map[byte]*node),
		paths:  make(map[int][]int),
	}
}

// addSuffixes adds suffixes of s
func (n *node) addSuffixes(s *suffixGenerator) {
	for !s.empty() {
		n.add(s.next())
	}
}

func (n *node) pprint(indent string) {
	ps := fmt.Sprintf("%v", n.paths)
	fmt.Printf("%s%s\n", indent, ps[3:])
	for c, next := range n.letter {
		fmt.Printf("%s%s\n", indent, string(c))
		next.pprint(indent + "  ")
	}
}

type stringIterator struct {
	s       string
	pos     int
	forward bool
	id      int
}

type letterGenerator struct {
	stringIterator
}

func (g *stringIterator) empty() bool {
	if g.forward {
		return g.pos == len(g.s)
	}
	return g.pos == -1
}

func (g *stringIterator) advance() {
	if g.forward {
		g.pos++
	} else {
		g.pos--
	}
}

func (g *letterGenerator) next() (byte, int, int) {
	defer func() { g.advance() }()
	return g.s[g.pos], g.pos, g.id
}

type suffixGenerator struct {
	stringIterator
}

func NewSuffixGenerator(s string, id int, forward bool) *suffixGenerator {
	pos := 0
	if !forward {
		pos = len(s) - 1
	}
	return &suffixGenerator{
		stringIterator{
			s:       s,
			pos:     pos,
			forward: forward,
			id:      id,
		}}
}

func (s *suffixGenerator) next() *letterGenerator {
	defer func() { s.advance() }()
	return &letterGenerator{stringIterator{
		s.s,
		s.pos,
		s.forward,
		s.id,
	}}
}

func (n *node) add(g *letterGenerator) {
	c, pos, id := g.next()
	next := n.letter[c]
	if next == nil {
		next = New()
		n.letter[c] = next
	}
	next.paths[id] = append(next.paths[id], pos)
	if !g.empty() {
		next.add(g)
	}
}
