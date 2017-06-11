package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	var qr, qc int
	fmt.Scanln(&qr, &qc)
	b := newBoard(n, qr, qc)
	for j := 0; j < k; j++ {
		var r, c int
		fmt.Scanln(&r, &c)
		b.add(r, c)
	}
	fmt.Println(b.totalAttacked())
}

type direction int

const (
	// rows increase to the north,
	// columns increase to the east.
	north direction = iota
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
)

var directionNames = []string{
	"north", "northEast",
	"east", "southEast",
	"south", "southWest",
	"west", "northWest"}

func (d direction) String() string {
	return directionNames[d]
}

type board struct {
	r, c int // queen pos
	// How many squares are attacked, by direction
	attacked [8]int
}

func newBoard(n, r, c int) *board {
	b := &board{r: r, c: c}
	b.attacked[north] = n - r
	b.attacked[northEast] = min(n-r, n-c)
	b.attacked[east] = n - c
	b.attacked[southEast] = min(r-1, n-c)
	b.attacked[south] = r - 1
	b.attacked[southWest] = min(r-1, c-1)
	b.attacked[west] = c - 1
	b.attacked[northWest] = min(n-r, c-1)
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (b *board) add(r, c int) {
	if dir, n, ok := isAttacked(r-b.r, c-b.c); ok && n < b.attacked[dir] {
		b.attacked[dir] = n
	}
}

func (b *board) totalAttacked() int {
	total := 0
	for _, n := range b.attacked {
		total += n
	}
	return total
}

func isAttacked(dr, dc int) (d direction, m int, ok bool) {
	defer func() {
		m--
	}()
	if dr == 0 && dc == 0 {
		panic("sunk my battleship")
	}
	switch {
	case dr == 0:
		switch {
		case dc < 0:
			return west, -dc, true
		default:
			return east, dc, true
		}
	case dc == 0:
		switch {
		case dr < 0:
			return south, -dr, true
		default:
			return north, dr, true
		}
	case dr < 0:
		switch {
		case dc == dr:
			return southWest, -dr, true
		case dc == -dr:
			return southEast, -dr, true
		}
	default:
		switch {
		case dc == dr:
			return northEast, dr, true
		case dc == -dr:
			return northWest, dr, true
		}
	}
	return 0, 0, false
}
