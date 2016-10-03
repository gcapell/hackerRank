package main

import "fmt"

type pos struct{ x, y int }

func main() {
	var size int
	fmt.Scanln(&size)
	grid := make(map[pos]bool)
	for j := 0; j < size; j++ {
		var line string
		fmt.Scanln(&line)
		for k, r := range line {
			if r == '.' {
				grid[pos{j, k}] = true
			}
		}
	}

	r := 0
	centres := grid
	for len(centres) > 0 {
		//fmt.Println(r, len(centres))
		r++
		centres = expand(centres, grid, r)
	}
	fmt.Println(r - 1)
}

func expand(centres, points map[pos]bool, r int) map[pos]bool {
	reply := make(map[pos]bool)
	qs := quadrantPerimeter(r)
	var radial [4]pos
outer:
	for c := range centres {
		for _, q := range qs {
			c.radial(q, radial[:])
			for _, p := range radial {
				if !points[p] {
					continue outer
				}
			}
		}
		reply[c] = true
	}
	return reply
}

func (p pos) radial(r pos, d []pos) {
	d[0] = pos{p.x + r.x, p.y + r.y}
	d[1] = pos{p.x + r.y, p.y - r.x}
	d[2] = pos{p.x - r.x, p.y - r.y}
	d[3] = pos{p.x - r.y, p.y + r.x}
}

var qpCache = make(map[int][]pos)

func quadrantPerimeter(r int) []pos {
	if reply, ok := qpCache[r]; ok {
		return reply
	}
	var reply []pos
	for j := 0; j < r; j++ {
		for k := 0; j*j+k*k <= r*r; k++ {
			if j*j+k*k > (r-1)*(r-1) {
				reply = append(reply, pos{j, k})
			}
		}
	}
	qpCache[r] = reply
	return reply
}
