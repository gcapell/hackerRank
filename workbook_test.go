package main

import "testing"

func TestA(t *testing.T) {
	data := []struct {
		page, problems, problemsPerPage int
	}{
		{2, 4, 2},
		{1, 3, 3},
		{5, 3, 3},
		{4, 6, 3},
		{5, 5, 3},
		{5, 6, 3},
		{5, 7, 3},
		{5, 8, 3},
		{5, 9, 3},
		{10, 9, 3},
		{29, 30, 29},
	}
	for _, d := range data {
		sd, nd := dumbCountSpecial(d.page, d.problems, d.problemsPerPage)
		s, n := countSpecial(d.page, d.problems, d.problemsPerPage)
		if s != sd || n != nd {
			t.Errorf("(%d,%d,%d) got %d,%d, want %d,%d",
				d.page, d.problems, d.problemsPerPage, s, n, sd, nd)
		}
	}
	page := 4
	problemsPerPage := 3
	for p := 1; p < 30; p++ {
		d, dnp := dumbCountSpecial(page, p, problemsPerPage)
		s, snp := countSpecial(page, p, problemsPerPage)
		if s != d {
			t.Errorf("(%d,%d,%d) got %d, want %d",
				page, p, problemsPerPage, s, d)
		}
		if dnp != snp {
			t.Errorf("nextpage %d/%d", dnp, snp)
		}
	}
}
