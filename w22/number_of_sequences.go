package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n+1)
	for j := range a {
		fmt.Scanf("%d", &a[j+1])
	}
	fmt.Println(sequences(a))
}

func sequences(a []int) uint64 {
	if a[1] != 0 {
		return 0
	}
	reply := uint64(1)
	composite := make([]bool, n)
	for j := 2; j < len(a); j++ {
		if composite[j] {
			continue
		}
		ok, mul := sieve(a, composite, j)
		if !ok {
			return 0
		}
		reply = (reply * mul) % (1e9 + 7)
	}
	return reply
}

func sieve(a []int, composite []int, p int) (bool, uint64) {
	if a[p] >= p {
		return false, 0
	}
	maxPower := p
	nextPower := p * p
	specified := false
	var mod, rem int

	for j := p; j < len(a); j += p {
		composite[j] = true
		if j == nextPower {
			maxPower, nextPower = nextPower, nextPower*p
		}
		if a[j] == -1 {
			continue
		}
		specified := true
		mod = maxPower
		rem = a[j] % mod
	}
	if !specified {
		return true, maxPower
	}

}
