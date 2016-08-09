package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	pals(a)
	pals(b)
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
func (n *node) addSuffixes(s string, id int) {
	for j := 0; j < len(s); j++ {
		n.add(s[j:], id, j)
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

func (n *node) add(s string, id, start int) {
	c, s := s[0], s[1:]
	next := n.letter[c]
	if next == nil {
		next = New()
		n.letter[c] = next
	}
	next.paths[id] = append(next.paths[id], start)
	if len(s) > 0 {
		next.add(s, id, start)
	}
}

// pals prints palindromes found in s
func pals(s string) {
	root := New()
	root.addSuffixes(s, 0)
	root.addSuffixes(reverse(s), 1)
	root.pprint("")
}

func reverse(s string) string {
	b := []byte(s)
	for j, k := 0, len(b)-1; j < k; j, k = j+1, k-1 {
		b[j], b[k] = b[k], b[j]
	}
	return string(b)
}
