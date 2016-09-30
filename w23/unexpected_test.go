package main

import (
	"reflect"
	"testing"
)

func TestHasPeriod(t *testing.T) {
	testData := []struct {
		s    string
		p    int
		want bool
	}{
		{"aaa", 1, true},
		{"aaa", 2, false},
		{"aba", 2, false},
		{"abab", 2, true},
		{"abab", 3, false},
	}
	for _, d := range testData {
		got := hasPeriod(d.s, d.p)
		if got != d.want {
			t.Errorf("hasPeriod(%s,%d) got %v, want %v",
				d.s, d.p, got, d.want)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	testData := []struct {
		n    int
		want []exponent
	}{
		{3, []exponent{{3, 1}}},
		{4, []exponent{{2, 2}}},
		{12, []exponent{{2, 2}, {3, 1}}},
		{5e9, []exponent{{2, 9}, {5, 10}}},
		{5e5 - 1, []exponent{{31, 1}, {127, 2}}},
	}
	for _, d := range testData {
		got := primeFactors(d.n)
		if !reflect.DeepEqual(got, d.want) {
			t.Errorf("primeFactors(%d) = %v, want %v",
				d.n, got, d.want)
		}
	}
}

func TestSortFactors(t *testing.T) {
	testData := []struct {
		primes  []exponent
		factors []int
	}{
		{
			[]exponent{{2, 3}, {3, 1}},
			[]int{1, 2, 3, 4, 6, 8, 12, 24},
		},
		{
			[]exponent{{17, 1}},
			[]int{1, 17},
		},
		{
			[]exponent{{31, 1}, {127, 2}},
			[]int{1, 31, 127, 3937, 16129, 499999},
		},
	}
	for _, d := range testData {
		got := sortFactors(d.primes)
		if !reflect.DeepEqual(got, d.factors) {
			t.Errorf("combinePrimes(%v) = %v, want %v",
				d.primes, got, d.factors)
		}
	}
}

func TestPeriod(t *testing.T) {
	testData := []struct {
		s string
		p int
	}{
		{"aaa", 1},
		{"aba", 3},
		{"abab", 2},
		{"abaaba", 3},
		{"abaabc", 6},
	}
	for _, d := range testData {
		got := period(d.s)
		if got != d.p {
			t.Errorf("period(%s) got %d, want %d",
				d.s, got, d.p)
		}
	}
}
