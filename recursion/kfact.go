package main

import "fmt"

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19}
var facts = [][]int{
	{},                 //0
	{},                 //1
	{1},                //2
	{0, 1},             //3
	{2},                //4
	{0, 0, 1},          //5
	{1, 1},             //6
	{0, 0, 0, 1},       //7
	{3},                //8
	{0, 2},             //9
	{1, 0, 1},          //10
	{0, 0, 0, 0, 1},    //11
	{2, 1},             //12
	{0, 0, 0, 0, 0, 1}, //13
	{1, 0, 0, 1},       //14
	{0, 1, 1},          //15
	{4},                //16
	{0, 0, 0, 0, 0, 0, 1}, //17
	{1, 2},                //18
	{0, 0, 0, 0, 0, 0, 0, 1}, //19
	{2, 0, 1},                //20
}

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	a := make([]int, k)
	for j := 0; j < k; j++ {
		fmt.Scan(&a[j])
	}

	// factorise n
	var factored []int
	for _, p := range primes {
		if n == 1 {
			break
		}
		count := 0
		for n%p == 0 {
			count++
			n /= p
		}
		factored = append(factored, count)
	}
	if n != 1 {
		fmt.Println(-1)
		return
	}

	fmt.Println(factored)
	for _, f := range a {
		fmt.Println(facts[f])
	}
}
