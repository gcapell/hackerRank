package main

import "fmt"

func main() {
	var m, n, r int

	fmt.Scanln(&m, &n, &r)
	matrix := readMatrix(m, n)
	rotated := rotate(matrix, r)
	printMatrix(rotated)
}

func readMatrix(rows, cols int) [][]int {
	var m [][]int
	for j := 0; j < rows; j++ {
		row := make([]int, cols)
		for k := 0; k < cols; k++ {
			fmt.Scan(&row[k])
		}
		m = append(m, row)
	}
	return m
}

func printMatrix(m [][]int) {
	for _, row := range m {
		for _, n := range row {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
}

type point struct{ row, col int }

func (p point) lessThan(o point) bool {
	return p.row < o.row && p.col < o.col
}

func (p *point) inc() {
	p.row++
	p.col++
}

func (p *point) dec() {
	p.row--
	p.col--
}

func rotate(m [][]int, r int) [][]int {
	rows, cols := len(m), len(m[0])
	dst := empty(rows, cols)
	tl := point{0, 0}
	br := point{rows - 1, cols - 1}
	for tl.lessThan(br) {
		rotateLayer(tl, br, m, dst, r)
		tl.inc()
		br.dec()
	}
	return dst
}

func empty(rows, cols int) [][]int {
	reply := make([][]int, rows)
	for r := 0; r < rows; r++ {
		reply[r] = make([]int, cols)
	}
	return reply
}

func rotateLayer(tl, br point, src, dst [][]int, r int) {
	s := newCursor(tl, br)
	d := newCursor(tl, br)
	d.advance(r % ((br.row - tl.row + br.col - tl.col) * 2))
	for {
		dst[d.row][d.col] = src[s.row][s.col]
		d.advance(1)
		s.advance(1)
		if s.point == tl {
			return
		}
	}
}

type direction int

const (
	down direction = iota
	right
	up
	left
)

type cursor struct {
	point
	tl, br point
	d      direction //
}

func newCursor(tl, br point) cursor {
	return cursor{point: tl, tl: tl, br: br}
}

func (c *cursor) advance(n int) {
	switch c.d {
	case down:
		dist := min(n, c.br.row-c.row)
		c.row += dist
		n -= dist
		if c.row == c.br.row {
			c.d = right
		}
	case right:
		dist := min(n, c.br.col-c.col)
		c.col += dist
		n -= dist
		if c.col == c.br.col {
			c.d = up
		}
	case up:
		dist := min(n, c.row-c.tl.row)
		c.row -= dist
		n -= dist
		if c.row == c.tl.row {
			c.d = left
		}
	case left:
		dist := min(n, c.col-c.tl.col)
		c.col -= dist
		n -= dist
		if c.col == c.tl.col {
			c.d = down
		}
	}
	if n > 0 {
		c.advance(n)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
