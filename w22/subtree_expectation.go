package main

import (
	"fmt"
	"strconv"
	"strings"
)

const debug = false

func dprint(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}

func main() {
	var q int
	fmt.Scanf("%d", &q)
	for j := 0; j < q; j++ {
		fmt.Println(query())
	}
}

func query() float64 {
	var vertices int
	fmt.Scanf("%d", &vertices)
	weights := nInts(vertices)
	scores := nInts(sum(weights) + 1)
	edges := nInts(2 * (vertices - 1))

	root := newTree(weights, edges)
	dist, excluding := root.weightDistribution(nil)
	dist.sum(excluding)
	dprint("dist", dist)
	var subtrees, score int
	for w, count := range dist {
		subtrees += count
		score += count * scores[w]
	}
	return float64(score) / float64(subtrees)
}

func newTree(weights, edges []int) *node {
	nodes := make([]node, len(weights))
	for j, w := range weights {
		nodes[j].w = w
		nodes[j].id = j
	}
	for j := 0; j < len(edges); j += 2 {
		nodes[edges[j]-1].link(&nodes[edges[j+1]-1])
	}
	if debug {
		fmt.Println("nodes")
		for _, n := range nodes {
			fmt.Println(n)
		}
	}
	return &nodes[0]
}

func (n *node) link(o *node) {
	n.peer = append(n.peer, o)
	o.peer = append(o.peer, n)
}

type node struct {
	w       int
	id      int
	visited bool
	peer    []*node
}

func (n node) String() string {
	var peerStrings []string
	for _, p := range n.peer {
		peerStrings = append(peerStrings, strconv.Itoa(p.id))
	}
	return fmt.Sprintf("%d:%d[%s]", n.id, n.w,
		strings.Join(peerStrings, ","))
}

func (n *node) safeId() int {
	if n == nil {
		return -1
	}
	return n.id
}

type dist map[int]int

// Return distribution of weights including/excluding n
func (n *node) weightDistribution(parent *node) (including, excluding dist) {
	dprint("weightDist", n.id, parent.safeId())
	including = map[int]int{n.w: 1}
	excluding = map[int]int{}
	n.visited = true

	for _, c := range n.peer {
		if c.visited {
			continue
		}
		cInc, cExc := c.weightDistribution(n)
		excluding.sum(cExc)
		excluding.sum(cInc)
		including = combine(including, cInc)
	}
	dprint("dist", n.id, including, excluding)
	return including, excluding
}

func combine(as, bs map[int]int) map[int]int {
	reply := make(map[int]int)
	for aw, ac := range as {
		reply[aw] = ac
	}

	for aw, ac := range as {
		for bw, bc := range bs {
			reply[aw+bw] += ac * bc
		}
	}
	dprint("combine(", as, ",", bs, ")->", reply)
	return reply
}

func (a *dist) sum(b dist) {
	for k, v := range b {
		(*a)[k] += v
	}
}

func nInts(n int) []int {
	reply := make([]int, n)
	for j := range reply {
		fmt.Scanf("%d", &reply[j])
	}
	return reply
}

func sum(w []int) int {
	total := 0
	for _, j := range w {
		total += j
	}
	return total
}
