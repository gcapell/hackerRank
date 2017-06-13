package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"sort"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	debug.SetGCPercent(-1) // disable GC
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var n int
	r := bufio.NewReaderSize(os.Stdin, 1<<20)
	fmt.Fscan(r, &n)
	gene := make([]string, n)
	for j := 0; j < n; j++ {
		fmt.Fscan(r, &gene[j])
	}
	health := make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Fscan(r, &health[j])
	}

	dfa := ahoCorasick(gene, health)

	var strands int
	fmt.Fscan(r, &strands)

	var result extremes
	for j := 0; j < strands; j++ {
		var start, end int
		var dna string
		fmt.Fscan(r, &start, &end, &dna)
		n := score(dfa, j+1, start, end, dna)
		//fmt.Printf("%s %d\n", dna, n)
		result.update(n)
	}
	fmt.Println(result.min, result.max)
}

type extremes struct {
	init     bool
	min, max int
}

func (e *extremes) update(n int) {
	if !e.init {
		e.init = true
		e.min = n
		e.max = n
	}
	if n > e.max {
		e.max = n
	}
	if n < e.min {
		e.min = n
	}
}

type geneScore struct {
	id, cumScore int
}

type scoreCache struct {
	id, val int
}

type node struct {
	suffix, parent, next *node

	edge       [26]*node
	health     []geneScore
	inEdge     byte
	scoreCache scoreCache
}

func ahoCorasick(gene []string, health []int) *node {
	root := new(node)
	for j := range gene {
		addTrie(root, j, gene[j], health[j])
	}
	bfs(root, addSuffixes)
	return root
}

func addSuffixes(child *node) {
	p, e := child.parent, child.inEdge

	for {
		if p.suffix == nil {
			child.suffix = p
			return
		}
		p = p.suffix
		if s := p.edge[e]; s != nil {
			child.suffix = s
			return
		}
	}
}

type queue struct {
	head, tail *node
}

var q queue

func bfs(n *node, f func(n *node)) {
	q.addChildren(n)
	for q.tail != nil {
		child := q.tail
		q.tail = q.tail.next
		q.addChildren(child)
		f(child)
	}
}

func (q *queue) addChildren(n *node) {
	for _, child := range n.edge {
		if child == nil {
			continue
		}
		if q.head == nil {
			q.head = child
			q.tail = child
			n.next = nil
		} else {
			q.head.next = child
			q.head = child
		}
	}
}

func addTrie(n *node, id int, gene string, score int) {
	if n == nil {
		panic("n is nil?")
	}
	if len(gene) == 0 {
		cumScore := score
		if len(n.health) > 0 {
			cumScore += n.health[len(n.health)-1].cumScore
		}
		n.health = append(n.health, geneScore{id, cumScore})
		return
	}
	c := gene[0] - 'a'
	child := n.edge[c]
	if child == nil {
		child = new(node)
		child.parent = n
		child.inEdge = c
		n.edge[c] = child
	}
	addTrie(child, id, gene[1:], score)
}

func score(n *node, id, start, end int, dna string) int {
	total := 0
	for len(dna) > 0 {
		n = n.follow(dna[0] - 'a')
		total += n.score(id, start, end)
		dna = dna[1:]
	}
	return total
}

func (n *node) follow(c byte) *node {
	for {
		if child := n.edge[c]; child != nil {
			return child
		}
		if n.suffix == nil {
			return n
		}
		n = n.suffix
	}
}

func (n *node) score(id, start, end int) int {
	total := 0
	for n != nil {
		if id == n.scoreCache.id {
			total += n.scoreCache.val
		} else if len(n.health) > 0 {
			subtotal := sumRange(n.health, start, end)
			n.scoreCache.val = subtotal
			n.scoreCache.id = id
			total += subtotal
		}
		n = n.suffix
	}
	return total
}

func sumRange(health []geneScore, start, end int) int {
	size := len(health)
	tooBig := sort.Search(size, func(j int) bool {
		return health[j].id > end
	})
	if tooBig == 0 {
		return 0
	}
	bigEnough := sort.Search(size, func(j int) bool {
		return health[j].id >= start
	})
	if bigEnough == size {
		return 0
	}

	maybeNotTooBig := tooBig - 1
	if health[maybeNotTooBig].id > end {
		return 0
	}
	subtotal := health[maybeNotTooBig].cumScore
	if bigEnough > 0 {
		subtotal -= health[bigEnough-1].cumScore
	}
	return subtotal
}
