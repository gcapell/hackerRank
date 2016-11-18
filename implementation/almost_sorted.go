package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	ds := make([]int, n+2)
	for j := 0; j < n; j++ {
		fmt.Scan(&ds[j+1])
	}
	ds[0] = -1
	ds[n+1] = 1e6 + 1
	fmt.Println(fix(ds))
}

func fix(ds []int) string {
	for j := 1; j < len(ds); j++ {
		if ds[j] < ds[j-1] {
			return swapOrReverse(ds, j-1)
		}
	}
	return "yes"
}

// ds[first+1] is < ds[first].  We have to swap or reverse.
func swapOrReverse(ds []int, first int) string {
	if ds[first+2] < ds[first+1] {
		return reverse(ds, first)
	}
	if ds[first+2] > ds[first] {
		if ds[first+1] > ds[first-1] && ascending(ds[first+2:]) {
			return fmt.Sprintf("yes\nswap %d %d", first, first+1)
		} else {
			return "no"
		}
	}
	// swap or bust
	for j := first + 2; j < len(ds); j++ {
		if ds[j] > ds[first] {
			return "no"
		}
		if ds[j] < ds[j-1] {
			if !(ds[j] > ds[first-1] && ds[j] < ds[first+1]) {
				return "no"
			}
			if ds[j+1] < ds[first] {
				return "no"
			}
			if ascending(ds[j+1:]) {
				return fmt.Sprintf("yes\nswap %d %d", first, j)
			}
		}
	}
	return "NEVER HAPPEN swapOrReverse"
}

// ds[first] > ds[first+1] > ds[first+2]
func reverse(ds []int, first int) string {
	for j := first + 3; j < len(ds); j++ {
		if ds[j] > ds[j-1] {
			if ds[first] > ds[j] || ds[j-1] < ds[first-1] || !ascending(ds[j:]) {
				return "no"
			}
			return fmt.Sprintf("yes\nreverse %d %d", first, j-1)
		}
	}
	return "NEVER HAPPEN reverse"
}

func ascending(ds []int) bool {
	for j := 1; j < len(ds); j++ {
		if ds[j] < ds[j-1] {
			return false
		}
	}
	return true
}
