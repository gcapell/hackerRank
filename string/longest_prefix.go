package prefix

// LongestSuffixAndPrefix calculates, for every prefix P of a string S,
// what's the longest proper suffix of P which is also a proper prefix of P.
// ('proper' meaning not counting 'ab' as a prefix or suffix of 'ab').
//
// For example, where P is "aabaab", "aab" is the longest suffix
// of P which is also a prefix of P.
//
// We represent this with an array A, where A[n] is the position
// in S of the end of the prefix of S matching the longest suffix
// ending at S[n]
//
// For example:
// P   a  a  b  a  a  b
// A  -1  0 -1  0  1  2
//
// A[0] will always be -1 (no proper suffix)
// A[1]=0 means the prefix ending at 0 (i.e. "a") matches the suffix ending at 1 "a"
// A[2]=-1 (neither 'ab' nor 'b' are prefix of S)
// A[3]=0 suffix ending at 3 = prefix ending at 0='a'
// A[4]=1 suffix ... 4 = prefix ... 1 = 'aa'
// A[5]=2 suffix ... 5 = prefix ... 2 = 'aab'

func LongestSuffixAndPrefix(s string) []int {
	if len(s) == 0 {
		return nil
	}
	a := make([]int, len(s))
	a[0] = -1
outer:
	for j := 1; j < len(s); j++ {
		prev := j - 1
		for a[prev] != -1 {
			if s[j] == s[a[prev]+1] {
				a[j] = a[prev] + 1
				continue outer
			}
			// Can we extend a _shorter_ suffix?
			// The prefix that a[prev] points to could also have a proper
			// suffix matching a (smaller) prefix.
			// e.g. having seen "abacaba" we now see "b"
			// Although we're working on the suffix/prefix "aba" and
			// initially hoping to see a "c" to extend to "abac",
			// "aba" itself has suffix/prefix "a" which can be extended with the "b"
			prev = a[prev]
		}
		if s[j] == s[0] {
			a[j] = 0
		} else {
			a[j] = -1
		}
	}
	return a
}

// Period returns the period of a cyclic string (len(s) if not cyclic).
// E.g. Period("abcabc") -> 3, Period("aaa") -> 1, Period("abc") -> 3
// Period ("abca") -> 3 Period("") = 1, Period("a") = 1
func Period(s string) int {
	a := LongestSuffixAndPrefix(s)
	period := 1
	for j := 1; j < len(s); j++ {
		if s[j] != s[j%period] {
			period = j - a[j]
		}
	}
	return period
}
