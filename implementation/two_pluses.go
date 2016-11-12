package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	goodCells := readGoodCells(m, n)
	//fmt.Println(goodCells)
	pluses := findPluses(goodCells)
	//fmt.Println(pluses)
	maxSize := largestKey(pluses)
	sortedSizePairs := sortSizePairs(maxSize)
	fmt.Println(sortedSizePairs)
	var area int
	for _, p := range sortedSizePairs {
		if findOneDisjoint(p, pluses, maxSize) {
			area = p.areaProduct()
			break
		}
	}
	fmt.Println(area)
}

func findOneDisjoint(p sizePair, pluses map[int][]point, maxSize int) bool {
	for _, c := range plusesOfSize(p.larger, pluses, maxSize) {
		for _, c2 := range plusesOfSize(p.smaller, pluses, maxSize) {
			if disjoint(c.delta(c2), p) {
				//fmt.Println("found", c, c2, p)
				return true
			}
		}
	}
	return false
}

func (a point) delta(b point) sizePair {
	dx, dy := absDelta(a.r, b.r), absDelta(a.c, b.c)
	if dy > dx {
		dy, dx = dx, dy
	}
	return sizePair{dx, dy}
}

func absDelta(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func disjoint(pos, size sizePair) bool {
	//fmt.Printf("disjoint pos:%v size:%v\n", pos, size)
	if pos.smaller == 0 {
		return size.larger+size.smaller < pos.larger
	}

	return pos.smaller > size.smaller || pos.larger > size.larger
}

func plusesOfSize(n int, pluses map[int][]point, maxSize int) []point {
	reply := pluses[n]
	for j := n + 1; j <= maxSize; j++ {
		reply = append(reply, pluses[j]...)
	}
	return reply
}

func sortSizePairs(maxSize int) []sizePair {
	var reply []sizePair
	for j := 0; j <= maxSize; j++ {
		for k := j; k <= maxSize; k++ {
			reply = append(reply, sizePair{k, j})
		}
	}
	sort.Sort(byAreaProduct(reply))
	return reply
}

func largestKey(pluses map[int][]point) int {
	max := 0
	for k := range pluses {
		if k > max {
			max = k
		}
	}
	return max
}

type point struct{ r, c int }
type sizePair struct{ larger, smaller int }

type byAreaProduct []sizePair

func (a byAreaProduct) Len() int      { return len(a) }
func (a byAreaProduct) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byAreaProduct) Less(i, j int) bool {
	return a[i].areaProduct() > a[j].areaProduct()
}

func (s sizePair) areaProduct() int {
	return (s.larger*4 + 1) * (s.smaller*4 + 1)
}

func readGoodCells(m, n int) map[point]bool {
	reply := make(map[point]bool)
	for row := 0; row < m; row++ {
		var line string
		fmt.Scanf("%s", &line)
		for col, ch := range line {
			if ch == 'G' {
				reply[point{row, col}] = true
			}
		}
	}
	return reply
}

func findPluses(cells map[point]bool) map[int][]point {
	reply := make(map[int][]point)
	for p := range cells {
		size := 0
		for {
			n := size + 1
			if !(cells[point{p.r, p.c + n}] &&
				cells[point{p.r, p.c - n}] &&
				cells[point{p.r + n, p.c}] &&
				cells[point{p.r - n, p.c}]) {
				break
			}
			size = n
		}
		reply[size] = append(reply[size], p)
	}
	return reply
}
