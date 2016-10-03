package main

import "fmt"

type node struct {
	parent   int
	children []int
	depths   []int
}

var nodes []node

func main() {
	var n int
	fmt.Scanln(&n)
	nodes = make([]node, n+1)
	for j := 2; j <= n; j++ {
		var parent int
		fmt.Scanf("%d", &parent)
		setParent(j, parent)
	}
	countDepths(1)
	pprint(1, "")

	var q int
	fmt.Scanln(&q)
	for j := 0; j < q; j++ {
		var u, v int
		_, err := fmt.Scanln(&u, &v)
		if err != nil {
			fmt.Println("scanln:", err)
		}
		fmt.Println(gravity(u, v))
	}
}

func gravity(u, v int) int {
	fmt.Printf("g(%d,%v)\n", u, v)
	p, isAncestor := path(u, v)
	fmt.Println(p, isAncestor)
	if !isAncestor {
		d := len(p) + 1
		reply := d * d
		fmt.Printf("adding %d^2\n", d)
		for pos, v := range nodes[v].depths {
			fmt.Printf("adding %d * %d^2\n", v, pos+d+1)
			reply += v * (pos + d + 1) * (pos + d + 1)
		}
		return reply
	}
	// ancestor is tricky

	return 0
}

// list nodes from u to v (excluding u and v).
// indicate if v is one of u's parents
func path(u, v int) ([]int, bool) {
	uParents := parents(u)
	uParentPos := make(map[int]int)
	for j, p := range uParents {
		// v is ancestor of u?
		if p == v {
			return uParents[:j], true
		}
		uParentPos[p] = j
	}
	vParents := parents(v)

	// is u an ancestor of v?
	for j, p := range vParents {
		if p == u {
			reverse(vParents[:j])
			return vParents[:j], false
		}
	}

	for j, p := range vParents {
		// First one of v's parents which is also one of u's parents
		if pos, ok := uParentPos[p]; ok {
			reverse(vParents[:j])
			return append(uParents[:pos+1], vParents[:j]...), false
		}
	}
	return nil, false
}

func reverse(s []int) []int {
	for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
		s[j], s[k] = s[k], s[j]
	}
	return s
}

func parents(n int) []int {
	var reply []int
	orig := n
	for n != 1 {
		reply = append(reply, nodes[n].parent)
		n = nodes[n].parent
	}
	fmt.Printf("parents(%d):%v\n", orig, reply)
	return reply
}

func pprint(n int, ind string) {
	fmt.Printf("%s%d %v\n", ind, n, nodes[n].depths)
	for _, c := range nodes[n].children {
		pprint(c, ind+"  ")
	}
}

// set my depthCounts to counts of how many nodes are below me.
func countDepths(n int) []int {
	if len(nodes[n].children) == 0 {
		return nil
	}
	var d []int
	for _, c := range nodes[n].children {
		combine(&d, countDepths(c))
	}
	nodes[n].depths = d
	return d
}

// add depth-counts from a child
func combine(d *[]int, ch []int) {
	for len(*d) < len(ch)+1 {
		*d = append(*d, 0)
	}
	(*d)[0]++
	for pos, v := range ch {
		(*d)[pos+1] += v
	}
}

func setParent(n, p int) {
	nodes[n].parent = p
	nodes[p].children = append(nodes[p].children, n)
}
