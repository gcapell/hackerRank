package main

import "fmt"

func main() {
	var queries int
	fmt.Scanln(&queries)
	for j := 0; j < queries; j++ {
		fmt.Println(query())
	}
}

func query() int {
	var nCities, nRoads, libCost, roadCost int
	fmt.Scanln(&nCities, &nRoads, &libCost, &roadCost)
	neighbours := make([][]int, nCities+1)
	for j := 0; j < nRoads; j++ {
		var a, b int
		fmt.Scanln(&a, &b)
		neighbours[a] = append(neighbours[a], b)
		neighbours[b] = append(neighbours[b], a)
	}
	if libCost <= roadCost {
		return nCities * libCost
	}

	visited := make(map[int]bool)

	reply := 0
	for j := 1; j <= nCities; j++ {
		if visited[j] {
			continue
		}
		reply += libCost
		r := bfs(neighbours, j, visited)
		reply += roadCost * r
	}
	return reply
}

func bfs(net [][]int, j int, visited map[int]bool) int {
	q := newQueue()
	q.push(j)
	reply := 0
	for !q.empty() {
		j := q.pop()
		visited[j] = true
		for _, neighbour := range net[j] {
			if visited[neighbour] {
				continue
			}
			q.push(neighbour)
			visited[neighbour] = true
			reply++
		}
	}
	return reply
}

type queue struct {
	buf        []int
	head, tail int
}

func newQueue() *queue {
	return &queue{buf: make([]int, 1000)}
}

func (q *queue) empty() bool {
	return q.head == q.tail
}

func (q *queue) push(n int) {
	q.buf[q.tail] = n
	q.tail = (q.tail + 1) % len(q.buf)
}

func (q *queue) pop() int {
	reply := q.buf[q.head]
	q.head = (q.head + 1) % len(q.buf)
	return reply
}
