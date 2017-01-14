package prefix

import (
	"fmt"
	"math/rand"
	"testing"
)

// TestLongest checks LongestSuffixAndPrefix against known results
func TestLongest(t *testing.T) {
	testData := []struct {
		s string
		a []int
	}{
		{"a", []int{-1}},
		{"aa", []int{-1, 0}},
		{"aab", []int{-1, 0, -1}},
		{"aabaaaba", []int{-1, 0, -1, 0, 1, 1, 2, 3}},
		{"abacabab", []int{-1, -1, 0, -1, 0, 1, 2, 1}},
	}
	for _, d := range testData {
		a := LongestSuffixAndPrefix(d.s)
		if !same(a, d.a) {
			t.Errorf("LongestSuffixAndPrefix(%s) got %v, want %v",
				d.s, a, d.a)
		}
	}
}

// TestDumb checks LongestSuffixAndPrefix against dumb implementation with handcrafted data.
func TestDumb(t *testing.T) {
	testData := []string{
		"a", "aa", "aab", "aabaaaba", "abacabab",
	}
	for _, s := range testData {
		a := LongestSuffixAndPrefix(s)
		dumbA := dumbLongestSuffixAndPrefix(s)
		if !same(a, dumbA) {
			t.Errorf("failure(%s): smart->%v, dumb->%v", s, a, dumbA)
		}
	}
}

// TestRandomAgainstDumb checks LongestSuffixAndPrefix against dumb implementation with random data.
func TestRandomAgainstDumb(t *testing.T) {
	for j := 0; j < 1000; j++ {
		s := randomString(20, "ab")
		a := LongestSuffixAndPrefix(s)
		dumbA := dumbLongestSuffixAndPrefix(s)
		if !same(a, dumbA) {
			t.Errorf("failure(%s): smart->%v, dumb->%v", s, a, dumbA)
		}
	}
}

func TestKnownPeriod(t *testing.T) {
	testData := []struct {
		s      string
		period int
	}{
		{"a", 1},
		{"aa", 1},
		{"aab", 3},
		{"aabaaaba", 4},
		{"abacabab", 6},
	}
	for _, d := range testData {
		period := Period(d.s)
		if period != d.period {
			t.Errorf("Period(%s) got %d want %d", d.s, period, d.period)
		}
	}
}

func TestDumbPeriodKnownValues(t *testing.T) {
	testData := []string{"a", "aa", "aab", "aabaaaba", "abcabca", "abacabab"}
	for _, s := range testData {
		p := Period(s)
		dp := dumbPeriod(s)
		if p != dp {
			t.Errorf("Period(%s) smart:%d, dumb:%d", s, p, dp)
		}
	}
}

func TestPeriodRandom(t *testing.T) {
	testParams := []struct {
		reps    int
		length  int
		choices string
	}{
		{1000, 20, "ab"},
		{1000, 20, "abc"},
		{100, 100, "abc"},
	}
	var acyclic, cyclic int
	for _, param := range testParams {
		for j := 0; j < param.reps; j++ {
			s := randomString(param.length, param.choices)
			p := Period(s)
			dp := dumbPeriod(s)
			if p == len(s) {
				acyclic++
			} else {
				cyclic++
			}
			if p != dp {
				t.Errorf("failure(%s): smart->%d, dumb->%d", s, p, dp)
			}
		}
	}
	fmt.Println("acyclic", acyclic, " cyclic:", cyclic)
}

func dumbPeriod(s string) int {
period:
	for p := 1; p < len(s); p++ {
		for j := p; j < len(s); j++ {
			if s[j] != s[j%p] {
				continue period
			}
		}
		return p
	}
	return len(s)
}

func dumbLongestSuffixAndPrefix(s string) []int {
	a := make([]int, len(s))
	a[0] = -1
outer:
	for j := 1; j < len(s); j++ {
		for l := j; l > 0; l-- {
			prefix := s[:l]
			suffix := s[j-l+1 : j+1]
			if prefix == suffix {
				a[j] = l - 1
				continue outer
			}
		}
		a[j] = -1

	}
	return a
}

func randomString(length int, choices string) string {
	buf := make([]byte, length)
	for j := 0; j < length; j++ {
		buf[j] = choices[rand.Int31n(int32(len(choices)))]
	}
	return string(buf[:])
}

func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for pos, v := range a {
		if v != b[pos] {
			return false
		}
	}
	return true
}
