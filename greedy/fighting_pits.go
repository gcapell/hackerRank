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
	cpuprofile      = flag.String("cpuprofile", "", "write cpu profile to file")
	expected        = flag.String("expected", "", "file of expected answers")
	expectedAnswers []int
)

type (
	soldier struct {
		army, strength int
	}
	squad struct {
		strength, size int
	}
	army []squad // in increasing strength

	cursor struct {
		a        army
		squad    int // within the team
		size     int // remaining in the squad
		strength int // of current squad
		live     bool
	}
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
	if *expected != "" {
		buf, err := xx.Readfile(*expected)
		if err != nil {
			log.Fatal(err)
		}
		expectedAnswers = mapInt(strings.Fields(string(buf)))
	}
	f := bufio.NewReaderSize(os.Stdin, 4e6)
	var n, k, q int
	fmt.Fscanln(f, &n, &k, &q)
	soldiers := make([]soldier, n)
	for j := 0; j < n; j++ {
		fmt.Fscanln(f, &soldiers[j].strength, &soldiers[j].army)
	}
	sort.Sort(byArmyThenStrength(soldiers))

	armies := make([]army, k+1)

	for _, s := range soldiers {
		armies[s.army] = armies[s.army].add(s.strength)
	}

	for j := 0; j < q; j++ {
		var a, b, c int
		fmt.Fscanln(f, &a, &b, &c)
		switch a {
		case 1:
			armies[c] = armies[c].add(b)
		case 2:
			if win(armies[b], armies[c]) {
				output(b)
			} else {
				output(c)
			}
		}
	}
}

func output(n int) {
	if len(expectedAnswers) > 0 {
		if n != expectedAnswers[0] {
			log.Fatalf("unexpected output, got %d, want %d", n, expectedAnswers[0])
		}
		expectedAnswers = expectedAnswers[1:]
	}
	fmt.Println(n)
}

func (a army) add(s int) army {
	if len(a) > 0 && a[len(a)-1].strength == s {
		a[len(a)-1].size++
		return a
	}
	if len(a) > 0 {
		assert(a[len(a)-1].strength < s)
	}
	return append(a, squad{strength: s, size: 1})
}

func win(aa, ab army) bool {
	a, b := aa.cursor(), ab.cursor()
	for a.live && b.live {
		aBefore, bBefore := a.squad, b.squad
		rounds := min(a.size/b.strength, (b.size-1)/a.strength)
		b.die(rounds * a.strength)
		a.die(rounds * b.strength)
		assert(b.live)
		if !a.live {
			break
		}
		b.die(a.strength)
		if !b.live {
			break
		}
		a.die(b.strength)

		if !(a.squad < aBefore || b.squad < bBefore) {
			log.Panic("progress?", a.squad, aBefore, b.squad, bBefore, a, b)
		}
	}
	return a.live
}

func assert(b bool) {
	if !b {
		log.Panic("assert")
	}
}

func (c *cursor) die(n int) {
	for n >= c.size {
		n -= c.size
		c.squad--
		if c.squad < 0 {
			c.live = false
			return
		}
		s := c.a[c.squad]
		c.size, c.strength = s.size, s.strength
	}
	assert(n < c.size)
	c.size -= n
}

func (a army) cursor() cursor {
	if len(a) < 1 {
		return cursor{live: false}
	}
	s := len(a) - 1

	return cursor{
		a:        a,
		squad:    s,
		size:     a[s].size,
		strength: a[s].strength,
		live:     s >= 0,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type byArmyThenStrength []soldier

func (a byArmyThenStrength) Len() int      { return len(a) }
func (a byArmyThenStrength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byArmyThenStrength) Less(i, j int) bool {
	return a[i].army < a[j].army ||
		(a[i].army == a[j].army && a[i].strength < a[j].strength)
}
