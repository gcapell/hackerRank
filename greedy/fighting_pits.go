package main

import (
	"fmt"
	"sort"
)

var (
	teams [][]int
	cache [][]*cacheEntry
)

func main() {
	var n, k, q int
	fmt.Scanln(&n, &k, &q)
	teams = make([][]int, k+1)
	cache = make([][]cacheEntry, k+1)
	for j := 1; j <= k; j++ {
		cache[j] = make([]cacheEntry, k+1)
		for e := 1; e <= k; e++ {
			cache[j][e] = newCacheEntry()
		}
	}
	for j := 0; j < n; j++ {
		var s, t int
		fmt.Scanln(&s, &t)
		teams[t] = append(teams[t], s)
	}
	for j := 1; j <= k; j++ {
		sort.Ints(teams[j])
	}
	for j := 0; j < q; j++ {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		switch a {
		case 1:
			teams[c] = append(teams[c], b)
		case 2:
			if win(b, c, cache[b][c], cache[c][b]) {
				fmt.Println(b)
			} else {
				fmt.Println(c)
			}
		}
	}
}

func win(teamA, teamB []int, cacheA, cacheB *cacheEntry) bool {
	if result, ok := cacheA.get(len(teamA), len(teamB)); ok {
		return result
	}
	var reply bool
	s := teamA[len(teamA)-1]
	if len(teamB) <= s {
		reply = true
	} else {
		reply = !win(teamB[:len(teamB)-s], teamA, cacheB, cacheA)
	}
	cache.put(len(teamA), len(teamB), reply)
	return reply
}

type (
	// given one count, what's the other count which guarantees a win or loss?
	winLoss struct {
		win, loss int
	}

	// cache maps from our counts to winLoss records and from their counts to winLoss records
	cacheEntry struct {
		our   map[int]winLose
		their map[int]winLose
	}
)

func newCacheEntry() *cacheEntry {
	return &cacheEntry{
		make(map[int]winLoss),
		make(map[int]winLoss),
	}
}

func (c *cacheEntry) get(ourLen, theirLen int) (bool, bool) {
	if e, ok := c.our[ourLen]; ok {
		if e.win > 0 && theirLen <= e.win {
			return true, true
		}
		if e.loss > 0 && theirLen >= e.loss {
			return false, true
		}
	}
	if e, ok := c.their[theirLen]; ok {
		if e.win > 0 && ourLen >= e.win {
			return true, true
		}
		if e.loss > 0 && ourLen <= e.loss {
			return false, true
		}
	}
	return false, false
}

func (c *cacheEntry) put(ourLen, theirLen int, win bool) {
	e := c.our[ourLen]
	if win {

	}

}
