package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var score int
	fmt.Scanf("%d", &score)
	best, worst := score, score
	n--
	var bestCount, worstCount int
	for j := 0; j < n; j++ {
		fmt.Scanf("%d", &score)
		if score > best {
			bestCount++
			best = score
		}
		if score < worst {
			worstCount++
			worst = score
		}
	}

	fmt.Println(bestCount, worstCount)
}
