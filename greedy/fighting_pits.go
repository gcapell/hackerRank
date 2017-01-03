package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	teams [][]int
	cache map[contest]*cacheEntry
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	f := bufio.NewReaderSize(os.Stdin, 4e6)
	var n, k, q int
	fmt.Fscanln(f, &n, &k, &q)
	teams = make([][]int, k+1)
	cache = make(map[contest]*cacheEntry)
	for j := 0; j < n; j++ {
		var s, t int
		fmt.Fscanln(f, &s, &t)
		teams[t] = append(teams[t], s)
	}
	for j := 1; j <= k; j++ {
		sort.Ints(teams[j])
	}
	for j := 0; j < q; j++ {
		var a, b, c int
		fmt.Fscanln(f, &a, &b, &c)
		switch a {
		case 1:
			teams[c] = append(teams[c], b)
		case 2:
			if win(teams[b], teams[c], getCache(b, c), getCache(c, b)) {
				fmt.Println(b)
			} else {
				fmt.Println(c)
			}
		}
	}
}

func getCache(first, second int) *cacheEntry {
	k := contest{first, second}
	e, ok := cache[k]
	if !ok {
		e = &cacheEntry{}
		cache[k] = e
	}
	return e
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
	cacheA.put(len(teamA), len(teamB), reply)
	return reply
}

type (
	contest struct {
		first, second int
	}

	// given one count, what's the other count which guarantees a win or loss?
	winLoss struct {
		win, loss int
	}

	// cache maps from our counts to winLoss records and from their counts to winLoss records
	cacheEntry struct {
		our   []winLoss
		their []winLoss
	}
)

func (c *cacheEntry) get(ourLen, theirLen int) (bool, bool) {
	if len(c.our) > ourLen {
		e := c.our[ourLen]
		if e.win > 0 && theirLen <= e.win {
			return true, true
		}
		if e.loss > 0 && theirLen >= e.loss {
			return false, true
		}
	}
	if len(c.their) > theirLen {
		e := c.their[theirLen]
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
	if len(c.our) <= ourLen {
		c.our = append(c.our, make([]winLoss, ourLen-len(c.our)+1)...)
	}
	e := c.our[ourLen]
	if win {
		if theirLen > e.win {
			e.win = theirLen
		}
	} else {
		if e.loss == 0 || theirLen < e.loss {
			e.loss = theirLen
		}
	}
	c.our[ourLen] = e

	if len(c.their) <= theirLen {
		c.their = append(c.their, make([]winLoss, theirLen-len(c.their)+1)...)
	}
	e = c.their[theirLen]
	if win {
		if e.win == 0 || ourLen < e.win {
			e.win = ourLen
		}
	} else {
		if ourLen > e.loss {
			e.loss = ourLen
		}
	}
	c.their[theirLen] = e
}
