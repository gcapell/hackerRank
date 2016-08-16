package main

import "fmt"

var primes = sieve(1.1e5)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n+1)
	for j := 0; j < n; j++ {
		fmt.Scanf("%d", &a[j+1])
	}
	fmt.Println(sequences(a))
}

func sieve(max int) []int {
	p := []int{2}
	composite := make([]bool, max)
	for j := 3; j < max; j += 2 {
		if composite[j] {
			continue
		}
		p = append(p, j)
		for k := j + j; k < max; k += j {
			composite[k] = true
		}
	}
	return p
}

type primeCount struct {
	maxPowerUnspecified int
	maxPowerSeen        int
}

var primeData = map[int]primeCount{}

func factors(n int, specified bool) {
	for j := 0; primes[j] <= n; j++ {
		p := primes[j]
		count := 0
		for n%p == 0 {
			n /= p
			count++
		}
		if count > 0 {
			primeData[p] = primeData[p].update(count, specified)
		}
	}
}

func (p primeCount) update(count int, specified bool) primeCount {
	if specified && count > p.maxPowerUnspecified {
		p.maxPowerUnspecified = count
	}
	if count > p.maxPowerSeen {
		p.maxPowerSeen = count
	}
	return p
}

func sequences(a []int) uint64 {
	reply := uint64(1)
	for j := 2; j < len(a); j++ {
		factors(j, a[j] != -1)
	}

	for p, d := range primeData {
		reply = reply * pow(p, d.maxPowerSeen-d.maxPowerUnspecified)
		reply %= 1e9 + 7
	}
	return reply
}

func pow(p, n int) uint64 {
	if n == 0 {
		return 1
	}
	reply := uint64(1)
	for j := 0; j < n; j++ {
		reply *= uint64(p)
	}
	return reply
}
