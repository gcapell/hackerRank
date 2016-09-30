package main

import "fmt"

type node struct {
	id    int
	child []*node
}

var evens int

func main() {
	var vertices, edges int
	fmt.Scanf("%d %d", &vertices, &edges)

	nodes := make([]node, vertices)
	for j := range nodes {
		nodes[j].id = j + 1
	}
	for j := 0; j < edges; j++ {
		var u, v int
		fmt.Scanf("%d %d", &u, &v)
		link(&nodes[u-1], &nodes[v-1])
	}
	size(&nodes[0], nil)
	fmt.Println(evens)
}

func link(a, b *node) {
	a.child = append(a.child, b)
	b.child = append(b.child, a)
}

func size(n, parent *node) int {
	if len(n.child) == 1 {
		return 1
	}
	total := 1
	for _, c := range n.child {
		if c == parent {
			continue
		}
		total += size(c, n)
	}
	if total%2 == 0 && parent != nil {
		evens++
	}
	return total
}
