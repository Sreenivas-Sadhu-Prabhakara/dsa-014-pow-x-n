package main

import (
	"math"
	"testing"
)

func almost(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestPow(t *testing.T) {
	cases := []struct {
		x    float64
		n    int
		want float64
	}{
		{2, 10, 1024},
		{2, -2, 0.25},
		{5, 0, 1},
		{-2, 3, -8},
		{-2, 4, 16},
		{1, 1 << 40, 1},
		{0.5, 3, 0.125},
	}
	for _, c := range cases {
		if got := Pow(c.x, c.n); !almost(got, c.want) {
			t.Errorf("Pow(%v,%d) = %v, want %v", c.x, c.n, got, c.want)
		}
	}
	// Agreement with the stdlib on a non-trivial case.
	if got, want := Pow(1.1, 47), math.Pow(1.1, 47); !almost(got, want) {
		t.Errorf("Pow(1.1,47) = %v, want %v", got, want)
	}
}
