package main

import "fmt"

type point struct {
	r, c int
}

const (
	north int = iota
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
)

type queen struct {
	point
	obstructions [8]point // 0=North, continue clockwise
}

func main() {
	var length, obstacles int
	fmt.Scanln(&length, &obstacles)
	var q queen
	fmt.Scanln(&q.r, &q.c)
	q.setMaxObstructions(length)

	for j := 0; j < obstacles; j++ {
		var o point
		fmt.Scanln(&o.r, &o.c)
		q.obstruct(o)
	}
	fmt.Println(q.canAttack())
}

func (q *queen) setMaxObstructions(length int) {
	directions := map[int]struct{ dr, dc int }{
		north:     {1, 0},
		south:     {-1, 0},
		east:      {0, 1},
		west:      {0, -1},
		northEast: {1, 1},
		northWest: {1, -1},
		southEast: {-1, 1},
		southWest: {-1, -1},
	}
	for d, delta := range directions {
		p := q.point
		for ok(p.r, length) && ok(p.c, length) {
			p.r += delta.dr
			p.c += delta.dc
		}
		q.obstructions[d] = p
	}
}

// sentinel obstructions
func (q *queen) setMaxObstructionsClever(length int) {
	q.obstructions[north] = point{length + 1, q.c}
	q.obstructions[south] = point{0, q.c}
	q.obstructions[east] = point{q.r, length + 1}
	q.obstructions[west] = point{q.r, 0}
	if q.c >= q.r {
		q.obstructions[northEast] = point{q.r + length + 1 - q.c, length + 1}
		q.obstructions[southWest] = point{0, q.c - q.r}
	} else {
		q.obstructions[northEast] = point{length + 1, q.c + length + 1 - q.r}
		q.obstructions[southWest] = point{q.r - q.c, 0}
	}
	if q.c+q.r > length {
		q.obstructions[northWest] = point{length + 1, q.c - (length + 1 - q.r)}
		q.obstructions[southEast] = point{q.r - (length + 1 - q.c), length + 1}
	} else {
		q.obstructions[northWest] = point{q.r + q.c, 0}
		q.obstructions[southEast] = point{0, q.c + q.r}
	}
	//fmt.Println(q.obstructions)
}

func (q *queen) obstruct(p point) {
	var dir int
	switch {
	case q.r == p.r:
		if p.c > q.c {
			dir = east
		} else {
			dir = west
		}
	case q.c == p.c:
		if p.r < q.r {
			dir = south
		} else {
			dir = east
		}
	case delta(q.r, p.r) == delta(q.c, p.c):
		if p.c > q.c {
			if p.r > q.r {
				dir = northEast
			} else {
				dir = southEast
			}
		} else {
			if p.r > q.r {
				dir = northWest
			} else {
				dir = southWest
			}
		}
	default:
		return
	}
	if q.distance(p) < q.distance(q.obstructions[dir]) {
		q.obstructions[dir] = p
	}
}

func ok(n, length int) bool {
	return n > 0 && n <= length
}

func (p point) distance(o point) int {
	if p.r != o.r {
		return delta(p.r, o.r)
	}
	return delta(p.c, o.c)
}

func delta(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (q *queen) canAttack() int {
	var squares int
	for _, o := range q.obstructions {
		squares += q.distance(o) - 1
	}
	return squares
}
