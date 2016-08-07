package main

import "fmt"

func main() {
	var chapters, problemsPerPage int
	fmt.Scanf("%d %d", &chapters, &problemsPerPage)
	if problemsPerPage == 1 {
		// Every problem in chapter 1 is special, and that's all.
		var problems int
		fmt.Scanf("%d", &problems)
		fmt.Println(problems)
		return
	}

	page, special := 1, 0
	for j := 0; j < chapters; j++ {
		var problems int
		fmt.Scanf("%d", &problems)
		nextPage := page + divCeil(problems, problemsPerPage)
		if page == 1 {
			special++
		} else if problems >= nextPage-1 {
			special++
			catchupPerPage := problemsPerPage - 1
			delta := page - problemsPerPage
			if delta >= 0 &&
				delta%catchupPerPage == 0 &&
				problems > problemsPerPage*(1+(delta/catchupPerPage)) {
				// The first 'special' problem is at the end of a page.
				// There's a subsequent problem, also special.
				special++
			}
		}
		page = nextPage
	}
	fmt.Println(special)
}

func divCeil(a, b int) int {
	r := a / b
	if a%b != 0 {
		r++
	}
	return r
}
