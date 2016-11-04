package main

import "fmt"

func main() {
	var r, c, n int
	fmt.Scanf("%d %d %d", &r, &c, &n)
	g := readGrid(r, c)
	h := g.emptyCopy()
	for j := 0; j < n; j++ {
		g.print()
		g.boom(h)
		h, g = g, h
		fmt.Println()
	}
}

type grid [][]bool

func readGrid(r, c int) grid {
	g := make([][]bool, r)
	for j := 0; j < r; j++ {
		g[j] = make([]bool, c)
		var row string
		fmt.Scanf("%s", &row)
		for k, ch := range row {
			if ch == 'O' {
				g[j][k] = true
			}
		}
	}
	return g
}

var chars = map[bool]string{
	false: ".",
	true:  "O",
}

func (g grid) print() {
	for _, r := range g {
		for _, c := range r {
			fmt.Print(chars[c])
		}
		fmt.Println()
	}
}

func (g grid) emptyCopy() grid {
	h := make([][]bool, len(g))
	for j := 0; j < len(h); j++ {
		h[j] = make([]bool, len(g[j]))
	}
	return h
}

func (g grid) boom(h grid) {
	for j := 0; j < len(g); j++ {
		for k := 0; k < len(g[j]); k++ {
			h[j][k] = !g.bombNear(j, k)
		}
	}
}

func (g grid) bombNear(j, k int) bool {
	return g[j][k] ||
		(j > 0 && g[j-1][k]) ||
		(j+1 < len(g) && g[j+1][k]) ||
		(k > 0 && g[j][k-1]) ||
		(k+1 < len(g[j]) && g[j][k+1])
}
