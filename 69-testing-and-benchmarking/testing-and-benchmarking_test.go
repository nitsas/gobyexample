package main

import (
	"fmt"
	"testing"
)

// Run the tests with `go test` or `go test -v`.

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d. Expected %d.", ans, -2)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b     int
		expected int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{3, 3, 3},
		{2, -2, -2},
		{-1, 0, -1},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("IntMin(%d,%d)", test.a, test.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(test.a, test.b)
			if ans != test.expected {
				t.Errorf("Got %d. Expected %d.", ans, test.expected)
			}
		})
	}
}

// Run this with: `go test -bench=.`
// Option `-bench=` can filter benchmark function names with a regex.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
