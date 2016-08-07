package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	distances := make(map[int]distance)
	var min int
	for j := 0; j < n; j++ {
		var a int
		fmt.Scanf("%d", &a)
		x := distances[a]
		d := x.update(j)
		distances[a] = x
		if d > 0 && (min == 0 || d < min) {
			min = d
		}
	}
	if min == 0 {
		min = -1
	}
	fmt.Println(min)
}

type distance struct {
	seen, minDistance, lastIndex int
}

func (d *distance) update(j int) int {
	switch d.seen {
	case 0:
		d.seen++
	case 1:
		d.seen++
		d.minDistance = j - d.lastIndex
	default:
		if j-d.lastIndex < d.minDistance {
			d.minDistance = j - d.lastIndex
		}
	}
	d.lastIndex = j
	return d.minDistance
}
